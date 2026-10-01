---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-mx-b5244a77
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-mx-b5244a77"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-mx-b5244a77.md
source_anchor: ""
source_lines: [1, 28]
sha256: b05f39a0f6f950efd24efbc1888ac256f4945a11ec32ddeea677b8b0ffec1e52
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-mx-b5244a77

MX to Sonicwall Site-to-Site VPN Setup
When setting up a non-Meraki Site-to-Site VPN between an MX Security Appliance and a Sonicwall, the following settings should be used on the Sonicwall to get the tunnel up and running.
General Tab
The settings configured on the General tab on the Sonicwall interface should follow the configuration below:
- Policy Type: Site to Site
- Authentication Method: IKE using Preshared Secret
- Name: Enter a name the security policy will be displayed as on the Sonicwall
- IPsec Primary Gateway Name or Address: Enter the public IP address of the MX.
- IPsec Secondary Gateway Name or Address: Use the address "0.0.0.0"
- Shared Secret: This should match the Preshared secret configured for this peer on the Security & SD-WAN > Configure > Site-to-site VPN page in Dashboard
- Local IKE ID: Select "IP Address" and enter the public IP address of the Sonicwall.
- Peer IKE ID: Select "IP Address" and enter the IP address configured on the MX's primary uplink. If the MX is relying on a cellular connection, use the IP address of the cellular modem. If the MX is behind a NATing device, this IP (unlike the public IP address that was used for "IPsec Primary Gateway Name or Address" mentioned above) will be the NATed IP address of the MX.
Proposals Tab
The configuration of this page should match the phase 1 and 2 parameters as configured on the MX, if the MX is utilizing custom IPsec policies. If the MX is using the default parameters, then the settings configured on the Proposals tab should follow the configuration below, and match the screenshot provided:
- Exchange: Main Mode
- DH Group: Group 2
- Encryption: AES-256
- Authentication: SHA1
- Life Time (seconds): 28800
- Protocol: ESP
- Encryption: AES-256
- Authentication: SHA1
- Enable Perfect Forward Secrecy: Enable, the box should be checked
- Life Time (seconds): 28800
Additional Notes
- On the Advanced tab, ensure the box for Enable Keepalive is checked.
- Make sure that the remote subnets configured on Sonicwall exactly match the VPN subnets configured on the MX.
- If Address Objects are used, select the appropriate type between 'Network' and 'Range'.
