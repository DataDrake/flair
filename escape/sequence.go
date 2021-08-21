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

var (
	Reset          = Sequence{Pre: "0"}
	Bold           = Sequence{Pre: "1"}
	Dim            = Sequence{Pre: "2"}
	Underline      = Sequence{Pre: "4"}
	Blink          = Sequence{Pre: "5"}
	Reverse        = Sequence{Pre: "7"}
	Hidden         = Sequence{Pre: "8"}
	ResetBold      = Sequence{Pre: "22"}
	ResetDim       = Sequence{Pre: "22"}
	ResetUnderline = Sequence{Pre: "24"}
	ResetBlink     = Sequence{Pre: "25"}
	ResetReverse   = Sequence{Pre: "27"}
	ResetHidden    = Sequence{Pre: "28"}
)

type Sequence struct {
	Pre  string
	Post string
}

const (
	start     = "\033["
	Separator = ";"
	end       = "m"
)

func (s Sequence) Func() style.Style {
	return func(original string) string {
		if len(s.Pre) > 0 {
			original = start + s.Pre + end + original
		}
		if len(s.Post) > 0 {
			original += start + s.Post + end
		}
		return original
	}
}

func (s Sequence) String() string {
	return s.Func()("")
}

func (s Sequence) Swap() Sequence {
	return Sequence{Pre: s.Post, Post: s.Pre}
}

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
