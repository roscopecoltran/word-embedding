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

package glove

type PairMap map[Pair]int
type Pair struct {
	former, latter int
}

func NewPairMap() PairMap {
	return make(PairMap)
}

func NewPair(w1, w2 int) *Pair {
	f, l := order(w1, w2)
	return &Pair{
		former: f,
		latter: l,
	}
}

func order(w1, w2 int) (f, l int) {
	if w1 <= w2 {
		f = w1
		l = w2
	} else {
		f = w2
		l = w1
	}
	return
}

func (p PairMap) Update(w1, w2 int) {
	pp := NewPair(w1, w2)
	p[*pp]++
}

func (p PairMap) Get(w1, w2 int) int {
	pp := NewPair(w1, w2)
	return p[*pp]
}
