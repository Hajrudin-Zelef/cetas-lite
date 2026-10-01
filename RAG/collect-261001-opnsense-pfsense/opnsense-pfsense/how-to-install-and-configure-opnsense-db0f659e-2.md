---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-install-and-configure-opnsense-db0f659e-2
title: "how-to-install-and-configure-opnsense-db0f659e"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["consumer", "memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-install-and-configure-opnsense-db0f659e.md
source_anchor: ""
source_lines: [31, 62]
sha256: fbb822a434e58567fe842bb1fe71a0afff692154520477fd2d40fc3c75a43c9e
---

# how-to-install-and-configure-opnsense-db0f659e

In this example, I am not going to assign any optional interfaces, but you could assign additional interfaces if you want to create multiple networks. Keep in mind that you can still add more interfaces later using the web interface after installation. You may find it easier to do the bare minimum configuration during the installation and do the rest of your configuration via the web interface.
Enter “Y” to continue with the installation.
You should now see both of your interfaces listed along with their IP addresses. By default, the LAN interface will be assigned the 192.168.1.1/24 network.
The WAN interface will use DHCPv4/DHCPv6 by default and you may see your public IPv4/IPv6 addresses assigned by your ISP. The reason I say ‘you may’ is that there may be some delay from when you unplug your old router and plug in your new router when the new MAC address will be automatically registered with your ISP. Also if you are not plugged directly into your modem during installation, you obviously will not see your public WAN IP address(es) yet.
You may notice the WAN IP address in my screenshot below is an internal network address instead of a public IP address. I have my virtual machine connected to my LAN network so it was automatically assigned an IP address in my LAN network via DHCP. You may be doing something similar if you are testing your OPNsense system behind another router. I generally recommend making OPNsense your primary router but that requires setting your modem/router provided by your ISP to bridge mode or you need to purchase your own modem to use (you may be able to save on rental fees if you do this).
Warning
If you are using a VM on a hypervisor such as Proxmox and you are trying out OPNsense behind another router, be careful which interface you choose for the LAN. If you are using a bridged interface on your same network as your primary router, you may encounter an issue with IP address conflicts if your primary router is using 192.168.1.1/24 since OPNsense will use that network by default on the LAN interface.
In my example, I used a virtual network on Proxmox (a bridge that is not assigned to any physical interfaces) which allows me to put any VM or container behind that virtual network so that I can test out the virtual LAN network on my OPNsense VM without conflicting with my main network.
Enter the username of installer and the password of opnsense in order to continue with the installation. Do not login as root because you will end up running a live version of OPNsense which will not be installed to your system. Live mode is nice if you just want to try out OPNsense without installing anything, but this is guide about installing OPNsense on your system so there is no need to run in live mode.
If you are a US user, you may simply press “Enter”. Otherwise, you will need to select your preferred keyboard layout.
You may choose if you wish to run UFS or ZFS. If you are a novice user, you may prefer not to use ZFS because it is a more advanced filesystem. ZFS is more robust than other filesystems so you may want to use ZFS even if you do not understand anything about ZFS. You may not even notice the difference between the two filesystems if your hardware is functioning properly.
One thing that is important to consider with ZFS is that it requires more RAM than other filesystems so if your system does not have a lot of memory, you should use UFS instead of ZFS.
To keep things simple for this example, I am using the default filesystem of UFS.
Select the disk which you wish to install OPNsense. In many router/firewall devices, there will be only one drive installed so you will only have one choice.
Select “Yes” for the recommended swap partition size. If you run out of system memory, it can lead to crashes so it is typically best to have some swap space. As you can see, the recommended size is relatively small so it should not impact your overall storage capacity by much. OPNsense does not require a lot of disk space unless you are doing a large amount of logging.
Press “Enter” to continue with the installation. I am assuming your disk is blank or you do not care about its contents.
OPNsense should now be installing the system files.
For security purposes, the recommendation is to change the default root user password. You should do this now so you do not forget later. The password can be changed later in the web interface if you decide to change it again.
Enter the new password.
Enter the password again to verify you entered it correctly.
Press “Enter” to exit and reboot your system.
OPNsense is now installed! You can unplug your USB drive or eject your DVD disc depending on the medium used to install OPNsense since you will no longer need it.
At this point if you have done already done so, I recommend you plug a network switch into the LAN port on your OPNsense system and plug at least one PC/laptop into the switch so that you can continue with the OPNsense configuration via the web interface. DHCP should automatically be configured for the LAN network so when you plug into a switch, your system should be able to obtain an IP address like it would with a consumer grade router.
Configure OPNsense
From the system connected to the LAN network of OPNsense, you can access the OPNsense web interface using the default hostname/domain name of the new OPNsense installation: https://opnsense.localdomain (or if you prefer IP addresses, you can use https://192.168.1.1). You should click the “Accept the Risk” prompt since OPNsense is using a self-signed certificate that is generated during the installation.
Login with the root user with the password you set during the installation process.
When you log into the OPNsense web user interface for the first time, you will be prompted to complete a general setup process. While it is not required to complete the wizard, I recommend new users go through the wizard to help guide you through a few basic settings that you may wish to change according to your preferences. Click “Next” to continue.
If you prefer, you may change the “Hostname” of OPNsense to some other name such as “router”.
Likewise, you can change the “localdomain” to some other domain. You can use any domain that is not a real domain name unless you own the domain name. The reason is that it would conflict with the real domain name if you happen to visit the website or any services that use that domain name.
For all of the DNS settings, if you leave everything at the default, your OPNsense installation will behave similar to a consumer grade router. Your ISP DNS servers will be used. That is what the “Override DNS” option does – it will prefer your ISP DNS over any DNS servers you provide. If you wish to use alternate DNS servers such as 1.1.1.1 or 8.8.8.8, you need to uncheck the “Override DNS” option and enter the DNS servers in the “Primary DNS Server” and “Secondary DNS Server” boxes. If you know your ISP or your specified DNS servers support DNSSEC, you can also check the “Enable DNSSEC Support” box (and hardening the DNSSEC data likely is ok to select unless the setting is incompatible with your DNS server).
I would recommend leaving all the DNS settings at the default settings unless you are comfortable changing them and know the impacts of such changes. Once you gain a greater understanding, you can change the DNS servers at a later time. Click “Next” to continue.
The main setting you may want to change on this screen is to set your local timezone. If you prefer to use other time servers, you can replace the default OPNsense timeservers. Click “Next”.
