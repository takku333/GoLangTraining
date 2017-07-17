@echo off
cd /d %~dp0

go build gopl.io/ch8/reverb1
start reverb1.exe
go build mynetcat.go
mynetcat.exe