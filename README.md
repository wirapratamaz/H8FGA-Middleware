# HTTP Server dengan Middleware

Middleware adalah program yang berfungsi sebagai perantara antara permintaan (request) dan respons (response). Middleware memungkinkan untuk menambahkan fungsionalitas tertentu ke dalam aplikasi web tanpa harus mengubah kode inti (core code) dari aplikasi tersebut.

Ini adalah contoh sederhana untuk HTTP server dengan middleware menggunakan bahasa pemrograman Go.

## Alur Flow

1. Membuat HTTP ServeMux baru
2. Membuat fungsi handler HTTP untuk endpoint "/greet"
3. Membuat dua fungsi middleware
4. Menggabungkan dua fungsi middleware dengan fungsi handler HTTP
5. Menunggu permintaan HTTP yang masuk pada port 3000
6. Ketika sebuah permintaan HTTP diterima, jalankan melalui rangkaian middleware
7. Fungsi middleware pertama menulis "Middleware pertama" ke konsol
8. Fungsi middleware kedua menulis "Middleware kedua" ke konsol
9. Fungsi handler HTTP menulis "Hello World" sebagai respons
10. HTTP server mengirimkan respons kembali ke klien
