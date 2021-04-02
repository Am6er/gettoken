package main

import (
        "crypto/hmac"
        "crypto/sha1"
        "encoding/base64"
        "fmt"
        "log"
        "os"
        "strings"
)

func computeHmac256(message string, secret string) string {
        key := []byte(secret)
        h := hmac.New(sha1.New, key)
        _, err := h.Write([]byte(strings.ToLower(message)))
        if err != nil {
                log.Fatal(err)
        }
        return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {
        if len(os.Args) != 3 {
                fmt.Println("Usage:\n",
                        "gettoken <param_concat> <secretkey>\n\n",
                        "Where:\n",
                        "param_concat: Concatenation of payload parameter values\n",
                        "secretkey: secret key string\n\n",
                        "Example:\n",
                        "gettoken UserPassword 606136046f08c75c6a017f12dae767e2\n")
                os.Exit(1)
        } else {
                fmt.Println(computeHmac256(os.Args[1], os.Args[2]))
        }
}
