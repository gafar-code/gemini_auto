# Clipboard File Generator

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/clipfile)](https://goreportcard.com/report/github.com/yourusername/clipfile)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Clipboard File Generator adalah utilitas Go yang memantau clipboard sistem Anda dan secara otomatis membuat atau memperbarui file berdasarkan konten yang disalin. Ini memungkinkan pembuatan file yang cepat dan mudah hanya dengan menyalin teks dengan format tertentu.

## Fitur Utama

- ✅ **Pemantauan Clipboard Real-time**: Memeriksa clipboard setiap 2 detik untuk konten baru
- ✅ **Format yang Fleksibel**: Mendukung dua format berbeda untuk menandai path file
- ✅ **Pembuatan Direktori Otomatis**: Membuat struktur direktori yang diperlukan jika belum ada
- ✅ **Aman**: Mencegah serangan path traversal dan validasi input
- ✅ **Pelaporan Status**: Menampilkan informasi tentang file yang dibuat atau diperbarui
- ✅ **Ringan**: Konsumsi sumber daya yang minimal

## Instalasi

### Persyaratan

- Go 1.18 atau lebih tinggi
- [golang.design/x/clipboard](https://pkg.go.dev/golang.design/x/clipboard) package

### Cara Instalasi

1. Clone repositori:
   ```bash
   git clone https://github.com/yourusername/clipfile.git
   cd clipfile
   ```

2. Build aplikasi:
   ```bash
   go build -o clipfile
   ```

3. (Opsional) Pindahkan binary ke lokasi PATH Anda:
   ```bash
   sudo mv clipfile /usr/local/bin/
   ```

## Penggunaan

1. Jalankan aplikasi dengan menentukan direktori target di mana file akan dibuat:
   ```bash
   ./clipfile /path/ke/direktori/target
   ```

2. Salin teks dengan format berikut ke clipboard Anda:

   **Format 1**: Menggunakan komentar garis miring (slash comments)
   ```
   // path/ke/file.txt
   Ini adalah konten file.
   Bisa multi-baris.
   ```

   **Format 2**: Menggunakan komentar HTML
   ```
   <!-- path/ke/file.txt -->
   Ini adalah konten file.
   Bisa multi-baris.
   ```

3. Aplikasi akan secara otomatis membuat atau memperbarui file di lokasi yang ditentukan relatif terhadap direktori target.

## Contoh Penggunaan

### Contoh 1: Membuat file konfigurasi

Salin teks berikut ke clipboard:
```
// config/settings.json
{
  "appName": "MyApp",
  "version": "1.0.0",
  "debug": true,
  "database": {
    "host": "localhost",
    "port": 5432
  }
}
```

Aplikasi akan membuat file `settings.json` di dalam folder `config` di direktori target Anda.

### Contoh 2: Membuat file HTML

Salin teks berikut ke clipboard:
```
<!-- web/index.html -->
<!DOCTYPE html>
<html>
<head>
    <title>Halaman Saya</title>
</head>
<body>
    <h1>Selamat Datang!</h1>
    <p>Ini adalah halaman contoh.</p>
</body>
</html>
```

Aplikasi akan membuat file `index.html` di dalam folder `web` di direktori target Anda.

## Batasan dan Keamanan

- Path file tidak boleh kosong
- Path file harus relatif (tidak boleh dimulai dengan `/` atau mengandung `..`)
- Path file divalidasi untuk mencegah serangan path traversal
- Aplikasi akan membuat direktori induk jika belum ada

## Cara Kerja

Aplikasi ini bekerja dengan:
1. Memantau clipboard sistem setiap 2 detik
2. Memeriksa apakah konten clipboard cocok dengan format yang ditentukan
3. Mengekstrak path file relatif dan konten file
4. Membuat struktur direktori yang diperlukan
5. Menulis atau memperbarui file target
6. Melaporkan status operasi (dibuat/diubah)

## Kontribusi

Kontribusi dan saran selalu diterima! Silakan buka issue atau pull request untuk perbaikan atau fitur baru.

## Lisensi

Proyek ini dilisensikan di bawah [Lisensi MIT](LICENSE).