package listener

import (
	"log"
	"main/golang/app/graph"
	"main/golang/app/graph/generated"
	"main/golang/app/middlewares"
	"main/golang/app/service/client"
	"main/golang/app/service/visitor"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

type GinServer interface {
	LintenAndServe(address string)
}

type ginImpl struct {
	myAccount   client.MyAccountService
	visitor     visitor.Visitor
	middlewares middlewares.Minddleware
}

func NewGinServer() GinServer {
	return &ginImpl{
		myAccount:   client.NewserviceMyAccount(),
		visitor:     visitor.NewVisitor(),
		middlewares: middlewares.NewMinddleware(),
	}
}

func (g *ginImpl) LintenAndServe(address string) {
	// Load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error when loading .env file")
	}

	// address := os.Getenv("HOST")

	// check redis client
	redisCli := redisClient()

	// gin setup if production
	// gin.SetMode(gin.ReleaseMode)

	// router init
	var router *gin.Engine = gin.New()
	var h *handler.Server = handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	}))

	// router.Use(gin.Recovery(), middlewares.Middleware(redisCli))
	router.LoadHTMLGlob("./views/*.html")

	client := router.Group("/account/")
	client.Use(g.middlewares.MyAccountMiddleware(redisCli))
	{
		client.GET("/", g.myAccount.MyAccount())
		client.POST("/data/all/", g.myAccount.GetAllData(h))
		client.POST("/data/products/product/", g.myAccount.InsertProduct())
		client.PUT("/data/products/product/", g.myAccount.UpdateProduct())
		client.DELETE("/data/products/product/:product_id", g.myAccount.DeleteProduct())
	}

	visitor := router.Group("/")
	{
		visitor.GET("/", g.visitor.Home())
		visitor.GET("/products/product/:pname/:pid", g.visitor.GetProductById())
		visitor.POST("/auth/login", g.visitor.PostLogin())
		visitor.POST("/auth/register", g.visitor.PostRegister())

	}

	log.Print("listen and router at " + address)
	router.Run(address)
}

func redisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
