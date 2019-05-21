# TTS 

## 说明

定时任务执行器



## 文件说明

tts 可执行程序,运行操作系统为Linux
tts.cnf 程序运行的配置文件

## 运行方式

### 服务方式运行

```sh
# 1. 将tts.service 复制到 /usr/lib/systemd/system 目录下
# 2. 激活服务运行
# 3. 查看状态

cp tts.service /usr/lib/systemd/system  # 复制文件
systemctl enable tts.service            # 激活服务
systemctl start tts                     # 启动服务
systemctl status tts.service            # 查看服务状态

```


```bash
chmod 777 tts

# $dir 程序存放目录

$dir/tts -c=tts.cnf

```


### 服务配置说明

```bash
[Unit]
Description=tts - 定时处理任务工具
After=network.target nss-lookup.target

[Service]
Type=forking
ExecStart=/bin/tts/tts -c=/bin/tts/tts.cnf -log=/bin/tts/tts.log -s=false
Restart=always
KillMode=process
RestartSec=1

[Install]
WantedBy=multi-user.target

```


## 检验

随系统启动了一个8081的web端口,可以在浏览器中打开看是否可以正常访问

运行好以后可以看启动第一次执行的日志,然后是每天设置的时间会执行

例如:

2019/04/01 13:44:13 tts.go:194: tts service start! start log file 
2019/04/01 13:44:14 tts.go:151: do worker name: [overall], request result: {"response":"received","status":"ok"}
