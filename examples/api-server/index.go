package main

import (
	"context"
	"fmt"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IndexP struct {
	Collection     string  `json:"collection"`
	Host  string  `json:"host"`
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/index")
	{
		v1.POST("/", createTodo)
	}
	router.Run(":1234")
}


// createTodo add a new todo
func createTodo(c *gin.Context) {
	var newIndexP IndexP
	if err := c.BindJSON(&newIndexP); err != nil {
		return
	}

	fmt.Println(newIndexP.Host,newIndexP.Collection)
	IndexByCollection(newIndexP.Collection,newIndexP.Host)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Index successfully!", "CollectionName": newIndexP.Collection})
}

func IndexByCollection(collection ,host string)  {
	// Milvus instance proxy address, may verify in your env/settings
	milvusAddr := host+`:19530`
	fmt.Println(milvusAddr)

	// setup context for client creation, use 2 seconds here
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	c, err := client.NewGrpcClient(ctx, milvusAddr)
	if err != nil {
		// handling error and exit, to make example simple here
		log.Fatal("failed to connect to milvus:", err.Error())
	}
	// in a main func, remember to close the client
	defer c.Close()

	// here is the collection name we use in this example

	// load collection with async=false
	err = c.LoadCollection(ctx, collection, false)
	if err != nil {
		log.Fatal("failed to load collection:", err.Error())
	}
	log.Println("load collection completed")

	idx, err := entity.NewIndexNANG(entity.L2, 200,220,12,25,200,40,50,0.6,0.6,10)
	if err != nil {
		log.Fatal("fail to create ivf flat index:", err.Error())
	}
	err = c.CreateIndex(ctx, collection, "embedding", idx, false)
	if err != nil {
		log.Fatal("fail to create index:", err.Error())
	}
}



