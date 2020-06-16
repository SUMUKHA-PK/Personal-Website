package routing

import (
	"fmt"
	"net/http"
)

func books(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entered books")

}
