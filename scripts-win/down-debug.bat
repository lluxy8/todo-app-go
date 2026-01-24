@echo off
set "VOLUME_FLAG="

choice /c yn /m "Remove volume?"
if errorlevel 2 goto RUN
set "VOLUME_FLAG=-v"

:RUN
docker compose ^
 -f ../docker-compose.yml ^
 -f ../docker-compose.debug.yml ^
 down %VOLUME_FLAG%

pause
