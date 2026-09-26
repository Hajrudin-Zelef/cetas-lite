---
id: collect-260926-mikrotik/mikrotik/tutorial-home-vlan-configuration-rb5009-cap-acs-multiple-ssids-2
title: "1. Motivation"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/tutorial-home-vlan-configuration-rb5009-cap-acs-multiple-ssids.md
source_anchor: ""
source_lines: [142, 327]
sha256: 69d532c2b0b35512bbf047d46bb9ef05ed5e3ea8e0f0915a0eb2a02b3a74060c
---

# 1. Motivation

## 4.2. Ethernet interfaces

We will configure the Ethernet interfaces to have descriptive names and comments for easy identification. This is not strictly necessary, but it will make it easier to follow the configuration and understand what each port is used for.

```
/interface ethernet
set [ find default-name=ether1 ] name=ether1-wan        comment="ISP PPPoE WAN"
set [ find default-name=ether2 ] name=ether2-ap1        comment="Access Point 1"
set [ find default-name=ether3 ] name=ether3-ap2        comment="Access Point 2"
set [ find default-name=ether4 ] name=ether4-laptop     comment="Personal laptop"
set [ find default-name=ether5 ] name=ether5-rpi        comment="Raspberry PI"
set [ find default-name=ether8 ] name=ether8-management comment="Dedicated unbridged port for management access"
set [ find default-name=sfp-sfpplus1 ] disabled=yes
```

 Customization

- Adjust the names and comments to match your setup.
- If you have more or fewer Ethernet ports, adjust the configuration accordingly.
- If you have an SFP+ port, you can keep it enabled it if you plan to use it.


## 4.3. VLANs

Create a bridge and enable VLAN filtering on it.

```
/interface bridge add name=bridge-lan vlan-filtering=yes comment="LAN bridge"
```

Create a VLAN interface for each VLAN we want to use. This allows us to use IP (Layer 3) services such as on each VLAN.

```
/interface vlan
add interface=bridge-lan name=vlan10-owner      vlan-id=10 comment="Owner VLAN"
add interface=bridge-lan name=vlan20-iot        vlan-id=20 comment="IoT VLAN"
add interface=bridge-lan name=vlan30-guest      vlan-id=30 comment="Guest VLAN"
add interface=bridge-lan name=vlan99-management vlan-id=99 comment="Management VLAN"
```

Next add the required bridge ports to the bridge.

```
/interface bridge
# The Wi-Fi access points ports are trunk ports which will carry traffic for mutliple VLANs.
port add bridge=bridge-lan frame-types=admit-only-vlan-tagged interface=ether2-ap1 comment="AP1 Trunk Port"
port add bridge=bridge-lan frame-types=admit-only-vlan-tagged interface=ether3-ap2 comment="AP2 Trunk Port"
# The Ethernet ports for personal devices and IoT devices are access ports which will carry traffic for a single VLAN.
port add bridge=bridge-lan frame-types=admit-only-untagged-and-priority-tagged interface=ether4-laptop pvid=10 comment="Laptop Access Port (VLAN 10)"
port add bridge=bridge-lan frame-types=admit-only-untagged-and-priority-tagged interface=ether5-rpi    pvid=20 comment="Raspberry PI Access Port (VLAN 30)"
```

 Do not add the management port (`ether8-management`) to the bridge, this port is used for unbridged management access only.


Next we configure the bridge VLAN table to restrict which VLANs are allowed on each port.

```
vlan add bridge=bridge-lan tagged=bridge-lan,ether2-ap1,ether3-ap2 untagged=ether4-laptop vlan-ids=10 comment="Owner VLAN with access port for laptop"
vlan add bridge=bridge-lan tagged=bridge-lan,ether2-ap1,ether3-ap2 untagged=ether5-rpi    vlan-ids=20 comment="IoT VLAN with access port for Raspberry PI"
vlan add bridge=bridge-lan tagged=bridge-lan,ether2-ap1,ether3-ap2                        vlan-ids=30 comment="Guest VLAN with access only via Wi-FI AP trunks"
vlan add bridge=bridge-lan tagged=bridge-lan,ether2-ap1,ether3-ap2                        vlan-ids=99 comment="Management VLAN"
```

## 4.4. Basic IP Services

Next we will configure various IP services on the router, such as DHCP and DNS.

### 4.4.1. DNS

I decided to use Cloudflare and Google DNS servers for my home network.

```
/ip dns set allow-remote-requests=yes servers=1.1.1.1,8.8.8.8
```

 Customization

