---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1058104-hyper-v-ubuntu-server-vm-not-connecting-to-unifi-usg-c8789d34
title: "questions-1058104-hyper-v-ubuntu-server-vm-not-connecting-to-unifi-usg-c8789d34"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1058104-hyper-v-ubuntu-server-vm-not-connecting-to-unifi-usg-c8789d34.md
source_anchor: ""
source_lines: [1, 11]
sha256: 51f3bd0917e7154011f926eeeed237d7beb33da9f3fbfdf1e0be24ae2969eda0
---

# questions-1058104-hyper-v-ubuntu-server-vm-not-connecting-to-unifi-usg-c8789d34

I have a network set up with 2 VLANs, one for personal interconnectivity and internet access for personal devices, and another one for non-trusted devices like IOT devices. Here's the network topology:
In short I have two subnets and two VLANs. 192.168.1.0/24, and 192.168.2.0/24 with VLANs 1 and 2 respectively. They are not allowed to communicate with each other. I have a Jellyfin Server running on Ubuntu Server 20.04.02 LTS VM that's bridged with a Hyper-V switch with static IP 192.168.1.223. It's reverse proxied for ease of access. (http://jellyfin/).
I'd like to make it so that anyone who's connected to the VPN (192.168.2.15) can type in Jellyfin's VM VLAN 2 IP (192.168.2.17) into a browser and access my jellyfin server. The apache2 side of things I can handle. My issue starts with where it seems that my jellyfin server can't seem to access the second VLAN. I set it up with:
sudo ip link add link eth0 name eth0.2 type vlan id 2
then I edited /etc/NetworkManager/NetworkManager.conf
[keyfile]
unmanaged-devices=*,except:type:wifi,except:type:wwan,except:type:ethernet,except:type:vlan
I added "except:type:vlan". I restarted Network Manager. Then I opened NMTUI and set the eth0.2 device with static IP 192.168.2.17, gateway 192.168.2.1, DNS 192.168.2.1, subnet 192.168.2.0/24
then I activated the device.
This broke the DNS records for the VM but I was still able to access http://jellyfin/ from any device on subnet 1. I could ping any IP address from VLAN 1 but VLAN 2 got HOST DESTINATION UNREACHABLE.
Any ideas on how I can improve? I'm willing to edit this with more information if I'm asked. Thank you!
