package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/lon9/ww/game"

	"github.com/jroimartin/gocui"
	pb "github.com/lon9/ww/proto"
	"github.com/lon9/ww/viewmanagers"
	uuid "github.com/satori/go.uuid"
	xcontext "golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	MainViewID       = "main_view"
	LeftViewID       = "left_view"
	RightViewID      = "right_view"
	DialogViewID     = "dialog_view"
	DefaultViewID    = MainViewID
	DefaultViewIndex = 1
)

type Client struct {
	clientUUID uuid.UUID
	managers   []viewmanagers.ViewManager
	activeView int
	players    []*pb.Player
	state      pb.State
	mu         *sync.Mutex
	personer   game.Personer
}

func NewClient() *Client {
	return &Client{
		mu: new(sync.Mutex),
	}
}

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

	g.Update(func(g *gocui.Gui) error {
		v, err := c.setCurrentViewOnTop(g, DialogViewID)
		if err != nil {
			return err
		}
		v.Clear()
		v.Title = "Put your name"
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
			c.personer = game.NewPersoner(int(res.GetId()), res.GetName(), res.GetKind())
			id, err := uuid.FromString(res.GetUuid())
			if err != nil {
				return err
			}
			c.personer.SetUUID(id)
			c.setDefaultView(g)
			// Reset dialog
			v.Clear()
			v.Editable = false
			g.DeleteKeybindings(DialogViewID)

			mainView, err := g.View(MainViewID)
			if err != nil {
				return err
			}
			fmt.Fprint(mainView, c.personer)
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
	fmt.Fprint(v, c.state)
	return nil
}
