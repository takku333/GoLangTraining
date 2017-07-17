@echo off
cd /d %~dp0

go build myfindlinks2.go links.go
myfindlinks2.exe -depth=3 http://gopl.io/ 