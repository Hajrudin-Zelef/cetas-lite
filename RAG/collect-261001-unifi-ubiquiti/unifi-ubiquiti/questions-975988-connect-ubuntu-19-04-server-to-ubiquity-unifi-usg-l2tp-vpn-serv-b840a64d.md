---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-975988-connect-ubuntu-19-04-server-to-ubiquity-unifi-usg-l2tp-vpn-serv-b840a64d
title: "questions-975988-connect-ubuntu-19-04-server-to-ubiquity-unifi-usg-l2tp-vpn-serv-b840a64d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-975988-connect-ubuntu-19-04-server-to-ubiquity-unifi-usg-l2tp-vpn-serv-b840a64d.md
source_anchor: ""
source_lines: [1, 32]
sha256: 9c89d82f76aefee9b77847b454ef3bfb051352616f5a120251f19094fab34265
---

# questions-975988-connect-ubuntu-19-04-server-to-ubiquity-unifi-usg-l2tp-vpn-serv-b840a64d

I am trying to connect a remote Ubuntu Server 19.04 instance to a Ubiquity Unifi USG L2TP VPN Server.
I have the VPN server set up as an L2TP server with PSK and verified that it is working by connecting from my Android phone.
It seems that Ubuntu Server 19.04 uses netplan to configure network interfaces; however, it also seems that netplan does not yet support l2tp vpn.
So I have installed network manager alongside netplan, as I have experiences that it will allow me to connect as an L2TP VPN client on a Ubuntu Desktop.
I created a network manager connection by writing the file /etc/NetworkManager/system-connections/VPN.nmconnection:
[connection]
id=VPN
uuid=f6019847-2f9b-4735-8f10-a235ea629576
type=vpn
autoconnect=false
permissions=
[vpn]
gateway=vpn-server-hostname
ipsec-psk=mypskkey
password-flags=2
user=cloudbox
service-type=org.freedesktop.NetworkManager.l2tp
[ipv4]
dns-search=
method=auto
[ipv6]
addr-gen-mode=stable-privacy
dns-search=
method=auto
After restarting network-manager I try to bring up the VPN, but I get the following error:
$ sudo nmcli -s c up VPN --ask
Error: Connection activation failed: Could not find source connection.
After googling that error a bit I find some information about the problem being that network-manager does not manage the network interface with the default gateway. So I create the empty file /etc/NetworkManager/conf.d/10-globally-managed-devices.conf and modify the file /etc/NetworkManager/NetworkManager.conf where I change the line managed=false to managed=true. Network-manager now manages and configures the primary interface correctly; however, I get the same error upon starting the VPN again:
$ sudo nmcli -s c up VPN --ask
Error: Connection activation failed: Could not find source connection.
At this point, I am completely stuck. Google does not seem to offer any suggestions.
My problem is simply: How can I configure my Ubuntu Server 19.04 to connect to my Ubiquity Unifi USG IPSec/L2TP PSK VPN Server?
