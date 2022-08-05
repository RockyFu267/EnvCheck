#!/bin/bash
# 获取CPU架构 判断用x86还是arm 待添加
# 目前默认环境均为统一的架构
# 赋予权限
chmod 777 /root/envcheck/envcheck
chmod 777 /root/envcheck/config.yaml
chmod 777 /root/envcheck/client-start.sh
chmod 777 /root/envcheck/gossh

cd /root/envcheck/
nohup /root/envcheck/envcheck &