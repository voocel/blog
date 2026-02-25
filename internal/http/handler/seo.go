package handler

import (
	"encoding/json"
	"html"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"blog/config"
	"blog/internal/entity"
	"blog/internal/usecase"
	"blog/pkg/log"
	"blog/pkg/markdown"

	"github.com/gin-gonic/gin"
)

// fallbackHTML is a minimal HTML page served when the frontend index.html is not available.
const fallbackHTML = `<!doctype html>
<html lang="en">
<head>
<meta charset="UTF-8"/>
<meta name="viewport" content="width=device-width,initial-scale=1.0"/>
<title>Voocel Journal</title>
</head>
<body>
<div id="root"></div>
</body>
</html>`

// pageMeta holds the dynamic meta values to inject into the HTML.
type pageMeta struct {
	Title       string
	Description string
	OGType      string // "website" or "article"
	OGImage     string
	URL         string
	JSONLD      string // JSON-LD script content
	Noscript    string // HTML content for crawlers
}

// SEOHandler serves frontend pages with dynamically injected meta tags,
// enabling search engine crawlers to index content without executing JavaScript.
type SEOHandler struct {
	postUseCase   *usecase.PostUseCase
	indexHTML     atomic.Value // stores string
	indexHTMLPath string
	siteURL      string
}

func NewSEOHandler(postUseCase *usecase.PostUseCase, indexHTMLPath string) *SEOHandler {
	h := &SEOHandler{
		postUseCase:   postUseCase,
		indexHTMLPath: indexHTMLPath,
		siteURL:       strings.TrimRight(config.GetConf().App.SiteURL, "/"),
	}
	h.loadTemplate()
	return h
}

func (h *SEOHandler) loadTemplate() {
	data, err := os.ReadFile(h.indexHTMLPath)
	if err != nil {
		log.Warnw("SEO: index.html not found, will retry on first request",
			log.Pair("path", h.indexHTMLPath),
			log.Pair("error", err.Error()),
		)
		return
	}
	h.indexHTML.Store(string(data))
}

func (h *SEOHandler) getTemplate() string {
	if v := h.indexHTML.Load(); v != nil {
		return v.(string)
	}
	// Lazy retry: the frontend container may not have started yet on first deploy.
	data, err := os.ReadFile(h.indexHTMLPath)
	if err != nil {
		return fallbackHTML
	}
	h.indexHTML.Store(string(data))
	return string(data)
}

// ServeHome handles GET /
func (h *SEOHandler) ServeHome(c *gin.Context) {
	h.servePage(c, pageMeta{
		Title:       "Voocel Journal",
		Description: "Voocel's personal blog exploring technology, design, and life.",
		OGType:      "website",
		URL:         h.siteURL,
	})
}

// ServePosts handles GET /posts
func (h *SEOHandler) ServePosts(c *gin.Context) {
	h.servePage(c, pageMeta{
		Title:       "Posts | Voocel Journal",
		Description: "All articles on technology, design, AI and creative development.",
		OGType:      "website",
		URL:         h.siteURL + "/posts",
	})
}

// ServePost handles GET /post/:slug
func (h *SEOHandler) ServePost(c *gin.Context) {
	slug := c.Param("slug")
	post, err := h.postUseCase.GetBySlug(c.Request.Context(), slug)
	if err != nil || post == nil || post.Status != "published" {
		// Post not found or unpublished: serve default page, let SPA handle 404.
		h.ServeFallback(c)
		return
	}

	ogImage := post.Cover
	if ogImage != "" && !strings.HasPrefix(ogImage, "http") {
		ogImage = h.siteURL + ogImage
	}

	meta := pageMeta{
		Title:       post.Title + " | Voocel Journal",
		Description: post.Excerpt,
		OGType:      "article",
		OGImage:     ogImage,
		URL:         h.siteURL + "/post/" + post.Slug,
		JSONLD:      h.buildArticleJSONLD(post),
		Noscript:    "<h1>" + html.EscapeString(post.Title) + "</h1>" + markdown.ToHTML(post.Content),
	}
	h.servePage(c, meta)
}

// ServeAbout handles GET /about
func (h *SEOHandler) ServeAbout(c *gin.Context) {
	h.servePage(c, pageMeta{
		Title:       "About | Voocel Journal",
		Description: "About Voocel - a developer exploring technology, open source, and creative engineering.",
		OGType:      "website",
		URL:         h.siteURL + "/about",
	})
}

// ServeFallback serves the unmodified index.html with default meta for any unmatched route.
func (h *SEOHandler) ServeFallback(c *gin.Context) {
	h.servePage(c, pageMeta{
		Title:       "Voocel Journal",
		Description: "Voocel's personal blog exploring technology, design, and life.",
		OGType:      "website",
		URL:         h.siteURL + c.Request.URL.Path,
	})
}

