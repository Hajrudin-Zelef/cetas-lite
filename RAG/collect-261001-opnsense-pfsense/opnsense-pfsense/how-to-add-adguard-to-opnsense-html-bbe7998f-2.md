---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-add-adguard-to-opnsense-html-bbe7998f-2
title: "how-to-add-adguard-to-opnsense-html-bbe7998f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-add-adguard-to-opnsense-html-bbe7998f.md
source_anchor: ""
source_lines: [38, 97]
sha256: 6aa4645ed173f4dfcb014aa56dc044c9b7621ee426227a91acfc749031c360b2
---

# how-to-add-adguard-to-opnsense-html-bbe7998f

By satisfying these prerequisites, you ensure a reliable and secure deployment of AdGuard within your OPNsense environment, enabling effective network-wide ad and malware blocking.
Preparing OPNsense for AdGuard Installation
Before installing AdGuard on OPNsense, proper preparation ensures a smooth setup process. Follow these steps to ready your system:
- Update OPNsense Firmware: Ensure your OPNsense system is running the latest firmware. Navigate to System > Firmware > Check for Updates and apply any available updates. This guarantees compatibility and security.
- Configure Network Interfaces: Identify the interface(s) you want AdGuard to monitor and filter. Usually, this will be the LAN interface. Go to Interfaces > Assignments to verify or set interface roles appropriately.
- Set Up a Dedicated DNS Resolver (Optional but Recommended): For optimal AdGuard performance, it’s advisable to run it as a DNS resolver or proxy. Consider installing the Unbound DNS package if you haven’t already. Navigate to System > Firmware > Available Packages, find unbound, and install it.
- Configure Firewall Rules: Ensure that your firewall rules allow traffic for AdGuard. Create rules permitting DNS (port 53) and any other necessary ports, depending on your configuration, on the LAN interface.
- Backup Configuration: Before making significant changes, back up your current configuration. Head to System > Configuration > Backups and download your current setup.
- Gather Necessary Files and Packages: Prepare the AdGuard binary or Docker image (if you plan to run it in a container). While OPNsense primarily runs FreeBSD, you may need to set up a separate VM or container if native installation isn’t straightforward. Alternatively, ensure you have access to a compatible plugin or script.
With these preparatory steps completed, your OPNsense system will be primed for a seamless AdGuard deployment. Proper setup minimizes troubleshooting and maximizes privacy and filtering efficiency.
Installing AdGuard on OPNsense
Adding AdGuard to OPNsense enhances your network security by providing robust ad blocking and privacy features. Follow these straightforward steps to install and configure AdGuard on your OPNsense device.
Prerequisites
- An active OPNsense installation (version 21.1 or later recommended)
- Administrative access to OPNsense web interface
- Basic knowledge of the command line and package management
Step-by-Step Installation
1. Enable the Package Repository
Ensure your OPNsense is up-to-date and has access to the package repositories. Navigate to System > Firmware > Update and confirm your system is current.
2. Install the pfSense-pkg-AdGuardHome Plugin
- Access the OPNsense shell via SSH or the local console.
- Run the following command to install the plugin:
pkg install os-adguardhome
3. Enable and Configure AdGuard
- Once installed, go to Services > AdGuard Home in the web interface.
- Click Enable AdGuard Home.
- Specify the listen interfaces—typically, localhost or your internal network interfaces.
- Set the DNS port (default 53) and configure upstream DNS servers.
- Adjust filtering options and blocklists as needed for your environment.
Finalizing Setup
After configuration, start the AdGuard service within the web interface. Ensure your network devices use the AdGuard DNS IP address for name resolution, enhancing ad blocking and privacy across your network.
Regularly update AdGuard Home and blocklists to maintain optimal filtering and security.
Step 1: Accessing OPNsense Web Interface
To begin integrating AdGuard with your OPNsense firewall, the first step is accessing the OPNsense web interface. This interface serves as the central hub for configuration and management. Ensure you have network connectivity to your OPNsense device before proceeding.
Recommended Free Tools
Open your preferred web browser. Enter the IP address of your OPNsense device into the address bar. Typically, this IP is set during initial setup, often something like https://192.168.1.1 unless changed. Make sure to include https:// to ensure a secure connection.
If you are on the local network, your browser should establish a connection to the OPNsense web interface. You will be prompted to log in. Enter your administrator username and password. If you haven’t customized these credentials, they are likely still set to defaults such as admin and opnsense. For security reasons, change default login details immediately after initial access.
Once logged in, you will see the OPNsense dashboard. From here, you can navigate to various configuration sections. Since your goal is to add AdGuard, familiarize yourself with the interface. Locate the Services menu on the sidebar, which grants access to various features related to DNS, DHCP, and firewall rules.
Before proceeding with AdGuard installation, it is advisable to verify your network settings—including interface configurations and DNS settings—within the web interface. This ensures a smooth setup process later. Confirm that your OPNsense device’s IP address is static, and note down the current DNS servers, as you may need to modify these during or after the setup.
Quick wins for a faster PC:
Fix the driver behind crashes, sound loss and screen glitchesFind Drivers →Clear out junk files and repair common Windows errorsFree Scan →
Having successfully accessed and logged into the OPNsense web interface, you are now ready to proceed with the installation and configuration of AdGuard. This foundational step ensures you can manage your network effectively and apply the necessary filters seamlessly.
Step 2: Installing Necessary Packages
Before you can integrate AdGuard with OPNsense, you need to install the required packages. These packages form the backbone that enables AdGuard’s filtering capabilities and ensures smooth operation within your firewall environment.
Start by accessing the OPNsense web interface and navigating to System > Firmware > Plugins. This section allows you to manage additional packages not included in the default installation.
- os-dnsspoof: This package helps redirect DNS queries through AdGuard, enabling content filtering at the DNS level.
- os-bind: Provides DNS server functionalities, allowing you to run local DNS services with filtering capabilities.
- os-fqdns: Facilitates advanced DNS filtering features, ensuring that unwanted domains are blocked effectively.
To install these plugins, locate each in the list or use the search function. Click the + icon next to each to initiate installation. Confirm your choices when prompted, and wait for the process to complete. The system may require a reboot to activate these plugins properly.
Do these 3 things before closing this tab:
1Repair Windows errors before they cause bigger problems2Scan for outdated or missing drivers - takes under a minute3Clear out junk files and repair common Windows errors
Additionally, ensure your system is up to date by navigating to System > Firmware > Updates and checking for the latest updates. Keeping your system current prevents compatibility issues and ensures you benefit from the latest security patches.
Once the packages are installed successfully, proceed to configure the DNS and firewall rules in subsequent steps. Proper installation of these components lays the groundwork for an effective AdGuard deployment on your OPNsense platform.
Step 3: Configuring AdGuard
Once AdGuard is installed on your OPNsense system, the next step is to configure it properly to start filtering and blocking unwanted content. Follow these instructions to ensure optimal setup and operation.
Accessing the AdGuard Dashboard
- Open your web browser and navigate to the AdGuard web interface, typically accessible at http://:3000 .
- Login using the credentials created during installation, or default if unchanged.
Initial Configuration
- Within the dashboard, navigate to the Settings tab.
