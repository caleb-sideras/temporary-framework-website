package metadata

import (
	// "net/http"

	// "calebsideras.com/temporary/src/components/server"
	"github.com/a-h/templ"
)

// not safe html. would be better if separate html file, json or templ
var Metadata []string = []string{
	`<link rel="author" href="https://calebsideras.com">`,
	`<meta name="author" content="Caleb">`,
	`<meta name="keywords" content="MusicGPT,Music,AI">`,
}

// ???
// func Metadata_() templ.Component {
// 	// fetch data
// 	return server.Todo()
// }

// func Metadata_() []string {
// 	// fetch data
// 	return []string{
// 		`<link rel="author" href="https://calebsideras.com">`,
// 		`<meta name="author" content="Caleb">`,
// 		`<meta name="keywords" content="MusicGPT,Music,AI">`,
// 	}
// }

// func Metadata(w http.ResponseWriter, r *http.Request) []string {
// 	// fetch data
// 	return []string{
// 		`<link rel="author" href="https://calebsideras.com">`,
// 		`<meta name="author" content="Caleb">`,
// 		`<meta name="keywords" content="MusicGPT,Music,AI">`,
// 	}
// }

func Page() templ.Component {
	return PageTempl()
}
