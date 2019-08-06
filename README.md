# 简介

这是一个用 golang 实现的外部排序的示例

`crtdata.go` 中的 main 用来创建数据

`sort.go`中的 main 为入口
在下一个版本中将进行重构

整个的使用流程为：

1. 创建数据
2. 用 几个 goroutine 来读取数据并在 memory 中进行排序
3. 然后将结果放入到 server 中去，等待客户端来取数
4. 开启了几个 client 用来合并上一个在 memory 中的排序的结构
5. 最后再输出到外部文件
 