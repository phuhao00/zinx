@echo off
echo ��������Go����...
 protoc -I=./ --go_out=./pb/  .\subscriberID.proto
pause