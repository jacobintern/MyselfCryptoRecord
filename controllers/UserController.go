package controllers

import (
	"net/http"
)

// 取得
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method Get User"))
}

// 新增
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method Post User"))
}

// 更新
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method Put User"))
}

// 刪除
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method Delete User"))
}
