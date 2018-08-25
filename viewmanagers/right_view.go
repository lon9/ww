package viewmanagers

import (
	"github.com/jroimartin/gocui"
)

// RightView is struct for right view
type RightView struct {
	BaseView
}

// NewRightView is constructor
func NewRightView(name string, editable bool) *RightView {
	return &RightView{
		*NewBaseView(name, editable),
	}
}

// Layout is interface method
func (w *RightView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(w.Name, maxX*4/5-1, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Memo"
		v.Wrap = true
		v.Editable = w.Editable
	}

	return nil
}
