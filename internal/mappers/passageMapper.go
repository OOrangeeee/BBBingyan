package mappers

import (
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
)

type PassageMapper struct {
}

func (pm *PassageMapper) AddNewPassage(passage *dataModels.Passage) error {
	result := utils.DB.Create(passage)
	return result.Error
}

func (pm *PassageMapper) DeletePassage(passage *dataModels.Passage) error {
	result := utils.DB.Delete(passage)
	return result.Error
}

func (pm *PassageMapper) UpdatePassage(passage *dataModels.Passage) error {
	result := utils.DB.Save(passage)
	return result.Error
}

func (pm *PassageMapper) GetAllPassages() ([]*dataModels.Passage, error) {
	var passages []*dataModels.Passage
	result := utils.DB.Find(&passages)
	return passages, result.Error
}

func (pm *PassageMapper) GetPassagesByID(passageID uint) ([]*dataModels.Passage, error) {
	var passages []*dataModels.Passage
	result := utils.DB.Find(&passages, "ID=?", passageID)
	return passages, result.Error
}

func (pm *PassageMapper) GetPassagesByPassageTitle(passageTitle string) ([]*dataModels.Passage, error) {
	var passages []*dataModels.Passage
	result := utils.DB.Find(&passages, "passage_title=?", passageTitle)
	return passages, result.Error
}

func (pm *PassageMapper) GetPassagesByPassageAuthorUserName(passageAuthorUserName string) ([]*dataModels.Passage, error) {
	var passages []*dataModels.Passage
	result := utils.DB.Find(&passages, "passage_author_user_name=?", passageAuthorUserName)
	return passages, result.Error
}

func (pm *PassageMapper) GetPassagesByPassageAuthorNickName(passageAuthorNickName string) ([]*dataModels.Passage, error) {
	var passages []*dataModels.Passage
	result := utils.DB.Find(&passages, "passage_author_nick_name=?", passageAuthorNickName)
	return passages, result.Error
}

func (pm *PassageMapper) GetPassagesByPassageAuthorId(passageAuthorId uint) ([]*dataModels.Passage, error) {
	var passages []*dataModels.Passage
	result := utils.DB.Find(&passages, "passage_author_id=?", passageAuthorId)
	return passages, result.Error
}

func (pm *PassageMapper) GetPassagesByPassageTag(passageTag string) ([]*dataModels.Passage, error) {
	var passages []*dataModels.Passage
	result := utils.DB.Find(&passages, "passage_tag=?", passageTag)
	return passages, result.Error
}

func (pm *PassageMapper) SearchPassagesByPassageTitle(passageTitle string) ([]*dataModels.Passage, error) {
	var passages []*dataModels.Passage
	// 根据点赞数降序排序
	result := utils.DB.Order("passage_be_liked_count desc").Find(&passages, "passage_title like ?", "%"+passageTitle+"%")
	return passages, result.Error
}

func (pm *PassageMapper) IfPassageExist(passageID uint) bool {
	var passage dataModels.Passage
	result := utils.DB.First(&passage, "ID=?", passageID)
	return result.Error == nil
}
