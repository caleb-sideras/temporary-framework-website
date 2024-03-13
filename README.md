<p align="center">
  <a href="https://temporary-framework.org">
    <picture>
      <img src="/static/assets/temporary.png" height="128">
    </picture>
    <h1 align="center">Temporary Website</h1>
  </a>
</p>

<p align="center">
  <a aria-label="Temporary logo" href="https://temporary-framework.org">
    <img src="https://img.shields.io/badge/MADE%20BY%20Caleb%20Sideras-000000.svg?style=for-the-badge&logo=Go&labelColor=000">
  </a>
  <a aria-label="License">
    <img alt="" src="https://img.shields.io/npm/l/next.svg?style=for-the-badge&labelColor=000000">
  </a>
</p>

## What is Temporary?

Temporary is a framework designed to make working with HTMX and Templ easier

## Getting Started

Visit <a href="https://temporary-framework.org">https://temporary-framework.org</a> to get started with Temporary.

## Documentation

Visit [https://temporary-framework.org/docs](https://temporary-framework.org/docs) to view the full documentation.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- Bun: [Download and Install Bun](https://bun.sh/)
- Go Programming Language: [Download and Install Go](https://golang.org/doc/install)
- Templ: [Download and Install Templ](https://templ.guide/quick-start/installation)


## Installing

1. Clone the repository to your local machine.

   ```bash
   git clone --recursive https://github.com/caleb-sideras/temporary-framework-website.git
   ```

2. Navigate to the project's directory.

   ```bash
   cd temporary-framework-website
   ```

3. Install the project dependencies.

   ```bash
   bun install
   ```

## Building the Application

Run the Temporary build process through Bun:

  ```bash
  bun run build
  ```

## Running the Application

Run the Temporary start process through Bun:

   ```bash
   bun run start
   ```

## License

This project is licensed under the MIT License.

## TODO

1. Example routes are always returning the whole page
    - navigating between dynamic routes and regular routes throws off temporary -> returning whole page
    - why are we using the path passed in from request object opposed to the path we have from searching the dir?

2. Delete/create default files on BUILD - > html/css/js delete 
    - so we dont want to just delete all the files in this dir - simply because users might add shit to this?
    - store files we saved so we can remove them?
    - give a warning to the user about storing files here? 

3. Create code for the handlers instead of the definitions.go??
    - could still have the same process but the routes would be stored in separate package?

4. Fix highlighting of focused color

5. Fix mobile titles length

6. dependency injection (it works?)

- Idea 1

Use of closures

NOTE
ok full rewrite needed (dont be consumed by it caleb, still chess, leet and jobs) - thanks!
but we want static routes/pages/indexs to ALL take a w, r, dep objects
so for indexs, they need to take the objects and pass it too their children (if it has)

1. user defines struct
2. user creates object of struct type
3. user passes object into `NewTemp` function
4. Temporary finds type of this object

```go
type MyStruct struct {
	Field string
}

func PrintType(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println("Type of v:", t)
}

func main() {
	myStructInstance := MyStruct{Field: "Hello, Go!"}
	PrintType(myStructInstance)  


// Type of v: main.MyStruct
```

5. Temporary (while traversing the dirs) finds functions that export a http.HandlerFunc vs templ.Component
6. If there is a param of this type, it passes this object into the handler func
NOTE: ok but it looks like we still need to checl the params of the closure function?

```go
package handlers

func Page(w http.ResponseWriter, r *http.Request) templ.Component{
		return Component()
}

func Page(userStruct UserStruct) templ.Component{
		return Component()
}

func Page(w http.ResponseWriter, r *http.Request, userStruct UserStruct) templ.Component{
		return Component()
}

func Page() templ.Component{
		return Component()
}


or

func Page(w http.ResponseWriter, r *http.Request) templ.Component {
}


func main() {
    db, _ := sql.Open("your_db_driver", "your_db_datasource")

    struct UserStruct{
      db *sql.DB
    }

    userStruct := UserStruct{
      db,
      other
    }
    
		t := temp.NewTemp(userStruct)

    func NewTemp(generic T) *Temp {
    	return &Temp{generic T}
    }

    // handled by temporary
    http.HandleFunc("/your-route", handlers.Page(userStruct))
    http.HandleFunc("/your-route", handlers.Page())
    http.ListenAndServe(":8080", nil)
}
```

7. remove mux handler

8. render paths to have handletype - tick

9. metadata api

10. more route level ui

11. nested indexs

12. chunking of js

## REGEX double example

```regex
regex={`^\/projects\/aproject(?:$|\/.*)`}
```
