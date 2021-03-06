# Model

## Word2Vec

Word2Vec is the generic term below modules:

```
model:
- Skip-Gram
- CBOW

optimizer:
- Hierarchical Softmax
- Negative Sampling
```

In training, select one `model` and one `optimizer` above. `model` and `optimizer` represent architecture of objective and the way of approximating its function respectively.

### Features

- [x] Skip-Gram
- [x] CBOW
- [x] Hierarchical Softmax
- [x] Negative Sampling
- [x] Subsampling
- [x] Update learning rate in training

### Usage

```
Embed words using word2vec

Usage:
  word-embedding word2vec [flags]

Flags:
      --batchSize int       Set the batch size to update learning rate (default 10000)
  -d, --dimension int       Set the dimension of word vector (default 10)
      --initlr float        Set the initial learning rate (default 0.025)
  -i, --inputFile string    Set the input file path to load corpus (default "example/input.txt")
      --lower               Whether the words on corpus convert to lowercase or not (default true)
      --maxDepth int        Set the number of times to track huffman tree, max-depth=0 means to track full path from root to word (using only hierarchical softmax)
      --model string        Set the model of Word2Vec. One of: cbow|skip-gram (default "cbow")
      --optimizer string    Set the optimizer of Word2Vec. One of: hs|ns (default "hs")
  -o, --outputFile string   Set the output file path to save word vectors (default "example/word_vectors.txt")
      --sample int          Set the number of the samples as negative (using only negative sampling) (default 5)
      --theta float         Set the lower limit of learning rate (lr >= initlr * theta) (default 0.0001)
      --threshold float     Set the threshold for subsampling (default 0.001)
  -w, --window int          Set the context window size (default 5)
```
