package algotable

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/pj-pj-pj/page-replacement-algorithms/algorithms"
)

// table to show the process of the selected algorithms

var Table = tview.NewTable().SetBorders(true)
var tableHeaders = strings.Split("Step_Page_Frames Content_Page Fault?_Faults Count", "_")

var FaultsTable = tview.NewTable().SetBorders(true)

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

	// faults table ends here ----------

	// ---- algorithm table starts here

	// table headers
	for c, header := range tableHeaders {
    Table.SetCell(0, c,
        tview.NewTableCell(header).
            SetTextColor(tcell.ColorOrange).
            SetAlign(tview.AlignCenter))
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


	

