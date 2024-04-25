package model

type Carrinho struct {
	ID              string
	UserId          string
	ProdutoInfo     []ProdutoInfos
	PrecoTotal      float64
	QuantidadeTotal int
}

type ProdutoInfos struct {
	ProdutoID     string
	ProductAmount int
}
