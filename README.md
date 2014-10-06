beego-pongo2.v2
============

A tiny little helper for using Pongo2 (v2) with Beego.

Status: With little tests. IF YOU FIND BUGS, PLS LET ME KNOW.

Documentation: http://godoc.org/github.com/ipfans/beego-pongo2.v2

Based on [https://github.com/oal/beego-pongo2](https://github.com/oal/beego-pongo2)

## Usage

```go
package controllers

import (
    "github.com/astaxie/beego"
    "github.com/ipfans/beego-pongo2.v2"
)

type MainController struct {
    beego.Controller
}

func (this *MainController) Get() {
    pongo2.Render(this.Ctx, "page.html", pongo2.Context{
        "ints": []int{1, 2, 3, 4, 5},
    })
}
```