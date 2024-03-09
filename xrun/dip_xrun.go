package xrun

import (
	"fmt"
	"sample_solid_go/solid/dip/impls"
	"sample_solid_go/solid/dip/utils"
)

func DIP() {

	fmt.Println("Result of DIP:")

	fw := impls.FileWriter{FileName: "data.txt"}
	processor := utils.Processor{Writer: fw}
	processor.ProcessAndWrite([]byte("Data is written to a file."))

	nw := impls.NetworkWriter{EndPoint: "localhost"}
	processor.Writer = nw
	processor.ProcessAndWrite([]byte("Data is written to"))

	fmt.Println()
}
