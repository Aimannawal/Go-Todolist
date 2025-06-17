## 📄 `README.md`

````markdown
# 📝 Go Native ToDo List API (MongoDB)

API ToDo List sederhana menggunakan Go **tanpa framework** (`net/http`) dan MongoDB sebagai database.

---

## 📦 Fitur

- ✅ Tambah catatan (POST)
- 📃 Lihat semua catatan (GET)
- ❌ Hapus catatan berdasarkan ID (DELETE)

---

## ⚙️ Cara Install & Setup (Linux)

### 1. Clone Proyek
```bash
git clone https://github.com/yourname/go-todolist-native.git
cd go-todolist-native
````

### 2. Install Golang

> Pastikan Go versi minimal `1.18` sudah terpasang.
> Jika belum, kamu bisa install:

```bash
sudo apt update
sudo apt install golang
```

### 3. Inisialisasi Module & Install Dependency

```bash
go mod tidy
```

### 4. Jalankan MongoDB

Pastikan MongoDB sudah berjalan di `localhost:27017`.

Jika belum punya MongoDB:

```bash
sudo apt install mongodb
sudo systemctl start mongodb
```

### 5. Jalankan Server

```bash
go run main.go
```

API berjalan di: `http://localhost:8080`

---

## 🔌 Dokumentasi API

### 🔹 `GET /notes`

Ambil semua catatan.

* **Method**: `GET`
* **URL**: `/notes`
* **Response**:

```json
[
  {
    "id": "666f2d...",
    "title": "Belajar Go Native",
    "done": false,
    "created_at": "2025-06-15T07:00:00Z"
  }
]
```

---

### 🔹 `POST /notes`

Tambah catatan baru.

* **Method**: `POST`
* **URL**: `/notes`
* **Header**:

  * `Content-Type: application/json`
* **Body**:

```json
{
  "title": "Belajar MongoDB"
}
```

* **Response**:

```json
{
  "id": "666f2d...",
  "title": "Belajar MongoDB",
  "done": false,
  "created_at": "2025-06-15T07:05:00Z"
}
```

---

### 🔹 `DELETE /notes/{id}`

Hapus catatan berdasarkan ID.

* **Method**: `DELETE`
* **URL**: `/notes/<id>`
* **Response**:

```json
{
  "message": "Note berhasil dihapus"
}
```

---

## 🐘 Contoh Tes via Curl

### Ambil Semua Notes

```bash
curl http://localhost:8080/notes
```

### Tambah Note Baru

```bash
curl -X POST http://localhost:8080/notes \
-H "Content-Type: application/json" \
-d '{"title": "Belajar Golang Native"}'
```

### Hapus Note

```bash
curl -X DELETE http://localhost:8080/notes/<id>
```

---

## ✨ Saran Pengembangan

* Tambahkan fitur `PUT` untuk edit catatan
* Tambahkan autentikasi JWT
* Tambahkan validasi input
* Deploy ke VPS (contoh: Ubuntu 22.04 + PM2 + nginx reverse proxy)

---

## 🧑‍💻 Dibuat oleh

Aiman Wafi’i An Nawal
SMK Telkom Sidoarjo – Backend Developer
