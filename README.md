HTTP Server dengan Middleware
middleware adalah program yang berfungsi sebagai perantara antara permintaan (request) dan respons (response). Middleware memungkinkan untuk menambahkan fungsionalitas tertentu ke dalam aplikasi web tanpa harus mengubah kode inti (core code) dari aplikasi tersebut.
Ini adalah contoh sederhana untuk HTTP server dengan middleware menggunakan bahasa pemrograman Go.

START
|
v
Membuat HTTP ServeMux baru
|
v
Membuat fungsi handler HTTP untuk endpoint "/greet"
|
v
Membuat dua fungsi middleware
|
v
Menggabungkan dua fungsi middleware dengan fungsi handler HTTP
|
v
Menunggu permintaan HTTP yang masuk pada port 3000
|
v
Ketika sebuah permintaan HTTP diterima, jalankan melalui rangkaian middleware
|
v
Fungsi middleware pertama menulis "Middleware pertama" ke konsol
|
v
Fungsi middleware kedua menulis "Middleware kedua" ke konsol
|
v
Fungsi handler HTTP menulis "Hello World" sebagai respons
|
v
HTTP server mengirimkan respons kembali ke klien
|
END
