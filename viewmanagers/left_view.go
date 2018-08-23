package viewmanagers

import (
	"github.com/jroimartin/gocui"
)

type LeftView struct {
	BaseView
}

func NewLeftView(name string, editable bool) *LeftView {
	return &LeftView{
		*NewBaseView(name, editable),
	}
}

func (w *LeftView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(w.Name, 0, 0, maxX/5-1, maxY-1)
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
