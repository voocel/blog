import React from 'react';

interface SEOProps {
    title: string;
    description?: string;
    image?: string;
    type?: 'website' | 'article';
    url?: string;
}

/**
 * SEO component using React 19 native document metadata support.
 * React 19 automatically hoists <title>, <meta>, <link> to <head>.
 */
const SEO: React.FC<SEOProps> = ({
    title,
    description = "A digital sanctuary for thoughts, aesthetics, and the silent rhythm of algorithms.",
    image,
    type = 'website',
    url
}) => {
    const siteTitle = "Voocel Journal";
    const fullTitle = title === siteTitle ? title : `${title} | ${siteTitle}`;

    return (
        <>
            {/* Standard Metadata - React 19 hoists these to <head> */}
            <title>{fullTitle}</title>
            <meta name="description" content={description} />

            {/* Open Graph / Facebook */}
            <meta property="og:type" content={type} />
            <meta property="og:title" content={fullTitle} />
            <meta property="og:description" content={description} />
            {image && <meta property="og:image" content={image} />}
            {url && <meta property="og:url" content={url} />}

            {/* Twitter */}
            <meta name="twitter:card" content="summary_large_image" />
            <meta name="twitter:title" content={fullTitle} />
            <meta name="twitter:description" content={description} />
            {image && <meta name="twitter:image" content={image} />}
        </>
    );
};

export default SEO;
