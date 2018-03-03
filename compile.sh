
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
 cp conf/opn.conf.example /etc/opn/opn.conf
 apt-get install -y postfix mysql-server masscan
 echo "Creating Database opn.."
 mysql -p -e "create database opn"
 echo "Creating Table for opn.."
 mysql opn -p -e "create table opnscans ( IP varchar(20),PORT varchar(20), createdDate datetime);"
