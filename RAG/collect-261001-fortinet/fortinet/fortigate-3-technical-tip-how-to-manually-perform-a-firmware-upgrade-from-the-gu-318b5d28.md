---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-how-to-manually-perform-a-firmware-upgrade-from-the-gu-318b5d28
title: "fortigate-3-technical-tip-how-to-manually-perform-a-firmware-upgrade-from-the-gu-318b5d28"
domain: fortinet
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-how-to-manually-perform-a-firmware-upgrade-from-the-gu-318b5d28.md
source_anchor: ""
source_lines: [1, 115]
sha256: 5f364bc20727ecd618a28cb4799aa057823c089e37bf7ccc31e1c6a7789f17c7
---

# fortigate-3-technical-tip-how-to-manually-perform-a-firmware-upgrade-from-the-gu-318b5d28

**Description**


This article describes how to manually perform a firmware upgrade from GUI and CLI.


**Scope**


From FortiGate v7.2 and above.**Solution**


- FortiGate administrator access is required with Read and Write permission.
- The recent firmware build from the Fortinet Technical Support>Firmware Download .

**Note:** 

Physical presence should be maintained on site while doing the upgrade and the console cable should be kept handy. The upgrade path can be checked from the support.fortinet.com site. Log into **support.fortinet.com -> Support -> Firmware Download -> Upgrade Path**.




- Known Issues can be found in the release note for the Firmware Version and information related to the Firmware can be found from the FortiGuard Site 'https://fortiguard.com/psirt'.
- To download the image files navigate to support portal: **support.fortinet.com -> Support -> Firmware Download -> Download:**


-  Navigate to the required firmware version and locate the image file as the below example.


Example: FortiGate 100F.

- The file name starts with -> FGT_100F.
- Select 'HTTPS' to start the download.

 

- Firmware upgrades are advisable to perform during the maintenance window, with onsite presence.
- If the Firewall is in High Availability then the instruction given at the url given below should be followed.


**Note**: '**pop-ups**' must be allowed by the web browser for the '**HTTPS**' button to work. 


**Related document**:


**Upgrading the firmware using the Firewall GUI manually.**

- Log into the Firewall GUI using super admin privilege.
- Select **System -> Firmware** and select the 'Browse' button to locate the firmware image file.
- Locate the file on the local computer and select the firmware image file.
- Select 'Backup config and upgrade' to back up the configuration and start a firmware upgrade.
- The FortiGate uploads the firmware image file, upgrades to the new firmware version, restarts, and displays the FortiGate login. This process takes a few minutes.


- The FortiGate uploads the firmware image file, upgrades to the new firmware version, restarts, and displays the FortiGate login. This process takes a few minutes.
- If the error 'Something went wrong while displaying this web page. Error Code: Out of Memory.' is seen in Google Chrome after selecting the firmware file, try another browser such as Firefox to proceed.


**Upgrading the firmware using the GUI for v7.4 to v7.6 from the FortiGuard.**

- From v7.4 the GUI interface for Firmware is changed slightly and the Process Given below can be followed to download and install the Firmware from the FortiGuard Servers.
- In the **GUI** , select**System -> Firmware and Registration.**




- If the FortiGate is the Root FortiGate at the Fabric then the 'Upgrade' option needs to be selected.




- Then 'Confirm and Backup Config' is selected, this option will download the Firmware automatically from the FortiGuard Server and the new Firmware will be installed.


**Note**:

The Firewall will reboot during the Firmware Upgrade Process.


**Upgrading the firmware using the CLI.**


- The TFTP server should be running and accessible to the FortiGate. The new firmware image file should be in the root directory of the TFTP server.
- Log into the firewall CLI using a console cable.
- To Verify the TFTP server reach from the FortiGate the 'execute ping' command could be used.
- For example, if the IP address of the TFTP server is 192.168.1.168, run the following:


**execute ping 192.168.1.168**


- Enter the following command to copy the firmware image from the TFTP server to the FortiGate:


**execute restore image tftp <filename> <tftp_ipv4>**


- The FortiGate will respond with the following message:


This operation will replace the current firmware version!

Do you want to continue? (y/n)


- Press 'Y'. The FortiGate will upload the firmware image file, upgrade to the new firmware version, and restart. 
This process takes a few minutes.


- If the Flash needs to be formatted then the article given below can be followed: Technical Tip: Formatting and loading FortiGate firmware image using TFTP

**Related articles**:

Technical Tip: How to upgrade FortiGate firmware
