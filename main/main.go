package main

import "github.com/fanofsherlock/BasicHttpServer_Go/TimeApiMuxServer"

func main() {
	/*
			handler := &PlayerServer{}

		Did we register a mux here?

		 What about this approach of declaring a handler right in the library

		       http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(200)
					fmt.Fprintf(w, "Hello to Here")
				})

				log.Fatal(http.ListenAndServe(":2000", nil))

		log.Fatal(http.ListenAndServe(":5000", handler))
	*/

	//BasicServer.StartBasicServer()

	TimeApiMuxServer.StartTimeServer()
}
