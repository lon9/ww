package viewmanagers

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// CursorDown move the cursor down
func CursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

// CursorUp move the cursor up
func CursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetLineIndex returns index of the cursor
func GetLineIndex(v *gocui.View) int {
	_, cy := v.Cursor()
	return cy
}

// DrawDeadView draws view for dead
func DrawDeadView(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		mainView, err := g.View(MainViewID)
		if err != nil {
			return err
		}
		mainView.Clear()
		fmt.Fprintln(mainView, "You're already dead")
		return nil
	})
}
