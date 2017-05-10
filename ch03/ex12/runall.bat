@echo off
cd /d %~dp0

go run checkAnagram.go abcd cdba
go run checkAnagram.go 125.1 25.11
go run checkAnagram.go abcd cdba ad 