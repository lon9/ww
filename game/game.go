package game

import (
	"fmt"

	pb "github.com/lon9/ww/proto"
	uuid "github.com/satori/go.uuid"
)

// Personer is interface for person
type Personer interface {
	GetID() int
	GetUUID() uuid.UUID
	SetUUID(uuid.UUID)
	GetKind() pb.Kind
	GetCamp() pb.Camp
	GetName() string
	GetIsDead() bool
	Vote(people []Personer) int
	NightAction()
}

// Person is struct for person
type Person struct {
	ID     int
	UUID   uuid.UUID
	Kind   pb.Kind
	Camp   pb.Camp
	Name   string
	IsDead bool
}

// GetID returns ID
func (p *Person) GetID() int {
	return p.ID
}

// GetUUID returns UUID
func (p *Person) GetUUID() uuid.UUID {
	return p.UUID
}

// SetUUID sets UUID
func (p *Person) SetUUID(id uuid.UUID) {
	p.UUID = id
}

// GetKind returns kind
func (p *Person) GetKind() pb.Kind {
	return p.Kind
}

// GetCamp returns camp
func (p *Person) GetCamp() pb.Camp {
	return p.Camp
}

// GetName returns name
func (p *Person) GetName() string {
	return p.Name
}

// GetIsDead returns is the person dead
func (p *Person) GetIsDead() bool {
	return p.IsDead
}

// NightAction defines action at night
func (p *Person) NightAction() {}

// Vote votes some player
func (p *Person) Vote(people []Personer) int {
	for _, v := range people {
		fmt.Printf("%d: %s\n", v.GetID(), v.GetName())
	}
	return 1
}

// NewPersoner is constructor for Person
func NewPersoner(id int, name string, kind pb.Kind) Personer {
	switch kind {
	case pb.Kind_CITIZEN:
		return &Citizen{
			Person{
				ID:   id,
				UUID: uuid.Must(uuid.NewV4()),
				Kind: kind,
				Camp: pb.Camp_GOOD,
				Name: name,
			},
		}
	case pb.Kind_WAREWOLF:
		return &Warewolf{
			Person{
				ID:   id,
				UUID: uuid.Must(uuid.NewV4()),
				Kind: kind,
				Camp: pb.Camp_EVIL,
				Name: name,
			},
		}
	case pb.Kind_TELLER:
		return &Teller{
			Person{
				ID:   id,
				UUID: uuid.Must(uuid.NewV4()),
				Kind: kind,
				Camp: pb.Camp_GOOD,
				Name: name,
			},
		}
	case pb.Kind_KNIGHT:
		return &Knight{
			Person{
				ID:   id,
				UUID: uuid.Must(uuid.NewV4()),
				Kind: kind,
				Camp: pb.Camp_GOOD,
				Name: name,
			},
		}
	}
	return &Citizen{
		Person{
			ID:   id,
			UUID: uuid.Must(uuid.NewV4()),
			Kind: pb.Kind_CITIZEN,
			Camp: pb.Camp_GOOD,
			Name: name,
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
func (w *Warewolf) NightAction() {

}

// Teller is struct for Teller
type Teller struct {
	Person
}

// Knight is struct for Kinght
type Knight struct {
	Person
}
