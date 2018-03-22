package web_ui

import (
	"html/template"
	"log"
	"net/http"
)

const (
	pageTemplateText = `<html><body>
<h1>{{.Header}}</h1>
<form action="/play" method="POST">
<label for="player1">P1</label>
<input name="player1" type="string" value="{{.Player1Value}}"/>
<br>
<label for="player2">P2</label>
<input name="player2" type="string" value="{{.Player2Value}}"/>
<br>
<input type="submit" value="Play" />
</form>
</body>
</html>`
)

var pageTemplate = template.Must(template.New("page").Parse(pageTemplateText))

type pageFields struct {
	Header       string
	Player1Value string
	Player2Value string
}

func renderPage(w http.ResponseWriter, fields pageFields) {
	err := pageTemplate.Execute(w, fields)
	if err != nil {
		log.Printf("error in template: %s\n", err)
	}
}
