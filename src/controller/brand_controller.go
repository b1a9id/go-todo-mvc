package controller

import (
	"net/http"
	"database/sql"
	"html/template"
	"github.com/labstack/gommon/log"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/b1a9id/go-todo-mvc/src/model"
		_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func GetBrands(w http.ResponseWriter, r *http.Request)  {
}

func GetBrand(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	brand, err := model.BrandByID(DB, uint(id))

	if err != nil {
		panic(err.Error())
	}

	t := template.Must(template.ParseFiles("src/view/describe.html"))
	if err := t.ExecuteTemplate(w, "describe.html", brand); err != nil {
		log.Fatal(err.Error())
	}
}

func CreateBrand(w http.ResponseWriter, r *http.Request)  {
	
}

func UpdateBrand(w http.ResponseWriter, r *http.Request)  {
	
}

func DeleteBrand(w http.ResponseWriter, r *http.Request)  {
	
}

func InitDb() {
	dbDriver := "mysql"
	dbUser := "root"
	dbName := "gtm"
	dbOption := "?parseTime=true"
	db, err := sql.Open(dbDriver, dbUser + "@/" + dbName + dbOption)
	if err != nil {
		log.Fatal(err.Error())
	}
	DB = db
}
