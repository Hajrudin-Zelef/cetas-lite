---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-34210126298775-self-hosting-unifi-b7056f9c
title: "hc-en-us-articles-34210126298775-self-hosting-unifi-b7056f9c"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-34210126298775-self-hosting-unifi-b7056f9c.md
source_anchor: ""
source_lines: [1, 83]
sha256: c0271bd2d97e0dc777bb3cd149d862335ba8c6aefadd6e4c59021ce53fd872cb
---

# hc-en-us-articles-34210126298775-self-hosting-unifi-b7056f9c

Self-Hosting UniFi
The UniFi OS Server is the new standard for self-hosting UniFi, replacing the legacy UniFi Network Server. While the Network Server provided basic hosting functionality, it lacked support for key UniFi OS features like Organizations, IdP Integration, or Site Magic SD-WAN. With a fully unified operating system, UniFi OS Server now delivers the same management experience as UniFi-native–including CloudKeys, Cloud Gateways, and Official UniFi Hosting–and is fully compatible with Site Manager for centralized, multi-site control.
| To install the legacy UniFi Network Server instead, refer to the appropriate guide below based on your operating system: | 
Requirements
Before setting up a self-hosted UniFi Network Server, ensure you meet the following requirements:
- 
Operating System:
  - 
Linux: Ubuntu 24.04 or later; Debian 13 or later; or any modern Linux distribution meeting the below requirements:
    - systemd
    - libc 2.31+
    - podman 4.9.3+ (using pasta networking)
    - slirp4netns 1.2+ (typically installed with Podman, but may require separate installation on some systems)
  - 
Windows: Windows 10 or later, Windows Server 2022
    - WSL (Windows Subsystem for Linux) version 2 will be installed during setup.
    - Note that installation within a guest VM in Hyper-V is currently not supported.
  - macOS: Catalina 13.1.0 or later
- 
Linux: Ubuntu 24.04 or later; Debian 13 or later; or any modern Linux distribution meeting the below requirements:
- 
Hardware Specifications:
  - CPU: x86-64 Processor
  - RAM: Minimum 2GB
  - Network: 100Mbps Wired Ethernet
  - Storage: At least 10GB free
- 
Additional Requirements:
  - Ensure necessary ports are open (e.g., TCP 8080, UDP 3478).
Installing UniFi OS Server
Windows/macOS
Linux
Windows as Service
- Download the Installer: Get the latest UniFi OS Server from Releases, or the Download page by right clicking the download and selecting copy link address. It should look something like this: https://fw-download.ubnt.com/data/unifi-os-server/8b93-linux-x64-4.2.23-158fa00b-6b2c-4cd8-94ea-e92bc4a81369.23-x6
- Install the Application: Launch the installer and follow the on-screen setup instructions.
- (Optional) Enable Remote Management: Once installed, create a UI Account (or sign in with an existing one) to enable remote management through Site Manager.
- Adopt Devices: Review UniFi Device Adoption for detailed instructions on adding UniFi devices like APs and switches to your deployment.
- 
Continue Running UniFi: In order to view and manage your adopted devices, the UniFi OS Server must continue running at all times. If it goes offline, your devices will continue to run using the last available configuration.
  - If you are unable to continuously run the UniFi OS Server, we recommend taking advantage of a CloudKey or Official UniFi Hosting as an alternative hosting method.
Note: If you’re migrating from the UniFi Network Server to UniFi OS Server, make sure to close the Network Server before installing UniFi OS Server.
Note: Do not clone VMs except for high availability or failover installations. Cloned VMs will use the same remote access tokens for Site Manager and Fabrics, leading to unexpected behavior. We recommend creating new installations as needed each time.
- Refresh package list and install required packages: sudo apt-get update && sudo apt-get install podman slirp4netns .
- Obtain the UniFi OS Server download link address from Releases, or the Download page by right clicking the download and selecting copy link address. It should look something like this: https://fw-download.ubnt.com/data/unifi-os-server/8b93-linux-x64-4.2.23-158fa00b-6b2c-4cd8-94ea-e92bc4a81369.23-x6
- Download the installer for your architecture using curl or wget for the URL in (2).
  - 
curl: curl -O <uos_server_download_link>
  - 
wget: wget <uos_server_download_link>
- 
curl: 
- Make the installer executable.chmod +x <local_path_to_uos_server_download>
- Run the installer.sudo <local_path_to_uos_server_download>
- (Optional) Enable Remote Management: Once installed, create a UI Account (or sign in with an existing one) to enable remote management through Site Manager.
- Adopt Devices: Review UniFi Device Adoption for detailed instructions on adding UniFi devices like APs and switches to your deployment.
Note: If you’re migrating from the UniFi Network Server to UniFi OS Server, make sure to close the Network Server before installing UniFi OS Server.
Note: Do not clone VMs except for high availability or failover installations. Cloned VMs will use the same remote access tokens for Site Manager and Fabrics, leading to unexpected behavior. We recommend creating new installations as needed each time.
Other Helpful Commands
- Stop UniFi OS Server: sudo systemctl stop uosserver
- Start UniFi OS Server: sudo systemctl start uosserver
- Disable Automatic Starting at System Boot: sudo systemctl disable uosserver
- Enable Automatic Starting at System Boot: sudo systemctl enable uosserver
Updating UniFi OS Server
Your UniFi OS Server and associated UniFi applications can be updated from the Update Manager on Site Manager, or directly from the local Control Plan settings. See UniFi Updates for more information.
Migrating From Network Server to UniFi OS Server
- Obtain a backup from your UniFi Network Server following the instructions here.
- Close the Network Server before installing UniFi OS Server.
- Follow the instructions above in Installing UniFi OS Server according to your operating system.
  - On macOS and Windows, the installer will automatically detect and offer to migrate your existing Network Server setup (if installed in the default location).
- On Linux, or if auto-migration doesn’t occur, you can manually migrate by installing UniFi OS Server and using the Site Export tool. See UniFi Backups and Migration to learn more.
Note: Captive portals will be served on port 8444, changed from port 8843 on Network Server. This port will serve the certificate configured in the UniFi OS Console settings, or if TLS is not configured, this port will be served directly from UniFI Network.
Frequently Asked Questions (FAQs)
Functionally, they are the same. The key difference is that UniFi OS Server is fully self-managed, meaning you are responsible for installation, maintenance, and uptime — unlike Official UniFi Hosting, which is hosted and maintained by us in the cloud.
See Getting Started with Official UniFi Hosting to learn more.
A CloudKey cannot be used to manage UniFi Network devices when paired with a UniFi OS Server, since both include the UniFi Network application and devices can only be managed by a single host. However, CloudKeys can supplement a UniFi OS Server by running additional UniFi applications that the UniFi OS Server does not currently support:
- UniFi Protect (cameras)
- UniFi Access (access control)
- UniFi Talk (VoIP phones)
- UniFi Connect (digital signage, lighting, etc.)
No. Cloud Gateways already include the UniFi Network application as part of their integrated control plane, which means they cannot be managed by an external UniFi OS Server.
If you’re looking to pair a gateway with a UniFi OS Server-managed site, you’ll need to use one of our Independent Gateways (UXG-series), which are designed without a built-in controller and are fully compatible with external UniFi OS deployments
Also make sure that the UniFi devices and the application are on the same network. If not, then remote adoption is required.
This can be safely ignored, however we also allow users to upload their own TLS certificate in Settings > Control Plane > Console > Certificates.
