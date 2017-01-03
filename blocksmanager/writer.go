package blocksmanager


func Writer(){
	for {
		select {
			case job := <- ToDo:
				WriteBlock(job)
				filename, err := GetFileName(job.File)					
				if err!= nil {
					filename = ""
				}
				job.Result <- filename
		
		}
	}
}
