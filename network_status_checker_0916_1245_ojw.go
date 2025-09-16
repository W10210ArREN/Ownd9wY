// 代码生成时间: 2025-09-16 12:45:25
package main

import (
    "net"
    "os"
    "revel"
)

// NetworkStatusChecker checks the network connection status
type NetworkStatusChecker struct {
    *revel.Controller
}

// Check method is used to check the network connection status of a given host
func (c NetworkStatusChecker) Check(host string) revel.Result {
    // Validate input host
    if host == "" {
        return c.RenderError(revel.NewError("Host is required"))
    }

    // Try to establish a TCP connection to the host on port 80
    conn, err := net.Dial("tcp", host+":80")
    if err != nil {
        // Return an error message if connection fails
        return c.RenderError(revel.NewError("Failed to connect to host: " + err.Error()))
    }
    defer conn.Close()

    // Return a success message if connection is established
    return c.RenderJson(map[string]string{
        "status": "Connected",
        "host": host,
    })
}

func init() {
    // Define the application's routes
    revel.Router(
        &NetworkStatusChecker{},
        ["Check/{host}"],
        func(c *NetworkStatusChecker, host string) revel.Result {
            // Call the Check method with the provided host
            return c.Check(host)
        },
    )
}
