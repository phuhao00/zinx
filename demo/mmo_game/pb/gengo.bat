@echo off
echo 正在生成Go代码...
 protoc.exe -I=./ --go_out=.  .\ID\(\w).proto
pause