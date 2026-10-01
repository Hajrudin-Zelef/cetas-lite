---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/therettom-self-hosted-guide-opnsense-3bc340bc-1
title: "Prune the oldest historical environment to control storage footprint"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["ethernet", "full-duplex", "intel", "throughput"]
source: docs/RAG/collect-261001-opnsense-pfsense/therettom-self-hosted-guide-opnsense-3bc340bc.md
source_anchor: ""
source_lines: [1, 41]
sha256: 430996ddec72bb46a84e4a2e41dd16e782f17a0035d65d3d0515eec0a3fa6862
---

# Prune the oldest historical environment to control storage footprint

OPNsense is an open-source, FreeBSD-based firewall and routing platform that provides features like traffic shaping, VPN support, and intrusion detection/prevention. It is designed for home, small business, and enterprise networks, offering a user-friendly web interface and a wide range of functionalities.
If you know little of networking, this whole guide may be a bit over your head. However, I'm going to do my best to aid you. I'm only covering IPv4 and not IPv6 as I haven't been forced to study up on configuring IPv6.
- 
You want to configure this router on an isolated device, as in do not connect to the existing local network. If you do, you'll cause conflicts and break internet connectivity to your whole infrastructure. You can use an ethernet cable from the device you're installing OPNsense on to another device to use the GUI. When it's ready, disconnect the current router and then use the OPNsense router. If you have two Ethernet ports in your device, plug in your secondary device into the second port. OPNsense assumes the numerically lowest port will be WAN.
- 
If your hardware you will be installing OPNsense on does not allow you to have more than one Ethernet port (lots of Intel NUCs and other mini-PCs), you have to do the "Router-on-a-Stick" method. That means you need to use a VLAN-capable switch. USB adapters will not work because they are not reliable. Ask me how I know. Because a single 1 Gbps network card on the firewall must handle both incoming internet traffic and outgoing local traffic simultaneously over the exact same physical wire, the total bandwidth of the port is split down the middle in a worst case scenario. Full-duplex contention will cap your real-world throughput to roughly half of the physical link capacity. Put simply: the single 1 Gbps port causes the bottleneck when sharing simultaneous ingress/egress traffic at 1 Gbps.
A robust guide to installing is already listed in their documentation.
You can download what you need from OPNsense's website.
I highly encourage you to verify the authenticity/integrity.
When you boot the OPNsense image (DVD, VGA, Serial formats), the firewall will be in the live environment with and without the use of OPNsense Importer. You can interact with the environment via console, GUI (HTTPS), or SSH, but for this guide, I'm just explaining it over console.
- To identify your disk to install on, we need to run a command before we install. The default username is root and the default password isopnsense . Press8 for8) Shell , thenEnter and run:
geom disk list
- 
This outputs a detailed block for every detected drive. Look for these specific keys: 
  - 
Geom name : This matches the identifier you will see in the installer (e.g.,nda0 for NVMe,ada0 for SATA,da0 for USB).
  - 
Media size : Shows the capacity (e.g., 250 GB). Match this against the known size of your internal drive.
  - 
Descr : Shows the manufacturer string and model name (e.g., Crucial CT500P3SSD8 or SanDisk Ultra). If you see "Kingston DataTraveler" or "JetFlash," that is most likely your USB installer stick.
- 
- 
Once you identified the correct storage, memorize or write the geom identifier.
- 
To use the installer, either run the command opnsense-installer here, or sign in with theinstaller user. To do that, pressCtrl +D to go back and press0 to log out. The password isopnsense . The installation process involves the following steps:1: Keymap selection: The default configuration should be fine for pretty much everyone using this guide.2: Install (UFS|ZFS): Choose UFS or ZFS filesystem. I highly encourage ZFS. It creates backups with little downsides and is very reliable. It does require at least a few gigabytes of storage, worst case scenario.3: Partitioning (ZFS): Choose a device type. The default option (stripe) is acceptable when using a single disk.4: Disk Selection (ZFS): Select the Storage device. e.g.da0 ornvd0 .5: Last Chance: Select Yes to continue with partitioning and to format the disk. However, doing so will destroy the contents of the disk.6: Continue with recommended swap (UFS): Yes is usually fine here unless the install target is very small (< 16GB)7: Select Root Password: Change and confirm the new root password. If you want to use a password manager, wait to do this until you have access to it. The more complex the password, the better.8: Select Complete Install: Exits the installer and reboots the machine. The system is now installed and ready for initial configuration.
- 
After rebooting into the system (not your USB), the system will prompt you for the interface assignment. If you ignore this, then the default settings will be applied. Installation ends with the login prompt. At this point, you can access the GUI. I recommend you configure things from the GUI from this point unless notated otherwise. The configured IP to access the GUI is 192.168.1.1 . The default username isroot and the default password isopnsense .
If your device has two ethernet ports: Cool, too easy. Interface assignment and configuration should be super easy. If your device only has one: Well this is fun! /s
One Ethernet Port
The GUI is a little wonky. Unless something has changed since version 26.1.5, do it the way I'm guiding you.
- 
Use the CLI and sign in as root . Use1) Assign interfaces . The configuration engine will initialize and scan your hardware.
- 
The wizard will first list your available physical interface names (on Intel NUCs, this is typically an interface name like em0 or igb0). Write this identifier down. This will be the parent device. The wizard will prompt you with the following sequence: 1: Do you want to configure LAGGs now? Typen and hitEnter .2: Do you want to configure VLANs now? Typey and hitEnter .3: Enter the parent interface name for the new VLAN: Type your physical interface name (e.g.,em0 ) and hit Enter.4: Enter the VLAN tag (1-4094): Create your LAN tag (e.g.,20 ) and hit Enter.5: Enter a description for this VLAN (optional): TypeLAN and hitEnter .6: Enter the parent interface name for the new VLAN: Type your physical interface name again (e.g.,em0 ) and hit Enter.7: Enter the VLAN tag (1-4094): Create your ISP/WAN tag (e.g.,10 ) and hitEnter .8: Enter a description for this VLAN (optional): TypeWAN and hitEnter .9: Enter the parent interface name for the new VLAN: Assuming you don't have any other core system trunking to establish right now, leave this line completely blank and hitEnter to finish the VLAN build phase.10: Enter the WAN interface name or 'a' for auto-detect: Type your WAN virtual identifier exactly:em0_vlan10 and hitEnter .11: Enter the LAN interface name or 'a' for auto-detect: Type your LAN virtual identifier exactly:em0_vlan20 and hitEnter .12: Enter the Optional interface 1 name: Leave this blank and hitEnter unless you want to bind additional isolated tags, like for an IoT network.
The screen will output a summary block detailing the pending network transformation:
The interfaces will be assigned as follows:
WAN  -> em0_vlan10
LAN  -> em0_vlan20
Do you want to proceed? [y/n]:
Type y and hit Enter. The console will cycle, tear down the legacy untagged routing tables, spin up the new virtual ones, and reload the core firewall rules.
- 
