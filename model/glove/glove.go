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

import (
	"bufio"
	"io"
	"runtime"
	"strings"

	"github.com/chewxy/lingo/corpus"
	"github.com/pkg/errors"

	"github.com/ynqa/word-embedding/model"
)

type s int

func (sl s) Start() int { return int(sl) }
func (sl s) End() int   { return int(sl) + 1 }
func (sl s) Step() int  { return 0 }

type GloVe struct {
	*corpus.Corpus
	*model.Config
	PairMap
	FactorMatrix
	Epoch int
	XMax  int
	Alpha float64
	proc  int
}

func NewGloVe(config *model.Config, epoch int, alpha float64) *GloVe {
	return &GloVe{
		Corpus:  corpus.New(),
		Config:  config,
		PairMap: NewPairMap(),
		Alpha:   alpha,
		Epoch:   epoch,
		proc:    runtime.GOMAXPROCS(-1),
	}
}

func (g *GloVe) update(wordIDs []int) {
	for index := range wordIDs {
		for i := 0; i < g.Config.Window; i++ {
			if index+i >= len(wordIDs) {
				break
			}
			pair := NewPair(wordIDs[index], wordIDs[index+i])
			g.PairMap.Update(pair.former, pair.latter)
		}
	}
	return
}

func (g *GloVe) Preprocess(f io.ReadSeeker) (io.ReadCloser, error) {
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		line = strings.ToLower(line)
		words := strings.Fields(line)
		wordIDs := g.toIDs(words)
		g.update(wordIDs)
		if err := s.Err(); err != nil && err != io.EOF {
			return nil, errors.Wrap(err, "Unable to complete scanning.")
		}
	}
	return nil, nil
}

func (g *GloVe) Train(f io.ReadCloser) error {
	if g.FactorMatrix.W == nil {
		return errors.Errorf("No initialize model parameters")
	}

	f.Close()

	g.Corpus.Size()
	for i := 0; i < g.Epoch; i++ {
		for p, f := range g.PairMap {
			g.train(p, f)
		}
	}
	return nil
}

func (g *GloVe) train(p Pair, f int) (err error) {
	return nil
}

func (g *GloVe) toIDs(words []string) []int {
	retVal := make([]int, len(words))
	for i, w := range words {
		retVal[i], _ = g.Id(w)
	}
	return retVal
}

func (g *GloVe) toWords(wordIDs []int) []string {
	retVal := make([]string, len(wordIDs))
	for i, w := range wordIDs {
		retVal[i], _ = g.Word(w)
	}
	return retVal
}
