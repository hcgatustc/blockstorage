package blocksmanager

import (
	"blockstorage/config"
	"log"
	"os"
	"strconv"
)

var currentid int64
var currentsize int64

func init() {
	currentid = 0
	currentsize = 0
	var err error
	for {
		filename := config.Config.BlockFolder + "/block" + strconv.FormatInt(currentid, 10) + ".dat"
		if fileinfo, err := os.Stat(filename); os.IsNotExist(err) {
			//文件不存在
			currentid--
			break
		} else {
			log.Printf("Found Block %s , %d", filename, fileinfo.Size())
			//文件存在
			currentsize = fileinfo.Size()
			currentid++
		}
	}
	if currentid < 0 {
		currentid = 0
	}
	BlockFile, err = OpenBlockForWrite(currentid)
	if err != nil {
		log.Fatal("OpenBlockForWrite ", currentid, " error")
	}
	log.Printf("Init Success currentid %d currentsize %d\n", currentid, currentsize)
	go Writer()
}

func OpenBlockForRead(id int64, offset int64) (*os.File, error) {
	var file *os.File
	var err error
	filename := config.Config.BlockFolder + "/block" + strconv.FormatInt(id, 10) + ".dat"
	file, err = os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	_, err = file.Seek(offset, os.SEEK_SET)
	if err != nil {
		file.Close()
		return nil, err
	}
	return file, nil
}

func OpenBlockForWrite(id int64) (*os.File, error) {
	var file *os.File
	var err error
	filename := config.Config.BlockFolder + "/block" + strconv.FormatInt(id, 10) + ".dat"
	file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_SYNC, 0666)
	if err != nil {
		return nil, err
	}
	_, err = file.Seek(0, os.SEEK_END)
	if err != nil {
		file.Close()
		return nil, err
	}
	return file, nil
}

type SmallFile struct {
	Blockid int64
	Offset  int64
	Size    int64
}

var BlockFile *os.File = nil

func WriteBlock(job *Job) error {
	//log.Printf("Before Write Block %d Size is %d\n",currentid,currentsize)
	var err error
	if currentsize > config.Config.MaxBlockSize {
		log.Printf("Block %d Full Size is %d\n", currentid, currentsize)
		if BlockFile != nil {
			BlockFile.Close()
		}
		BlockFile, err = OpenBlockForWrite(currentid + 1)
		if err != nil {
			return err
		}
		currentid++
		currentsize = 0
	}
	job.File.Blockid = currentid
	job.File.Offset = currentsize
	count, err := BlockFile.Write(job.Data)
	if err != nil {
		log.Printf("Write BlockFile Failed %s\n", err.Error())
		return err
	}
	/*err = BlockFile.Sync()
	if err != nil {
		log.Printf("Sync BlockFile Failed %s\n", err.Error())
		return err
	}*/
	job.File.Size = int64(count)
	currentsize += int64(count)
	return nil
}
