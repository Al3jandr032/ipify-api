
package api

import (
    "fmt"
	"github.com/julienschmidt/httprouter"
    "net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "HealthCheck!\n")
}

