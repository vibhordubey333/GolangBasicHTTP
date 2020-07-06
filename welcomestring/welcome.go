package welcome

import (
	"fmt"
	"net/http"
)

type WelcomeInterface interface {
	WelcomeString(w http.ResponseWriter, r *http.Request)
}

func WelcomeString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Alien")
}
