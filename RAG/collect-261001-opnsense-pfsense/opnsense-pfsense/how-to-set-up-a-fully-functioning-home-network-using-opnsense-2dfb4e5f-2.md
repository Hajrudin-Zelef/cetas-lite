---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f-2
title: "how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f.md
source_anchor: ""
source_lines: [49, 124]
sha256: ad1df824d1bd03991cc11cd15aaf747ba5354879423f4149d581ccae158adbbb
---

# how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f

To keep this guide shorter, I will not include screenshots in this step like I did with my installation guide.
- Do not press any key when prompted for the configuration importer
- Press any key for manual interface assignment (I prefer to manually configure than using the automatic process)
- Press “Enter” to skip the LAG configuration (will do this later in the web interface)
- Press “Enter” to skip creating VLANs (also will do this later)
- Enter igb0 for the WAN interface name (for the Protectli VP2410 – your interface names may be different)
- Enter igb1 for the LAN interface name (for the Protectli VP2410)
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
When it is finished, it will reboot and you should see a login screen with the IP addresses of your interfaces.
Configure OPNsense
Before configuring OPNsense, make sure you have your PC, laptop, or other device you wish to use to configure OPNsense plugged into the LAN interface. In our example, that would be the second interface from the right side of the box. The LAN interface will be used as the dedicated interface to manage OPNsense and other network infrastructure.
Warning
If you are setting up OPNsense on new hardware and you are connecting the WAN interface to your existing home network, you need to be careful when configuring your interfaces.
One issue I accidentally encountered is configuring an interface which has the same IP address range of the network where the WAN interface is attached. I was testing configuration similar to my main network and did not think about this issue. This caused an issue with DNS because the WAN of my test system was using the DNS server (192.168.20.1) of the VLAN on my main network and one of my VLANs on my test system was set up to use the 192.168.20.1/24 network. The local interfaces take priority so the WAN DNS no longer worked properly.
For the OPNsense configuration, you will need to login to the default https://192.168.1.1 URL. You could use the default opnsense.local domain, but if you decide to change the host/domain name later when configuring OPNsense, you will have to switch to the new hostname in your browser to continue the configuration. If you use the IP address instead, you will not be interrupted due to an OPNsense host/domain name change.
Configuration Wizard
You will be presented with the configuration wizard when you first log into OPNsense. The configuration wizard is completely optional. In this guide, I will skip using the configuration wizard so that you will know where all of the configuration options are located in case you want to make changes in the future. Click on the OPNsense logo in the upper left hand corner of the page to skip the wizard.
System Configuration
The system settings is a good place to start with configuring OPNsense. I will cover the most common settings you will like want to change, but of course if your needs are different, you can deviate from this guide.
Settings: General
On the “System > Settings > General” page, you can customize a few network-wide settings such as the domain name and DNS settings, which can be confusing since there are a few places where you can configure various DNS options. Of course, these settings are free for you to change to your own personal preferences. I do not expect you to use homenetworkguy.com as your domain name, for instance!
| Option | Value | 
|---|---|
| Hostname | router (or choose your own hostname) | 
| Domain | homenetworkguy.com (choose your own domain name) | 
| Time zone | America/New_York (choose your own time zone) | 
| DNS servers | Leave blank | 
| DNS server options | Leave both boxes unchecked | 
Note
If you are unchecking the default DNS server options and you have your OPNsense router plugged into your main router, DNS may not work until properly until you plug your new OPNsense router into your modem as your main router since I am assuming this router will be on the edge of your network when it is fully configured.
Alternatively, you could leave the default DNS server options in the “General Settings” and enable “Use System Nameservers” on the “Services > Unbound DNS > Query Forwarding” page so that DNS queries are forwarded to your main network’s DNS server. You will have to remember to change them back later once you use the newly configured OPNsense box as your primary router.
Settings: Administration
The “System > Settings > Administration” page contains useful configuration for how you wish to access OPNsense (web, SSH, and console):
Web GUI:
| Option | Value | 
|---|---|
| Protocol | HTTPS (should be the default value) | 
| HTTP Strict Transport Security | Check “Enable HTTP Strict Transport Security” | 
| TCP port | Change to 8443 if exposing an internal web server to the Internet (securely, of course) | 
| HTTP Redirect | Leave unchecked unless you are exposing port 80/HTTP to the Internet for the purpose of redirecting to 443/HTTPS | 
| DNS Rebind Check | Leave unchecked for security (may interfere with certain name resolutions) | 
| HTTP Compression | High (if you have a reasonably fast system) | 
| Listen Interfaces | Leave at “All (recommmended)” | 
Secure Shell:
| Option | Value | 
|---|---|
| Secure Shell Server | Check “Enable Secure Shell” (to allow access to OPNsense via SSH) | 
| Root Login | Check “Permit root user login” (since only SSH keys are allowed via the next option below, allowing the root login is much more secure, but you may disable this for extra security) | 
| Authentication Method | Uncheck “Permit password login” to improve security (only SSH keys are allowed) | 
| SSH port | Leave blank for default port of 22 | 
| Listen Interfaces | Leave at “All (recommmended)” | 
Console:
| Option | Value | 
|---|---|
| Console driver | Check “Use the virtual terminal driver (vt)” | 
| Primary Console | Select “VGA Console” (if you plan to connect to a monitor, which I recommend leaving enabled) | 
| Secondary Console | Select “Serial Console” (if your device has a serial console) | 
| Serial Speed | Most likely the default of 115200 should be ok (but consult your user manual) | 
| USB-based serial | Check “Use USB-based serial ports” (if your serial console uses USB) | 
| Console menu | Check “Password protect the console menu” to add more protection to console access | 
Authentication:
| Option | Value | 
|---|---|
| Server | Leave at “Local Database” unless you are using multi-factor authentication | 
| Sudo | Choose “Ask password” if you wish to allow other admin accounts you create to have sudo privileges for higher level system access | 
| User OTP seed | You can leave this option as “Nothing selected” even with multi-factor authentication enabled if you do not want users to generate their own codes (if you control all user accounts this option is not needed) | 
