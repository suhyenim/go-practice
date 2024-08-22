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

# mongodb express
http://localhost:8081 (admin/pass)

# docker compose
docker-compose up --build  
docker-compose ps  
docker-compose logs  
docker-compose down
http://localhost:8080  




---
# reference
gin crud: https://k-in.tistory.com/11    
go project layout: https://github.com/golang-standards/project-layout   
go & mongoDB
- https://ifelsehero.medium.com/create-read-update-and-delete-using-golang-and-mongodb-5241c319895c
- https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
- https://ccambo.blogspot.com/2020/12/golang-golang-mongodb.html
- https://welcome1208.tistory.com/50
