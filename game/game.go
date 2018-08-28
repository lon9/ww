package game

import (
	"fmt"
	"math/rand"

	"github.com/jroimartin/gocui"
	pb "github.com/lon9/ww/proto"
	"github.com/lon9/ww/viewmanagers"
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
	ConvertPersoners(Personers) []*pb.Player

	UpdateInfo(*gocui.Gui, []*pb.Player) error
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

// ConvertPersoners converts Personers to []*pb.Player
func (p *Person) ConvertPersoners(personers Personers) []*pb.Player {
	players := make([]*pb.Player, len(personers))

	for i, v := range personers {
		player := &pb.Player{
			Id:     int32(v.GetID()),
			Name:   v.GetName(),
			IsDead: v.GetIsDead(),
		}
		players[i] = player
	}
	return players
}

// UpdateInfo updates info view
func (p *Person) UpdateInfo(g *gocui.Gui, players []*pb.Player) error {
	// Update left view
	v, err := g.View(viewmanagers.LeftViewID)
	if err != nil {
		return err
	}
	v.Clear()
	for _, player := range players {
		fmt.Fprintf(v, "%d: %s ", player.GetId(), player.GetName())
		if player.GetIsDead() {
			fmt.Fprint(v, "Dead")
		} else {
			fmt.Fprint(v, "Alive")
		}
		fmt.Fprintln(v)
	}
	return nil
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

// ConvertPersoners converts Personers to []*pb.Player (Override)
func (w *Warewolf) ConvertPersoners(personers Personers) []*pb.Player {
	players := make([]*pb.Player, len(personers))

	for i, v := range personers {
		player := &pb.Player{
			Id:     int32(v.GetID()),
			Name:   v.GetName(),
			IsDead: v.GetIsDead(),
		}
		if v.GetKind() == pb.Kind_WAREWOLF {
			player.Kind = v.GetKind()
		}
		players[i] = player
	}
	return players
}

// UpdateInfo updates information of left view (Override)
func (w *Warewolf) UpdateInfo(g *gocui.Gui, players []*pb.Player) error {
	// Update left view
	v, err := g.View(viewmanagers.LeftViewID)
	if err != nil {
		return err
	}
	v.Clear()
	for _, player := range players {
		fmt.Fprintf(v, "%d: %s ", player.GetId(), player.GetName())
		if player.GetIsDead() {
			fmt.Fprint(v, "Dead")
		} else {
			fmt.Fprint(v, "Alive")
		}
		if player.GetKind() == pb.Kind_WAREWOLF {
			fmt.Fprint(v, " W")
		}
		fmt.Fprintln(v)
	}
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
