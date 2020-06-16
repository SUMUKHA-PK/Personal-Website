package routing

import (
	"fmt"
	"net/http"
)

func blogs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entered blogs")
}
