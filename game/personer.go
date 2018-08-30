package game

import (
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
	Update([]*pb.Player)
	ConvertPersoners(Personers) []*pb.Player
	ConvertAfter(Personers) []*pb.Player

	UpdateInfo(*gocui.Gui, []*pb.Player)
	MorningAction(*gocui.Gui, pb.WWClient, []*pb.Player)
	NightAction(*gocui.Gui, pb.WWClient, []*pb.Player)
	AfterAction(*gocui.Gui, pb.WWClient, []*pb.Player)
	RestartAction(*gocui.Gui, pb.WWClient)
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
	case pb.Kind_WEREWOLF:
		return &Werewolf{
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
