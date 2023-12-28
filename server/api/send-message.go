package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	c, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading body", err)
		fmt.Fprintf(w, "Not-OK")
		return
	}

	fmt.Println("BODY", string(c))
	fmt.Fprintf(w, "Server Echo %s", string(c))
}
