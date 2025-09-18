package ginitem

import (
	"my-app/common"
	"my-app/modules/item/biz"
	"my-app/modules/item/model"
	"my-app/modules/item/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TodoItemCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewCreateItemBiz(store)

		if err := business.CreateNewItem(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
