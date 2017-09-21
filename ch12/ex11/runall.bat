@echo off
cd /d %~dp0

go build search/main.go
start main.exe
go build gopl.io/ch1/fetch
fetch.exe "http://localhost:12345/search"
fetch.exe "http://localhost:12345/search?l=golang&l=programming"
fetch.exe "http://localhost:12345/search?l=golang&l=programming&max=100"
fetch.exe "http://localhost:12345/search?x=true&l=golang&l=programming"
