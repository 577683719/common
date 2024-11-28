package tree_util

import "ecs_cloud/idl/gen/v1/device"

/**
 * 递归模式，数组转tree
 * @param arr 目标数组
 * @param pid 第一级 目标id
 * @returns {*[]} tree
 * @constructor
 */
type TreeListStr struct {
	Name     string        `json:"name"`
	Id       string        `json:"id"`
	Pid      string        `json:"pid"`
	Data     interface{}   `json:"data"`
	Children []TreeListStr `json:"children"`
}
type TreeList struct {
	Name     string     `json:"name"`
	Id       int        `json:"id"`
	Pid      int        `json:"pid"`
	Children []TreeList `json:"children"`
}

func DiskArrayToTreeStr(arr []*device.QueryDiskHistoryBasedOn, pid string) []*device.QueryDiskHistoryBasedOn {
	var newArr []*device.QueryDiskHistoryBasedOn
	for _, v := range arr {
		if v.SnapshotID == v.Pid {
			panic("数据错误")
		}
		if v.Pid == pid {
			v.Children = DiskArrayToTreeStr(arr, v.Id)
			newArr = append(newArr, v)
		}
	}
	return newArr
}
func ArrayToTreeStr(arr []TreeListStr, pid string) []TreeListStr {
	var newArr []TreeListStr
	for _, v := range arr {
		if v.Pid == pid {
			v.Children = ArrayToTreeStr(arr, v.Id)
			newArr = append(newArr, v)
		}
	}
	return newArr
}
func ArrayToTree(arr []TreeList, pid int) []TreeList {
	var newArr []TreeList
	for _, v := range arr {
		if v.Pid == pid {
			v.Children = ArrayToTree(arr, v.Id)
			newArr = append(newArr, v)
		}
	}
	return newArr
}
