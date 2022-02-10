package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store PlayerStore
}

type PlayerStore interface {
	GetPlayerScore(name string) string
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	write(w, p, player)
}

func write(w http.ResponseWriter, p *PlayerServer, player string) {
	score := p.store.GetPlayerScore(player)
	if score == "" {
		w.WriteHeader(404)
	} else {
		_, err := fmt.Fprintf(w, score)
		if err != nil {
			panic(err)
		}
	}
}

//func GetPlayerScore(name string) string {
//	switch name {
//	case "shivam":
//		return "10"
//	case "anu":
//		return "20"
//	default:
//		return ""
//	}
//}
