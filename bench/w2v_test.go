package bnc

import (
	"os"
	"os/exec"
	"testing"

	"github.com/ynqa/word-embedding/builder"
)

func BenchmarkText8(b *testing.B) {
	w2v := builder.NewWord2VecBuilder()
	mod, _ := w2v.Build()
	file, _ := os.Open("text8")
	f, _ := mod.Preprocess(file)
	mod.Train(f)
}

func BenchmarkGensim(b *testing.B) {
	exec.Command("python", "gensim_w2v.py").Run()
}
