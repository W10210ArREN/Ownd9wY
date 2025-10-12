// 代码生成时间: 2025-10-12 20:53:48
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "regexp"
    "time"

    "github.com/hybridgroup/go-ocr"
    "github.com/revel/revel"
    "image"
    "image/jpeg"
)

// OCRService holds the necessary parameters for OCR operations.
type OCRService struct {
    // Empty struct for now, may be extended with configuration options in the future.
}	

// NewOCRService creates a new OCRService instance.
func NewOCRService() *OCRService {
    return &OCRService{}
}

// ProcessImage performs OCR on an uploaded image file and returns the recognized text.
func (service *OCRService) ProcessImage(imagePath string) (string, error) {
    // Check if the file exists
    if _, err := os.Stat(imagePath); os.IsNotExist(err) {
        return "", err
    }

    // Initialize the OCR client
    ocrClient := ocr.NewTesseract()

    // Perform OCR
    text, err := ocrClient.RecognizeFile(imagePath)
    if err != nil {
        return "", err
    }

    // Return the recognized text
    return text, nil
}

// OCRController handles HTTP requests for the OCR service.
type OCRController struct {
    *revel.Controller
    Service *OCRService
}

// NewOCRController initializes a new OCRController with an OCRService instance.
func NewOCRController() *OCRController {
    return &OCRController{
        Service: NewOCRService(),
    }
}

// UploadImage handles the image upload and performs OCR on the uploaded file.
func (c OCRController) UploadImage() revel.Result {
    // Check if a file has been uploaded
    if c.Params.Files == nil {
        return c.RenderError(revel.NewError("No file uploaded."))
    }

    // Get the uploaded image file
    imageFile := c.Params.Files["image"]

    // Validate the file
    if imageFile.Size == 0 || imageFile.ContentType != "image/jpeg" {
        return c.RenderError(revel.NewError("Invalid file. Please upload a JPEG image."))
    }

    // Save the uploaded file to a temporary location
    tmpFile, err := os.CreateTemp(".", "ocr-*.jpg")
    if err != nil {
        return c.RenderError(err)
    }
    defer os.Remove(tmpFile.Name())
    _, err = tmpFile.Write(imageFile.Content)
    if err != nil {
        return c.RenderError(err)
    }
    tmpFile.Close()

    // Perform OCR on the uploaded image
    recognizedText, err := c.Service.ProcessImage(tmpFile.Name())
    if err != nil {
        return c.RenderError(err)
    }

    // Return the recognized text as a JSON response
    return c.RenderJSON(revel.Map{
        "filename": imageFile.Filename,
        "recognizedText": recognizedText,
    })
}

func init() {
    // Register the OCRController
    revel.Router(
        (*OCRController)(nil),
        ["POST", "/ocr/upload-image"],
        "UploadImage",
    )
}

func main() {
    revel.Run()
}