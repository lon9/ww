package game

import (
	"math/rand"

	"github.com/jroimartin/gocui"
	pb "github.com/lon9/ww/proto"
	uuid "github.com/satori/go.uuid"
)

// Personer is interface for person
type Personer interface {
	GetID() int
	SetID(int)
	GetUUID() uuid.UUID
	SetUUID(uuid.UUID)
	GetKind() pb.Kind
	GetCamp() pb.Camp
	GetName() string
	SetName(string)
	GetVotes() int
	SetVotes(int)
	IncVotes()
	GetIsDead() bool
	SetIsDead(bool)
	GetAliveWill() int
	SetAliveWill(int)
	IncAliveWill()
	GetDeadWill() int
	SetDeadWill(int)
	IncDeadWill()
	Init()

	MorningAction(*gocui.Gui, pb.WWClient, []*pb.Player) error
	NightAction(*gocui.Gui, pb.WWClient, []*pb.Player) error
}

// Person is struct for person
type Person struct {
	id        int
	uid       uuid.UUID
	kind      pb.Kind
	camp      pb.Camp
	name      string
	votes     int
	isDead    bool
	deadWill  int
	aliveWill int
}

// GetID returns id
func (p *Person) GetID() int {
	return p.id
}

// SetID set id
func (p *Person) SetID(id int) {
	p.id = id
}

// GetUUID returns uid
func (p *Person) GetUUID() uuid.UUID {
	return p.uid
}

// SetUUID sets uid
func (p *Person) SetUUID(id uuid.UUID) {
	p.uid = id
}

// GetKind returns kind
func (p *Person) GetKind() pb.Kind {
	return p.kind
}

// GetCamp returns camp
func (p *Person) GetCamp() pb.Camp {
	return p.camp
}

// GetName returns name
func (p *Person) GetName() string {
	return p.name
}

// SetName set name
func (p *Person) SetName(name string) {
	p.name = name
}

// GetVotes returns votes
func (p *Person) GetVotes() int {
	return p.votes
}

// SetVotes set votes
func (p *Person) SetVotes(n int) {
	p.votes = n
}

// IncVotes inclements votes
func (p *Person) IncVotes() {
	p.votes++
}

// GetIsDead returns is the person dead
func (p *Person) GetIsDead() bool {
	return p.isDead
}

// SetIsDead is setter for isDead
func (p *Person) SetIsDead(b bool) {
	p.isDead = b
}

// GetAliveWill returns aliveWill
func (p *Person) GetAliveWill() int {
	return p.aliveWill
}

// SetAliveWill is setter for AliveWill
func (p *Person) SetAliveWill(b int) {
	p.aliveWill = b
}

// IncAliveWill inclements aliveWill
func (p *Person) IncAliveWill() {
	p.aliveWill++
}

// GetDeadWill returns deadWill
func (p *Person) GetDeadWill() int {
	return p.deadWill
}

// SetDeadWill is setter for DeadWill
func (p *Person) SetDeadWill(b int) {
	p.deadWill = b
}

// IncDeadWill inclements deadWill
func (p *Person) IncDeadWill() {
	p.deadWill++
}

// Init initializes wills
func (p *Person) Init() {
	p.votes = 0
	p.deadWill = 0
	p.aliveWill = 0
}

// NightAction defines action at night
func (p *Person) NightAction(g *gocui.Gui, c pb.WWClient, players []*pb.Player) error {
	return nil
}

// MorningAction votes some player
func (p *Person) MorningAction(g *gocui.Gui, c pb.WWClient, players []*pb.Player) error {
	return nil
}

// NewPersoner is constructor for Person
func NewPersoner(id int, name string, kind pb.Kind) Personer {
	switch kind {
	case pb.Kind_CITIZEN:
		return &Citizen{
			Person{
				id:   id,
				uid:  uuid.Must(uuid.NewV4()),
				kind: kind,
				camp: pb.Camp_GOOD,
				name: name,
			},
		}
	case pb.Kind_WAREWOLF:
		return &Warewolf{
			Person{
				id:   id,
				uid:  uuid.Must(uuid.NewV4()),
				kind: kind,
				camp: pb.Camp_EVIL,
				name: name,
			},
		}
	case pb.Kind_TELLER:
		return &Teller{
			Person{
				id:   id,
				uid:  uuid.Must(uuid.NewV4()),
				kind: kind,
				camp: pb.Camp_GOOD,
				name: name,
			},
		}
	case pb.Kind_KNIGHT:
		return &Knight{
			Person{
				id:   id,
				uid:  uuid.Must(uuid.NewV4()),
				kind: kind,
				camp: pb.Camp_GOOD,
				name: name,
			},
		}
	}
	return &Citizen{
		Person{
			id:   id,
			uid:  uuid.Must(uuid.NewV4()),
			kind: pb.Kind_CITIZEN,
			camp: pb.Camp_GOOD,
			name: name,
		},
	}
}

// Citizen is struct for citizen
type Citizen struct {
	Person
}

// Warewolf is struct for warewolf
type Warewolf struct {
	Person
}

// NightAction defines warewolf's action at night
func (w *Warewolf) NightAction(g *gocui.Gui, c pb.WWClient, players []*pb.Player) error {
	return nil
}

// Teller is struct for Teller
type Teller struct {
	Person
}

// Knight is struct for Kinght
type Knight struct {
	Person
}

// Personers is slice of Personer
type Personers map[int]Personer

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

// ConvertPersoners converts Personers to []*pb.Player
func (ps Personers) ConvertPersoners() []*pb.Player {
	var players []*pb.Player

	for _, v := range ps {
		player := &pb.Player{
			Id:     int32(v.GetID()),
			Name:   v.GetName(),
			IsDead: v.GetIsDead(),
		}
		players = append(players, player)
	}
	return players
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
