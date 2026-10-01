---
id: collect-261001-huawei/huawei/network-ptmngsys-web-tsrev-ar-en-content-ar-06-edesk-internet-users-can-not-acce-a245c549
title: "network-ptmngsys-web-tsrev-ar-en-content-ar-06-edesk-internet-users-can-not-acce-a245c549"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/network-ptmngsys-web-tsrev-ar-en-content-ar-06-edesk-internet-users-can-not-acce-a245c549.md
source_anchor: ""
source_lines: [1, 18]
sha256: e9c682965007d2cfc91943d68425bf6c3e5c542bae18a6d2fc38b487ede62245
---

# network-ptmngsys-web-tsrev-ar-en-content-ar-06-edesk-internet-users-can-not-acce-a245c549

<Huawei> display nat static
  Nat Server Information: 
  Interface  : GigabitEthernet0/0/0                                     //Interface configured with the NAT server
    Global IP/Port     : current-interface/80(www) (Real IP : 1.1.1.1)  //Public IP address and port number
    Inside IP/Port     : 10.10.10.3/80(www)                             //Private IP address and port number
    Protocol : 6(tcp)                                                   //Application protocol
    VPN instance-name  : ----                                           //VPN instance name
    Acl number         : ----                                           //ACL number
    Vrrp id            : ----                                           //VRRP ID
    Netmask  : 255.255.255.255                                          //Natwork mask
    Description : ----                                                  //Description
 
  Total :    1
 
      If the translated private and public IP addresses, port numbers, and application protocol and other information are incorrect, modify them. If the network access mode is PPPoE, configure the NAT server on the dialer interface associated with the public network interface.
<Huawei> system-view
[Huawei] interface dialer 0
[Huawei-Dialer0] nat static protocol tcp global current-interface 80 inside 10.10.10.3 80
