package xrun

import (
	"fmt"
	"sample_solid_go/solid/isp/impls"
	"sample_solid_go/solid/isp/utils"
)

const (
	PRINT          = 0
	SCAN           = 1
	FAX            = 2
	SCAN_PRINT     = 3
	SCAN_FAX       = 4
	SCAN_PRINT_FAX = 5
)

type Client struct {
	Name string
	Task int
}

func ISPx() {

	fmt.Println("Result of ISPx:")

	var clients []Client

	clients = append(clients, Client{Name: "A", Task: 0})
	clients = append(clients, Client{Name: "B", Task: 1})
	clients = append(clients, Client{Name: "C", Task: 2})
	clients = append(clients, Client{Name: "D", Task: 3})
	clients = append(clients, Client{Name: "E", Task: 4})
	clients = append(clients, Client{Name: "F", Task: 5})

	clientX := &impls.RegularClient{Name: ""}
	docPrinter := &utils.DocumentPrinter{Printer: clientX}

	clientY := &impls.AvgClient{Name: ""}
	docWorker := &utils.DocumentWorker{Printer: clientY, Scanner: clientY}

	clientZ := &impls.VipClient{Name: ""}
	docWorkerVip := &utils.DocumentWorkerVip{Printer: clientZ, Scanner: clientZ, FaxMachine: clientZ}

	for _, v := range clients {

		switch v.Task {

		case 0:
			clientX.Name = v.Name
			docPrinter.Printer = clientX
			docPrinter.Print("Document:" + fmt.Sprintf("%d", v.Task))

		case 1:
			clientY.Name = v.Name
			docWorker.Scanner = clientY
			docWorker.Scan()

		case 2:
			clientZ.Name = v.Name
			docWorkerVip.FaxMachine = clientZ
			docWorkerVip.Fax("Document:" + fmt.Sprintf("%d", v.Task))

		case 3:
			clientZ.Name = v.Name
			docWorkerVip.Scanner = clientZ
			docWorkerVip.Printer = clientZ
			docWorkerVip.Scan()
			docWorkerVip.Print("Document:" + fmt.Sprintf("%d", v.Task))

		case 4:
			clientZ.Name = v.Name
			docWorkerVip.Scanner = clientZ
			docWorkerVip.FaxMachine = clientZ
			docWorkerVip.Scan()
			docWorkerVip.Fax("Document:" + fmt.Sprintf("%d", v.Task))

		case 5:
			clientZ.Name = v.Name
			docWorkerVip.Scanner = clientZ
			docWorkerVip.Printer = clientZ
			docWorkerVip.FaxMachine = clientZ
			docWorkerVip.Scan()
			docWorkerVip.Print("Document:" + fmt.Sprintf("%d", v.Task))
			docWorkerVip.Fax("Document:" + fmt.Sprintf("%d", v.Task))

		default:
			fmt.Println("Unknown Task!")
		}
	}

	fmt.Println()
}

func ISP() {

	fmt.Println("Result of ISP:")

	// only print
	regularClient := &impls.RegularClient{Name: "Client-1"}
	docPrinter := &utils.DocumentPrinter{Printer: regularClient}
	docPrinter.Print("Document-1")

	// only scan or only print or (scan and print)
	avgClient1 := &impls.AvgClient{Name: "Client-2"}
	docWorker := &utils.DocumentWorker{Printer: avgClient1, Scanner: avgClient1}
	docWorker.Scan()
	docWorker.Print("Document-2")

	// only scan or only print or (scan and print)
	avgClient2 := &impls.AvgClient{Name: "Client-3"}
	docWorker.Scanner = avgClient2
	docWorker.Scan()

	// scan | print | fax
	vipClient := &impls.VipClient{Name: "Client-4"}
	docWorkerVip := &utils.DocumentWorkerVip{Printer: vipClient, Scanner: vipClient, FaxMachine: vipClient}
	docWorkerVip.Scan()
	docWorkerVip.Print("Document-3")
	docWorkerVip.Fax("Document-3")

	fmt.Println()
}
