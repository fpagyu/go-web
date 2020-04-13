package dao

// 此文件定义缓存key的格式

const (
	Book_List_Key = "books:v:%d"   // books:v:<版本号>
	Book_Key      = "book:%v:v:%d" // book:<book_id>:v:<版本号>
)
