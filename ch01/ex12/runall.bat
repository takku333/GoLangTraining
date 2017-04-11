@echo off
cd /d %~dp0

go run lissajousURL.go http://localhost:8000/?cycles=30