@echo off
cd /d %~dp0

go build myreverb2.go
start myreverb2.exe
mynetcat.exe
