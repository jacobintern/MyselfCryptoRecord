package controller

import (
	"encoding/json"
	"log"
	"net/http"

	userService "github.com/jacobintern/MyselfCryptoRecord/service/user"
)

// 取得
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method Get User"))
}

// 新增
func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body userService.CreateUserModel
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}
	res, err := userService.UserCreate(&body)

	if err != nil {
		log.Fatal(err)
	}

	JosnRes(w, res)
}

// 更新
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method Put User"))
}

// 刪除
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method Delete User"))
}

func JosnRes(w http.ResponseWriter, model interface{}) {

	jsonRes, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Contect-Type", "application/json")
	w.Write(jsonRes)
}
