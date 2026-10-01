---
id: collect-261001-general-networking/general-networking/switching-ms-switches-install-and-get-started-installation-guides-ms410-series-i-d61aaa2b-1
title: "switching-ms-switches-install-and-get-started-installation-guides-ms410-series-i-d61aaa2b"
domain: general-networking
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-install-and-get-started-installation-guides-ms410-series-i-d61aaa2b.md
source_anchor: ""
source_lines: [1, 93]
sha256: cc188839fea182658aa83bf56b55c59a4c64d020755b81a0821b6c47eeefa55a
---

# switching-ms-switches-install-and-get-started-installation-guides-ms410-series-i-d61aaa2b

MS410 Series Installation Guide
About this Guide
This guide provides instruction on how to install and configure your MS410 series switch. This guide also provides mounting instructions and limited troubleshooting procedures. For more switch installation guides, refer to the switch installation guides section on our documentation website.
Product Overview
Models
| Model number | Description | 
|---|---|
| MS410-16 | Layer-3 16-port 1Gbe SFP and 2-port 10Gbe SFP+ aggregation switch with 1 management interface and hot-swappable power supplies / fans | 
| MS410-32 | Layer-3 32-port 1Gbe SFP and 4-port 10Gbe SFP+ aggregation switch with 1 management interface and hot-swappable power supplies / fans | 
Physical Specifications
|  | MS410-16 | MS410-32 | 
| 1GbE SFP | 16 | 32 | 
| 10GbE SFP+ uplink | 2 | 4 | 
| 40G QSFP+ Stacking ports | 2 | 2 | 
| Dedicated Mgmt Interface | 1 | 1 | 
| Hot Swap Power Supply | Yes, Dual | Yes, Dual | 
| Hot Swap Fans | Yes, 2x | Yes, 2x | 
| Power Input | 100 - 240 VAC, 47-63 Hz | 100 - 240 VAC, 47-63 Hz | 
| Power Consumption | 50-85W | 50-85W | 
| Operating Temperature | 32°F - 104 °F 0°C - 40 °C | 32°F - 104 °F 0°C - 40 °C | 
| Storage and Transportation Temperature | -4°F - 158°F -20°C - 70°C | -4°F - 158°F -20°C - 70°C | 
| Humidity | 5% to 95% | 5% to 95% | 
| Mounting | 1U Rack Mount | 1U Rack Mount | 
Product View and Physical Features
MS410-16 Series front panel
MS410-32 Series front panel
In addition, there is a RESTORE button available on the front panel.
Insert a paper clip if a restore is required.
- A brief, momentary press: To delete a downloaded configuration and reboot.
- Press and hold for more than 10 sec: To force the unit into a full factory restore.
Ports and Status Indicators
The MS uses LEDs to inform the user of the device's status. When the device powers on, the main LED will be amber. Additional functions are described below, from left to right.
| Item | Function | LED Status | Meaning | 
| 1 | Restore | N/A | Restore button to clear switch IP and local configuration settings | 
| 2 | Fan | Orange | One or more of the system fans has malfunctioned or is missing | 
| 3 | Power (PSU) | Orange | One of the system PSUs (power supplies) has malfunctioned or is missing | 
| 2 | Status | Orange | Switch is unable to or has not yet connected to the Meraki cloud | 
|  |  | Flashing green | Firmware upgrade in process | 
|  |  | White | Switch is fully operational and connected to the Meraki cloud | 
|  |  | Rainbow | Switch is booting, searching for uplink to Meraki Cloud | 
|  |  | Off | Switch does not have power | 
| 3 | Switch Port LEDs | Off | No client connected | 
|  |  | Solid orange | 1 Gbps on SFP+ | 
|  |  | Solid green | 1 Gbps on SFP/10 Gbps on SFP+ | 
Factory Reset Button
If the button is pressed and held for at least 10 seconds and then released, the switch will reboot and be restored to its original factory settings by deleting all configuration information stored on the unit.
Insert a paper clip if a restore is required.
- A brief, momentary press: To delete a downloaded configuration and reboot.
- Press and hold for more than 10 sec: To force the unit into a full factory restore.
Back Panel
| Power input | Power cords may be ordered separately. | 
| Function | LED Status | Meaning | 
| Fan | Orange | One or more of the system fans has malfunctioned or is missing | 
| Power (PSU) | Orange | One of the system PSUs (power supplies) has malfunctioned or is missing | 
| Management Port | Green | Connected, used for easy access to the local status page | 
Equipment is to be used only in a restricted access location and installed/operated only by trained service personnel.
Package contents
In addition to the MS switch, the following are provided (mounting kit provided with 1U models only):
- US 12-24 mounting screws and cage nuts, 5 of each
- INTL M5 mounting screws and cage nuts, 5 of each
- INTL M6 mounting screws and cage nuts, 5 of each
- Mounting washers
The MS410 series will ship with all fans and a single power supply included, additional accessories including spare fans and power supplies can be purchased separately.
Safety and Warnings
These operations are to be taken with respect to all local laws. Please take the following into consideration for safe operation:
- Power off the unit before you begin. Read the installation instructions before connecting the system to the power source.
- Before you work on any equipment, be aware of the hazards involved with electrical circuitry and be familiar with standard practices for preventing accidents.
- Read the mounting instructions carefully before beginning installation. Failure to use the correct hardware or to follow the correct procedures could result in a hazardous situation to people and damage to the system.
- This product relies on the building’s installation for short-circuit (overcurrent) protection. Ensure that the protective device is rated not greater than: 15 A, 125 Vac, or 10A, 240 Vac.
- Please only power the device with the provided power cables to ensure regulatory compliance.
Pre-install Preparation
You should complete the following steps before going on-site to perform an installation.
Configure your Dashboard Network
The following is a brief overview only of the steps required to add a switch to your network. For detailed instructions about creating, configuring and managing Meraki networks, refer to the online documentation (documentation.meraki.com).
- Login to http://dashboard.meraki.com. If this is your ﬁrst time, create a new account.
- Find the network to which you plan to add your switches or create a new network.
- Add your switches to your network. You will need your Meraki order number (found on your invoice) or the serial number of each switch, which looks like Qxxx-xxxx-xxxx, and is found on the bottom of the unit. You will also need your Enterprise license key, which you should have received via email.
- Go to the map / ﬂoor plan view and place each switch on the map by clicking and dragging it to the location where you plan to mount it.
Check and Set Firmware
To ensure your switch performs optimally immediately following installation, it is recommended that you facilitate a ﬁrmware upgrade prior to mounting your switch.
- Attach your switch to power and a wired Internet connection.
- The switch will turn on and the power LED will glow solid orange.
- If the unit requires an upgrade, the power LED will begin blinking white until the upgrade is complete, at which point the LED will turn solid white. You should allow at least a few minutes for the ﬁrmware upgrade to complete, depending on the speed of your internet connection.
Check and Configure Upstream Firewall Settings
If a ﬁrewall is in place, it must allow outgoing connections on particular ports to particular IP addresses. The most current list of outbound ports and IP addresses for your particular organization can be found on the firewall configuration page in your dashboard.
Assigning an IP Address
All switches must be assigned routable IP addresses. These IP addresses can be dynamically assigned via DHCP or statically assigned.
Dynamic Assignment
When using DHCP, the DHCP server should be configured to assign a static IP address for each MAC address belonging to a Meraki switch. Other features of the network, such as 802.1X authentication, may rely on the property that the switches have static IP addresses.
Static Assignment
Static IPs are assigned using the local web server on each switch. The following procedure describes how to set the static IP:
- Using a client machine (e.g., a laptop), connect to the switch over a wired connection.
- Using a web browser on the client machine, access the switch’s built-in web server by browsing to http://my.meraki.com. Alternatively, browse to http://1.1.1.100
