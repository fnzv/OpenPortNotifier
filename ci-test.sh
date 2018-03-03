# OPN Installer

 go get ./src
 go build -o binaries/opn src/opn.go
 sudo cp binaries/opn /bin
 sudo mkdir /etc/opn
 sudo cp conf/opn.conf.example /etc/opn/opn.conf
 sudo apt-get install -y postfix mysql-server masscan
