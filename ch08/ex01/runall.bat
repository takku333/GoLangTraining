@echo off
cd /d %~dp0

start myclock.exe -port 8010
start myclock.exe -port 8020 
start myclock.exe -port 8030
clockwall.exe NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030