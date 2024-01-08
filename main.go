package main

import (
	"os"

	"github.com/Vodka479/space-shop-tutorial/config"
	"github.com/Vodka479/space-shop-tutorial/modules/servers"
	"github.com/Vodka479/space-shop-tutorial/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 { /*Args คือ argument*/
		return ".env"
	} else {
		return os.Args[1] /*1 คือ Args ตัวที่ 2*/
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close() // defer จะทำงานท้ายสุดก่อนจะ return กลับไป

	servers.Newserver(cfg, db).Start()
}
