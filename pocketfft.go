// Package `pocketfft` provides efficient algorithms for FFT and IFFT.

package pocketfft

import (
	"math"

	"github.com/manuelmenzella/go-pocketfft/internal/pocketfft"
)

// Plan represents a strategy for computing FFTs/IFFTs for a given input size.
type Plan struct {
	plan *pocketfft.Plan
}

// NewPlan creates a new FFT/IFFT plan for a given input size.
// Created plans may be reused across multiple FFTs/IFFTs.
// Users must call `Destroy` to release resources.
func NewPlan(size int) *Plan {
	return &Plan{plan: pocketfft.NewPlan(size)}
}

// Destroy must be called to release resources associated with the `Plan`.
// No further calls to any bound functions are allowed after calling `Destroy`.
func (p *Plan) Destroy() {
	p.plan.Destroy()
}

// FFT performs a single FFT calculation on the input data,
// which remains unchanged.
func (p *Plan) FFT(data []complex128) []complex128 {
	out := make([]complex128, len(data))
	copy(out, data)
	p.plan.Forward(out, math.Sqrt(1/float64(len(data))))
	return out
}

// DestructiveFFT performs a single FFT calculation on the input data,
// which may be modified for efficiency.
func (p *Plan) DestructiveFFT(data []complex128) []complex128 {
	p.plan.Forward(data, math.Sqrt(1/float64(len(data))))
	return data
}

// IFFT performs a single IFFT calculation on the input data,
// which remains unchanged.
func (p *Plan) IFFT(data []complex128) []complex128 {
	out := make([]complex128, len(data))
	copy(out, data)
	p.plan.Backward(out, math.Sqrt(1/float64(len(data))))
	return out
}

// DestructiveIFFT performs a single IFFT calculation on the input data,
// which may be modified for efficiency.
func (p *Plan) DestructiveIFFT(data []complex128) []complex128 {
	p.plan.Backward(data, math.Sqrt(1/float64(len(data))))
	return data
}
