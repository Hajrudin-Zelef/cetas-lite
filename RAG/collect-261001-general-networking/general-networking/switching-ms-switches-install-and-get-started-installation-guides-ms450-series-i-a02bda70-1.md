---
id: collect-261001-general-networking/general-networking/switching-ms-switches-install-and-get-started-installation-guides-ms450-series-i-a02bda70-1
title: "switching-ms-switches-install-and-get-started-installation-guides-ms450-series-i-a02bda70"
domain: general-networking
role: reference
task: reference
actors: ["EU", "United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-install-and-get-started-installation-guides-ms450-series-i-a02bda70.md
source_anchor: ""
source_lines: [1, 67]
sha256: 58297e39455bfdd62313511eca9d8c5cb73bc127eacd4c22ddade07d9710a632
---

# switching-ms-switches-install-and-get-started-installation-guides-ms450-series-i-a02bda70

MS450 Series Installation Guide
About this Guide
This guide provides instruction on how to install and configure your MS450 series switch. This guide also provides mounting instructions and limited troubleshooting procedures. For more switch installation guides, refer to the switch installation guides section on our documentation website.
Models
| Model number | Description | 
|---|---|
| MS450-12 | Layer-3 12-port 40GbE QSFP+ aggregation switch with two 100GbE QSFP28 ports and 1 management interface, hot-swappable power supplies / fans | 
Product Overview
Physical Specifications
|  | MS450-12 | 
| 40GbE QSFP+ | 12 | 
| 100Gbe QSFP28 uplink ports | 2 | 
| 100G Hardware Stack Port | 2 | 
| Dedicated Mgmt Interface | 1 | 
| Hot Swap Power Supply | Yes, Dual | 
| Hot Swap Fans | Yes, 3x | 
| Power Input | 100 - 240 VAC, 47-63 Hz | 
| Power Consumption | 47-138W | 
| Operating Temperature | 32°F - 113 °F 0°C - 45 °C | 
| Storage and Transportation Temperature | -4°F - 158°F -20°C - 70°C | 
| Humidity | 5% to 95% | 
| Mounting | 1U Rack Mount | 
Product View and Physical Features
MS450-12 Series front panel   
MS450-12 Series back panel
Ports and Status Indicators
Front Panel
The MS uses LEDs to inform the user of the device's status. When the device powers on, all the Internet LEDs flash twice. Additional functions are described below, from left to right.
| Item | Function | LED Status | Meaning | 
|---|---|---|---|
| 1 | Power | Solid orange | Switch is unable to connect to the Meraki cloud | 
|  |  | Flashing white | Firmware upgrade in process | 
|  |  | Solid white | Switch is fully operational and connected to the Meraki cloud | 
|  |  | Off | Switch does not have power | 
| 2 | Switch Ports | Off | No client connected | 
|  |  | Solid orange | 10/100 Mbps (1000 Mbps on SFP+) | 
|  |  | Solid green | 1000/2500/5000/10000 Mbps (10000 Mbps on SFP+) | 
Back Panel
In addition, there is a RESTORE button available on the back panel.
Insert a paperclip if a restore is required.
- 
    A brief, momentary press: To delete a downloaded configuration and reboot.
- 
    Press and hold for more than 10 sec: To force the unit into a full factory restore.
| Item | Function | LED Status | Meaning | 
|---|---|---|---|
| 1 | Restore | N/A | Restore button to clear switch IP and local configuration settings | 
| 2 | Management Interface | Green | Connected, used for easy access to the local status page | 
| 3 | Stack Ports | N/A | Stack Cables are connected here | 
| 4 | Redundant Fans | Green | Active and operational | 
| 5 | Redundant Power Supplies | Green | Active and functional power supplies | 
Region-specific power cords are not included in the box. Order the appropriate power cord separately:
- MA-PWR-CORD-US
- MA-PWR-CORD-EU
- MA-PWR-CORD-UK
- MA-PWR-CORD-CN
- MA-PWR-CORD-IN
- MA-PWR-CORD-BR
- MA-PWR-CORD-TW
- MA-PWR-CORD-AU
- MA-PWR-CORD-AR
- MA-PWR-CORD-JP
Equipment is to be used only in a restricted access location and installed/operated only by trained service personnel.
Package Contents
In addition to the MS switch, the following are provided (mounting kit provided with 1U models only):
- 4-post Rack Mount Kit includes:
    
