package main

import "fmt"
import "github.com/gin-gonic/gin"


func main(){
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.JSON(200,gin.H{
			"message": "starting of build of Url midget",
		})
	})

	err := r.Run(":9808")

	if err!=nil{
		panic(fmt.Sprintf("fucking failed to heat up the server %v", err))
	}
	fmt.Printf("lets build URL shortner")
}