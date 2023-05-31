package blog

import (
	"gorm.io/gorm/clause"
	"inception/api/internal/global"
	"inception/api/internal/model"
)

//todo 待优化

// GetInsertTags
//
//	@Description: 有新标签则存入。返回所有更新完成后的标签数据
//	@param tagsStr
//	@return tags
func GetInsertTags(tagsStr []string) []*model.Tag {
	global.Log.Info("tagsStr", tagsStr)
	tags := toTags(tagsStr)
	global.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{}),
	}).Create(&tags)
	global.DB.Where("name IN ?", tagsStr).Find(&tags)
	return tags
}

func toTags(tagStr []string) []*model.Tag {
	tags := make([]*model.Tag, 0)
	for _, v := range tagStr {
		tags = append(tags, &model.Tag{
			Name: v,
		})
	}
	return tags
}
