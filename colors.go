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
)

var (
	DefaultFG = color.DefaultFG.Func()
	DefaultBG = color.DefaultBG.Func()

	Black   = color.Black.FG().Func()
	BlackBG = color.Black.BG().Func()

	Red         = color.Red.FG().Func()
	RedBG       = color.Red.BG().Func()
	Green       = color.Green.FG().Func()
	GreenBG     = color.Green.BG().Func()
	Yellow      = color.Yellow.FG().Func()
	YellowBG    = color.Yellow.BG().Func()
	Blue        = color.Blue.FG().Func()
	BlueBG      = color.Blue.BG().Func()
	Magenta     = color.Magenta.FG().Func()
	MagentaBG   = color.Magenta.BG().Func()
	Cyan        = color.Cyan.FG().Func()
	CyanBG      = color.Cyan.BG().Func()
	LightGray   = color.LightGray.FG().Func()
	LightGrayBG = color.LightGray.BG().Func()
	DarkGray    = color.DarkGray.FG().Func()
	DarkGrayBG  = color.DarkGray.BG().Func()

	LightRed       = color.LightRed.FG().Func()
	LightRedBG     = color.LightRed.BG().Func()
	LightGreen     = color.LightGreen.FG().Func()
	LightGreenBG   = color.LightGreen.BG().Func()
	LightYellow    = color.LightYellow.FG().Func()
	LightYellowBG  = color.LightYellow.BG().Func()
	LightBlue      = color.LightBlue.FG().Func()
	LightBlueBG    = color.LightBlue.BG().Func()
	LightMagenta   = color.LightMagenta.FG().Func()
	LightMagentaBG = color.LightMagenta.BG().Func()
	LightCyan      = color.LightCyan.FG().Func()
	LightCyanBG    = color.LightCyan.BG().Func()

	White   = color.White.FG().Func()
	WhiteBG = color.White.BG().Func()
)
