package html

import (
	"html/template"
	"io"
	"log"

	"github.com/Ben-harder/estate/choreManager/chore"
	"github.com/Ben-harder/estate/household"
)

func MainPage(chores []chore.ChoreInterface, household household.HouseholdInterface, writer io.Writer) {
	const tpl = `
<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- Bootstrap CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-KyZXEAg3QhqLMpG8r+8fhAXLRk2vvoC2f3B09zVXn8CA5QIVfZOJ3BCsw2P0p/We" crossorigin="anonymous">

    <title>{{.Title}}</title>
  </head>
  <body>
  <div class="container">
	<div class="row py-5">
	  <div class="col">
	  	<div class="container">
		  <h1 class="display-4">Wonderful Estate</h1>
		  <p class="lead">The official website of The Estate</p>
		</div>
	  </div>
	</div>
  	<div class="row">
      <div class="col">
    	<div class="container">
		  <h1>Housemates</h1>
		  {{ .MemberNames }} 
		</div>
	  </div>
	<div class="row py-4">
	  <div class="col">
	    <div class="container">
		  <h1>Chores</h1>
		  <table class="table-responsive-lg">
			<thead>
				<tr>
				<th scope="col">Schedule</th>
				<th scope="col">Responsibilities</th>
				<th scope="col">Date</th>
				<th scope="col">Whose Turn</th>
				</tr>
			</thead>
			<tbody>
			{{range .Chores}}
			  <tr>
				<th scope="row">{{ .Schedule }}</th>
				<td>{{ .Responsibilities }}</td>
				<td>{{ .Date }}</td>
				<td>{{ .WhoseTurn }}</td>
			  </tr>
			{{end}}
			</tbody>
			</table>
	    </div>
	  </div>
	</div>
  	</div>
  </div>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-U1DAWAznBHeqEIlVSCgzq+c9gqGAJn5c/t99JyeKa9xxaYpSvHU5awsuZVVFIhvj" crossorigin="anonymous"></script>
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
		MemberNames string
		Chores      []chore.ChoreInterface
	}{
		Title:       "Wonderful Estate",
		Header:      "The Estate Chore Manager",
		MemberNames: household.String(),
		Chores:      chores,
	}

	err = t.Execute(writer, data)
	check(err)

}
