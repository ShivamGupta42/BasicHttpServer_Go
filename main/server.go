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
	GetPlayerScore(name string) int
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	write(w, player)
}

func write(w http.ResponseWriter, player string) {
	_, err := fmt.Fprintf(w, GetPlayerScore(player))
	if err != nil {
		panic(err)
	}
}

func GetPlayerScore(name string) string {
	switch name {
	case "shivam":
		return "10"
	case "anu":
		return "20"
	default:
		return ""
	}
}
