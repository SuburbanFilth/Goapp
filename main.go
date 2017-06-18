package main

import ( 
		"html/template"
		"fmt"
		"net/http"
		//"io/ioutil"
		//"os"
		)

//Entry Point
func main() {
	http.ListenAndServe(":8080",New())
}

//This function returns our object
func New() http.Handler {
	var router = http.NewServeMux()
	var app = App{router}
	router.HandleFunc("/", app.index)
	router.HandleFunc("/register", app.CreateAccount)
	return app
}

//The main app object that holds the methods for the router and the actual router
type App struct { 
	router *http.ServeMux
}

//Implementing the Interface method
//NOTE> This is a non-pointer receiver. The rest of the methods are pointers
func (app App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w,r)
}

//Method responsible for the main page
func (app *App) index(w http.ResponseWriter, r *http.Request){
	var IndexTemplate,err = template.New("IndexTemplate").ParseFiles("Templates/base.html")
	//var temp, _ = ioutil.ReadFile("Templates/base.html")
	fmt.Println(IndexTemplate)
	if err != nil {
		w.Write([]byte("fsadsad"))
	} else {
		fmt.Println("dsa")
		IndexTemplate.Execute(w,nil)
	}
}

//Method responsible for the register page
func (app *App) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var bit = []byte("you should be able to create an account here. Not implemented")
	w.Write(bit)
}