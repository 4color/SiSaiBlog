## 简介
《四嗮小博客》是一款采用Golang开发的博客程序，它只包含最基本的博客发布功能，评论模块采用搜狐畅言模块。
## 为什么
为什么这个博客程序只有这么少的功能？因为对我来说，能有一个地方独立发布博客就行了，然后还满足了自己的修改欲望。

之前也采用Wordpress等架设博客，但是太庞大，耗资源，老是被服务商停机。正好Golang是一个能快速开发的语言，采用Gin框架，我3天就写好了这个功能。


## 将来

将来要加什么功能，还不知道！


## 源代码编译使用
> 一共四步
前提：请先安装golang的sdk

1. 从Github下克隆或下载代码
2. 设置goproxy代理，推荐用http://goproxy.cn
3. 运行go mod download 下载依赖包
4. 运行go build 生成安装包  （由于数据库采用了sqlite，所以不能交叉编译）

windows下编译生成exe程序，linux下编译生成二进制运行包
其它静态文件和配置只要
- cmd
- configs
- web

三个文件夹即可，删除里面的go文件


### 配置修改
找到文件：configs\yaml\application.yaml

## 官方站点
http://4color.cn

## Github
https://github.com/4color/SiSaiBlog/
欢迎Follow


