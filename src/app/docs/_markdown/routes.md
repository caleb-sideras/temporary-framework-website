# Routes

A route is UI that is unique to a route. You can define a route by exporting a function of return type `templ.Component` from a `route.go` file. The names of the exported functions defined in this file will resolve to the URL, paired with the current directory. 

Create your first route by adding a `route.go` file inside the app directory:

```
app                  (1)
 ├─ index.go       
 └─ example          (2)
    └─ route.go  
```

```go

package example

import (
	"net/http"
	"github.com/a-h/templ"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) templ.Component {
	return header1("Hello World")
}

func Hello() templ.Component {
	return header1("Hello")
}

templ header1(text string){
  <h1>{ text }</h1>
}
  
```

These resolve to:

1. __example/hello-world__ ← `func HelloWorld()`

2. __example/hello__ ← `func Hello()`


### Good to know:

- Camel case function names are converted to Hyphen-separated case URLs. 
- Routes are Dynamic Server Components by default.
- Routes can fetch data. 
- Routes should return a `Templ.component`.
- Routes are NOT parsed with Indexs. 
