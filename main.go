package main

import (
	"app/pkg/config"
	"fmt"
)

func main() {

	fmt.Println("The app name:", config.Cfg.GetString("appName"))

}
