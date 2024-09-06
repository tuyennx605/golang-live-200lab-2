package main

import (
	"log"
	"net/http"
	"os"
	"todo-list/component/tokenprovider/jwt"
	"todo-list/middleware"
	"todo-list/module/item/transport/ginitem"
	"todo-list/module/upload"
	"todo-list/module/user/storage"
	usergin "todo-list/module/user/transport/gin"
	ginuserlikeitem "todo-list/module/userlikeitem/transport/gin"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// function

func main() {
	// load .env file
	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatal("error loading environment variables")
	}

	// connect db
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err) // log and exit program
	}

	// turn on debug sql
	db = db.Debug()

	if err := runService(db); err != nil {
		log.Fatal(err) // log and exit program
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	r.Use(middleware.Recover())

	// đặt folder là static public
	r.Static("/static", "./static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD

	SECRET_KEY := os.Getenv("SECRET_KEY")

	authStore := storage.NewSQLStore(db)
	tokenProvider := jwt.NewTokenJWTProvider("jwt", SECRET_KEY)
	middlewareAuth := middleware.RequiredAuth(authStore, tokenProvider)

	v1 := r.Group("/v1")

	{
		v1.PUT("upload", upload.Upload(db))

		v1.POST("register", usergin.Register(db))
		v1.POST("login", usergin.Login(db, tokenProvider))
		v1.GET("profile", middlewareAuth, usergin.Profile())

		items := v1.Group("/items")

		{
			// POST /v1/items (create anew item)
			items.POST("", middlewareAuth, ginitem.CreateItem(db))
			// GET /v1/items (list items) /v1/items?page=1...
			items.GET("", middlewareAuth, ginitem.ListItem(db))
			// GET /v1/items/:id (get item detail by id)
			items.GET("/:id", ginitem.GetItem(db))
			// (PUT || PATCH) /v1/items/:id (update by id)
			items.PATCH("/:id", middlewareAuth, ginitem.UpdateItem(db))
			// DELETE /v1/items/:id (delete by id)
			items.DELETE("/:id", middlewareAuth, ginitem.DeleteItem(db))

			//// like
			items.POST("/:id/like", middlewareAuth, ginuserlikeitem.LikeItem(db))
			items.DELETE("/:id/unlike", middlewareAuth, ginuserlikeitem.UnLikeItem(db))
			items.GET("/:id/liked-users", middlewareAuth, ginuserlikeitem.ListUserLikeItem(db))
		}

		rpc := v1.Group("/rpc")
		{
			rpc.POST("/get_item_likes", ginuserlikeitem.GetItemLikes(db))
		}
	}

	return r.Run(":3005") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//

//

//
//
// type TodoItem struct {
// 	Id          int        `json:"id"` // tag json cho biet Id quy doi sang json thanh id và ngược lại
// 	Title       string     `json:"title"`
// 	Description string     `json:"description"`
// 	Status      string     `json:"status"`
// 	CreatedAt   *time.Time `json:"created_at"`
// 	UpdatedAt   *time.Time `json:"updated_at"`
// }

// func main() {
// 	now := time.Now().UTC()

// 	item := TodoItem{
// 		Id:          1,
// 		Title:       "taks 1",
// 		Description: "content 1",
// 		Status:      "Doing",
// 		CreatedAt:   &now,
// 		UpdatedAt:   &now,
// 	}

// 	// hàm chuyển từ struct sang json
// 	jsonData, err := json.Marshal(item)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	log.Println(string(jsonData))

// 	jsonString := `{"id":1,"title":"taks 1","description":"content 1","status":"Doing","created_at":"2024-08-08T14:40:56.464805349Z","updated_at":"2024-08-08T14:40:56.464805349Z"}`

// 	var item2 TodoItem
// 	// hàm chuyển từ json sang struct
// 	if err := json.Unmarshal([]byte(jsonString), &item2); err != nil {
// 		log.Fatalln(err)
// 	}
// 	log.Println(item2)
// }
