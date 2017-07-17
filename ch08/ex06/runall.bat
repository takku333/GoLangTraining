@echo off
cd /d %~dp0

go build myfindlinks.go
myfindlinks.exe -depth=3 http://gopl.io/ 