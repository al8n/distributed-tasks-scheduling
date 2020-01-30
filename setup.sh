sudo chmod 777 /etc/apk
sudo rm /etc/apk/repositories
sudo echo https://mirrors.aliyun.com/alpine/v3.10/main >> /etc/apk/repositories
sudo echo https://mirrors.aliyun.com/alpine/v3.10/community >> /etc/apk/repositories
sudo chmod 755 /etc/apk
sudo apk update
sudo apk --no-cache add docker nginx py-pip python-dev libffi-dev openssl-dev gcc git libc-dev make
sudo adduser -D -g "nginx" nginx
sudo mkdir /nginx
sudo chown -R nginx:nginx /var/lib/nginx
sudo chown -R nginx:nginx /nginx
sudo rc-service nginx start
sudo rc-update add docker boot
sudo service docker start
sudo chmod 777 /etc/hosts
sudo echo 192.30.253.112  github.com >> /etc/hosts
sudo echo 199.232.5.194   github.global.ssl.fastly.net >> /etc/hosts
sudo chmod -wx /etc/hosts
sudo curl -L "https://github.com/docker/compose/releases/download/1.25.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
