# GoLang处理页面CSS

使用正则表达式分析页面css和id，提取到单独的文件。

## 正则规则

```
(id|class)=["'](.*?)["']
```

> 可以在这里验证正则匹配情况[regexr](https://regexr.com/)

## 执行

```
go-prune-css -origin "code.html" -dist "./dist.txt"
```
> `-origin` 要修改的文件
- `-dist` 需要写入的文件