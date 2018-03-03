# OpenPortNotifier
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

```root:# opn```

### Use cases
- Continuos scanning of your Networks and monitor service exposure (All scans are saved into Mysql and/or sent via email/tg)
- Constant Alerting of critical service or when a firewall (could be software or hardware) stops working
- Tracking service/hosts from your VPN/Allowed cidr to know where your services are without running everytime the slow Nmap 

### Compile

To compile and install binaries run: 

```bash compile.sh```
