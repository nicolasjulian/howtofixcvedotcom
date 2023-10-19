package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
)


// Logging request
func LogRequest(c echo.Context) {
	req := c.Request()

	// Print the request method and URI
	fmt.Printf("Request %s %s\n", req.Method, req.RequestURI)

	// Print the request body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading request body: %v\n", err)
	} else {
		log.Printf("Request Body: %s\n", string(body))
	}

	// Reset the request body
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
}
