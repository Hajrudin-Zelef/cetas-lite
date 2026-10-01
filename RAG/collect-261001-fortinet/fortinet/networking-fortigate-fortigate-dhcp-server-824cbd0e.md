---
id: collect-261001-fortinet/fortinet/networking-fortigate-fortigate-dhcp-server-824cbd0e
title: "Configure FortiGate as a DHCP Server via GUI and CLI"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-fortinet/networking-fortigate-fortigate-dhcp-server-824cbd0e.md
source_anchor: ""
source_lines: [1, 73]
sha256: 92081b357542bf7c4beb73bab9155208ad999dc5bb874d66dae516398c5995bf
---

# Configure FortiGate as a DHCP Server via GUI and CLI

*Source : https://defencedev.com/networking/fortigate/fortigate-dhcp-server/*

Configuring FortiGate as a DHCP server allows network administrators to easily manage IP address distribution within their network. In this article, we will explore how to configure FortiGate as a DHCP server both via the GUI and CLI, including some advanced configuration options.
A DHCP (Dynamic Host Configuration Protocol) server is used to dynamically assign IP addresses to devices on a network. FortiGate firewalls, aside from their security capabilities, can also act as DHCP servers.
To configure FortiGate as a DHCP server through the graphical user interface (GUI), follow these steps:
If you need to make IP ADress Reservation you can add MAC Adress of device and IP Adress in the “MAC Reservatrion” Table:
For administrators who prefer the command-line interface (CLI) or need to automate the configuration, setting up a DHCP server via CLI is also straightforward.
Use the following command to configure the DHCP server on a specific interface (e.g., internal):
fgt-remote1 # config system dhcp server
fgt-remote1 (server) # edit 1
fgt-remote1 (1) # get
id                  : 1
status              : enable
lease-time          : 604800
mac-acl-default-action: assign
forticlient-on-net-status: enable
dns-service         : specify
wifi-ac1            : 0.0.0.0
wifi-ac2            : 0.0.0.0
wifi-ac3            : 0.0.0.0
ntp-service         : specify
domain              :
wins-server1        : 0.0.0.0
wins-server2        : 0.0.0.0
default-gateway     : 192.168.20.1
next-server         : 0.0.0.0
netmask             : 255.255.255.0
interface           : internal
ip-range:
    == [ 1 ]
    id:     1
timezone-option     : default
tftp-server         :
filename            :
options:
server-type         : regular
conflicted-ip-timeout: 1800
auto-configuration  : enable
ddns-update         : disable
vci-match           : disable
exclude-range:
reserved-address:
dns-server1         : 8.8.8.8
dns-server2         : 8.8.4.4
dns-server3         : 0.0.0.0
ntp-server1         : 0.0.0.0
ntp-server2         : 0.0.0.0
ntp-server3         : 0.0.0.0
fgt-remote1 (1) #
In this configuration:
internal).
After configuring the DHCP server, you can verify the settings with the following command: show system dhcp server
fgt-remote1 # show system dhcp server
config system dhcp server
    edit 1
        set default-gateway 192.168.20.1
        set netmask 255.255.255.0
        set interface "internal"
        config ip-range
            edit 1
                set start-ip 192.168.20.110
                set end-ip 192.168.20.210
            next
        end
        set timezone-option default
        set dns-server1 8.8.8.8
        set dns-server2 8.8.4.4
    next
end
Configuring FortiGate as a DHCP server, either via the GUI or CLI, is a simple yet powerful way to manage IP address allocation for your network devices. By following the steps outlined above, you can easily set up a DHCP server, customize the IP address range, and apply advanced options like static IP mapping, TFTP server configuration, and DHCP relay. The flexibility offered by FortiGate allows you to fine-tune your DHCP settings to meet the specific needs of your network.
In next article we will write about DHCP special atributes. You can check my other FortiGate post on this link.
