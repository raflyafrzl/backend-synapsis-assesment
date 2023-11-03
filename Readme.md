## Synapsis Backend
merupakan hasil dari challenge yang diberikan oleh synapsis

### HOW TO USE IT 

1. Pertama-tama install dependencies terlebih dahulu

```shell
go mod tidy
```

2. kemudian jalankan image postgres pada docker compose 
```shell
docker-compose up
```

3. tambahkan .env file  pada project app dan isikan nilainya. Untuk contoh environtment variable yang disediakan bisa lihat ``.env.example``

4. jalankan migrasi golang, namun sebelum menjalankan pastikan sudah menginstall depdency nya sebagai CLI
 ```shell
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
5. jalankan file migration up nya(merujuk ke docker-compose)
```shell
    migrate -database "postgres://postgres:postgres@localhost:5431/synapsisdb?sslmode=disable" -path ./migrations up
```
6. setelah semua step dijalankan kita bisa menjalankan app nya
```shell
go run main.go
```


## Endpoint 

1. ``GET /`` halaman utama
2. ``GET /api/customers`` melihat list kustomer
3. ``POST /api/customers/register`` untuk mendaftarkan akun customers. payload yang dibutuhkan adalah: `username dan password`
4. ``POST /api/customers/login`` untuk melakukan login dan mendapatkan token. payload yang dibutuhkan adalah: `username dan password`
5.  ``GET /api/products/{category}`` untuk mendapatkan list of products berdasarkan category yang dimasukan di URLParams <br> <br>
=======Untuk URL dibawah ini membutuhkan token(request header authorization)=======
6.  ``DELETE /api/products/ `` menghapus product dalam shopping cart, membutuhkan query params `id`
7.  ``POST /api/products/cart`` menambahkan product kedalam cart. payload yang dibutuhkan adalah `product_id`
8.  ``GET /api/products/list `` melihat list pada cart yang sudah ditambahkan oleh customers
9.  ``GET /api/products/checkout`` checkout products yang ada didalam cart.


## ERD
<p align="center">
  <img src="./erd.png" width="80%">
</p>


    