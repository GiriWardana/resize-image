package main

import (
    "image"
    "image/jpeg"
    "image/png"
    "os"
    "log"

    "github.com/nfnt/resize"
)

func main() {
    // Open the file
    file, err := os.Open("input.jpg")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Decode the image
    img, _, err := image.Decode(file)
    if err != nil {
        log.Fatal(err)
    }

    // Resize the image to width 800 using Lanczos resampling
    // and preserve the aspect ratio
    m := resize.Resize(800, 0, img, resize.Lanczos3)

    // Create a new file
    out, err := os.Create("output.jpg")
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()

    // Write the resized image to the new file
    jpeg.Encode(out, m, nil)
}
