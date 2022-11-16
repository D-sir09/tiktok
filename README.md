## 上手指南

- 在 /tiktok 目录下创建 publish 目录，以存放上传的视频。
- 在 /dao/init.go 中修改自己的 mysql 连接参数
- 运行 `go mod tidy`

```shell
go run main.go router.go
```

### 说明

* 视频上传后会保存到本地 publish 目录中，访问时用 127.0.0.1:9999/static/video_name 即可
* [客户端参考][https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7]
* [各功能接口参考][https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145]

### 测试数据

使用 apache-ab 测试 

测试环境 i5-9300H@2.4G 4C8T

video表中有10条数据时：

```

ab -n 10000 -c 1 http://127.0.0.1:9999/douyin/feed/
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            9999

Document Path:          /douyin/feed/
Document Length:        3853 bytes

Concurrency Level:      1
Time taken for tests:   34.562 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      39560000 bytes
HTML transferred:       38530000 bytes
Requests per second:    289.33 [#/sec] (mean)
Time per request:       3.456 [ms] (mean)
Time per request:       3.456 [ms] (mean, across all concurrent requests)
Transfer rate:          1117.77 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     1    3   2.3      2      64
Waiting:        1    3   2.3      2      63
Total:          1    3   2.3      2      64

Percentage of the requests served within a certain time (ms)
  50%      2
  66%      4
  75%      4
  80%      5
  90%      7
  95%      8
  98%      9
  99%     10
 100%     64 (longest request)

```

