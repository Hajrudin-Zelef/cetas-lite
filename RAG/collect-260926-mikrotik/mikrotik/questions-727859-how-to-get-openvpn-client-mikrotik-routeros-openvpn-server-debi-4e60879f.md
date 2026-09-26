---
id: collect-260926-mikrotik/mikrotik/questions-727859-how-to-get-openvpn-client-mikrotik-routeros-openvpn-server-debi-4e60879f
title: "cat /etc/openvpn/server.conf"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-727859-how-to-get-openvpn-client-mikrotik-routeros-openvpn-server-debi-4e60879f.md
source_anchor: ""
source_lines: [1, 145]
sha256: 8648b415dbb70817204f6cfe953bd64ca207809be12417ecd35d34f0f13ce007
---

# cat /etc/openvpn/server.conf

I have some issues with making MT to work with OpenVPN server (Debian). I can make successfull connection to OVPN server, but traffic is not routed through OVPN server. Here is my configuration.
Setup - https://i.sstatic.net/AAH9Y.jpg
OpenVPN server (Debian/Linux) configuration
# cat /etc/openvpn/server.conf
local 95.2.171.3
port 1194
proto tcp
dev tun
ca ca.crt
cert server.crt
key server.key
dh dh.pem
server 10.8.0.0 255.255.255.0
ifconfig-pool-persist ipp.txt
client-config-dir ccd
route 192.168.81.0/24 255.255.255.0
keepalive 10 120
tun-mtu 1500
mssfix 1450
cipher AES-256-CBC
auth sha1
persist-key
persist-tun
status /var/log/openvpn-status.log
log-append /var/log/openvpn.log
verb 5
crl-verify /etc/openvpn/easy-rsa/pki/crl.pem
# cat /etc/openvpn/ccd/client
iroute 192.168.81.0 255.255.255.0 10.8.0.2
ifconfig-push 10.8.0.2 10.8.0.1
# cat /proc/sys/net/ipv4/ip_forward
1
# netstat -an | grep 1194
tcp        0      0 95.2.171.3:1194       0.0.0.0:*               LISTEN
tcp        0      0 95.2.171.3:1194       81.190.190.100:62973    ESTABLISHED
# ifconfig
eth0   Link encap:Ethernet  HWaddr 20:cf:30:f2:a8:76
          inet addr:95.2.171.3  Bcast:95.2.171.31  Mask:255.255.255.224
          inet6 addr: fe80::22cf:30ff:fef2:a876/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:255189 errors:0 dropped:0 overruns:0 frame:0
          TX packets:333054 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000
          RX bytes:34521411 (32.9 MiB)  TX bytes:367074147 (350.0 MiB)
          Interrupt:26 Base address:0x8000
lo       Link encap:Local Loopback
          inet addr:127.0.0.1  Mask:255.0.0.0
          inet6 addr: ::1/128 Scope:Host
          UP LOOPBACK RUNNING  MTU:16436  Metric:1
          RX packets:15579 errors:0 dropped:0 overruns:0 frame:0
          TX packets:15579 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:1326071 (1.2 MiB)  TX bytes:1326071 (1.2 MiB)
tun0   Link encap:UNSPEC  HWaddr 00-00-00-00-00-00-00-00-00-00-00-00-00-00-00-00
          inet addr:10.8.0.1  P-t-P:10.8.0.2  Mask:255.255.255.255
          UP POINTOPOINT RUNNING NOARP MULTICAST  MTU:1500  Metric:1
          RX packets:57 errors:0 dropped:0 overruns:0 frame:0
          TX packets:6 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:100
          RX bytes:6669 (6.5 KiB)  TX bytes:504 (504.0 B)
