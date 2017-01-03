package blocksmanager

import (
	"io"
	"os"
	"net/http"
	"encoding/json"
	"ioutil"
)

var folder string  = "blocks/"
func init() {
	currentid = 0
	currentsize = 0
	for {
		filename := folder + "/block" +  string(currentid) + ".dat"
		if fileinfo, err = os.Stat(filename); os.IsNotExist(err) {
			//文件不存在 
			break
		}else{
			currentsize = fileinfo.Size()
			currentid++
		}
	}
	currentid--
}

func OpenBlock(id int64) (*os.File,error) {
	var file *os.File
	filename := folder + "/block" + string(id) + ".dat"
	var err error
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		file, err = os.Create(filename) //创建文件
	} else {
		file, err = os.OpenFile(filename, O_APPEND)
	}
	if err != nil {
		return nil, err
	}
	return file
}

type SmallFile {
	Blockid int64
	Offset int64
	Size int64
}
var BlockFile *os.File = nil

func WriteBlock(job *Job)(*os.File,error){
	if currentsize > 100*1024*1024 {
		if BlockFile != nil{
			BlockFile.close()
		}
		BlockFile , err := OpenBlock(currentid + 1)
		if err != nil {
			return nil, err
		}
		currentid ++
                currentsize = 0
	}
	job.smallfile.Blockid = currentid
	job.smallfile.Offset = currentsize
	count , err := file.write((*job).Data)
	job.smallfile.Size = count
	currentsize += count
	return file , nil
}
