
##  Pengertian tentang Fungsi dalam Golang

Fungsi adalah blok kode yang memiliki nama tertentu dan dapat digunakan untuk melakukan tugas tertentu. Fungsi membantu mengorganisir kode menjadi bagian-bagian yang terpisah dan dapat dipanggil dari tempat lain dalam program.

### Mendefinisikan Fungsi dengan Menentukan Nama, Parameter, dan Tipe Kembalian

- Nama fungsi adalah identitas yang digunakan untuk memanggil fungsi.
- Parameter adalah variabel yang digunakan untuk menerima nilai saat memanggil fungsi.
- Tipe kembalian adalah tipe data yang menggambarkan jenis nilai yang dikembalikan oleh fungsi.

Contoh:

```go
func tambah(a, b int) int {
    return a + b
}
```

### Contoh Pembuatan Fungsi Sederhana

Berikut adalah contoh fungsi sederhana yang menghitung luas persegi:

```go
func hitungLuasPersegi(sisi int) int {
    return sisi * sisi
}
```

### Menggunakan Fungsi di dalam Program Utama

Fungsi-fungsi yang telah didefinisikan dapat digunakan di dalam program utama untuk melakukan tugas tertentu. Berikut adalah contoh penggunaan fungsi dalam program utama:

```go
func main() {
    hasil := tambah(3, 4)
    fmt.Println("Hasil penjumlahan:", hasil)
}
```

### Pemanggilan Fungsi dengan Memberikan Argumen ke Parameter

Saat memanggil fungsi, argumen dapat diberikan ke parameter sesuai dengan tipe data yang diharapkan. Berikut adalah contoh pemanggilan fungsi dengan memberikan argumen ke parameter:

```go
func main() {
    luas := hitungLuasPersegi(5)
    fmt.Println("Luas persegi:", luas)
}
```

##  Membuat Package dan Mengimpor Package Eksternal

### Konsep Package dalam Golang dan Manfaatnya

Package adalah cara untuk mengorganisir dan membagi kode ke dalam unit yang lebih kecil. Package membantu dalam memperjelas struktur kode, memungkinkan penggunaan kode yang telah ada, dan memfasilitasi kolaborasi dalam pengembangan perangkat lunak.

### Mengimpor Package Eksternal yang Telah Tersedia

Golang memungkinkan penggunaan package eksternal yang telah tersedia di repositori resmi Go atau di sumber pihak ketiga. Package-package ini dapat digunakan untuk memperluas fungsionalitas aplikasi.

Contoh:

```go
// File: main2.go
package main

import (
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
```

Catatan: Pastikan Anda telah menjalankan beberapa perintah seperti `go mod init 'namamodul'`, `go mod tidy`, dan `go run namafile.go` 
untuk menjalankan program ini. Anda juga perlu membuka browser dan mengakses `localhost:8080` untuk melihat hasil dari program.
