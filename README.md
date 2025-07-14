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
- **POST /api/v1/vendor** : Tambah vendor
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **BODY:**
    ```json
    {
      "name": "Vendor Name",
      "description": "Vendor Description"
    }
    ```
- **GET /api/v1/vendor** : List semua vendor
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **Query Parameters:**
    - `page`: Nomor halaman (default 1)
    - `limit`: Jumlah data per halaman (default 10)
- **GET /api/v1/vendor?page=1&limit=10** : List vendor dengan   pagination
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **Query Parameters:**
    - `page`: Nomor halaman (default 1)
    - `limit`: Jumlah data per halaman (default 10)
- **GET /api/v1/vendor/{id}** : Detail vendor
  - **Bearers:**
    - **Authorization:** Bearer token dari login
- **PUT /api/v1/vendor/{id}** : Update vendor
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **BODY:**
    ```json
    {
      "name": "Updated Vendor Name",
      "description": "Updated Vendor Description"
    }
    ```
- **DELETE /api/v1/vendor/{id}** : Hapus vendor
  - **Bearers:**
    - **Authorization:** Bearer token dari login

#### CRUD Produk Vendor
- **GET /api/v1/vendor/product** : List semua produk vendor
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **Query Parameters:**
    - `page`: Nomor halaman (default 1)
    - `limit`: Jumlah data per halaman (default 10)
- **GET /api/v1/vendor/product?page=1&limit=10** : List produk vendor dengan pagination
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **Query Parameters:**
    - `page`: Nomor halaman (default 1)
    - `limit`: Jumlah data per halaman (default 10)
- **POST /api/v1/vendor/product** : Tambah produk
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **BODY:**
    ```json
    {
      "product_name":"karbu",
      "product_price": 100000,
      "product_description":"otomootip",
      "product_category_id":"da698079-ba94-4064-bb6f-f894871f9711",
      "vendor_id":"12d17ede-8dc7-4c7f-ad63-0a9f457c01a3"
    }
    ```
- **GET /api/v1/vendor/product/{id}** : Detail produk
  - **Bearers:**
    - **Authorization:** Bearer token dari login

- **PUT /api/v1/vendor/product/{id}** : Update 
produk
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **BODY:**
    ```json
    {
      "product_name":"Updated Product Name",
      "product_price": 120000,
      "product_description":"Updated Description",
      "product_category_id":"da698079-ba94-4064-bb6f-f894871f9711",
    }
    ```
- **DELETE /api/v1/vendor/product/{id}** : 
Hapus produk
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  

### 3. CRUD User
- **GET /api/v1/user** : Detail user
  - **Bearers:**
    - **Authorization:** Bearer token dari login
    - **Query Parameters:**
    - `Id`: ID produk yang akan dihapus
    (default 10)
- **PUT /api/v1/user/{id}** : Update user
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **BODY:**
    ```json
    {
      "name": "Updated User Name",
      "email": "EMAI UPDATED",
      "password": "newpassword"
    }
    ```
- **DELETE /api/v1/user** : Hapus user
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **Query Parameters:**
    - `Id`: ID user yang akan dihapus
- **PUT /api/v1/user/password** : Update user
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **BODY:**
    ```json
    {
      "old_password": "oldpassword",
      "new_password": "newpassword"
    }
    ```
## 4. Kategori Produk
- **GET /api/v1/category** : List semua kategori produk
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **Query Parameters:**
    - `page`: Nomor halaman (default 1)
    - `limit`: Jumlah data per halaman (default 10)
- **GET /api/v1/category?page=1&limit=10** : List kategori produk dengan pagination
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **Query Parameters:**
    - `page`: Nomor halaman (default 1)
    - `limit`: Jumlah data per halaman (default 10)
- **POST /api/v1/category** : Tambah kategori produk
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **BODY:**
    ```json
    {
      "category_name": "Category Name",
      "category_description": "Category Description"
    }
    ```
- **Put /api/v1/category/{id}** : Update kategori produk
  - **Bearers:**
    - **Authorization:** Bearer token dari login
  - **BODY:**
    ```json
    {
      "category_name": "Updated Category Name",
      "category_description": "Updated Category Description"
    }
    ```
- **DELETE /api/v1/category/{id}** : Hapus kategori produk
  - **Bearers:**
    - **Authorization:** Bearer token dari login
## Catatan
- Pastikan environment database sudah berjalan.
- Gunakan tools seperti Postman untuk menguji endpoint API.

---

Kontribusi dan saran sangat terbuka untuk pengembangan lebih lanjut.
