package spinner

import (
    "fmt"
    "io"
)

type Spinner int

var spinBytes = []byte{'-', '/', '|', '\\'}

func (s *Spinner) Tick() {
    *s = (*s + 1) % 4
    fmt.Printf("\b%s", spinBytes[*s:*s+1])
}

func (s *Spinner) Done() {
    fmt.Print("\b ")
}

type SpinReadCloser struct {
    io.ReadCloser
    Spinner
}

func (r *SpinReadCloser) Read(b []byte) (int, error) {
    r.Spinner.Tick()
    return r.ReadCloser.Read(b)
}

func (r *SpinReadCloser) Close() error {
    r.Spinner.Done()
    return r.ReadCloser.Close()
}
