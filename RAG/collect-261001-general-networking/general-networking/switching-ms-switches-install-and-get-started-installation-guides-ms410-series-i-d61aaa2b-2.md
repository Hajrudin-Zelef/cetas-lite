---
id: collect-261001-general-networking/general-networking/switching-ms-switches-install-and-get-started-installation-guides-ms410-series-i-d61aaa2b-2
title: "switching-ms-switches-install-and-get-started-installation-guides-ms410-series-i-d61aaa2b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["packaging"]
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-install-and-get-started-installation-guides-ms410-series-i-d61aaa2b.md
source_anchor: ""
source_lines: [94, 154]
sha256: b6fc0cefab4bd84fc01e0c61a184c028ec945dc3f19a10083fe15e5b522898a9
---

# switching-ms-switches-install-and-get-started-installation-guides-ms410-series-i-d61aaa2b

- Click on the “Uplink Conﬁguration” tab. Log in. The default login is the serial number (e.g. Qxxx-xxxx-xxxx), with no password (e.g., Q2DD-551C-ZYW3).
- Configure the static IP address, net mask, gateway IP address and DNS servers that this switch will use on its management connection.
- If necessary, reconnect the switch to the LAN.
Static IP via DHCP Reservations
Instead of associating to each Meraki switch individually to configure static IP addresses, an administrator can assign static IP addresses on the upstream DHCP server. Through “DHCP reservations,” IP addresses are “reserved” for the MAC addresses of the Meraki switches. Please consult the documentation for the DHCP server to conﬁgure DHCP reservations.
Installation Instructions
Note: Each switch comes with a graphical instruction pamphlet within the box. This pamphlet contains detailed step by step guides and images to assist in the physical install of the switch.
1. Install the mounting cage nuts in the rack being used for the switch.
2. Attach the switch face plate to the cage nuts on the rack.
3. Insert the power supply unit into the back of the switch. After it has been securely installed, you can connect power to the power supply unit.
Mounting hardware
The mounting hardware includes integrated mounting ears for standard 1U racks. When installing the device, make sure that there is sufficient space between the rear of the rack and other obstacles to ensure adequate airflow.
Bringing your Stack Online
The steps below explain how to prepare a group of switches for physical stacking, how to stack them together, and how to configure the stack in Dashboard.
- Add the switches into a Dashboard network. This can be a new Dashboard network for these switches, or an existing network with other switches. Do not configure the stack in Dashboard yet.
- Connect each switch with individual uplinks to bring them both online and ensure they can check in with the Meraki Dashboard.
- Download the latest firmware build using the Firmware Upgrade Manager under Organization > Monitor > Firmware Upgrades, if they are not already set for this. This helps ensure each switch is running the same firmware build.
- With all switches powered off and links disconnected, connect the switches together via stacking cables in a ring topology (as shown in the following image). To create a full ring, start by connecting switch 1/stack port 1 to switch 2/stack port 2, then switch 2/stack port 1 to switch 3/stack port 2 and so forth, with the bottom switch connecting to the top switch to complete the ring.
- Connect one uplink for the entire switch stack.
- Power on all the switches, then waits several minutes for them to download the latest firmware and updates from Dashboard. The switches may reboot during this process.
    
  - The power LEDs on the front of each switch will blink during this process.
  - Once the switches are done downloading and installing the firmware, their power LEDs will stay solid white or green.
- Navigate to Switch > Monitor > Switch stacks.
- Configure the switch stack in Dashboard. If Dashboard has already detected the correct stack under Detected potential stacks, click Provision this stack to automatically configure the stack.
Otherwise, to configure the stack manually:
- Navigate to Switch > Monitor > Switch stacks.
- Click add one / Add a stack:
- Select the checkboxes of the switches you would like to stack, name the stack, and then click Create.:
- The configuration is complete and the stack should be up and running.
Basic Troubleshooting
The following steps can be used for troubleshooting basic connectivity issues with your switch.
- 
    Reset the switch
- Factory reset the switch by holding the factory reset button for 5 seconds
- Try switching cables, or testing your cable on another device
If your switch still does not connect, the following link may be useful, depending on your issue: Troubleshooting an MS Switch
Reference https://documentation.meraki.com/MS for additional information and troubleshooting tips.
If you are still experiencing hardware issues, please open a case via Meraki dashboard, click the ? icon in the top right of the menu bar and select Support Center. Select the relevant Product or Platform tile, then follow the prompts to enter your case details and submit. For full steps, see Ways to Contact Meraki Support.
Warranty
MS Warranty coverage periods are as follows:
|  | Time Period | Comments | 
| MS410 | Lifetime |  | 
| MS Accessories | 1 Year | The following are considered accessories: SFP Modules, twinax/SFP+ cables, stacking cables, all mounting kits and stands, antennas, interface modules, additional power cords, PoE injectors | 
Note: The above table is a general guideline for warranty terms and is not final. Warranty terms are subject to printed warranty information on the relevant online Meraki data sheets.
If your Cisco Meraki device fails and the problem cannot be resolved by troubleshooting, contact support to address the issue. Once support determines that the device is in a failed state, they can process an RMA and send out a replacement device free of charge. In most circumstances, the RMA will include a pre-paid shipping label so the faulty equipment can be returned.
In order to initiate a hardware replacement for non-functioning hardware that is under warranty, you must have access to the original packaging the hardware was shipped in. The original hardware packaging includes device serial number and order information, and may be required for return shipping.
Meraki MS410 devices have been tested and found to comply with the limits for a Class A digital device, pursuant to part 15 of the FCC rules. A digital device that is marketed for use in a residential environment notwithstanding use in commercial, business and industrial environments.
Additional warranty information can be found on: https://meraki.cisco.com/support#process:warranty
Support and Additional Information
If issues are encountered with device installation or additional help is required, contact Meraki Support by logging in to dashboard.meraki.com and opening a case by visiting the Get Help section.
- The equipment is intended for industrial or other commercial activities.
- The equipment is used in areas without exposure to harmful and dangerous production factors, unless otherwise specified in the operational documentation and/or on the equipment labeling.
- The equipment is not for domestic use. The equipment is intended for operation without the constant presence of maintenance personnel.
- The equipment is subject to installation and maintenance by specialists with the appropriate qualifications, sufficient specialized knowledge, and skills.
- Rules and conditions for the sale of equipment are determined by the terms of contracts concluded by Cisco or authorized Cisco partners with equipment buyers.
- Disposal of a technical device at the end of its service life should be carried out in accordance with the requirements of all state regulations and laws.
- Do not throw in the device with household waste. The technical equipment is subject to storage and disposal in accordance with the organization's disposal procedure.
- The equipment should be stored in its original packaging in a room protected from atmospheric precipitation. The permissible temperature and humidity ranges during storage are specified in the Operation (Installation) Manual.
- Transportation of equipment should be carried out in the original packaging in covered vehicles by any means of transport. The temperature and humidity during transportation must comply with the permissible established ranges of temperature and humidity during storage (in the off state) specified in the Operation Manual (Installation).
For additional information on Meraki hardware and for other installation guides, please refer to documentation.meraki.com.
