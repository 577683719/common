package dict_util

import (
	"ecs_cloud/common/model/system"
	"gorm.io/gorm"
)

// 获取字典数据
func GetDictData(tx *gorm.DB, dict_type, dict_value string) system.SysDictData {
	var dictData system.SysDictData
	tx.Raw("select * from sys_dict_data where status = '0' and dict_type = ? and dict_value = ? ", dict_type, dict_value).Scan(&dictData)
	return dictData
}

// 获取字典标签
func GetDictLabel(tx *gorm.DB, dict_type, dict_value string) string {
	var dictData system.SysDictData
	tx.Raw("select dict_label from sys_dict_data where status = '0'  and dict_type = ? and dict_value = ? ", dict_type, dict_value).Scan(&dictData)
	return dictData.DictLabel
}
