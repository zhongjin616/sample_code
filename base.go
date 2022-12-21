package code

//go:generate codegen -type=int
//go:generate codegen -type=int -doc -output ./error_code_generated.md

// 通用：数据库类错误
const (
	// ErrDatabase - 500: Database error.
	ErrDatabase int = iota + 100100
	ErrNoRow
)

// 通用：认证授权类错误
const (
	ErrBadRequest   int = iota + 200101
	ErrUserNotFound int = iota + 200102
)

// init register error codes defines in this source code to `github.com/zhongjin616/errors`
func init() {
	register(ErrDatabase, 500, "Database error")
	register(ErrNoRow, 500, "now rows found")

	register(ErrBadRequest, 400, "Invalid request")
	register(ErrUserNotFound, 400, "User not found")
}
