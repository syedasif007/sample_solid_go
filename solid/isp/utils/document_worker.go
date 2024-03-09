package utils

import "sample_solid_go/solid/isp/tools"

type DocumentWorker struct {
	Printer tools.Printer
	Scanner tools.Scanner
}

func (dw *DocumentWorker) Print(doc string) {
	dw.Printer.Print(doc)
}

func (dw *DocumentWorker) Scan() {
	dw.Scanner.Scan()
}
