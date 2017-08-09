@echo off
cd /d %~dp0

go build myjpeg.go
myjpeg.exe -f=gif < gopher.png > pngTogif.gif
myjpeg.exe -f=png < gopher.jpeg > jpegTopng.png
myjpeg.exe -f=jpeg < gopher.gif > gifTojpeg.jpeg