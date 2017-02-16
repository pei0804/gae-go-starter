# gae-go-starter

## Required Install 
- Python 2.7.x
- [goapp](https://cloud.google.com/appengine/docs/go/download)

## How to use

### Install

```bash
make dev
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
