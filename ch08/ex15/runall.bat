@echo off
cd /d %~dp0

go build mychat4.go
go build gopl.io/ch8/netcat3
start mychat4.exe
start netcat3.exe
start netcat3.exe
start netcat3.exe