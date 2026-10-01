---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab10-networkservices-md-f5b58b9b-2
title: "Create VLANs"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab10-networkservices-md-f5b58b9b.md
source_anchor: ""
source_lines: [322, 586]
sha256: b5a10f0f2ec9c3e12397f719e7df280b6b197289cd937067df4c898cd9fab7c3
---

# Create VLANs

 Client Request          : 8
  Dhcp Discover          : 4
  Dhcp Request           : 4
  Dhcp Decline           : 0
  Dhcp Release           : 0
  Dhcp Inform            : 0
  Server Reply           : 8
  Dhcp Offer             : 4
  Dhcp Ack               : 4
  Dhcp Nak               : 0
  Bad Messages           : 0 PC1> ipconfig
PC2> ipconfig
PC3> ipconfig
PC4> ipconfig /renew
EdgeR1
ping 192.168.137.1
 Reply from 192.168.137.1: bytes=56 Sequence=2 ttl=128 time=10 ms
ping 8.8.8.8
 Request time out
Default Static Route
ip route-static 0.0.0.0 0.0.0.0 192.168.137.1
display cu | include staticping 8.8.8.8
 Reply from 8.8.8.8: bytes=56 Sequence=1 ttl=107 time=90 ms
Advertise the Default Route
ospf 1
 default-route-advertise
 quit<EdgeR1> display ip routing-table
Destination/Mask   Proto   Pre   Cost   Flags   NextHop         Interface
       0.0.0.0/0   Static  60    0      RD      192.168.137.1   GigabitEthernet 0/0/1<C1> display ip routing-table
Destination/Mask   Proto   Pre   Cost   Flags   NextHop         Interface
       0.0.0.0/0   O_ASE   150   1      D       10.1.1.101      GigabitEthernet 0/0/0
<C1> display ospf routing
 Routing for ASEs
 Destination        Cost      Type       Tag         NextHop         AdvRouter
 0.0.0.0/0          1         Type2      1           10.1.1.101      50.1.1.1acl 2000
 rule permit source 172.16.111.0 0.0.0.255
 rule permit source 172.16.112.0 0.0.0.255
 quit
int g0/0/1
 nat outbound 2000
 quit
Verify Configuration
display cu section acl
display nat outboundPC1> ping 8.8.8.8
PC2> ping 8.8.8.8
PC3> ping 8.8.8.8
PC4> ping 8.8.8.8
 From 8.8.8.8: bytes=32 seq=3 ttl=105 time=156 msPC1> ping google.com
PC3> ping google.com
 From 142.250.181.238: bytes=32 seq=1 ttl=106 time=156 ms
NAT Table
[EdgeR1] display nat session all verbose
A1 Switch
# Create VLANIF interface
interface vlanif 50
 ip address 10.1.50.101 24
 quit
display ip int brief# Default Gateway
ip route-static 0.0.0.0 0.0.0.0 10.1.50.254
A2 Switch
# Create VLANIF interface
interface vlanif 50
 ip address 10.1.50.102 24
 quit
display ip int brief# Default Gateway
ip route-static 0.0.0.0 0.0.0.0 10.1.50.254
A1 Switch
ping 10.1.50.102
 Reply from 10.1.50.102: bytes=56 Sequence=4 ttl=255 time=30 ms
ping 50.3.3.3
 Reply from 50.3.3.3: bytes=56 Sequence=3 ttl=255 time=20 ms
ping 50.4.4.4
 Reply from 50.4.4.4: bytes=56 Sequence=5 ttl=255 time=60 ms
ping 50.2.2.2
 Reply from 50.2.2.2: bytes=56 Sequence=1 ttl=254 time=60 ms
ping 50.1.1.1
 Reply from 50.1.1.1: bytes=56 Sequence=3 ttl=253 time=70 ms
ping 50.5.5.5
 Reply from 50.5.5.5: bytes=56 Sequence=3 ttl=253 time=70 ms
Step1: Enable SSH/Telnet
stelnet server enable
display ssh server statusdisplay telnet server status
telnet server enable
Step2: Generate RSA Key
rsa local-key-pair create
Warning: Confirm to replace them! Continue? [Y/N] Y
Input the bits in the modulus[default = 1024]: 2048
display rsa local-key-pair public
Step3: Configure Local User Authentication and Authorization
aaa
 local-user student password cipher Huawei@123
 local-user student privilege level 15
 local-user student service-type terminal ssh telnet
 quit
Step4: Configure SSH User Settings
ssh user student authentication-type password
ssh user student service-type stelnet
Step5: Configure VTY Lines
user-interface vty 0 4
 authentication-mode aaa
 protocol inbound all
 quitdisplay cu | include ssh
Verify SSH Connectivity
# Configure SSH Client Settings
[A1] ssh client first-time enable
[A1] stelnet 50.3.3.3
Please input the username: student
The server is not authenticated. Continue to access it? (y/n)[n]: y
Save the server's public key? (y/n)[n]: y
Enter password: Huawei@123
<D1> system-view
[D1] quit
<D1> quit
[A1][A1] stelnet 50.4.4.4
[A1] stelnet 50.2.2.2
[A1] stelnet 50.1.1.1
[A1] stelnet 50.5.5.5
[A1] stelnet 10.1.50.102
Verify Telnet Connectivity
<A1> telnet 50.3.3.3
 Username: student
 Password: Huawei@123
