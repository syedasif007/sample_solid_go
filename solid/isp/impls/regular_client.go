package impls

import "fmt"

type RegularClient struct {
	Name string
}

func (rc *RegularClient) Print(doc string) {
	fmt.Println("Service: Regular (Print), Client:", rc.Name, "Document:", doc)
}
