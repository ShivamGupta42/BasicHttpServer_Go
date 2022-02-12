package TimeApiMuxServer

import (
	"encoding/json"
	"net/http"
	"time"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if len(q) == 0 {
		ContHeaderAndResCode(w, "application/json", 200)
		JsonEncodeData(w, map[string]string{"current_time": time.Now().String()})
	} else {
		parseQueries(w, q)
	}
}

func parseQueries(w http.ResponseWriter, q map[string][]string) {
	v, ok := q["tz"]

	if !ok {
		ContHeaderAndResCode(w, "application/json", 404)
	} else {
		parseAndLoadTzs(w, q, v)
	}
}

func parseAndLoadTzs(w http.ResponseWriter, q map[string][]string, v []string) {
	times := make(map[string]string, len(v))

	for _, s := range v {
		loc, err := time.LoadLocation(s)
		if err != nil {
			times[s] = "tz not found"
		} else {
			times[s] = time.Now().In(loc).String()
		}
	}

	ContHeaderAndResCode(w, "application/json", 200)
	JsonEncodeData(w, times)

}

func JsonEncodeData(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func ContHeaderAndResCode(w http.ResponseWriter, value string, resCode int) {
	w.Header().Add("Content-Type", value)
	w.WriteHeader(resCode) //This needs to come last
}
