# OpenPortNotifier
[![Build Status](https://travis-ci.org/fnzv/OpenPortNotifier.svg?branch=master)](https://travis-ci.org/fnzv/OpenPortNotifier) <br>
Simple tool to monitor network changes over time and trigger alerts

## Demo
[![asciicast](https://asciinema.org/a/cUL1ksv8JaNrZvM2PgNUjyuyj.png)](https://asciinema.org/a/cUL1ksv8JaNrZvM2PgNUjyuyj)

## Requirements

- Golang (optional if you compile the sources)
- masscan 
- postfix/smtp relay (optional for want E-mail reports)
- mysql (optional if you want to track changes over time)


### Quickstart

Run the bash script (install.sh) to install all the required dependencies.

```bash install.sh```

Downloads the binaries and install them into your system, after you changed the example configuration (/etc/opn/opn.conf) you can run the scans with the command:

```opn```

### Use cases
- Continuos scanning of your Networks and monitor service exposure (All scans are saved into Mysql and/or sent via email/tg).. "pro-tip" -> run on crontab -> http://crontab.guru/
- Constant Alerting of critical service or when a firewall (could be software or hardware) stops working
- Tracking service/hosts from your VPN/Allowed cidr to know where your services are without running everytime the slow Nmap 

### Compile

To compile and install binaries run: 

```bash compile.sh```

### How OPN works
- Reads from /etc/opn/opn.conf configuration, if these values (SMTPhost,MysqlAuth,Telegram) are "" will be skipped and the feature is omitted
- Scans the Networks using Masscan (because nmap is slow)
- Send output to screen and trigger alerts if ports open are on "CriticalPorts" list
- All the scans are saved into the configured database 

