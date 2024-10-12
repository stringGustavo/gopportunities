package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utilizando as configs padrões do gin
	router := gin.Default()
	// Definindo uma Rota
	router.GET("/ping", func(c *gin.Context) { //Função Handler
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Rodando a API
	router.Run(":8080")
}
