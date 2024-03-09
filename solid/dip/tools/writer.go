package tools

type Writer interface {
	Write(data []byte) error
}
