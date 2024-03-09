package utils

import "sample_solid_go/solid/dip/tools"

type Processor struct {
	Writer tools.Writer
}

func (p Processor) ProcessAndWrite(data []byte) error {
	return p.Writer.Write(data)
}
