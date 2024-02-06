
// Code generated by Temporary; DO NOT EDIT.
package temporary
import (
	"calebsideras.com/temporary/src/app/examples/_slug_/dynamic-routes"
	"calebsideras.com/temporary/src/app/examples/static-render"
	"calebsideras.com/temporary/src/app/home_"
	"calebsideras.com/temporary/src/app/blog"
	"calebsideras.com/temporary/src/app"
	"calebsideras.com/temporary/src/app/examples/todo"
	"calebsideras.com/temporary/src/app/examples/suspense"
)

var IndexList = map[string]IndexDefaultFunc{
	"/docs/index" : app.Index,
	"/docs/pages/static-pages" : app.Index,
	"/docs/routing/static-pages-and-routes" : app.Index,
	"/docs/routing" : app.Index,
	"/docs/routing/defining-routes" : app.Index,
	"/examples/todo" : app.Index,
	"/blog/temporarystandard" : app.Index,
	"/examples/suspense" : app.Index,
	"/docs/project-structure" : app.Index,
	"/examples/{slug}" : app.Index,
	"/blog" : app.Index,
	"/docs/routing/pages-and-index" : app.Index,
	"/examples" : app.Index,
	"/examples/{slug}/dynamic-routes" : app.Index,
	"/docs/routing/caching" : app.Index,
	"/docs/routing/dynamic-routes" : app.Index,
	"/" : app.Index,
	"/examples/static-render" : app.Index,
	"/docs/routing/hx-boost" : app.Index,
	"/docs" : app.Index,
	"/docs/pages/templ" : app.Index,
	"/docs/routing/suspense" : app.Index,
	"/docs/index/templ" : app.Index,
	"/docs/pages" : app.Index,
}

var PageRenderList = []RenderDefault{
	{"/", home.Page_},
	{"/blog", blog.Page_},
}

var RouteRenderList = []RenderDefault{
	
}

var PageHandleList = []Handler{
}

var RouteHandleList = []Handler{
	{"/examples/todo/addtask", examples_todo.AddTask, ResReqHandler},
	{"/examples/todo/code", examples_todo.Code, DefaultHandler},
	{"/examples/todo/example", examples_todo.Example, DefaultHandler},
	{"/examples/suspense/example", examples_suspense.Example, DefaultHandler},
	{"/examples/suspense/code", examples_suspense.Code, DefaultHandler},
	{"/examples/static-render/example", examples_staticrender.Example, DefaultHandler},
	{"/examples/static-render/code", examples_staticrender.Code, DefaultHandler},
	{"/examples/{slug}/dynamic-routes/example", examples_dynamicroutes.Example, ResReqHandler},
	{"/examples/{slug}/dynamic-routes/code", examples_dynamicroutes.Code, DefaultHandler},
}
