package handler

import (
	"github.com/ti-click/golang-tidb-realworld-example-app/article"
	"github.com/ti-click/golang-tidb-realworld-example-app/user"
)

type Handler struct {
	userStore    user.Store
	articleStore article.Store
}

func NewHandler(us user.Store, as article.Store) *Handler {
	return &Handler{
		userStore:    us,
		articleStore: as,
	}
}
