package model

type Produto struct {
	ID         string  `json:"id"`
	Tipo       string  `json:"tipo"`
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidadeEmEstoque"`
	Imagem     string  `json:"imagem"`
}
