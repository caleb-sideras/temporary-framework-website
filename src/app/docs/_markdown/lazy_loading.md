# Lazy Loading

Lazy loading in Temporary helps improve the initial loading performance of an application by decreasing the amount of JavaScript needed to be loaded with a route.

It allows you to defer loading of imported libraries and custom Javascript, and only include them in the client when they're needed.

Using Temporary's [metadata API](/docs/optimizing/metadata), you can include script tags in your Index if Javascript is shared between mutiple pages, or directly in your Pages if Javascript is isolated to a single page. 

Consider the following files inside the app directory:

```
app                  
 ├─ index.go       
 ├─ 1                
 │  └─ page.go  
 └─ 2
    └─ page.go  
```
### index.go

```go
var Metadata []string = []string{
	`<script src="/path/to/shared.js"></script>`,
}

func Index(w http.ResponseWriter, r *http.Request) templ.Component {
	return IndexTempl()
}
}
```

### 1/page.go

```go
func Page() templ.Component {
  return Component()
}
```

### 2/page.go

```go
var Metadata []string = []string{
	`<script src="/path/to/isolated.js"></script>`,
}

func Page() templ.Component {
  return Component()
}
```

Based on the above examples, `<script src="/path/to/shared.js"></script>` will be loaded for both `1/page.go` and `2/page.go`, whereas `<script src="/path/to/isolated.js"></script>` is only loaded for `2/page.go`.
