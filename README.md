# GO 处理css

使用正则表达式分析页面css和id，提取到单独的文件。

## 匹配的正则规则

```
(id|class)=["'](.*?)["']
```

## 执行

```
go-prune-css -origin "code.html" -dist "./dist.txt"
```