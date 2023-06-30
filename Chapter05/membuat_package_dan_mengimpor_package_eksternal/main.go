// File: main2.go
package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	router.Run(":8080")
}

//untuk menjalankan file ini, anda harus menjalankan beberapa perintah, berikut perintahnya: go mod init 'namamodul' tekan enter lalu 'go mod tidy' tekan enter lagi kemudian 'go run namafile.go'.
// setelah file berhasil di run untuk mengcek apakah program sudah berjalan atau tidaknya anda harus membuka browser lalu masukkan didalam url 'localhost:8080' lalu tekan enter, nanti akan muncul isi dari message yaitu "Hello, World" jika sudah muncul maka program berhasil dijalankan.