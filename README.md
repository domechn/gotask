# gotask

### 轮询任务

```go
package main

import (
    "time"
    
    "github.com/domgoer/gotask"
)

func main()  {
     tk := gotask.NewTask(time.Second*20,func() {
            // do ... 
     })
     gotask.AddToTaskList(tk)
}
```

#### 注意

当task的轮询速度快于执行速度时，需要设置合理的结束时间，来防止gorouting的泄露。

### 动态修改轮询任务的执行时间

```go
package main

import (
    "time"
    
    "github.com/domgoer/gotask"
)

func main()  {
     tk := gotask.NewTask(time.Second*20,func() {
            // do ... 
     })
     gotask.AddToTaskList(tk)
     
     // 修改执行时间，立即生效
     tk.SetInterval(time.Second*30)
}
```

### 定时任务

```go
package main

import (
    "github.com/domgoer/gotask"
)

func main()  {
     tkDay,_ := gotask.NewDayTask("12:20:00",func() {
            // do ... 
     })
     tkMonth,_ := gotask.NewMonthTask("20 12:20:00",func() {
             // do ... 
      })
     gotask.AddToTaskList(tkDay)
     gotask.AddToTaskList(tkMonth)
}
```

> 多任务

```go
package main

import (
    "github.com/domgoer/gotask"
)

func main()  {
     tkDays,_ := gotask.NewDayTasks([]string{"12:20:00","10:10:10"},func() {
            // do ... 
     })
     tkMonths,_ := gotask.NewMonthTasks([]string{"20 12:20:00","21 10:10:10"},func() {
             // do ... 
      })
     gotask.AddToTaskList(tkDays...)
     gotask.AddToTaskList(tkMonths...)
}
```

### 停止

```go
package main

import (
    "github.com/domgoer/gotask"
)

func main()  {
     gotask.Stop("task.ID()")
}
```
