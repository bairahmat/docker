# Docker Swarm

## Overview Docker machine
<p align="center"><img src="images/1.png"/><br>Gambar 1.1</p>
Gambar di atas menjelaskan sekilas tentang arsitektur dalam docker swarm, di mana terdapat 1 manager dan 3 workers

### kenapa docker machine

  Docker-machine adalah sebuah tool yang kita install pada docker -engine.
  Docker-machine berjalan di atas virtual host. Fungsinya adalah untuk memanage hosts di komputer local kita. Dengan docker-machine kita bisa membuat host  di komputer local, atau virtual box, dapat gunakan untuk data-center, AWS, dan Digital Ocean, dll.

## How to install docker machine

* curl -L https://github.com/docker/machine/releases/download/v0.8.2/docker-machine-`uname -s`-`uname -m` >/usr/local/bin/docker-machine

* chmod +x /usr/local/bin/docker-machine

* docker-machine -v

  jika sudah berhasil kita bisa mengeceknya menggunakan perintah di atas untuk cek
  versi docker-machine

  <p align="center"><img src="images/2.png"/><br>Gambar 1.2</p>

## How to create node on docker-machine

  Untuk membuat node pada docker-machine, di sistem anda harus terinstall virtual machine seperti (virtual box, qemu, vmware, dll), untuk tutorial ini saya menggunakan virtual box.

  berikut cara membuat workers atau node dalam virtual machine:

* docker-machine create --driver virtualbox node1
* docker-machine create --driver virtualbox node2
* docker-machine create --driver virtualbox node3

Gunakan perintah di bawah ini mengatur resourcenya, seperti hdd, ram, processor dll.

* docker-machine create --driver virtualbox --virtualbox-memory "512" --virtualbox-disk-size "5000" node4

selanjutnya cek docker-machine anda dengan menggunakan perintah docker-machine ls
<p align="center"><img src="images/3.png"/><br>Gambar 1.3</p>

## How to join
  Join di gunakan untuk menggabungkan beberapa node dan menjadikan salah satu sebagai leader dan lainya sebagai workers.

### konfigurasi pada manager

  * docker swarm init –advertise-addr (ip address manager)
    ip ini saya ambil dari ip virtual-box saya.

    <p align="center"><img src="images/4.png"/><br>Gambar 1.4</p>

**inisialisasi**

  <p align="center"><img src="images/5.png"/><br>Gambar 1.5</p>
  `docker swarm init --advertise-addr 192.168.99.1`

**mengecek token**

  <p align="center"><img src="images/6.png"/><br>Gambar 1.6</p>
  `docker swarm join-token -q worker`

### konfigurasi pada node untuk melakukan join.

  **Mengakses ke masing masing node**
  `docker-machine ssh nama node`
  <p align="center"><img src="images/7.png"/><br>Gambar 1.7</p>

* docker swarm join (ip address manager):2377 --token(token pada manager)
  `docker swarm join 192.168.99.1:2377 --token SWMTKN-1-0aj4hqxy0gb04kq29r6kcxh8v3mz8z6k0wccpwftnvkv4lxtcp-cuexjpxih5oyycypc85u5ropv`

  <p align="center"><img src="images/8.png"/><br>Gambar 1.8</p>

  selanjutnya cek status node pada para terminaluntuk mengetahui leader dan workers, gunakan perintah
  `docker node ls`
  <p align="center"><img src="images/9.png"/><br>Gambar 1.9</p>

## Deploy service to Docker-Swarm

