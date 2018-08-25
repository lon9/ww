package viewmanagers

import (
	"github.com/jroimartin/gocui"
)

// MainView is struct for main view
type MainView struct {
	BaseView
}

// NewMainView is constructor
func NewMainView(name string, editable bool) *MainView {
	return &MainView{
		*NewBaseView(name, editable),
	}
}

// Layout is interface method
func (w *MainView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(w.Name, maxX/5-1, 0, maxX*4/5-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Editable = w.Editable
	}

	return nil
}
