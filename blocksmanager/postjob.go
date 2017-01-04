package blocksmanager

var ToDo chan *Job = make(chan *Job)

func PostJob(data []byte) (*Job, error) {
	var job Job
	job.Data = data
	job.Result = make(chan string, 1)
	ToDo <- &job
	return &job, nil
}
