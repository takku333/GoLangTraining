@echo off
cd /d %~dp0

go build packageDependency.go
packageDependency.exe fmt...
