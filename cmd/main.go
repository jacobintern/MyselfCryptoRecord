package main

import (
	"log"
	"net/http"

	routes "github.com/jacobintern/MyselfCryptoRecord/router"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../.env")
}

func main() {
	// data, err := json.Marshal(coinmarketcapapi.GetcryptoDataList().Data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", data)
	//coinmarketcapapi.MappingMyList()

	router := routes.NewRouter()

	err := http.ListenAndServe(":2000", router)

	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
