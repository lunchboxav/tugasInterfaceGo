Ini adalah contoh penggunaan interface dalam package di Go.

Dalam contoh ini, ada 2 struct, yakni `Sepeda` dan `Motor`. Keduanya adalah bagian dari package `model`.
Selain itu, ada sebuah interface yang berisi beberapa method yang akan diimplementasikan di setiap struct ini. Perhatikan bahwa ada perbedaan minor di bagaimana sebuah interface diimplementasikan di kedua struct.

Dengan cara ini, kita memisahkan interface dan implementasinya, sehingga kita bisa merasakan fleksibilitas saat melakukan implementasi. Interface tetap dibutuhkan, hanya untuk mendefinisikan kontrak untuk implementasi sebuah method.

Di dalam `main.go`, kita kemudian cukup memanggil interface yang diimplementasikan di sebuah struct. Kita tidak lagi memanggil method langsung milik sebuah struct, tapi implementasinya. Dengan demikian, ini menghindarkan kita dari memanggil method yang belum diimplementasi di dalam struct, yang bisa menjadi indikator tentang celah di dalam kontrak yang kita buat di interface.
