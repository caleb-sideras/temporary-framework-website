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
			Href:        "/docs/routing/pages-and-index",
			Title:       "Server Components",
			Description: "Add Templ Components without the need to handle rendering or routing.",
			HxBoost:     true,
		},
		{
			Href:        "/docs/routing/dynamic-routes",
			Title:       "Dynamic Routes",
			Description: "Create routes from dynamic data known at request-time.",
			HxBoost:     true,
		},
		{
			Href:        "/docs/routing/suspense",
			Title:       "Lazy Loading",
			Description: "Instantly load UI from the server and lazy load your content.",
			HxBoost:     true,
		},
		{
			Href:        "/docs/routing/caching",
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
