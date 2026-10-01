---
id: collect-261001-meraki/meraki/questions-810327-connecting-to-meraki-client-vpn-from-ubuntu-w-strongswan-and-xl-d9bb8195
title: "ipsec.conf - strongSwan IPsec configuration file"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-meraki/questions-810327-connecting-to-meraki-client-vpn-from-ubuntu-w-strongswan-and-xl-d9bb8195.md
source_anchor: ""
source_lines: [1, 85]
sha256: 007a6fb08b5cde5c6bbdc8d5f3208f334fc249277046e19468bbc87aa3419f7d
---

# ipsec.conf - strongSwan IPsec configuration file

I have a virtual appliance running Ubuntu 14 that I would like to connect to a remote network via Meraki Client VPN.
I have found instructions for installing and configuring strongswan and xl2tpd on Ubuntu 16 which get me most of the way there, but after establishing the VPN I cannot seem to establish a route.
The goal is allow traffic from this single Ubuntu VM to a subnet on the remote LAN via the VPN.
(Ubuntu 192.168.5.99/32) =VPN> (Meraki appliance 20.20.20.20) => (Subnet 192.168.1.0/24)
https://gist.github.com/psanford/42c550a1a6ad3cb70b13e4aaa94ddb1c http://www.jasonernst.com/2016/06/21/l2tp-ipsec-vpn-on-ubuntu-16-04/
/etc/ipsec.conf:
# ipsec.conf - strongSwan IPsec configuration file
# basic configuration
config setup
        # strictcrlpolicy=yes
        # uniqueids = no
# Add connections here.
# Sample VPN connections
conn %default
        ikelifetime=60m
        keylife=20m
        rekeymargin=3m
        keyingtries=1
        keyexchange=ikev1
        authby=secret
        ike=aes128-sha1-modp1024,3des-sha1-modp1024!
        esp=aes128-sha1-modp1024,3des-sha1-modp1024!
conn meraki-vpn
     keyexchange=ikev1
     left=%defaultroute
     auto=add
     authby=secret
     type=transport
     leftprotoport=17/1701
     rightprotoport=17/1701
     # set this to the ip address of your meraki vpn
     right=20.20.20.20
/etc/xl2tpd/xl2tpd.conf :
[lac meraki]
lns = 20.20.20.20
ppp debug = yes
pppoptfile = /etc/ppp/options.l2tpd.client
length bit = yes
I can bring up the VPN:
root@zzz:~# ipsec up meraki-vpn
...
connection 'meraki-vpn' established successfully
Login:
echo "c meraki me@mydomain.com mypassword" > /var/run/xl2tpd/l2tp-control
But I cannot get a route setup. Attempting to follow the magic recipes result errors. I can't figure if this is some xl2tpd config problem, or if I just can't figure the correct syntax to add the route. Any insight would be appreciated.
root@zzz:~# ip route add 192.168.1.0/24 dev ppp0
Cannot find device "ppp0"
ipsec statusall :
Status of IKE charon daemon (strongSwan 5.1.2, Linux 4.2.0-42-generic, x86_64):
  uptime: 20 hours, since Oct 19 19:30:38 2016
  malloc: sbrk 2433024, mmap 0, used 352240, free 2080784
  worker threads: 11 of 16 idle, 5/0/0/0 working, job queue: 0/0/0/0, scheduled: 4
  loaded plugins: charon test-vectors aes rc2 sha1 sha2 md4 md5 rdrand random nonce x509 revocation constraints pkcs1 pkcs7 pkcs8 pkcs12 pem openssl xcbc cmac hmac ctr ccm gcm attr kernel-netlink resolve socket-default stroke updown eap-identity addrblock
Listening IP addresses:
  192.168.5.99
Connections:
  meraki-vpn:  %any...20.20.20.20  IKEv1
  meraki-vpn:   local:  [192.168.5.99] uses pre-shared key authentication
  meraki-vpn:   remote: [20.20.20.20] uses pre-shared key authentication
  meraki-vpn:   child:  dynamic[udp/l2f] === dynamic[udp/l2f] TRANSPORT
Security Associations (1 up, 0 connecting):
  meraki-vpn[4]: ESTABLISHED 8 minutes ago, 192.168.5.99[192.168.5.99]...20.20.20.20[20.20.20.20]
  meraki-vpn[4]: IKEv1 SPIs: 2121ded4dc4dc94e_i* 3d5d1cc8bf935f19_r, pre-shared key reauthentication in 48 minutes
  meraki-vpn[4]: IKE proposal: 3DES_CBC/HMAC_SHA1_96/PRF_HMAC_SHA1/MODP_1024
  meraki-vpn{5}:  INSTALLED, TRANSPORT, ESP in UDP SPIs: c330d0b3_i 06749ffa_o
  meraki-vpn{5}:  3DES_CBC/HMAC_SHA1_96, 0 bytes_i, 0 bytes_o, rekeying in 8 minutes
  meraki-vpn{5}:   192.168.5.99/32[udp/l2f] === 20.20.20.20/32[udp/l2f]
ifconfig :
eth0      Link encap:Ethernet  HWaddr 00:0c:29:5f:9c:73  
          inet addr:192.168.5.99  Bcast:192.168.5.255  Mask:255.255.255.0
          inet6 addr: fe80::20c:29ff:fe5f:9c73/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:459939 errors:0 dropped:0 overruns:0 frame:0
          TX packets:362145 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:78969068 (78.9 MB)  TX bytes:52179997 (52.1 MB)
lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          inet6 addr: ::1/128 Scope:Host
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:427152 errors:0 dropped:0 overruns:0 frame:0
          TX packets:427152 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:25278435 (25.2 MB)  TX bytes:25278435 (25.2 MB) 
Thanks!
