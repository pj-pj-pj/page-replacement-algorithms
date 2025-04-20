package ui

import (
	"fmt"

	"Image/jpeg"
	"bytes"
	"encoding/base64"

	"math/rand"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/pj-pj-pj/page-replacement-algorithms/assets"
	"github.com/pj-pj-pj/page-replacement-algorithms/ui/algotable"

	"github.com/pj-pj-pj/page-replacement-algorithms/algorithms"
)

// generates prs based on the length given
func generatePageRefString(length int) []int {
	result := make([]int, length) // make is like array but it has dynamic size, it makes "slice"

	for i := 0; i < length; i++ {
		result[i] = rand.Intn(9) // this thing returns integers from 0 to 9 (inclusive)
	}

	return result
}

// green for hackerist vibes
var primaryColor = tcell.ColorLimeGreen

// insert Image for hacker vibes
var Image = tview.NewImage().SetImage(photo) 
var b, _ = base64.StdEncoding.DecodeString(assets.Hackerist) 
var photo, _ = jpeg.Decode(bytes.NewReader(b))



// this is to make creating texts easier, just use newMainText("blablabla")
// instead of writing very long stuff
// i might create non-centered text later so this is named "main" text
var NewMainText = func(text string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text).SetTextColor(tcell.Color153) // blue
}

var NewText = func(text string) *tview.TextView {
	return tview.NewTextView().SetText(text).
	SetTextColor(primaryColor)
}




// ----- algo values that menu needs to access so i need to initialize them here at the top

var selectedAlgo = "First-In, First-Out (FIFO)"
var selectedAlgoDisplay = NewText("Algorithm: \t" + selectedAlgo)

var selectedFrames = 3
var selectedFramesDisplay = NewText(fmt.Sprintf("Frames: \t%d", selectedFrames))

var selectedRange = 9
var selectedRangeDisplay = NewText(fmt.Sprintf("PRS Range: \t0 - %d", selectedRange))

var prs = generatePageRefString(9)
var generatedPageReferenceString = NewText(fmt.Sprint("Generated String: \t", prs))

// ----- algo values end here


// --------------------------- menu grid and lists ui stuff (starts here)

// this is title text
var menu = NewMainText("\nMenu")

// this will appear on menu and users can select which page-replacement-algo to use
var algoType = tview.NewList().
	ShowSecondaryText(false).
	AddItem("First-In, First-Out (FIFO)", "", 0, func() { redrawAlgoType(2, "First-In, First-Out (FIFO)") }).
	AddItem("Least Recently Used (LRU)", "", 0, func() { redrawAlgoType(8, "Least Recently Used (LRU)") }).
	AddItem("Optimal Algorithm (OPT)", "", 0, func() { redrawAlgoType(256, "Optimal Algorithm (OPT)") })

// options of number of frames for users to choose from
var frames = tview.NewList().
	ShowSecondaryText(false).
	AddItem("3 (Default)", "", 0, func() { redrawFrames(4, 3) }).
	AddItem("9", "", 0, func() { redrawFrames(0, 9) }).
	AddItem("15", "", 0, func() { redrawFrames(256, 15) }).
	AddItem("18", "", 0, func() {  redrawFrames(1, 18) })

	// options of prs ranges for users to choose from
var pageRefString *tview.List = tview.NewList().
	ShowSecondaryText(false).
	AddItem("9 pages (Default)", "", 0, func() { redrawPRS(0, 9) }).
	AddItem("16 pages", "", 0, func() {  redrawPRS(9, 16) }).
	AddItem("20 pages", "", 0, func() {  redrawPRS(1, 20) })

// set up list's title and borders
// somehow i cant do it when i initialize them so here it is
func SetUpLists() {
	frames.SetTitle(" Number of Frames ").SetBorder(true)
	pageRefString.SetTitle(" Length of Page Reference String (0–9) ").SetBorder(true)
	algoType.SetTitle(" Algorithms ").SetBorder(true)

	// out of place buut it needs to be setup along with the lists
	// this sets up the table
	algotable.PopulateTable(prs, selectedFrames, algorithms.Fifo(prs, selectedFrames))
}

// these functions update ui
// i put here the repetitive functions that is used
// whenever a selection is made from the list because they are now very long and just repetitive
func redrawAlgoType(imgNumber int, algo string) {
	Image.SetColors(imgNumber)

	selectedAlgo = algo
	selectedAlgoDisplay.SetText("Algorithm: \t" + selectedAlgo)

	switch selectedAlgo {
	case "First-In, First-Out (FIFO)" :
		algotable.PopulateTable(prs, selectedFrames, algorithms.Fifo(prs, selectedFrames))
	case "Least Recently Used (LRU)" :
		algotable.PopulateTable(prs, selectedFrames, algorithms.Lru(prs, selectedFrames))
	case "Optimal Algorithm (OPT)" :
		algotable.PopulateTable(prs, selectedFrames, algorithms.Opt(prs, selectedFrames))
}
}

