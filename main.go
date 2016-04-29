package main

import (
	"fmt"
	//"github.com/hoisie/web"
	//"io/ioutil"
	//"log"
	//"net/http"
	//"os"
	//"runtime"
	"time"
)

func main() {

	gotime := time.Now().Unix()
	fmt.Println(fmt.Sprintf("DwarfGUI setting up. Unixtime:%v", gotime))

	remote()

	site_initialize()
}