- You can use any DNS servers you prefer, such as your ISP's DNS servers or a local DNS server.


### 4.4.2. IP Addresses

Add IP addresses to the VLAN interfaces we created earlier.

 IP capacity planning is a complex topic which I will not provide guidance on here.


For my home network I decided to use a `/24` subnet for each VLAN, which provides sufficient capacity for my needs. I also decided to use `10.<VLAN_ID>.0.0/24` as the network for each VLAN, which makes it easy to identify the VLAN based on the IP address.

```
/ip address
add address=10.10.0.1/24 comment="Owner VLAN addresses"      interface=vlan10-owner      network=10.10.0.0
add address=10.20.0.1/24 comment="IoT VLAN addresses"        interface=vlan20-iot        network=10.20.0.0
add address=10.30.0.1/24 comment="Guest VLAN addresses"      interface=vlan30-guest      network=10.30.0.0
add address=10.99.0.1/24 comment="Management VLAN addresses" interface=vlan99-management network=10.99.0.0
```

 Customization

- Adjust the IP addresses and networks to suit your needs.


### 4.4.3. DHCP Server

Each VLAN will have its own DHCP server to assign IP addresses to devices on the network.

I decided to only use `10.<VLAN_ID>.0.100-10.10.0.199` as the DHCP range for each VLAN, which leaves room for static IP addresses in the `10.<VLAN_ID>.0.2-10.<VLAN_ID>.0.99` and `10.<VLAN_ID>.0.200 10.<VLAN_ID>.0.254` ranges, which suits my needs.

Create the IP pools for each DHCP server to use.

```
/ip pool
add name=pool-owner ranges=10.10.0.100-10.10.0.199
add name=pool-iot ranges=10.20.0.100-10.20.0.199
add name=pool-guest ranges=10.30.0.100-10.30.0.199
add name=pool-management ranges=10.99.0.100-10.99.0.199
```

 Customization

- Adjust the IP pools to suit your needs, but I recommend reserving a range for static IP addresses.


Create each DHCP server and add a DHCP network for each.

```
/ip dhcp-server
add address-pool=pool-owner      interface=vlan10-owner      lease-time=1d name=dhcp-owner
add address-pool=pool-iot        interface=vlan20-iot        lease-time=1d name=dhcp-iot
add address-pool=pool-guest      interface=vlan30-guest      lease-time=1h name=dhcp-guest
add address-pool=pool-management interface=vlan99-management lease-time=1d name=dhcp-management
network add address=10.10.0.0/24 dns-server=10.10.0.1 gateway=10.10.0.1 comment="Owner DHCP Network"
network add address=10.20.0.0/24 dns-server=10.20.0.1 gateway=10.20.0.1 comment="IoT DHCP Network"
network add address=10.30.0.0/24 dns-server=10.30.0.1 gateway=10.30.0.1 comment="Guest DHCP Network"
network add address=10.99.0.0/24 dns-server=10.99.0.1 gateway=10.99.0.1 comment="Management DHCP Network"
```

 Customization

- Adjust the lease times to suit your needs. Typically guest devices will have shorter lease times, while permanent devices will have longer lease times.


### 4.5 PPPoE Client

Now we will configure the PPPoE client to connect to the internet.

```
/interface pppoe-client
add add-default-route=yes allow=pap,chap comment="ISP PPPoE" interface=ether1-wan max-mru=1480 max-mtu=1480 mrru=1600 name=pppoe-out user=your_user@isp.com password=your_password
```

 Customization

- The `allow` ,`max-mru` ,`max-mtu` and`mrru` values are what my ISP requires, adjust these values to suit your ISP's requirements.
- Specify your ISP's PPPoE username and password in the `user` and`password` fields.
- If you do not use a PPPoE connection, you can skip this step and configure your internet connection accordingly.


## 4.6. Wi-Fi

Now we will configure the Wi-Fi, including the special configuration considerations for the cAP ac APs.

The cAP ac APs do not support dynamic VLAN assignment based on datapath when using the`wifi-qcom-ac` package.

 If you are using cAP ax APs, please see this official example configuration and adapt the steps below accordingly.


### 4.6.1. Datapath

Create a single generic datapath.

```
/interface wifi datapath add bridge=bridge-lan disabled=no name=dp-ac-generic comment="Generic datapath for cAP ac APs"
```

### 4.6.2. Band steering

I decided to enable band steering on the cAP ac APs to encourage dual-band capable devices to connect to the 5GHz band, which has less interference and higher speeds.

```
/interface wifi steering add disabled=no name=steering-main rrm=yes wnm=yes
```

 Customization

