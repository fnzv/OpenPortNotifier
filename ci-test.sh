# OPN Installer

 go get ./src
 go build -o binaries/opn src/opn.go
 cp binaries/opn /bin
 mkdir /etc/opn
 cp conf/opn.conf.example /etc/opn/opn.conf
 apt-get install -y postfix mysql-server masscan
