# Niceware

This is a port of [niceware](https://github.com/diracdeltas/niceware) from Node.js to Go.

## Example Usage

```
package main

import (
        "fmt"
        "github.com/superp00t/niceware"
        "log"
)

func main() {
        input := []byte{0xDE, 0xAD, 0xBE, 0xEF}

        output, err := niceware.BytesToString(input)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Encoded:", output)

        decoded, err := niceware.StringToBytes(output)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Printf("Decoded: 0x%X\n", decoded)
}
```
