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

package escape

import (
	"github.com/DataDrake/flair/style"
)

// CSI Character Attributes (SGR)
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html
var (
	Reset          = Sequence{Pre: "0"}  // Reset all Character Attributes to defaults
	Bold           = Sequence{Pre: "1"}  // Bold text
	Dim            = Sequence{Pre: "2"}  // Dim text, opposite of Bold
	Underline      = Sequence{Pre: "4"}  // Underline text
	Blink          = Sequence{Pre: "5"}  // Blink text at cursor rate
	Reverse        = Sequence{Pre: "7"}  // Reverse FG and BG colors
	Hidden         = Sequence{Pre: "8"}  // Hidden suppresses typed characters (e.g. password entry)
	ResetBold      = Sequence{Pre: "22"} // ResetBold turns off Bold/Dim formatting only
	ResetDim       = Sequence{Pre: "22"} // ResetDim turns off Bold/Dim formatting only
	ResetUnderline = Sequence{Pre: "24"} // ResetUnderline turns off Underline formatting only
	ResetBlink     = Sequence{Pre: "25"} // ResetBlink turns off Blink formatting only
	ResetReverse   = Sequence{Pre: "27"} // ResetReverse turns off Reverse formatting only
	ResetHidden    = Sequence{Pre: "28"} // ResetHidden turns off Hidden formatting only
)

// Sequence Represents one or more CSI Character Attributes (SGR)
type Sequence struct {
	// Pre is added at the start of the formatted string
	Pre string
	// Post is added at the end of the formatted string
	Post string
}

// ECMA-48 strings for constructing CSI SGR directives
const (
	Start     = "\033[" // Start is the CSI escape code
	Separator = ";"     // Separator allows multiple attributes to be strung together
	End       = "m"     // End terminates the CSI sequence and indicates that this is an SGR command
)

// Func forms a closure around a sequence which allows it to be called as a formatting function
func (s Sequence) Func() style.Style {
	return func(original string) string {
		if len(s.Pre) > 0 {
			original = Start + s.Pre + End + original
		}
		if len(s.Post) > 0 {
			original += Start + s.Post + End
		}
		return original
	}
}

// String generates the sequence, without and associated text to format
func (s Sequence) String() string {
	return s.Func()("")
}

// Swap exchanges the Pre and Post fields, useful for the last Sequence(s) in a combination
func (s Sequence) Swap() Sequence {
	return Sequence{Pre: s.Post, Post: s.Pre}
}

// Combine concatenates one or more Sequences together to form a single pre-calculated Sequence
func Combine(seqs ...Sequence) (next Sequence) {
	for _, seq := range seqs {
		if len(seq.Pre) > 0 {
			if len(next.Pre) == 0 {
				next.Pre = seq.Pre
			} else {
				next.Pre += Separator + seq.Pre
			}
		}
		if len(seq.Post) > 0 {
			if len(next.Post) == 0 {
				next.Post = seq.Post
			} else {
				next.Post += Separator + seq.Post
			}
		}
	}
	return
}
