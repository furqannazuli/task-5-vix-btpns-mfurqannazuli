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
<ul>
    <li>postgresup (untuk melakukan pull postgres docker image dan set up volume untuk persistance storage)</li>
    <li>postgresdown (untuk menghentikan proses postgres docker image yang sedang berjalan pada suatu container dan membuang cache container)</li>
    <li>psql (untuk memulai interactive shell dari container postgres yang sedang berjalan)</li>
</ul>
`Perintah : make <target>, Contoh : make postgresup`

## Cara Menjalankan Program
<p style="margin-left:10px">
    Terdapat beberapa tahapan dalam menjalankan <i>Restful API</i> ini, yaitu: Development, Testing dan Production.
</p>

- Development 
    - Untuk menjalankan tahapan development ini, harus melalui beberapa konfigurasi yaitu:
        - Ubah DB_HOST menjadi localhost (file <b>Dockerfile</b>)
        - Ubah DB_PORT menjadi 5433 (file <b>Dockerfile</b>)
        - Ubah STAGE menjadi development (file <b>Dockerfile</b>)
        - Jalankan perintah `make postgresup` di terminal
        - Jalankan perintah `go run main.go` di terminal
- Testing
    - Untuk menjalankan tahapan testing ini, harus melalui beberapa konfigurasi yaitu:
        - Ubah DB_HOST menjadi localhost (file <b>Dockerfile</b>)
        - Ubah DB_PORT menjadi 5433 (file <b>Dockerfile</b>)
        - Ubah STAGE menjadi testing (file <b>Dockerfile</b>)
        - Jalankan perintah `make test` di terminal
- Production
    - Untuk menjalankan tahapan production (dengan container), harus melalui beberapa konfigurasi yaitu:
        - Anda tidak perlu melakukan konfigurasi pada file .env, karena untuk keperluan demonstrasi, saya telah melakukan konfigurasi ENV Variable pada Dockerfile
        - Jalankan perintah `make build` di terminal
        - Jalankan perintah `make run` di terminal