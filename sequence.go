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

var (
	Reset          = escape.Reset.Func()
	Bold           = escape.Bold.Func()
	Dim            = escape.Dim.Func()
	Underline      = escape.Underline.Func()
	Blink          = escape.Blink.Func()
	Reverse        = escape.Reverse.Func()
	Hidden         = escape.Hidden.Func()
	ResetBold      = escape.ResetBold.Func()
	ResetDim       = escape.ResetDim.Func()
	ResetUnderline = escape.ResetUnderline.Func()
	ResetBlink     = escape.ResetBlink.Func()
	ResetReverse   = escape.ResetReverse.Func()
	ResetHidden    = escape.ResetHidden.Func()
)
