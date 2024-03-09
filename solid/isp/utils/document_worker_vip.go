package utils

import "sample_solid_go/solid/isp/tools"

type DocumentWorkerVip struct {
	Printer    tools.Printer
	Scanner    tools.Scanner
	FaxMachine tools.FaxMachine
}

func (dwv *DocumentWorkerVip) Print(doc string) {
	dwv.Printer.Print(doc)
}

func (dwv *DocumentWorkerVip) Scan() {
	dwv.Scanner.Scan()
}

func (dwv *DocumentWorkerVip) Fax(doc string) {
	dwv.FaxMachine.Fax(doc)
}
