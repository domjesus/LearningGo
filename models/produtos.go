package models

import (
	"fmt"

	"github.com/domjesus/webapp/db"
	_ "github.com/domjesus/webapp/db"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Id, Quantidade  int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	p := Produto{}
	produtos := []Produto{}

	query, err := db.Query("select * from produtos")

	if err != nil {
		fmt.Println("Erro ao selecionar produtos: ", err.Error())
	} else {
		// fmt.Println("Query Ok!", query)
		for query.Next() {
			var id, quantidade int
			var nome, descricao string
			var preco float64

			err = query.Scan(&id, &nome, &descricao, &preco, &quantidade)

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

	}

	defer db.Close()

	return produtos

}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()
	// p := Produto{Nome: nome, Descricao: desc, Preco: preco, Quantidade: qtde}

	insereDadosSql, err := db.Prepare("insert into produtos (nome, descricao, preco,quantidade) values($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosSql.Exec(nome, descricao, preco, quantidade)

	// fmt.Println(p)

}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	query, err := db.Prepare("DELETE FROM produtos WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	query.Exec(id)
}
