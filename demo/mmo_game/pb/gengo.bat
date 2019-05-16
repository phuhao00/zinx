@echo off
echo 正在生成Go代码...
protoc --go_out=. ./SubscriberID/subscriberID.proto
pause