
// Code generated by Temporary; DO NOT EDIT.
package temporary
import (
	"calebsideras.com/temporary/src/app/docs/routing/pages-and-index"
	"calebsideras.com/temporary/src/app/docs/pages/static-pages"
	"calebsideras.com/temporary/src/app/home_"
	"calebsideras.com/temporary/src/app/docs/routing/defining-routes"
	"calebsideras.com/temporary/src/app/docs/routing/dynamic-routes"
	"calebsideras.com/temporary/src/app/docs/routing/hx-boost"
	"calebsideras.com/temporary/src/app/blog"
	"calebsideras.com/temporary/src/app/examples/suspense"
	"calebsideras.com/temporary/src/app/blog/temporarystandard"
	"calebsideras.com/temporary/src/app/docs/pages"
	"calebsideras.com/temporary/src/app/docs/routing"
	"calebsideras.com/temporary/src/app/docs/index/templ"
	"calebsideras.com/temporary/src/app/docs/routing/static-pages-and-routes"
	"calebsideras.com/temporary/src/app/docs/index"
	"calebsideras.com/temporary/src/app/docs/pages/templ"
	"calebsideras.com/temporary/src/app/examples/todo"
	"calebsideras.com/temporary/src/app/examples/static-render"
	"calebsideras.com/temporary/src/app/docs/routing/suspense"
	"calebsideras.com/temporary/src/app/docs/introduction_"
	"calebsideras.com/temporary/src/app"
	"calebsideras.com/temporary/src/app/docs/project-structure"
	"calebsideras.com/temporary/src/app/examples/_slug_/dynamic-routes"
	"calebsideras.com/temporary/src/app/docs/routing/caching"
)

var IndexList = map[string]IndexDefaultFunc{
	"/docs/routing" : app.Index,
	"/examples/{slug}" : app.Index,
	"/docs/routing/pages-and-index" : app.Index,
	"/examples/suspense" : app.Index,
	"/docs/routing/hx-boost" : app.Index,
	"/docs/index" : app.Index,
	"/docs/routing/suspense" : app.Index,
	"/docs/routing/dynamic-routes" : app.Index,
	"/blog/temporarystandard" : app.Index,
	"/docs/pages" : app.Index,
	"/docs/routing/defining-routes" : app.Index,
	"/docs/routing/caching" : app.Index,
	"/docs/routing/static-pages-and-routes" : app.Index,
	"/docs/project-structure" : app.Index,
	"/examples/static-render" : app.Index,
	"/docs/pages/templ" : app.Index,
	"/examples" : app.Index,
	"/examples/{slug}/dynamic-routes" : app.Index,
	"/examples/todo" : app.Index,
	"/" : app.Index,
	"/docs" : app.Index,
	"/blog" : app.Index,
	"/docs/pages/static-pages" : app.Index,
	"/docs/index/templ" : app.Index,
}

var PageRenderList = []RenderDefault{
	{"/blog", blog.Page_},
}

var RouteRenderList = []RenderDefault{
	
}

var PageHandleList = []Handler{
	{"/docs/routing/static-pages-and-routes", docs_routing_staticpagesandroutes.Page, DefaultHandler},
	{"/docs/index", docs_index.Page, DefaultHandler},
	{"/docs/project-structure", docs_projectstructure.Page, DefaultHandler},
	{"/docs/routing/pages-and-index", docs_routing_pagesandindex.Page, DefaultHandler},
	{"/docs/pages/static-pages", docs_pages_staticpages.Page, DefaultHandler},
	{"/docs/pages/templ", docs_pages_templ.Page, DefaultHandler},
	{"/examples/{slug}/dynamic-routes", examples_dynamicroutes.Page, ResReqHandler},
	{"/examples/todo", examples_todo.Page, ResReqHandler},
	{"/examples/static-render", examples_staticrender.Page, DefaultHandler},
	{"/examples/suspense", examples_suspense.Page, ResReqHandler},
	{"/", home.Page, ResReqHandler},
	{"/docs/routing/suspense", docs_routing_suspense.Page, DefaultHandler},
	{"/blog/temporarystandard", temporarystandard.Page, ResReqHandler},
	{"/docs/pages", docs_pages.Page, DefaultHandler},
	{"/docs", docs_introduction.Page, DefaultHandler},
	{"/docs/routing/defining-routes", docs_routingfundamentals.Page, DefaultHandler},
	{"/docs/routing/dynamic-routes", docs_routing_dynamicrouting.Page, DefaultHandler},
	{"/docs/routing/hx-boost", docs_routing_hxboost.Page, DefaultHandler},
	{"/docs/index/templ", docs_index_templ.Page, DefaultHandler},
	{"/docs/routing", docs_routing.Page, DefaultHandler},
	{"/docs/routing/caching", docs_routing_caching.Page, DefaultHandler},
}

var RouteHandleList = []Handler{
	{"/examples/{slug}/dynamic-routes/example", examples_dynamicroutes.Example, ResReqHandler},
	{"/examples/{slug}/dynamic-routes/code", examples_dynamicroutes.Code, DefaultHandler},
	{"/examples/todo/example", examples_todo.Example, DefaultHandler},
	{"/examples/todo/addtask", examples_todo.AddTask, ResReqHandler},
	{"/examples/todo/code", examples_todo.Code, DefaultHandler},
	{"/examples/static-render/example", examples_staticrender.Example, DefaultHandler},
	{"/examples/static-render/code", examples_staticrender.Code, DefaultHandler},
	{"/examples/suspense/example", examples_suspense.Example, DefaultHandler},
	{"/examples/suspense/code", examples_suspense.Code, DefaultHandler},
}
