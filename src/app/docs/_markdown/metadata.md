# Metadata

Temporary has a Metadata API that can be used to define your application metadata (e.g. `meta` and `link` tags inside your HTML `head` element) for improved SEO and web shareability.

Currently Temporary only supports static metadata with plans to add dynamic metadata very soon.

- Static Metadata: Export a variable named `Metadata` of type `[]string` in a `index.go` or `page.go` file, where each `string` element in this list correlates to a metadata tag.

Temporary will automatically generate the relevant `<head>` elements for your pages. 

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
	`<meta name="description" content="Metadata Example">`,
}

func Index(w http.ResponseWriter, r *http.Request) templ.Component {
	return IndexTempl()
}

templ IndexTempl(){
	<!DOCTYPE html>
	<html lang="en">
    <head>
  		<meta charset="UTF-8" hx-preserve="true"/>
  		<meta http-equiv="X-UA-Compatible" content="IE=edge" hx-preserve="true"/>
  		<meta name="viewport" content="width=device-width, initial-scale=1.0" hx-preserve="true"/>
    </head>
    <body>
      <main>
        { children... }
      </main>
    </body>
  </html>
}
```

### 1/page.go

```go
var Metadata []string = []string{
	"<title> 1 | Page </title>",
	`<script src="/path/to.js"></script>`,
}

func Page() templ.Component {
  return Component()
}
```

### 2/page.go

```go
var Metadata []string = []string{
	"<title> 2 | Page </title>",
}

func Page() templ.Component {
  return Component()
}
```

## Full Page Responses

Navigating individually to the following paths will resolve to:

1. __/1__ 

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta name="description" content="Metadata Example">

		<meta charset="UTF-8" hx-preserve="true"/>
		<meta http-equiv="X-UA-Compatible" content="IE=edge" hx-preserve="true"/>
		<meta name="viewport" content="width=device-width, initial-scale=1.0" hx-preserve="true"/>

  	<title> 1 | Page </title>
  	<script src="/path/to.js"></script>
  </head>
  <body>
    <main>
      <!--  1/page.go output -->
    </main>
  </body>
</html>
```

2. __/2__

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta name="description" content="Metadata Example">

		<meta charset="UTF-8" hx-preserve="true"/>
		<meta http-equiv="X-UA-Compatible" content="IE=edge" hx-preserve="true"/>
		<meta name="viewport" content="width=device-width, initial-scale=1.0" hx-preserve="true"/>

  	<title> 2 | Page </title>
  </head>
  <body>
    <main>
      <!-- 2/page.go output -->
    </main>
  </body>
</html>
```

## Partial Responses

Due to Temporary's [hx-boost](/docs/routing/hx-boost) override, navigation with hx-boost between Pages that share the same Index will return HTML partials. 

If you navigate to path `/2` from path `/1`, the following will be returned:

```html
<head>
  <meta name="description" content="Metadata Example">
	<title> 2 | Page </title>
</head>
<body>
  <main>
    <!-- 2/page.go output -->
  </main>
</body>

```

Temporary uses the [head-support](https://htmx.org/extensions/head-support/) htmx extension to merge the response into the DOM. All conventions described in the head-support documentation can be applied here.
