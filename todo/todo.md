# go
go mod init todo  
go run .  
go mod tidy  

# gin
go get -u github.com/gin-gonic/gin  
http://localhost:8080  

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

---
# reference
gin crud: https://k-in.tistory.com/11    
go project layout: https://github.com/golang-standards/project-layout   
go & mongoDB
- https://ifelsehero.medium.com/create-read-update-and-delete-using-golang-and-mongodb-5241c319895c
- https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
- https://ccambo.blogspot.com/2020/12/golang-golang-mongodb.html
- https://welcome1208.tistory.com/50

---
# api
아래는 각 API 엔드포인트의 상태 코드와 그 설명을 표로 정리한 것입니다.

| HTTP Method | API Endpoint         | 상태 코드 | 설명                                               |
|-------------|-----------------------|------------|----------------------------------------------------|
| `GET`       | `/api/todos`          | 200        | 모든 Todo 항목 조회 성공.                          |
|             |                       | 500        | 서버 오류 또는 데이터 조회 실패.                   |
| `GET`       | `/api/todos/:id`      | 200        | 특정 Todo 항목 조회 성공.                          |
|             |                       | 404        | 지정된 ID의 Todo 항목이 존재하지 않음.               |
|             |                       | 500        | 서버 오류 또는 데이터 조회 실패.                   |
| `POST`      | `/api/todos`          | 201        | 새 Todo 항목 생성 성공.                            |
|             |                       | 400        | 요청 본문 JSON 형식이 올바르지 않음.                |
|             |                       | 500        | 서버 오류 또는 데이터 삽입 실패.                   |
| `PUT`       | `/api/todos/:id`      | 200        | Todo 항목 업데이트 성공.                          |
|             |                       | 400        | 요청 본문 JSON 형식이 올바르지 않음.                |
|             |                       | 500        | 서버 오류 또는 데이터 업데이트 실패.               |
| `DELETE`    | `/api/todos/:id`      | 200        | Todo 항목 삭제 성공.                              |
|             |                       | 500        | 서버 오류 또는 데이터 삭제 실패.                   |

### 상태 코드 설명

- **200 OK**: 요청이 성공적으로 처리되었음을 나타냅니다. 응답 본문에 요청된 데이터가 포함됩니다.
- **201 Created**: 요청이 성공적으로 처리되었으며, 새 리소스가 생성되었음을 나타냅니다.
- **400 Bad Request**: 클라이언트의 요청이 잘못되었음을 나타냅니다. 주로 요청 데이터의 형식이 잘못되었을 때 발생합니다.
- **404 Not Found**: 요청한 리소스가 서버에서 찾을 수 없음을 나타냅니다. 주로 잘못된 ID나 URL로 요청했을 때 발생합니다.
- **500 Internal Server Error**: 서버에서 예기치 않은 오류가 발생했음을 나타냅니다. 주로 서버의 문제나 데이터베이스 오류가 원인입니다.