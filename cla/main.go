package main
import (
	"cfa/native/config"
	"os"
)

func report(value string){}

func main(){
	config.FetchAndValid(os.Args[1],os.Args[2],true,report)
}

