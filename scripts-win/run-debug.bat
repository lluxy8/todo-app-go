@echo off
echo Starting DEBUG environment...
docker compose ^
 -f ../docker-compose.yml ^
 -f ../docker-compose.debug.yml ^
 up -d
pause
