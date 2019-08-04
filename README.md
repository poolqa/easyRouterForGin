# 推薦使用gin-swagger
https://github.com/swaggo/gin-swagger

# easyRouterForGin

1. 簡化gin router配置，使層級清晰
2. 支持無限層級group
3. 規則

| ex           |    |
| ------------ |:--:|
| group->group | OK |
| group->path  | OK |
| path->group  | NG |

## example
### main.go
```go
import (
	"github.com//easyRouterForGin"
	"github.com/gin-gonic/gin"
)

func main() {
	srv := gin.Default()
	easyRouter.SetRoutersToGinRouter(srv,
		easyRouter.NewRouter("/", []string{"GET"}, helloWorld),
		easyRouter.NewRouterGroup("/v1").AddRouters(
			easyRouter.NewRouterGroup("/user").AddRouters(
				easyRouter.NewRouter("/", []string{"get"}, user),
				easyRouter.NewRouter("/:name", []string{"get", "post"}, user),
				),
			easyRouter.NewRouter("/ping", []string{"GET"}, pong),
			),
		easyRouter.NewRouterGroup("/v2").AddRouters(
			easyRouter.NewRouterGroup("/user").AddRouters(
				easyRouter.NewRouter("/", []string{"get"}, user),
				easyRouter.NewRouter("/:name", []string{"get", "post"}, user),
			),
			easyRouter.NewRouter("/ping", []string{"GET"}, pong),
		),
		easyRouter.NewRouter("/ping", []string{"GET"}, pong),
	)

	srv.Run(":8888")
}

func helloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func user(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user list:" + c.Param("name"),
	})
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
```

### routers print
```txt
[GIN-debug] GET    /                         --> main.helloWorld (3 handlers)
[GIN-debug] GET    /v1/user/                 --> main.user (3 handlers)
[GIN-debug] GET    /v1/user/:name            --> main.user (3 handlers)
[GIN-debug] POST   /v1/user/:name            --> main.user (3 handlers)
[GIN-debug] GET    /v1/ping                  --> main.pong (3 handlers)
[GIN-debug] GET    /v2/user/                 --> main.user (3 handlers)
[GIN-debug] GET    /v2/user/:name            --> main.user (3 handlers)
[GIN-debug] POST   /v2/user/:name            --> main.user (3 handlers)
[GIN-debug] GET    /v2/ping                  --> main.pong (3 handlers)
[GIN-debug] GET    /ping                     --> main.pong (3 handlers)
[GIN-debug] Listening and serving HTTP on :8888
```
