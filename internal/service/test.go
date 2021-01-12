package service

import (
	"github.com/muyehub/gin-api/internal/model"
)

type TestRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) GetArticle(param *TestRequest) (*model.Test, error) {
	_, err := svc.dao.GetTest(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	return &model.Test{
		Model:         nil,
		Title:         "",
		Desc:          "",
		Content:       "",
		CoverImageURL: "",
		State:         0,
	}, nil
}
