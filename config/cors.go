package config

var allowedOrigins = []string{
	"https://localhost:3000",
	"https://indrariksa.github.io",
	"http://localhost:5173",
	"http://127.0.0.1:8088/",
	"https://fe-programming-iii.vercel.app/",
}

func GetAllowedOrigin() []string {
	return allowedOrigins
}
