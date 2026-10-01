---
id: collect-261001-fortinet/fortinet/2018-08-fortinet-firewall-fortigate-cli-html-37ec343f-1
title: "2018-08-fortinet-firewall-fortigate-cli-html-37ec343f"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/2018-08-fortinet-firewall-fortigate-cli-html-37ec343f.md
source_anchor: ""
source_lines: [1, 270]
sha256: 02eaa6e6af5b446dfd7992703b233c527566e13633ab7b4950afe3e62ecd4e6d
---

# 2018-08-fortinet-firewall-fortigate-cli-html-37ec343f

FortiGate firewall always surprise me with his rich embedded features, prices and performance. FortiOS is a security-hardened, purpose-built operating system that is the software foundation of FortiGate products. With this one unified intuitive OS, we can control all the security and networking capabilities across all of your Fortigate products.

I put some of useful commands or configurations in following two posts:

- FortiOS Configuration for FortiGate Firewalls (Commands) 1
- FortiOS Configuration for FortiGate Firewalls (Commands) 2

## 1. Debugging and Diagnostic your system

```
diag debug enable
diag debug console timestamp enable
diag sniffer packet wan 'host 8.8.8.8' 1
diag debug disable
diag debug reset
```
diag debug cli cmd will show you the "cli commands" for actions that you take from the gui.

diag debug enable
diag debug cli 8 

FWF60D # diag sys flash list 
Partition  Image                                     TotalSize(KB)  Used(KB)  Use%  Active
1          FWF60D-5.06-FW-build1486-170816                  253871     48877   19%  No    
2          FWF60D-6.00-FW-build0163-180725                  253871     53598   21%  Yes   
3          ETDB-1.00000                                    3368360     81632    2%  No    
Image build at Jul 25 2018 19:06:34 for b0163

diag sys tcpsock command will show you the active opening ports in your system.

FWF60D # diag sys tcpsock 
0.0.0.0:10400->0.0.0.0:0->state=listen err=0 sockflag=0x8 rma=0 wma=0 fma=0 tma=0
0.0.0.0:5060->0.0.0.0:0->state=listen err=0 sockflag=0x102 rma=0 wma=0 fma=0 tma=0
0.0.0.0:135->0.0.0.0:0->state=listen err=0 sockflag=0x2 rma=0 wma=0 fma=0 tma=0
0.0.0.0:1004->0.0.0.0:0->state=listen err=0 sockflag=0x8 rma=0 wma=0 fma=0 tma=0
0.0.0.0:1005->0.0.0.0:0->state=listen err=0 sockflag=0x2 rma=0 wma=0 fma=0 tma=0
0.0.0.0:7822->0.0.0.0:0->state=listen err=0 sockflag=0x1 rma=0 wma=0 fma=0 tma=0
0.0.0.0:910->0.0.0.0:0->state=listen err=0 sockflag=0x1 rma=0 wma=0 fma=0 tma=0
0.0.0.0:80->0.0.0.0:0->state=listen err=0 sockflag=0x1 rma=0 wma=0 fma=0 tma=0
0.0.0.0:80->0.0.0.0:0->state=listen err=0 sockflag=0x2 rma=0 wma=0 fma=0 tma=0
0.0.0.0:10000->0.0.0.0:0->state=listen err=0 sockflag=0x8 rma=0 wma=0 fma=0 tma=0
0.0.0.0:2000->0.0.0.0:0->state=listen err=0 sockflag=0x102 rma=0 wma=0 fma=0 tma=0
0.0.0.0:1010->0.0.0.0:0->state=listen err=0 sockflag=0x2 rma=0 wma=0 fma=0 tma=0
......

1.1 Debug VPN

Enable Debugging

FWF60D #diag debug enable
FWF60D #diag debug console timestamp enable 

`FWF60D #diag vpn ike log-filter dst-addr4` 
```
FWF60D #diag debug application ike -1 
```
| `FWF60D # tree diag vpn ike gateway` | 
| `-- gateway -- list -- name -- <name>  (0)` | 
|            `\|- clear -- name -- <name>  (0)` | 
|            `+- flush -- name -- <name>  (0)` | 

Disable Debugging

FWF60D #diag debug disable
FWF60D #diag debug console timestamp disable
FWF60D #diag debug application ike 0

Show active tunnel and gateway list

FWF60D #diag vpn tunnel list
FWF60D #diag vpn gw list

## **2. Get system configuraiton**

get system arp          // ARP Table

get system dns // DNS Configuration

get system dhcp server // DHCP server configuration


FGT30D # get system setting

opmode : nat

firewall-session-dirty: check-all

