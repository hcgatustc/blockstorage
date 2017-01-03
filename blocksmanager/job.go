package blocksmanager

type Job struct{
       Data []byte
       Result chan string //应为缓存为1的 不能阻塞
       File SmallFile
}