func redrawFrames(imgNumber int, frames int) {
	Image.SetColors(imgNumber)
		
	selectedFrames = frames
	selectedFramesDisplay.SetText(fmt.Sprintf("Frames: \t%d", selectedFrames))

	switch selectedAlgo {
	case "First-In, First-Out (FIFO)" :
		algotable.PopulateTable(prs, selectedFrames, algorithms.Fifo(prs, selectedFrames))
	case "Least Recently Used (LRU)" :
		algotable.PopulateTable(prs, selectedFrames, algorithms.Lru(prs, selectedFrames))
	case "Optimal Algorithm (OPT)" :
		algotable.PopulateTable(prs, selectedFrames, algorithms.Opt(prs, selectedFrames))
}
}

func redrawPRS(imgNumber int, prsRange int) {
	Image.SetColors(imgNumber)
		
	selectedRange = prsRange
	selectedRangeDisplay.SetText(fmt.Sprintf("PRS Range: \t0 - %d", selectedRange))

	prs = generatePageRefString(selectedRange)
	generatedPageReferenceString.SetText(fmt.Sprint("Generated String: \t", prs))

	switch selectedAlgo {
		case "First-In, First-Out (FIFO)" :
			algotable.PopulateTable(prs, selectedFrames, algorithms.Fifo(prs, selectedFrames))
		case "Least Recently Used (LRU)" :
			algotable.PopulateTable(prs, selectedFrames, algorithms.Lru(prs, selectedFrames))
		case "Optimal Algorithm (OPT)" :
			algotable.PopulateTable(prs, selectedFrames, algorithms.Opt(prs, selectedFrames))
	}
}

// we put lists on an array of type boxes
// this is just to put in a for loop that can be seen on main
// the loop puts keyboard navigation on the list
// i just put it here to lessen more code on main as much as possible
var Selections = []*tview.Box{algoType.Box, frames.Box, pageRefString.Box}

// grid that holds the lists of options together
var MenuGrid = tview.NewGrid().
	SetBorders(false).
	SetColumns(2, 0, 2).
	SetRows(3, 0, 1, 0, 1, 0, 6).
	AddItem(menu, 0, 1, 1, 1, 0, 0, true).
	AddItem(algoType, 1, 1, 1, 1, 0, 0, true).
	AddItem(NewMainText(""), 2, 1, 1, 1, 0, 0, true).
	AddItem(frames, 3, 1, 1, 1, 0, 0, true).
	AddItem(NewMainText(""), 4, 1, 1, 1, 0, 0, true).
	AddItem(pageRefString, 5, 1, 1, 1, 0, 0, true).
	AddItem(NewText("\n[!] Navigation: Use Mouse Keys or\n\tArrow keys [↑, ↓] to change option,\n\t[Tab] to switch lists,\n\t[Enter] key to select option\n"), 6, 1, 1, 1, 0, 0, true)

// --------------------------- menu grid and lists (ends here)




// --------------------------- algorithm panel starts here

var tableInfo1 = tview.NewTextView().SetText("[!]\tIf PRS Range is 0-16 or 0-20, use [Mouse Scroll] or Arrow keys [↑, ↓] to move the table when some rows become hidden").SetTextColor(tcell.ColorRed)

var AlgoGrid = tview.NewGrid().
	SetBorders(false).
	SetColumns(2, 6, 0, 0, 0).
	SetRows(1, 1, 1, 1, 2, 2, 1, 0, 3).
	AddItem(NewText(""), 0, 1, 1, 4, 0, 0, true).
	AddItem(selectedAlgoDisplay, 1, 1, 1, 4, 0, 0, true).
	AddItem(selectedFramesDisplay, 2, 1, 1, 4, 0, 0, true).
	AddItem(selectedRangeDisplay, 3, 1, 1, 4, 0, 0, true).
	AddItem(generatedPageReferenceString, 4, 1, 1, 5, 0, 0, true).
	AddItem(algotable.Table, 6, 2, 3, 5, 0, 0, true).
	AddItem(tableInfo1, 5, 2, 1, 3, 0, 0, true).
	AddItem(algotable.FaultsTable, 1, 6, 5, 2, 0, 0, true)

// --------------------------- algorithm panel ends here