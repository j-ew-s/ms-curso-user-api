# ms-curso-user-api


# User Mod
go mod init github.com/j-ew-s/ms-curso-user-api

# User API porta
:5100


# Comunica com outros serviços:

1 - AUTH GRPC   pela porta :5500 


# Pacotes que utilizamos 
go get -u github.com/go-sql-driver/mysql
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite



# DOCKER 

docker pull mysql/mysql-server:tag
docker run --name user-mysql -e MYSQL_ROOT_PASSWORD=12345 -d mysql/mysql-server:latest


docker build --tag user-api:v0.0.2 -f Dockerfile.multistage   .


docker compose build

docker compose up