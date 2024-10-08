package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// )

type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	IsCompleted bool `json:"isCompleted"`
}

// サイズを指定していないから、スライス
var todos = []Todo {
	{Id: 1, Name: "Buy milk", IsCompleted: false},
	{Id: 2, Name: "Buy eggs", IsCompleted: true},
	{Id: 3, Name: "Buy bread", IsCompleted: false},
}


var templates = map[string]*template.Template{}

// 初期化に使用される関数 mainよりも早く呼ばれる
func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index.html"] = template.Must(template.ParseFiles("index.html"))
	templates["todo.html"] = template.Must(template.ParseFiles("todo.html"))
}

// todoがpostされた時の処理
func submitTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	isCompleted := r.PostFormValue("completed") == "true"	
	todo:= Todo{Id: len(todos) + 1, Name: name, IsCompleted: isCompleted}
	
	todos = append(todos,todo)
	tmpl := templates["todo.html"]
	// これtodosじゃなくていいのか
	tmpl.ExecuteTemplate(w,"todo.html",todo)
	
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	json,err := json.Marshal(todos)
	
	if err != nil {
		log.Fatal(err)
	}
	tmpl := templates["index.html"]
	// jsonでtodosを渡す
	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "index.html", map[string]template.JS{"Todos": template.JS(json)})
}

func main() {
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/submit-todo/",submitTodoHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}