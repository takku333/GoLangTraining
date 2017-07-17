@echo off
cd /d %~dp0

set GOMAXPROCS=1
go test -bench=.
set GOMAXPROCS=2
go test -bench=.
set GOMAXPROCS=3
go test -bench=.
set GOMAXPROCS=4
go test -bench=.