package controllers

import (
    "github.com/nicolasjulian/howtofixcvedotcom/backend/backend/utils"
    "github.com/nicolasjulian/howtofixcvedotcom/backend/configs"

    "encoding/json"
    "log"
    "io/ioutil"
    "fmt"
    "net/http"

    "github.com/labstack/echo/v4"
)


type CVE struct {
    ID         string `json:"id"`
    Summary    string `json:"summary"`
    CreatedAt  string `json:"created_at"`
    UpdatedAt  string `json:"updated_at"`
}

func Latest(e echo.Context) error {
    utils.LogRequest(e)

    req, err := http.NewRequest("GET", configs.CVE_API  + "/api/cve", nil)
    if err != nil {
        fmt.Println(err)
        return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create request"})
    }

    // Add Basic Authentication header to the request
    req.SetBasicAuth(configs.AUTH_USER, configs.AUTH_PASSWORD)

    // Send the request
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send request"})
    }
    defer res.Body.Close()

    // Check the response status code
    if res.StatusCode != http.StatusOK {
        log.Printf("Unexpected status code: %d\n", res.StatusCode)
        return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data"})
    }

    // Read the response body
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read response"})
    }

    // Parse the response JSON 
    var cves []CVE
    if err := json.Unmarshal(body, &cves); err != nil {
        fmt.Println(err)
        return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse JSON"})
    }

    return e.JSON(http.StatusOK, cves)
}


func Info(e echo.Context) error {
    utils.LogRequest(e)

    cveid := e.Param("*")

    req, err := http.NewRequest("GET", configs.CVE_API + "/api/cve/" + cveid, nil)
    if err != nil {
        fmt.Println(err)
        return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create request"})
    }

    // Add Basic Authentication header to the request
    req.SetBasicAuth(configs.AUTH_USER, configs.AUTH_PASSWORD)

    // Send the request
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send request"})
    }
    defer res.Body.Close()

    // Check the response status code
    if res.StatusCode != http.StatusOK {
        log.Printf("Unexpected status code: %d\n", res.StatusCode)
        return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data"})
    }

    // Read the response body
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read response"})
    }

    // Parse the response JSON 
    var cves CVE
    if err := json.Unmarshal(body, &cves); err != nil {
        fmt.Println(err)
        return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse JSON"})
    }

    return e.JSON(http.StatusOK, cves)
}


