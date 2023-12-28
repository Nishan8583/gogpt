package server

import (
	"embed"
	"fmt"
	"gogpt/server/api"
	"log"
	"net/http"
	"text/template"
)

func StartServer(content embed.FS) {
	log.Println("Starting server ...")

	tmplHandler := func(w http.ResponseWriter, r *http.Request) {
		tmplData, err := content.ReadFile("templates/chat.html")
		if err != nil {
			log.Println("could not access the template")
			w.WriteHeader(500)
			fmt.Fprintf(w, "error")
			return
		}

		tmpl, err := template.New("index").Parse(string(tmplData))
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "error")
			return
		}
		tmpl.Execute(w, "")
	}

	http.HandleFunc("/", tmplHandler)
	http.HandleFunc("/api/v1/send-message", api.SendMessage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
