---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-final-laboratory-md-0b92a2ff-3
title: "Create VLANs"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-final-laboratory-md-0b92a2ff.md
source_anchor: ""
source_lines: [592, 665]
sha256: ce1dbbea7f958ff03e27e09236cdfb5e73196149ec6bc4fab263ef3b2f6fb7c5
---

# Create VLANs

student@ubuntu:~$ sudo chmod -R 755 /srv/tftp
student@ubuntu:~$ ls -ld /srv/tftp
drwxr-xr-x 2 tftp nogroup /srv/tftp
Step3: Configure the Firewall
student@ubuntu:~$ sudo ufw enable
student@ubuntu:~$ sudo ufw allow from 172.16.128.0/24 to any port 69 proto udp
student@ubuntu:~$ sudo ufw deny 69/udp
student@ubuntu:~$ sudo ufw reload
student@ubuntu:~$ sudo ufw statusstudent@ubuntu:~$ ss -tulpna
немесе
student@ubuntu:~$ netstat -tulpna
Step4: Testing the TFTP Server
Download and Upload files
get - Download file from TFTP server
put - Upload file from TFTP server
student@ubuntu:~$ sudo touch /srv/tftp/f1.conf
student@ubuntu:~$ tftp 172.16.128.69 -c get f1.conf
student@ubuntu:~$ pwd
/home/student
student@ubuntu:~$ ls -l
-rw-rw-r-- 1 student student f1.conf
Huawei VRP Router/Switch
# Ping from EdgeR1 to TFTP Server
<EdgeR1> ping 172.16.128.69
  Reply from 172.16.128.69: bytes=56 Sequence=2 ttl=64 time=10 ms# Download file from TFTP server
tftp <tftp-server-ip> get <remote-file>
<EdgeR1> tftp 172.16.128.69 get f1.conf
TFTP: Downloading the file successfully
<EdgeR1> dir
  Idx  Attr     Size(Byte)  Date        Time(LMT)  FileName 
    0  -rw-              0  May 03 2026 11:37:46   f1.conf
tftp 172.16.128.10 get f1.conf
tftp 172.16.128.10 get f1.conf f11.cfg
# Upload file from TFTP server
tftp <tftp-server-ip> put <local-file>
<EdgeR1> tftp 172.16.128.69 put vrpcfg.zip
TFTP: Uploading the file successfully
student@ubuntu:~$ ls -lh /srv/tftp/
-rw-r--r-- 1 root root f1.conf
-rw-rw-rw- 1 tftp tftp vrpcfg.zip
Table - WLAN Data Plan
| Item | Value | 
|---|---|
| Management VLAN for APs | VLAN 43 | 
| Service VLAN for STAs | SSID Staff: VLAN 100, SSID Guest: VLAN 200 | 
| Default Gateway for AP | 10.1.43.254 | 
| DHCP Pool for APs | 10.1.43.101 - 10.1.43.200/24 | 
| Default Gateway for Staff | 192.168.100.254 | 
| DHCP Pool for Staff | 192.168.100.11 - 192.168.100.250/24 | 
| Default Gateway for Guest | 192.168.200.254 | 
| DHCP Pool for Guest | 192.168.200.11 - 192.168.200.250/24 | 
| AP Name | AP1, AP2 | 
| AP Group | Name: ap-group1 | 
|  | Referenced profiles: VAP profile VAP-Staff and VAP-Guest, Regulatory domain profile default | 
| Regulatory Domain Profile | Name: default | 
|  | Country code: KZ | 
| SSID Profile (Staff) | Name: WLAN-Staff | 
|  | SSID name: Staff-WiFi | 
| Security Profile (Staff) | Name: WLAN-Staff | 
|  | Security policy: WPA-WPA2+PSK+AES | 
|  | Password: Huawei@123 | 
| SSID Profile (Guest) | Name: WLAN-Guest | 
|  | SSID name: Guest-WiFi | 
| Security Profile (Guest) | Name: WLAN-Guest | 
|  | Security policy: WPA-WPA2+PSK+AES | 
|  | Password: Huawei@123 | 
| VAP Profile (Staff) | Name: VAP-Staff | 
|  | Forwarding mode: Direct forwarding | 
|  | Service VLAN: 100 | 
|  | Referenced profiles: SSID profile WLAN-Staff and Security profile WLAN-Staff | 
| VAP Profile (Guest) | Name: VAP-Guest | 
|  | Forwarding mode: Direct forwarding | 
|  | Service VLAN: 200 | 
|  | Referenced profiles: SSID profile WLAN-Guest and Security profile WLAN-Guest |
