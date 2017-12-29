# 01-print-log-using-stardard-library

### 使用Go的标准库`fmt`

1. 使用`fmt.Printf()`方法打印的日志在没有换行符`\n`情况下，stackdriver Logging 输出日志如下：

![](https://raw.githubusercontent.com/mrdulin/pic-bucket-01/master/WX20190703-124942.png)

可以看到，只打印出了一行日志，这行日志的`textPayload`的字段包含了所有日志内容。显然，这样的日志可读性很差。

2. 使用`fmt.Printf()`方法打印的日志在有换行符`\n`情况下，stackdriver Logging 输出日志如下：

![](https://raw.githubusercontent.com/mrdulin/pic-bucket-01/master/WX20190703-130242.png)

可以看到，加入换行符以后，每条日志打印一行，有较好的可读性，很容易区分每条日志。

### 使用Go的标准库`log`

1. 使用`log.Printf()`方法打印的日志在有换行符`\n`情况下，stackdriver Logging 输出日志如下：

![](https://raw.githubusercontent.com/mrdulin/pic-bucket-01/master/WX20190703-133308.png)

可以看到，加入换行符以后，每条日志打印一行，有较好的可读性，很容易区分每条日志，但是又额外打印出了时间戳，和stackdriver Logging提供的日志时间戳功能重复。

### 使用GCP提供的logging package

打印的日志即可以是简单的字符串`textPayload`字段，也可以是结构化日志`jsonPayload`， 并且可以设置log level，如下：

![](https://raw.githubusercontent.com/mrdulin/pic-bucket-01/master/WX20190703-152125.png)

### 参考

- https://cloud.google.com/logging/docs/setup/go
- https://godoc.org/cloud.google.com/go/logging
- https://github.com/GoogleCloudPlatform/cloud-functions-go/issues/13#issuecomment-323556031