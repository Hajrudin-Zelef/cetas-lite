---
id: collect-261001-general-networking/general-networking/wireless-troubleshooting-and-support-troubleshooting-troubleshooting-local-conne-df19f7e1-1
title: "wireless-troubleshooting-and-support-troubleshooting-troubleshooting-local-conne-df19f7e1"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/wireless-troubleshooting-and-support-troubleshooting-troubleshooting-local-conne-df19f7e1.md
source_anchor: ""
source_lines: [1, 102]
sha256: 6bfc8d06b026ce5554a4da37407e903294967d32797d1981720055497eb138cd
---

# wireless-troubleshooting-and-support-troubleshooting-troubleshooting-local-conne-df19f7e1

Troubleshooting Meraki AP Cloud and Gateway Connectivity
Click 日本語 for Japanese
Overview
This article explains how to diagnose and resolve common connectivity issues encountered with Cisco Meraki MR access points (APs). It provides step-by-step instructions to identify and fix problems that prevent an AP from contacting the Meraki Cloud, operating as a gateway, or being configured locally.
The guide covers four related issue categories:
- An AP reporting difficulty contacting the Meraki Cloud.
- A gateway AP switching to repeater mode.
- A gateway AP failing to rejoin as a gateway after reconnection.
- Connecting to a Meraki AP locally using its default SSID when it cannot reach the cloud.
Environment
- Hardware: Cisco Meraki MR series access points (including MR46 and other Wi-Fi 6 and newer APs)
- Network services: DHCP, DNS, ARP, outbound firewall access to the Meraki Cloud
- Local configuration tools: ap.meraki.com and my.meraki.com local status pages
Troubleshooting Meraki Cloud connectivity issue
An alert may appear on your AP detail page stating:
"This device is having difficulty contacting the Meraki Cloud. Please make sure your wired network allows outgoing connections to x.x.x.x and x.x.x.x on ports 443, 7734, 7351 and 7752."
When this happens, the AP icon (located in dashboard under Wireless > Monitor > Access Points) turns yellow, the connectivity graph stays green, and the AP does not download the latest firmware or configuration from the Meraki Cloud.
Possible causes
- The AP sits behind a firewall that blocks outbound access to the Meraki Cloud.
- The DNS servers the AP uses are unreachable, do not respond, or send invalid DNS responses.
- The upstream ISP-provided modem has traffic inspection security enabled.
- Dashboard is experiencing a temporary outage.
Troubleshooting steps
- 
    Check your firewall and confirm it allows outbound access to the cloud on ports 443, 7734, 7351, and 7752.
- 
    If the firewall settings are correct, investigate DNS and perform following steps:
- Change the DNS servers used by your AP to a public server (such as Google Public DNS)
- Confirm your firewall allows outbound DNS traffic (UDP port 53). If your AP uses a static IP address, refer to the static DNS settings documentation.
- 
    If the issue persists perform following steps:
- Enable traffic inspection security setting on your ISP-supplied modem or router.
- Reboot the modem or router, mostly fixes the issue temporarily.
- Comcast modems and SMC-manufactured routers ship with the Gateway Smart Packet Detection feature enabled by default, which is known to cause this condition. Contact your ISP or modem/router manufacturer for help disabling this security feature.
Expected outcome
The AP icon returns to green, the AP downloads the latest firmware and configuration, and the alert no longer appears on the AP detail page.
Troubleshooting gateway AP repeater-mode issue
A gateway AP is an access point with a wired interface configured with an IP address, connected to the LAN, and with a route to the Internet.
If the Internet is unavailable and the SSID allows LAN access, the AP continues to act as a gateway because it still holds a valid IP and can reach the local router or firewall.
Possible causes
The AP converts to a repeater only when one of these conditions is true:
- The AP cannot receive an ARP reply packet from the default gateway on the LAN (usually a local firewall or router).
- The AP cannot obtain a valid IP address via DHCP.
If the AP uses a static IP, it advertises "<ssid name>-bad gateway" when the default gateway is unavailable. With no other APs to mesh with that provide a route to the Internet, the AP remains offline. If it finds another AP with a route to the Internet, it acts as a repeater, but the dashboard reports an invalid IP configuration.
Troubleshooting steps
If an AP joins the dashboard as a repeater but you expect it to be a gateway, complete these verifications:
- 
    Connect a laptop with Wireshark installed to the switch port where the AP connects. Confirm the laptop receives an IP by DHCP and can ping its gateway.
- 
    If the laptop does not receive an IP by DHCP, troubleshoot until it successfully gets an IP address. Then reconnect the AP; it should join the dashboard as a gateway.
- 
    If the laptop receives an IP by DHCP but does not receive ARP replies from its default gateway, troubleshoot the default gateway accordingly.
- 
    If the AP needs a static IP address because no DHCP server is available on site, use local status page to configure it.
Expected outcome
The AP obtains a valid IP and reaches its default gateway, then joins dashboard as a gateway rather than a repeater.
Troubleshooting gateway AP reconnection issue
All Meraki access points dynamically monitor their uplink port for Ethernet connectivity. In a few scenarios, an AP that was once a gateway will not become a gateway again after you reconnect its Ethernet cable. When this happens, the AP signal LEDs scan back and forth, and an SSID appended with "-scanning" may appear in your wireless network list.
Possible causes
- A bad cable run or a cable run that is too long.
- A faulty PoE injector.
- The AP has no IP address
- A Layer 1, 2, or 3 problem on the switch port.
- The switch port is bad or administratively shut down.
Troubleshooting steps
- 
    Check the cable run: 
  - Verify the cable is securely connected on both ends.
  - Replace the current cable run, and confirm the total run is 100m or less (a physical limitation of CAT5/Ethernet cables).
- 
    Check the PoE injector. Swap the PoE injector in use with a known good PoE injector.
- 
    Check if AP is missing the IP address. If the AP is set to obtain an IP address automatically, the DHCP server may not be responding, may be unreachable, or may be out of IP addresses. 
  - Verify the DHCP server is running and reachable.
  - Verify the DHCP pool has addresses available for lease.
  - Configure the AP with a static IP address to see if it becomes a gateway; success indicates a problem with the DHCP service on the LAN.
- 
    Verify Layer 1, 2, and 3 on the switch port using a laptop. 
  - Disconnect the AP from the switch port.
  - Plug the laptop into the same switch port the AP used.
  - Verify the laptop obtains a DHCP address and can ping hosts on the Internet.
- 
    Check the switch port. If the port is bad or administratively shut down, connect the AP to a different port on the switch.
Expected outcome
The AP obtains an IP address, passes Layer 1, 2, and 3 checks on the switch port, and rejoins as a gateway without the "-scanning" SSID appearing.
Troubleshooting local connection using the default SSID
Both ap.meraki.com and my.meraki.com are locally hosted sites useful for configuring an access point (AP) when it cannot reach the Meraki Cloud.
Possible causes
- The AP is on a static, non-DHCP network.
- Strict firewall rules block the AP's connection to the Meraki Cloud.
- The AP has lost its Internet connection while still powered, causing it to broadcast a default SSID.
Troubleshooting steps
You can connect to the default SSID for administrative tasks in the Local Status Page by completing these steps:
- 
    Physically inspect the AP. 
  - Check that the AP has power (refer to the LED codes section of the MR installation Guides).
  - Copy the MAC address (refer to the Locating the MAC Address of Cisco Meraki Devices article).
- 
    Check for available wireless networks. 
  - Check whether a known default SSID is being broadcast.
  - If the AP has no configuration from the Meraki Cloud controller, the following is expected behavior:  
        
