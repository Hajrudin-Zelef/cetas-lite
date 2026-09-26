---
id: collect-260926-mikrotik/mikrotik/tutorial-home-vlan-configuration-rb5009-cap-acs-multiple-ssids-1
title: "1. Motivation"
domain: mikrotik
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/tutorial-home-vlan-configuration-rb5009-cap-acs-multiple-ssids.md
source_anchor: ""
source_lines: [1, 141]
sha256: 0b709d850fff86812a26379cd51e0866bf6f13466f00ce208eb21436acd2e2c6
---

# 1. Motivation

Hi all,

I recently decided to explore the idea of segmenting my home network into multiple VLANs. I am by no means a Mikrotik expert and my networking knowledge was mostly confined to cloud provider abstractions. Due to this lack of experience it was a frustrating and time-consuming experience.

Instead of just dumping my config here and asking the community if it's correct, I decided I'd rather write this up as a guide so that other people can also benefit from my learnings. I would of course appreciate any recommendations or corrections, and I'll update the guide to reflect these.

 **Disclaimer**: I am not a networking or Mikrotik expert, and this guide is based on my personal experience and learning.

This guide is opinionated and does not claim to be error-free. Please consult the official Mikrotik documentation to verify any commands or configurations that you apply.


# 1. Motivation

There are many reasons why you might want to segment your network with VLANs, but my specific goals were:

1. Control IoT device communication, including the ability to prevent IoT devices from initiating communication with security sensitive devices on my network (such as my laptop).
2. Allow guest devices to access the internet and nothing else (no communication with other devices on the network, including other guest devices).
3. Prepare for various other things I'd like to implement, such as secure access to my home network with a reverse tunnel.
4. Learn more about networks and generally just play around with RouterOS for fun.

There are almost certainly other approaches to achieve these goals, but I wanted to do it with VLANs since this felt like the most flexible and future-proof solution.

# 2. Overview

## 2.1. VLANs

The end result of this guide is a network with the following VLANs:

- **VLAN 10** : Owner Devices  - This VLAN is for my personal devices, such as my laptop, TV and phone.
- **VLAN 20** : IoT Devices  - This VLAN is for my IoT devices, such as a Raspberry Pi (physically connected via ethernet) and Victron Cerbo GX (
connected via Wi-Fi).
- **VLAN 30** : Guest Devices  - This VLAN is for guest devices, such as visitors' phones that connect to Wi-Fi.
- **VLAN 99** : Management  - This VLAN is for management of the router and access points.

Even though this is likely a common setup, you can adjust the VLANs to suit your needs. For example, you might want to create a separate VLAN TVs or place TVs in the IoT VLAN (I decided against this since I wanted to simplify casting from my phone to the TV).

## 2.2. Wi-Fi SSIDs

To assign devices to the correct VLANs, the following Wi-Fi SSIDs will be created:

- **OwnerWifi** : This SSID will be used for my personal devices and will be assigned to VLAN 10.
- **IoTWifi** : This SSID will be used for IoT devices and will be assigned to VLAN 20.
- **GuestWifi** : This SSID will be used for guest devices and will be assigned to VLAN 30.  - This SSID will have client isolation enabled, so guest devices will not be able to communicate with each other.

 Instinctively you might want to create a single SSID for all devices and use MAC address filtering to assign devices to the correct VLANs, but this is generally not recommended.


## 2.3. Hardware

Even though many of the concepts in this guide applies to RouterOS in general, there are some things that are specific to the hardware I am using.

- Router: **Mikrotik RB5009UPr+S+IN**  - The RB5009**UG** +S+IN should work exactly the same - I chose the UPr since I wanted multiple PoE outputs.
- Wireless Access Points: **Mikrotik RBcAPGi-5acD2nD** (x2)  - I already had two of these, so I made do with them.
  - These cAP ac APs unfortunately don't support CAPsMAN based VLAN provisioning at the time of writing, which caused a lot of frustration, but the guide explains how make VLANs work with them.
  - If you use cAP ax APs, you'll have less configuration to do on the APs themselves (more on this later).
- Optical Network Terminal: Huawei EchoLife HG8240H
  - This is the ONT provided by my ISP, it facilitates the connection to the fiber network via GPON and PPPoE.
  - If you use something other than PPPoE to connect to the internet you will need to adjust the configuration accordingly.
  - Configuration changes on the ONT was not necessary in my case, and generally shouldn't be required.
- Computer:
  - Any computer with internet access and an Ethernet port should work.
- Cabling:
  - You will need an Ethernet cable to connect your computer to the router and access points.

## 2.4. Software

I used the latest stable version of RouterOS at the time of writing, but this guide should work with any 7.x version.

See the official Mikrotik download page for the latest versions: https://mikrotik.com/download

- Router:
  - RouterOS: **7.19.2** (`routeros-7.19.2-arm64.npk` in my case)
- Wireless Access Points:
  - RouterOS: **7.19.2** (`routeros-7.19.2-arm.npk` in my case)
  - Wifi QCom AC (`wifi-qcom-ac-7.19.2-arm.npk` in my case)    - It's important to use this package instead of the older `wireless` package if you are using cAP ac APs.
- Configuration software:
  - Windows
    - WinBox is required for this guide since we will use Layer 2 (MAC) based communication for configuration.
    - If you only have Linux available, you can consider using `wine` or virtualization (such as VirtualBox) to run Winbox.
  - WinBox: **3.42**    - The latest version at the time of writing, but any recent version should work.

# 3. Initial Setup

This guide assumes that you are starting with a fresh RouterOS installation without any defaults on both the router and access points.

First we will reset the router and access points.

 This will erase all existing configurations on the router and access points. Make sure to back up any existing configurations you want to retain before proceeding.


1. Connect the router to your computer using an Ethernet cable.
  1. Use **ethernet port 8** , since this port will be a dedicated port for configuration in this guide.
2. Open WinBox and connect to the router using the MAC address.
3. Reset the router to factory defaults```
 /system reset-configuration no-defaults=yes skip-backup=yes  
```
Your router will reboot, and you will need to reconnect to it using WinBox.
4. Configure your username and password
  1. It's best practice to use a non-default username and a strong password.

Repeat the same steps for each of your access points, but connect the access points to your computer using **ethernet port 2**, since this port will be a dedicated port for configuration of the access points.

The next step is to install the required packages on the router and access points.

Since the devices do not have any internet at this point, we will perform a manual upgrade using WinBox.

Please follow this guide to install the packages as specified in the "Software" section above.

 It's highly recommended to also update the RouterBoard firmware after upgrading RouterOS. See the link above.

 If you are using cAP ac APs, make sure to uninstall the older `wireless` package if it's installed.


# 4. Router Configuration

Now that we have a clean RouterOS installation on the router and access points, we can start configuring the router.

We will use the terminal in WinBox to configure the router, but you can also use the GUI if you prefer.

## 4.1. Basic Configuration

```
/system 
# Set the identity
identity set name=router
# Setup initial timezone and time
clock set time-zone-name=Africa/Johannesburg
clock set date="jun/28/2025"
clock set time="15:00:00"
# Use Cloudflare NTP servers for accurate timekeeping
ntp client set enabled=yes servers=time.cloudflare.com
```

 Customization

- Set the `time-zone-name` to your local timezone. You can find a list of timezones
here.
- Set the `date` and`time` to the current date and time (anything close is fine).
- It's not necessary to use an NTP server, and you can use your ISP's NTP server if you prefer.


