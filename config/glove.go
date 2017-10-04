// Copyright © 2017 Makoto Ito
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

type GloVeConfig int

const (
	Iteration GloVeConfig = iota
	Alpha
	XMax
)

const (
	DefaultIteration int     = 50
	DefaultAlpha     float64 = 0.75
	DefaultXMax      int     = 100
)

func (g GloVeConfig) String() string {
	switch g {
	case Iteration:
		return "iter"
	case Alpha:
		return "alpha"
	case XMax:
		return "xmax"
	default:
		return "unknown"
	}
}