<D1> system-view
[D1] quit
<D1> quit
<A1> <A1> telnet 50.4.4.4
<A1> telnet 50.2.2.2
<A1> telnet 50.1.1.1
<A1> telnet 50.5.5.5
<A1> telnet 10.1.50.102
DNS Server
Basic Config:
 Local Address: 172.16.128.53
 Subnet Mask: 255.255.255.0
 Gateway: 172.16.128.1
 DNS: 8.8.8.8                        // Public DNS Server
Server info:
 Hostname: lab.local
 IP Address: 172.16.128.80          // Web Server
 "Add" батырмасын басамыз!
 DNSServer ➜ Service ➜ Start
HTTP Server
Basic Config:
 Local Address: 172.16.128.80
 Subnet Mask: 255.255.255.0
 Gateway: 172.16.128.1
 DNS: 172.16.128.53                 // Local DNS Server
Server info:
 Root Path: C:\Users\student\Documents\www\
 HTTPServer ➜ Service ➜ Start
C:\Users\student\Documents\www\index.html
<!DOCTYPE html>
<html>
<head>
   	 <meta charset="UTF-8">
   	 <title>Example</title>
</head>
<body>
   	 <h1>Welcome to Almaty!</h1>
</body>
</html>
HTTP Client
Basic Config:
 Local Address: 172.16.111.10
 Subnet Mask: 255.255.255.0
 Gateway: 172.16.111.254
 DNS: 172.16.128.53
 HTTPClient ➜ URL: http://lab.local
немесе
 HTTPClient ➜ URL: http://172.16.128.80
Нәтиже:
HTTP/1.1 200 OK
Server: ENSP HttpServer
Auth: HUAWEI
Cache-Control: private
Content-Type: text/html
Content-Length: 179
FTP Server
Basic Config:
 Local Address: 172.16.128.21
 Subnet Mask: 255.255.255.0
 Gateway: 172.16.128.1
 DNS: 172.16.128.53
Server info:
 Root Path: C:\Users\student\Documents\tftpboot\file1.txt
 FtpServer ➜ Service ➜ Start
Download and Upload files
get - Download file from FTP server
put - Upload file from FTP server
Example #1: Download file from FTP server
<EdgeR1> ftp 172.16.128.21
User(172.16.128.21:(none)): ENTER
Enter password: ENTER
[EdgeR1-ftp]
[EdgeR1-ftp] ?
[EdgeR1-ftp] dir
-rwxrwxrwx  1  nogroup  0 May 3  2026  file1.txt
[EdgeR1-ftp] get file1.txt
226 Transfer finished successfully. Data connection closed.
[EdgeR1-ftp] bye
<EdgeR1> dir
Idx  Attr     Size(Byte)  Date        Time(LMT)  FileName 
  6  -rw-              0  May 03 2026 03:53:41   file1.txt
Example #2: Upload file from FTP server
<EdgeR1> save
Are you sure to continue? (y/n)[n]: y
<EdgeR1> dir
  Idx  Attr     Size(Byte)  Date        Time(LMT)  FileName 
    8  -rw-            864  May 03 2026 03:58:19   vrpcfg.zip
<EdgeR1> ftp 172.16.128.21
User(172.16.128.21:(none)): ENTER
Enter password: ENTER
[EdgeR1-ftp]
[EdgeR1-ftp] put vrpcfg.zip
226 Transfer finished successfully. Data connection closed.
[EdgeR1-ftp] dir
-rwxrwxrwx  1  nogroup  0 May 3  2026  file1.txt
-rwxrwxrwx  1  nogroup  864 May 3  2026  vrpcfg.zip
About the System
student@ubuntu:~$ uname -rs
Linux 6.8.0-101-generic x86_64 GNU/Linux
student@ubuntu:~$ lsb_release -a
Ubuntu 24.04.4 LTS
Codename: noble
Желілік интерфейсті конфигурациялау
student@ubuntu:~$ sudo nano /etc/netplan/50-cloud-init.yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    ens32:
      dhcp4: true
    ens34:
      dhcp4: false
      addresses:
        - 172.16.128.69/24
CTRL+O, ENTER, CTRL+Xstudent@ubuntu:~$ sudo netplan applystudent@ubuntu:~$ ip address
Step1: installation of TFTP Server
Package атауы: tftpd-hpa
Daemon/Service атауы: tftpd-hpa
student@ubuntu:~$ sudo apt update
student@ubuntu:~$ sudo apt install -y tftpd-hpa tftp-hpa
tftpd-hpa – HPA's TFTP Server
tftp-hpa – HPA's TFTP Client
Status the tftpd-hpa Service/Daemon
student@ubuntu:~$ sudo systemctl status tftpd-hpa
active (running)
student@ubuntu:~$ sudo systemctl is-enabled tftpd-hpa
enabled
Step2: Configure the TFTP Server
Edit tftpd-hpa Configuration File
student@ubuntu:~$ sudo nano /etc/default/tftpd-hpa
TFTP_USERNAME="tftp"
TFTP_DIRECTORY="/srv/tftp"
TFTP_ADDRESS="172.16.128.69:69"
TFTP_OPTIONS="--secure --create --listen --verbose"
CTRL+O, ENTER, CTRL+X
Restart the tftpd-hpa Service/Daemon
student@ubuntu:~$ sudo systemctl restart tftpd-hpa
Modify Permission/Ownership on TFTP Root Directory
student@ubuntu:~$ ls -ld /srv/tftp
drwxr-xr-x 2 root nogroup /srv/tftp
Modify Ownership
student@ubuntu:~$ sudo chown -R tftp /srv/tftp
Modify Permission
