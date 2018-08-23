package client

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/lon9/ww/viewmanagers"
)

type Client struct {
	managers []viewmanagers.ViewManager
	active   int
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Run() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	mainView := viewmanagers.NewMainView("main", false)
	leftView := viewmanagers.NewLeftView("left", false)
	rightView := viewmanagers.NewRightView("memo", true)
	c.managers = append(c.managers, leftView)
	c.managers = append(c.managers, mainView)
	c.managers = append(c.managers, rightView)

	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	g.SetManager(mainView, leftView, rightView)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, c.quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, c.nextView); err != nil {
		log.Panicln(err)
	}

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

func (c *Client) nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (c.active + 1) % len(c.managers)
	manager := c.managers[nextIndex]

	if _, err := c.setCurrentViewOnTop(g, manager.GetName()); err != nil {
		return err
	}

	if manager.GetEditable() {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	c.active = nextIndex
	return nil
}
