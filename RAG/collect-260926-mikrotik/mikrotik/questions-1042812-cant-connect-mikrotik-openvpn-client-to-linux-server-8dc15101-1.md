---
id: collect-260926-mikrotik/mikrotik/questions-1042812-cant-connect-mikrotik-openvpn-client-to-linux-server-8dc15101-1
title: "ip addr"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2020-11-17"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-1042812-cant-connect-mikrotik-openvpn-client-to-linux-server-8dc15101.md
source_anchor: ""
source_lines: [1, 134]
sha256: d94410c3fc2e57aa719b267d8ebb19ae87149c6a2729760a74925f0e7e06144c
---

# ip addr

Greethigs!
I would really appreciate any help!
For the last two days i cannot make Mikrotik router connect to Debian 10 OpenVPN server.
For each connection attempt router shows the following error message:
Status: terminating... - could not add address: netmask cannot be /0 (6)
Here's my config:
Client: Mikrotik RB2011UiAS-2HnD-IN / RouterOS: 7.1b2 / 6.46.8 (tried both firmwares)
"White" IP: 195.111.111.2/30    Gateway: 195.111.111.1
[admin@MikroTik] > /interface print                                                                                                                                     
Flags: R - RUNNING; S - SLAVE      
   #      NAME            TYPE      ACTU  L2MT  MAX-  MAC-ADDRESS          
   0  R   ether1-gateway  ether     1500  1598  4074  D4:CA:6D:00:E0:03    
   1  RS  ether2          ether     1500  1598  4074  D4:CA:6D:00:E0:04    
   2  RS  ether3          ether     1500  1598  4074  D4:CA:6D:00:E0:05    
   3  RS  ether4          ether     1500  1598  4074  D4:CA:6D:00:E0:06    
   4   S  ether5          ether     1500  1598  4074  D4:CA:6D:00:E0:07    
   5   S  ether6          ether     1500  1598  2028  D4:CA:6D:00:E0:08    
   6   S  ether7          ether     1500  1598  2028  D4:CA:6D:00:E0:09    
   7   S  ether8          ether     1500  1598  2028  D4:CA:6D:00:E0:0A    
   8   S  ether9          ether     1500  1598  2028  D4:CA:6D:00:E0:0B    
   9   S  ether10         ether     1500  1598  2028  D4:CA:6D:00:E0:0C  
  10      sfp1            ether     1500  1598  4074  D4:CA:6D:00:E0:02    
  11  RS  wlan1           wlan      1500  1600  2290  D4:CA:6D:00:E0:0D    
  12      office          ovpn-out                    02:FF:F4:DF:C3:9A    
  13  R   bridge1         bridge    1500  1598        D4:CA:6D:00:E0:0D
[admin@MikroTik] > /interface ovpn-client print                                                                                             
Flags: X - disabled; R - running     
 0    name="office" mac-address=02:F2:F2:2F:CC:1A max-mtu=1500 connect-to=195.222.222.2 port=1198 mode=ip protocol=tcp user="mikrotik" password="" profile=default     
      certificate=cert_2 verify-server-certificate=no auth=sha1 cipher=aes256 use-peer-dns=no add-default-route=no 
[admin@MikroTik] > /ip route print                                                                                                          
Flags: D - DYNAMIC; X - DISABLED, I - INACTIVE, A - ACTIVE; C - CONNECT, S - STATIC, m - MODEM        
  #       DST-ADDRESS        GATEWAY         D    
     AS   0.0.0.0/0          195.111.111.1  1    
     DAC  192.168.47.0/24    bridge1         0    
     DAC  195.111.111.2/30   ether1-gateway  0
[admin@MikroTik] > /ip firewall nat print     
Flags: X - disabled, I - invalid; D - dynamic     
 0    chain=srcnat action=masquerade dst-address=0.0.0.0 out-interface=ether1-gateway log=no log-prefix="" 
Server: Debian 10 + Xen 4.11 bridged
# ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: eno1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq master xenbr0 state UP group default qlen 1000
    link/ether 2c:2f:6b:20:2e:24 brd ff:ff:ff:ff:ff:ff
3: eno2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    link/ether 2c:2f:2b:20:2e:25 brd ff:ff:ff:ff:ff:ff
    inet 195.222.222.2/30 brd 195.222.222.1 scope global eno2
       valid_lft forever preferred_lft forever
4: xenbr0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 2c:2f:2b:20:2e:24 brd ff:ff:ff:ff:ff:ff
    inet 192.168.48.1/24 brd 192.168.48.255 scope global xenbr0
       valid_lft forever preferred_lft forever
5: vif1.0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq master xenbr0 state UP group default qlen 32
    link/ether fe:ff:ff:ff:ff:ff brd ff:ff:ff:ff:ff:ff
6: tun1: <POINTOPOINT,MULTICAST,NOARP,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UNKNOWN group default qlen 500
    link/none
    inet 10.101.0.1/24 scope global tun1
       valid_lft forever preferred_lft forever
