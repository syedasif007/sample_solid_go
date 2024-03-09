package impls

import "fmt"

type NetworkWriter struct {
	EndPoint string
}

func (nw NetworkWriter) Write(data []byte) error {
	fmt.Printf("%s %s\n", string(data), nw.EndPoint)
	return nil
}
