---
id: collect-261001-cisco/cisco/github-lotsofproblems-cisco-commands-cheatsheet-for-ccna-commands-2
title: "Cisco Commands Cheat Sheet"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/github-lotsofproblems-cisco-commands-cheatsheet-for-ccna-commands.md
source_anchor: ""
source_lines: [247, 472]
sha256: 34f4ae830f51d5e351e8a46b4e7b0164f4713ef6a748590c3e3926fd5330e2fc
---

# Cisco Commands Cheat Sheet

```
crypto isakmp policy 10
 encryption aes 256
 hash sha
 authentication pre-share
 group 5
 lifetime 86400
exit
crypto isakmp key SECRET-KEY address 203.0.113.1
```

##### Configure IPsec Transform Set (Encryption)

```
configure terminal
crypto ipsec transform-set VPN-SET esp-aes esp-sha-hmac
 mode tunnel
exit
```

##### Create Crypto Map and apply Transform Set

Keep in mind ACL number

```
crypto map VPN-MAP 10 ipsec-isakmp
 set peer 203.0.113.1
 set transform-set VPN-SET
 match address 110
exit
```

##### Define an ACL for Traffic Encryption

```
access-list 110 permit gre host 192.168.1.1 host 203.0.113.1
exit
```

##### Apply the Crypto Map to the WAN Interface

```
interface GigabitEthernet0/0
 crypto map VPN-MAP
exit
```

### Class Mapping and Marking


```
class-map match-all CRITICAL
 match protocol ospf
class-map match-any MANAGEMENT
 match protocol telnet
 match protocol ssh
class-map match-any WEB
 match protocol http
!
policy-map MARKING
 class CRITICAL
  set precedence 7
 class MANAGEMENT
  set precedence 5
 class WEB
  set precedence 3
interface g0/0/0
service-policy output MARKING
```


| **Value** | **Description** | 
| `0` | Routine - Match packets with routine precedence | 
| `1` | Priority - Match packets with priority precedence | 
| `2` | Immediate - Match packets with immediate precedence | 
| `3` | Flash - Match packets with flash precedence | 
| `4` | Flash Override - Match packets with flash override precedence | 
| `5` | Critical - Match packets with critical precedence | 
| `6` | Internet - Match packets with internetwork control precedence | 
| `7` | Network - Match packets with network control precedence | 


```
class-map ?
  WORD       class-map name
  match-all  Logical-AND all matching statements under this classmap
  match-any  Logical-OR all matching statements under this classmap
  type       type of the class-map
```


```
match ?
  access-group         Access group
  any                  Any packets
  class-map            Class map
  cos                  IEEE 802.1Q/ISL class of service/user priority values
  destination-address  Destination address
  input-interface      Select an input interface to match
  ip                   IP specific values
  not                  Negate this match result
  precedence           Match Precedence in IP(v4) and IPv6 packets
  protocol             Protocol
  qos-group            Qos-group
```



##### Disable (enabled by default)

```
show cdp neighbors
show cdp neighbors detail
```



##### Enable (disabled by default on terminal console)

```
lldp receive
lldp transmit
```

```
logging on
terminal monitor
logging buffered <size> <level>
service timestamps log datetime msec
Levels range from 0 (emergencies) to 7 (debugging):
0 - emergencies
1 - alerts
2 - critical
3 - errors
4 - warnings
5 - notifications
6 - informational
7 - debugging
```

```
logging host <IP_Address>
logging host 192.168.1.100
logging trap informational
snmp-server enable traps syslog
```



```
copy running-config tftp:100.10.1.10 R_R-C
```

```
copy <TFTPpath> tftp:<IPdest> <filename>
```

````
# Cisco IOS NAT (Network Address Translation) Commands
## 1. Enable NAT
```sh
configure terminal
````

## 2. Static NAT Configuration

ip nat inside source static <inside-local-ip> <inside-global-ip>

Example:

ip nat inside source static 192.168.1.10 203.0.113.10

## 3. Dynamic NAT Configuration

ip nat pool NAT_POOL <start-ip> <end-ip> netmask <subnet-mask>
access-list <acl-number> permit <source-ip> <wildcard-mask>
ip nat inside source list <acl-number> pool NAT_POOL

Example:

ip nat pool NAT_POOL 203.0.113.100 203.0.113.200 netmask 255.255.255.0
access-list 1 permit 192.168.1.0 0.0.0.255
ip nat inside source list 1 pool NAT_POOL

## 4. Port Address Translation (PAT) / NAT Overload

ip nat inside source list <acl-number> interface <outside-interface> overload

Example:

access-list 1 permit 192.168.1.0 0.0.0.255
ip nat inside source list 1 interface GigabitEthernet0/0 overload

## 5. NAT Interface Configuration

interface <interface-name>
ip nat inside  # For inside interface
exit
interface <interface-name>
ip nat outside  # For outside interface
exit

Example:

interface GigabitEthernet0/1
 ip nat inside
 exit
interface GigabitEthernet0/0
 ip nat outside
 exit

## 6. NAT Verification Commands

show ip nat translations  # Display NAT translation table
show ip nat statistics  # Display NAT statistics
clear ip nat translation *  # Clear all NAT translations

## 7. Remove NAT Configuration

no ip nat inside source static <inside-local-ip> <inside-global-ip>
no ip nat inside source list <acl-number> pool NAT_POOL
no ip nat inside source list <acl-number> interface <outside-interface> overload

- **Static NAT** : Maps one inside local IP to one public IP.
- **Dynamic NAT** : Uses a pool of public IPs for internal hosts.
- **PAT (Overload)** : Maps multiple internal IPs to a single public IP using different port numbers.
