package concurrent

import (
	"fmt"
	"math/rand"
	"time"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

// goroutine池演示
// 打印100个随机各位的和
func goPoolDemo() {
	// 需要2个管道
	// job管道
	jobCh := make(chan *Job, 128)
	// 结果管道
	resCh := make(chan *Result, 128)
	// 创建工作池
	createPool(64, jobCh, resCh)
	num := 1000
	// 开个打印的协程
	var end bool
	go func(resCh chan *Result) {
		var cnt int
		// 遍历结果管道打印
		for result := range resCh {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
			cnt++
			if cnt == num {
				break
			}
		}
		end = true
	}(resCh)
	var id int
	// 循环创建job，输入到管道
	for i := 0; i < num; i++ {
		id++
		// 生成随机数
		rNum := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: rNum,
		}
		jobCh <- job
	}
	for !end {
		time.Sleep(time.Millisecond * 10)
	}
}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobCh chan *Job, resCh chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobCh chan *Job, resCh chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobCh {
				// 随机数接过来
				rNum := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for rNum != 0 {
					tmp := rNum % 10
					sum += tmp
					rNum /= 10
				}
				// 想要的结果是Result
				res := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resCh <- res
			}
		}(jobCh, resCh)
	}
}
