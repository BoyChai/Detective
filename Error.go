package blast_rar

type ErrorCode int

const Pass ErrorCode = 1

type Error struct {
	Is   bool
	Msg  string
	Code ErrorCode
	Err  error
}
