package main
//blog: xiaorui.cc
import (
    "github.com/robfig/cron"
    "log"
)

/*cron详解
* * * * * 要执行的命令
----------------
| | | | |
| | | | ---- 周当中的某天 (0 - 7) (周日为 0 或 7)
| | | ------ 月份 (1 - 12)
| | -------- 一月当中的某天 (1 - 31)
| ---------- 小时 (0 - 23)
------------ 分钟 (0 - 59)
*/
func main() {
    i := 0
    c := cron.New()
    spec := "0 0 16 * * *"
    c.AddFunc(spec, func() {
        i++
        log.Println("start", i)
    })
    c.Start()
    select{} //阻塞主线程不退出
}