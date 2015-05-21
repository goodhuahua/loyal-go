package main

import (
	"log"
	"html/template"
	"net/http"
	"io"
)

func main() {
	log.Println("正在开启服务。。。")
	err := http.ListenAndServe(":8080", handlerMux())
	if err!=nil {
		log.Println("服务启动失败。。。")
		log.Println(err.Error())
	}
	log.Println("服务已启动，正在监听888端口...")
}

func handlerMux() (*http.ServeMux) {
	//创建Mux实例
	mux := http.NewServeMux()

	//静态文件处理
	mux.Handle("/assets/", http.FileServer(http.Dir("")))

	//创建Mux实例
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/login", loginHandler)

	//创建Mux实例
	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Println(err.Error())
	}
	w.Header().Set("Content-Type", "text/html")
	err1 := t.Execute(w, nil)
	if err1 != nil {
		log.Println(err1)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "登录")
}

