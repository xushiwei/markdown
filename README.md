Go+ domain-specific text: markdown
=====

This is a demo of Go+ domain-specific text customization. 

## example

```go
import (
    "os"

    "github.com/xushiwei/markdown"
)

md := markdown`
# Title

Hello world
`

md.convert os.Stdout
```
