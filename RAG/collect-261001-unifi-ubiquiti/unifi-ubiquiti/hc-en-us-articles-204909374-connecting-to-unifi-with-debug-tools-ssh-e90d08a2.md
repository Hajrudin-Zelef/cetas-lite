---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-204909374-connecting-to-unifi-with-debug-tools-ssh-e90d08a2
title: "hc-en-us-articles-204909374-connecting-to-unifi-with-debug-tools-ssh-e90d08a2"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-204909374-connecting-to-unifi-with-debug-tools-ssh-e90d08a2.md
source_anchor: ""
source_lines: [1, 26]
sha256: 92fb55213a3fc3f3857887c4e4c77b834f89073a429b881ceaf329bed24f53b2
---

# hc-en-us-articles-204909374-connecting-to-unifi-with-debug-tools-ssh-e90d08a2

Connecting to UniFi with Debug Tools & SSH
Users have options to connect directly to their UniFi device. The recommended method is to use the built-in Debug Console in UniFi Network. Advanced users can also connect via SSH, however we do not recommend using SSH unless instructed by one of our Support Engineers as part of advanced troubleshooting. Inexperienced users risk making changes that may degrade network performance, or even worse, completely break your deployment. Proceed with caution.
Simple Mode: Debug Console
The Debug Console is the recommended method for connecting directly to a UniFi console or other device.
Note: This requires an operational UniFi OS, so if your network is down, proceed to the SSH option below. Debug is also not recommended if pasting long commands in; be sure to check commands that you paste into it before pressing enter.
- Navigate to the UniFi Devices page and select your device.
- In the device panel, open its Settings, scroll to the bottom and select Debug.
- The debug panel will open and connect directly to the device.
Advanced Mode: SSH
Requirements
1. You are connected to the same local network as the device/console you plan to connect with via SSH. This may consist of using a laptop connected to the same WiFi network, or hardwired directly to the device.
2. SSH is enabled. UniFi Network devices and UniFi Consoles (Dream Machines, CloudKeys, etc.) have independent SSH settings.
- UniFi Consoles: SSH is disabled by default. To enable it, navigate to Settings > Control Plane > Console and check SSH.
- UniFi Network Devices: SSH is enabled by default. The credentials consist of a random string of characters. View and configure them by navigating to UniFi Network > UniFi Devices > Device Updates and Settings (bottom-left corner) and check Device SSH Authentication under Device SSH Settings drop down menu.
3. The device you are using has a command-line interface (CLI) capable of establishing a Secure Shell (SSH) connection. Linux and macOS devices can use their native terminal. Windows OS requires PowerShell or PuTTY.
Establishing an SSH Connection
The format of the command used to establish an SSH connection as follows:
ssh <username>@<ip-address>
The <username> for UniFi Consoles (UDM Pro / UNVR / CloudKey) and UniFi Gateways (UXG Pro) is always ‘root’. For example, a Dream Machine Pro (which is a UniFi Cloud Gateway) with an IP address of 192.168.1.1 can be accessed as follows:
ssh root@192.168.1.1
Note: The UXG will use <username> = ‘root’, but the <password> will be the shared password set in your UniFi Network Application.
Default Credentials
Prior to setup/adoption, devices have a set of default credentials.
- UniFi Consoles - root / ui (root / ubnt on older devices)
- UniFi Gateways - root / ui (root / ubnt on older devices)
- UniFi Devices - ui / ui (ubnt / ubnt on older devices)
