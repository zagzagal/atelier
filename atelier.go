package main

import (
	"flag"
)

func main() {
	filePath := flag.String("file", "", "dot file name")
	ui := flag.String("ui", "cli", "Ui selection, web or cli")
	flag.Parse()
	switch *ui {
	case "cli":
		startCLIUI(*filePath)
	case "web":
		webui(*filePath)
	case "default":
	}
}
