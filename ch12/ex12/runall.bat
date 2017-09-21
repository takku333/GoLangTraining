@echo off
cd /d %~dp0

go build search/main.go
start main.exe
go build gopl.io/ch1/fetch
fetch.exe "http://localhost:12345/search?maddr=example@gmail.com&ccn=4980015105862555"
fetch.exe "http://localhost:12345/search?maddr=example"
fetch.exe "http://localhost:12345/search?ccn=897fa"
