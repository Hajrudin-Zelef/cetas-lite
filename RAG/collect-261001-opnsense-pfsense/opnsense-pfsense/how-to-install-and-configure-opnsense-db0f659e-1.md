---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-install-and-configure-opnsense-db0f659e-1
title: "how-to-install-and-configure-opnsense-db0f659e"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "ethernet"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-install-and-configure-opnsense-db0f659e.md
source_anchor: ""
source_lines: [1, 30]
sha256: 491fc9f719f66817d67fe2826b2e536a923818588b3f58048b5ea5e396f7642c
---

# how-to-install-and-configure-opnsense-db0f659e

How to Install and Configure OPNsense
Table of Contents
While much of the content on this site focuses on more advanced home networking topics, I thought it may be helpful for new users to write a guide on installing and configuring OPNsense. This information should be a good starting point for beginning the OPNsense journey.
For new users, I often recommend starting with the basics before diving deeper into more complex configurations. Once the basics are well understood, implement one additional feature at a time until you achieve your desired goals. Not only does this allow you to grow in knowledge about each feature, but it also aids with troubleshooting your network. If your network is functioning properly and you try implementing something new which breaks your network, you will know that the configuration for the new feature is the root cause of your issue(s).
If you are interested in reading this guide, I am going to assume that you have some level of understanding of why you want to use OPNsense and you are now ready to try it out. With that said, let us get started!
Choose Your Hardware
The first step is to choose the hardware in which you are running OPNsense. You can either run OPNsense directly on the system (bare metal) or in a virtual machine on a hypervisor such as Proxmox, ESXi, etc. Many home users will either choose a dedicated mini-PC firewall appliance or a virtual machine on a hypervisor.
If you are not familiar with how to set up virtual machines, I recommend you either gain a solid of that first or start with a mini-PC firewall appliance. Mini-PC firewall appliances are compact, power efficient, silent, and powerful systems (compared to many consumer grade routers).
I personally prefer to have a dedicated device rather than using a virtual machine since it allows me to tinker with my Proxmox server without worrying about taking down my network. If you need some insight on the types of hardware options that are available, you may refer to my hardware recommendations page.
Download OPNsense
You may download OPNsense on their download page. I recommend downloading the default “vga” version of the installer since the mini-PC firewalls do not have CD/DVD drives and the installer will also work if you are installing inside a virtual machine. The “vga” installer lets you install the image on a USB drive so you can boot the installer from that drive.
Choose a mirror that is close to your location so the file will download faster and then click “Download”.
Flash USB Drive
Once OPNsense is downloaded, you will need to flash the USB drive with the installer. I prefer to use Etcher because it is a simple tool which works great. One nice feature is that you do not need to extract the image from a compressed file (if it is a format recognized by Etcher), which saves some time and disk space. Simply choose the compressed OPNsense file, the USB drive, and then click “Flash!”.
Install OPNsense
If you are installing OPNsense directly on the system, you may proceed with the installation. I am going to assume your disk is empty or you do not care about erasing the contents of the drive installed on your system.
However, if you are installing in a virtual machine, you will need to prepare the VM before you can continue with the installation of OPNsense. In this guide, I will leave the virtual machine configuration up to the user since many new users will likely be installing to a bare metal system (VM configuration is a topic for another guide). Once the virtual machine is set up, the installation steps are the same as installing on bare metal.
When booting up the OPNsense installer, you will see the default menu below. It only shows for a few seconds and you do not need to enter any options to proceed.
You will be prompted if you want to start the configuration importer. This option is very useful if you wish to restore from a previous configuration backup if your system crashed due to a hardware failure or you decide to reinstall from scratch to return the system to a known working state. Since I am discussing a new installation, wait a few seconds for the installer to continue.
Press any key to manually assign interfaces. You only have a few seconds to hit any key. The reason you should manually assign interfaces is that I have found that it often chooses the incorrect interfaces you wish to use as the WAN or LAN interface. The automatic interface assignment will choose the first interface it encounters in the hardware, which may always not reflect its physical location on your router/firewall device.
If you are familiar with consumer grade routers, the leftmost Ethernet port is often used as the WAN interface so you may want to do the same on your device especially if your ports are not labeled WAN, LAN, etc. Some firewall devices such as Protectli actually have the WAN port on the right side of the device. If your device has labels, you should try to make sure it matches to minimize confusion. However, if your device is simply numbered 1-4, for example, it does not matter which one you use as the WAN interface so you can choose the left or right side or one in the middle if you like living on the edge.
To keep this installation guide more basic, you may enter “N” or press “Enter” since “N” is the default value to skip configuring LAGGs.
You may skip configuring VLANs as well by entering “N” or pressing “Enter” since you may configure them through the web interface after OPNsense installed. Also that would go beyond the basic installation of OPNsense and requires you to have network switches and wireless access points which have VLAN support. If you start with a simple LAN network when you are first learning, you can expand into having multiple networks as you gain more knowledge and experience with networking.
In this step, you should see a list of your network interfaces. Type in the name of the interface to select the WAN interface. On my test system, I have two interfaces named vtnet0 and vtnet1. Since I am using a virtual machine to demonstrate the installation process, I can easily tell which interface is the WAN and which interface is the LAN by the numbers at the end of the interface name. It corresponds with the network adapters net0 and net1 I added in Proxmox.
If you have different network adapters installed in your system, the interfaces may be easier to distinguish based upon the names of the network interfaces. Worst case scenario is that you have to reassign your interfaces later after you finish installing OPNsense.
If you are directly connected to your firewall device with a monitor and keyboard, you can easily reassign interfaces after installation since you do not have to worry about losing network connectivity while making configuration changes to the interfaces.
After selecting the WAN interface, you will need to select the LAN interface. The most basic network for home users will only have a single WAN and LAN interface.
Note
If you need to connect more than one device to your router/firewall device, you will need to use a network switch to plug more devices into it. Unlike a consumer grade router, by default you cannot use all of the extra Ethernet ports as a network switch since the interfaces are treated individually. It is possible to bridge the extra interfaces together so that they act like a network switch, but the packets are routed in software rather than in hardware like a network switch.
You will find that the general recommendation is to avoid bridging interfaces due to decreased network performance under heavy loads for certain hardware configurations. If you are still interested in bridging, you should experiment to see how much performance is decreased when there is heavy traffic on your network before committing to that decision. Otherwise you may be disappointed in performance if your router cannot handle large amounts of network traffic.
