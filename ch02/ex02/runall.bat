@echo off
cd /d %~dp0

go build cf.go conv.go
cf.exe 30 40 50
echo 30| cf.exe