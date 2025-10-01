// 代码生成时间: 2025-10-02 03:30:24
// face_recognition_system.go
// This file contains the main application structure for a face recognition system.

package main

import (
    "encoding/json"
    "fmt"
    "github.com/revel/revel"
    "gocv.io/x/gocv"
)

// FaceRecognitionApp is the main application struct.
type FaceRecognitionApp struct {
    revel.Controller
}

// Index handles the GET request to the root path and serves the index page.
func (c FaceRecognitionApp) Index() revel.Result {
    return c.Render()
}

// DetectFace handles the POST request to /detect-face and performs face detection.
func (c FaceRecognitionApp) DetectFace() revel.Result {
    var req FaceRecognitionRequest

    // Decode the JSON request body into the FaceRecognitionRequest struct.
    if err := json.Unmarshal(c.Params.Form["json"], &req); err != nil {
        return c.RenderError(err)
    }

    // Initialize the face detector.
    net := gocv.NewDNNObjectDetectionModel("deploy.prototxt", "res10_300x300_ssd_iter_140000_fp16.caffemodel")
    defer net.Close()

    // Load the image from the request.
    img, err := gocv.IMRead(req.ImageFile, gocv.IMReadUnchanged)
    if err != nil {
        return c.RenderError(err)
    }
    defer img.Close()

    // Perform face detection.
    classes, confidences, rects := net.Detect(img, nil, gocv.NewMat(), 10, 0.6)

    // If no faces are detected, return an error.
    if len(classes) == 0 {
        err := fmt.Errorf("No faces detected in the image")
        return c.RenderError(err)
    }

    // Marshal the detection results into JSON.
    results := FaceDetectionResults{
        Classes:    classes,
        Confidences: confidences,
        Rectangles: rects,
    }
    return c.RenderJSON(results)
}

// FaceRecognitionRequest represents the request data for face recognition.
type FaceRecognitionRequest struct {
    ImageFile string `json:"image_file"`
}

// FaceDetectionResults represents the results of the face detection.
type FaceDetectionResults struct {
    Classes    []string  `json:"classes"`
    Confidences []float32 `json:"confidences"`
    Rectangles []gocv.Rect `json:"rectangles"`
}

func init() {
    revel.Filters = []revel.Filter{
        // Add your global filters or middleware here, e.g.,
gocv.
        revel.recoverPanic,
        revel.Router,
        revel.Params,
        revel.Session,
        revel.Flash,
        revel.I18n,
        revel.Auth,
        revel.Interceptor,
        revel.ActionInvoker,
        revel.TemplateRenderer,
        revel.Profiler,
    }
}
