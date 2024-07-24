package patterns

import "fmt"

func Usecase() {
	h1 := &InfoHandler{}
	h2 := &WarningHandler{}
	h3 := &ErrorHandler{}

	h1.SetNext(h2)
	h2.SetNext(h3)

	for _, request := range []string{"info", "warning", "error", "unknown"} {
		h1.Handle(request)
	}

}

type IHandler interface {
	Handle(request string)
	SetNext(h IHandler)
}

type BaseHandler struct {
	next IHandler
}

func (h *BaseHandler) SetNext(nh IHandler) {
	h.next = nh
}

func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	} else {
		fmt.Printf("No handler could handle the request: %s\n", request)
	}
}

type InfoHandler struct {
	BaseHandler
}

func (h *InfoHandler) Handle(request string) {
	if request == "info" {
		fmt.Printf("Info handler: %s\n", request)
	} else {
		fmt.Println("InfoHandler: Passing request to next handler")
		h.BaseHandler.Handle(request)
	}
}

type WarningHandler struct {
	BaseHandler
}

func (h *WarningHandler) Handle(request string) {
	if request == "warning" {
		fmt.Printf("Warning handler: %s\n", request)
	} else {
		fmt.Println("WarningHandler: Passing request to next handler")
		h.BaseHandler.Handle(request)
	}
}

type ErrorHandler struct {
	BaseHandler
}

func (h *ErrorHandler) Handle(request string) {
	if request == "error" {
		fmt.Printf("Error handler: %s\n", request)
	} else {
		fmt.Println("ErrorHandler: Passing request to next handler")
		h.BaseHandler.Handle(request)
	}
}
