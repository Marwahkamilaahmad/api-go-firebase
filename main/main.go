package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"go-firebase/handlers"
	"go-firebase/firebase"

)

func main() {
	client, err := firebase.CreateClient()
	if err != nil {
		log.Printf("error creating client: %v", err)
		return
	}

	route := gin.Default()
	route.POST("/api/addTodo", func(c *gin.Context) { handlers.AddTodoHandler(c, client) })
	route.GET("/api/getTodo/:id", func(c *gin.Context) { handlers.GetTodoHandler(c, client) })

	log.Fatal(route.Run(":8080"))
}



// package main

// import (
	// "context"
	// "log"
	// "github.com/gin-gonic/gin"
	// "go-firebase/handlers"

	// firebase "firebase.google.com/go"
	// firestore "cloud.google.com/go/firestore"
// )

// func main(){

// 	client, err := createClient()
// 	if err != nil {
// 		log.Printf("error creating client", err)
// 		return
// 	}

// 	route := gin.Default()
// 	route.GET("api/addTodo", handlers.eventAddTodo)


// }

// func createClient() (*firestore.Client, error) {
// 	ctx := context.Background()
// 	conf := &firebase.Config{ProjectID: "uas-c251d"}
// 	app, err := firebase.NewApp( ctx, conf)

// 	if err != nil {
// 		log.Printf("error initialize app:", err)
// 		return nil, err
// 	}

// 	client, err := app.Firestore(ctx)
// 	if err != nil {
// 		log.Print("error creating client")
// 		log.Fatalln(err)

// 	}

	// ref := client.Collection("todos").NewDoc()
	// result , err := ref.Set(ctx, map[string]interface{}{
	// 	"title" : "random todo",
	// 	"description" : "random task",
	// })

	// if err != nil {
	// 	log.Printf("error occured when creating a todo")
	// }

	// log.Printf("Result is ", result)

// 	return client, nil

// }