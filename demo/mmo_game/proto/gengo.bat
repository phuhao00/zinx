@echo off
echo  later
 protoc -I=./ --go_out=./ID/  .\messageid.proto
pause