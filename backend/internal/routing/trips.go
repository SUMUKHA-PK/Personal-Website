package routing

import (
	"fmt"
	"net/http"
)

func trips(w http.ResponseWriter, r *http.Request) {

	data := `
	[
      {
        "id": 1,
        "name": "Leanne Graham",
        "email": "Sincere@april.biz",
        "company": {
          "catchPhrase": "Multi-layered client-server neural-net",
        }
      }
    ]
	`
	fmt.Println("FG")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
}
