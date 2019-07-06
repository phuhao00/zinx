package api
import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	_"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_"github.com/garyburd/redigo/redis"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
	_"github.com/json-iterator/go"
)
type DbWorker struct {
	//mysql data source name
	Dsn string
}

func TestLogin( t *testing.T) {
	dbw := DbWorker{
		Dsn: "user:password@tcp(127.0.0.1:3306)/gogs",
	}
	db, err := sql.Open("mysql",
		dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()
}
func TestMongoDB(t *testing.T)  {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}