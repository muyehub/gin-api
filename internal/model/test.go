package model

import (
	"github.com/jinzhu/gorm"
	"github.com/muyehub/gin-api/pkg/app"
)

type Test struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageURL string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

// TableName 表名方法
func (t Test) TableName() string {
	return "test"
}

type TestSwagger struct {
	List  []*Test
	Pager *app.Pager
}

func (t Test) Create(db *gorm.DB) (*Test, error) {
	if err := db.Create(&t).Error; err != nil {
		return nil, err
	}

	return &t, nil
}

func (t Test) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (t Test) Get(db *gorm.DB) (Test, error) {
	var Test Test
	db = db.Where("id = ? AND state = ? AND is_del = ?", t.ID, t.State, 0)
	err := db.First(&Test).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Test, err
	}

	return Test, nil
}

func (t Test) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error; err != nil {
		return err
	}

	return nil
}

type TestRow struct {
	TestID        uint32
	TagID         uint32
	TagName       string
	TestTitle     string
	TestDesc      string
	CoverImageUrl string
	Content       string
}