#### Deploy menggunakan images docker.
  Untuk menjalankan service hal pertama yang harus anda punya adalah image App yang akan di jalankan, pada panduan ini saya menggunakan images saya sendiri yaitu `jiharalgifari/web-nginx:v1` yang saya akan jadikan sebuah service. berikut implementasinya

  `docker service create --name web-saya -p 80:80 --replicas 2 jiharalgifari/web-nginx:v1`

  perintah di atas bermaksud untuk menjalankan service dengan nama `web-saya` yang akan di publish ke port 80 dengan replicas 2.

  untuk mengecek service yang berjalan bisa menggunakan perintah `docker service ls`
  <p align="center"><img src="images/10.png"/><br>Gambar 1.10</p>
  selanjutnya kita bisa mengecek service yang berjalan di setiap node1

  1. Pada komputer lokal localhost (localhost:80)
    <p align="center"><img src="images/11.png"/><br>Gambar 1.11</p>

  2. pada node1 (http://192.168.99.100/)
    <p align="center"><img src="images/12.png"/><br>Gambar 1.12</p>

  3. pada node2 (http://192.168.99.101/)
    <p align="center"><img src="images/13.png"/><br>Gambar 1.13</p>

  4. pada node3 (http://192.168.99.102/)
    <p align="center"><img src="images/14.png"/><br>Gambar 1.14</p>

#### Deploy menggunakan docker compose
Jika anda ingin menjalankan service yang ada dalam docker compose gunakan perintah di bawah ini:
  ~~~bash
    docker stack deploy
  ~~~
**Demo menggunakan wordpress**

Untuk menjalankan service dalam docker swarm minimal docker-compose.yml nya versi 3.0 berikut contohnnya
~~~yml
version: '3.0'

services:
   db:
     image: mysql:5.7
     volumes:
       - db_data:/var/lib/mysql
     restart: always
     environment:
       MYSQL_ROOT_PASSWORD: wordpress
       MYSQL_DATABASE: wordpress
       MYSQL_USER: wordpress
       MYSQL_PASSWORD: wordpress

   wordpress:
     depends_on:
       - db
     image: wordpress:latest
     ports:
       - "8000:80"
     restart: always
     environment:
       WORDPRESS_DB_HOST: db:3306
       WORDPRESS_DB_PASSWORD: wordpress
volumes:
    db_data:
~~~

sebelumnya coba anda menjalankan secara manual, untuk memastikan apakah docker composenya jalan atau tidak.

`docker-compose up`

jika sudah jalan anda bisa berpindah ke docker swarm. sebelumnya anda harus mematika semua service yang ada dulu menggunakan perintah `docker stop $(docker ps -a -q)` agar tidak ada tabrakan saat menjalankan service. berikut adalah contoh sebelum menggunakan swarm:
  <p align="center"><img src="images/18.png"/><br>Gambar 1.15</p>


Sekarang kita akan menjalankan service docker-compose dalam swarm.
1. docker stack deploy --compose-file docker-compose.yml web
<p align="center"><img src="images/19.png"/></p>
2. docker stack ls
  <p align="center"><img src="images/20.png"/></p>
3. selanjutnya anda bisa mengeceknya melalui browser seperti (chrome, firefox) dengan memanggil ip:port (192.168.99.100:8000) atau node lainya. berdasarkan ip. berikut contoh dalam localhost:8000.  <p align="center"><img src="images/18.png"/></p>



## How to scale

  untuk melakukan scale service anda bisa menggunakan perintah
  `docker service scale web-saya=8`
  dan untuk mengecek hasil scale bisa menggunakan perintah `docker service ps web-saya`

  di sini kita bisa melihat node apa saja yang menangani service, nama image yang jalan, dan waktu ketika service start.

  <p align="center"><img src="images/15.png"/><br>Gambar 1.15</p>

## Configuring resource

Konfigurasi resource bermaksud untuk mengkonfigurasi kebutuhan Ram, Processor  dengan dengan batas tertentu yang di pakai saat service sedang jalan.

untuk menlakukan konfigurasi bisa menggunakan perintah
`docker service update --limit-cpu 1 --limit-memory 512mb web-saya`

maksud dari perintah di atas adalah untuk mengatur service web-saya dengan limit CPU 1 dan limit memory 512MB.
<p align="center"><img src="images/16.png"/><br>Gambar 1.16</p>

Untuk mengeceknya bisa menggunakan perintah `docker service inspect --pretty web-saya`
<p align="center"><img src="images/17.png"/><br>Gambar 1.17</p>
