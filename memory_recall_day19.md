# Day 19 - Dockerfile Setup
- Dockerfile defines how to build and run the app inside a container.
- Two-stage build: builder (compile) + runtime (slim image).
- EXPOSE 8080 tells Docker which port to open.
- CMD runs the compiled binary.
- docker build -t notes-api . builds the image.
- docker run -p 8080:8080 --env-file .env notes-api runs it.
- Next: add docker-compose to connect Postgrest + API. 

# What to memorize now
- Docker builds containers from Dockerfiles.
- Containers are isolated environments for running apps.
- "docker build" creates an image.
- "docker run" starts the container.
- "-p" exposes ports, "--env-file" loads env vars.
- The Dockerfile just says: base image -> copy code -> build -> run.

# What not to memorize yet
- Exact Dockerfile syntax or advanced flags.
- Compose YAML keys and configs 