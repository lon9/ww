package viewmanagers

import (
	"github.com/jroimartin/gocui"
)

type MainView struct {
	BaseView
}

func NewMainView(name string, editable bool) *MainView {
	return &MainView{
		*NewBaseView(name, editable),
	}
}

func (w *MainView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(w.Name, maxX/5-1, 0, maxX*4/5-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = w.Name
		v.Wrap = true
		v.Editable = w.Editable
	}

	return nil
}
