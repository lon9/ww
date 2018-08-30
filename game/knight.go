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

// Knight is struct for Kinght
type Knight struct {
	Person
}

// NightAction is action at night (Override)
func (k *Knight) NightAction(g *gocui.Gui, c pb.WWClient, players []*pb.Player) {
	// If already dead
	if k.GetIsDead() {
		if err := k.deadAction(g, c); err != nil {
			log.Println(err)
		}
		return
	}

	// Make player list that excludes myself and dead peoples and my kind
	var selectablePlayers []*pb.Player
	for _, player := range players {
		if !player.GetIsDead() && int(player.GetId()) != k.GetID() {
			selectablePlayers = append(selectablePlayers, player)
		}
	}

	k.drawSelectablePlayerList(
		g,
		viewmanagers.MainViewID,
		"Select a player you want to protect",
		selectablePlayers,
		func(g *gocui.Gui, v *gocui.View, selected *pb.Player) error {

			// Sending protect request
			req := &pb.ProtectRequest{
				SrcUuid: k.GetUUID().String(),
				DstId:   selected.GetId(),
			}
			ctx, cancel := xcontext.WithTimeout(xcontext.Background(), 30*time.Second)
			defer cancel()

			_, err := c.Protect(ctx, req)
			if err != nil {
				return err
			}
			g.DeleteKeybindings(viewmanagers.MainViewID)
			v.Highlight = false
			v.Clear()
			fmt.Fprintf(v, "You'll protect a player named %s (%d)", selected.GetName(), selected.GetId())
			return nil
		})
}
