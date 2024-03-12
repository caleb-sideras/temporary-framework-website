package main

import (
	temp "calebsideras.com/temporary/temporary"
	"fmt"
	"net/http"
	"os"
	// "reflect"
	"github.com/gorilla/mux"
)

// func test(param interface{}) {
// 	if param == nil {
// 		fmt.Println("We need some generic type")
// 		return
// 	}
// 	t := reflect.TypeOf(param)
// 	fmt.Println("Type of v:", t)
// 	fmt.Println("Package Path:", t.PkgPath())
// }

func main() {
	// t := temp.Temp{}
	// test(nil)
	// test(temporary)
	type hello struct {
		name string
	}

	if len(os.Args) > 1 {

		t := temp.NewTemp(hello{"bossman"})

		switch os.Args[1] {
		case "build":
			// t.Build()

		case "render":
			t.Render()

		case "run":

			http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

			t.Run(mux.NewRouter(), ":8000")

		default:
			fmt.Println("Invalid argument. Use 'build' or 'run'.")
		}
	} else {
		fmt.Println("Please provide an argument: 'build' or 'run'.")
	}
}
