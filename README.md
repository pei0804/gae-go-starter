# gae-go-starter

## Required Install 
- Python 2.7.x
- [goapp](https://cloud.google.com/appengine/docs/go/download)

## How to use

### Install

```bash
// Required package download
make dev

// MySQL start
make docker/build
make docker/start

// Create Database
make migrate/init 
make migrate/up 
```

### local run

```bash
make local
```

- [here](http://localhost:8080)

### Deploy

#### first time

```
gcloud init
gcloud auth login
gcloud config set project "your-project-id"
gcloud app deploy
```
