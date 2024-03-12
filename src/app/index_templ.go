// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package app

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func IndexTempl() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><script type=\"module\" src=\"/static/js/index.js\"></script><link href=\"/static/css/output.css\" rel=\"stylesheet\"><!-- <script src=\"https://unpkg.com/htmx.org@1.9.6\" integrity=\"sha384-FhXw7b6AlE/jyjlZH5iHa/tTe9EpJ1Y55RjcgPbjeWMskSxZt1v9qkxLJWNJaGni\" crossorigin=\"anonymous\"></script> --><script src=\"https://unpkg.com/htmx.org@2.0.0-alpha1/dist/htmx.min.js\"></script><link href=\"https://fonts.googleapis.com/css2?family=Roboto:wght@400;500;700&amp;display=swap\" rel=\"stylesheet\"><link href=\"https://cdnjs.cloudflare.com/ajax/libs/prism/1.20.0/themes/prism.min.css\" rel=\"stylesheet\"><link rel=\"stylesheet\" href=\"https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200\"><!-- HTML Meta Tags --><title>Temporary</title><meta name=\"description\" content=\"The HTMX and Templ Framework you are looking for.\"><!-- Facebook Meta Tags --><meta property=\"og:url\" content=\"https://temporary-framework.org\"><meta property=\"og:type\" content=\"website\"><meta property=\"og:title\" content=\"Temporary\"><meta property=\"og:description\" content=\"The HTMX and Templ Framework you are looking for.\"><meta property=\"og:image\" content=\"https://temporary-framework.org/static/assets/temporary.png\"><!-- Twitter Meta Tags --><meta name=\"twitter:card\" content=\"summary_large_image\"><meta property=\"twitter:domain\" content=\"temporary-framework.org\"><meta property=\"twitter:url\" content=\"https://temporary-framework.org\"><meta name=\"twitter:title\" content=\"Temporary\"><meta name=\"twitter:description\" content=\"The HTMX and Templ Framework you are looking for.\"><meta name=\"twitter:image\" content=\"https://temporary-framework.org/static/assets/temporary.png\"></head><body class=\"dark\"><header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = MobileNavigation().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Navigation().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</header><div class=\"main_container\"><main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Navigation() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<t-navigation><t-navigation-rail slot=\"rail\"><temporary-list-item interactive type=\"link\" href=\"/\" regex=\"^/?$\" hx-boost=\"true\">Home <md-icon slot=\"start\" class=\"material-symbols-filled\">crop_square</md-icon></temporary-list-item> <temporary-list-item interactive type=\"link\" href=\"/examples/todo\" hx-boost=\"true\" regex=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(`^/?\/examples(\/?|\/.*)?$`))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Examples <md-icon slot=\"start\" class=\"material-symbols-filled\">code_blocks</md-icon></temporary-list-item> <temporary-list-item interactive type=\"link\" href=\"/docs\" regex=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(`^/?\/docs(\/?|\/.*)?$`))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-boost=\"true\">Docs <md-icon slot=\"start\" class=\"material-symbols-filled\">developer_guide</md-icon></temporary-list-item></t-navigation-rail> <t-navigation-drawer slot=\"drawer\"><temporary-list-item interactive border type=\"link\" href=\"/docs\" hx-boost=\"true\">Introduction</temporary-list-item> <temporary-list-item interactive border type=\"link\" href=\"/docs/installation\" hx-boost=\"true\">Installation</temporary-list-item> <temporary-list-item interactive border type=\"link\" href=\"/docs/project-structure\" hx-boost=\"true\">Project Structure</temporary-list-item> <t-dropdown><t-dropdown-title border hide-event slot=\"title\"><div slot=\"headline\">Routing</div><md-icon slot=\"active\">arrow_drop_up</md-icon> <md-icon slot=\"inactive\">arrow_drop_down</md-icon></t-dropdown-title> <t-dropdown-list slot=\"content\"><t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing\" hx-boost=\"true\">Fundamentals</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/defining-routes\" hx-boost=\"true\">Defining Routes</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/pages-and-index\" hx-boost=\"true\">Pages and Index</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/hx-boost\" hx-boost=\"true\">hx-boost</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/static-pages-and-routes\" hx-boost=\"true\">Static Pages</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/caching\" hx-boost=\"true\">Caching</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/dynamic-routes\" hx-boost=\"true\">Dynamic Routes</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/suspense\" hx-boost=\"true\">Suspense</t-dropdown-list-item></t-dropdown-list></t-dropdown> <t-dropdown><t-dropdown-title border hide-event slot=\"title\"><div slot=\"headline\">Pages</div><md-icon slot=\"active\">arrow_drop_up</md-icon> <md-icon slot=\"inactive\">arrow_drop_down</md-icon></t-dropdown-title> <t-dropdown-list slot=\"content\"><t-dropdown-list-item interactive border type=\"link\" href=\"/docs/pages\" hx-boost=\"true\">Definition</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/pages/templ\" hx-boost=\"true\">Templ</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/pages/static-pages\" hx-boost=\"true\">Static Rendering</t-dropdown-list-item></t-dropdown-list></t-dropdown> <t-dropdown><t-dropdown-title border hide-event slot=\"title\"><div slot=\"headline\">Index</div><md-icon slot=\"active\">arrow_drop_up</md-icon> <md-icon slot=\"inactive\">arrow_drop_down</md-icon></t-dropdown-title> <t-dropdown-list slot=\"content\"><t-dropdown-list-item interactive border type=\"link\" href=\"/docs/index\" hx-boost=\"true\">Definition</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/index/templ\" hx-boost=\"true\">Templ</t-dropdown-list-item></t-dropdown-list></t-dropdown></t-navigation-drawer> <t-navigation-drawer slot=\"drawer\"><temporary-list-item interactive border type=\"link\" href=\"/examples/todo\" hx-boost=\"true\">Todo</temporary-list-item> <temporary-list-item interactive border type=\"link\" href=\"/examples/suspense\" hx-boost=\"true\">Suspense</temporary-list-item> <temporary-list-item interactive border type=\"link\" regex=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(`^/examples/.*/dynamic-routes$`))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" href=\"/examples/_replace-me_/dynamic-routes\" hx-boost=\"true\">Dynamic Routing</temporary-list-item> <temporary-list-item interactive border type=\"link\" href=\"/examples/static-render\" hx-boost=\"true\">Static Render</temporary-list-item></t-navigation-drawer><div slot=\"footer\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = NavFooter().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></t-navigation>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func MobileNavigation() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = MobileAppBar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<t-mobile-navigation><t-mobile-navigation-rail slot=\"rail\"><temporary-list-item interactive type=\"link\" href=\"/\" regex=\"^/?$\" hx-boost=\"true\">Home <md-icon slot=\"start\" class=\"material-symbols-filled\">crop_square</md-icon></temporary-list-item> <temporary-list-item interactive type=\"button\" href=\"/examples/todo\" regex=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(`^/?\/examples(\/?|\/.*)?$`))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Examples <md-icon slot=\"start\" class=\"material-symbols-filled\">code_blocks</md-icon> <md-icon slot=\"end\" class=\"material-symbols-filled\">arrow_forward</md-icon></temporary-list-item> <temporary-list-item interactive type=\"button\" href=\"/docs\" regex=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(`^/?\/docs(\/?|\/.*)?$`))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Docs <md-icon slot=\"start\" class=\"material-symbols-filled\">developer_guide</md-icon> <md-icon slot=\"end\" class=\"material-symbols-filled\">arrow_forward</md-icon></temporary-list-item></t-mobile-navigation-rail> <t-mobile-navigation-drawer slot=\"drawer\"><temporary-list-item interactive border type=\"link\" href=\"/docs\" hx-boost=\"true\">Introduction</temporary-list-item> <temporary-list-item interactive border type=\"link\" href=\"/docs/installation\" hx-boost=\"true\">Installation</temporary-list-item> <temporary-list-item interactive border type=\"link\" href=\"/docs/project-structure\" hx-boost=\"true\">Project Structure</temporary-list-item> <t-dropdown><t-dropdown-title border hide-event slot=\"title\"><div slot=\"headline\">Routing</div><md-icon slot=\"active\">arrow_drop_up</md-icon> <md-icon slot=\"inactive\">arrow_drop_down</md-icon></t-dropdown-title> <t-dropdown-list slot=\"content\"><t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing\" hx-boost=\"true\">Fundamentals</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/defining-routes\" hx-boost=\"true\">Defining Routes</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/pages-and-index\" hx-boost=\"true\">Pages and Index</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/hx-boost\" hx-boost=\"true\">hx-boost</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/static-pages-and-routes\" hx-boost=\"true\">Static Pages</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/caching\" hx-boost=\"true\">Caching</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/dynamic-routes\" hx-boost=\"true\">Dynamic Routes</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/routing/suspense\" hx-boost=\"true\">Suspense</t-dropdown-list-item></t-dropdown-list></t-dropdown> <t-dropdown><t-dropdown-title border hide-event slot=\"title\"><div slot=\"headline\">Pages</div><md-icon slot=\"active\">arrow_drop_up</md-icon> <md-icon slot=\"inactive\">arrow_drop_down</md-icon></t-dropdown-title> <t-dropdown-list slot=\"content\"><t-dropdown-list-item interactive border type=\"link\" href=\"/docs/pages\" hx-boost=\"true\">Definition</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/pages/templ\" hx-boost=\"true\">Templ</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/pages/static-pages\" hx-boost=\"true\">Static Rendering</t-dropdown-list-item></t-dropdown-list></t-dropdown> <t-dropdown><t-dropdown-title border hide-event slot=\"title\"><div slot=\"headline\">Index</div><md-icon slot=\"active\">arrow_drop_up</md-icon> <md-icon slot=\"inactive\">arrow_drop_down</md-icon></t-dropdown-title> <t-dropdown-list slot=\"content\"><t-dropdown-list-item interactive border type=\"link\" href=\"/docs/index\" hx-boost=\"true\">Definition</t-dropdown-list-item> <t-dropdown-list-item interactive border type=\"link\" href=\"/docs/index/templ\" hx-boost=\"true\">Templ</t-dropdown-list-item></t-dropdown-list></t-dropdown></t-mobile-navigation-drawer> <t-mobile-navigation-drawer slot=\"drawer\"><temporary-list-item interactive border type=\"link\" href=\"/examples/todo\" hx-boost=\"true\">Todo</temporary-list-item> <temporary-list-item interactive border type=\"link\" href=\"/examples/suspense\" hx-boost=\"true\">Suspense</temporary-list-item> <temporary-list-item interactive border type=\"link\" regex=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(`^/examples/.*/dynamic-routes$`))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" href=\"/examples/_replace-me_/dynamic-routes\" hx-boost=\"true\">Dynamic Routing</temporary-list-item> <temporary-list-item interactive border type=\"link\" href=\"/examples/static-render\" hx-boost=\"true\">Static Render</temporary-list-item></t-mobile-navigation-drawer><div slot=\"footer\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = NavFooter().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></t-mobile-navigation>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func MobileAppBar() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<t-mobile-top-app-bar><md-icon slot=\"navigationIcon\" class=\"material-symbols-filled\">menu</md-icon></t-mobile-top-app-bar>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func NavFooter() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<t-mode-toggle></t-mode-toggle> <temporary-list-item interactive type=\"link\" href=\"https://github.com/caleb-sideras/\">github <md-icon slot=\"end\" class=\"material-symbols-filled-light\">north_east</md-icon></temporary-list-item> <temporary-list-item interactive type=\"link\" href=\"https://www.linkedin.com/in/caleb-james-sideras-9555b0196/\">linkedin <md-icon slot=\"end\" class=\"material-symbols-filled-light\">north_east</md-icon></temporary-list-item>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
