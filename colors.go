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

package flair

import (
	"github.com/DataDrake/flair/color"
	"github.com/DataDrake/flair/escape"
)

// CSI Character Attributes (SGR)
// Helper functions for changing colors
var (
	// Defaults
	DefaultFG = color.DefaultFG.String()
	DefaultBG = color.DefaultBG.String()

	// Normal Colors (8 color mode)
	Black       = escape.Combine(color.Black.FG(), color.DefaultFG.Swap()).Func()
	BlackBG     = escape.Combine(color.Black.BG(), color.DefaultBG.Swap()).Func()
	Red         = escape.Combine(color.Red.FG(), color.DefaultFG.Swap()).Func()
	RedBG       = escape.Combine(color.Red.BG(), color.DefaultBG.Swap()).Func()
	Green       = escape.Combine(color.Green.FG(), color.DefaultFG.Swap()).Func()
	GreenBG     = escape.Combine(color.Green.BG(), color.DefaultBG.Swap()).Func()
	Yellow      = escape.Combine(color.Yellow.FG(), color.DefaultFG.Swap()).Func()
	YellowBG    = escape.Combine(color.Yellow.BG(), color.DefaultBG.Swap()).Func()
	Blue        = escape.Combine(color.Blue.FG(), color.DefaultFG.Swap()).Func()
	BlueBG      = escape.Combine(color.Blue.BG(), color.DefaultBG.Swap()).Func()
	Magenta     = escape.Combine(color.Magenta.FG(), color.DefaultFG.Swap()).Func()
	MagentaBG   = escape.Combine(color.Magenta.BG(), color.DefaultBG.Swap()).Func()
	Cyan        = escape.Combine(color.Cyan.FG(), color.DefaultFG.Swap()).Func()
	CyanBG      = escape.Combine(color.Cyan.BG(), color.DefaultBG.Swap()).Func()
	LightGray   = escape.Combine(color.LightGray.FG(), color.DefaultFG.Swap()).Func()
	LightGrayBG = escape.Combine(color.LightGray.BG(), color.DefaultBG.Swap()).Func()

	// Extended Colors (16 color mode)
	DarkGray       = escape.Combine(color.DarkGray.FG(), color.DefaultFG.Swap()).Func()
	DarkGrayBG     = escape.Combine(color.DarkGray.BG(), color.DefaultBG.Swap()).Func()
	LightRed       = escape.Combine(color.LightRed.FG(), color.DefaultFG.Swap()).Func()
	LightRedBG     = escape.Combine(color.LightRed.BG(), color.DefaultBG.Swap()).Func()
	LightGreen     = escape.Combine(color.LightGreen.FG(), color.DefaultFG.Swap()).Func()
	LightGreenBG   = escape.Combine(color.LightGreen.BG(), color.DefaultBG.Swap()).Func()
	LightYellow    = escape.Combine(color.LightYellow.FG(), color.DefaultFG.Swap()).Func()
	LightYellowBG  = escape.Combine(color.LightYellow.BG(), color.DefaultBG.Swap()).Func()
	LightBlue      = escape.Combine(color.LightBlue.FG(), color.DefaultFG.Swap()).Func()
	LightBlueBG    = escape.Combine(color.LightBlue.BG(), color.DefaultBG.Swap()).Func()
	LightMagenta   = escape.Combine(color.LightMagenta.FG(), color.DefaultFG.Swap()).Func()
	LightMagentaBG = escape.Combine(color.LightMagenta.BG(), color.DefaultBG.Swap()).Func()
	LightCyan      = escape.Combine(color.LightCyan.FG(), color.DefaultFG.Swap()).Func()
	LightCyanBG    = escape.Combine(color.LightCyan.BG(), color.DefaultBG.Swap()).Func()
	White          = escape.Combine(color.White.FG(), color.DefaultFG.Swap()).Func()
	WhiteBG        = escape.Combine(color.White.BG(), color.DefaultBG.Swap()).Func()
)
