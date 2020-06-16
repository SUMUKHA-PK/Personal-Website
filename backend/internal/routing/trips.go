package routing

import (
	"fmt"
	"net/http"
)

func trips(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entered trips")
}
