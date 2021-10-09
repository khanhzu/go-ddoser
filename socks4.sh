while true; do
./SOCKS4 $1 $2 $3 $4 > /dev/null 2>&1 & 
sleep 10s
sudo pkill SOCKS4
done
