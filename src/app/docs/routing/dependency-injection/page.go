package docs_routing_dependencyinjecton

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/dependency_injection.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
