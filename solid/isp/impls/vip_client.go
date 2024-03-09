package impls

import "fmt"

type VipClient struct {
	Name string
}

func (vp *VipClient) Print(doc string) {
	fmt.Println("Service: VIP (Print), Client:", vp.Name, "Document:", doc)
}

func (vp *VipClient) Scan() {
	fmt.Println("Service: VIP (Scan), Client:", vp.Name)
}

func (vp *VipClient) Fax(doc string) {
	fmt.Println("Service: VIP (Fax), Client:", vp.Name, "Document:", doc)
}
