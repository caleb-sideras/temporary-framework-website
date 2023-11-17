package home

import (
	"github.com/a-h/templ"
	components "github.com/caleb-sideras/gox2/src/app/home_/components"
)

// Server Component
func Page_() templ.Component {
	// perform server-side logic -> fetch data, mutations, etc

	return components.HomeTempl(VarHomeCards, VarHomeSections)
}