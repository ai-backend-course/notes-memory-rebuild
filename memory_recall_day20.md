# Day 20 - Docker compose Networking
- docker-compose.yml defines multi-container setup.
- Services share a private network.
- API connects to Postgres at host "db".
- docker compose up --build starts all containers.
- docker compose down stops them.
- Volumes persist Postgres data between runs. 
- Both local Postgres and Docker Postgres have the same DB name (ai_backend) but live in different environments.

