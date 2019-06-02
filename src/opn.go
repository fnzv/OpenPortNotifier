package main

import (
        "bufio"
        "bytes"
        "database/sql"
        "log"
        "net/smtp"
        "os"
        "os/exec"
        "strings"

        "github.com/BurntSushi/toml"
        _ "github.com/Go-SQL-Driver/MySQL"
        "gopkg.in/telegram-bot-api.v4"
)

var (
        networks       []string
        whitelist      []string
        ports          string
        critical_ports string
        telegram       string
        chatid         int64
        mailsrc        string
        maildst        string
        mailhost       string
        mysql_auth     string
)

type Config struct {
        Networks      []string
        Whitelist     []string
        Ports         string
        CriticalPorts string
        Telegram      string
        Chatid        int64
        MailFrom      string
        MailTo        string
        SMTPhost      string
        MysqlAuth     string
}

func checkErr(err error) {
        if err != nil {
                panic(err)
        }
}

func SendMail(mailsrc string, maildst string, mailhost string, text string) {
        c, err := smtp.Dial(mailhost)
        if err != nil {
                log.Fatal(err)
        }
        defer c.Close()
        c.Mail(mailsrc)
        c.Rcpt(maildst)
        wc, err := c.Data()
        if err != nil {
                log.Fatal(err)
        }
        defer wc.Close()
        buf := bytes.NewBufferString(text)
        if _, err = buf.WriteTo(wc); err != nil {
                log.Fatal(err)
        }
}

func ReadConfig() Config {
        var configfile = "/etc/opn/opn.conf"
        _, err := os.Stat(configfile)
        if err != nil {
                log.Fatal("Config file is missing: ", configfile)
        }

        var config Config
        if _, err := toml.DecodeFile(configfile, &config); err != nil {
                log.Fatal(err)
        }
        return config
}

func exec_shell(command string) string {
        out, err := exec.Command("/bin/bash", "-c", command).Output()
        if err != nil {
                log.Fatal(err)
        }
        return string(out)
}

func main() {

        var config = ReadConfig()

        f, err := os.OpenFile("/var/log/opn.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
        if err != nil {
                log.Fatal(err)
        }
        defer f.Close()
        log.SetOutput(f)
        log.Println("OPN started")

        networks = config.Networks
        whitelist = config.Whitelist
        ports = config.Ports
        critical_ports = config.CriticalPorts
        telegram = config.Telegram
        chatid = config.Chatid
        mailsrc = config.MailFrom
        maildst = config.MailTo
        mailhost = config.SMTPhost
        mysql_auth = config.MysqlAuth
        complete_text := "\r\nMessage Body:\n"

        i2 := 0
        for range whitelist {
                host := whitelist[i2]
                cmd := "echo '" + host + "' >> /tmp/whitelist"
                exec_shell(cmd)
        }

        i := 0
        for range networks {
                network := networks[i]
                cmd := "masscan --excludefile /tmp/whitelist -p" + ports + " " + network + " -oL tempoutput"
                log.Println("Scanning Ports " + ports + " on Network " + network)
                exec_shell(cmd)
                exec_shell("sed -i '/masscan/d' tempoutput")
                exec_shell("sed -i '/end/d' tempoutput")
                file, err := os.Open("tempoutput")
                if err != nil {
                        log.Fatal(err)
                }
                defer file.Close()

                scanner := bufio.NewScanner(file)
                for scanner.Scan() {
                        raw := scanner.Text()
                        line := strings.Fields(raw)
                        port := line[2]
                        ip := line[3]
                        if mysql_auth != "" {
                                db, err := sql.Open("mysql", mysql_auth+"?charset=utf8")
                                stmt, err := db.Prepare("INSERT opnscans SET IP=?,PORT=?,createdDate=(SELECT now())")
                                checkErr(err)
                                stmt.Exec(ip, port)
                                stmt.Close()
                        }
                        if strings.Contains(critical_ports, port) {
                                log.Println("CRITICAL PORT " + port + " FOUND OPEN ON HOST " + ip)
                                complete_text = complete_text + " \xE2\x9A\xA0 CRITICAL PORT " + port + " FOUND OPEN ON HOST " + ip + "\n"
                                text := " \xE2\x9A\xA0  CRITICAL PORT " + port + " FOUND OPEN ON HOST " + ip
                                if telegram != "" {
                                        bot, err := tgbotapi.NewBotAPI(telegram)
                                        if err != nil {
                                                log.Panic(err)
                                        }
                                        msg := tgbotapi.NewMessage(chatid, text)
                                        bot.Send(msg)
                                }
                        }
                }

                if mailhost != "" && complete_text != "\r\nMessage Body:\n" {
                        text := "Subject: Open Port Notifier - Report" + "\r\n " + complete_text
                        log.Println("Sending email report")
                        SendMail(mailsrc, maildst, mailhost, text)
                }

                if err := scanner.Err(); err != nil {
                        log.Fatal(err)
                }
                i++
        }
}
