package impls

import "fmt"

type AvgClient struct {
	Name string
}

func (ac *AvgClient) Print(doc string) {
	fmt.Println("Service: Average (Print), Client:", ac.Name, "Document:", doc)
}

func (ac *AvgClient) Scan() {
	fmt.Println("Service: Average (Scan), Client:", ac.Name)
}
