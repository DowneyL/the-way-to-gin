package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(page int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(page).Limit(pageSize).Find(&tags)

	return
}

func GetTagsTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)

	return tag.ID > 0
}

func AddTag(name string, state int, createdBy string) bool {
	wdb.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)

	return tag.ID > 0
}

func EditTag(id int, data interface{}) bool {
	wdb.Model(Tag{}).Where("id = ?", id).Updates(data)

	return true
}

func DeleteTag(id int) bool {
	wdb.Where("id = ?", id).Delete(&Tag{})

	return true
}

func CleanTag() bool {
	wdb.Unscoped().Where("deleted_on != ?", 0).Delete(&Tag{})

	return true
}

