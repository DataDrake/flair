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
	"github.com/DataDrake/flair/escape"
)

// CSI Character Attributes (SGR) - Text Formatting

// Helper functions for changing text styles (Auto-reset at string end)
var (
	Bold      = escape.Combine(escape.Bold, escape.ResetBold.Swap()).Func()
	Dim       = escape.Combine(escape.Dim, escape.ResetDim.Swap()).Func()
	Underline = escape.Combine(escape.Underline, escape.ResetUnderline.Swap()).Func()
	Blink     = escape.Combine(escape.Blink, escape.ResetBlink.Swap()).Func()
	Reverse   = escape.Combine(escape.Reverse, escape.ResetReverse.Swap()).Func()
	Hidden    = escape.Combine(escape.Hidden, escape.ResetHidden.Swap()).Func()
)

// Helper strings for clearing specific formatting
var (
	Reset          = escape.Reset.String() // Reset clears all formatting, including colors
	ResetBold      = escape.ResetBold.String()
	ResetDim       = escape.ResetDim.String()
	ResetUnderline = escape.ResetUnderline.String()
	ResetBlink     = escape.ResetBlink.String()
	ResetReverse   = escape.ResetReverse.String()
	ResetHidden    = escape.ResetHidden.String()
)
