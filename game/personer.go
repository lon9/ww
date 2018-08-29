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
	ConvertPersoners(Personers) []*pb.Player
	ConvertAfter(Personers) []*pb.Player

	UpdateInfo(*gocui.Gui, []*pb.Player)
	MorningAction(*gocui.Gui, pb.WWClient, []*pb.Player)
	NightAction(*gocui.Gui, pb.WWClient, []*pb.Player)
	AfterAction(*gocui.Gui, pb.WWClient, []*pb.Player)
}
