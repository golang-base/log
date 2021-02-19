package main

func main() {
	config := Config{
		Development: false,
		OutputDir:   "./log/",
	}
	Init(&config)

	log.Info("test info")
	log.Error("test error")
}
