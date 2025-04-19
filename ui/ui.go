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
)

// generates prs based on the length given
// it says length but that is also range
func generatePageRefString(length int) []int {
	result := make([]int, length) // make is like array but it has dynamic size

	for i := 0; i < length; i++ {
		result[i] = rand.Intn(length) // this thing returns integers from 0 to 9 (inclusive)
	}

	return result
}



// green for hackerist vibes
var primaryColor = tcell.ColorLimeGreen

var Image = tview.NewImage().SetImage(photo) // insert Image for hacker vibes
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
	AddItem("First-In, First-Out (FIFO)", "", 0, func() { 
		Image.SetColors(2)

		selectedAlgo = "First-In, First-Out (FIFO)"
		selectedAlgoDisplay.SetText("Algorithm: \t" + selectedAlgo)
	}).
	AddItem("Least Recently Used (LRU)", "", 0, func() { 
		Image.SetColors(8) 

		selectedAlgo = "Least Recently Used (LRU)"
		selectedAlgoDisplay.SetText("Algorithm: \t" + selectedAlgo)
	}).
	AddItem("Optimal Algorithm (OPT)", "", 0, func() { 
		Image.SetColors(tview.TrueColor) 

		selectedAlgo = "Optimal Algorithm (OPT)"
		selectedAlgoDisplay.SetText("Algorithm: \t" + selectedAlgo)
	})

// options of number of frames for users to choose from
var frames = tview.NewList().
	ShowSecondaryText(false).
	AddItem("3 (Default)", "", 0, func() {
		Image.SetColors(1)
		
		selectedFrames = 3
		selectedFramesDisplay.SetText(fmt.Sprintf("Frames: \t%d", selectedFrames))

		algotable.PopulateTable(selectedRange, prs, selectedFrames)
	}).
	AddItem("9", "", 0, func() {
		Image.SetColors(10)
		
		selectedFrames = 9
		selectedFramesDisplay.SetText(fmt.Sprintf("Frames: \t%d", selectedFrames))

		algotable.PopulateTable(selectedRange, prs, selectedFrames)
	}).
	AddItem("15", "", 0, func() {
		Image.SetColors(4)
		
		selectedFrames = 15
		selectedFramesDisplay.SetText(fmt.Sprintf("Frames: \t%d", selectedFrames))

		algotable.PopulateTable(selectedRange, prs, selectedFrames)
	}).
	AddItem("18", "", 0, func() {
		Image.SetColors(1)
		
		selectedFrames = 18
		selectedFramesDisplay.SetText(fmt.Sprintf("Frames: \t%d", selectedFrames))

		algotable.PopulateTable(selectedRange, prs, selectedFrames)
	})

	// options of prs ranges for users to choose from
var pageRefString *tview.List = tview.NewList().
	ShowSecondaryText(false).
	AddItem("0 - 9 (Default)", "", 0, func() {
		Image.SetColors(0)
		
		selectedRange = 9
		selectedRangeDisplay.SetText(fmt.Sprintf("PRS Range: \t0 - %d", selectedRange))

		prs = generatePageRefString(selectedRange)
		generatedPageReferenceString.SetText(fmt.Sprint("Generated String: \t", prs))

		algotable.PopulateTable(selectedRange, prs, selectedFrames)
	}).
	AddItem("0 - 16", "", 0, func() {
		Image.SetColors(9)
		
		selectedRange = 16
		selectedRangeDisplay.SetText(fmt.Sprintf("PRS Range: \t0 - %d", selectedRange))

		prs = generatePageRefString(selectedRange)
		generatedPageReferenceString.SetText(fmt.Sprint("Generated String: \t", prs))

		algotable.PopulateTable(selectedRange, prs, selectedFrames)
	}).
	AddItem("0 - 20", "", 0, func() {
		Image.SetColors(1)
		
		selectedRange = 20
		selectedRangeDisplay.SetText(fmt.Sprintf("PRS Range: \t0 - %d", selectedRange))

		prs = generatePageRefString(selectedRange)
		generatedPageReferenceString.SetText(fmt.Sprint("Generated String: \t", prs))

		algotable.PopulateTable(selectedRange, prs, selectedFrames)
	})

// set up list's title and borders
// somehow i cant do it when i initialize them so here it is
func SetUpLists() {
	frames.SetTitle(" Number of Frames ").SetBorder(true)
	pageRefString.SetTitle(" Page Reference String (PRS) Range ").SetBorder(true)
	algoType.SetTitle(" Algorithms ").SetBorder(true)

	algotable.PopulateTable(9, prs, selectedFrames)

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
	AddItem(NewText("\n[!] Navigation:\n\tArrow keys [↑, ↓] to change option,\n\t[Tab] to switch lists,\n\t[Enter] key to select option\n"), 6, 1, 1, 1, 0, 0, true)

// --------------------------- menu grid and lists (ends here)




// --------------------------- algorithm panel starts here

var tableInfo = NewText("\n☆☆☆ Page Faults: 6 ☆☆☆\n\n\n\n[!]\tWhen selected PRS Range is 0-19 or 0-20, use [Mouse Scroll] or Arrow keys [↑, ↓] to move the table when some cells become hidden\n\nShift focus on the table using [Right Mouse] key and press [Enter] to select table cells")

var AlgoGrid = tview.NewGrid().
	SetBorders(false).
	SetColumns(2, 25, 0, 0, 0).
	SetRows(1, 1, 1, 1, 2, 1, 0).
	AddItem(NewText(""), 0, 1, 1, 5, 0, 0, true).
	AddItem(selectedAlgoDisplay, 1, 1, 1, 4, 0, 0, true).
	AddItem(selectedFramesDisplay, 2, 1, 1, 4, 0, 0, true).
	AddItem(selectedRangeDisplay, 3, 1, 1, 4, 0, 0, true).
	AddItem(generatedPageReferenceString, 4, 1, 1, 4, 0, 0, true).
	AddItem(tableInfo, 5, 1, 2, 1, 0, 0, true).
	AddItem(algotable.Table, 5, 2, 2, 4, 0, 0, true)

// --------------------------- algorithm panel ends here