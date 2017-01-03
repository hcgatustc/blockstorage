package httphandler

import (
	"fmt"
	"net/http"
	"blockstorage/blocksmanager"
	"io/ioutil"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	data,_ := ioutil.ReadAll(r.Body)
	if len(data) == 0 {
		return 
	}
	job, _ := blocksmanager.PostJob(data)
	
	fmt.Fprintf(w, "filename %s \n", <-job.Result)
}
