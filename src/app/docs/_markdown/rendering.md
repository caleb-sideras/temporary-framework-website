# Static Rendering

By __default__, Temporary treats all routes defined in `indexs.go`, `pages.go` and `route.go` as dynamic. This approach ensures that your functions are executed with each request, allowing all operations (like data fetching and mutations) to occur in real-time, thereby guaranteeing the most current data.

The following example illustrates a typical dynamic route scenario:


```go
package home

import (
	"net/http"
	"github.com/a-h/templ"
)

func Page(w http.ResponseWriter, r *http.Request) templ.Component {
	someData := fetchNames()

	return NamesList(someData)
}

templ NamesList(names []string){
  <ul>
		for _, n := range names {
			<li>
				{ n }
			</li>
		}
	</ul>
  }  
```

## Static

However, not all routes require real-time data processing. Some serve consistent HTML content, where executing extensive build processes for each request can hamper performance. For these scenarios, Temporary offers a build-time rendering feature. By simply appending an underscore `_` to your function names in your  `index.go`, `page.go` and `route.go`, you can render these routes at build-time. This not only generates the full HTML pages but also prepares the partial content necessary for Temporary's `hx-boost` override, optimizing performance for static content.

Below is an example of a route with an extensive build process, such as converting markdown to a Templ component. This route is used for the page you are currently viewing.

```go
package home

import (
	"github.com/a-h/templ"
	"calebsideras.com/temporary/src/utils"
)

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("path/rendering.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
  
```

### Parameters

Additionally, static rendering does not require you to omit any Temporary supported parameters from your functions. The following code below will build.

```go
package home

import (
	"github.com/a-h/templ"
	"calebsideras.com/temporary/src/utils"
)

func Page_(w http.ResponseWriter, r *http.Request, d dep) templ.Component {

	newTempl, err := utils.MdFileToTempl("path/rendering.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
  
```
__NOTE:__ A dummy request and object are passed in to statisfy the `*http.Request` and the `http.ResponseWriter` parameters. Using them in your statically generated functions may cause errors. Alternatively, your custom dependency type can be used without issue.


## Nested Rendering

Temporary optimizes performance by explicitly handling various types of nested rendering for static content. 

Consider the following project layout:

```
app                  (1)
 ├─ index.go       
 └─ example          (2)
    └─ page.go
```

### Static Index and Dynamic Page [`Index_()` -> `Page()`] 

The Index `Index_()` is pre-rendered at build time. Temporary dynamically renders the Page `Page()` upon each request and parses its output into the pre-rendered Index HTML.

Ideal for scenarios where the Index structure remains consistent, but the Page content is dynamic or user-specific.

### Dynamic Index and Static Page [`Index()` -> `Page()_`]

The Page `Page_()` is pre-rendered at build time. Temporary then dynamically renders the Index `Index()` for each request and parses the pre-rendered Page HTML into it.

Suitable when the Page content is static, but the Index needs to dynamically display user-specific data or content.

### Static Index and Static Page [`Index_()` -> `Page()_`]

Both the Index `Index_()` and the Page `Page_()` are pre-rendered at build time. Temporary combines these static HTML outputs once and serves the fully static HTML for each request.

Best for entirely static websites where both the Index and Page do not change or only change infrequently.
