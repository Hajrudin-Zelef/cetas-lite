---
id: collect-261001-cisco/cisco/github-vanchhai-fortigate-cisco-redundancy-2
title: "Full Network Redundancy (FortiGate Firewall  +  Cisco Core Switch)"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "voice"]
source: docs/RAG/collect-261001-cisco/github-vanchhai-fortigate-cisco-redundancy.md
source_anchor: ""
source_lines: [179, 528]
sha256: f15a358611c89e56fd52ef2a4c05c19da9b084710cfc9c4738d94209b1b2ee92
---

# Full Network Redundancy (FortiGate Firewall  +  Cisco Core Switch)

```
# LACP aggregate to core (port1 + port2)
config system interface
    edit "agg1"
        set vdom "root"
        set type aggregate
        set member "port1" "port2"
        set lacp-mode active
        set lacp-speed fast
        set min-links 1
    next
end
 
# Sub-interface VLANs on the aggregate
config system interface
    edit "vl10-users"
        set vdom "root"
        set interface "agg1"
        set vlanid 10
        set ip 10.10.10.2 255.255.255.0
        set allowaccess ping https ssh
    next
    edit "vl20-servers"
        set vdom "root"
        set interface "agg1"
        set vlanid 20
        set ip 10.10.20.2 255.255.255.0
        set allowaccess ping
    next
end
```
```
# WAN interfaces for dual ISP
config system interface
    edit "wan1"
        set alias "ISP-A"
        set mode static
        set ip 203.0.113.2 255.255.255.252
        set allowaccess ping
    next
    edit "wan2"
        set alias "ISP-B"
        set mode static
        set ip 198.51.100.2 255.255.255.252
        set allowaccess ping
    next
end
 
# SD-WAN with health checks
config system sdwan
    set status enable
    config members
        edit 1
            set interface "wan1"
            set gateway 203.0.113.1
        next
        edit 2
            set interface "wan2"
            set gateway 198.51.100.1
        next
    end
end
```
### FortiGate — Routing, SD-WAN Rule & Policy (Failover policy and outbound NAT)

```
# Static default routes (priority-based failover)
config router static
    edit 1
        set dst 0.0.0.0 0.0.0.0
        set gateway 203.0.113.1
        set device "wan1"
        set priority 10
    next
    edit 2
        set dst 0.0.0.0 0.0.0.0
        set gateway 198.51.100.1
        set device "wan2"
        set priority 20
    next
end
 
# SD-WAN health check via ISP
config system sdwan
    config health-check
        edit "ping-google"
            set server "8.8.8.8" "1.1.1.1"
            set members 1 2
            set failtime 3
            set recoverytime 5
        next
    end
end
```
```
# SD-WAN rule: best quality wins
config system sdwan
    config service
        edit 1
            set name "to-internet"
            set mode sla
            set dst "all"
            set src "all"
            config sla
                edit "ping-google"
                    set id 1
                next
            end
            set priority-members 1 2
        next
    end
end
 
# Outbound firewall policy with NAT
config firewall policy
    edit 100
        set name "LAN-to-INET"
        set srcintf "agg1"
        set dstintf "virtual-wan-link"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set service "ALL"
        set nat enable
    next
end
```
## Cisco Core — StackWise-Virtual

### Two physical switches, one logical control plane

- Two Catalyst 9300/9500 switches behave as ONE logical switch
- Active + Standby supervisors with Stateful Switchover (SSO)
- StackWise-Virtual Link (SVL) — 2× 10/40G between members
- Dual-Active Detection (DAD) link prevents split-brain
- Multichassis EtherChannel (MEC) — LACP across both members
- Single config, single management IP, single STP root
- Sub-second convergence on member or link failure
- Upgrade with ISSU (In-Service Software Upgrade)

- Single Logical Core Switch
- StackWise-Virtual ID 10
- Mgmt IP 10.255.0.10

## Cisco IOS XE — StackWise-Virtual & VLANs

# On Switch-1 (will become Active)
!
configure terminal
stackwise-virtual
 domain 10
!
interface range TenGigabitEthernet1/1/1-2
 stackwise-virtual link 1
!
interface TenGigabitEthernet1/0/48
 stackwise-virtual dual-active-detection
!
! On Switch-2
stackwise-virtual
 domain 10
