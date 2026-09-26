---
id: collect-260926-mikrotik/mikrotik/questions-635247-mikrotik-ipsec-vpn-routing-045bdee6
title: "questions-635247-mikrotik-ipsec-vpn-routing-045bdee6"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-635247-mikrotik-ipsec-vpn-routing-045bdee6.md
source_anchor: ""
source_lines: [1, 27]
sha256: 40130d3dd44fa54427b0e5c9bb7895ab91fd3a801131978989ac905c4883f83c
---

# questions-635247-mikrotik-ipsec-vpn-routing-045bdee6

I am trying to setup an IPSec VPN tunnel so as to secure communication between my private LAN and a destination host. Any device within my private LAN should be able to initiate connection to the destination host. However, if the destination host wants to connect to my network (directed to my public IP address), I want to forward that connection to just one particular server - 192.168.1.65.
My setup
Private LAN: 192.168.1.1/24
Public IP: 50.X.X.X
Destination Host IP: 173.X.X.X (using CISCO ASA)
I got the basic setup working fine and my internal LAN is able to access the internet. My attempt at the IPSec configuration is as follows:
/interface ipip
add comment="" disabled=no local-address=50.X.X.X mtu=1460 name=ipip1 \
    remote-address=173.X.X.X
/ip address
add address=192.168.1.1/24 broadcast=192.168.1.255 comment="" disabled=no \
    interface=ipip1 network=192.168.1.0
/ip ipsec peer
add address=173.X.X.X/32 auth-method=pre-shared-key comment="" \
    dh-group=modp1024 disabled=no dpd-interval=disable-dpd \
    enc-algorithm=3des exchange-mode=main generate-policy=no \
    hash-algorithm=md5 lifebytes=0 lifetime=1d nat-traversal=no port=500 \
    proposal-check=obey secret=SECRETKEY send-initial-contact=yes
/ip ipsec policy
add action=encrypt comment="" disabled=no dst-address=173.X.X.X/32 dst-port=any \
    ipsec-protocols=esp level=require priority=0 proposal=ipsec protocol=all \
    sa-dst-address=173.X.X.X sa-src-address=50.X.X.X \
    src-address=50.X.X.X/32:any tunnel=yes
/ip ipsec proposal
set default auth-algorithms=md5 comment="" disabled=no enc-algorithms=3des \
    lifetime=60m name=ipsec pfs-group=none
What to do next? How to configure such that my devices can initiate connection to the host and for the host to initiate connection to just one particular server? Would it be IP firewall NAT, masquerade or IP route?
