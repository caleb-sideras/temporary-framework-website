package home

import (
	"github.com/a-h/templ"
)

type TLink2 struct {
	Title       string
	Description string
	Href        string
	HxBoost     bool
}

var Metadata []string = []string{
	// HTML Tags
	`<title>Temporary</title>`,
	`<meta name="description" content="The HTMX and Templ Framework you are looking for."/>`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org"/>`,
	`<meta property="og:type" content="website"/>`,
	`<meta property="og:title" content="Temporary"/>`,
	`<meta property="og:description" content="The HTMX and Templ Framework you are looking for."/>`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png"/>`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image"/>`,
	`<meta property="twitter:domain" content="temporary-framework.org"/>`,
	`<meta property="twitter:url" content="https://temporary-framework.org"/>`,
	`<meta name="twitter:title" content="Temporary"/>`,
	`<meta name="twitter:description" content="The HTMX and Templ Framework you are looking for."/>`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png"/>`,
}

func Page_() templ.Component {

	links := []TLink2{
		{
			Href:        "/docs/routing",
			Title:       "File-Based Routing",
			Description: "Temporary uses a file-system based router where folders are used to define routes.",
			HxBoost:     true,
		},
		{
			Href:        "/docs/routing/defining-routes",
			Title:       "Route Handling",
			Description: "Route specific UI that perform unique actions.",
			HxBoost:     true,
		},
		{
			Href:        "/docs/optimizing/rendering",
			Title:       "Server Components",
			Description: "Add Templ Components without the need to handle a rendering pipeline or routing.",
			HxBoost:     true,
		},
		{
			Href:        "/docs/routing/dynamic-routes",
			Title:       "Dynamic Routes",
			Description: "Create routes from dynamic data known at request-time.",
			HxBoost:     true,
		},
		{
			Href:        "/docs/optimizing/suspense",
			Title:       "Lazy Loading",
			Description: "Instantly load UI from the server and lazy load your content.",
			HxBoost:     true,
		},
		{
			Href:        "/docs/optimizing/caching",
			Title:       "Caching",
			Description: "Never have redundant or stale HTML sent to the client.",
			HxBoost:     true,
		},
		{
			Href:        "/docs/routing/hx-boost",
			Title:       "hx-boost",
			Description: "hx-boost overide to provide an even better SPA feel.",
			HxBoost:     true,
		},
		{
			Href:        "/docs/optimizing/metadata",
			Title:       "hx-boost",
			Description: "Dynamically set metadata.",
			HxBoost:     true,
		},
		{
			Href:        "/docs/routing/dependency-injection",
			Title:       "Dependency Injection",
			Description: "Inject dependencies that are accessible in your handlers.",
			HxBoost:     true,
		},
	}

	links2 := []TLink2{
		{
			Title:       "temporary-starter",
			Description: "Temporary is an open-source project and far from production ready. Feel free to contribute!",
			Href:        "https://github.com/caleb-sideras/temporary-framework-starter",
			HxBoost:     false,
		},
		{
			Title:       "temporary-website",
			Description: "The code for this current website is public.",
			Href:        "https://github.com/caleb-sideras/temporary-framework-website",
			HxBoost:     false,
		},
	}

	return HomeTempl(links, links2)
}
