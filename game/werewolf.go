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

// Werewolf is struct for werewolf
type Werewolf struct {
	Person
}

// ConvertPersoners converts Personers to []*pb.Player (Override)
func (w *Werewolf) ConvertPersoners(personers Personers) []*pb.Player {
	players := make([]*pb.Player, len(personers))

	for i, v := range personers {
		player := &pb.Player{
			Id:     int32(v.GetID()),
			Name:   v.GetName(),
			IsDead: v.GetIsDead(),
		}
		if w.GetUUID() == v.GetUUID() {
			player.Uuid = v.GetUUID().String()
		}
		if v.GetKind() == pb.Kind_WEREWOLF {
			player.Kind = v.GetKind()
		}
		players[i] = player
	}
	return players
}

// NightAction is action at night (Override)
func (w *Werewolf) NightAction(g *gocui.Gui, c pb.WWClient, players []*pb.Player) {
	// If already dead
	if w.GetIsDead() {
		if err := w.deadAction(g, c); err != nil {
			log.Println(err)
		}
		return
	}

	// Make player list that excludes myself and dead peoples and my kind
	var selectablePlayers []*pb.Player
	for _, player := range players {
		if !player.GetIsDead() && player.GetUuid() != w.GetUUID().String() && player.GetKind() != pb.Kind_WEREWOLF {
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
func (w *Werewolf) AfterAction(g *gocui.Gui, c pb.WWClient, players []*pb.Player) {
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
