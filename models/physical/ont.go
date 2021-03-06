/*
   Copyright 2017 the original author or authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package physical

/*
Ont represents a single ont/onu connect to a splitter on a Port
*/
type Ont struct {
	Number       int      `json:",omitempty"`
	Svlan        uint32   `,json:",omitempty"`
	Cvlan        uint32   `,json:",omitempty"`
	SerialNumber string   `,json:",omitempty"`
	Parent       *PONPort `json:"-" bson:"-"`
	Active       bool     `json:",omitempty"`
	NasPortID    string   `json:",omitempty"`
	CircuitID    string   `json:",omitempty"`
	TechProfile  string   `json:",omitempty"`
	SpeedProfile string   `json:",omitempty"`
}
