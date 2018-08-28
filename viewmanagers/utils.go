package viewmanagers

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// CursorDown moves the cursor down
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

// CursorUp moves the cursor up
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

// CursorDownWithRange moves the cursor with range
func CursorDownWithRange(v *gocui.View, bottom int) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy == bottom {
			return nil
		}
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}

	}
	return nil
}

// CursorUpWithRange moves the cursor with range
func CursorUpWithRange(v *gocui.View, top int) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy == top {
			return nil
		}
		ox, oy := v.Origin()
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
func DrawDeadView(g *gocui.Gui, viewID string) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(viewID)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, "You're already dead")
		return nil
	})
}
