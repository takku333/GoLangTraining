@echo off
cd /d %~dp0

go build gopl.io/ch1/fetch
fetch.exe http://gopl.io/ch1/helloworld?go-get=1
