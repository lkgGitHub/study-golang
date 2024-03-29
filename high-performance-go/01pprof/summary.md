
# 性能分析类型
1. CPU 性能分析，runtime 每隔 10 ms 中断一次，记录此时正在运行的 goroutines 的堆栈信息
2. 内存性能分析，记录堆内存分配时的堆栈信息，忽略栈内存分配信息，默认每 1000 次采样 1 次
3. 阻塞性能分析，GO 中独有的，记录一个协程等待一个共享资源花费的时间
4. 锁性能分析，记录因为锁竞争导致的等待或延时

# 1. CPU 性能分析
使用原生 runtime/pprof 包，通过在 main 函数中添加代码运行可生成性能分析报告：

pprof.StartCPUProfile(os.Stdout)
defer pprof.StopCPUProfile()
可通过 go tool pprof -http=:9999 cpu.pprof 在 web 页面查看分析数据

可通过 go tool pprof cpu.prof 交互模式查看分析数据，可使用 help 查看支持的命令和选项

内存性能分析

使用 pkg/profile 库，通过在 main 函数中添加代码运行可生成性能分析报告：

defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
同样可通过 web 页面或交互模式查看分析数据

benchmark 生成 profile

可通过在 go test  中添加参数 -cpuprofile=$FILE,-memprofile=$FILE,-blockprofile=$FILE 生成相应的 profile 文件
生成的 profile 文件同样可通过 web 页面或交互模式查看分析数据