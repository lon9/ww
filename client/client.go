package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/lon9/ww/consts"
	"github.com/lon9/ww/game"

	"github.com/jroimartin/gocui"
	pb "github.com/lon9/ww/proto"
	"github.com/lon9/ww/viewmanagers"
	uuid "github.com/satori/go.uuid"
	xcontext "golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// MainViewID is ID for main view
	MainViewID = "main_view"
	// LeftViewID is ID for left view
	LeftViewID = "left_view"
	// RightViewID is ID for right view
	RightViewID = "right_view"
	// DialogViewID is ID for dialog view
	DialogViewID = "dialog_view"
	// DefaultViewID is ID for default view
	DefaultViewID = MainViewID
	// DefaultViewIndex is index for default view
	DefaultViewIndex = 1
)

// Client is struct for client
type Client struct {
	clientUUID uuid.UUID
	managers   []viewmanagers.ViewManager
	activeView int
	players    []*pb.Player
	state      pb.State
	mu         *sync.Mutex
	personer   game.Personer
}

// NewClient is constructor
func NewClient() *Client {
	return &Client{
		mu: new(sync.Mutex),
	}
}

// Run runs client
func (c *Client) Run(addr, port string) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	mainView := viewmanagers.NewMainView(MainViewID, false)
	leftView := viewmanagers.NewLeftView(LeftViewID, false)
	rightView := viewmanagers.NewRightView(RightViewID, true)
	dialogView := viewmanagers.NewDialogView(DialogViewID, true)
	c.managers = append(c.managers, leftView)
	c.managers = append(c.managers, mainView)
	c.managers = append(c.managers, rightView)

	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	g.SetManager(mainView, leftView, rightView, dialogView)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, c.quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, c.nextView); err != nil {
		log.Panicln(err)
	}

	conn, err := grpc.Dial(addr+":"+port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewWWClient(conn)
	go c.gameLoop(g, client)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func (c *Client) quit(g *gocui.Gui, v *gocui.View) error {
	fmt.Println("quit")
	return gocui.ErrQuit
}

func (c *Client) setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func (c *Client) setDefaultView(g *gocui.Gui) (*gocui.View, error) {
	v, err := c.setCurrentViewOnTop(g, DefaultViewID)
	if err != nil {
		return nil, err
	}
	c.activeView = DefaultViewIndex
	return v, nil
}

func (c *Client) nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (c.activeView + 1) % len(c.managers)
	manager := c.managers[nextIndex]

	if _, err := c.setCurrentViewOnTop(g, manager.GetName()); err != nil {
		return err
	}

	if manager.GetEditable() {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	c.activeView = nextIndex
	return nil
}

func (c *Client) stateLoop(g *gocui.Gui, client pb.WWClient) {
	ctx, cancel := xcontext.WithCancel(context.Background())
	defer cancel()
	stream, err := client.State(ctx, new(pb.StateRequest))
	if err != nil {
		panic(err)
	}
	defer stream.CloseSend()
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			// read done.
			break
		}
		if err != nil {
			panic(err)
		}
		c.mu.Lock()
		c.players = res.GetPlayers()
		c.state = res.GetState()
		c.mu.Unlock()
		g.Update(func(g *gocui.Gui) error {
			return c.updateState(g)
		})
	}
}

func (c *Client) gameLoop(g *gocui.Gui, client pb.WWClient) {

	// Connect to server, sending hello request
	g.Update(func(g *gocui.Gui) error {
		v, err := c.setCurrentViewOnTop(g, DialogViewID)
		if err != nil {
			return err
		}
		v.Clear()
		v.Title = "Put your name and press enter"
		v.Editable = true
		if err := g.SetKeybinding(DialogViewID, gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {

			// Get line
			line, err := v.Line(0)
			if err != nil {
				return err
			}

			// Hello request
			req := &pb.HelloRequest{
				Name: line,
			}
			ctx, cancel := xcontext.WithTimeout(xcontext.Background(), time.Second*30)
			defer cancel()
			res, err := client.Hello(ctx, req)
			if err != nil {
				return err
			}

			// Initialize personer
			c.personer = game.NewPersoner(int(res.GetId()), res.GetName(), res.GetKind())
			id, err := uuid.FromString(res.GetUuid())
			if err != nil {
				return err
			}
			c.personer.SetUUID(id)

			// Reset dialog
			v.Clear()
			v.Editable = false
			g.DeleteKeybindings(DialogViewID)

			mainView, err := c.setDefaultView(g)
			if err != nil {
				return err
			}
			kind, err := consts.GetKind(c.personer.GetKind())
			if err != nil {
				return err
			}
			camp, err := consts.GetCamp(c.personer.GetCamp())
			if err != nil {
				return err
			}
			fmt.Fprintf(mainView, "Your job is %s (%s)", kind, camp)
			go c.stateLoop(g, client)
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
}

func (c *Client) updateState(g *gocui.Gui) error {

	v, err := g.View(LeftViewID)
	if err != nil {
		return err
	}
	v.Clear()
	for _, player := range c.players {
		fmt.Fprintf(v, "%d: %s ", player.GetId(), player.GetName())
		if player.GetIsDead() {
			fmt.Fprintf(v, "Dead")
		} else {
			fmt.Fprintf(v, "Alive")
		}
	}
	return nil
}
