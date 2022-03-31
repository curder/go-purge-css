**当前项目已经归档，[更多查看](https://github.com/ava-cn/purge-css)**

# GoLang处理页面CSS

使用正则表达式分析页面css和id，提取到单独的文件。

## 正则规则

```
(id|class)=["'](.*?)["']
```

> 可以在这里验证正则匹配情况[regexr](https://regexr.com/)

## 安装

下载仓库中的二进制文件`go-purge-css`，添加到当前电脑的`$PATH`目录。

```
cp go-purge-css /usr/local/bin/.
```

> 授权执行权限 `chmod +x /usr/local/bin/go-purge-css`

重新启动一个命令行终端执行`go-purge-css`命令。

## 执行

```
go-purge-css -origin "code.html" -dist "./dist.txt"
```
> `-origin` 要修改的文件
> `-dist` 需要写入的文件
