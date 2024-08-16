# gFly View - Template engine

    Copyright © 2023, gFly
    https://www.gfly.dev
    All rights reserved.

# gFly View - Pongo engine

### Usage

Install
```bash
go get -u github.com/gflydev/view/pongo@v1.0.2
```


Quick usage `main.go`
```go
import (
    "github.com/gflydev/core"
    "github.com/gflydev/view/pongo"	
)

// Register view
core.RegisterView(pongo.New())
```

### Page & View
```go
// =========================================================================================
//                                     About page
// =========================================================================================

// NewAboutPage As a constructor to create a About Page.
func NewHomePage() *AboutPage {
return &AboutPage{}
}

type AboutPage struct {
core.Page
}

func (m *AboutPage) Handle(c *core.Ctx) error {
return c.View("about", core.Data{
"title": "About Us",
})
}
```

Template file `resources/views/about.tpl`
```html
<center><h2>{{ title }}</h2></center>
```