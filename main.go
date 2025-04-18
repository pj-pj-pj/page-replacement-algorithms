package main

import (
	"fmt"
	"math/rand"

	"bytes"
	"encoding/base64"

	"image/jpeg"

	"github.com/gdamore/tcell/v2"
	"github.com/pj-pj-pj/page-replacement-algorithms/assets"
	"github.com/rivo/tview"
)


func generateRange(length int) []int {
	result := make([]int, length)

	for i := 0; i < length; i++ {
		result[i] = rand.Intn(length) // this thing returns integers from 0 to 9 (inclusive)
	}

	return result
}

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
			SetText(text).SetTextColor(tcell.Color153) // blue
	}

	newText := func(text string) *tview.TextView {
		return tview.NewTextView().SetText(text).
		SetTextColor(primaryColor)
	}

	// ----- algo values that menu needs to access so i need to initialize them here at the top

	selectedAlgo := "First-In, First-Out (FIFO)"
	selectedAlgoDisplay := newText("Algorithm: \t\t\t" + selectedAlgo)

	selectedFrames := 3
	selectedFramesDisplay := newText(fmt.Sprintf("Frames: \t\t\t%d", selectedFrames))

	selectedRange := 9
	selectedRangeDisplay := newText(fmt.Sprintf("PRS Range: \t\t\t0 - %d", selectedRange))

	generatedPageReferenceString := newText(fmt.Sprint("Generated String: \t", generateRange(9)))

	// ----- algo values end here

	// --------------------------- menu grid and lists (starts here)
	// TODO: Create a module for menu later to reduce clutter on main.go

	menu := newMainText("\nMenu")

	// this will appear on menu and users can select which page-replacement-algo to use
	algoType := tview.NewList().
		ShowSecondaryText(false).
		AddItem("First-In, First-Out (FIFO)", "", 0, func() { 
			image.SetColors(2)

			selectedAlgo = "First-In, First-Out (FIFO)"
			selectedAlgoDisplay.SetText("Algorithm: \t\t\t" + selectedAlgo)
	  }).
		AddItem("Least Recently Used (LRU)", "", 0, func() { 
			image.SetColors(8) 

			selectedAlgo = "Least Recently Used (LRU)"
			selectedAlgoDisplay.SetText("Algorithm: \t\t\t" + selectedAlgo)
		}).
		AddItem("Optimal Algorithm (OPT)", "", 0, func() { 
			image.SetColors(tview.TrueColor) 

			selectedAlgo = "Optimal Algorithm (OPT)"
			selectedAlgoDisplay.SetText("Algorithm: \t\t\t" + selectedAlgo)
		})
		// AddItem("256 colors", "", 0, func() { image.SetColors(256) })
	algoType.SetTitle(" Algorithms ").SetBorder(true)

	frames := tview.NewList().
		ShowSecondaryText(false).
		AddItem("3 (Default)", "", 0, func() {
			image.SetColors(1)
			
			selectedFrames = 3
			selectedFramesDisplay.SetText(fmt.Sprintf("Frames: \t\t\t%d", selectedFrames))
		}).
		AddItem("9", "", 0, func() {
			image.SetColors(4)
			
			selectedFrames = 9
			selectedFramesDisplay.SetText(fmt.Sprintf("Frames: \t\t\t%d", selectedFrames))
		}).
		AddItem("15", "", 0, func() {
			image.SetColors(3)
			
			selectedFrames = 15
			selectedFramesDisplay.SetText(fmt.Sprintf("Frames: \t\t\t%d", selectedFrames))
		}).
		AddItem("18", "", 0, func() {
			image.SetColors(3)
			
			selectedFrames = 18
			selectedFramesDisplay.SetText(fmt.Sprintf("Frames: \t\t\t%d", selectedFrames))
		}).
		AddItem("25", "", 0, func() {
			image.SetColors(3)
			
			selectedFrames = 25
			selectedFramesDisplay.SetText(fmt.Sprintf("Frames: \t\t\t%d", selectedFrames))
		})
	frames.SetTitle(" Number of Frames ").SetBorder(true)

	pageRefString := tview.NewList().
		ShowSecondaryText(false).
		AddItem("0 - 9 (Default)", "", 0, func() {
			image.SetColors(8)
			
			selectedRange = 9
			selectedRangeDisplay.SetText(fmt.Sprintf("PRS Range: \t\t\t0 - %d", selectedRange))

			generatedPageReferenceString.SetText(fmt.Sprint("Generated String: \t", generateRange(9)))
		}).
		AddItem("0 - 16", "", 0, func() {
			image.SetColors(7)
			
			selectedRange = 16
			selectedRangeDisplay.SetText(fmt.Sprintf("PRS Range: \t\t\t0 - %d", selectedRange))

			generatedPageReferenceString.SetText(fmt.Sprint("Generated String: \t", generateRange(16)))
		}).
		AddItem("0 - 21", "", 0, func() {
			image.SetColors(122)
			
			selectedRange = 21
			selectedRangeDisplay.SetText(fmt.Sprintf("PRS Range: \t\t\t0 - %d", selectedRange))

			generatedPageReferenceString.SetText(fmt.Sprint("Generated String: \t", generateRange(21)))
		})
	pageRefString.SetTitle(" Page Reference String (PRS) Range ").SetBorder(true)

	menuGrid := tview.NewGrid().
		SetBorders(false).
		SetColumns(2, 0, 2).
		SetRows(3, 0, 1, 0, 1, 0, 6).
		AddItem(menu, 0, 1, 1, 1, 0, 0, true).
		AddItem(algoType, 1, 1, 1, 1, 0, 0, true).
		AddItem(newMainText(""), 2, 1, 1, 1, 0, 0, true).
		AddItem(frames, 3, 1, 1, 1, 0, 0, true).
		AddItem(newMainText(""), 4, 1, 1, 1, 0, 0, true).
		AddItem(pageRefString, 5, 1, 1, 1, 0, 0, true).
		AddItem(newText("\n[!] Navigation:\n\tArrow keys [↑, ↓] to change option,\n\t[Tab] to switch lists,\n\t[Enter] key to select option\n"), 6, 0, 1, 3, 0, 0, true)

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
			}).
			SetBorderColor(tcell.Color153)
			// SetBackgroundColor(tcell.Color153)
		})(i)
	}

	// --------------------------- menu grid and lists (ends here)

	// --------------------------- algorithm panel starts here

	algo := newMainText("\nAlgorithm\n")

	algoGrid := tview.NewGrid().
		SetBorders(false).
		SetColumns(3, 0, 3).
		SetRows(3, 1, 1, 1, 1, 1).
		AddItem(algo, 0, 1, 1, 1, 0, 0, true).
		AddItem(selectedAlgoDisplay, 1, 1, 1, 1, 0, 0, true).
		AddItem(newMainText(""), 2, 1, 1, 1, 0, 0, true).
		AddItem(selectedFramesDisplay, 3, 1, 1, 1, 0, 0, true).
		AddItem(selectedRangeDisplay, 4, 1, 1, 1, 0, 0, true).
		AddItem(generatedPageReferenceString, 5, 1, 1, 1, 0, 0, true)

	// --------------------------- algorithm panel ends here


	grid := tview.NewGrid().
		SetRows(3, 0, 5).
		SetColumns(40, 0, 0, 0).SetBorders(true).
		AddItem(newText("･*\t･ﾟ･･ﾟ\t･ﾟ*･ﾟ:･ﾟ･ﾟ✧ ･ﾟ✧･ﾟ*･ﾟ✧\t･ﾟ*･ﾟ  :･ﾟ✧\t･ﾟ*･ﾟ:･ﾟ･ﾟ*･ﾟ✧\t･ﾟ*･ﾟ  :･ﾟ✧   ･ﾟ*･ﾟ*･ﾟ･ﾟ\n:･ﾟ✧ ･･･ﾟ✧\t･ﾟ*･ﾟ  :･ﾟ✧ ✧\t･ﾟ*･ﾟ  :･ﾟ✧･ﾟ*･ﾟ･ﾟ*･ﾟ･ﾟ*･ﾟ*･ﾟpage-replacement-algorithms･✧✧\t･ﾟ*･ﾟ  :･ﾟ✧･ﾟ*･ﾟ*･ﾟ\n･ﾟ✧\t･ﾟ･ﾟ*✧･ﾟ:･ﾟ✧✧\t･ﾟ*･ﾟ  :･ﾟ✧   ･ﾟ*･ﾟ*･ﾟ･ﾟ✧\t･ﾟ*･ﾟ:✧･ﾟ✧･ﾟ✧･ﾟ:･ﾟ\t･ﾟ･ﾟ"), 0, 0, 1, 4, 0, 0, false). // stars for extravagance
		AddItem(newMainText("\npaula-joyce-ucol\nbscs-3b\n--- please fullscreen for better experience ---"), 2, 2, 1, 2, 0, 0, false).
		AddItem(newMainText("\nPress [Q] to exit\n\nPress [Ctrl + C] to force exit"), 2, 0, 1, 2, 0, 0, false)

	// func (g *Grid) AddItem(p Primitive, row, column, rowSpan, colSpan, minGridHeight, minGridWidth int, focus bool) *Grid 
	// ----> for guide to know how to layout items

	// Layout for screens narrower than 150 cells.
	grid.AddItem(image, 0, 0, 0, 0, 0, 0, false).
		AddItem(menuGrid, 1, 0, 1, 1, 0, 0, false).
		AddItem(algoGrid, 1, 1, 1, 3, 0, 0, false)

	// Layout for screens wider than 150 cells.
	grid.AddItem(image, 1, 0, 1, 1, 0, 150, false).
		AddItem(menuGrid, 1, 1, 1, 1, 0, 150, false).
		AddItem(algoGrid, 1, 2, 1, 2, 0, 150, false)


	if err := app.SetRoot(grid, true).SetFocus(selections[0]).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}