@echo off
cd /d %~pd0

go run handleGithub.go test
go run handleGithub.go -action read test