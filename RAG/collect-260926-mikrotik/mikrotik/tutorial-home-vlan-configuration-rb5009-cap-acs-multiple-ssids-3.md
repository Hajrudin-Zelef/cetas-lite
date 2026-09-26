---
id: collect-260926-mikrotik/mikrotik/tutorial-home-vlan-configuration-rb5009-cap-acs-multiple-ssids-3
title: "1. Motivation"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/tutorial-home-vlan-configuration-rb5009-cap-acs-multiple-ssids.md
source_anchor: ""
source_lines: [328, 443]
sha256: 026f1b27b1e3df90916e690a1d6bac1ab4d92ded8fea99420e4ffbccf614eca1
---

# 1. Motivation

- Band steering is optional, do your own research to determine if it's suitable for your environment.


### 4.6.3. Security Profiles

Create security profiles for each SSID we will create. These profiles are provisioned to the access points using CAPsMAN.

```
/interface wifi security
add authentication-types=wpa2-psk,wpa3-psk disabled=no ft=yes name=sec-owner comment="Owner WiFi security"
add authentication-types=wpa2-psk,wpa3-psk disabled=no ft=no  name=sec-iot   comment="IoT WiFi security"
add authentication-types=wpa2-psk,wpa3-psk disabled=no ft=yes name=sec-guest comment="Guest WiFi security"
```

 Customization

- I decided to use WPA2/WPA3 PSK for all SSIDs. You can use WPA3 only if you prefer, but this will prevent older devices from connecting.
- Fast Transition (FT) is optional. I enabled FT for the owner and guest SSIDs, while I disabled it for the IoT SSID since these devices are not expected to roam between access points.


### 4.6.4. Configuration Profiles

Create configuration profiles for each SSID we want to create. These profiles are provisioned to the access points using CAPsMAN.

```
/interface wifi configuration
add country="South Africa" datapath=dp-ac-generic disabled=no name=config-owner-wifi security=sec-owner ssid=OwnerWifi steering=steering-main comment="Config for Owner WiFi SSID"
add country="South Africa" datapath=dp-ac-generic disabled=no name=config-iot-wifi   security=sec-iot   ssid=IoTWifi                          comment="Config for IoT WiFi SSID"
add country="South Africa" datapath=dp-ac-generic disabled=no name=config-guest-wifi security=sec-guest ssid=GuestWifi steering=steering-main comment="Config for Guest WiFi SSID"
```

 Customization

- Adjust the `country` field to your country name.
- Adjust the `ssid` field to your desired SSID names.
- The
 `steering` field is optional, I enabled it for the owner and guest SSIDs to encourage dual-band capable devices to connect to the 5GHz band.


### 4.6.5. CAPsMAN

Enable and configure CAPsMAN to manage the access points.

 The `action` is not set to `create-dynamic-enabled` (instead we use `create-enabled`) since we are somewhat manually managing the CAPs due to the "wifi-qcom-ac" package not supporting automatic VLAN provisioning.


```
/interface wifi
capsman set enabled=yes ca-certificate=auto certificate=auto interfaces=vlan99-management
provisioning add action=create-enabled disabled=no master-configuration=config-owner-wifi slave-configurations=config-iot-wifi,config-guest-wifi comment="Provision config to all APs"
```

## 4.7. Security

 Security is a complex topic and this section will not cover all aspects of securing your home network. See the Securing your router and Firewall documentation for more information.


### 4.7.1. Interface lists

Create interface lists to group interfaces based on their purpose. This will make it easier to manage firewall rules and other configurations.

```
/interface list
add name=management comment="Management interface list"
member add list=management interface=ether8-management comment="Dedicated management port"
member add list=management interface=vlan99-management comment="Management VLAN"
add name=vlan comment="All VLANs excluding management VLAN"
member add list=vlan interface=vlan10-owner    comment="Owner VLAN"
member add list=vlan interface=vlan20-iot      comment="IoT VLAN"
member add list=vlan interface=vlan30-guest    comment="Guest VLAN"
add name=wan comment="WAN interface list"
member add list=wan interface=pppoe-out comment="PPPoE WAN interface"
```

### 4.7.2. Address lists

I wanted my laptop and phone to be able to access the router via HTTP, and I assigned static IP addresses to these devices. I thus have an address list for them which is used in the firewall rules.

```
/ip firewall address-list
add address=10.10.0.2 list=owner-webfig comment="Owner laptop"
add address=10.10.0.3 list=owner-webfig comment="Owner phone"
```

### 4.7.3. Disable unused services and tools

Disable any unused services and tools to reduce the attack surface of the router.

```
/tool 
# Disable the bandwidth test server
bandwidth-server set enabled=no
# Restrict Layer 2 (MAC) access to the management interface list
mac-server mac-winbox set allowed-interface-list=management
mac-server set allowed-interface-list=management
/ip
# Only allow neighbor discovery from the management interface list
neighbor discovery-settings set discover-interface-list=management
# Disable unused services
service disable telnet,ftp,api,api-ssl
# Disable DynDNS
cloud set ddns-enabled=auto update-time=no
# Use stronger encryption for SSH
ssh set strong-crypto=yes
```

 Customization

- The configuration above is opinionated and may not suit your needs, so please review it carefully.
- I decided to only allow Layer 2 (MAC) access to the management interface list. If you want to allow Layer 2 access via WinBox to other interfaces, you can adjust the `allowed-interface-list` accordingly.


### 4.7.4. Firewall

The firewall configuration below is based on various sources and research, and is by no means comprehensive or a definitive reference for securing your home network. Again, please consult the official Mikrotik documentation.

