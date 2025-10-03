// 代码生成时间: 2025-10-03 17:39:46
package ocr

import (
    "errors"
    "fmt"
    "os"
    "path/filepath"
    "revel"
    "strings"
    "time"

    // Importing Tesseract OCR
    "gocv.io/x/gocv"
)

// OCRService is the struct that represents the OCR service.
type OCRService struct {
    *revel.Controller
}

// NewOCRService creates a new instance of OCRService.
func NewOCRService() *OCRService {
    return &OCRService{
        Controller: &revel.Controller{},
    }
}

// RecognizeText performs OCR on an image file and returns the recognized text.
//
// @param filename is the path to the image file to perform OCR on.
// @return The recognized text and an error if any.
func (o *OCRService) RecognizeText(filename string) (string, error) {
    // Check if the file exists
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        return "", errors.New("file does not exist")
    }

    // Open the image file
    img, err := gocv.IMRead(filename, gocv.IMReadColor)
    if err != nil {
        return "", fmt.Errorf("failed to read image: %v", err)
    }
    defer img.Close()

    // Create a new OCR engine
    recognizer := gocv.NewTesseract(gocv.DefaultTesseractPath)
    if err := recognizer.SetImage(img); err != nil {
        return "", fmt.Errorf("failed to set image to OCR engine: %v", err)
    }
    defer recognizer.Clear()

    // Perform OCR and obtain text
    text, err := recognizer.Text()
    if err != nil {
        return "", fmt.Errorf("failed to recognize text: %v", err)
    }
    return text, nil
}

// RecognizeTextAction is the Revel action that handles the OCR request.
//
// @param filename is the path to the image file to perform OCR on.
// @Render text The recognized text or an error message.
func (o *OCRService) RecognizeTextAction(filename string) revel.Result {
    text, err := o.RecognizeText(filename)
    if err != nil {
        return o.RenderError(err)
    }
    return o.RenderText(text)
}
