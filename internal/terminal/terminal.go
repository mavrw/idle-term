package terminal

import "fmt"

type Terminal interface {
	ReadInput() string
	WriteOutput(output string)
}

type DefaultTerminal struct{} // Add width and height?

func NewDefaultTerminal() *DefaultTerminal {
	return &DefaultTerminal{}
}

func (dt *DefaultTerminal) ReadInput() string {
	return ""
}

func (dt *DefaultTerminal) WriteOutput(output string) {
	fmt.Println(output)
}
