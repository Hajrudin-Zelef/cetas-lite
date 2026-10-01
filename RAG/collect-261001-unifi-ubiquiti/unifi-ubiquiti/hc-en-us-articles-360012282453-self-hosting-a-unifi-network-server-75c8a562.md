---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-360012282453-self-hosting-a-unifi-network-server-75c8a562
title: "hc-en-us-articles-360012282453-self-hosting-a-unifi-network-server-75c8a562"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AMD", "Apple", "Google", "Intel"]
dates: []
keywords: ["amd", "ethernet", "intel"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-360012282453-self-hosting-a-unifi-network-server-75c8a562.md
source_anchor: ""
source_lines: [1, 54]
sha256: 4bdcd569f263fe9553c4a09d95513c7a4e8aa2a7fdf69da74b2f15ffdbe285f6
---

# hc-en-us-articles-360012282453-self-hosting-a-unifi-network-server-75c8a562

(Legacy) Self-Hosting a UniFi Network Server
| Looking for the next generation of UniFi self-hosting? The UniFi OS Server is the new standard for self-hosting UniFi, offering support for advanced features like Organizations, IdP Integration, and Site Magic SD-WAN. Learn more. | 
Self-hosting a UniFi Network Server should only be done by experienced network administrators. For the best and most streamlined experience, we recommend running UniFi Network on a dedicated UniFi Console compatible with UniFi OS. This will eliminate compatibility risks associated with third-party software and hardware.* Other benefits of UniFi Consoles include:
- Optimized hosting and management performance
- Automatic scheduling of backups and updates
- Seamless email and mobile push notifications for important system activity
*UniFi also offers the CloudKey Enterprise (on-premise) and an Official UniFi Hosting subscription service for scaleable deployments managing hundreds of UniFi APs.
If you choose to continue, see below for instructions to install the UniFi Network Application on your device.
Instructions
Windows
Linux
macOS
- Download and install the latest version of the UniFi Network application (UniFi-installer.exe) from the Download page.
  - Note: MongoDB 3.6 is the minimum supported version and is automatically bundled with the download.
- Enable TCP Port 8080, TCP Port 8843, UDP Port 10001, and UDP Port 3478 on any local firewall (including Windows Defender) or antivirus software. See our Required Ports Reference to learn more.
- Start the newly installed UniFi Network application.
- Open a browser, navigate to https://localhost:8443, and continue if you see a security warning.
- Proceed with the setup wizard.
- Adopt your first UniFi device.
Please see the Help Center article here.
- Verify that the Security & Privacy settings on macOS are set to App Store and identified developers.
- Download and install the latest version of the UniFi Network application (UniFi-Network-Server.dmg) from the Download page.
  - Note: MongoDB 3.6 is the minimum supported version and is automatically bundled with the download.
- Open Finder, control-click the UniFi.pkg file and select Open.
- Select Open if you get a malicious software warning.
- If unable to launch UniFi Network, navigate to your Apple menu > System Settings > Privacy & Security and select Open Anyway (you may need to scroll down).
- Follow the UniFi Network application installation process and start it. If prompted, install Rosetta as well.
- Enable TCP Port 8080, TCP Port 8843, UDP Port 10001, and UDP Port 3478 on any local firewall or antivirus software. See our Required Ports Reference to learn more.
- Open a browser, navigate to https://localhost:8443 and proceed when seeing the security warning.
- Proceed with the setup wizard.
- Adopt your first UniFi device.
Running as a Windows Service
Some advanced users may want to run the UniFi Server as a Windows service. To do so, follow the instructions here.
Frequently Asked Questions
However, it is highly recommended to keep the UniFi Network application running at all times. This enables you gather accurate statistics and make configuration changes whenever needed.
If you are not able to keep the self-hosting Network Server running continuously, then we recommend to upgrade to a UniFi Console, which is optimized to run UniFi Network and other applications.
Please be aware that self-hosting the UniFi Network application is an advanced setup for those familiar with network administration. If you are experiencing poor performance, we recommend upgrading to a dedicated UniFi Console compatible with UniFi OS, which is optimized to run UniFi Network and other applications.
If you continue to experience issues, try installing the application on a different machine.
Please be aware that self-hosting the UniFi Network application is an advanced setup for those familiar with network administration. If you are unable to get the UniFi Network application running on your machine, we recommend upgrading to a dedicated UniFi Console compatible with UniFi OS, which is optimized to run UniFi Network and other applications.
Also make sure that the UniFi devices and the application are on the same network. If not, then remote adoption is required.
This can be safely ignored. Proceed to the next page.
At a bare minimum, we recommend the following system requirements (make sure to read the Release Notes for more details about a particular version):
- Operating system:
  - Linux: Ubuntu Desktop / Server 22.04; Debian 11 "Bullseye"
  - Windows: Windows 10; Windows Server 2016
  - macOS: Mavericks 10.9, 10.10 Yosemite, 10.11 El Capitan, 10.12 Sierra, 10.13 High Sierra, 10.14 Mojave, 10.15 Catalina.
- CPU: x86-64 Processor (Intel / AMD x64 Processors)
- RAM: 2GB
- Network: 100Mbps Wired Ethernet
- HDD: Minimum 10GB free (20GB or more preferred)
- Java is no longer required as of UniFi Network 7.5
- Web Browser: Google Chrome
- MongoDB: version 3.6 or later. Mongo is offered bundled: default is 3.6 (for macOS and Windows only).
Note: You will need to continually increase your system specs as you begin to adopt and manage more devices.
