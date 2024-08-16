# gFly View - Template engine

    Copyright © 2023, gFly
    https://www.gfly.dev
    All rights reserved.

# Use Pongo template engine

### Install
```bash
go get -u github.com/gflydev/view/pongo@latest
```

### Register
```go
import (
    "github.com/gflydev/core"
    "github.com/gflydev/view/pongo"	
)

// Register view
core.RegisterView(pongo.New())
```
