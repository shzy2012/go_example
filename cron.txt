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

/*
	Entry                  | Description                                | Equivalent To
	-----                  | -----------                                | -------------
	@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
	@monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
	@weekly                | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
	@daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
	@hourly                | Run once an hour, beginning of hour        | 0 0 * * * *
*/

/*
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
*/
func job(){
    i := 0
    c := cron.New()
    spec := "0 0 15 * * *" //19点 5-10分之内，每分钟发次
    c.AddFunc(spec, func() {
        i++
        log.Println("start", i)
    })
    c.Start()
}
func main() {
    
    go job()
    select{} //阻塞主线程不退出
}