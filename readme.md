工具运行的端口在 `main.go`里面修改，默认是12321

具体使用可观察如下命令：

```
curl -G -d 'user=用户' -d 'prompt=你好，我是来调用你的' 127.0.0.1:12321/gpt
```

直接返回接口的文本，请求失败会返回的对应错误消息请在 `api.go` 里面看

然后是调用CV的：

```
curl -X POST -d '{
	"content":"这里放简历文本"
	}'	127.0.0.1:12321/cv
```
