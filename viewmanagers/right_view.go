package viewmanagers

import (
	"github.com/jroimartin/gocui"
)

type RightView struct {
	BaseView
}

func NewRightView(name string, editable bool) *RightView {
	return &RightView{
		*NewBaseView(name, editable),
	}
}

func (w *RightView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(w.Name, maxX*4/5-1, 0, maxX-1, maxY-1)
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
