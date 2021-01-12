package dao

import (
	"github.com/muyehub/gin-api/internal/model"
)

type Test struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d *Dao) GetTest(id uint32, state uint8) (model.Test, error) {
	test := model.Test{Model: &model.Model{ID: id}, State: state}
	return test.Get(d.engine)
}
