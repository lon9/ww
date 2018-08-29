package game

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/lon9/ww/consts"
	pb "github.com/lon9/ww/proto"
	"github.com/lon9/ww/viewmanagers"
	uuid "github.com/satori/go.uuid"
	xcontext "golang.org/x/net/context"
)

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

// Update updates status
func (p *Person) Update(players []*pb.Player) {
	for _, player := range players {
		if p.GetID() == int(player.GetId()) {
			p.SetIsDead(player.GetIsDead())
		}
	}
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

// ConvertAfter convert Personers to player with all info
func (p *Person) ConvertAfter(personers Personers) []*pb.Player {
	players := make([]*pb.Player, len(personers))

	for i, v := range personers {
		player := &pb.Player{
			Id:     int32(v.GetID()),
			Name:   v.GetName(),
			IsDead: v.GetIsDead(),
			Camp:   v.GetCamp(),
			Kind:   v.GetKind(),
		}
		players[i] = player
	}
	return players
}

// UpdateInfo updates info view
func (p *Person) UpdateInfo(g *gocui.Gui, players []*pb.Player) {
	g.Update(func(g *gocui.Gui) error {
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
	})
}

// NightAction defines action at night
func (p *Person) NightAction(g *gocui.Gui, c pb.WWClient, players []*pb.Player) {
	// Common night action

	// If already dead
	if p.GetIsDead() {
		viewmanagers.DrawDeadView(g, viewmanagers.MainViewID)
		return
	}

	// Sending sleep request
	ctx, cancel := xcontext.WithTimeout(xcontext.Background(), 30*time.Second)
	defer cancel()
	_, err := c.Sleep(ctx, &pb.SleepRequest{
		SrcUuid: p.GetUUID().String(),
	})
	if err != nil {
		log.Println(err)
		return
	}

	g.Update(func(g *gocui.Gui) error {
		mainView, err := g.View(viewmanagers.MainViewID)
		if err != nil {
			return err
		}
		mainView.Clear()
		fmt.Fprintln(mainView, "Waiting for morning")
		return nil
	})
}

// MorningAction votes some player
func (p *Person) MorningAction(g *gocui.Gui, c pb.WWClient, players []*pb.Player) {
	// Common morning action

	// If already dead
	if p.GetIsDead() {
		viewmanagers.DrawDeadView(g, viewmanagers.MainViewID)
		return
	}

	for i := consts.DiscussionTime; i > 0; i-- {

		// Discussion time 60 seconds
		g.Update(func(g *gocui.Gui) error {
			mainView, err := g.View(viewmanagers.MainViewID)
			if err != nil {
				return err
			}
			mainView.Clear()
			fmt.Fprintf(mainView, "Vote in %d seconds, discuss with other players", i)
			return nil
		})
		time.Sleep(1 * time.Second)
	}

	// Make player list that excludes myself and dead peoples
	var selectablePlayers []*pb.Player
	for _, player := range players {
		if !player.GetIsDead() && int(player.GetId()) != p.GetID() {
			selectablePlayers = append(selectablePlayers, player)
		}
	}

	// Vote a player
	p.drawSelectablePlayerList(
		g,
		viewmanagers.MainViewID,
		"Vote a player",
		selectablePlayers,
		func(g *gocui.Gui, v *gocui.View, selected *pb.Player) error {

			// Sending vote request
			req := &pb.VoteRequest{
				SrcUuid: p.GetUUID().String(),
				DstId:   selected.GetId(),
			}
			ctx, cancel := xcontext.WithTimeout(xcontext.Background(), 30*time.Second)
			defer cancel()
			_, err := c.Vote(ctx, req)
			if err != nil {
				return err
			}
			g.DeleteKeybindings(viewmanagers.MainViewID)
			v.Highlight = false
			v.Clear()
			fmt.Fprintln(v, "Waiting for other players")
			return nil
		})
}

// AfterAction is action did when game is finished
func (p *Person) AfterAction(g *gocui.Gui, c pb.WWClient, players []*pb.Player) {
	personers := make(Personers)
	personers.FromPlayers(players)
	wonCamp, err := personers.WhichWon()
	if err != nil {
		log.Println(err)
		return
	}
	if wonCamp == pb.Camp_GOOD {
		p.drawAfterView(g, viewmanagers.MainViewID, "You won", players)
	} else if wonCamp == pb.Camp_EVIL {
		p.drawAfterView(g, viewmanagers.MainViewID, "You lose", players)
	}
}

func (p *Person) drawSelectablePlayerList(g *gocui.Gui,
	viewID,
	msg string,
	players []*pb.Player,
	onSelected func(g *gocui.Gui, v *gocui.View, player *pb.Player) error) {

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(viewID)
		if err != nil {
			return err
		}
		v.Clear()

		// Message
		fmt.Fprintln(v, msg)

		// Player list
		for _, player := range players {
			fmt.Fprintf(v, "%d: %s\n", player.GetId(), player.GetName())
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		// Setting selector to top of selectable
		if err := v.SetCursor(0, 1); err != nil {
			return err
		}

		// Setting key bindings
		if err = g.SetKeybinding(viewmanagers.MainViewID, gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			return viewmanagers.CursorDownWithRange(v, len(players))
		}); err != nil {
			return err
		}
		if err = g.SetKeybinding(viewmanagers.MainViewID, gocui.KeyArrowUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			return viewmanagers.CursorUpWithRange(v, 1)
		}); err != nil {
			return err
		}
		err = g.SetKeybinding(viewmanagers.MainViewID, gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			index := viewmanagers.GetLineIndex(v)
			if index < 1 || index > len(players) {
				return errors.New("Index out of range")
			}
			return onSelected(g, v, players[index-1])
		})
		return err
	})
}

func (p *Person) drawAfterView(g *gocui.Gui, viewID, msg string, players []*pb.Player) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(viewID)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, msg)
		for _, player := range players {
			kind, err := consts.GetKind(player.GetKind())
			if err != nil {
				return err
			}
			if player.GetIsDead() {
				fmt.Fprintf(v, "%d: %s %s %s\n", player.GetId(), player.GetName(), kind, "Dead")
			} else {
				fmt.Fprintf(v, "%d: %s %s %s\n", player.GetId(), player.GetName(), kind, "Alive")
			}
		}
		return nil
	})
}