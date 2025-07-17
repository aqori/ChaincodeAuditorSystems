// cmd/chaincodeauditorsystems/main.go
package main

import (
"flag"
"log"
"os"

"chaincodeauditorsystems/internal/chaincodeauditorsystems"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := chaincodeauditorsystems.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
