# go user login
Go user api with Gin , MySql and Docker

## 1. Run with DockerCompose

1. **Run**

```shell script
docker-compose up -d
```
_______

## 2. Generate Docs

```shell script
# Get swag
go get -u github.com/swaggo/swag/cmd/swag

# Generate docs
swag init --dir cmd/api --parseDependency --output docs
```

Run and go to **http://localhost:8000/docs/index.html**
