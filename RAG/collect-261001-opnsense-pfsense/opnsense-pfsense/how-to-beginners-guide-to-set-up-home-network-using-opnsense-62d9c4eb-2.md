---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb-2
title: "how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["AMD", "Intel", "United States"]
dates: []
keywords: ["amd", "intel", "memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb.md
source_anchor: ""
source_lines: [42, 115]
sha256: 1ba24014602c96ca0a967f6603b788ba1470b75d74d12be24a7c7c6a0daa4562
---

# how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb

To keep this guide shorter, I will not include screenshots in this step like I did with my installation guide.
- Do not press any key when prompted for the configuration importer
- Press any key for manual interface assignment (I prefer to manually configure than using the automatic process)
- Press “Enter” to skip the LAGG configuration
- Press “Enter” to skip creating VLANs (will do this later in the web interface)
- Enter igc0 for the WAN interface name (for the Protectli VP2420 – your interface name may be different such asigb0 )
- Enter igc1 for the LAN interface name (for the Protectli VP2420 – your interface name may be different such asigb1 )
- When prompted for an optional interface name, press “Enter” to skip configuration (will do this later)
- Press “y” and “Enter” to continue
- At the login prompt, enter the username installer and the passwordopnsense to continue with the installation
- Press “Enter” to continue with the default keymap (if you are using the US keyboard, otherwise select the appropriate option)
- Select the “Install (ZFS)” option to use the ZFS filesystem
- Select the disk where you want to install it (you should be able to tell the difference between the USB and internal hard drive based on the name and size)
- Select “Yes” to continue with the recommended swap partition size of 8 GB
- Select “Yes” to continue to destroy the current contents of the disk
- Select “Change root password” now so you do not forget (you will use this to log into the web interface or console)
- Enter the password twice
- Select “Complete Install” to finish installing OPNsense
When it is finished, it will reboot and you should see a login screen which lists the IP addresses of all your interfaces.
Configure OPNsense
Before configuring OPNsense, make sure you have your PC, laptop, or other device you wish to use to configure OPNsense plugged into the LAN interface. In our example, that would be the second interface from the right side of the box.
The LAN interface will be used to configure OPNsense, but once the network switch has been configured, you will plug your network switch into the LAN interface so you can connect more than one device to your network. You will still be able to manage OPNsense if you are connected to any interface on the switch which is set to the default VLAN 1.
For the OPNsense configuration, you will need to login to the default https://192.168.1.1 URL. You could use the default opnsense.local hostname, but if you decide to change the default host/domain name later when configuring OPNsense, you will have to switch to the new hostname in your browser to continue the configuration. If you use the IP address instead, you will not be interrupted due to an OPNsense host/domain name change.
Configuration Wizard
You will be presented with the configuration wizard when you first log into OPNsense. The configuration wizard is completely optional. In this guide, I will skip using the configuration wizard so that you will know where all of the configuration options are located in case you want to make changes in the future. Click on the OPNsense logo in the upper left hand corner of the page to skip the wizard.
System Configuration
The system settings is a good place to start with configuring OPNsense. I will cover the most common settings you will like want to change, but of course if your needs are different, you can deviate from this guide.
Settings: General
On the “System > Settings > General” page, you can customize a few network-wide settings such as the domain name and DNS settings, which can be confusing since there are a few places where you can configure various DNS options. Of course, these settings are free for you to change to your own personal preferences. For instance, I do not expect you to use homenetworkguy.com as your domain name for your network!
| Option | Value | 
|---|---|
| Hostname | router (or choose your own hostname) | 
| Domain | homenetworkguy.com (choose your own domain name) | 
| Time zone | America/New_York (choose your own time zone) | 
| DNS servers | Leave blank | 
| DNS server options | Check “Allow DNS server list to be overridden by DHCP/PPP on WAN” | 
Settings: Administration
The “System > Settings > Administration” page contains useful configuration for how you wish to access OPNsense (web, SSH, and console):
Web GUI:
| Option | Value | 
|---|---|
| Protocol | HTTPS (should be the default value) | 
| HTTP Strict Transport Security | Check “Enable HTTP Strict Transport Security” | 
| TCP port | Leave as 443 | 
| HTTP Redirect | Leave unchecked | 
| DNS Rebind Check | Leave unchecked for security (may interfere with certain name resolutions) | 
| HTTP Compression | High (if you have a reasonably fast system or “Low” otherwise) | 
| Listen Interfaces | Leave at “All (recommmended)” | 
You may configure the SSH/console access but for simplicity for beginners, I am not including those extra settings since most beginners will likely be accessing OPNsense via the web interface.
Settings: Miscellaneous
The “System > Settings > Miscellaneous” page has a few options you may want to tweak such as the CPU type for the thermal sensors widget on the “Dashboard”. There are some periodic backup options, power savings options, and memory/swap options.
| Option | Value | 
|---|---|
| Thermal Sensors Hardware | Intel Core CPU (unless you have AMD hardware) | 
| Periodic RRD Backup | 24 hours (optional) | 
| Periodic DHCP Leases Backup | 24 hours (optional) | 
| Periodic NetFlow Backup | 24 hours (optional) | 
| Use PowerD | Checked (if you have power saving options enabled in the BIOS) | 
| Power Mode | Hiadaptive (to favor performance over power savings) | 
If you are using a SSD or a traditional hard disk, you should not need to adjust any of the disk/memory settings at the bottom of the page since those options are more designed for systems where you want to minimize wear on the disk or if disk space is very constrained. Modern SSDs can handle a lot of writes before the disks wear out.
Interface Configuration
Next up is configuring the interfaces. For this simplified guide which does not include LAG configuration, the only interface you will need to create is a VLAN interface since the WAN and LAN interfaces were created during the initial installation.
Interfaces: Settings
| Option | Value | 
|---|---|
| Hardware CRC | Check “Disable hardware checksum offload” (if not already checked) | 
| Hardware TSO | Check “Disable hardware TCP segmentation offload” (if not already checked) | 
| Hardware LRO | Check “Disable hardware large receive offload” (if not already checked) | 
| VLAN Hardware Filtering | Choose the “Disable VLAN Hardware Filtering” option | 
I have often seen the recommendation to disable hardware offloading on the network interfaces due to various issues that may be encountered. For most users, it is always best to leave it off unless you have thoroughly tested out these configuration options.
In a home network, hardware offloading (if it works) is likely less impactful than using it on a heavily saturated business or enterprise network unless you regularly saturate your network’s bandwidth.
If you are using IDS/IPS services such as Suricata or Zenarmor, hardware offloading should be disabled since it is incompatible with netmap.
Other Types: VLAN
A VLAN is a virtual network that resides on top of a physical network. VLANs allow you to create multiple virtual networks on physical hardware so that you have the flexibility of separating network traffic (each have their own broadcast domains) without needing to purchase additional hardware for each physical network.
