package glove

import (
	"github.com/chewxy/gorgonia/tensor"
)

type FactorMatrix struct {
	W        tensor.Tensor
	poolGrad tensor.Tensor
}

func NewFactorMatrix(vocab, dim int) *FactorMatrix {
	return &FactorMatrix{
		// f, q, bias_f, bias_q
		W: tensor.New(tensor.WithShape(vocab * (dim + 1) * 2)),
		// store grad sqrt for AdaGrad
		poolGrad: tensor.New(tensor.WithShape(vocab * (dim + 1) * 2)),
	}
}

type Factor struct {
	Vector   tensor.Tensor
	PoolGrad tensor.Tensor
}

func (fm *FactorMatrix) getVector() {}
