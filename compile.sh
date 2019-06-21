
# Golang quick install

 apt-get update
 wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz
 sudo tar -xvf go1.9.2.linux-amd64.tar.gz
 sudo mv go /usr/local
 echo "export GOROOT=/usr/local/go" >> /root/.bashrc
 echo "export GOPATH=$HOME/Projects" >> /root/.bashrc
 echo "export PATH=$GOPATH/bin:$GOROOT/bin:$PATH" >> /root/.bashrc
 echo "Checking Golang version\n"
 ln -s /usr/local/go/bin/go /usr/bin/go

# OPN Installer

 go get ./src
 go build -o binaries/opn src/opn.go
 cp binaries/opn /bin
 mkdir /etc/opn
 sudo apt-get install -y masscan
 cp conf/opn.conf.example /etc/opn/opn.conf
 echo "Do you wish to install mysql-server and postfix on this server?"
 select yn in "Yes" "No"; do
    case $yn in
        Yes ) sudo apt-get install -y postfix mysql-server; echo "Creating Database opn..";  mysql -p -e "create database opn";  echo "Creating Table for opn.."; mysql opn  -e "create table opnscans (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, IP varchar(20),PORT varchar(20), createdDate datetime);";
 break;;
        No ) exit;;
    esac
done
 
