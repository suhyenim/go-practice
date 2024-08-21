# go
go mod init todo
go run .
go mod tidy

# gin
go get -u github.com/gin-gonic/gin

# swagger
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go install github.com/swaggo/swag/cmd/swag@latest
swag init
http://localhost:8080/swagger/index.html

# mongodb
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson

# docker compose
docker-compose up --build
docker-compose ps
docker-compose logs
docker-compose down <- 서비스 중지
docker-compose down -v <- 데이터 볼륨, 네트워크까지 삭제
docker-compose down --rmi all <- 모든 이미지까지 삭제
http://localhost:8080



### reference
gin crud: https://k-in.tistory.com/11
go & mongoDB: https://soyoung-new-challenge.tistory.com/107
go project layout: https://github.com/golang-standards/project-layout

