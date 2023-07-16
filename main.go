package main

import (
	"fmt"
	go_module_name "github.com/fauziihsan/go-module-name/v2"
	alamat "github.com/fauziihsan/app-say-hello/helpers"
	db "github.com/fauziihsan/app-say-hello/config"
)

func main() {
	fmt.Println(go_module_name.SayName("Fauzi"))
	fmt.Println(alamat.GetAlamat("Siak"))
	fmt.Println(db.Getdb("DB_GOLANG"))
}