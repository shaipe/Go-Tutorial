#!/bin/sh

# create directory
mkdir -p /usr/bin/tts

# copy file to tatget
cp tts /usr/bin/tts/
cp tts.service /usr/lib/systemd/system  # 复制文件
systemctl enable tts.service            # 激活服务
systemctl start tts                     # 启动服务
systemctl status tts.service            # 查看服务状态