# gorn

# 知识点

* 快速使用 
-   go get :go get github.com/roylee0704/gron
-   创建：gron.New()
-   添加任务：AddFunc()或Add()
    -   无参数启动、实现任务接口
-   启动：Start()
-   时间间隔：gron.Every()
-   保证主 goroutine 不退出：sync.WaitGroup

* 时间格式
- 类型：time.Duration、Second/Minute/Hour、Day/Week
- 精度：1s 最小
- 多久执行一次：gron.Every()
- 某个时间点执行：gron.Every().At()

* 自定义任务
-   实现接口：gron.Job
-   添加自定义任务：Add()