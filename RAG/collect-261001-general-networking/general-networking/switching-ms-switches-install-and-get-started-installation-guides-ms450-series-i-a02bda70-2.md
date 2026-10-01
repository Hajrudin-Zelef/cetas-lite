---
id: collect-261001-general-networking/general-networking/switching-ms-switches-install-and-get-started-installation-guides-ms450-series-i-a02bda70-2
title: "switching-ms-switches-install-and-get-started-installation-guides-ms450-series-i-a02bda70"
domain: general-networking
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["license", "packaging"]
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-install-and-get-started-installation-guides-ms450-series-i-a02bda70.md
source_anchor: ""
source_lines: [68, 135]
sha256: 0fd73db61c49c90abe71b76a982155aa197b5fa0ae59ef808d563264393a1fef
---

# switching-ms-switches-install-and-get-started-installation-guides-ms450-series-i-a02bda70

  - US 12-24 mounting screws and cage nuts, 10 of each
  - INTL M5 mounting screws and cage nuts, 10 of each
  - INTL M6 mounting screws and cage nuts, 10 of each
  - Mounting washers
  - 2 Rack mount rails
  - Rail kit screws
- 250WAC Power Supply Unit
- 3 Pre-installed Fans
Note: The MS450 does not include stacking cables. Stacking cables sold separately.
The MS450 series will ship with all fans and a single power supply included, additional accessories including spare fans and power supplies can be purchased separately.
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
- Click on the “Uplink Conﬁguration” tab. Log in. The default login is the serial number (e.g. Qxxx-xxxx-xxxx), with no password (e.g., Q2DD-551C-ZYW3).
- Configure the static IP address, net mask, gateway IP address and DNS servers that this switch will use on its management connection.
- If necessary, reconnect the switch to the LAN.
Static IP via DHCP Reservations
Instead of associating to each Meraki switch individually to configure static IP addresses, an administrator can assign static IP addresses on the upstream DHCP server. Through “DHCP reservations,” IP addresses are “reserved” for the MAC addresses of the Meraki switches. Please consult the documentation for the DHCP server to conﬁgure DHCP reservations.
Installation Instructions
Note: Each MS450 comes with an instruction pamphlet within the box. This pamphlet contains detailed step by step guides and images to assist in the physical install of the switch.
1. Install the mounting cage nuts in the rack being used for the switch.
2. Separate the rack mount rails, and install the rack mount rails channel onto the rack.
3. Attach the rack mount rail to the sides of the switch.
4. Insert the rack mount rail into the rack mount rail channel.
5. Attach the switch face plate to the cage nuts on the rack.
6. Secure the rack mount rail to the rack mount rail channel.
7. Insert the power supply unit into the back of the switch. After it has been securely installed, you can connect power to the power supply unit.
8. (Optional) Install additional SFP+, QSFP+ or QSFP28 units as needed, depending on the compatibility of your model.
Mounting hardware
The mounting hardware includes a rack mount kit for standard 1U racks. When installing the device, make sure that there is sufficient space between the rear of the rack and other obstacles to ensure adequate airflow.
Warranty
MS Warranty coverage periods are as follows:
|  | Tme Period | Comments | 
| MS450 | Lifetime |  | 
| MS Accessories | 1 Year | The following are considered accessories: SFP Modules, twinax/SFP+ cables, stacking cables, all mounting kits and stands, antennas, interface modules, additional power cords, PoE injectors | 
Note: The above table is a general guideline for warranty terms and is not final. Warranty terms are subject to printed warranty information on the relevant online Meraki data sheets.
If your Cisco Meraki device fails and the problem cannot be resolved by troubleshooting, contact support to address the issue. Once support determines that the device is in a failed state, they can process an RMA and send out a replacement device free of charge. In most circumstances, the RMA will include a pre-paid shipping label so the faulty equipment can be returned.
In order to initiate a hardware replacement for non-functioning hardware that is under warranty, you must have access to the original packaging the hardware was shipped in. The original hardware packaging includes device serial number and order information, and may be required for return shipping.
Meraki MS450 devices have been tested and found to comply with the limits for a Class A digital device, pursuant to part 15 of the FCC rules. A digital device that is marketed for use in a residential environment notwithstanding use in commercial, business and industrial environments.
Additional warranty information can be found on: https://meraki.cisco.com/support#process:warranty
Support and Additional Information
