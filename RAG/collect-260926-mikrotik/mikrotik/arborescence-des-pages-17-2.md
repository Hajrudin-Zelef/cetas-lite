---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-17-2
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "license"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-17.md
source_anchor: ""
source_lines: [75, 135]
sha256: 755f8549db4be7ff742ac343f038693a2f13f4de79248ae92b47484fb6b2f641
---

# Introduction

| Parameter | Meaning | 
|---|---|
| **-r** | Reinstall device and apply default configuration script. | 
| **-e** | Reinstall device without applying default configuration script - leave device with empty config | 
| **-b** | Option to discard the currently installed branding package from the device, otherwise it will kept | 
| **-m** | Enables multiple device reinstallation. **Available options:** **No "-m", no "-o" chosen** – Only one successful installation will proceed, and Netinstall will close afterward. **"-m" only** – Enables multiple device reinstallation. The same device will be reinstalled as many times as BOOTP requests are sent. **"-o" only** – Functions the same as when neither "-m" nor "-o" is chosen: only one successful installation will proceed, and Netinstall will close afterward. **Both "-m" and "-o"** – Enables multiple device reinstallation. The same device will be reinstalled only once per Netinstall run. | 
| **-o** | When using the netinstall tool with the "-o" option, devices can only be installed once per netinstall run. This means that during the netinstall process, the tool will keep track of the MAC addresses of devices that were successfully installed. If a device with the same MAC address tries to reinstall during the same run, the tool will ignore it and not respond to its BOOTP requests. | 
| **-f** | Ignore size constraints. Netinstall-cli checks the storage size on the router. If the total size of the selected packages exceeds the device's available storage, an error will be displayed: **"Ignoring XX:XX:XX:XX:XX:XX, not enough space (override with -f)"** | 
| **-c** | Allow to run multiple Netinstall instances on the same computer. | 
| **-v** | Verbose mode | 
| **-k <keyfile>** | Provides the device with a license key in .KEY format (optional). | 
| **-s <userscript>** | Specifies **Configure script** to be installed on device, this replaces replaces with RouterOS supplied default configuration script. `/system/default-configuration/custom-script/print` This script will be kept during RouterOS updates and used after further configuration resets till device will be reinstalled with new script or removed if no script provided. | 
| **-sm <modescript>** | Specifies a one‑time custom script to run on the device’s first boot after installation. Use this script  to configure **device‑mode** and**protected‑routerboot** and other settings during device deployment. Mode script executes before any custom or default configuration scripts. Upon completion, script is automatically removed from the device. If the script modifies the device-mode, the device will be reboot immediately after execution. This feature requires RouterOS and Netinstall version 7.22 or newer. | 
| **--mac <mac address>** | Specifies MAC address which will be allowed to be installed. When a MAC address is provided, all other BOOTP requests are disregarded. | 
| **-i <interface>** | Allows you to specify an interface (optional). | 
| **-a <IP address>** | Uses a specific IP address that the Netinstall server will assign to the device. Mandatory, but can be auto-assigned if interface parameter used. | 
| **PACKAGE** | Specify a list of RouterOS.NPK format packages that Netinstall will try to install on the device (mandatory). **The system package must be listed first.** | 

If the "-r" or "e-" parameter is not specified, *netinstall-cli* will reinstall RouterOS  and will keep the current configuration by downloading current configuration database from the router, reinstalling the router (including disk formatting), and uploading the configuration back to it, the same as  Netinstall **"Keep old configuration"** option. However, it's important to note that this process solely applies to the configuration itself and does not impact the files, including databases like the User Manager database, Dude database, and others.

# Quick start guide for Windows


- Download the **Stable** or**Testing** version of the**Netinstall** utility from the downloads page;
- Download the RouterOS **Main package** from the downloads page;You need to select a RouterOS version, preferably one marked as **Stable** . Additionally, choose the appropriate architecture (ARM, MIPS, SMIPS, TILE, etc.). If unsure, you can download the RouterOS package for all architectures, and Netinstall will determine the correct one for your device.
- Disable all computer network interfaces (WiFi, Ethernet, LTE, or any other type of connection) except for the one to be used for installation. Netinstall will only function with one active interface on your computer. It's strongly recommended to deactivate any other network interfaces to ensure Netinstall selects the correct one.
- Configure a static IP address for your Ethernet interface, open **Start,** and select**Settings** :

Netinstall can run also on a local network, in such case you could skip setting a static IP address, but it is highly recommended that you set a static IP address if you are not familiar with Netinstall.

- Open **Network & Internet** and select**Change adapter options**

-  Right-click on your Ethernet interface and select **Properties**

- Select **Internet Protocol Version 4 (TCP/IPv4)** and click**Properties**

- Check Use the following IP address and fill out the fields as shown in the image below

- Open your Downloads folder (or wherever you saved the downloaded files) and extract the Netinstall *.zip file to a convenient place

- Make sure that the Ethernet interface is running and launch Netinstall.exe. If you followed the guide precisely, then you should not have any Internet connection on your computer, Windows 10 wants to verify all apps that it runs, but will not be able to do it since lack of an Internet connection, for this reason, a warning might pop up, you should click **Run** .

Netinstall requires administrator rights, there should be a window asking for permissions to run Netinstall, you must accept these permissions in order for Netinstall to work properly.

- Allow access for Netinstall in **Public** networks and configure**Net booting** settings and fill out the required fields as shown in the image below

- Connect your device to your computer using an ethernet cable directly (without any other devices in-between), plug the Ethernet cable into your device's Etherboot port (see the next "Warning" in this article before connecting your Netinstall network).
- MikroTik devices are able to use Netinstall from their **first** port (Ether1), or from the port marked with "**BOOT** ".

Some computers have a network interface (especially USB Ethernet adapters) that tend to create an extra link flap, which is enough for Netinstall to fail to detect a device that is in Etherboot mode. In such a case you can use a switch between your device and your computer or a router in bridge mode to prevent this issue. If you use RouterOS powered router in bridge mode, then make sure that you disable any DHCP clients on the router bridge interface and disable Detect Internet feature.

Netinstall uses bootp packets, which are using the same port numbers as DHCP packets. If you're using a switch between your PC and the device to be Netinstalled, ensure that the ports in the bridge are not blocked by other network devices.

If you have dhcp-snooping enabled, make sure to enable "trusted" on the bridge ports facing the Netinstall PC.

- Power up your device and put it into etherboot mode

There are multiple ways how to put your device into Etherboot mode. Make sure you read the Etherboot manual before trying to put the device into this mode. Methods vary between different MikroTik devices.

- Wait for the device to show up in Netinstall, then select it and click **Browse.**  Navigate to your**Downloads** folder (or wherever you saved your RouterOS packages) and press**OK.**

