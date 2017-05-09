@echo off
cd /d %~dp0

go run surface_all.go eggBox >eggBox.xml
go run surface_all.go mogle >mogle.xml
go run surface_all.go saddle >saddle.xml
