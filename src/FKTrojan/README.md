# 大小码设计文档
## 需求
为解决控制远端windows机器而设计

## 原理
通过在受控端执行任意exe，达到控制受控端目的

## 功能
服务端监听tcp端口，受控端访问此端口并获取执行命令，命令包括执行二进制文件，上传文件和下载文件

> ### 命令
命令分为：执行，上传，下载

> ### 执行
受控端将命令二进制以压缩加密保存于本地，获取到命令后，
受控端判断文件是否存在于本地，如果不存在，则先请求下载，
解密解压，执行命令，执行结果回传，删除执行文件

> ### 上传，下载
受控端获取到命令后，上传下载对应文件

## 模块
简要介绍代码的核心模块

> ### [main](main)
主模块，server.go是服务端，client.go是客户端，
分别使用go build编译成二进制可执行文件。

> ### [connect](connect)
connect模块为tcp加密模块，控制端口与传输端口均使用加密传输
加密算法为aes-cfb，并使用动态协商iv，由于密码为提前约定，所以代码省掉验证，
借鉴shadowsocks最新协议中“加密及验证”的先进方法，受控端与服务端所有数据都是经过加密

> ### [hide_client](hide_client)
此模块目的是隐藏受控端，受控端以windows普通服务运行，注册服务时将文件名，desc，displayname设置为
与windows现有服务中极为相似的名字，以此迷惑

> ### [service_command](service_command)
命令，此模块功能为命令发送与接收，涉及受控端与服务端通信

> ### [service_transfer](service_transfer)
传输，文件上传下载模块，此模块用于传输大字节流，借鉴ftp控制命令与数据传输端口分离

> ### [client_tools](client_tools)
此模块下，是具体的执行命令，包括run_bat,mysql_tools,scan_dir等，在受控端最终执行的exe

> ### [client_tools/user_interface](client_tools/user_interface)
此模块为服务端控制UI，此模块是c#实现，将命令下发，文件传输功能可视化控制，运行于服务端

> ### [file_crypto](file_crypto)
加解密执行文件

> ### [stream_utils](stream_utils)
服务端与受控端流式加解密

> ### [registry_crypto](registry_crypto)
加解密受控端使用到的注册表项

## 服务端部署
> ### 安装配置mysql
普通的mysql安装即可

> ### database初始化
>>下载[database初始化文件](../../bin/database.sql)

>>命令行下运行：

      cd {mysql的安装目录}
      mysql.exe -u{用户名} -p{密码} < {脚本路径}\\database.sql

> ### 服务端运行
> >下载[服务端二进制文件](../../bin/server.exe)

> >下载[服务端配置文件](../../bin/config.json)

> >配置文件解释：

    {
      "mysql_user": "root",                   // mysql用户名
      "mysql_pass": "qazwsx",                 // mysql密码
      "mysql_host": "tcp(127.0.0.1:3306)",    // mysql host
      "mysql_name": "panel",                  // dbname
      "base_data_dir":"d:/bin/data",          // 日志存放目录
      "cmd_port":7778,                        // 命令监听端口
      "trans_port":7779                       // 传输监听端口
    }

>>二进制与配置文件放入同一个目录，命令行下执行`server.exe`


