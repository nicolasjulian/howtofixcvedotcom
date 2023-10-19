package configs

import (
    "encoding/base64"
    "os"
    "fmt"
)

var (
    CVE_API      string
    AUTH_USER    string
    AUTH_PASSWORD string
)

func init() {
    CVE_API = os.Getenv("CVE_API")
    AUTH_USER = os.Getenv("AUTH_USER")
    AUTH_PASSWORD = os.Getenv("AUTH_PASSWORD")
}

func GetBasicAuthHeader() string {
    auth := AUTH_USER + ":" + AUTH_PASSWORD
    authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
    return "Basic " + authEncoded
}

func main() {
    // Print the Basic Authentication header
    fmt.Println(GetBasicAuthHeader())
}
