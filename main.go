package logger

func main() {
	config := Config{
		Development: true,
		FileDir:     "./log/",
	}
	Init(&config)

	log.Sugar().Debug("aafefef", 2323, "abddee")
	log.Sugar().Error("2323", "aabd", "abc")
	//Info("test info")
	//Error("test error")
	//Info("abcddddddd", zap.String("abc", "aaaaaaaaa"))
	//Error("234533error", zap.Int("abc", 333))
}
