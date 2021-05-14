package errcode

// a collection for handling tag error
var (
	ErrorGetTagListFail = NewError(20010001, "取得標籤列表失敗")
	ErrorCreateTagFail  = NewError(20010002, "建立標籤失敗")
	ErrorUpdateTagFail  = NewError(20010003, "更新標籤失敗")
	ErrorDeleteTagFail  = NewError(20010004, "刪除標籤失敗")
	ErrorCountTagFail   = NewError(20010005, "統計標籤失敗")
)

// a collection for handling article error
var (
	ErrorGetArticleListFail = NewError(30010001, "取得文章列表失敗")
	ErrorCreateArticleFail  = NewError(30010002, "建立文章失敗")
	ErrorUpdateArticleFail  = NewError(30010003, "更新文章失敗")
	ErrorDeleteArticleFail  = NewError(30010004, "刪除文章失敗")
	ErrorCountArticleFail   = NewError(30010005, "統計文章失敗")
)

// Error for upload file
var (
	ErrorUploadFileUrl = NewError(40010001, "upload file fail")
)
