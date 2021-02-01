package main

import (
	"github.com/robin019/url-detection/router"
	_ "github.com/robin019/url-detection/src/psersistence/sql"
)

func main() {
	router.Route()
}
