package game

import (
	"fmt"
	"log"
	"time"

	"github.com/jroimartin/gocui"
	pb "github.com/lon9/ww/proto"
	"github.com/lon9/ww/viewmanagers"
	xcontext "golang.org/x/net/context"
)

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
func (w *Warewolf) UpdateInfo(g *gocui.Gui, players []*pb.Player) {
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
			if player.GetKind() == pb.Kind_WAREWOLF {
				fmt.Fprint(v, " W")
			}
			fmt.Fprintln(v)
		}
		return nil
	})
}

// NightAction is action at night (Override)
func (w *Warewolf) NightAction(g *gocui.Gui, c pb.WWClient, players []*pb.Player) {
	// If already dead
	if w.GetIsDead() {
		viewmanagers.DrawDeadView(g, viewmanagers.MainViewID)
		return
	}

	// Make player list that excludes myself and dead peoples and my kind
	var selectablePlayers []*pb.Player
	for _, player := range players {
		if !player.GetIsDead() && int(player.GetId()) != w.GetID() && player.GetKind() != pb.Kind_WAREWOLF {
			selectablePlayers = append(selectablePlayers, player)
		}
	}

	w.drawSelectablePlayerList(
		g,
		viewmanagers.MainViewID,
		"Select a player you want to bite",
		selectablePlayers,
		func(g *gocui.Gui, v *gocui.View, selected *pb.Player) error {

			// Sending bite request
			req := &pb.BiteRequest{
				SrcUuid: w.GetUUID().String(),
				DstId:   selected.GetId(),
			}
			ctx, cancel := xcontext.WithTimeout(xcontext.Background(), 30*time.Second)
			defer cancel()

			_, err := c.Bite(ctx, req)
			if err != nil {
				return err
			}
			g.DeleteKeybindings(viewmanagers.MainViewID)
			v.Highlight = false
			v.Clear()
			fmt.Fprintf(v, "You'll bite a player named %s (%d)\n", selected.GetName(), selected.GetId())
			return nil
		})
}

// AfterAction is action when the game is finished (Override)
func (w *Warewolf) AfterAction(g *gocui.Gui, c pb.WWClient, players []*pb.Player) {
	personers := make(Personers)
	personers.FromPlayers(players)
	wonCamp, err := personers.WhichWon()
	if err != nil {
		log.Println(err)
		return
	}
	if wonCamp == pb.Camp_GOOD {
		w.drawAfterView(g, viewmanagers.MainViewID, "You lose", players)
	} else if wonCamp == pb.Camp_EVIL {
		w.drawAfterView(g, viewmanagers.MainViewID, "You won", players)
	}
}
