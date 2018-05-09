package ginsession

import (
	"context"
	"sync"

	"github.com/gin-gonic/gin"
	"gopkg.in/session.v2"
)

var (
	// DefaultKey Keys stored in the context
	DefaultKey      = "github.com/go-session/gin-session"
	once            sync.Once
	internalManager *session.Manager
	internalError   = func(ctx *gin.Context, err error) {
		ctx.AbortWithError(500, err)
	}
)

// SetErrorHandler Set error handling
func SetErrorHandler(handler func(ctx *gin.Context, err error)) {
	internalError = handler
}

func manager(opt ...session.Option) *session.Manager {
	once.Do(func() {
		internalManager = session.NewManager(opt...)
	})
	return internalManager
}

// New Create a session middleware
func New(opt ...session.Option) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		store, err := manager(opt...).Start(context.Background(), ctx.Writer, ctx.Request)
		if err != nil {
			internalError(ctx, err)
			return
		}
		ctx.Set(DefaultKey, store)
		ctx.Next()
	}
}

// FromContext Get session storage from context
func FromContext(ctx *gin.Context) session.Store {
	return ctx.MustGet(DefaultKey).(session.Store)
}

// Destroy a session
func Destroy(ctx *gin.Context) error {
	return manager().Destroy(context.Background(), ctx.Writer, ctx.Request)
}

// Refresh a session and return to session storage
func Refresh(ctx *gin.Context) (session.Store, error) {
	return manager().Refresh(context.Background(), ctx.Writer, ctx.Request)
}
