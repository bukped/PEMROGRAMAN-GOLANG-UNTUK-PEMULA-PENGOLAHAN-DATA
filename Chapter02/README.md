# README.md

## Persyaratan Sistem untuk Menginstal Golang

Sebelum menginstal Golang di Windows, pastikan komputer Anda memenuhi persyaratan sistem berikut:

1. Sistem operasi Windows (Versi 7 atau yang lebih baru).
2. Setidaknya memiliki 2GB ruang penyimpanan yang tersedia.
3. Koneksi internet yang stabil (untuk mengunduh paket instalasi).

## Mengunduh dan Menginstal Golang di Windows

Berikut adalah langkah-langkah untuk mengunduh dan menginstal Golang di sistem operasi Windows:

1. Buka browser web Anda dan kunjungi situs resmi Go di [https://go.dev/dl/](https://go.dev/dl/).
2. Di halaman tersebut, Anda akan melihat beberapa versi Go yang tersedia. Pilih versi terbaru yang sesuai dengan sistem operasi Anda, misalnya Windows.
3. Setelah memilih versi, unduh file instalasi Go untuk Windows (berupa file berekstensi .msi).
4. Setelah selesai mengunduh, buka file .msi tersebut dan ikuti langkah-langkah instalasi yang ditampilkan di layar.
5. Pada langkah instalasi, Anda akan ditanyakan tentang direktori tempat Go akan diinstal. Secara default, Go akan diinstal di `C:\Go\`. Jika Anda ingin mengubah lokasi instalasi, Anda dapat memilih direktori lain.
6. Klik "Finish" untuk menyelesaikan proses instalasi.

## Konfigurasi PATH dan GOPATH

Setelah menginstal Go, Anda perlu mengatur beberapa variabel lingkungan untuk menggunakan Go dengan benar, yaitu mengatur PATH dan GOPATH. Berikut adalah langkah-langkah konfigurasinya:

1. Ketik "Environment Variables" pada fitur pencarian di taskbar dan buka program tersebut.
2. Ketika program sudah terbuka, pilih "Environment Variables".
3. Di bagian "User Variables", cari variabel bernama "Path" dan klik "Edit".
4. Tambahkan jalur direktori Go (`C:\Go\bin` sesuai dengan direktori instalasi) ke variabel "Path".
5. Klik "OK" untuk menyimpan perubahan.
6. Selanjutnya, Anda perlu mengatur variabel lingkungan GOPATH. Klik "New" pada bagian "System Variables" dan buat variabel baru dengan nama "GOPATH". Isi nilainya dengan jalur direktori di mana Anda ingin menyimpan proyek Go Anda, misalnya `C:\GoProjects\` (sesuaikan direktori dengan saat instalasi Golang).
7. Klik "OK" untuk menyimpan perubahan.

## Verifikasi Instalasi Golang

Setelah melakukan instalasi, Anda dapat memverifikasi apakah Go telah terinstal dengan benar atau belum. Berikut adalah langkah untuk memverifikasi instalasi Golang:

1. Buka Command Prompt Windows.
2. Ketik perintah `go version` dan tekan Enter.
3. Jika Anda melihat versi Go yang terinstal, misalnya `go version go1.17 windows/amd64`, itu berarti instalasi Go Anda berhasil.

## Menggunakan Go Module

Go Module adalah fitur yang memungkinkan manajemen dependensi yang lebih baik dalam proyek Go. Untuk menggunakan Go Module, Anda perlu menginisialisasi modul Go pada proyek Anda dengan perintah `go mod init nama_modul`. Setelah modul diinisialisasi, Go Module akan mengatur dependensi proyek Anda dengan mempertahankan versi yang konsisten dan memastikan proyek dapat dibangun dengan benar. Anda juga dapat menjalankan perintah `go mod tidy` pada terminal VSCode untuk membersihkan dan memperbarui dependensi proyek.

---
