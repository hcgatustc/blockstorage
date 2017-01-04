package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ConfigType struct {
	BlockFolder    string
	MaxBlockSize   int64
	MaxConcurrency int64
}

var Config ConfigType

func InitDefaultConfig() {
	Config.BlockFolder = "blocks"
	Config.MaxBlockSize = 100 * 1024 * 1024
	Config.MaxConcurrency = 1000
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	file, err := os.OpenFile("config.json", os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		log.Fatalf("Open Config File Error : %s\n", err.Error())
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Read Config File Error : %s\n", err.Error())
	}
	if len(data) == 0 {
		InitDefaultConfig()
		data, err = json.MarshalIndent(Config, "", " ")
		if err != nil {
			log.Fatalf("Encode Config File Error : %s\n", err.Error())
		}
		_, err = file.Write(data)
		if err != nil {
			log.Fatalf("Write Default Config Error : %s\n", err.Error())
		}
		/*		err = file.Sync()
				if err != nil {
					log.Fatalf("Sync Config Data To Disk Error : %s\n", err.Error())
				}*/
	} else {
		err = json.Unmarshal(data, &Config)
		if err != nil {
			log.Fatalf("Config File Format Error : %s\n", err.Error())
		}
	}
}
