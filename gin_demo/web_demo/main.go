package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Student struct {
	Name   string
	Age    int
	Gender string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//获取模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	Stu01 := Student{
		Name:   "猪八戒",
		Age:    20,
		Gender: "男",
	}

	//解析模板
	err = t.Execute(w, Stu01)
	if err != nil {
		fmt.Printf("template execute failed, err:%v\n", err)
		return
	}

}

func main() {

	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}

}
