---
id: collect-261001-general-networking/general-networking/wp-content-uploads-2026-06-preview-v1-pdf-24e88431-5
title: "wp-content-uploads-2026-06-preview-v1-pdf-24e88431"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-general-networking/wp-content-uploads-2026-06-preview-v1-pdf-24e88431.md
source_anchor: ""
source_lines: [384, 548]
sha256: 1bd07ee34dd50ae363a8afb1eb129ba07b5056e4dc789fae0013467a9916e84b
---

# wp-content-uploads-2026-06-preview-v1-pdf-24e88431

OOB Subnets PC IP Address Gateway IP Address Clients 
172.20.20.0/24 172.20.20.100 172.20.20.1 OOB Client 
2.3.5 Core Infrastructure  
In this LAB, Fortinet Firewall acts as the gateway for traffic betw een zones and the external network. We will 
discuss zones in detail later in session 2.4.12. Port2 is assigned for connection to LAN zone using Static routing 
in the subnet 192.168.22.0/30. Port3 to Port4 connect FW-core to DMZ an d port5 to port6 connect FW-core to 
Data Center zone. 
2.4 Basic Setup from factory default 
2.4.1 Getting Familiar with CLI 
In this lab, we will log in to the CLI and configure the device's manage ment IP address so we can access the 
graphical interface. Start by logging in to the FortiGate for the first  time, then explore its CLI. Double-click the 
FortiGate icon to open the CLI console. The default username is admin with a blank password. Upon first login, 
FortiGate prompts you to set a new password. 
FortiGate-VM64-KVM login: admin 
Password:  
You are forced to change your password. Please input a new password. 
New Password:  
Confirm Password:  
Welcome! 
As you can see in the configuration panel, port1 is currently set to DHCP. Since we don’t have a DHCP server 
in this lab, change the mode to static and assign the IP address 172.20.20.1 to the management port. 
FortiGate-VM64-KVM #show system interface port1  
config system interface 
    edit "port1" 
        set vdom "root" 
        set mode dhcp 
        set allowaccess ping 
        set type physical 
        set snmp-index 1 
    next 
end

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-14 
 
FortiGate-VM64-KVM #config system interface  
 
FortiGate-VM64-KVM (interface) #edit port1  
FortiGate-VM64-KVM (port1) #set mode static 
FortiGate-VM64-KVM (port1) #set ip 172.20.20.1 255.255.255.0 
FortiGate-VM64-KVM (port1) #set allowaccess ping http https 
FortiGate-VM64-KVM (port1) #end 
To view the graphical user interface, enter "http://172.20.20.1" in OOB-client browser. Please note that activating 
the HTTP service using the highlighted command is required to access the GUI. The end command in FortiGate 
is used to exit the current configuration menu and save any changes, while the next command is used to move 
to the next entry or object within the current configuration menu. 
2.4.2 Getting Familiar with GUI 
After activating the web service, the GUI loads for the first time, as shown in the following figure. The login 
page is displayed, enter your username and password, then press Enter. 
 
Afterward, next figure will appear.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-15 
 
 
Click "Begin," then on the next screen press "Later."

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-16 
 
 
Press Save and continue. 
 
Click OK in the next few figures.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-17 
 
 
 
Upon first login to the FortiGate GUI, the dashboard is displayed. Configuration will be performed step-by-step. 
The FortiGate dashboard provides an overview of system information, licensing, allocated vCPU and RAM. The 
firewall runs firmware v7.4.1 in NAT mode, with an uptime of over 9 days. It has 1% vCPU usage and 55% RAM 
usage.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-18 
 
 
From the left-hand menu, select Network, then navigate to the Interfaces menu, ports, 1 to 7 are visible in the 
figure. We can configure several items such as IP address, subnet mask, role ( WAN, LAN, or undefined), and 
administrative access for each port. Double-click the port1 which was configured in the CLI part. 
 
Port1 configuration is shown in the following figure.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-19 
 
 
As displayed, the addressing mode is Manual. Choose a name for the interface and enter it in the Alias field, 
then click OK. 
As shown in previous figure, administrative access protocols such as PING, HTTPS, SSH,  and HTTP can be 
enabled or disabled on a per-interface basis to allow connections from specific trusted  hosts or subnets. In the 
previous section, we enabled HTTP, PING, and HTTPS access through the command-line interface (CLI).  
For the next step, click Network > Interfaces in the left-hand menu, then double-click Port 2, which is 
assigned to the LAN. Enter the address 192.168.22.1/255.255.255.0 in the IP/Netmask field, as shown in the 
following figure.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-20 
 
 
Different interface roles on a FortiGate device, such as LAN, WAN, or DMZ, dictate which settings are displayed 
on the GUI, and prevent accidental misconfigurations. When an interface is confi gured as WAN, for example, 
settings for a DHCP server and device detection are not available, as these a re typically used for internal LAN 
configurations. If an unusual case requires the use of options hidden by the current role, the role can be switched 
to Undefined which will display all options. 
With all the configuration we have done so far, we can ping sw-core from the firewall using the following 
command. The FortiGate firewall and sw-core can ping each other because they have connected interfaces. 
However, can we ping the clients? 
FortiGate-VM64-KVM #execute ping 192.168.22.2 
PING 192.168.22.2 (192.168.22.2): 56 data bytes 
64 bytes from 192.168.22.2: icmp_seq=0 ttl=255 time=1.2 ms 
64 bytes from 192.168.22.2: icmp_seq=1 ttl=255 time=0.7 ms 
FortiGate-VM64-KVM # execute ping 192.168.33.100 
PING 192.168.33.100 (192.168.33.100): 56 data bytes 
sendto failed

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-21 
 
 
2.4.3 Static Route 
Since the firewall has no route to the core switch, a static route is required. Configure it as shown in the following 
figure.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-22 
 
 
 
Now, the ping is successful. 
2.4.4 Cloud Network 
Based on the topology, we can observe that the firewall is connected to the cloud. For this scenario, we assume 
that the cloud represents the internet. The cloud network has an IP address range  of 172.24.47.0/24, with the 
gateway set to 172.24.47.254.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-23 
 
Ideally, we should assign a public IP address to the internet-facing interface. However, since this is a lab 
environment, we will assume that the firewall is connected to the cloud using a public IP address. 
In fact, since it is a lab environment, we are using a private IP address (172 .24.47.100) to simulate a public IP 
for testing purposes. 
Configure Port 7 as shown in the following figure. 
 
To connect to the cloud, we need a static route like the one shown in the following figure.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-24 
 
 
Now, we can ping the global DNS servers from the FortiGate CLI, confirming Internet access.
