package game

import (
	"fmt"
	"log"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/lon9/ww/consts"
	pb "github.com/lon9/ww/proto"
	"github.com/lon9/ww/viewmanagers"
	xcontext "golang.org/x/net/context"
)

// Teller is struct for Teller
type Teller struct {
	Person
}

// NightAction is action at night (Override)
func (t *Teller) NightAction(g *gocui.Gui, c pb.WWClient, players []*pb.Player) {
	// If already dead
	if t.GetIsDead() {
		if err := t.deadAction(g, c); err != nil {
			log.Println(err)
		}
		return
	}

	// Make player list that excludes myself and dead peoples
	var selectablePlayers []*pb.Player
	for _, player := range players {
		if !player.GetIsDead() && player.GetUuid() != t.GetUUID().String() {
			selectablePlayers = append(selectablePlayers, player)
		}
	}

	t.drawSelectablePlayerList(
		g,
		viewmanagers.MainViewID,
		"Select a player you want to tell",
		selectablePlayers,
		func(g *gocui.Gui, v *gocui.View, selected *pb.Player) error {

			// Sending tell request
			req := &pb.TellRequest{
				SrcUuid: t.GetUUID().String(),
				DstId:   selected.GetId(),
			}
			ctx, cancel := xcontext.WithTimeout(xcontext.Background(), 30*time.Second)
			defer cancel()

			res, err := c.Tell(ctx, req)
			if err != nil {
				return err
			}
			g.DeleteKeybindings(viewmanagers.MainViewID)
			v.Highlight = false
			v.Clear()
			camp, err := consts.GetCamp(res.GetCamp())
			if err != nil {
				return err
			}
			fmt.Fprintf(v, "%s (%d) is %s\n", selected.GetName(), selected.GetId(), camp)
			return nil
		})
}
