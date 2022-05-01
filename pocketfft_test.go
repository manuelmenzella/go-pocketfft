package pocketfft

import (
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func randomData(length int) []complex128 {
	data := make([]complex128, length)
	for i := 0; i < length; i++ {
		data[i] = complex(rand.Float64(), rand.Float64())
	}
	return data
}

func TestFFTAndIFFT(t *testing.T) {
	const length int = 7
	original := randomData(length)

	plan := NewPlan(length)
	defer plan.Destroy()

	fft := plan.FFT(original)
	ifft := plan.IFFT(fft)

	complexToFloats :=
		cmp.Transformer("ComplexToFloats", func(in complex128) (out struct{ Real, Imag float64 }) {
			out.Real, out.Imag = real(in), imag(in)
			return out
		})
	if !cmp.Equal(original, ifft, complexToFloats, cmpopts.EquateApprox(1e-6, 1e-6)) {
		t.Errorf("expected %v, got %v", original, ifft)
	}
}

func BenchmarkFFTAndIFFT(b *testing.B) {
	const length int = 100_000
	data := randomData(length)

	plan := NewPlan(length)
	defer plan.Destroy()

	for i := 0; i < b.N; i++ {
		data = plan.FFT(data)
		data = plan.IFFT(data)
	}
}

func BenchmarkDestructiveFFTAndDestructiveIFFT(b *testing.B) {
	const length int = 100_000
	data := randomData(length)

	plan := NewPlan(length)
	defer plan.Destroy()

	for i := 0; i < b.N; i++ {
		data = plan.DestructiveFFT(data)
		data = plan.DestructiveIFFT(data)
	}
}
