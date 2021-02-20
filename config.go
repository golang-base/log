package logger

type Config struct {
	FileDir       string // 日志输出目录
	OutputConsole bool   // 是否console输出
	Development   bool   // 是否是开发模式
}
