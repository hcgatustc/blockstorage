package blocksmanager

func Writer(){
	for {
		select {
			case job := <- ToDo:
				var err error
				var filename string
				err = WriteBlock(job)
				if err != nil {
					job.Result<-""
					continue
				}
				filename, err = GetFileName(job.File)					
				if err!= nil {
					job.Result<-""
					continue
				}
				job.Result <- filename
		
		}
	}
}
