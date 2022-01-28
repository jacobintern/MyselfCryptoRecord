package main

import (
	"github.com/jacobintern/MyselfCryptoRecord/coinmarketcapapi"
)

func main() {
	// data, err := json.Marshal(coinmarketcapapi.GetcryptoDataList().Data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", data)
	coinmarketcapapi.MappingMyList()
}
