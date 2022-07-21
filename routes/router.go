package routes

import (
	"fmt"

	"github.com/edrank/edrank_backend/db"
	"github.com/gin-gonic/gin"
)

type tt struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func InitRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		database := db.GetDatabase()

		rows, err := database.Query("select * from test")
		if err != nil {
			fmt.Println(err)
		}
		a, e := rows.Columns()

		fmt.Println(a, e)

		var t []tt
		for rows.Next() {
			var id int
			var name string
			// var alb tt
			if err := rows.Scan(&id, &name); err != nil {
				c.JSON(200, gin.H{
					"data": t,
				})
			}
			fmt.Println(id, name)
			t = append(t, tt{Id: id, Name: name})
		}
		c.JSON(200, gin.H{
			"data": t,
		})
	})
}

func InitPublicRoutes(r *gin.RouterGroup) {

}

func InitPrivateRoutes(r *gin.RouterGroup) {

}
