@echo off
cd /d %~dp0

go run myfetch.go https://www.google.co.jp/ https://www.yahoo.co.jp/ http://golang.jp/
