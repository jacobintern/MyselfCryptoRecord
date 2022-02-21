package main

import (
	"log"
	"net/http"

	"github.com/jacobintern/MyselfCryptoRecord/controllers"
)

func main() {
	// data, err := json.Marshal(coinmarketcapapi.GetcryptoDataList().Data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", data)
	//coinmarketcapapi.MappingMyList()

	http.HandleFunc("api/users", controllers.GetUser)
	err := http.ListenAndServe(":2000", nil)

	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
