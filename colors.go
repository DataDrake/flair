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

var (
	DefaultFG = color.DefaultFG.String()
	DefaultBG = color.DefaultBG.String()

	Black       = escape.Combine(color.Black.FG(), color.DefaultFG.Swap())
	BlackBG     = escape.Combine(color.Black.BG(), color.DefaultBG.Swap())
	DarkGray    = escape.Combine(color.DarkGray.FG(), color.DefaultFG.Swap())
	DarkGrayBG  = escape.Combine(color.DarkGray.BG(), color.DefaultBG.Swap())
	LightGray   = escape.Combine(color.LightGray.FG(), color.DefaultFG.Swap())
	LightGrayBG = escape.Combine(color.LightGray.BG(), color.DefaultBG.Swap())
	White       = escape.Combine(color.White.FG(), color.DefaultFG.Swap())
	WhiteBG     = escape.Combine(color.White.BG(), color.DefaultBG.Swap())

	Red       = escape.Combine(color.Red.FG(), color.DefaultFG.Swap())
	RedBG     = escape.Combine(color.Red.BG(), color.DefaultBG.Swap())
	Green     = escape.Combine(color.Green.FG(), color.DefaultFG.Swap())
	GreenBG   = escape.Combine(color.Green.BG(), color.DefaultBG.Swap())
	Yellow    = escape.Combine(color.Yellow.FG(), color.DefaultFG.Swap())
	YellowBG  = escape.Combine(color.Yellow.BG(), color.DefaultBG.Swap())
	Blue      = escape.Combine(color.Blue.FG(), color.DefaultFG.Swap())
	BlueBG    = escape.Combine(color.Blue.BG(), color.DefaultBG.Swap())
	Magenta   = escape.Combine(color.Magenta.FG(), color.DefaultFG.Swap())
	MagentaBG = escape.Combine(color.Magenta.BG(), color.DefaultBG.Swap())
	Cyan      = escape.Combine(color.Cyan.FG(), color.DefaultFG.Swap())
	CyanBG    = escape.Combine(color.Cyan.BG(), color.DefaultBG.Swap())

	LightRed       = escape.Combine(color.LightRed.FG(), color.DefaultFG.Swap())
	LightRedBG     = escape.Combine(color.LightRed.BG(), color.DefaultBG.Swap())
	LightGreen     = escape.Combine(color.LightGreen.FG(), color.DefaultFG.Swap())
	LightGreenBG   = escape.Combine(color.LightGreen.BG(), color.DefaultBG.Swap())
	LightYellow    = escape.Combine(color.LightYellow.FG(), color.DefaultFG.Swap())
	LightYellowBG  = escape.Combine(color.LightYellow.BG(), color.DefaultBG.Swap())
	LightBlue      = escape.Combine(color.LightBlue.FG(), color.DefaultFG.Swap())
	LightBlueBG    = escape.Combine(color.LightBlue.BG(), color.DefaultBG.Swap())
	LightMagenta   = escape.Combine(color.LightMagenta.FG(), color.DefaultFG.Swap())
	LightMagentaBG = escape.Combine(color.LightMagenta.BG(), color.DefaultBG.Swap())
	LightCyan      = escape.Combine(color.LightCyan.FG(), color.DefaultFG.Swap())
	LightCyanBG    = escape.Combine(color.LightCyan.BG(), color.DefaultBG.Swap())
)