!
interface range TenGigabitEthernet2/1/1-2
 stackwise-virtual link 1
!
interface TenGigabitEthernet2/0/48
 stackwise-virtual dual-active-detection
!
write memory
reload          # reboot both members
!

# VTP & VLAN database (after stack forms)
!
vtp mode transparent
!
vlan 10
 name USERS
vlan 20
 name SERVERS
vlan 30
 name VOICE
vlan 99
 name MGMT
vlan 999
 name NATIVE-UNUSED
!
# Spanning tree — core is root
!
spanning-tree mode rapid-pvst
spanning-tree vlan 1-4094 priority 4096
spanning-tree extend system-id
!
# Mgmt interface (logical)
!
interface Vlan99
 ip address 10.255.0.10 255.255.255.0
 no shutdown
 !

## Cisco IOS XE — LACP Uplinks to FortiGate

### Multichassis EtherChannel (MEC) across both stack members

# Port-channel to FortiGate-A (Po10)
!
interface Port-channel10
 description >>> FGT-A agg1 <<<
 switchport
 switchport mode trunk
 switchport trunk native vlan 999
 switchport trunk allowed vlan 10,20,30,99
 spanning-tree guard root
!
interface TenGigabitEthernet1/0/1
 description FGT-A port1
 channel-group 10 mode active
!
interface TenGigabitEthernet2/0/1
 description FGT-A port2
 channel-group 10 mode active
!

# Port-channel to FortiGate-B (Po20)
!
interface Port-channel20
 description >>> FGT-B agg1 <<<
 switchport
 switchport mode trunk
 switchport trunk native vlan 999
 switchport trunk allowed vlan 10,20,30,99
 spanning-tree guard root
!
interface TenGigabitEthernet1/0/2
 description FGT-B port1
 channel-group 20 mode active
!
interface TenGigabitEthernet2/0/2
 description FGT-B port2
 channel-group 20 mode active
!

## Cisco IOS XE — HSRP Gateway Redundancy

### Virtual IPs on the core stack point clients at the FortiGate VIP

# Users SVI with HSRPv2
!
interface Vlan10
 description USERS GATEWAY
 ip address 10.10.10.252 255.255.255.0
 standby version 2
 standby 10 ip 10.10.10.1
 standby 10 priority 110
 standby 10 preempt delay minimum 60
 standby 10 timers msec 250 msec 750
 standby 10 authentication md5 key-string CORE-HSRP
 standby 10 track 1 decrement 20
!
# Servers SVI
!
interface Vlan20
 description SERVERS GATEWAY
 ip address 10.10.20.252 255.255.255.0
 standby version 2
 standby 20 ip 10.10.20.1
 standby 20 priority 110
 standby 20 preempt
!

# Object tracking — uplink to FortiGate HA VIP
!
track 1 ip route 10.10.10.2 255.255.255.255 reachability
!
# Default route to FortiGate cluster VIP
!
ip route 0.0.0.0 0.0.0.0 10.10.10.2 name TO-FGT
!
# Recommended L3 hardening
!
ip cef
no ip source-route
no ip http server
ip http secure-server
service password-encryption
!
# SSH + AAA
!
username netadmin privilege 15 secret S3cret!2026
ip domain name corp.local
crypto key generate rsa modulus 2048
line vty 0 15
 transport input ssh
 login local
 exec-timeout 10 0
!

## Verification — Operational Commands (Confirm HA, stack, channels and gateway state)

# Cluster health
!
get system ha status
diagnose sys ha status
diagnose sys ha checksum cluster
!
 
# Heartbeat & sync
!
diagnose sys ha dump-by group
diagnose sys session sync
!
 
# Interface & LACP
!
diagnose netlink aggregate name agg1
get system interface physical
!
 
# SD-WAN & routing
diagnose sys sdwan health-check
get router info routing-table all
 
# Force failover for testing
diagnose sys ha reset-uptime

# Stack & redundancy
!
show stackwise-virtual
show stackwise-virtual link
show stackwise-virtual dual-active-detection
show redundancy
show switch detail
!
# Port-channels
!
show etherchannel summary
show lacp neighbor
show interfaces Port-channel10 trunk
!
# HSRP & STP
!
show standby brief
show spanning-tree summary
show ip route
!
# CPU / health
!
show platform resources
show logging | include HA|FAIL
!
