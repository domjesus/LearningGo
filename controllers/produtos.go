package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/domjesus/webapp/db"
	"github.com/domjesus/webapp/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.BuscaTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", produtos)

}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))

		models.CriarNovoProduto(nome, descricao, preco, quantidade)

	}
	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idToDelete := r.URL.Query().Get("id")

	models.DeletaProduto(idToDelete)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProd := r.URL.Query().Get("id")
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	produto, err := db.Prepare("SELECT * FROM produtos WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	produto.Exec(idProd)

	temp.ExecuteTemplate(w, "Edit", produto)
}
