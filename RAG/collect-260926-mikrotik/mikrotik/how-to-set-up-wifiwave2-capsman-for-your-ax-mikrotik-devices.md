---
id: collect-260926-mikrotik/mikrotik/how-to-set-up-wifiwave2-capsman-for-your-ax-mikrotik-devices
title: "how-to-set-up-wifiwave2-capsman-for-your-ax-mikrotik-devices"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/how-to-set-up-wifiwave2-capsman-for-your-ax-mikrotik-devices.md
source_anchor: ""
source_lines: [1, 77]
sha256: 1cee941c1337c139bd205fb109bddf562efdd33271e490508e7525faf76dd871
---

# how-to-set-up-wifiwave2-capsman-for-your-ax-mikrotik-devices

##### How to set up WifiWave2 CAPsMAN for your AX MikroTik devices

WifiWave2 is a new package for MikroTik RouterOS that introduces a ton of new features and optimisations for 802.11ac Wave2 and 802.11ax wireless RouterBOARD interfaces. In addition to new features like WPA3, OWE, and Beamforming, WifiWave2 includes a new and improved CAPSMAN. WifiWave2 CAPsMAN allows applying wireless settings to multiple MikroTik WifiWave2 AP devices from a central configuration interface. This enables seamless management of your entire MikroTik Wi-Fi network from a single management plane.

**WifiWave2 CAPsMAN requirements**

- Any RouterOS device, that supports the WifiWave2 package, can be a controlled wireless access point (CAP) as long as it has at least a Level 4 RouterOS license.
- WifiWave2 CAPsMAN server can be installed on any RouterOS device that supports the WifiWave2 package, even if the device itself does not have a wireless interface
- Unlimited CAPs (access points) supported by CAPsMAN

 

**Note** - WifiWave2 CAPsMAN can only control WifiWave2 interfaces, and WifiWave2 CAPs can join only WifiWave2 CAPsMAN, similarly, regular CAPsMAN only supports non-WifiWave2 caps.

**Setting up WifiWave2 CAPsMAN**

Our lab made use of a hAP ax³ as the CAPsMAN server but you can use any RouterBOARD device with the WifiWave2 package installed, it does not have to have wireless interfaces to act as your server. For the managed cAP we used a hAP ax² but you can use any AX-based MikroTik access point or any of the devices on the WifiWave2 supported devices list - https://help.mikrotik.com/docs/display/ROS/WifiWave2#WifiWave2-Compatibility


**Note** - All devices were running the latest stable release of RouterOS (7.11.2)

**On your CAPsMAN server**

#Configure a security profile

/interface wifiwave2 security

add authentication-types=wpa3-psk name=testsec passphrase=testing123


#configure configuration profiles for both 2GHz and 5GHz bands

/interface wifiwave2 configuration

add country="South Africa" name=5G security=testsec ssid=CAPTest5G

add country="South Africa" name=2G security=testsec ssid=CAPTest2G


#configure provisioning templates for your remote CAPs

/interface wifiwave2 provisioning

add action=create-enabled master-configuration=5G supported-bands=5ghz-n

add action=create-enabled master-configuration=2G supported-bands=2ghz-n


#Enable CAPsMAN on your CAPsMAN server

/interface wifiwave2 capsman

set ca-certificate=auto enabled=yes


**On your remote CAP/s**

#enable CAP service

/interface/wifiwave2/cap set enabled=yes


**Note** - If your CAPsMAN server is not on the same LAN as your CAP you can specify a caps-man-addresses value here (ex. caps-man-addresses=172.16.254.1)


#set configuration.manager on the wireless interfaces that should be managed by CAPsMAN

/interface/wifiwave2/set wifi1,wifi2 configuration.manager=capsman-or-local


Still stuck? Get in touch with the team support@miro.co.za or 012 657 0960 to chat with one of our expert support specialists.

Get your hands on the latest hardware from MikroTik from MiRO here.

## Comments

View Comments
