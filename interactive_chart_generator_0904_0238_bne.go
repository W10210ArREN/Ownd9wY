// 代码生成时间: 2025-09-04 02:38:01
 * interactive_chart_generator.go
 * This is a simple application using Revel framework to create an interactive chart generator.
 */

package main

import (
    "encoding/json"
    "github.com/revel/revel"
    "os"
    "os/exec"
)

type Chart struct {
    // Chart structure
    Title string `json:"title"`
    Width  int    `json:"width"`
    Height int    `json:"height"`
    Data   []DataPoint
}

type DataPoint struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

type ChartResponse struct {
    Filename string `json:"filename"`
    Content  string `json:"content"`
}

// ChartController handles the chart generation
type ChartController struct {
    *revel.Controller
}

func (c ChartController) AddPoint(chartId string, point DataPoint) revel.Result {
    // Retrieve chart from storage and add point
    // For simplicity, we'll just simulate this with a file
    filename := "charts/" + chartId + ".json"
    file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        return c.RenderError(err)
    }
    defer file.Close()

    var chart Chart
    if err := json.NewDecoder(file).Decode(&chart); err != nil {
        return c.RenderError(err)
    }

    chart.Data = append(chart.Data, point)
    if _, err := file.Seek(0, 0); err != nil {
        return c.RenderError(err)
    }
    if err := json.NewEncoder(file).Encode(chart); err != nil {
        return c.RenderError(err)
    }

    return c.RenderJSON(ChartResponse{filename, "Point added successfully"})
}

func (c ChartController) GenerateChart(chart Chart) revel.Result {
    // Generate chart image using an external tool like gnuplot
    // For simplicity, we'll just simulate this with a shell command
    filename := "charts/" + chart.Title + ".png"
    cmd := exec.Command("gnuplot", "-p", "-e", "set terminal png; set output '"+filename+