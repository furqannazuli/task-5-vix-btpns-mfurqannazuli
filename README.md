# **Final Project Virtual Internship Rakamin Academy BTPN Syariah** 

<hr/>
<br/>

## Description
<p style="margin-left: 10px">
    Program ini merupakan tugas akhir dari program <b>Virtual Internship Rakamin Academy di BTPN Syariah.</b>
    Para peserta diarahkan untuk membentuk API berdasarkan kasus yang telah diberikan. Menggunakan bahasa pemrograman Golang dan bantuan Docker.
</p>

## Task
<p style="margin-left: 10px">
    Berdasarkan data yang telah diolah oleh tim Data Analysts, bahwa untuk meningkatkan engagement user pada aplikasi m-banking adalah meningkatkan aspek memiliki user pada aplikasi tersebut. Saran yang diberikan oleh tim data analysts adalah membentuk fitur personalize user, salah satunya adalah memungkinkan user dapat mengupload gambar untuk dijadikan foto profilnya. Tim developer bertanggung jawab untuk mengembangkan fitur ini, dan kalian diberikan tugas untuk merancang API pada fitur upload, dan menghapus gambar. Beberapa ketentuannya antara lain :
    <ul>
        <li>User dapat menambahkan foto profile</li>
        <li>Sistem dapat mengidentifikasi User ( log in / sign up)</li>
        <li>Hanya user yang telah log in / sign up yang dapat melakukan delete / tambah foto profil</li>
        <li>User dapat menghapus gambar yang telah di post</li>
        <li>User yang berbeda tidak dapat menghapus / mengubah foto yang telah dibuat oleh user lain</li>
    </ul>
</p>

## Set up
<p style="margin-left: 10px">
    Dalam program ini saya tambahkan file <b>JalanAPI</b> yang memiliki fungsi dalam menjalankan beberapa init script yang terkait dengan <i>Golang Restful API.</i>
</p>
Beberapa perintah yang ada yaitu sebagai berikut :
- postgresup (untuk melakukan pull postgres docker image dan set up volume untuk persistance storage)
- postgresdown (untuk menghentikan proses postgres docker image yang sedang berjalan pada suatu container dan membuang cache container)
- psql (untuk memulai interactive shell dari container postgres yang sedang berjalan)

`Perintah : make <target>, Contoh : make postgresup`

<p style="margin-left: 10px">
    Berikut adalah step-by-step untuk men-setup docker image <i>Golang Restful API</i> dan <i>PostgreSQL</i> secara cepat. Agar dapat secara langsung me-running <i>Restful API</i> nya.
</p>

- `make build` | untuk mem-build image dari <i>Golang Restful API </i> nya
- `make run` | me-running image Golang dan Postgres nya secaras otomatis.

That's it, Restful API nya sudah bisa di testing sekarang, Enjoy :

<br />

## How to Use 
<p style="margin-left:10px">
    There are several stage in this <i>Restful API</i> (Development, Testing & Production)
</p>

- Development 
    - to running development stage you can use this configurations
        - change DB_HOST to localhost
        - change DB_PORT to 5433
        - change STAGE to development
        - run `make postgresup` in your terminal
        - run `go run main.go` in your terminal
- Testing
    - to running testing stage you can use this configurations
        - change DB_HOST to localhost
        - change DB_PORT to 5433
        - change STAGE to testing
        - run `make test` in your terminal
- Production
    - to running production (with container) stage you can use this configurations
        - you dont need to configure .env files, because for demonstration purpose, Iam already Hard-Code the ENV Variable in Dockerfile
        - run `make build` in your terminal
        - run `make run` in your terminal

`If container or image already created before, you can use make delete_container or make delete_image command in Makefile`


## System Design
<br />
<p align="center">
    <img 
        src="./assets/system-design-api.png" width="75%">
</p>

<p style="margin-left:10px">
    User akan di berikan dua <i>Endpoint</i>, yakni <i>User Endpoint</i> dan <i>Photo Endpoint</i>. Dapat di lihat <i>Flow Process</i> dari <i>Restful API</i> yang telah di buat. User akan mengirimkan Request kepada <i>Restful API</i> terkait, lalu akan di teruskan kepada <i>Endpoint</i> yang telah di tuju sebelumnya. Jika <i>Endpoint</i> User, maka akan di suguhkan oleh dua route, yakni Login dan Register. Jika Endpoint Photo, maka akan di suguhkan sebuah <i>CRUD Method</i>. Namun untuk mengubah dan menghapus photo di butuhkan <i>Permission</i> khusus, yakni hanya User yang membuat / mengupload photo tersebut. <i>Middleware</i> akan mem-filter siapa saja User yang dapat melakukan perubahan dan penghapusan terkait dengan photo terkait. 
</p>

<br />

## API Spec

### Authentication
<p style="margin-left:10px">
    Semua <i>API Endpoint</i> harus menggunakan <i>Authentication</i> ini.
</p>

Request :
- Header : 
    - Authorization : JWT Token


### User Register :

Request :
- Method : POST
- Endpoint : `/api/v1/users/register`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :
    ```json
        {
            "id" : "primary key, required",
            "username" : "required",
            "email" : "unique, required",
            "password" : "required, min-length: 6",
            "photo" : "Many-to-One, CASCADE | optional",
        }
    ```

Response :

```json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "string",
            "email" : "string",
            "username" : "string",
            "created_at" : "timestamp"
        }
    }
```

### User Login

Request :

- Method : POST
- Endpoint : `/api/v1/users/login`
- Header : 
    - Accept: application/json
- Body:
    ```json
        {
            "email" : "string",
            "password" : "string"
        }
    ```

Response :

```json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "token" : "string"
        }
    }
```

### Photo 
- Method : GET 
- Endpoint : `/api/v1/photo`
- Header : 
    - Accept : application/json
    - 
Response :

```json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "uint",
            "title" : "string",
            "photo_url" : "string",
            "caption" : "string",
            "user_id" : "uint",
            "user" : "Many-to-One"
        }
    }
```

- Method : POST 
- Endpoint : `/api/v1/photo`
- Header : 
    - Accept : application/json
    - Authorization : JWT-Token

Request :
```json
    {
        "title" : "string",
        "photo_url" : "string",
        "caption" : "string | optional",
        "user_id" : "uint",
    }
```

Response :

```json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "uint",
            "title" : "string",
            "photo_url" : "string",
            "user_id" : "uint",
            "user" : "Many-to-One",
            "created_at" : "timestamp",
        }
    }
```

- Method : PUT 
- Endpoint : `/api/v1/photo`
- Header : 
    - Accept : application/json

Request :
```json
    {
        "user_id" : "uint",
    }
```

Response :

```json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "uint",
            "title" : "string",
            "photo_url" : "string",
            "user_id" : "uint",
            "user" : "Many-to-One",
            "updated_at" : "timestamp",
        }
    }
```

- Method : DELETE 
- Endpoint : `/api/v1/photo`
- Header : 
    - Accept : application/json

Request :
```json
    {
        "user_id" : "uint",
    }
```

Response :

```json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "message" : "string"
        }
    }
```

## API Test

<p style="margin-left:10px">
    Terkait dengan <i>Restful API</i> nya sendiri dapat di running atau testing dengan menggunakan file <b>test.http</b>. Dengan memanfaatkan RestClient Extension yang terdapat pada <b>VSCode</b> atau <b>IntelliJ IDEA</b>
</p>