package handler

import (
	"main/src/code"
	"main/src/templates"
	"net/http"
	"text/template"

	"github.com/sourcegraph/syntaxhighlight"
)

type TCodeTemplate struct {
	Code string
	Lang string
}

func Handler(writer http.ResponseWriter, request *http.Request) {
	highlightedCode, _ := syntaxhighlight.AsHTML([]byte(code.Js))

	var codeTemplate = TCodeTemplate{
		Code: string(highlightedCode),
		Lang: "javascript",
	}

	htmlTemplate, _ := template.New("show-off-term").Parse(templates.Html)
 
	writer.Header().Set("Content-Type","image/svg+xml")
	htmlTemplate.Execute(writer, &codeTemplate)
}
