package pocketfft

// #include "pocketfft.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type Plan struct {
	plan   C.cfft_plan
	length int
}

func NewPlan(length int) *Plan {
	return &Plan{plan: C.make_cfft_plan(C.size_t(length)), length: length}
}

func (p *Plan) Destroy() {
	C.destroy_cfft_plan(p.plan)
}

func (p *Plan) Forward(data []complex128, fct float64) {
	p.checkLen(data)
	p.checkReturnValue(
		C.cfft_forward(p.plan, (*C.double)(unsafe.Pointer(&data[0])), C.double(fct)))
}

func (p *Plan) Backward(data []complex128, fct float64) {
	p.checkLen(data)
	p.checkReturnValue(
		C.cfft_backward(p.plan, (*C.double)(unsafe.Pointer(&data[0])), C.double(fct)))
}

func (p *Plan) checkLen(data []complex128) {
	if len(data) != p.length {
		panic(fmt.Sprintf("expected slice of length %d, got slice of length %d",
			p.length, len(data)))
	}
}

func (p *Plan) checkReturnValue(status C.int) {
	if int(status) != 0 {
		panic(fmt.Sprintf("expected return value of 0, got return value of %d",
			int(status)))
	}
}
