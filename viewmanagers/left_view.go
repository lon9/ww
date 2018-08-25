package viewmanagers

import (
	"github.com/jroimartin/gocui"
)

// LeftView is struct for view of left
type LeftView struct {
	BaseView
}

// NewLeftView is constructor
func NewLeftView(name string, editable bool) *LeftView {
	return &LeftView{
		*NewBaseView(name, editable),
	}
}

// Layout is interface method
func (w *LeftView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(w.Name, 0, 0, maxX/5-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Info"
		v.Wrap = true
		v.Editable = w.Editable
	}

	return nil
}
