package controllers

import (
	"html/template"
	"lojaGoingGo/models"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosProdutos := models.BuscarTodosProdutos()
	temp.ExecuteTemplate(w, "Index", todosProdutos)

}
