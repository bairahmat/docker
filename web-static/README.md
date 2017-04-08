# build image
Build image adalah proses di mana ketika kita ingin membuat image di docker. Image adalah istilah yang di berikan Docker Inc.
~~~
docker build -t username/name-image:version .
~~~


# Menjalankan docker image jiharalgifari/web-nginx:v1
Jika anda ingin menjalankan images berikut langkah langkanya.
~~~
  docker pull jiharalgifari/web-nginx:v1
  docker run -d -p 80:80 jiharalgifari/web-nginx:v1
~~~
setelah itu anda bisa mngecek di browser url:'ip address:80'. ex : http://127.0.0.1:80
