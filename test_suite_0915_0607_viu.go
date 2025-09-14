// 代码生成时间: 2025-09-15 06:07:06
@author: Your Name
@date: YYYY-MM-DD
*/

package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "revel"
    "testing"
)

func init() {
    revel.Filters = []revel.Filter{
        revel.PanicFilter,             // Recover from panics and display an error page instead.
        revel.ActionInvoker,           // Invoke action.
    }
}

// TestSuite is a test suite application.
type TestSuite struct {
    *revel.Controller
}

// Index returns a simple greeting.
func (t *TestSuite) Index() revel.Result {
    return t.RenderText("Hello, Revel Test Suite!")
}

// TestExample is a test function that tests the Index method.
func TestExample(t *testing.T) {
    c := &TestSuite{}
    c.StartTime = revel.Now
    
    result, err := c.Index()
    if err != nil {
        t.Fatal(err)
    } else {
        resp := result.(*revel.ResultText)
        if resp.ContentType != "text/plain" || resp.Content != "Hello, Revel Test Suite!" {
            t.Errorf("Expected 'Hello, Revel Test Suite!', got '%s'", resp.Content)
        }
    }
}

// TestMain is the entry point for the test suite.
func TestMain(m *testing.M) {
    revel.Init(
        os.Args[1:],
        revel.WithRandomAppSecret(),
        revel.WithLoggingEnabled(false), // Disable logging for tests.
    )
    os.Exit(m.Run())
}
