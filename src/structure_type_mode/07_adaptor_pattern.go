package structure_type_mode

import "fmt"

// Adaptee is the interface that needs to be adapted
type Adaptee interface {
	SpecificRequest() string
}

// AdapteeImpl is a concrete implementation of Adaptee
type AdapteeImpl struct{}

func (a *AdapteeImpl) SpecificRequest() string {
	return "specific request"
}

// Target is the target interface that we want to adapt to.
type Target interface {
	Request() string
}

// Adapter is the adapter that adapts Adaptee to Target.
type Adaptor struct {
	Adaptee
}

func (a *Adaptor) Request() string {
	return a.SpecificRequest()
}

// Client is the consumer of Target interface.
func Client(t Target) string {
	return t.Request()
}

func CallAdaptorPattern() {
	adaptee := &AdapteeImpl{}
	target := &Adaptor{Adaptee: adaptee}
	output := Client(target)
	fmt.Println(output)
}
