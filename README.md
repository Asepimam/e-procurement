# E-Procurement Multi Vendor

E-Procurement adalah aplikasi backend berbasis Go untuk sistem multi vendor, memungkinkan pengelolaan produk, vendor, dan user secara efisien. Sistem ini menyediakan fitur CRUD (Create, Read, Update, Delete) untuk produk, vendor, dan user, serta autentikasi login dan registrasi.

## Fitur Utama
- **Multi Vendor:** Mendukung banyak vendor untuk mengelola katalog produk masing-masing.
- **CRUD Produk Vendor:** Vendor dapat membuat, membaca, memperbarui, dan menghapus produk mereka.
- **CRUD User:** Admin dapat mengelola data user.
- **Autentikasi:** Login dan registrasi user.

## Struktur Project
- `cmd/` : Entry point aplikasi
- `internals/` : Business logic, delivery, domain, repository, usecase
- `pkg/` : Package utilitas (database, JWT, response)

## Instalasi
1. **Clone repository**
   ```bash
   git clone <repo-url>
   cd e-procurement
   ```
2. **Install dependency**
   ```bash
   go mod tidy
   ```
3. **Konfigurasi database**
   - Edit konfigurasi database di `pkg/connections/db.go` sesuai kebutuhan.
4. **Jalankan aplikasi**
   ```bash
   go run cmd/main.go
   ```

## Dokumentasi API

### 1. Autentikasi
#### Login
- **Endpoint:** `POST /api/v1/login`
- **Body:**
  ```json
  {
    "email": "user@example.com",
    "password": "yourpassword"
  }
  ```
- **Response:** JWT Token

#### Register
- **Endpoint:** `POST /api/v1/register`
- **Body:**
  ```json
  {
    "name": "User Name",
    "email": "user@example.com",
    "password": "yourpassword"
  }
  ```
- **Response:** Data user terdaftar

### 2. Vendor & Katalog Produk
#### CRUD Vendor
- **GET /api/v1/vendor** : List semua vendor
- **POST /api/v1/vendor** : Tambah vendor
- **GET /api/v1/vendor/{id}** : Detail vendor
- **PUT /api/v1/vendor/{id}** : Update vendor
- **DELETE /api/v1/vendor/{id}** : Hapus vendor

#### CRUD Produk Vendor
- **GET /api/v1/vendor/catalog** : List semua produk vendor
- **POST /api/v1/vendor/catalog** : Tambah produk
- **GET /api/v1/vendor/catalog/{id}** : Detail produk
- **PUT /api/v1/vendor/catalog/{id}** : Update produk
- **DELETE /api/v1/vendor/catalog/{id}** : Hapus produk

### 3. CRUD User
- **GET /api/v1/user** : List semua user
- **POST /api/v1/user** : Tambah user
- **GET /api/v1/user/{id}** : Detail user
- **PUT /api/v1/user/{id}** : Update user
- **DELETE /api/v1/user/{id}** : Hapus user

## Catatan
- Pastikan environment database sudah berjalan.
- Gunakan tools seperti Postman untuk menguji endpoint API.

---

Kontribusi dan saran sangat terbuka untuk pengembangan lebih lanjut.
