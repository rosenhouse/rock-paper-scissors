package web_ui

import (
	"fmt"
	"net/http"

	"github.com/desmondrawls/rock-paper-scissors/play"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		renderPage(w, pageFields{Header: "Record a game"})
		return
	}
	if r.URL.Path == "/play" && r.Method == "POST" {
		throws := play.Inputs{
			Player1Name:  "player1",
			Player2Name:  "player2",
			Player1Throw: r.FormValue("player1"),
			Player2Throw: r.FormValue("player2"),
		}
		play.Play(throws, &web_ui{
			ResponseWriter: w,
		})
	}
}

type web_ui struct {
	http.ResponseWriter
}

func (w web_ui) Winner(name string) {
	w.Write([]byte(fmt.Sprintf("<body>%s <br> WINS!</body>", name)))
}

func (w web_ui) Draw() {
	w.Write([]byte("TIE!"))
}

func (w web_ui) Invalid(throws play.Inputs) {
	renderPage(w, pageFields{
		Header:       "Invalid input",
		Player1Value: throws.Player1Throw,
		Player2Value: throws.Player2Throw,
	})
}
