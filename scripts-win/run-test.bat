@echo off
echo Starting TEST environment...
docker compose ^
 -f ../docker-compose.yml ^
 -f ../docker-compose.test.yml ^
 up -d --build
pause