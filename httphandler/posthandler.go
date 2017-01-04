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
		w.WriteHeader(http.StatusBadRequest)
                fmt.Fprintf(w, "No Post Data Found")
		return
	}
	job, _ := blocksmanager.PostJob(data)
	
	fmt.Fprintf(w, "%s", <-job.Result)
}
