package spinner

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "time"
)

// Spin for 5 seconds at 20 frames per second.
func ExampleSpinner() {
    fmt.Print("Spinning... ")
    defer fmt.Println()

    var s Spinner
    defer s.Done()

    ticker := time.NewTicker(50 * time.Millisecond)
    defer ticker.Stop()
    done := time.After(5 * time.Second)
    for {
        select {
        case <-ticker.C:
            s.Tick()
        case <-done:
            return
        }
    }
}

// Read 100 megabytes through the SpinReadCloser.
func ExampleSpinReadCloser() {
    f, err := os.Open("/dev/random")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print("Reading... ")
    defer fmt.Println()

    rc := &SpinReadCloser{ReadCloser: f}
    defer rc.Close()

    _, err = io.CopyN(ioutil.Discard, rc, 100<<20)
    if err != nil {
        log.Fatal(err)
    }
}
