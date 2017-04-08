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
  jika sudah berhasil kita bisa mengeceknya menggunakan perintah di atas untuk cek versi docker-machine
  <p align="center"><img src="images/2.png"/><br>Gambar 1.2</p>
## How to create node on docker-machine
    Untuk membuat node pada docker-machine, di sistem anda harus terinstall virtual machine seperti (virtual box, qemu, vmware, dll), untuk tutorial ini saya menggunakan virtual box.
    barikut cara membuat workers atau node dalam virtual machine:
      * docker-machine create --driver virtualbox node1
      * docker-machine create --driver virtualbox node2
      * docker-machine create --driver virtualbox node3
    Jika di setiap node anda ingin mengatur resourcenya, seperti hdd, ram, processor dll. bisa menggunakan perintah di bawah ini
      * docker-machine create --driver virtualbox --virtualbox-memory "512" --virtualbox-disk-size "5000" node4

    selanjutnya cek docker-machine anda dengan menggunakan perintah docker-machine ls

    <p align="center"><img src="images/3.png"/><br>Gambar 1.3</p>
## How to join
### konfigurasi pada manager
  * docker swarm init –advertise-addr (ip address manager)
    ip ini saya ambil dari ip virtual-box saya.

    <p align="center"><img src="images/4.png"/><br>Gambar 1.4</p>

    jalanakan perintah berikut untuk melakukan inisialisasi
    <p align="center"><img src="images/5.png"/><br>Gambar 1.5</p>
    `docker swarm init --advertise-addr 192.168.99.1`

    jalanakan perintah berikut untuk mengecek token
    <p align="center"><img src="images/6.png"/><br>Gambar 1.6</p>
    `docker swarm join-token -q worker`

### konfigurasi pada node untuk melakukan join.
  akses ke stiap node untuk melakukan join
  `docker-machine ssh nama node`
    <p align="center"><img src="images/7.png"/><br>Gambar 1.7</p>
  * docker swarm join (ip address manager):2377 --token(token pada manager)
    `docker swarm join 192.168.99.1:2377 --token SWMTKN-1-0aj4hqxy0gb04kq29r6kcxh8v3mz8z6k0wccpwftnvkv4lxtcp-cuexjpxih5oyycypc85u5ropv`

  <p align="center"><img src="images/8.png"/><br>Gambar 1.8</p>

## Deploy service to Docker-Swarm

## How to scale

## Configuring resource
