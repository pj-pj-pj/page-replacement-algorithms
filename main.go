package main

import (
	// these two packages provide stuff to create
	// rich interactive terminal program for go lang woooo
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	// a lot of code comes from here but they are mostly ui stuff
	// i put them in a module to reduce clutter here
	// but this module contains the function for generating the page reference string
	"github.com/pj-pj-pj/page-replacement-algorithms/ui"

	"github.com/pj-pj-pj/page-replacement-algorithms/ui/algotable"
)

func main() {
	// used to initialize tview, 
	// tview is a library for creating rich interactive terminal programs which is what im trying to create here
	app := tview.NewApplication()

	// this is the main grid of the program
	grid := tview.NewGrid().
		SetRows(1, 0, 1).
		SetColumns(40, 0, 0, 66).SetBorders(true).
		AddItem(ui.NewMainText("╰☆☆ Ｐａｇｅ Ｒｅｐｌａｃｅｍｅｎｔ Ａｌｇｏｒｉｔｈｍｓ ☆☆╮"), 0, 0, 1, 3, 0, 0, false). // stars for extravagance
		AddItem(algotable.AlgoFaults, 0, 3, 1, 1, 0, 0, false).
		AddItem(ui.NewMainText("--- please fullscreen for better experience ---"), 2, 0, 1, 2, 0, 0, false).
		AddItem(ui.NewMainText("╰☆☆ Ｐａｇｅ Ｒｅｐｌａｃｅｍｅｎｔ Ａｌｇｏｒｉｔｈｍｓ ☆☆╮"), 2, 2, 1, 2, 0, 0, false)

	// func (g *Grid) AddItem(p Primitive, row, column, rowSpan, colSpan, minGridHeight, minGridWidth int, focus bool) *Grid 
	// ----> for guide to know how to layout items

	// Layout for screens narrower than 150 cells.
	grid. // image removed for v1.1
		AddItem(ui.MenuGrid, 1, 0, 1, 1, 0, 0, true).
		AddItem(ui.AlgoGrid, 1, 1, 1, 3, 0, 0, true)

	// this is for setting up the titles and borders of the list
	// idk why but if I dont set them up separately and
	// initialize the title and border, their type becomes tview.Box and not tview.List
	// so separate setup
	ui.SetUpLists()

	// to enable navigation with arrows and tab on menu panel
	// i got this code from tview documentation and applied it here
	// link: https://github.com/rivo/tview/wiki/Image
	for i, box := range ui.Selections {
		(func(index int) {
			box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				switch event.Key() {
				case tcell.KeyTab:
					app.SetFocus(ui.Selections[(index+1)%len(ui.Selections)])
					return nil
				case tcell.KeyBacktab:
					app.SetFocus(ui.Selections[(index+len(ui.Selections)-1)%len(ui.Selections)])
					return nil
				}
				return event
			}).
			SetBorderColor(tcell.ColorCornflowerBlue)
			// SetBackgroundColor(tcell.Color153)
		})(i)
	}

	// first, setting the app root to grid, which means show grid to the users
	// setting focus on selection because that is the menu
	// enabling mouse if they get confused with keyboard navigation
	if err := app.SetRoot(grid, true).SetFocus(ui.Selections[0]).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
