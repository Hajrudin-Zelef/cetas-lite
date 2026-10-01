---
id: collect-261001-fortinet/fortinet/support-forum-92-aws-fortigate-vm-to-aws-virtual-private-gateway-vpn-112207-41a74e7d
title: "support-forum-92-aws-fortigate-vm-to-aws-virtual-private-gateway-vpn-112207-41a74e7d"
domain: fortinet
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-fortinet/support-forum-92-aws-fortigate-vm-to-aws-virtual-private-gateway-vpn-112207-41a74e7d.md
source_anchor: ""
source_lines: [1, 150]
sha256: 6d6cea980259ad30421370b531cb3e84d6fd95ba3631c822b0873e7763e67a30
---

# support-forum-92-aws-fortigate-vm-to-aws-virtual-private-gateway-vpn-112207-41a74e7d

AWS Fortigate VM to AWS Virtual Private Gateway VPN
Hello all,
I have deployed a Fortigate VM (Version 7.2.0) in one of the VPC in AWS. This is a new deployment and no additional configurations have been done so far. I have also created a Virtual Private Gateway(VGW) and attached a Customer Gateway(CGW) to the VGW. Finally, I am trying to establish an IPsec tunnel between the Fortigate VM in AWS and AWS Virtual Private Gateway. I downloaded the configuration from AWS and configured the Fortigate in AWS for the IPsec tunnel. Unfortunately, I am unable to get this up and running and I need the help of the experts here.
Can someone please guide me why the VPN is not coming up? I can see that the fortigate is trying to initiate the VPN connection and does not get beyond SA_INIT.
EUDC2-Transit-FW01 # diagnose vpn ike status detailed
vd: root/0
name: VPN_to_AWS_VGW
version: 2
connection: 1/1687
IKE SA: created 1/1687
IPsec SA: created 1/1687
EUDC2-Transit-FW01 # diagnose vpn ike config
vd: root/0
name: VPN_to_AWS_VGW
serial: 1
version: 2
status.admin: up
status.operational: down
type: static
local: <elastic IP address of fortigate>
remote: <Virtual Private Gateway Public IP address>
mode: main
dpd: on-idle retry-count 3 interval 20000ms
auth: psk
dhgrp: 14 2
xauth: none
interface: port1
virtual-interface-addr: 169.254.127.162 -> 169.254.127.161
auto-discovery-sender: disable
auto-discovery-receiver: disable
phase2s:
vpn-085017f3dc9294023-0 proto 0 src 0.0.0.0/0.0.0.0:0 dst 0.0.0.0/0.0.0.0:0 dhgrp 14 2 replay keep-alive auto-negotiate route-new
policy: yes
EUDC2-Transit-FW01 # diagnose vpn ike gatewa
vd: root/0
name: VPN_to_AWS_VGW
version: 2
interface: port1 3
addr: <Elastic IP address of Fortigate>:500 -> <Virtual Private Gateway Public IP address>:500
tun_id: <Virtual Private Gateway Public IP address>/::<Virtual Private Gateway Public IP address>
remote_location: 0.0.0.0
virtual-interface-addr: 169.254.127.162 -> 0.0.0.0
created: 22s ago
PPK: no
IKE SA: created 1/1
IPsec SA: created 1/1
id/spi: 1727 495a17243cf25b39/0000000000000000
direction: responder
status: connecting, state 3, started 22s ago
EUDC2-Transit-FW01 # show vpn ipsec phase1-interface VPN_to_AWS_VGW
config vpn ipsec phase1-interface
edit "VPN_to_AWS_VGW"
set interface "port1"
set ike-version 2
set local-gw <Elastic IP address of Fortigate>
set keylife 28800
set peertype any
set net-device disable
set proposal aes128-sha1 aes256-sha256 aes128gcm-prfsha256 aes256gcm-prfsha384 chacha20poly1305-prfsha256
set dpd on-idle
set comments "VPN to AWS VGW connecting with Direct Connect"
set dhgrp 14 2
set nattraversal disable
set remote-gw <Virtual Private Gateway Public IP address>
set psksecret ENC <PSK>
next
end
EUDC2-Transit-FW01 # show system interface VPN_to_AWS_VGW
config system interface
edit "VPN_to_AWS_VGW"
set vdom "root"
set ip 169.254.127.162 255.255.255.255
set allowaccess ping
set type tunnel
set tcp-mss 1379
set remote-ip 169.254.127.161 255.255.255.252
set description "Tunnel with AWS VGW"
set snmp-index 9
set mtu-override enable
set mtu 1427
set interface "port1"
next
end
EUDC2-Transit-FW01 # show firewall security-policy 1
config firewall security-policy
edit 1
set uuid caeb7206-11a3-51ed-ec10-d39d55569809
set name "Connection_to_VPN"
set comments "VPN connection"
set srcintf "port2"
set dstintf "VPN_to_AWS_VGW"
set srcaddr "all"
set dstaddr "all"
set action accept
set schedule "always"
set logtraffic all
set av-profile "default"
set webfilter-profile "default"
set ips-sensor "default"
next
end
EUDC2-Transit-FW01 # show firewall security-policy 2
config firewall security-policy
edit 2
set uuid 319f79fc-11a4-51ed-d632-617ccbdc4aec
set name "Traffic_From_VPN"
set comments "Traffic from VPN"
set srcintf "VPN_to_AWS_VGW"
set dstintf "port2"
set srcaddr "all"
set dstaddr "all"
set action accept
set schedule "always"
set logtraffic all
set av-profile "default"
set webfilter-profile "default"
set ips-sensor "default"
next
end
EUDC2-Transit-FW01 # show router bgp
config router bgp
set as 65000
set router-id <Elastic IP address of Fortigate>
config neighbor
edit "169.254.127.161"
set remote-as 64512
next
end
config network
edit 1
set prefix 10.90.224.0 255.255.255.0
next
end
EUDC2-Transit-FW01 # show router static
config router static
edit 1
set dst 10.90.192.0 255.255.224.0
set gateway 10.90.224.1
set device "port2"
set comment "Route to Production VPC"
next
edit 2
set gateway 10.90.226.1
set device "port1"
set comment "Internet access"
next
end
I hope that I have given enough information here to help me troubleshoot this.
Regards,
Anand
