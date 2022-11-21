package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"main.go/src/teams"
)

func main() {

	http.HandleFunc("/runSweepstake", PostHandler)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
	}

}

func PostHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Received. Body of the request is: ", &req.Body)

	var players []string

	decodeRequestBody(req.Body, &players)

	fmt.Println(players)

	if len(players) < 1 || len(players) > 32 {
		rw.WriteHeader(400)
		rw.Write([]byte("Invalid number of players must be between 1 and 32 players"))
		return
	}

	playersWithTeams := teams.AllocateTeamsToPlayers(players, teams.FetchWorldCupTeams())

	resp, err := json.Marshal(playersWithTeams)
	if err != nil {
		fmt.Println("Error marshalling response")
	}

	rw.WriteHeader(200)
	rw.Write(resp)

	fmt.Println(playersWithTeams)
}

func decodeRequestBody(body io.ReadCloser, players *[]string) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&players)
	if err != nil {
		log.Fatal("error decoding request body")
	}
	return
}
