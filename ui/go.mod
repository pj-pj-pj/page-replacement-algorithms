module github.com/pj-pj-pj/page-replacement-algorithms/ui

go 1.24.2

require (
	github.com/gdamore/tcell/v2 v2.8.1
	github.com/pj-pj-pj/page-replacement-algorithms/algorithms v0.0.0-00010101000000-000000000000
	github.com/pj-pj-pj/page-replacement-algorithms/assets v0.0.0-00010101000000-000000000000
	github.com/pj-pj-pj/page-replacement-algorithms/ui/algotable v0.0.0-00010101000000-000000000000
	github.com/rivo/tview v0.0.0-20250330220935-949945f8d922
)

require (
	github.com/gdamore/encoding v1.0.1 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/term v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)

replace github.com/pj-pj-pj/page-replacement-algorithms/assets => ../assets

replace github.com/pj-pj-pj/page-replacement-algorithms/ui/algotable => ./algotable

replace github.com/pj-pj-pj/page-replacement-algorithms/algorithms => ../algorithms