# cat /etc/openvpn/server.conf:
local 195.222.222.2
port 1198
proto tcp
dev tun1
ca /etc/openvpn/keys/ca.crt
cert /etc/openvpn/keys/server.crt
key /etc/openvpn/keys/server.key  
dh /etc/openvpn/keys/dh.pem
data-ciphers AES-256-CBC
cipher AES-256-CBC
auth SHA1
topology subnet
server 10.101.0.0 255.255.255.0
ifconfig-pool-persist ipp_tcp.txt
client-config-dir ccd
push "route 192.168.48.0 255.255.255.0"
push "route 192.168.47.0 255.255.255.0"
route 192.168.47.0/24 255.255.255.0
client-to-client
keepalive 10 120
;comp-lzo
;max-clients 100
user nobody
group nogroup
persist-key
persist-tun
status  /var/log/openvpn/openvpn-status.log
log     /var/log/openvpn/openvpn.log
verb 6
;mute 20
#/etc/openvpn/ccd/mikrotik
iroute 192.168.47.0 255.255.255.0 10.101.0.2
ifconfig-push 10.101.0.2 10.101.0.1 255.255.255.252
# iptables -S
-P INPUT DROP
-P FORWARD ACCEPT
-P OUTPUT ACCEPT
-A INPUT -p tcp -m tcp --dport 1198 -m state --state NEW -j ACCEPT
-A INPUT -i lo -j ACCEPT
-A INPUT -i xenbr0 -j ACCEPT
-A INPUT -m state --state RELATED,ESTABLISHED -j ACCEPT
-A INPUT -p icmp -m icmp --icmp-type 8 -j ACCEPT
-A INPUT -m state --state INVALID -j DROP
-A OUTPUT -o lo -j ACCEPT
-A OUTPUT -o xenbr0 -j ACCEPT
-A OUTPUT -m state --state RELATED,ESTABLISHED -j ACCEPT
# iptables -t nat -S
-P PREROUTING ACCEPT
-P INPUT ACCEPT
-P POSTROUTING ACCEPT
-P OUTPUT ACCEPT
-A POSTROUTING -s 192.168.48.0/24 -d 192.168.48.0/24 -j ACCEPT
-A POSTROUTING -s 192.168.48.0/24 -d 10.0.0.0/8 -j ACCEPT
-A POSTROUTING -s 192.168.48.0/24 -j SNAT --to-source 195.222.222.2
OpenVPN server log:
2020-11-17 01:53:40 us=127869 MULTI: multi_create_instance called
2020-11-17 01:53:40 us=127918 Re-using SSL/TLS context
2020-11-17 01:53:40 us=127965 Control Channel MTU parms [ L:1623 D:1210 EF:40 EB:0 ET:0 EL:3 ]
2020-11-17 01:53:40 us=127977 Data Channel MTU parms [ L:1623 D:1450 EF:123 EB:406 ET:0 EL:3 ]
2020-11-17 01:53:40 us=128002 Local Options String (VER=V4): 'V4,dev-type tun,link-mtu 1559,tun-mtu 1500,proto TCPv4_SERVER,cipher AES-256-CBC,auth SHA1,key$
2020-11-17 01:53:40 us=128010 Expected Remote Options String (VER=V4): 'V4,dev-type tun,link-mtu 1559,tun-mtu 1500,proto TCPv4_CLIENT,cipher AES-256-CBC,aut$
2020-11-17 01:53:40 us=128035 TCP connection established with [AF_INET]195.111.111.2:34720
2020-11-17 01:53:40 us=128045 TCPv4_SERVER link local: (not bound)
2020-11-17 01:53:40 us=128053 TCPv4_SERVER link remote: [AF_INET]195.111.111.2:34720
2020-11-17 01:53:40 us=132062 195.111.111.2:34720 TCPv4_SERVER READ [14] from [AF_INET]195.111.111.2:34720: P_CONTROL_HARD_RESET_CLIENT_V2 kid=0 [ ] pid=0$
2020-11-17 01:53:40 us=132100 195.111.111.2:34720 TLS: Initial packet from [AF_INET]195.111.111.2:34720, sid=d0dc4e5c 860c785f
2020-11-17 01:53:40 us=132124 195.111.111.2:34720 TCPv4_SERVER WRITE [26] to [AF_INET]195.111.111.2:34720: P_CONTROL_HARD_RESET_SERVER_V2 kid=0 [ 0 ] pid=$
2020-11-17 01:53:40 us=136840 195.111.111.2:34720 TCPv4_SERVER READ [22] from [AF_INET]195.111.111.2:34720: P_ACK_V1 kid=0 [ 0 ]
2020-11-17 01:53:40 us=183931 195.111.111.2:34720 TCPv4_SERVER READ [116] from [AF_INET]195.111.111.2:34720: P_CONTROL_V1 kid=0 [ ] pid=1 DATA len=102
2020-11-17 01:53:40 us=184699 195.111.111.2:34720 TCPv4_SERVER WRITE [1196] to [AF_INET]195.111.111.2:34720: P_CONTROL_V1 kid=0 [ 1 ] pid=1 DATA len=1170
2020-11-17 01:53:40 us=184741 195.111.111.2:34720 TCPv4_SERVER WRITE [1078] to [AF_INET]195.111.111.2:34720: P_CONTROL_V1 kid=0 [ ] pid=2 DATA len=1064
2020-11-17 01:53:40 us=190449 195.111.111.2:34720 TCPv4_SERVER READ [22] from [AF_INET]195.111.111.2:34720: P_ACK_V1 kid=0 [ 1 ]
2020-11-17 01:53:40 us=235548 195.111.111.2:34720 TCPv4_SERVER READ [22] from [AF_INET]195.111.111.2:34720: P_ACK_V1 kid=0 [ 2 ]
