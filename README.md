# flair
Simple console formatting library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/DataDrake/flair.svg)](https://pkg.go.dev/github.com/DataDrake/flair) [![Go Report Card](https://goreportcard.com/badge/github.com/DataDrake/flair)](https://goreportcard.com/report/github.com/DataDrake/flair) [![license](https://img.shields.io/github/license/DataDrake/flair.svg)]() 

## Motivation

There are definitely a bunch of other libraries like this, but I wanted one I knew would be maintained and properlyversioned. I *might* also have a strong preference about how clean the code that I'm importing is.

## Goals

 * KISS
 * Extensibility
 * A+ Rating on [Report Card](https://goreportcard.com/report/github.com/DataDrake/flair)
 
## Usage

The main package exposes helper functions for formatting text:

```
package main

import (
    "github.com/DataDrake/flair"
)

func main() {
    println(flair.Bold("This is Bolded Text"))
    println(flair.Green("This is Green Text"))
    println(flair.RedBG("This is Text with a Red background"))
}
```

You can pretty easily create custom [colors](https://misc.flogisoft.com/bash/tip_colors_and_formatting) and functions:

```
package main

import (
    "github.com/DataDrake/flair/color"
    "github.com/DataDrake/flair/escape"
)

func main() {
    Purple   := color.Color{56}
    purple   := escape.Combine(Purple.FG(), color.DefaultFG.Swap()).Func()
    purpleBG := escape.Combine(Purple.BG(), color.DefaultBG.Swap()).Func()

    println(purple("This is Purple Text"))
    println(purpleBG("This is Text with a Purple background"))
}
```

And you can easily combine multiple formatting directives into one:

```
package main

import (
    "fmt"
    "github.com/DataDrake/flair/color"
    "github.com/DataDrake/flair/escape"
)

func main() {
    label := escape.Combine(escape.Bold, color.White.FG(), color.Red.FG(), escape.Reset)
    println(label(" Bold Red Label ") + " Default Text")

    // Or with a formatting string
    msgFmt := label(" %s ") + " %s"
    fmt.Printf(msgFmt, "Like", "So")
}
```


## License
 
Copyright 2021 Bryan T. Meyers <root@datadrake.com>
 
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 
http://www.apache.org/licenses/LICENSE-2.0
 
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
