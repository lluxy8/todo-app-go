@echo off
echo Starting DEV environment...
docker compose ^
 -f ../docker-compose.yml ^
 -f ../docker-compose.dev.yml ^
 up -d
pause