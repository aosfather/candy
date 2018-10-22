package search

import "encoding/json"

type Field struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PageSearchResult struct {
	Id    string `json:"uuid"` //查询的请求id
	Index int64  `json:"page"` //页码
	Data  []TargetObject
}

type TargetObject struct {
	Id   string          `json:"id"`
	Data json.RawMessage `json:"data"`
}
type SourceObject struct {
	TargetObject
	Fields map[string][]string `json:"fields"`
}

type FieldType byte

const (
	FT_TEXT   FieldType = 11 //文本类型
	FT_NUMBER FieldType = 9  //数字
	FT_ENUM   FieldType = 8  //枚举
	FT_ID     FieldType = 7  //id唯一标识
	FT_DATE   FieldType = 6  //日期
)

//索引的元数据信息
type IndexMeta struct {
	Name   string      //索引名称
	Fields []FieldMeta //字段
}

type FieldMeta struct {
	Name string    //字段名称
	Type FieldType //类型
}
