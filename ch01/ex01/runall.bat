@echo off
cd /d %~dp0

cd echo1
go run main.go hellow world
cd ..

cd echo2
go run main.go hellow world
cd ..

cd echo3
go run main.go hellow world
cd ..
