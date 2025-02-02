// Copyright 2020-2021 Matt Layher
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netstate

// A Change is a bitmask of possible changes to a network interface's state.
type Change uint

// Possible state changes which may occur to a network interface.
const (
	// RFC 2863 "ifOperStatus" values which indicate network interface
	// state changes.
	LinkUp Change = 1 << iota
	LinkDown
	LinkTesting
	LinkUnknown
	LinkDormant
	LinkNotPresent
	LinkLowerLayerDown

	// LinkAny is a convenience bitmask which indicates interest in all
	// Link* Changes.
	LinkAny Change = LinkUp |
		LinkDown |
		LinkTesting |
		LinkUnknown |
		LinkDormant |
		LinkNotPresent |
		LinkLowerLayerDown
)

// changeStrings is a list of string values for a Change.
var changeStrings = []string{
	"link up",
	"link down",
	"link testing",
	"link unknown",
	"link dormant",
	"link not present",
	"link lower layer down",
}

// String returns the string representation of a Change.
func (c Change) String() string {
	// Special case values.
	switch c {
	case LinkAny:
		return "link ANY"
	}

	// Concatenate bitmask string names.
	s := ""
	for i, name := range changeStrings {
		if c&(1<<uint(i)) != 0 {
			if s != "" {
				s += "|"
			}
			s += name
		}
	}

	if s == "" {
		s = "0"
	}

	return s
}
