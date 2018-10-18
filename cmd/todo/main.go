package main

import (
	"github.com/gobuffalo/envy"

	"github.com/artistomin/planner/infrastructure/mongo"
)

func main() {
	envy.Load()

	var (
		dburl      = envy.Get("MONGODB_URL", "127.0.0.1")
		dbname     = envy.Get("MONGODB_NAME", "todo")
		dbusername = envy.Get("MONGO_USERNAME", "root")
		dbpassword = envy.Get("MONGO_PASSWORD", "root")
	)

	session := mongo.New(dburl, dbusername, dbpassword)
	defer session.Close()

	mongo.NewTodoRepository(dbname, "todo", session)
}
