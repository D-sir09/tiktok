## 上手指南

- 在 /tiktok 目录下创建 publish 目录，以存放上传的视频。
- 在 /dao/init.go 中修改自己的 mysql 连接参数
- 运行 `go mod tidy`

```shell
go run main.go router.go
```

### 说明

* 视频上传后会保存到本地 publish 目录中，访问时用 127.0.0.1:9999/static/video_name 即可

### 测试数据

使用 go-stress-testing 测试 

测试环境 i5-9300H@2.4G 4C8T

```
并发数:10000 请求数:1 请求参数: 
request:
 form:http 
 url:http://127.0.0.1:9999/douyin/feed 
 method:GET 
 headers:map[Content-Type:application/x-www-form-urlencoded; charset=utf-8] 
 data: 
 verify:statusCode 
 timeout:30s 
 debug:false 
 http2.0：false 
 keepalive：false 
 maxCon:1 


─────┬───────┬───────┬───────┬────────┬────────┬────────┬────────┬────────┬────────┬────────
 耗时│ 并发数│ 成功数│ 失败数│   qps  │最长耗时│最短耗时│平均耗时│下载字节│字节每秒│ 状态码
─────┼───────┼───────┼───────┼────────┼────────┼────────┼────────┼────────┼────────┼────────
   1s│      0│      0│      0│    0.00│    0.00│    0.00│    0.00│        │        │
   2s│    214│    214│      0│ 5912.63│ 1919.49│ 1520.69│ 1691.30│  58,636│  29,317│200:214
   3s│   4252│   4252│      0│ 4500.08│ 2680.50│ 1520.69│ 2222.18│ 683,693│ 227,867│200:4252
   4s│  10000│  10000│      0│ 3768.73│ 3603.50│ 1520.69│ 2653.41│1,783,042│ 476,585│200:10000


*************************  结果 stat  ****************************
处理协程数量: 10000
请求总数（并发数*请求数 -c * -n）: 10000 总请求时间: 3.741 秒 successNum: 10000 failureNum: 0
tp90: 3157.000
tp95: 3195.000
tp99: 3291.000
*************************  结果 end   ****************************

```

