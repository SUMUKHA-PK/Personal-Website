package routing

import (
	"fmt"
	"net/http"
)

func tv(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entered TV")
}
