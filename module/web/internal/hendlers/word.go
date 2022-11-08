package hendlers

import (
	"github.com/dimayaschhu/vocabulary/module/web/internal/repo"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type WordHandler struct {
	wordRepo *repo.WordRepo
}

func NewWordHandler(wordRepo *repo.WordRepo) *WordHandler {
	return &WordHandler{wordRepo: wordRepo}
}
func (h *WordHandler) GetRoute() map[string]func(c *gin.Context) {
	return make(map[string]func(c *gin.Context))
}

func (h *WordHandler) PostRoute() map[string]func(c *gin.Context) {
	r := make(map[string]func(c *gin.Context))
	r["/word/create"] = h.Create()
	return r
}

type CreateWordRequest struct {
	Name      string
	Lesson    int
	Translate string
}

func (h *WordHandler) Create() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req CreateWordRequest
		if err := ctx.BindJSON(&req); err != nil {
			sendError(ctx, err)
			return
		}

		//if err := h.validateCreateReq(&req); err != nil {
		//	h.sendError(ctx, err)
		//return
		//}

		if h.wordRepo.ExistStudent(req.Name) {
			sendError(ctx, errors.New("exist student.Name: "+req.Name))
			return
		}

		if err := h.wordRepo.CreateWord(req.Name, req.Translate, req.Lesson); err != nil {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"name": req.Name,
		})
	}
}

//func (h *StudentHandler) validateCreateReq(req *CreateStudentRequest) error {
//	return validation.ValidateStruct(req,
//		validation.Field(
//			&req.Name,
//			validation.Required,
//		),
//		validation.Field(
//			&req.Lesson,
//			validation.Required,
//		),
//		validation.Field(
//			&req.ChatId,
//			validation.Required,
//		),
//	)
//}
