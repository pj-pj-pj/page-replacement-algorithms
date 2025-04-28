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


var AlgoFaults = tview.NewTable()


var FifoFaultsText = tview.NewTextView().SetText("")
var LruFaultsText = tview.NewTextView().SetText("")
var OptFaultsText = tview.NewTextView().SetText("")



// this function puts the contents inside the table
func PopulateTable(prs []int, frames int) { 
	fifoResult, FifoFaults := algorithms.Fifo(prs, frames)
	lruResult, LruFaults := algorithms.Lru(prs, frames)
	optResult, OptFaults := algorithms.Opt(prs, frames)


	FifoFaultsText.SetText(fmt.Sprintf("「 First-In, First-Out Page Faults: %d 」", FifoFaults))
	LruFaultsText.SetText(fmt.Sprintf("「 Least Recently Used Page Faults: %d 」", LruFaults))
	OptFaultsText.SetText(fmt.Sprintf("「 Optimal Page Faults: %d 」", OptFaults))

	AlgoFaults.SetCell(0, 0,
		tview.NewTableCell("  [limegreen]「 Fault Statistics:  | ").
				SetAlign(tview.AlignCenter))

	AlgoFaults.SetCell(0, 1,
			tview.NewTableCell(fmt.Sprintf("[darkorange]FIFO: %d  [limegreen]| ", FifoFaults)).
					SetAlign(tview.AlignCenter))

	AlgoFaults.SetCell(0, 2,
		tview.NewTableCell(fmt.Sprintf(" [greenyellow]LRU: %d  [limegreen]| ", LruFaults)).
				SetAlign(tview.AlignCenter))

	AlgoFaults.SetCell(0, 3,
		tview.NewTableCell(fmt.Sprintf(" [dodgerblue]OPT: %d  [limegreen]」", OptFaults)).
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



	// lru table
	for c, val := range prs {
    TableStringsLru.SetCell(0, c,
        tview.NewTableCell(fmt.Sprintf("  %d", val)).
            SetAlign(tview.AlignCenter))}

	for i, val := range lruResult {
		prevFaultsCount := 0
		if i > 0 {
				prevFaultsCount = lruResult[i-1].FaultsCount
		}

		for j, framesVal := range val.Frames {
			if val.FaultsCount == prevFaultsCount {
				TableFramesLru.SetCell(j, i,
					tview.NewTableCell(fmt.Sprintf(" [darkslategray]%d ", framesVal)).
							SetAlign(tview.AlignCenter))
			} else if val.Page == framesVal && val.PageFault {
				TableFramesLru.SetCell(j, i,
						tview.NewTableCell(fmt.Sprintf(" [greenyellow]%d ", framesVal)).
								SetAlign(tview.AlignCenter))
			} else {
				TableFramesLru.SetCell(j, i,
						tview.NewTableCell(fmt.Sprintf(" %d ", framesVal)).
								SetAlign(tview.AlignCenter))
			}
		}
	}


	

	// opt table
	for c, val := range prs {
    TableStringsOpt.SetCell(0, c,
        tview.NewTableCell(fmt.Sprintf("  %d", val)).
            SetAlign(tview.AlignCenter))}

	for i, val := range optResult {
		prevFaultsCount := 0
		if i > 0 {
				prevFaultsCount = optResult[i-1].FaultsCount
		}

		for j, framesVal := range val.Frames {
			if val.FaultsCount == prevFaultsCount {
				TableFramesOpt.SetCell(j, i,
					tview.NewTableCell(fmt.Sprintf(" [darkslategray]%d ", framesVal)).
							SetAlign(tview.AlignCenter))
			} else if val.Page == framesVal && val.PageFault {
				TableFramesOpt.SetCell(j, i,
						tview.NewTableCell(fmt.Sprintf(" [dodgerblue]%d ", framesVal)).
								SetAlign(tview.AlignCenter))
			} else {
				TableFramesOpt.SetCell(j, i,
						tview.NewTableCell(fmt.Sprintf(" %d ", framesVal)).
								SetAlign(tview.AlignCenter))
			}
		}
	}
		
	
	// cool navigation stuff i found on the internet
	TableFramesFifo.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			TableFramesFifo.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		TableFramesFifo.GetCell(row, column).SetTextColor(tcell.ColorRed)
		TableFramesFifo.SetSelectable(false, false)
	})

	TableFramesLru.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			TableFramesLru.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		TableFramesLru.GetCell(row, column).SetTextColor(tcell.ColorRed)
		TableFramesLru.SetSelectable(false, false)
	})

	TableFramesOpt.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			TableFramesOpt.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		TableFramesOpt.GetCell(row, column).SetTextColor(tcell.ColorRed)
		TableFramesOpt.SetSelectable(false, false)
	})
}
