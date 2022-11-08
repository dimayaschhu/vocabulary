package web

import (
	"github.com/dimayaschhu/vocabulary/module/web/internal/hendlers"
	"github.com/dimayaschhu/vocabulary/module/web/internal/repo"
	"github.com/dimayaschhu/vocabulary/pkg/httpserver"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		repo.NewWordRepo,
		repo.NewStudentRepo,
		hendlers.NewWordHandler,
		hendlers.NewStudentHandler,
	),
	fx.Invoke(func(router *httpserver.Router, wordHandler *hendlers.WordHandler, studentHandler *hendlers.StudentHandler) {
		for path, method := range wordHandler.PostRoute() {
			router.AddPostHandler(path, method)
		}
		for path, method := range studentHandler.PostRoute() {
			router.AddPostHandler(path, method)
		}
	}),
)
