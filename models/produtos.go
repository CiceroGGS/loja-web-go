package models

import (
	"lojaGoingGo/db"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscarTodosProdutos() []Produto {
	db := db.ConectarComDB()

	selectProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id int
		var nome, descricao string
		var preco float64
		var quantidade int

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome string, descricao string, preco float64, quamtidade int) {
	db := db.ConectarComDB()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quamtidade)

	defer db.Close()
}

func DeletarProdutoPorId(id string) {
	db := db.ConectarComDB()

	deletarProduto, err := db.Prepare("DELETE FROM produtos WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)

	defer db.Close()
}

func EditarProduto(id string) Produto {
	db := db.ConectarComDB()

	produto, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoAtualizado := Produto{}

	for produto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoAtualizado.Id = id
		produtoAtualizado.Nome = nome
		produtoAtualizado.Descricao = descricao
		produtoAtualizado.Preco = preco
		produtoAtualizado.Quantidade = quantidade
	}
	defer db.Close()
	return produtoAtualizado
}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectarComDB()

	AtualizaProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
