package client

import (
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

	// Initialize views
	mainView := viewmanagers.NewMainView(viewmanagers.MainViewID, false)
	leftView := viewmanagers.NewLeftView(viewmanagers.LeftViewID, false)
	rightView := viewmanagers.NewRightView(viewmanagers.RightViewID, true)
	dialogView := viewmanagers.NewDialogView(viewmanagers.DialogViewID, true)

	// Add view managers
	c.managers = append(c.managers, leftView)
	c.managers = append(c.managers, mainView)
	c.managers = append(c.managers, rightView)

	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	g.SetManager(mainView, leftView, rightView, dialogView)

	// Quit
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, c.quit); err != nil {
		log.Panicln(err)
	}

	// Change focus
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, c.nextView); err != nil {
		log.Panicln(err)
	}

	// Connect to server
	conn, err := grpc.Dial(addr+":"+port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewWWClient(conn)

	// Initialize
	go c.initialize(g, client)

	// Start UI loop
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
	v, err := c.setCurrentViewOnTop(g, viewmanagers.DefaultViewID)
	if err != nil {
		return nil, err
	}
	c.activeView = viewmanagers.DefaultViewIndex
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
	ctx, cancel := xcontext.WithCancel(xcontext.Background())
	defer cancel()
	stream, err := client.State(ctx, &pb.StateRequest{
		Uuid: c.personer.GetUUID().String(),
	})
	if err != nil {
		panic(err)
	}
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
		c.personer.Update(c.players)
		c.mu.Unlock()
		c.doAction(g, client)
	}
}

func (c *Client) initialize(g *gocui.Gui, client pb.WWClient) {

	// Connect to server, sending hello request
	g.Update(func(g *gocui.Gui) error {
		v, err := c.setCurrentViewOnTop(g, viewmanagers.DialogViewID)
		if err != nil {
			return err
		}
		v.Clear()
		v.Title = "Put your name and press enter"
		v.Editable = true
		err = g.SetKeybinding(
			viewmanagers.DialogViewID,
			gocui.KeyEnter,
			gocui.ModNone,
			func(g *gocui.Gui, v *gocui.View) error {

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
				g.DeleteKeybindings(viewmanagers.DialogViewID)

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

				// Start state loop
				go c.stateLoop(g, client)
				return nil
			})
		if err != nil {
			return err
		}
		return nil
	})
}

func (c *Client) doAction(g *gocui.Gui, client pb.WWClient) {
	c.personer.UpdateInfo(g, c.players)
	// Do specific action
	switch c.state {
	case pb.State_MORNING:
		c.personer.MorningAction(g, client, c.players)
	case pb.State_NIGHT:
		c.personer.NightAction(g, client, c.players)
	case pb.State_AFTER:
		c.personer.AfterAction(g, client, c.players)
	}
}
