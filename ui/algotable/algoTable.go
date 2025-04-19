package algotable

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

	var Table = tview.NewTable().
		SetBorders(true)
	var tableHeaders = strings.Split("Step_Page_Frames Content_Page Fault?", "_")

	func PopulateTable(prsRange int, prs []int, frames int) {
		var cols = 4 // steps, page, frame contents, page fault?
		var rows = prsRange + 1
		Table.Clear()

		for c := 0; c < cols; c++ {
			Table.SetCell(0, c,
				tview.NewTableCell(tableHeaders[c]).
					SetTextColor(tcell.ColorOrange).
					SetAlign(tview.AlignCenter))
		}

		for steps := 0; steps < rows; steps++ {
			Table.SetCell(steps +1, 0,
				tview.NewTableCell(fmt.Sprintf("%d", steps)).
					SetTextColor(tcell.ColorWhite).
					SetAlign(tview.AlignCenter))
		}

		// page column
		for i := 0; i < prsRange; i++ {
			Table.SetCell(i + 2, 1,
				tview.NewTableCell(fmt.Sprintf("%d", prs[i])).
					SetTextColor(tcell.ColorOrange).
					SetAlign(tview.AlignCenter))
		}

		// TODO: frame column

		Table.SetCell(1, 2,
			tview.NewTableCell(strings.Repeat("[ ]", frames)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter))

		Table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				Table.SetSelectable(true, true)
			}
		}).SetSelectedFunc(func(row int, column int) {
			Table.GetCell(row, column).SetTextColor(tcell.ColorRed)
			Table.SetSelectable(false, false)
		})
	}

	

