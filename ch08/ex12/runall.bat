@echo off
cd /d %~dp0

go build mychat.go
go build gopl.io/ch8/netcat3
start mychat.exe
start netcat3.exe
start netcat3.exe
start netcat3.exe