package blocksmanager

import (
	"io"
	"os"
	"net/http"
	"log"
)

func Writer(){
	file,err := openblock()
	if err != nil {
		log.Fatal("open block error")
		return
	}
	for {
		select {
			case job := <- ToDo:
				WriteBlock(job)
				
		
		}
	}
}
