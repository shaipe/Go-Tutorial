# TTS 

## 说明

定时任务执行器



## 文件说明

tts 可执行程序,运行操作系统为Linux
tts.cnf 程序运行的配置文件

## 运行方式

```bash
chmod 777 tts

# $dir 程序存放目录

$dir/tts -c=tts.cnf

```

## 检验

运行好以后可以看启动第一次执行的日志,然后是每天设置的时间会执行

例如:

2019/04/01 13:44:13 tts.go:194: tts service start! start log file 
2019/04/01 13:44:14 tts.go:151: do worker name: [overall], request result: {"response":"received","status":"ok"}
