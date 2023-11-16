# text-captcha-breaker
Text Captcha Breaker client in golang. see https://huggingface.co/spaces/docparser/Text_Captcha_breaker

# Install
```
go get github.com/dvwzj/text-captcha-breaker
```

# Usage

```go
package main

import (
    tcb "github.com/dvwzj/text-captcha-breaker"
)

func main() {
    text, err := tcb.Solve("data:image/png;base64,R0lGODlhZAAeAPUvAGZmZ...")
    // tcb.Solve("path/to/file") is possible too
    if err != nil {
        panic(err)
    }
    println(text)
}
```