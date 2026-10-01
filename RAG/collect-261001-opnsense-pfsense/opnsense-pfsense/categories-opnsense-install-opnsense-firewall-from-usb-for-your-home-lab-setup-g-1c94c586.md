---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/categories-opnsense-install-opnsense-firewall-from-usb-for-your-home-lab-setup-g-1c94c586
title: "categories-opnsense-install-opnsense-firewall-from-usb-for-your-home-lab-setup-g-1c94c586"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "ethernet", "memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/categories-opnsense-install-opnsense-firewall-from-usb-for-your-home-lab-setup-g-1c94c586.md
source_anchor: ""
source_lines: [1, 83]
sha256: f2a39624c9abb891ecb6fc9011f20abf74cb6c995d6080d1a7f5f4d86961612e
---

# categories-opnsense-install-opnsense-firewall-from-usb-for-your-home-lab-setup-g-1c94c586

How to install OPNsense firewall from USB for your Home Lab. Setup guide for installing and configuring OPNsense: Assign WAN and LAN interfaces, configure interface IP addresses, set timezone and change the root password
The hardware requirements for running OPNsense firewall can be found in the docs
https://docs.opnsense.org/manual/hardware.html
In this example, I'm installing OPNsense firewall on a refurbished HP T610 Plus Thin Client.
The HP T610 Plus is a reasonably cheap and widely available thin client with enough CPU and memory to run OPNsense firewall in a small network environment or home lab.
I've added an HP NC365T 4 port gigabit ethernet NIC, so I've got extra firewall interfaces.
I'll use the onboard NIC for the LAN interface and port1 of the HP NC365T as the WAN interface.
I've removed the 16GB multi-level cell (MLC) flash drive and installed a 256GB SSD. 60GB or 120GB would also be OK; 256GB is the smallest hard drive I had on hand.
The hard drive is held in place using cable ties; if anyone knows where I can get a 2.5" hard drive bracket for the HP T610 Plus, please let me know in the comments.
AMD G-T56N Processor 1650 Mhz (2 cores)
4GB memory
256GB SSD
HP NC365T 4 port gigabit ethernet NIC
https://opnsense.org/download
Download the amd64 vga image for installing from USB
Extract the iso file from the compressed bz2 file using 7-Zip
OPNsense-22.7-OpenSSL-vga-amd64.img.bz2 > OPNsense-22.7-OpenSSL-vga-amd64.img 
UseRufus to create a bootable USB drive for installing OPNsense
Select the OPNsense .img file
Click start to create the bootable USB
When you boot from the USB drive, the OPNsense live environment will automatically start, and interfaces will also be automatically assigned.
You don't need to change any settings at this stage. Just wait for OPNsense to load.
Login with username installer to start the OPNsense installation
username: installer
password: opnsense
Choose your keyboard layout
Select keymap: United Kingdom
Continue with selected keymap
Install (UFS) GPT/UEFI Hybrid
Select hard drive ada0
Continue with recommended swap partition? Yes
Installation in progress
Complete install
Login with the OPNsense default username and password
username: root
password: opnsense
1 Assign interfaces
Do you want to configure LAGGs now? No
Do you want to configure VLANs now? No
Enter the WAN interface name: igb0
igb0 is port 1 of the HP NC365T network card
bge0 is the onboard NIC
Enter the optional interface name - press [Enter] for none
Interfaces will be assigned as follows:
WAN -> igb0
LAN -> bge0
Do you want to proceed? Yes
2 Set interface IP address
1 LAN (bge0)
Configure IPv4 address LAN interface via DHCP? No
Enter the new LAN ipv4 address: 192.168.1.254
Enter the new LAN IPv4 subnet bit count (subnet mask) 24
Upstream gateway address - for a LAN - press [Enter] for none
Configure IPv6 address LAN interface via WAN tracking? No
Configure IPv6 address LAN interface via DHCP? No
Enter the new LAN IPv6 address - press [Enter] for none
Do you want to enable the DHCP server on LAN ? No
Do you want to change th web GUI protocol to HTTPS? Yes
Restore web GUI access defaults? Yes
You can now access the web GUI by browsing to http://192.168.1.254
Access the web GUI by browsing to http://192.168.1.254
Login with the OPNsense default username and password
username: root
password: opnsense
Click Next to start the setup wizard
General Information
Change the hostname (optional)
Leave the other settings at the defaults
Set the timezone
Configure WAN Interface using DHCP
This will usually be set to DHCP for most home broadband connections with a dynamic WAN IP address
Configure WAN Interface using static IP address
In this example, we're configuring a static WAN IP address
IPv4 configuration: Static
Enter your WAN IP address in CIDR format and upstream gateway IP address
You should get all the IP addressing details from your Internet provider
Leave the other settings as the defaults
Subnet Mask: 24
Change the root password
Click reload to apply the changes
Reload in progress
OPNsense dashboard
Comments
