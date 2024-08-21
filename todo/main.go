package main

import (
	"context"
	"log"
	"net/http"
	"time"

	_ "backend/docs" // docs 패키지 경로
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Todo 구조체 정의
type Todo struct {
	TodoID      int       `json:"todo_id" bson:"todo_id"`
	Content     string    `json:"content" bson:"content"`
	IsCompleted bool      `json:"is_completed" bson:"is_completed"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

// setupSwagger configures the Swagger endpoint
func setupSwagger(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// MongoDB 연결 함수
func ConnectMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// MongoDB에 연결 확인
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// MongoDB 컬렉션 참조
var collection *mongo.Collection

// @Summary Get all todos
// @Description Retrieve a list of all todos
// @Tags todos
// @Produce json
// @Success 200 {array} Todo
// @Router /api/todos [get]
func getTodos(c *gin.Context) {
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	var todos []Todo
	if err = cursor.All(context.TODO(), &todos); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, todos)
}

// @Summary Get a specific todo
// @Description Retrieve details of a specific todo by ID
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} Todo
// @Router /api/todos/{id} [get]
func getTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	err := collection.FindOne(context.TODO(), bson.M{"todo_id": id}).Decode(&todo)
	if err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(200, todo)
}

// @Summary Create a new todo
// @Description Create a new todo item
// @Tags todos
// @Accept json
// @Produce json
// @Success 201 {object} Todo
// @Router /api/todos [post]
func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	_, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, todo)
}

// @Summary Update an existing todo
// @Description Update details of an existing todo by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} Todo
// @Router /api/todos/{id} [put]
func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todo.UpdatedAt = time.Now()
	filter := bson.M{"todo_id": id}
	update := bson.M{
		"$set": todo,
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, todo)
}

// @Summary Delete a todo
// @Description Delete a todo item by ID
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} map[string]string
// @Router /api/todos/{id} [delete]
func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"todo_id": id})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"deleted_id": id})
}

func main() {
	r := gin.Default()

	// MongoDB 연결
	client, err := ConnectMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	collection = client.Database("db").Collection("todo")

	setupSwagger(r)

	r.GET("/api/todos", getTodos)
	r.GET("/api/todos/:id", getTodo)
	r.POST("/api/todos", createTodo)
	r.PUT("/api/todos/:id", updateTodo)
	r.DELETE("/api/todos/:id", deleteTodo)

	r.Run() // 기본적으로 localhost:8080 실행
}
