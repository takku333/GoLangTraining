@echo off
cd /d %~dp0

go build myreverb3.go
start myreverb3.exe
mynetcat.exe
