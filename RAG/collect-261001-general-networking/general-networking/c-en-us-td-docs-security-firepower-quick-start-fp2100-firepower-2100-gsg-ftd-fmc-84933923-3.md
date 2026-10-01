---
id: collect-261001-general-networking/general-networking/c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923-3
title: "c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "licenses", "parameters"]
source: docs/RAG/collect-261001-general-networking/c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923.md
source_anchor: ""
source_lines: [77, 110]
sha256: 35b2e977a4d2acdc9b01333b45288b35b9a163cefc4ac43ad397ba4850ceca94
---

# c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923

| Step 4 | The first time you log in to the Firewall Threat Defense, you are prompted to accept the End User License Agreement (EULA) and, if using an SSH connection, to change the admin password. You are then presented with the CLI setup script. Defaults or previously entered values appear in brackets. To accept previously entered values, press Enter. See the following guidelines:  Example:  You must accept the EULA to continue. Press <ENTER> to display the EULA: End User License Agreement [...]  Please enter 'YES' or press <ENTER> to AGREE to the EULA:   System initialization in progress.  Please stand by. You must change the password for 'admin' to continue. Enter new password: ******** Confirm new password: ******** You must configure the network to continue. Configure at least one of IPv4 or IPv6 unless managing via data interfaces. Do you want to configure IPv4? (y/n) [y]: Do you want to configure IPv6? (y/n) [y]:n Configure IPv4 via DHCP or manually? (dhcp/manual) [manual]: Enter an IPv4 address for the management interface [192.168.45.45]: 10.10.10.15 Enter an IPv4 netmask for the management interface [255.255.255.0]: 255.255.255.192 Enter the IPv4 default gateway for the management interface [data-interfaces]: 10.10.10.1 Enter a fully qualified hostname for this system [firepower]: ftd-1.cisco.com Enter a comma-separated list of DNS servers or 'none' [208.67.222.222,208.67.220.220,2620:119:35::35]: Enter a comma-separated list of search domains or 'none' []:cisco.com If your networking information has changed, you will need to reconnect. Disabling IPv6 configuration: management0 Setting DNS servers: 208.67.222.222,208.67.220.220,2620:119:35::35 Setting DNS domains:cisco.com Setting hostname as ftd-1.cisco.com Setting static IPv4: 10.10.10.15 netmask: 255.255.255.192 gateway: 10.10.10.1 on management0 Updating routing tables, please wait... All configurations applied to the system. Took 3 Seconds. Saving a copy of running network configuration to local disk. For HTTP Proxy configuration, run 'configure network http-proxy'  Manage the device locally? (yes/no) [yes]: no DHCP server is already disabled DHCP Server Disabled Configure firewall mode? (routed/transparent) [routed]: Configuring firewall mode ...   Device is in OffBox mode - disabling/removing port 443 from iptables. Update policy deployment information     - add device configuration     - add network discovery     - add system policy  You can register the sensor to a Firepower Management Center and use the Firepower Management Center to manage it. Note that registering the sensor to a Firepower Management Center disables on-sensor Firepower Services management capabilities.  When registering the sensor to a Firepower Management Center, a unique alphanumeric registration key is always required.  In most cases, to register a sensor to a Firepower Management Center, you must provide the hostname or the IP address along with the registration key. 'configure manager add [hostname \| ip address ] [registration key ]'  However, if the sensor and the Firepower Management Center are separated by a NAT device, you must enter a unique NAT ID, along with the unique registration key. 'configure manager add DONTRESOLVE [registration key ] [ NAT ID ]'  Later, using the web interface on the Firepower Management Center, you must use the same registration key and, if necessary, the same NAT ID when you add this sensor to the Firepower Management Center. >  | 
| Step 5 | Identify the Firewall Management Center that will manage this Firewall Threat Defense. configure manager add {hostname \| IPv4_address \| IPv6_address \| DONTRESOLVE} reg_key [nat_id]  Example:  > configure manager add MC.example.com 123456 Manager successfully configured.If the Firewall Management Center is behind a NAT device, enter a unique NAT ID along with the registration key, and specify DONTRESOLVE instead of the hostname, for example: Example:  > configure manager add DONTRESOLVE regk3y78 natid90 Manager successfully configured.If the Firewall Threat Defense is behind a NAT device, enter a unique NAT ID along with the Firewall Management Center IP address or hostname, for example: Example:  > configure manager add 10.70.45.5 regk3y78 natid56 Manager successfully configured. | 
| Note | If the password was already changed, and you do not know it, you must reimage the device to reset the password to the default. See the FXOS troubleshooting guide for the reimage procedure. | 
| Note | You cannot repeat the CLI setup wizard unless you clear the configuration; for example, by reimaging. However, all of these settings can be changed later at the CLI using configure network commands. See Cisco Secure Firewall Threat Defense Command Reference. | 
Register your firewall to the Firewall Management Center.
Use the Firewall Management Center to configure and monitor the Firewall Threat Defense.
| Step 1 | Using a supported browser, enter the following URL. https://fmc_ip_address | 
| Step 2 | Enter your username and password. | 
| Step 3 | Click Log In. | 
All licenses are supplied to the Firewall Threat Defense by the Firewall Management Center. You can purchase the following licenses:
Essentials—Required
IPS
Malware Defense
URL Filtering
Cisco Secure Client
For a more detailed overview on Cisco Licensing, go to cisco.com/go/licensingguide
When you bought your device from Cisco or a reseller, your licenses should have been linked to your Smart Software License account. If you don't have an account on the Smart Software Manager, click the link to set up a new account.
If you have not already done so, register the Firewall Management Center with the Smart Software Manager. Registering requires you to generate a registration token in the Smart Software Manager. See the Cisco Secure Firewall Management Center Administration Guide for detailed instructions.
| Step 1 | Make sure your Smart Licensing account contains the available licenses you need. If you need to add licenses yourself, go to Cisco Commerce Workspace and use the Search All field. Choose Products & Services from the results. Search for the following license PIDs:  | 
| Step 2 | If you have not already done so, register the Firewall Management Center with the Smart Licensing server. Registering requires you to generate a registration token in the Smart Software Manager. See the Cisco Secure Firewall Management Center Administration Guide for detailed instructions. | 
| Note | If a PID is not found, you can add the PID manually to your order. | 
Register the Firewall Threat Defense to the Firewall Management Center manually using the device IP address or hostname.
| Step 1 | In the Firewall Management Center, choose . | 
| Step 2 | From the Add drop-down list, choose Add Device. Set the following parameters:  | 
| Step 3 | Click Register, and confirm a successful registration. If the registration succeeds, the device is added to the list. If it fails, you will see an error message. If the Firewall Threat Defense fails to register, check the following items:  For more troubleshooting information, see https://cisco.com/go/fmc-reg-error. | 
This section describes how to configure a basic security policy with the following settings:
Inside and outside interfaces—Assign a static IP address to the inside interface, and use DHCP for the outside interface.
DHCP server—Use a DHCP server on the inside interface for clients.
Default route—Add a default route through the outside interface.
NAT—Use interface PAT on the outside interface.
Access control—Allow traffic from inside to outside.
To configure a basic security policy, complete the following tasks.
When you use the Firewall Device Manager for initial setup, the following interfaces are preconfigured:
If you performed additional interface-specific configuration within Firewall Device Manager before registering with the Firewall Management Center, then that configuration is preserved.