bfd : disable

utf8-spam-tagging : enable

wccp-cache-engine : disable

vpn-stats-log :

vpn-stats-period : 0

v4-ecmp-mode : source-ip-based

gui-default-policy-columns:

asymroute : disable

ses-denied-traffic : disable

strict-src-check : disable

asymroute6 : disable

per-ip-bandwidth : disable

sip-helper : enable

sip-nat-trace : enable

status : enable

sip-tcp-port : 5060

sip-udp-port : 5060

sccp-port : 2000

multicast-forward : enable

multicast-ttl-notchange: disable

allow-subnet-overlap: disable

deny-tcp-with-icmp : disable

ecmp-max-paths : 10

discovered-device-timeout: 28

email-portal-check-dns: enable

show system interface wan1 | grep -A2 ip // Show WAN and interface information.

get system info admin status // Show logged in users

get system status // Show system hardware/software update versions

get hardware status // Detailed hardware model information

get system performance status // Check System Uptime

FGT30D3X12001671 $ get system performance status 
CPU states: 0% user 0% system 0% nice 100% idle
CPU0 states: 0% user 0% system 0% nice 100% idle
Memory states: 21% used
Average network usage: 3 kbps in 1 minute, 0 kbps in 10 minutes, 0 kbps in 30 minutes
Average sessions: 14 sessions in 1 minute, 11 sessions in 10 minutes, 11 sessions in 30 minutes
Average session setup rate: 0 sessions per second in last 1 minute, 0 sessions per second in last 10 minutes, 0 sessions per second in last 30 minutes
Virus caught: 0 total in 1 minute
IPS attacks blocked: 0 total in 1 minute
Uptime: 106 days,  0 hours,  8 minutes

get system performance top

show system interface

diagnose hardware deviceinfo nic // Interface Statistics/Settings

diagnose hardware sysinfo memory

diag debug crashlog read

diag hardware sysinfo shm // Device should be in 0, if (>0) then conservemode

get system global | grep -i timer // Show tcp and udp timers for halfopen and idle

get system session-ttl // System default tcp-idle session timeout

get hardware nic

get system interface physical

diagnose ip address list

diagnose ip arp list

diagnose sys session list

diagnose sys session clear

diagnose sys kill 9 <id>

## **3. Change Bult-in Internal Switch to Interface mode**

In Switch mode, all the internal interfaces are part of the same subnet and treated as a single interface, called either lan or internal by default, depending on the FortiGate model. Switch mode is used when the network layout is basic, with most users being on the same subnet.
In Interface mode, the physical interfaces of the FortiGate unit are handled individually, with each interface having its own IP address. Interfaces can also be combined by configuring them as part of either hardware or software switches, which allow multiple interfaces to be treated as a single interface.

**a. Command to change the FortiGate to switch mode:**

config system global

set internal-switch-mode switch

end

**b. Command to change the FortiGate to interface mode:**

config system global

set internal-switch-mode interface

end

After changed internal switch from switch mode to interface mode, you will be able to move some interface out of Internal switch and they will become routing interfaces for you to do configuration.

Note: How to Change Switch Mode to Interface Mode in Fortigate FortiOS 5

## **4. Daily System Scheduled Reboot**

config system global
set daily-restart enable
set restart-time 05:06
end

Note: For weekly reboot, you will need expect command with a script.

## **5. Some HA Commands**

**Manual Failover HA**

diagnose sys ha reset-uptime


**Mange Cluster Member from Console**

```
Test-1 # get system ha status 
Model: FortiGate-60D
Mode: a-p
Group: 0
Debug: 0
ses_pickup: disable
Master:250 Test-1    FGT60D4614041798 1
Slave : 50 Test-2    FGT60D4Q15005710 0
number of vcluster: 1
vcluster 1: work 169.254.0.2
Master:0 FGT60D4614041798
Slave :1 FGT60D4Q15005710
Test-1 # execute ha manage 0
 
Test-2 $ 
Test-2 $ execute reboot 
This operation will reboot the system !
Do you want to continue? (y/n)y
```
## 6. One to One Inbound NAT Configuration

WAN IP Address : 192.168.20.200
LAN IP Address: 192.168.2.200

Rule: Allow Any External IP Address to access LAN Server 1921.68.2.200 on SMB (TCP 445 - File Sharing Port), but not expose LAN IP Address.

6.1 Create VIP address

6.2 Create firewall Rule

Note: No NAT configuration. NAT will be taken care by VIP configuration in step 6.1

Destination will be External IP Address.

## 6. Tree command

Tree command can be used to display the command tree for a configuration section.

