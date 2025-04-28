package algotable

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/pj-pj-pj/page-replacement-algorithms/algorithms"
)

// tables to show the process of the algorithms
var TableStringsFifo = tview.NewTable()
var TableFramesFifo = tview.NewTable().SetBorders(true)

var TableStringsLru = tview.NewTable()
var TableFramesLru = tview.NewTable().SetBorders(true)

var TableStringsOpt = tview.NewTable()
var TableFramesOpt = tview.NewTable().SetBorders(true)

// format frames to use square brackets like this: [ ][ ]
// and have different colors as well for the replaced frame
func FormatFrames(frames []int, framesLength int, currentPage int, isPageFault bool) string {
	result := ""
	for i := 0; i < framesLength; i++ {
		if i < len(frames) {
			if frames[i] == currentPage && isPageFault {
				result += fmt.Sprintf("[mediumspringgreen][%d]", frames[i])
			} else {
				result += fmt.Sprintf("[white][%d]", frames[i])
			}
		} else {
			result += "[gray][ ]"
		}
	}
	return strings.TrimSpace(result)
}

// this function puts the contents inside the table
func PopulateTable(prs []int, frames int, result []algorithms.PageStep) {
	var rows = len(prs) + 1 // number of rows depends on the range selected + 1, because step 0 exists
	
	// clear the table every time frames and prs range changes to
	// remove previous data 
	Table.Clear()
	FaultsTable.Clear()

	// ----- faults table starts here

	FaultsTable.SetCell(0, 0,
			tview.NewTableCell(" Total Page Faults ").
					SetAlign(tview.AlignCenter))

	FaultsTable.SetCell(0, 1,
		tview.NewTableCell(fmt.Sprintf(" %d ", result[len(result) - 1].FaultsCount)).
				SetTextColor(tcell.ColorLimeGreen).
				SetAlign(tview.AlignCenter))


	// clear the table every time frames and prs range changes to
	// remove previous data 
	TableStringsFifo.Clear()
	TableFramesFifo.Clear()
	TableStringsLru.Clear()
	TableFramesLru.Clear()
	TableStringsOpt.Clear()
	TableFramesOpt.Clear()


	// fifo table
	for c, val := range prs {
    TableStringsFifo.SetCell(0, c,
        tview.NewTableCell(fmt.Sprintf("  %d", val)).
            SetAlign(tview.AlignCenter))}

	for i, val := range fifoResult {
		prevFaultsCount := 0
		if i > 0 {
				prevFaultsCount = fifoResult[i-1].FaultsCount
		}

		// If current faultsCount did not increase, skip filling this step
		for j, framesVal := range val.Frames {
			if val.FaultsCount == prevFaultsCount {
				TableFramesFifo.SetCell(j, i,
					tview.NewTableCell(fmt.Sprintf(" [darkslategray]%d ", framesVal)).
							SetAlign(tview.AlignCenter))
			} else if val.Page == framesVal && val.PageFault {
				TableFramesFifo.SetCell(j, i,
						tview.NewTableCell(fmt.Sprintf(" [darkorange]%d ", framesVal)).
								SetAlign(tview.AlignCenter))
			} else {
				TableFramesFifo.SetCell(j, i,
						tview.NewTableCell(fmt.Sprintf(" %d ", framesVal)).
								SetAlign(tview.AlignCenter))
			}
		}
	}

	// steps --- starts from 0
	for steps := range rows {
		Table.SetCell(steps +1, 0,
			tview.NewTableCell(fmt.Sprintf("%d", steps)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter))
	}

	// page column
	for i, page := range prs {
		Table.SetCell(i +2, 1,
			tview.NewTableCell(fmt.Sprintf("%d", page)).
				SetTextColor(tcell.ColorOrange).
				SetAlign(tview.AlignCenter))
	}

	// put empty frames at step 0
	Table.SetCell(1, 2,
		tview.NewTableCell((strings.Repeat("[gray][ ]", frames))).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter))

	
	for f, currentStep := range result {
		Table.SetCell(f +2, 2,
			tview.NewTableCell(FormatFrames(currentStep.Frames, frames, currentStep.Page, currentStep.PageFault)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter))
	}

	for u, currentStep := range result {
		if result[u].PageFault {
			Table.SetCell(u +2, 3,
				tview.NewTableCell(fmt.Sprintf("[mediumspringgreen]%v", currentStep.PageFault)).
					SetTextColor(tcell.ColorWhite).
					SetAlign(tview.AlignCenter))
		} else {
			Table.SetCell(u +2, 3,
				tview.NewTableCell(fmt.Sprintf("%v", currentStep.PageFault)).
					SetTextColor(tcell.ColorWhite).
					SetAlign(tview.AlignCenter))
		}
	}

	for g, currentStep := range result {
		Table.SetCell(g +2, 4,
			tview.NewTableCell(fmt.Sprintf("%v", currentStep.FaultsCount)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter))
	}

	// added cool navigation stuff i found on the internet
	Table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			Table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		Table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		Table.SetSelectable(false, false)
	})
}


	

