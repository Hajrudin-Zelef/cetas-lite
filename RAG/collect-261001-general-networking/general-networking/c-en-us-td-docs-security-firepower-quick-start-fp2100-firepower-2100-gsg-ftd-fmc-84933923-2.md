---
id: collect-261001-general-networking/general-networking/c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923-2
title: "c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923.md
source_anchor: ""
source_lines: [53, 76]
sha256: 8abf3c67b3b7ce9e23c375e0c130536731d5c590701e61bd0a603b0eec8d4cd1
---

# c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923

Cisco recommends running a Gold Star release indicated by a gold star next to the release number on the software download page. You can also refer to the release strategy described in https://www.cisco.com/c/en/us/products/collateral/security/firewalls/bulletin-c25-743178.html; for example, this bulletin describes short-term release numbering (with the latest features), long-term release numbering (maintenance releases and patches for a longer period of time), or extra long-term release numbering (maintenance releases and patches for the longest period of time, for government certification).
| Step 1 | Connect to the CLI. See Access the Firewall Threat Defense and FXOS CLI for more information. This procedure shows using the console port, but you can use SSH instead. Log in with the admin user and the default password, Admin123. You connect to the FXOS CLI. The first time you log in, you are prompted to change the password. This password is also used for the Firewall Threat Defense login for SSH. Example:  firepower login: admin Password: Admin123 Successful login attempts for user 'admin' : 1  [...]  Hello admin. You must change your password. Enter new password: ******** Confirm new password: ******** Your password was updated successfully.  [...]  firepower#    | 
| Step 2 | At the FXOS CLI, show the running version. scope ssa show app-instance Example:  Firepower# scope ssa Firepower /ssa # show app-instance  Application Name     Slot ID    Admin State     Operational State    Running Version Startup Version Cluster Oper State -------------------- ---------- --------------- -------------------- --------------- --------------- ------------------ ftd                  1          Enabled         Online               7.6.0.65        7.6.0.65        Not Applicable   | 
| Step 3 | If you want to install a new version, perform these steps. | 
| Note | If the password was already changed, and you do not know it, you must perform a factory reset to reset the password to the default. See the FXOS troubleshooting guide for the factory reset procedure. | 
You can complete the Firewall Threat Defense initial configuration using the CLI or Firewall Device Manager.
When you use the Firewall Device Manager for initial setup, the following interfaces are preconfigured in addition to the Management interface and manager access settings. Note that other settings, such as the DHCP server on inside, access control policy, or security zones, are not configured.
Ethernet 1/1—"outside", IP address from DHCP, IPv6 autoconfiguration
Ethernet 1/2— "inside", 192.168.95.1/24
Default route—Obtained through DHCP on the outside interface
If you perform additional interface-specific configuration within Firewall Device Manager before registering with the Firewall Management Center, then that configuration is preserved.
When you use the CLI, only the Management interface and manager access settings are retained (for example, the default inside interface configuration is not retained).
| Step 1 | Log in to the Firewall Device Manager. | 
| Step 2 | Use the setup wizard when you first log into the Firewall Device Manager to complete the initial configuration. You can optionally skip the setup wizard by clicking Skip device setup at the bottom of the page. After you complete the setup wizard, in addition to the default configuraton for the inside interface (Ethernet1/2), you will have configuration for an outside (Ethernet1/1) interface that will be maintained when you switch to Firewall Management Center management. | 
| Step 3 | (Might be required) Configure a static IP address for the Management interface. See the Management interface on . If you want to configure a static IP address, for example for an edge deployment where there is not DHCP server on the network yet, be sure to also set the default gateway to be a unique gateway instead of the data interfaces. If you use DHCP, you do not need to configure anything. | 
| Step 4 | If you want to configure additional interfaces, including an interface other than outside or inside, choose Device, and then click the link in the Interfaces summary. See Configure the Firewall in the Firewall Device Manager for more information about configuring interfaces in the Firewall Device Manager. Other Firewall Device Manager configuration will not be retained when you register the device to the Firewall Management Center. | 
| Step 5 | Choose , and click Proceed to set up the Firewall Management Center management. | 
| Step 6 | Configure the Management Center/Security Cloud Control Details. | 
| Step 7 | Configure the Connectivity Configuration. | 
| Step 8 | Click Connect. The Registration Status dialog box shows the current status of the switch to the Firewall Management Center. After the Saving Management Center/CDO Registration Settings step, go to the Firewall Management Center, and add the firewall. If you want to cancel the switch to the Firewall Management Center, click Cancel Registration. Otherwise, do not close the Firewall Device Manager browser window until after the Saving Management Center/CDO Registration Settings step. If you do, the process will be paused, and will only resume when you reconnect to the Firewall Device Manager. If you remain connected to the Firewall Device Manager after the Saving Management Center/CDO Registration Settings step, you will eventually see the Successful Connection with Management Center/CDO dialog box, after which you will be disconnected from the Firewall Device Manager. | 
Set the Management IP address, gateway, and other basic networking settings using the setup wizard. The dedicated Management interface is a special interface with its own network settings. In 6.7 and later: If you do not want to use the Management interface for the manager access, you can use the CLI to configure a data interface instead. You will also configure the Firewall Management Center communication settings. When you perform initial setup using the Firewall Device Manager (7.1 and later), all interface configuration completed in the Firewall Device Manager is retained when you switch to the Firewall Management Center for management, in addition to the Management interface and manager access interface settings. Note that other default configuration settings, such as the access control policy, are not retained.
| Step 1 | Connect to the Firewall Threat Defense CLI, either from the console port or using SSH to the Management interface, which obtains an IP address from a DHCP server by default. If you intend to change the network settings, we recommend using the console port so you do not get disconnected. The console port connects to the FXOS CLI. The SSH session connects directly to the Firewall Threat Defense CLI. | 
| Step 2 | Log in with the username admin and the password Admin123. At the console port, you connect to the FXOS CLI. The first time you log in to FXOS, you are prompted to change the password. This password is also used for the Firewall Threat Defense login for SSH. Example:  firepower login: admin Password: Admin123 Successful login attempts for user 'admin' : 1  [...]  Hello admin. You must change your password. Enter new password: ******** Confirm new password: ******** Your password was updated successfully.  [...]  firepower#    | 
| Step 3 | If you connected to FXOS on the console port, connect to the Firewall Threat Defense CLI. connect ftd Example:  firepower# connect ftd >   | 