// servePage injects meta tags into the HTML template and sends the response.
func (h *SEOHandler) servePage(c *gin.Context, meta pageMeta) {
	tpl := h.getTemplate()

	// Remove existing static meta that will be replaced.
	tpl = removeMeta(tpl, `<meta name="description"`)
	tpl = removeMeta(tpl, `<meta property="og:title"`)
	tpl = removeMeta(tpl, `<meta property="og:description"`)
	tpl = removeMeta(tpl, `<meta property="og:type"`)
	tpl = removeMeta(tpl, `<meta property="og:image"`)
	tpl = removeMeta(tpl, `<meta property="og:url"`)

	// Replace <title> content.
	if meta.Title != "" {
		tpl = replaceTitle(tpl, meta.Title)
	}

	// Build injection block before </head>.
	var inject strings.Builder
	inject.WriteString(`<meta name="description" content="` + html.EscapeString(meta.Description) + `"/>` + "\n")
	inject.WriteString(`<meta property="og:title" content="` + html.EscapeString(meta.Title) + `"/>` + "\n")
	inject.WriteString(`<meta property="og:description" content="` + html.EscapeString(meta.Description) + `"/>` + "\n")
	inject.WriteString(`<meta property="og:type" content="` + meta.OGType + `"/>` + "\n")
	if meta.OGImage != "" {
		inject.WriteString(`<meta property="og:image" content="` + html.EscapeString(meta.OGImage) + `"/>` + "\n")
	}
	if meta.URL != "" {
		inject.WriteString(`<meta property="og:url" content="` + html.EscapeString(meta.URL) + `"/>` + "\n")
		inject.WriteString(`<link rel="canonical" href="` + html.EscapeString(meta.URL) + `"/>` + "\n")
	}
	inject.WriteString(`<meta name="twitter:card" content="summary_large_image"/>` + "\n")
	inject.WriteString(`<meta name="twitter:title" content="` + html.EscapeString(meta.Title) + `"/>` + "\n")
	inject.WriteString(`<meta name="twitter:description" content="` + html.EscapeString(meta.Description) + `"/>` + "\n")
	if meta.OGImage != "" {
		inject.WriteString(`<meta name="twitter:image" content="` + html.EscapeString(meta.OGImage) + `"/>` + "\n")
	}
	if meta.JSONLD != "" {
		inject.WriteString(`<script type="application/ld+json">` + meta.JSONLD + `</script>` + "\n")
	}

	tpl = strings.Replace(tpl, "</head>", inject.String()+"</head>", 1)

	// Inject noscript content after <div id="root"></div>.
	if meta.Noscript != "" {
		noscript := "\n<noscript><article>" + meta.Noscript + "</article></noscript>"
		tpl = strings.Replace(tpl, `<div id="root"></div>`, `<div id="root"></div>`+noscript, 1)
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, tpl)
}

// buildArticleJSONLD generates JSON-LD structured data for a blog post.
func (h *SEOHandler) buildArticleJSONLD(post *entity.PostResponse) string {
	data := map[string]any{
		"@context":    "https://schema.org",
		"@type":       "BlogPosting",
		"headline":    post.Title,
		"description": post.Excerpt,
		"author": map[string]string{
			"@type": "Person",
			"name":  post.Author,
		},
		"datePublished": post.PublishAt.Format(time.RFC3339),
		"url":           h.siteURL + "/post/" + post.Slug,
		"mainEntityOfPage": map[string]string{
			"@type": "WebPage",
			"@id":   h.siteURL + "/post/" + post.Slug,
		},
	}
	if post.Cover != "" {
		cover := post.Cover
		if !strings.HasPrefix(cover, "http") {
			cover = h.siteURL + cover
		}
		data["image"] = cover
	}
	b, _ := json.Marshal(data)
	return string(b)
}

// removeMeta removes a <meta> tag line that starts with the given prefix.
func removeMeta(html, prefix string) string {
	idx := strings.Index(html, prefix)
	if idx < 0 {
		return html
	}
	end := strings.Index(html[idx:], "/>")
	if end < 0 {
		end = strings.Index(html[idx:], ">")
		if end < 0 {
			return html
		}
		end++ // include >
	} else {
		end += 2 // include />
	}
	// Also remove trailing newline if present.
	absEnd := idx + end
	if absEnd < len(html) && html[absEnd] == '\n' {
		absEnd++
	}
	return html[:idx] + html[absEnd:]
}

// replaceTitle replaces the content of the <title> tag.
func replaceTitle(h, newTitle string) string {
	start := strings.Index(h, "<title>")
	if start < 0 {
		return h
	}
	end := strings.Index(h[start:], "</title>")
	if end < 0 {
		return h
	}
	return h[:start] + "<title>" + html.EscapeString(newTitle) + "</title>" + h[start+end+len("</title>"):]
}
