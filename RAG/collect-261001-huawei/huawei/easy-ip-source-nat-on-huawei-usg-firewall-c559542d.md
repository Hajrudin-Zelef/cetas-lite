---
id: collect-261001-huawei/huawei/easy-ip-source-nat-on-huawei-usg-firewall-c559542d
title: "easy-ip-source-nat-on-huawei-usg-firewall-c559542d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/easy-ip-source-nat-on-huawei-usg-firewall-c559542d.md
source_anchor: ""
source_lines: [1, 95]
sha256: d58e97140bcb8b186ddccd4290255ee18b9110253f06317074ac618a0b59302d
---

# easy-ip-source-nat-on-huawei-usg-firewall-c559542d

A short video describing **EASY-IP** source **NAT** on **Huawei’s USG** firewall to assure that our intranet users can acccess the Internet. 

The easy-ip is the easiest way to translate private IP addresses on your local area network when you have only one public IP address or the public IP address is assigned dynamically.

## USG firewall configuration script:

`#`

dhcp enable

#

interface GigabitEthernet1/0/0

undo shutdown

ip address 10.0.0.1 255.255.255.0

service-manage ping permit

dhcp select interface

dhcp server excluded-ip-address 10.0.0.100

dhcp server static-bind ip-address 10.0.0.200 mac-address 5489-98b4-6a79

dhcp server dns-list 10.0.0.100

#

interface GigabitEthernet1/0/2

undo shutdown

ip address 5.0.0.2 255.255.255.252

#

firewall zone trust

set priority 85

add interface GigabitEthernet1/0/0

#

firewall zone untrust

set priority 5

add interface GigabitEthernet1/0/2

#

ip route-static 0.0.0.0 0.0.0.0 5.0.0.1

#

security-policy

rule name ALLOW

source-zone local

destination-zone trust

destination-zone untrust

action permit

rule name NAT_EASY

source-zone trust

destination-zone untrust

source-address 10.0.0.0 mask 255.255.255.0

action permit

#

nat-policy

rule name NAT_EASY

source-zone trust

destination-zone untrust

source-address 10.0.0.0 mask 255.255.255.0

action source-nat easy-ip

#