# netstat -rn
Kernel IP routing table
Destination     Gateway         Genmask         Flags   MSS Window  irtt Iface
10.8.0.2        0.0.0.0         255.255.255.255 UH        0 0          0 tun0
95.2.171.0    0.0.0.0         255.255.255.224 U         0 0          0 eth0
192.168.81.0    10.8.0.2        255.255.255.0   UG        0 0          0 tun0
10.8.0.0        10.8.0.2        255.255.255.0   UG        0 0          0 tun0
0.0.0.0         95.2.171.30   0.0.0.0         UG        0 0          0 eth0
# iptables -S
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT
-A INPUT -i lo -j ACCEPT
-A INPUT -d 127.0.0.0/8 -i !lo -j REJECT --reject-with icmp-port-unreachable
-A INPUT -i tun0 -j ACCEPT
-A INPUT -m state --state RELATED,ESTABLISHED -j ACCEPT
-A INPUT -p tcp -m tcp --dport 1194 -j ACCEPT
-A INPUT -m limit --limit 5/min -j LOG --log-prefix "iptables denied: " --log-level 7
-A INPUT -j REJECT --reject-with icmp-port-unreachable
-A OUTPUT -j ACCEPT
# iptables -t nat -S
-P PREROUTING ACCEPT
-P POSTROUTING ACCEPT
-P OUTPUT ACCEPT
-A POSTROUTING -s 10.8.0.0/24 -j SNAT --to-source 95.2.171.3
-A POSTROUTING -s 10.8.0.0/24 -j SNAT --to-source 95.2.171.3
-A POSTROUTING -s 10.8.0.0/24 -j SNAT --to-source 95.2.171.3
# ping 8.8.8.8
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_req=1 ttl=55 time=12.9 ms
64 bytes from 8.8.8.8: icmp_req=2 ttl=55 time=12.8 ms
This is all my config on OpenVPN Server (Debian/Linux).
OpenVPN Client side (Mikrotik RouterOS 6) configuration
/interface print
Flags: D - dynamic, X - disabled, R - running, S - slave
 #     NAME                                TYPE       ACTUAL-MTU L2MTU  MAX-L2MTU MAC-ADDRESS
 0  R  ether1                              ether            1500  1600       4076 D4:CA:6D:31:14:F4
 1   S ether2                              ether            1500  1598       2028 D4:CA:6D:31:14:F5
 2   S ether3                              ether            1500  1598       2028 D4:CA:6D:31:14:F6
 3   S ether4                              ether            1500  1598       2028 D4:CA:6D:31:14:F7
 4   S ether5                              ether            1500  1598       2028 D4:CA:6D:31:14:F8
 5  RS wlan1                               wlan             1500  1600            D4:CA:6D:31:14:F9
 6  R  bridge1                             bridge           1500  1598            D4:CA:6D:31:14:F5
 7  R  ovpn-out1                           ovpn-out         1500                  FE:3E:27:7D:61:8C
 /interface bridge print
Flags: X - disabled, R - running
 0  R name="bridge1" mtu=auto actual-mtu=1500 l2mtu=1598 arp=enabled mac-address=D4:CA:6D:31:14:F5 protocol-mode=rstp priority=0x8000 auto-mac=yes admin-mac=00:00:00:00:00:00 max-message-age=20s forward-delay=15s transmit-hold-count=6 ageing-time=5m
/interface bridge port print
Flags: X - disabled, I - inactive, D - dynamic
 #    INTERFACE    BRIDGE        PRIORITY  PATH-COST    HORIZON
 0 I  ether2       bridge1                       0x80         10       none
 1 I  ether3       bridge1                       0x80         10       none
 2 I  ether4       bridge1                       0x80         10       none
 3 I  ether5       bridge1                       0x80         10       none
 4    wlan1        bridge1                       0x80         10       none
 /ip address print
Flags: X - disabled, I - invalid, D - dynamic
 #   ADDRESS            NETWORK         INTERFACE
 0   192.168.81.1/24    192.168.81.0    bridge1
 1 D 192.168.7.200/24   192.168.7.0     ether1
 2 D 10.8.0.2/32        10.8.0.1        ovpn-out1
 /ip firewall nat print
Flags: X - disabled, I - invalid, D - dynamic
 0    chain=srcnat action=masquerade to-addresses=0.0.0.0 out-interface=ether1 log=no log-prefix=""
 /ip route print
Flags: X - disabled, A - active, D - dynamic, C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme, B - blackhole, U - unreachable, P - prohibit
 #      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
 0 ADS  0.0.0.0/0                          192.168.7.1               0
 1 ADC  10.8.0.1/32        10.8.0.2        ovpn-out1                 0
 2 ADC  192.168.7.0/24     192.168.7.200   ether1                    0
 3 ADC  192.168.81.0/24    192.168.81.1    bridge1                   0
 /interface ovpn-client print
Flags: X - disabled, R - running
 0  R name="ovpn-out1" mac-address=FE:3E:27:7D:61:8C max-mtu=1500 connect-to=195.13.171.3 port=1194 mode=ip user="client" password="" profile=default certificate=Client auth=sha1 cipher=aes256 add-default-route=no
 /ping 10.8.0.1
  SEQ HOST                                     SIZE TTL TIME  STATUS
    0 10.8.0.1                                   56  64 6ms
    1 10.8.0.1                                   56  64 9ms
    2 10.8.0.1                                   56  64 7ms
    3 10.8.0.1                                   56  64 6ms
    sent=4 received=4 packet-loss=0% min-rtt=6ms avg-rtt=7ms max-rtt=9ms
As you can see, I can ping OpenVPN server from Mikrotik. But when I use internet from Local PC it shows 81.190.190.100 IP address, not the one I would like to see - OpenVPN servers IP - 95.2.171.3.
I can successfully ping/traceroute to 10.8.0.1 from Laptop (192.168.81.100/24), but cannot understand why it's not routed through VPN tunnel. I think I'm missing something with routing either on server (Linux) or client (mikrotik).
Thanks for your help! I'm playing whit this for a while now, and cannot get it running :(
Have a good day!
