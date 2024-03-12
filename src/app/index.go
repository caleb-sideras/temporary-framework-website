package app

import (
	"github.com/a-h/templ"
)

// Server Component
func Index_() templ.Component {
	// perform server-side logic -> fetch data, mutations, etc

	return IndexTempl()
}
