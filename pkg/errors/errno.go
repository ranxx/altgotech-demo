package errors

// ErrNumber ...
type ErrNumber int

// const 错误码
const (
	Err ErrNumber = (10000)

	// 参数解析失败
	ErrBadRequest ErrNumber = (10400)

	// 鉴权失败
	ErrAuthVerify ErrNumber = (10401)

	// 没有权限
	ErrPermissionDenied ErrNumber = (10403)

	// 不存在
	ErrNotFound ErrNumber = (10404)

	// 用户已存在
	ErrUserEmailExist ErrNumber = 11000
	ErrUserNotFound   ErrNumber = 11001
	ErrUserPasswd     ErrNumber = 11002

	// 板块
	ErrSpaceNameExist  ErrNumber = 12000
	ErrSpaceNotFound   ErrNumber = 12001
	ErrSpaceJoined     ErrNumber = 12002
	ErrSpaceNoJoinSelf ErrNumber = 12003
	ErrSpaceNoOwner    ErrNumber = 12004
	ErrSpaceNoTop      ErrNumber = 12005
	ErrSpaceNoAccess   ErrNumber = 12006

	// 文章
	ErrArticleNotFound ErrNumber = 13000
)
