package handler

import (
	"encoding/xml"
	"net/http"
	"time"

	"blog/config"
	"blog/internal/usecase"

	"github.com/gin-gonic/gin"
)

// XML sitemap structures
type urlset struct {
	XMLName xml.Name    `xml:"urlset"`
	Xmlns   string      `xml:"xmlns,attr"`
	URLs    []sitemapURL `xml:"url"`
}

type sitemapURL struct {
	Loc        string  `xml:"loc"`
	Lastmod    string  `xml:"lastmod,omitempty"`
	Changefreq string  `xml:"changefreq,omitempty"`
	Priority   float32 `xml:"priority,omitempty"`
}

type SitemapHandler struct {
	postRepo usecase.PostRepo
}

func NewSitemapHandler(postRepo usecase.PostRepo) *SitemapHandler {
	return &SitemapHandler{postRepo: postRepo}
}

func (h *SitemapHandler) GenerateSitemap(c *gin.Context) {
	ctx := c.Request.Context()
	siteURL := config.GetConf().App.SiteURL

	sitemap := urlset{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}

	// Homepage
	sitemap.URLs = append(sitemap.URLs, sitemapURL{
		Loc:        siteURL,
		Changefreq: "daily",
		Priority:   1.0,
	})

	// Posts list page
	sitemap.URLs = append(sitemap.URLs, sitemapURL{
		Loc:        siteURL + "/posts",
		Changefreq: "daily",
		Priority:   0.9,
	})

	// All published posts
	filters := map[string]any{"status": "published"}
	posts, _, err := h.postRepo.List(ctx, filters, 1, 1000)
	if err == nil {
		for _, post := range posts {
			lastmod := post.UpdatedAt
			if lastmod.IsZero() {
				lastmod = post.PublishAt
			}

			sitemap.URLs = append(sitemap.URLs, sitemapURL{
				Loc:        siteURL + "/post/" + post.Slug,
				Lastmod:    lastmod.Format(time.RFC3339),
				Changefreq: "weekly",
				Priority:   0.8,
			})
		}
	}

	c.Header("Content-Type", "application/xml; charset=utf-8")
	c.XML(http.StatusOK, sitemap)
}

func (h *SitemapHandler) RobotsTxt(c *gin.Context) {
	siteURL := config.GetConf().App.SiteURL

	robots := `User-agent: *
Allow: /

Sitemap: ` + siteURL + `/sitemap.xml
`
	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.String(http.StatusOK, robots)
}
