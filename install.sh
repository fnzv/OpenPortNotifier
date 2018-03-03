# OPN Installer
cp binaries/* /bin/
mkdir /etc/opn/
cp conf/opn.conf.example /etc/opn/opn.conf
apt-get install -y postfix mysql-server masscan
echo "Creating Database opn.."
mysql -p -e "create database opn"
echo "Creating Table for opn.."
mysql opn -p -e "create table opnscans ( IP varchar(20),PORT varchar(20), createdDate datetime);"
