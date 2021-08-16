package html

import (
	"html/template"
	"io"
	"log"

	"github.com/Ben-harder/estate/choreManager"
)

func MainPage(memberNames []string, chores []choreManager.ChoreInterface, writer io.Writer) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
		<h1>{{.Header}}<h1>
	</head>
	<body>
		{{range .MemberNames}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title       string
		Header      string
		MemberNames []string
	}{
		Title:       "Wonderful Estate",
		Header:      "The Estate Chore Manager",
		MemberNames: memberNames,
	}

	err = t.Execute(writer, data)
	check(err)

}
