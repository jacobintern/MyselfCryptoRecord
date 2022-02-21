package controllers

import (
	"fmt"
	"net/http"
)

// 取得
func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method Get User")
}

// 新增
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method Post User")
}

// 更新
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method Put User")
}

// 刪除
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method Delete User")
}
