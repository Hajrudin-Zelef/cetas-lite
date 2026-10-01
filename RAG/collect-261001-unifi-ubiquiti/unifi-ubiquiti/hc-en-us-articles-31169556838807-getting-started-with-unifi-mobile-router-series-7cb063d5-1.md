---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-31169556838807-getting-started-with-unifi-mobile-router-series-7cb063d5-1
title: "hc-en-us-articles-31169556838807-getting-started-with-unifi-mobile-router-series-7cb063d5"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-31169556838807-getting-started-with-unifi-mobile-router-series-7cb063d5.md
source_anchor: ""
source_lines: [1, 109]
sha256: ce8156d6fefdb28aa3fcc593eb941d8d1839c436fae910996b2611626931cea5
---

# hc-en-us-articles-31169556838807-getting-started-with-unifi-mobile-router-series-7cb063d5

Getting Started with UniFi Mobile Router Series
Getting Started with UMR Series
Welcome to the UniFi Mobile Router (UMR) Series! This guide will walk you through the adoption process and key features of UMR, UMR-Industrial, and UMR-Ultra to help you set up and manage your device effectively.
UMR is a fully standalone mobile router that operates independently without requiring adoption by a UniFi Console. Users have the flexibility to manage it locally via WiFi or Ethernet or remotely through the optional Mobility Cloud subscription (mobility.ui.com).
Ports & Buttons
| Component | Description | 
| WAN Port | Connects to an external internet source (Ethernet WAN). | 
| LAN Port(s) | Provides wired network connections. | 
| SIM Slot | Insert a nano-SIM for LTE connectivity. | 
| USB-C Port | Alternative power supply option. | 
| Reset Button | Press for 10+ seconds to factory reset and enter recovery mode | 
| LED Indicators | Show device status, connection strength, and errors. | 
LED indicators
| LED status | Description | 
| LED Off | The device is powered off. | 
| Flashing White (1.5s interval) | The device is initializing and booting up. | 
| Rapid flashing white | The device is connecting. | 
| Steady White | The device is ready for adoption or connected via LTE. | 
| Steady Blue | The device is adopted and in normal operating mode. | 
| Alternating White/ Blue | The device is upgrading. | 
| Quickly Flashing White / Blue | The device is in recovery mode. | 
| Flashing Blue (3s On, 0.2s Off) | Network/Wi-Fi error (Device is on but not connected to the Internet). | 
| Flashing White (3s On, 0.2s Off) | SIM card initialization failed (SIM card locked or deactivated). | 
Device Adoption
Adoption is the process of connecting a new Mobility device to Mobility Cloud (mobility.ui.com) for centralized management. Once adopted, you will receive a free trial period for remote access:
- UMR & UMR Industrial: 30-day free trial
- UMR Ultra: 180-day free trial
During this period, you can manage and monitor your devices remotely through Mobility Cloud. If you do not require remote access, you can configure your devices locally via the Local Portal after activation. Please note certain features are only available through the Mobility Cloud (see feature comparison below).
| Feature | Mobility Cloud (Subscription Required) | Local Portal (On-Device Management) | 
| Remote access to all UMR devices | ✅ Yes | ❌ No (Local access only) | 
| Cellular Data reset & Data limit warning | ✅ Yes | ❌ No | 
| Device location reporting (Map view) | ✅ Yes (UMR, UMR Industrial only) | ❌ No | 
| VPN, firewall and port forwarding configurations | ✅ Yes (set by group) | ✅ Yes (by device only) | 
| Static routing | ✅ Yes | ❌ No | 
| S2S VPN with UniFi Gateway | ✅ Yes (Effortless setup) | ✅ Yes (manual setting only) | 
| Global settings | ✅ Yes | ❌ No | 
| Role management | ✅ Yes | ❌ No | 
| Ownership transfer | ✅ Yes | ❌ No | 
| Client management | ✅ Yes | ❌ No | 
| System logs | ✅ Yes | ❌ No | 
| Email notifications | ✅ Yes | ❌ No | 
| DDNS Management | ✅ Yes | ❌ No | 
How to Adopt a Device
Adoption is the process of connecting your UMR, UMR-Industrial, or UMR-Ultra to the Mobility Cloud (mobility.ui.com) for remote management. To ensure a smooth adoption process, follow the steps below.
1. Powering On
Connect the UMR device to power using PoE input (UMR & UMR-Industrial only) or USB-C port.
Wait for the LED to flash white, indicating that the device is booting and connecting.
2. Connecting to the Internet
Ensure that your UMR has internet access by using one of the following methods:
- Insert a SIM card with an active data plan. (For UMR-US, you can subscribe to our pre-installed SIM data plan.)
- Connect an Ethernet cable to the WAN port for a wired internet connection.
- Wait for the LED to turn steady white, indicating that the device is online and ready for adoption.
3. Activating Your Device on Mobility Cloud
Once the device is connected to the internet, activate it on Mobility Cloud to enable remote management.
Via Web UI:
- Go to mobility.ui.com.
- Click "Activate Device" and enter the 9-digit AC code found on the device label.
- The device will appear on your dashboard. Wait for it to go online.
- The LED will turn steady blue, indicating the device is successfully activated.
Via Mobile App (QR Code Scan):
- Scan the QR code located on the device.
- Follow the on-screen instructions to activate the device.
- Once the device goes online, the LED will turn steady blue.
4. Updating to the Latest Firmware
After adoption, ensure your device runs the latest firmware for security updates and feature enhancements. In Mobility Cloud, navigate to Device Settings > Firmware Update and install the latest version. If managing locally, update through the Local Portal.
Troubleshooting Adoption Issues
If you're experiencing issues adopting your UMR device, try the following steps:
Basic Checks
1. Ensure the device is powered on and connected to the internet via SIM or Ethernet.
2. If you activate the device after 24 hours of first powering it on:
- Remove the device from mobility.ui.com.
- Perform a factory reset by holding the reset button for 10+ seconds until the LED starts flashing.
- Re-attempt the activation process.
Check Local Network Settings
- Log into the Local Portal at https://192.168.105.1.
- Ensure that APN settings are correct and data roaming is enabled if using the device abroad.
Collect Support Data for Further Assistance
If the issue persists, download the Support File for troubleshooting:
- Connect your computer via LAN or your mobile device via UMR’s Wi-Fi.
- Open a browser and go to https://192.168.105.1.
- Log in with the admin password (found on mobility.ui.com or default: "ui").
- Navigate to Local Portal > Settings > Download Support File.
- Submit a support request via Mobility Technical Support.
Managing UMR Remotely via Mobility Cloud
Once adopted, your UMR device can be fully managed via the Mobility Cloud (mobility.ui.com), enabling you to effortlessly administer large-scale deployments with Mobility products, which allows for comprehensive monitoring and configuration of all devices.:
- Monitor real-time network status from anywhere, anytime.
- Effortlessly set up S2S VPN connections on the go for secure site-to-site communication.
- Configure VPN, firewall rules, and port forwarding profiles for all your devices remotely.
- Easily invite other administrators to supervise and manage your network remotely.
- Stay informed with email notifications and system logs to track important network events.
Using the Mobility Cloud provides a centralized and seamless way to manage your UMR devices while ensuring secure remote access and network control.
Managing UMR via Local Portal
If remote access is not required, you can configure and manage your UMR locally through the Local Portal.
Accessing the Local Portal:
- Connect your computer via LAN to the UMR’s LAN port or connect your mobile device to UMR’s Wi-Fi.
- Open a browser and go to https://192.168.105.1.
- Log in with the admin password (found on mobility.ui.com or use the default password: "ui").
- Navigate to Local Portal Settings to configure network settings, APN, and other features.
Remote Access via Public IP (If ISP Allows)
Important: If your ISP uses CG-NAT (which most mobile carriers do), remote access via a public IP will not be possible.
Steps for Remote Access via Public IP (If Supported by ISP):
- Verify if your ISP provides a publicly routable IP (not behind CG-NAT).
- Enable port forwarding on the UMR:
  - Important: 
    - Avoid using port 443 as the external port, as this is where the HTTPs runs.
    - External Port can be any value between 1-65535 but should avoid 443 for security reasons.
    - Example:
      - External Port : Any value (recommended: 8787, 8888, etc., but not 443)
      - Internal Port : 443
