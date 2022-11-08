package hendlers

import (
	"errors"
	"github.com/dimayaschhu/vocabulary/module/web/internal/repo"
	"github.com/gin-gonic/gin"
	//validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
)

type StudentHandler struct {
	studentRepo *repo.StudentRepo
}

type CreateStudentRequest struct {
	Name   string
	Lesson int
	ChatId int
}

func NewStudentHandler(studentRepo *repo.StudentRepo) *StudentHandler {
	return &StudentHandler{studentRepo: studentRepo}
}
func (h *StudentHandler) GetRoute() map[string]func(c *gin.Context) {
	return make(map[string]func(c *gin.Context))
}

func (h *StudentHandler) PostRoute() map[string]func(c *gin.Context) {
	r := make(map[string]func(c *gin.Context))
	r["/student/create"] = h.Create()
	return r
}

func (h *StudentHandler) Create() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req CreateStudentRequest
		if err := ctx.BindJSON(&req); err != nil {
			sendError(ctx, err)
			return
		}

		//if err := h.validateCreateReq(&req); err != nil {
		//	h.sendError(ctx, err)
		//return
		//}

		if h.studentRepo.ExistStudent(req.Name) {
			sendError(ctx, errors.New("exist student.Name: "+req.Name))
			return
		}

		if err := h.studentRepo.CreateStudent(req.Name, req.Lesson, req.ChatId); err != nil {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"name": req.Name,
		})
	}
}

func sendError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
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
