package controllers

import (
	"html/template"
	"log"
	"lojaGoingGo/models"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosProdutos := models.BuscarTodosProdutos()
	temp.ExecuteTemplate(w, "Index", todosProdutos)

}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preco", err)
		}

		quantidadeParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao do quantidade", err)
		}
		models.CriarNovoProduto(nome, descricao, precoParaFloat, quantidadeParaInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletarProdutoPorId(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditarProduto(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na concersao do ID para int:")
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na concersao do preco para float:")
		}

		quantidadeConvertidoParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na concersao do quantidade para int:")
		}

		models.AtualizarProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidoParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
