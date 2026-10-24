package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/00xshashank/rtchess/backend/backend/db"
	"github.com/00xshashank/rtchess/backend/backend/types"
	"github.com/gorilla/websocket"
)

var clientsToUsersMap map[string]*websocket.Conn

func init() {
	clientsToUsersMap = make(map[string]*websocket.Conn)
}

func JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Executing JoinGameHandler\n")
	userContext := r.Context().Value("user")
	if userContext == nil {
		log.Printf("Unauthenticated user called join\n")
		w.Write([]byte("Please authenticate first"))
		return
	}
	var user types.User = userContext.(types.User)
	log.Printf("User connected to websocket: %v\n", user)

	upgrader := websocket.Upgrader{}

	ws_c, ws_err := upgrader.Upgrade(w, r, nil)
	if ws_err != nil {
		log.Printf("Error while upgrading to websocket: %s\n", ws_err)
		return
	}

	clientsToUsersMap[user.Username] = ws_c
	opponent, opp_err := db.CheckForOpponentAndQueue(user)
	if opp_err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Printf("Returned from CheckForOpponentAndQueue: %s; %s\n", opponent, opp_err)
}

func StartGame(player1 types.User, player2 types.User) error {
	game_id, err := db.CreateGame(player1, player2)
	if err != nil {
		return err
	}

	clientsToUsersMap[player1.Username].WriteMessage(1, fmt.Appendf(nil, "%d", game_id))
	clientsToUsersMap[player2.Username].WriteMessage(1, fmt.Appendf(nil, "%d", game_id))

	return nil
}
