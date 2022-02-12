package TimeApiMuxServer

import (
	"encoding/json"
	"net/http"
	"time"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {

	if q := r.URL.Query(); len(q) == 0 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200) //This needs to come last
		json.NewEncoder(w).Encode(map[string]string{"current_time": time.Now().String()})
	} else {
		v, ok := q["tz"]
		if !ok {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(404)
		} else {
			times := make(map[string]string, len(v))

			for _, s := range v {
				loc, err := time.LoadLocation(s)
				if err != nil {
					times[s] = "tz not found"
				} else {
					times[s] = time.Now().In(loc).String()
				}
			}

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200) //This needs to come last
			json.NewEncoder(w).Encode(times)

		}
	}

}
