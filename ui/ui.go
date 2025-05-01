package ui

import (
	"fmt"
	"image"

	"Image/gif"
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
func generatePageRefString(length int) []int {
	result := make([]int, length) // make is like array but it has dynamic size, it makes "slice"

	for i := 0; i < length; i++ {
		result[i] = rand.Intn(9) // this thing returns integers from 0 to 9 (inclusive)
	}

	return result
}

// green for hackerist vibes
var primaryColor = tcell.ColorLimeGreen

var a, _ = base64.StdEncoding.DecodeString(assets.Juice)
var b, _ = base64.StdEncoding.DecodeString(assets.Comp)
var c, _ = base64.StdEncoding.DecodeString(assets.Moonlight)
var d, _ = base64.StdEncoding.DecodeString(assets.Coffee)
var e, _ = base64.StdEncoding.DecodeString(assets.Banana)

var juicebox, _ = jpeg.Decode(bytes.NewReader(a))
var computer, _ = gif.Decode(bytes.NewReader(b))
var moon, _ = gif.Decode(bytes.NewReader(c))
var coffee, _ = gif.Decode(bytes.NewReader(d))
var banana, _ = gif.Decode(bytes.NewReader(e))

var imgList = []image.Image{juicebox, computer, moon, coffee, banana}

var Image = tview.NewImage().SetImage(computer).SetColors(256)




// this is to make creating texts easier, just use newMainText("blablabla")
// instead of writing very long stuff
// i might create non-centered text later so this is named "main" text
var NewMainText = func(text string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text).SetTextColor(primaryColor)
}

var NewText = func(text string) *tview.TextView {
	return tview.NewTextView().SetText(text).
	SetTextColor(primaryColor)
}




// ----- algo values that menu needs to access so i need to initialize them here at the top

var selectedFrames = 3

var selectedRange = 9

var prs = generatePageRefString(9)
var generatedPageReferenceString = NewText(fmt.Sprint("Generated Pages: \n", prs))

// ----- algo values end here




// --------------------------- menu grid and lists ui stuff (starts here)

// this is title text
var menu = NewMainText("\nMenu")

// this will appear on menu and users can select which to use
// options of number of frames for users to choose from
var frames = tview.NewList().
	ShowSecondaryText(false).
	AddItem("3 frames (Default)", "", 0, func() { redrawFrames(3) }).
	AddItem("1 frame", "", 0, func() { redrawFrames(1) }).
	AddItem("2 frames", "", 0, func() { redrawFrames(2) }).
	AddItem("4 frames", "", 0, func() {  redrawFrames(4) }).
	AddItem("5 frames", "", 0, func() {  redrawFrames(5) })

	// options of prs ranges for users to choose from
var pageRefString *tview.List = tview.NewList().
	ShowSecondaryText(false).
	AddItem("9 pages (Default)", "", 0, func() { redrawPRS(9) }).
	AddItem("14 pages", "", 0, func() {  redrawPRS(19) }).
	AddItem("19 pages", "", 0, func() {  redrawPRS(19) }).
	AddItem("28 pages", "", 0, func() {  redrawPRS(30) }).
	AddItem("39 pages", "", 0, func() {  redrawPRS(39) })




// set up list's title and borders
// somehow i cant do it when i initialize them so here it is
func SetUpLists() {
	frames.SetTitle(" Number of Frames ").SetBorder(true)
	pageRefString.SetTitle(" Generate Pages (0–9) ").SetBorder(true)
	// algoType.SetTitle(" Algorithms ").SetBorder(true)

	// out of place buut it needs to be setup along with the lists
	// this sets up the table
	algotable.PopulateTable(prs, selectedFrames)
}


// these functions update ui
// i put here the repetitive functions that is used
// whenever a selection is made from the list because they are now very long and just repetitive
func redrawPRS( prsRange int) {
	Image.SetImage(imgList[rand.Intn(4)])
	selectedRange = prsRange

	prs = generatePageRefString(selectedRange)
	generatedPageReferenceString.SetText(fmt.Sprint("Generated Pages: \n", prs))

	algotable.PopulateTable(prs, selectedFrames)
}

func redrawFrames(frames int) {
	Image.SetImage(imgList[rand.Intn(4)])
	selectedFrames = frames

	algotable.PopulateTable(prs, selectedFrames)
}



// we put lists on an array of type boxes
// this is just to put in a for loop that can be seen on main
// the loop puts keyboard navigation on the list
// i just put it here to lessen more code on main as much as possible
var Selections = []*tview.Box{pageRefString.Box, frames.Box}

// grid that holds the lists of options together
var MenuGrid = tview.NewGrid().
	SetBorders(false).
	SetColumns(2, 0, 2).
	SetRows(3, 7, 4, 1, 7, 1, 6, 0, 1).
	AddItem(menu, 0, 1, 1, 1, 0, 0, true).
	AddItem(pageRefString, 1, 1, 1, 1, 0, 0, true).
	AddItem(generatedPageReferenceString, 2, 1, 1, 1, 0, 0, true).
	AddItem(NewMainText(""), 3, 1, 1, 1, 0, 0, true).
	AddItem(frames, 4, 1, 1, 1, 0, 0, true).
	AddItem(NewMainText(""), 5, 1, 1, 1, 0, 0, true).
	AddItem(NewText("[!] Navigation: Use Mouse Keys or Arrow keys [↑, ↓] to change option, [Tab] to switch lists, and [Enter] key to select option\n").SetTextColor(tcell.ColorRed), 6, 1, 1, 1, 0, 0, true).
	AddItem(Image, 7, 1, 1, 1, 0, 0, true)

// --------------------------- menu grid and lists (ends here)




// --------------------------- algorithm panel starts here


var AlgoGrid = tview.NewGrid().
	SetBorders(false).
	SetColumns(2, 2, 0).
	SetRows(1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0).
	AddItem(algotable.FifoFaultsText.SetTextColor(tcell.ColorDarkOrange), 1, 2, 1, 1, 0, 0, true).
	AddItem(algotable.TableStringsFifo, 3, 1, 1, 2, 0, 0, true).
	AddItem(algotable.TableFramesFifo, 4, 1, 1, 2, 0, 0, true).
	AddItem(algotable.LruFaultsText.SetTextColor(tcell.ColorGreenYellow), 6, 2, 1, 1, 0, 0, true).
	AddItem(algotable.TableStringsLru, 7, 1, 1, 2, 0, 0, true).
	AddItem(algotable.TableFramesLru, 8, 1, 1, 2, 0, 0, true).
	AddItem(algotable.OptFaultsText.SetTextColor(tcell.ColorDodgerBlue), 9, 2, 1, 1, 0, 0, true).
	AddItem(algotable.TableStringsOpt, 11, 1, 1, 2, 0, 0, true).
	AddItem(algotable.TableFramesOpt, 12, 1, 1, 2, 0, 0, true)


// --------------------------- algorithm panel ends here