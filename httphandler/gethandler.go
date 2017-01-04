package httphandler

import (
	"fmt"
	"net/http"
	"blockstorage/blocksmanager"
	"strings"
	"io"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	slices := strings.Split(r.URL.Path,"/")
	if len(slices) < 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found")         	
		return
	}
	filename :=  slices[len(slices)-1]
	filename = strings.Split(filename,".")[0]
	var err error
	smallfile, err := blocksmanager.GetSmallFile(filename)
	if err != nil {
                w.WriteHeader(http.StatusBadRequest)
                fmt.Fprintf(w, err.Error())
                return
	}
	file ,err :=  blocksmanager.OpenBlockForRead(smallfile.Blockid,smallfile.Offset)
        if err != nil {
                w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
                return
        }
	_, err = io.CopyN(w,file,smallfile.Size)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
                fmt.Fprintf(w, err.Error())
                return
	}
}
