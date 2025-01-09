package server

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"gitnub.com/premwut/todo-service/model"
)

func Start() {
	e := echo.New()

	db, err := ConnectDB()
	if err != nil {
		log.Fatalln("cannot connect database", err)
	} else {
		fmt.Println("DB connected:", db)
	}

	// * Auto Migrations
	err = db.AutoMigrate(&model.Project{})
	if err != nil {
		log.Fatalln("cannot migrate Project table", err)
	}
	err = db.AutoMigrate(&model.Task{})
	if err != nil {
		log.Fatalln("cannot migrate Task table", err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalln("cannot migrate User table", err)
	}

	c := NewUserController()
	e.GET("/hello", c.HelloWorld)
	e.GET("/user/:id", c.GetUser)

	// projects
	pc := NewProjectController(db)
	e.GET("/projects/:projectId", pc.getProject)
	e.POST("/projects", pc.createProject)
	e.POST("/projects/:projectId/tasks", pc.createTask)
	e.Start(":8000")

}
