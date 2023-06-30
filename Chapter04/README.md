
## Deklarasi Variabel
Dalam Golang, variabel dideklarasikan menggunakan keyword `var` diikuti dengan nama variabel dan tipe datanya. Contoh deklarasi variabel:

```go
var nama string
var umur int
```

## Operator dan Ekspresi
Golang mendukung berbagai operator aritmatika, relasional, logika, dan bitwise untuk memanipulasi nilai dalam ekspresi. Contoh penggunaan operator dan ekspresi:

```go
a := 10
b := 5
c := a + b
fmt.Println("Hasil:", c) // Output: Hasil: 15
```

## Pengendalian Aliran Program
Golang menyediakan struktur pengendalian aliran program seperti kondisional (`if`), perulangan (`for`), dan pernyataan `switch` untuk mengatur alur eksekusi program. Contoh penggunaan pengendalian aliran program:

```go
umur := 20

if umur >= 18 {
	fmt.Println("Anda sudah dewasa")
} else {
	fmt.Println("Anda masih di bawah umur")
}

for i := 1; i <= 5; i++ {
	fmt.Println(i)
}

switch umur {
case 18:
	fmt.Println("Anda baru berusia 18 tahun")
case 20:
	fmt.Println("Anda berusia 20 tahun")
default:
	fmt.Println("Umur Anda tidak diketahui")
}
```

## Tipe Data Dasar
Golang memiliki tipe data dasar seperti numerik (integers, floats), boolean, string, karakter (rune), array, slice, map, dan struct. Berikut adalah penjelasan singkat tentang tipe data dasar tersebut:

- Numerik: Golang menyediakan tipe data untuk bilangan bulat (int) dan bilangan desimal (float32, float64).
- Boolean: Digunakan untuk menyimpan nilai kebenaran (true/false).
- String: Digunakan untuk menyimpan teks.
- Karakter (Rune): Digunakan untuk menyimpan karakter Unicode.
- Array: Digunakan untuk menyimpan kumpulan elemen dengan tipe data yang sama dan ukuran tetap.
- Slice: Digunakan untuk menyimpan kumpulan elemen dengan tipe data yang sama dan ukuran yang bisa berubah.
- Map: Digunakan untuk menyimpan pasangan nilai kunci (key-value pairs).
- Struct: Digunakan untuk menggabungkan beberapa nilai dengan tipe data yang berbeda menjadi satu kesatuan.
