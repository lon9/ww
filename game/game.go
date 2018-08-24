package game

import (
	"fmt"

	pb "github.com/lon9/ww/proto"
	uuid "github.com/satori/go.uuid"
)

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

type Person struct {
	ID     int
	UUID   uuid.UUID
	Kind   pb.Kind
	Camp   pb.Camp
	Name   string
	IsDead bool
}

func (p *Person) GetID() int {
	return p.ID
}

func (p *Person) GetUUID() uuid.UUID {
	return p.UUID
}

func (p *Person) SetUUID(id uuid.UUID) {
	p.UUID = id
}

func (p *Person) GetKind() pb.Kind {
	return p.Kind
}

func (p *Person) GetCamp() pb.Camp {
	return p.Camp
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetIsDead() bool {
	return p.IsDead
}

func (p *Person) NightAction() {}

func (p *Person) Vote(people []Personer) int {
	for _, v := range people {
		fmt.Printf("%d: %s\n", v.GetID(), v.GetName())
	}
	return 1
}

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

type Citizen struct {
	Person
}

type Warewolf struct {
	Person
}

func (w *Warewolf) NightAction() {

}

type Teller struct {
	Person
}

type Knight struct {
	Person
}
