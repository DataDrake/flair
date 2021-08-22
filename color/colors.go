//
// Copyright 2021 Bryan T. Meyers <root@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package color

import (
	"github.com/DataDrake/flair/escape"
	"strconv"
)

// CSI Character Attributes (SGR) - Colors
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html

// Default Colors
var (
	DefaultFG = escape.Sequence{Pre: "39", Post: ""} // DefaultFG resets the foreground color
	DefaultBG = escape.Sequence{Pre: "49", Post: ""} // DefaultBG resets the background color
)

// 16 colors
const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGray
	DarkGray
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White
)

const (
	fg = "38" + escape.Separator + "5" + escape.Separator
	bg = "48" + escape.Separator + "5" + escape.Separator
)

// Color represents a 256 color mode sequence
type Color uint8

// seq generates the necessary escape Sequence for this color, using the smallest encoding
func (c Color) seq(suffix string, eight, sixteen uint8) escape.Sequence {
	var pre string
	switch {
	case c < 8, c == 9:
		pre = strconv.Itoa(int(uint8(c) + eight))
	case c < 16:
		pre = strconv.Itoa(int(uint8(c) + sixteen - 8))
	default:
		pre = suffix + strconv.Itoa(int(c))
	}
	return escape.Sequence{Pre: pre, Post: ""}
}

// FG generates an escape Sequence which will set this color as the Foreground
func (c Color) FG() escape.Sequence {
	return c.seq(fg, 30, 90)
}

// BG generates an escape Sequence which will set this color as the Background
func (c Color) BG() escape.Sequence {
	return c.seq(bg, 40, 100)
}
