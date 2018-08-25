package viewmanagers

import (
	"github.com/jroimartin/gocui"
)

// DialogView is view for dialog
type DialogView struct {
	BaseView
}

// NewDialogView is constructor
func NewDialogView(name string, editable bool) *DialogView {
	return &DialogView{
		*NewBaseView(name, editable),
	}
}

// Layout is interface Method
func (w *DialogView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(w.Name, maxX/2-15, maxY/2, maxX/2+15, maxY/2+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Editable = w.Editable
	}
	return nil
}
