package main

import (
	"bytes"
	"encoding/base64"

	"image/jpeg"

	"github.com/gdamore/tcell/v2"
	"github.com/pj-pj-pj/page-replacement-algorithms/assets"
	"github.com/rivo/tview"
)


func main() {
	// used to initialize tview, 
	// tview is a library for creating rich interactive terminal programs which is what im trying to create here
	app := tview.NewApplication()

	// green for hackerist vibes
	primaryColor := tcell.ColorForestGreen

	image := tview.NewImage() // insert image for hacker vibes
	b, _ := base64.StdEncoding.DecodeString(assets.Hackerist) 
	photo, _ := jpeg.Decode(bytes.NewReader(b))
	image.SetImage(photo)

	// this is to make creating texts easier, just use newMainText("blablabla")
	// instead of writing very long stuff
	// i might create non-centered text later so this is named "main" text
	newMainText := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text).SetTextColor(primaryColor) // green
	}

	newText := func(text string) tview.Primitive {
		return tview.NewTextView().SetText(text).
		SetTextColor(tcell.ColorRed)
	}


	// --------------------------- menu grid and lists (starts here)
	// TODO: Create a module for menu later to reduce clutter on main.go

	menu := newMainText("\nMenu")

	// this will appear on menu and users can select which page-replacement-algo to use
	algoType := tview.NewList().
		ShowSecondaryText(false).
		AddItem("First-In, First-Out (FIFO)", "", 0, func() { image.SetColors(2) }).
		AddItem("Least Recently Used (LRU)", "", 0, func() { image.SetColors(8) }).
		AddItem("Optimal Algorithm", "", 0, func() { image.SetColors(tview.TrueColor) })
		// AddItem("256 colors", "", 0, func() { image.SetColors(256) })
	algoType.SetTitle(" Algorithms ").SetBorder(true)

	frames := tview.NewList().
		ShowSecondaryText(false).
		AddItem("3 (Default)", "", 0, func() { }).
		AddItem("9", "", 0, func() { }).
		AddItem("15", "", 0, func() { })
	frames.SetTitle(" Number of Frames ").SetBorder(true)

	pageRefString := tview.NewList().
		ShowSecondaryText(false).
		AddItem("0 - 9 (Default)", "", 0, func() { }).
		AddItem("0 - 16", "", 0, func() { }).
		AddItem("0 - 21", "", 0, func() { })
	pageRefString.SetTitle(" Page Reference String Range ").SetBorder(true)

	menuGrid := tview.NewGrid().
		SetBorders(false).
		SetRows(3, 0, 1, 0, 1, 0, 5).
		AddItem(menu, 0, 0, 1, 1, 0, 0, true).
		AddItem(algoType, 1, 0, 1, 1, 0, 0, true).
		AddItem(newMainText(""), 2, 0, 1, 1, 0, 0, true).
		AddItem(frames, 3, 0, 1, 1, 0, 0, true).
		AddItem(newMainText(""), 4, 0, 1, 1, 0, 0, true).
		AddItem(pageRefString, 5, 0, 1, 1, 0, 0, true).
		AddItem(newText("\n[!] Navigation: [Tab] to switch lists,\n\tarrow keys [↑, ↓] to change option,\n\t[Enter] key to select option\n"), 6, 0, 1, 1, 0, 0, true)

	selections := []*tview.Box{algoType.Box, frames.Box, pageRefString.Box}
	for i, box := range selections {
		(func(index int) {
			box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				switch event.Key() {
				case tcell.KeyTab:
					app.SetFocus(selections[(index+1)%len(selections)])
					return nil
				case tcell.KeyBacktab:
					app.SetFocus(selections[(index+len(selections)-1)%len(selections)])
					return nil
				}
				return event
			})
		})(i)
	}

	// --------------------------- menu grid and lists (ends here)

	// --------------------------- algorithm panel starts here

	algo := newMainText("\nAlgorithm")
	

	// --------------------------- algorithm panel ends here


	grid := tview.NewGrid().
		SetRows(3, 0, 5).
		SetColumns(40, 0, 0, 0).SetBorders(true).
		AddItem(newMainText("･✧･ﾟ*\t･ﾟ･･ﾟ✧\t･ﾟ*･ﾟ\n:･ﾟ✧･ﾟ* page-replacement-algorithms･✧･ﾟ*✧･ﾟ:･ﾟ\n✧･ﾟ✧･ﾟ✧･ﾟ:･ﾟ\t･ﾟ･ﾟ"), 0, 0, 1, 4, 0, 0, false). // stars for extravagance
		AddItem(newMainText("\npaula-joyce-ucol\nbscs-3b\n--- please fullscreen for better experience ---"), 2, 2, 1, 2, 0, 0, false).
		AddItem(newMainText("\nPress [Q] to exit\n\nPress [Ctrl + C] to force exit"), 2, 0, 1, 2, 0, 0, false)

	// func (g *Grid) AddItem(p Primitive, row, column, rowSpan, colSpan, minGridHeight, minGridWidth int, focus bool) *Grid 
	// ----> for guide to know how to layout items

	// Layout for screens narrower than 150 cells.
	grid.AddItem(image, 0, 0, 0, 0, 0, 0, false).
		AddItem(menuGrid, 1, 0, 1, 1, 0, 0, false).
		AddItem(algo, 1, 1, 1, 3, 0, 0, false)

	// Layout for screens wider than 150 cells.
	grid.AddItem(image, 1, 0, 1, 1, 0, 150, false).
		AddItem(menuGrid, 1, 1, 1, 1, 0, 150, false).
		AddItem(algo, 1, 2, 1, 2, 0, 150, false)


	if err := app.SetRoot(grid, true).SetFocus(selections[0]).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}