# Last.fm CLI 🎧

This is a simple project I made using the Last.fm API.  
I wanted to study how to consume APIs without having to deal with all the front-end, so I decided to create a CLI.

---

## 🚀 Getting Started (Docker)

The fastest way to run this project is through Docker. This ensures you don't need to worry about local Go environments or dependencies.

### 1. Build the image
From the project root, run:
```bash
docker build -t lastfm-cli .
```
### 2. Run the application

You will need your own Last.fm API Key. The application will interactively ask for the username once it starts.

```bash

docker run -it -e LASTFM_API_KEY="YOUR_API_KEY_HERE" lastfm-cli topTracks

```

### 🔧 Prerequisites
- Docker  
- A Last.fm API Key (https://www.last.fm/api)

## 🛠 Technologies Used

- Go 1.21  
- [Cobra](https://github.com/spf13/cobra) – CLI framework  
- Docker – For containerized execution  
- Last.fm API – Fetching music and artist data

## 📂 Project Structure

- `cmd/` – CLI commands (`topTracks`, `topArtists`)  
- `model/` – Data models and structs for the API  
- `service/` – API consumption and data handling logic  
- `main.go` – Entry point of the application  
- `Dockerfile` – Docker configuration  
- `.dockerignore` / `.gitignore` – Ignored files  
- `go.mod` / `go.sum` – Go dependencies
