package utils

import "sample_solid_go/solid/isp/tools"

type DocumentPrinter struct {
	Printer tools.Printer
}

func (dp *DocumentPrinter) Print(doc string) {
	dp.Printer.Print(doc)
}
