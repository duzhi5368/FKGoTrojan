# 简介

>  Real programmers don't comment their code. 
>  if it was hard to write,
>  it should be hard to understand and even harder to modify.
>  Real programmers don't read manuals. reliance on a reference is
>  the hallmark of a novice and a coward.
>> *--- FreeKnight*



该大马包括客户端和服务器两部分。
- 客户端部分直接运行Client.go，通过小马上传部署运行在目标肉鸡机器上，实现各种非常规功能。
- 服务器部分直接运行Server.go，对受木马感染的肉鸡机器进行远程控制。

一个服务器可管理多个木马客户端，以Http/Https命令方式控制木马行为，并通过Http/Https方式接收木马发送的数据信息。

## Client

Client使用代码包括Client.go以及components目录内全部源码。


##### 基本性功能 

- 看门狗进程保护
- 文件备份,文件隐藏,窗口隐藏,进程隐藏
- 注册表自启动和Windows定时任务自启动
- 自动防火墙注册
- 侦查调试软件和反调试
- 反病毒检查
- 反内存扫描
- Http/Https通讯支持
- 全消息自定义模糊和加密
- 支持UPnp端口映射
- 支持木马自我传播
- 支持自我版本更新

##### 功能性功能

- 支持进程扫描，活动窗口扫描，键盘记录，剪贴板记录的回传
- 支持定时截屏回传
- 支持常规系统信息获取和回传（IP,CPU,GPU,Wifi,操作系统信息,用户信息,网络配置,安装软件列表,进程列表等）
- 支持命令控制后台指定程序的执行
- 支持命令控制后台进行文件Http下载
- 支持用户Host文件修改，注册表修改
- 支持服务器安全狗的麻醉

##### 加强型功能

- 支持远程控制肉鸡机执行多种DDos攻击
- 支持从服务器远程上传BT种子，并于肉鸡机后台下载和执行
- 支持从服务器远程上传文件到肉鸡机并静默执行
- 支持内嵌WebServer并进行反向代理，允许外界访问
- 支持VB脚本，Bat脚本，MicrosoftPowerShell脚本静默执行
- 支持用户控制系统(UAC)绕过
- 支持控制肉鸡机打开指定网站，重启，桌面图片修改等测试功能

## Server

Server使用代码包括Server.go以及server目录内全部源码。

#### 功能列表

- 支持服务器用户管理
- 支持肉鸡客户端管理：注册，销毁等
- 支持用户信息于数据库的永久保存
- 支持命令注册机制
- 支持对肉鸡机的单体控制，组控制和全局控制

## TODOList

1. 需要上传功能,可以上传文件到网站目录。
2. 需要下载功能,可以下载对方网站目录的文件。
3. 需要修改和插入功能，可以修改或插入js,htm,html,php,asp,asa等常见网站文件
4. 需要数据库功能，可以获取到数据库的结构。修改，删除表结构等。