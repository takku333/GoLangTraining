@echo off
cd /d %~dp0

go build mychat2.go
go build gopl.io/ch8/netcat3
start mychat2.exe
start netcat3.exe
start netcat3.exe
start netcat3.exe