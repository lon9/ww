package game

import (
	"errors"
	"fmt"
	"math/rand"

	pb "github.com/lon9/ww/proto"
)

// Personers is slice of Personer
type Personers map[int]Personer

// FromPlayers initialize Personers from []*pb.FromPlayers
func (ps Personers) FromPlayers(players []*pb.Player) {
	for _, v := range players {
		ps[int(v.GetId())] = NewPersoner(int(v.GetId()), v.GetName(), v.GetKind())
		ps[int(v.GetId())].SetIsDead(v.GetIsDead())
	}
}

// NumAlive returns the number of alive persons
func (ps Personers) NumAlive() int {
	var n int
	for _, v := range ps {
		if !v.GetIsDead() {
			n++
		}
	}
	return n
}

// NumAliveGood returns the number of alive good people
func (ps Personers) NumAliveGood() int {
	var n int
	for _, v := range ps {
		if !v.GetIsDead() && v.GetCamp() == pb.Camp_GOOD {
			n++
		}
	}
	return n
}

// NumAliveEvil returns the number of alive evil people
func (ps Personers) NumAliveEvil() int {
	var n int
	for _, v := range ps {
		if !v.GetIsDead() && v.GetCamp() == pb.Camp_EVIL {
			n++
		}
	}
	return n
}

// IsFinish returns whether the game is finish or not
func (ps Personers) IsFinish() bool {
	nAliveGood := ps.NumAliveGood()
	nAliveEvil := ps.NumAliveEvil()
	if nAliveEvil == 0 {
		return true
	}
	if nAliveGood <= nAliveEvil {
		return true
	}
	return false
}

// WhichWon returns which camp won
func (ps Personers) WhichWon() (pb.Camp, error) {
	nAliveGood := ps.NumAliveGood()
	nAliveEvil := ps.NumAliveEvil()
	if nAliveEvil == 0 {
		return pb.Camp_GOOD, nil
	}
	if nAliveGood <= nAliveEvil {
		return pb.Camp_EVIL, nil
	}
	return -1, errors.New("Is not finish the game")
}

// Init initializes wills of members
func (ps Personers) Init() {
	for k := range ps {
		ps[k].Init()
	}
}

// ResolveDeadOrAliveAtNight calls DeadOrAlive of the member of the slice
func (ps Personers) ResolveDeadOrAliveAtNight() {
	var candice []Personer
	for k := range ps {
		if ps[k].GetDeadWill() > 0 && ps[k].GetAliveWill() == 0 {
			candice = append(candice, ps[k])
		}
	}
	candice[rand.Intn(len(candice))].SetIsDead(true)
}

// ResolveDeadOrAliveAtMorning resolve dead or alive at morning state
func (ps Personers) ResolveDeadOrAliveAtMorning() {
	var maxVotes int
	var candice []Personer
	for k := range ps {
		nVotes := ps[k].GetVotes()
		if nVotes > maxVotes {
			maxVotes = nVotes
		}
	}
	for k := range ps {
		if ps[k].GetVotes() == maxVotes {
			candice = append(candice, ps[k])
		}
	}
	candice[rand.Intn(len(candice))].SetIsDead(true)
}

// ValidKind valids kind with uuid
func (ps Personers) ValidKind(uid string, kind pb.Kind) bool {
	for _, v := range ps {
		if v.GetUUID().String() == uid && v.GetKind() == kind {
			return true
		}
	}
	return false
}

// FindPersonerByUUID finds Personer by UUID
func (ps Personers) FindPersonerByUUID(uid string) (Personer, error) {
	for _, v := range ps {
		if v.GetUUID().String() == uid {
			return v, nil
		}
	}
	return nil, fmt.Errorf("Not found personer of the uuid: %s", uid)
}
