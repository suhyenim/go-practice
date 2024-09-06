package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive" // BSON 문서 형식을 지원하는 패키지
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "todo/docs" // Swagger 문서의 자동 생성

	"github.com/gin-gonic/gin"                 // Gin 프레임워크
	swaggerFiles "github.com/swaggo/files"     // Swagger 파일을 위한 패키지
	ginSwagger "github.com/swaggo/gin-swagger" // Swagger를 Gin과 통합하기 위한 패키지
	"go.mongodb.org/mongo-driver/bson"         // BSON 문서 작업을 위한 패키지
)

// Todo 구조체 정의
type Todo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`                    // MongoDB의 ObjectId로 매핑
	Content     string             `json:"content" bson:"content"`           // Todo의 내용
	IsCompleted bool               `json:"is_completed" bson:"is_completed"` // 완료 여부
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`     // 생성 시간
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`     // 수정 시간
}

type CreateTodoRequest struct {
	Content string `json:"content" bson:"content"`
}

type UpdateTodoRequest struct {
	Content     string `json:"content" bson:"content"`
	IsCompleted bool   `json:"is_completed" bson:"is_completed"`
}

// MongoDB 컬렉션 참조
var collection *mongo.Collection

// setupSwagger 함수는 Swagger 엔드포인트를 설정합니다.
func setupSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // Swagger UI 핸들러 등록
}

// getTodos 함수는 모든 Todo 항목을 반환합니다.
// @Summary Get all todos
// @Tags Todo
// @Produce json
// @Success 200 {array} Todo "List of all Todo items"               // 성공 시 모든 Todo 항목을 배열로 반환
// @Failure 500 {object} map[string]string "Internal server error"  // 서버 내부 오류 발생 시 500 오류 반환
// @Router /api/todos [get]
func getTodos(c *gin.Context) {
	cursor, err := collection.Find(context.TODO(), bson.D{}) // 모든 Todo 항목을 조회
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()}) // 오류 발생 시 500 상태 코드 반환
		return
	}
	defer cursor.Close(context.TODO()) // 함수 종료 시 cursor 닫기

	var todos []Todo
	if err = cursor.All(context.TODO(), &todos); err != nil { // 조회한 Todo 항목을 모두 읽어 리스트에 저장
		c.JSON(500, gin.H{"error": err.Error()}) // 오류 발생 시 500 상태 코드 반환
		return
	}

	c.JSON(200, todos) // 조회된 Todo 항목을 JSON 형식으로 반환
}

// getTodo 함수는 특정 Todo 항목을 ID로 조회하여 반환합니다.
// @Summary Get a specific todo
// @Tags Todo
// @Produce json
// @Param id path string true "Todo ID"                              // 경로 파라미터로 Todo 항목의 ID를 받음
// @Success 200 {object} Todo "Todo item retrieved successfully"     // 성공 시 조회된 Todo 항목 반환
// @Failure 400 {object} map[string]string "Invalid ID format"       // ID 형식 오류 시 400 오류 반환
// @Failure 404 {object} map[string]string "Todo not found"          // 해당 ID의 Todo 항목이 없을 경우 404 오류 반환
// @Failure 500 {object} map[string]string "Internal server error"   // 서버 내부 오류 발생 시 500 오류 반환
// @Router /api/todos/{id} [get]
func getTodo(c *gin.Context) {
	id := c.Param("id") // 요청에서 ID 추출
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"}) // ID 형식 오류 발생 시 400 상태 코드 반환
		return
	}

	var todo Todo
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&todo) // ID로 Todo 항목 조회 및 디코딩
	if err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"}) // Todo 항목이 없으면 404 상태 코드 반환
		return
	}

	c.JSON(200, todo) // 조회된 Todo 항목을 JSON 형식으로 반환
}

// createTodo 함수는 새 Todo 항목을 생성하여 데이터베이스에 저장합니다.
// @Summary Create a new todo
// @Tags Todo
// @Accept json
// @Produce json
// @Param todo body CreateTodoRequest true "Todo item to create" // 클라이언트로부터 입력받을 데이터의 구조체를 정의
// @Success 201 {object} Todo "Created Todo item"                // 성공적으로 생성된 Todo 항목을 반환
// @Failure 400 {object} map[string]string "Invalid request body" // 요청 본문이 유효하지 않을 경우 400 오류 반환
// @Failure 500 {object} map[string]string "Internal server error" // 서버 내부 오류 발생 시 500 오류 반환
// @Router /api/todos [post]
func createTodo(c *gin.Context) {
	var request CreateTodoRequest
	if err := c.BindJSON(&request); err != nil { // 클라이언트 요청 본문을 CreateTodoRequest 구조체로 바인딩
		c.JSON(400, gin.H{"error": err.Error()}) // 바인딩 오류 발생 시 400 상태 코드 반환
		return
	}

	// 새로운 Todo 객체 생성
	todo := Todo{
		ID:          primitive.NewObjectID(), // 새로운 ObjectID 생성
		Content:     request.Content,         // 클라이언트로부터 받은 content 설정
		IsCompleted: false,                   // 기본값으로 설정
		CreatedAt:   time.Now(),              // 현재 시간으로 생성 시간 설정
		UpdatedAt:   time.Now(),              // 현재 시간으로 수정 시간 설정
	}

	_, err := collection.InsertOne(context.TODO(), todo) // 새 Todo 항목을 데이터베이스에 삽입
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()}) // 삽입 오류 발생 시 500 상태 코드 반환
		return
	}

	c.JSON(201, todo) // 성공적으로 생성된 Todo 항목을 JSON 형식으로 반환
}

// updateTodo 함수는 특정 Todo 항목을 ID로 조회하여 업데이트합니다.
// @Summary Update an existing todo
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"                                   // 경로 파라미터로 Todo 항목의 ID를 받음
// @Param todo body UpdateTodoRequest true "Todo item to update"          // 요청 본문으로 업데이트할 Todo 항목을 받음
// @Success 200 {object} Todo "Updated Todo item"                         // 성공 시 업데이트된 Todo 항목 반환
// @Failure 400 {object} map[string]string "Invalid ID format or request" // ID 형식 또는 요청 본문 오류 시 400 오류 반환
// @Failure 500 {object} map[string]string "Internal server error"        // 서버 내부 오류 발생 시 500 오류 반환
// @Router /api/todos/{id} [put]
func updateTodo(c *gin.Context) {
	id := c.Param("id") // 요청에서 ID 추출
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"}) // ID 형식 오류 발생 시 400 상태 코드 반환
		return
	}

	var request UpdateTodoRequest
	if err := c.BindJSON(&request); err != nil { // 요청 본문에서 JSON을 UpdateTodoRequest 구조체로 바인딩
		c.JSON(400, gin.H{"error": err.Error()}) // 바인딩 오류 발생 시 400 상태 코드 반환
		return
	}

	update := bson.M{
		"$set": bson.M{
			"content":      request.Content,
			"is_completed": request.IsCompleted,
			"updated_at":   time.Now(),
		},
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update) // Todo 항목 업데이트
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()}) // 업데이트 오류 발생 시 500 상태 코드 반환
		return
	}

	c.JSON(200, update) // 성공적으로 업데이트된 정보 반환
}

// deleteTodo 함수는 특정 Todo 항목을 ID로 삭제합니다.
// @Summary Delete a todo
// @Tags Todo
// @Produce json
// @Param id path string true "Todo ID"                                  // 경로 파라미터로 Todo 항목의 ID를 받음
// @Success 200 {object} map[string]string "Todo item deleted successfully" // 성공 시 삭제된 Todo 항목의 ID 반환
// @Failure 400 {object} map[string]string "Invalid ID format"            // ID 형식 오류 시 400 오류 반환
// @Failure 500 {object} map[string]string "Internal server error"        // 서버 내부 오류 발생 시 500 오류 반환
// @Router /api/todos/{id} [delete]
func deleteTodo(c *gin.Context) {
	id := c.Param("id") // 요청에서 ID 추출
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"}) // ID 형식 오류 발생 시 400 상태 코드 반환
		return
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID}) // ID로 Todo 항목 삭제
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()}) // 삭제 오류 발생 시 500 상태 코드 반환
		return
	}

	c.JSON(200, gin.H{"deleted_id": id}) // 성공적으로 삭제된 Todo 항목의 ID를 JSON 형식으로 반환
}

func main() {
	router := gin.Default() // Gin 엔진 초기화

	// 접속할 MongoDB 주소 설정
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// MongoDB 연결
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 연결 확인
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// "todo-db" 데이터베이스 내 todo라는 컬렉션과 연결
	collection = client.Database("todo-db").Collection("todo")

	setupSwagger(router) // Swagger 설정

	// REST API 엔드포인트 설정
	router.GET("/api/todos", getTodos)
	router.GET("/api/todos/:id", getTodo)
	router.POST("/api/todos", createTodo)
	router.PUT("/api/todos/:id", updateTodo)
	router.DELETE("/api/todos/:id", deleteTodo)

	router.Run() // 기본적으로 localhost:8080에서 서버 실행
}
