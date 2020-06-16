package routing

import (
	"fmt"
	"net/http"
)

func movies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entered movies")
}
