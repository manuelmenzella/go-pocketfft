# PocketFFT for Go

`go-pocketfft` provides a thin wrapper over the popular
[`pocketfft`](https://gitlab.mpcdf.mpg.de/mtr/pocketfft) library.
Performance is roughly comparable with that of [FFTW](http://www.fftw.org),
but the `pocketfft` library is published with a friendlier license.

The underlying `pocketfft` is written in C and included in `go-pocketfft`, so
there is no need to install it for neither compilation or execution.

## Installation

```sh
go get -u github.com/manuelmenzella/go-pocketfft
```

## Example

```go
package main

import (
	"math/rand"

    "github.com/manuelmenzella/go-pocketfft/internal/pocketfft"
)

func main() {
	const length int = 1023

	data := make([]complex128, length)
	for i := 0; i < length; i++ {
		data[i] = complex(rand.Float64(), rand.Float64())
	}

    // Create a plan for appropriate length
	plan := pocketfft.NewPlan(length)
	defer plan.Destroy()

	// FFT & IFFT
	fft := plan.FFT(data)
	_ = plan.IFFT(fft)

	// Or, more efficiently, if the input data can be modified...
	fft = plan.DestructiveFFT(data)
	_ = plan.DestructiveIFFT(fft)
}
```
