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