package easyRouter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RouterInterFace interface {
	GetRelativePath() string
	GetMethods() []string
	GetHandlers() []gin.HandlerFunc
	AddRouters(...RouterInterFace) RouterInterFace
	GetRouters() []RouterInterFace //
}

type Routers []RouterInterFace

type Router struct {
	RelativePath string
	Method       []string
	Handlers     []gin.HandlerFunc
	Routers      []RouterInterFace
}

func NewRouter(relativePath string, methods []string, handlerFunc ...gin.HandlerFunc) *Router {
	return &Router{
		RelativePath: relativePath,
		Method:       methods,
		Handlers:     handlerFunc,
		Routers:      []RouterInterFace{},
	}
}

func NewRouterGroup(relativePath string, handlerFunc ...gin.HandlerFunc) *Router {
	return &Router{
		RelativePath: relativePath,
		Method:       []string{},
		Handlers:     handlerFunc,
		Routers:      []RouterInterFace{},
	}
}

func (r *Router) GetRelativePath() string {
	return r.RelativePath
}

func (r *Router) GetMethods() []string {
	return r.Method
}

func (r *Router) GetHandlers() []gin.HandlerFunc {
	return r.Handlers
}

func (r *Router) AddRouters(iRouters ...RouterInterFace) RouterInterFace {
	for _, router := range iRouters {
		r.Routers = append(r.Routers, router)
	}
	return r
}

func (r *Router) GetRouters() []RouterInterFace {
	return r.Routers
}

func SetRoutersToGinRouter(engine gin.IRouter, iRouters ...RouterInterFace) error {
	for _, router := range iRouters {
		if len(router.GetRouters()) <= 0 {
			methods := router.GetMethods()
			if len(methods) >= 7 {
				engine.Any(router.GetRelativePath(), router.GetHandlers()...)
			} else {
				for _, method := range methods {
					switch strings.ToUpper(method) {
					case http.MethodGet:
						engine.GET(router.GetRelativePath(), router.GetHandlers()...)
					case http.MethodHead:
						engine.HEAD(router.GetRelativePath(), router.GetHandlers()...)
					case http.MethodPost:
						engine.POST(router.GetRelativePath(), router.GetHandlers()...)
					case http.MethodPut:
						engine.PUT(router.GetRelativePath(), router.GetHandlers()...)
					case http.MethodPatch:
						engine.PATCH(router.GetRelativePath(), router.GetHandlers()...)
					case http.MethodDelete:
						engine.DELETE(router.GetRelativePath(), router.GetHandlers()...)
					case http.MethodOptions:
						engine.OPTIONS(router.GetRelativePath(), router.GetHandlers()...)
					}
				}
			}
		} else {
			group := engine.Group(router.GetRelativePath(), router.GetHandlers()...)
			SetRoutersToGinRouter(group, router.GetRouters()...)
		}
	}
	return nil
}
