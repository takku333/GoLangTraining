@echo off
cd /d %~dp0

go run printSHA.go -SHAbit=256 abc
go run printSHA.go -SHAbit=384 abc
go run printSHA.go -SHAbit=512 abc
go run printSHA.go abc adc