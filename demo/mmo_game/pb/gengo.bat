@echo off
echo ��������Go����...
 protoc.exe -I=./ --go_out=.  .\ID\(\w).proto
pause