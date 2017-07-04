@echo off
cd /d %~dp0

go build gopl.io/ch1/fetch
go build main.go
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./main div div h2