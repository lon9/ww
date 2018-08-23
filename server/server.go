package server

import (
	"strconv"

	"github.com/lon9/ww/game"
)

type Server struct {
	people []game.Personer
}

func NewTestServer() *Server {
	var people []game.Personer

	people = []game.Personer{
		game.NewPerson(1, game.ETeller, "a"),
		game.NewPerson(2, game.EKnight, "b"),
	}
	for i := 0; i < 2; i++ {
		people = append(people, game.NewPerson(3+i, game.EWarewolf, "c"+strconv.Itoa(i)))
	}
	for i := 0; i < 6; i++ {
		people = append(people, game.NewPerson(5+i, game.ECitizen, "d"+strconv.Itoa(i)))
	}
	return &Server{
		people,
	}
}

func (s *Server) Run() {
}
