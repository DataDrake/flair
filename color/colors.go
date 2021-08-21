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

var (
	DefaultFG    = escape.Sequence{Pre: "39", Post: ""}
	DefaultBG    = escape.Sequence{Pre: "49", Post: ""}
	Black        = Color{0}
	Red          = Color{1}
	Green        = Color{2}
	Yellow       = Color{3}
	Blue         = Color{4}
	Magenta      = Color{5}
	Cyan         = Color{6}
	LightGray    = Color{7}
	DarkGray     = Color{8}
	LightRed     = Color{9}
	LightGreen   = Color{10}
	LightYellow  = Color{11}
	LightBlue    = Color{12}
	LightMagenta = Color{13}
	LightCyan    = Color{14}
	White        = Color{15}
)

const (
	fg = "38" + escape.Separator + "5" + escape.Separator
	bg = "48" + escape.Separator + "5" + escape.Separator
)

type Color struct {
	Code uint8
}

func (c Color) seq(suffix string, eight, sixteen uint8) escape.Sequence {
	var pre string
	code := c.Code
	switch {
	case code < 8, code == 9:
		pre = strconv.Itoa(int(code + eight))
	case code < 16:
		pre = strconv.Itoa(int(code + sixteen))
	default:
		pre = suffix + strconv.Itoa(int(code))
	}
	return escape.Sequence{Pre: pre, Post: ""}
}

func (c Color) FG() escape.Sequence {
	return c.seq(fg, 30, 92)
}

func (c Color) BG() escape.Sequence {
	return c.seq(bg, 40, 102)
}
