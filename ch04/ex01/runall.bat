@echo off
cd /d %~dp0

go run popCountLoop.go xorCountSHA256.go abc abc
go run popCountLoop.go xorCountSHA256.go abc adc
go run popCountLoop.go xorCountSHA256.go 1 9

