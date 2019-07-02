@echo off
echo 正在生成Go代码...
 protoc -I=./ --go_out=./pb/  .\subscriberID.proto
pause