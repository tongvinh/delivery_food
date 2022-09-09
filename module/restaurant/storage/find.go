package restaurantstorage

import (
	"context"
	"gorm.io/gorm"
	"myapp/common"
	restaurantmodel "myapp/module/restaurant/model"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*restaurantmodel.Restaurant, error) {

	var data restaurantmodel.Restaurant

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
