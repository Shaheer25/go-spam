package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

const (
	hostName       string = "localhost:27017"
	dbName         string = "demo_todo"
	collectionName string = "todo"
)

type (
	todoModel struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		Title     string        `bson:"title"`
		Completed bool          `bson:"completed"`
		CreatedAt time.Time     `bson:"createAt"`
	}

	todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func init() {
	sess, err := mgo.Dial(hostName)
	checkErr(err)
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(dbName)
}

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tpl", nil)
}

func createTodo(c *gin.Context) {
	var t todo

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusProcessing, err)
		return
	}

	// simple validation
	if t.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The title field is required",
		})
		return
	}

	// if input is okay, create a todo
	tm := todoModel{
		ID:        bson.NewObjectId(),
		Title:     t.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	if err := db.C(collectionName).Insert(&tm); err != nil {
		c.JSON(http.StatusProcessing, gin.H{
			"message": "Failed to save todo",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo created successfully",
		"todo_id": tm.ID.Hex(),
	})
}

func updateTodo(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))

	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The id is invalid",
		})
		return
	}

	var t todo

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusProcessing, err)
		return
	}

	// simple validation
	if t.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The title field is required",
		})
		return
	}

	// if input is okay, update a todo
	if err := db.C(collectionName).
		Update(
			bson.M{"_id": bson.ObjectIdHex(id)},
			bson.M{"title": t.Title, "completed": t.Completed},
		); err != nil {
		c.JSON(http.StatusProcessing, gin.H{
			"message": "Failed to update todo",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo updated successfully",
	})
}

func fetchTodos(c *gin.Context) {
	todos := []todoModel{}

	if err := db.C(collectionName).
		Find(bson.M{}).
		All(&todos); err != nil {
		c.JSON(http.StatusProcessing, gin.H{
			"message": "Failed to fetch todo",
			"error":   err,
		})
		return
	}

	todoList := []todo{}
	for _, t := range todos {
		todoList = append(todoList, todo{
			ID:        t.ID.Hex(),
			Title:     t.Title,
			Completed: t.Completed,
			CreatedAt: t.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todoList,
	})
}

func deleteTodo(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))

	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The id is invalid",
		})
		return
	}

	if err := db.C(collectionName).RemoveId(bson.ObjectIdHex(id)); err != nil {
		c.JSON(http.StatusProcessing, gin.H{
			"message": "Failed to delete todo",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted successfully",
	})
}

func main() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	r := gin.Default()
	r.LoadHTMLGlob("static/*.tpl")

	r.GET("/", homeHandler)

	todoGroup := r.Group("/todo")
	{
		todoGroup.GET("/", fetchTodos)
		todoGroup.POST("/", createTodo)
		todoGroup.PUT("/:id", updateTodo)
		todoGroup.DELETE("/:id", deleteTodo)
	}

	srv := &http.Server{
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Listening on port")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown error:", err)
	}
	log.Println("Server gracefully stopped!")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err) // respond with error page or message
	}
}
