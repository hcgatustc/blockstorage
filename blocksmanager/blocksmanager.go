package blocksmanager

import (
	"os"
)

var folder string  = "blocks/"
var currentid int64
var currentsize int64
func init() {
	currentid = 0
	currentsize = 0
	for {
		filename := folder + "/block" +  string(currentid) + ".dat"
		if fileinfo, err := os.Stat(filename); os.IsNotExist(err) {
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
		file, err = os.OpenFile(filename, os.O_APPEND, 0666)
	}
	if err != nil {
		return nil, err
	}
	return file,nil
}

type SmallFile struct {
	Blockid int64
	Offset int64
	Size int64
}
var BlockFile *os.File = nil

func WriteBlock(job *Job)error{
	if currentsize > 100*1024*1024 {
		if BlockFile != nil{
			BlockFile.Close()
		}
		var err error
		BlockFile , err = OpenBlock(currentid + 1)
		if err != nil {
			return err
		}
		currentid ++
                currentsize = 0
	}
	job.File.Blockid = currentid
	job.File.Offset = currentsize
	count , _ := BlockFile.Write(job.Data)
	job.File.Size = int64(count)
	currentsize += int64(count)
	return  nil
}
