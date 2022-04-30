package timezone

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type CurrentTime struct {
	Time string `json:"current_time" xml:"current_time"`
}

func getTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["tz"])

	timezone, ok := r.URL.Query()["tz"]
	if !ok || len(timezone[0]) < 1 {
		fmt.Printf("NO: Timezone provided. Return UTC time\n")
		currTime := CurrentTime{
			Time: time.Now().UTC().String(),
		}
		json.NewEncoder(w).Encode(currTime)
	} else {
		// for loop for keys
		tzones := strings.Split(timezone[0], ",")
		for _, t := range tzones {
			fmt.Printf("Timezone provided: %s. Return location time.\n", t)
			loc, err := time.LoadLocation(t)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte("Invalid location: "))
				http.NotFound(w, r)
			} else {
				w.Header().Add("Content-Type", "application/json")
				currTime := CurrentTime{
					Time: time.Now().In(loc).String(),
				}
				json.NewEncoder(w).Encode(currTime)
			}
		}
	}
}
