---
id: collect-261001-meraki/meraki/prod-img-cisco-meraki-whitepaper-air-marshal-pdf-b396bacc-1
title: "prod-img-cisco-meraki-whitepaper-air-marshal-pdf-b396bacc"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "containment", "copyright"]
source: docs/RAG/collect-261001-meraki/prod-img-cisco-meraki-whitepaper-air-marshal-pdf-b396bacc.md
source_anchor: ""
source_lines: [1, 123]
sha256: ac59f7cf048cd51a19c6d132d51b4c20257e6a862f5aaad5ea3371457d8ae7fc
---

# prod-img-cisco-meraki-whitepaper-air-marshal-pdf-b396bacc

White Paper
Air Marshal 
SEPTEMBER 2013
This document discusses potential security threats in a WiFi 
environment, and outlines how enterprises can use a best-in-class 
Wireless Intrusion Prevention System (WIPS) such as Meraki’s Air 
Marshal to create a secure network with preemptive protection policies 
and security alerts.

Copyright 
© 2013 Cisco Systems, Inc. All rights reserved
Trademarks 
Meraki® is a registered trademark of Cisco Systems, Inc. 
Table of Contents Introduction     3  
Wireless Threats     5  
Threat Remediation     7  
Configuration      10  
Conclusion      13  
 
 
Cisco Systems, Inc.  |  500 Terry A. Francois Blvd, San Francisco, CA 94158  |  (415) 432-1000  |  sales@meraki.com
2

Introduction
Wireless Security Threats in an Enterprise Environment
Secure WiFi access has become a critical component of enterprise networking. WiFi Internet 
access is critical for corporate communication in verticals including financial services, retail, and 
distributed enterprise. Due to the widespread use of WiFi and variety of use cases (e.g., point-of-
sale (POS) communications, corporate access, warehouse inventory, asset tracking, WiFi services 
for targeted advertising), the wealth of information transmitted across the wireless medium 
has skyrocketed. Data transmitted over wireless increasingly contains sensitive personal and 
financial data. Unfortunately, the tremendous growth in wireless has been accompanied with 
an increasingly widespread ability to obtain open-source hacking tools that can compromise a 
wireless network through impersonation of client devices and access points.
Examples of common threats in a modern WiFi environment include:
Network impersonation: achievable by purchasing any consumer-grade access point and 
copying an SSID, “tricking” clients into thinking that this SSID is available and snooping on their 
information transactions. 
Figure 1: Example of SSID spoofing threat in a retail environment
Legitimate SSID Malicious SSID Unsuspecting user  
connects to malicious 
SSID
Cisco Systems, Inc.  |  500 Terry A. Francois Blvd, San Francisco, CA 94158  |  (415) 432-1000  |  sales@meraki.com
3

Wired network compromise: achieved by an unsuspecting employee or student plugging in a 
consumer-grade access point into the wired infrastructure and exposing the LAN to hackers.
To successfully protect an enterprise network, a Wireless Intrusion Prevention System (WIPS) 
should provide powerful wireless intrusion scanning capabilities, enabling detection and 
classification of different types of wireless threats, including rogue access points and wireless 
hackers. Access points can be put in either dedicated scanning mode or sensor mode for 
real-time intrusion detection and threat remediation. Additionally, a WIPS system should be 
configurable with intuitive auto-containment policies to facilitate pre-emptive action against rogue 
devices. Once a threat has been detected, the WIPS platform should kick into gear to enact 
powerful policies, including intelligent auto-disablement of APs matching a pre-defined criteria 
and generating different tiers of e-mail alarms based on the type of threat in your airspace. 
In addition to protecting airspace against hackers with malicious intent, it is valuable to be able 
to mark or group high-value ‘VIP’ wireless clients into a special category, where they can be 
tracked, to ensure they never leave the wireless network. Examples of VIP clients include high-
value corporate assets — these devices belong to the organization and should never associate 
to a wireless network other than your own (e.g., corporate issued laptops, point-of-sale registers 
or barcode scanners in a retail environment, etc.) These VIP clients should only associate to the 
corporate network. If they do stray to a foreign network, it would be classified as an ‘accidental 
association.’ The ability to detect and generate alerts when a VIP client strays over to a rogue 
infrastructure can be invaluable in security conscious environments.
Cisco Meraki’s Air Marshal mode allows network administrators to meet these requirements and 
design an airtight network architecture that provides an industry-leading WIPS platform in order to 
completely protect the airspace from wireless attacks. The remainder of this document describes 
in greater depth wireless threats and the necessary security measures required to remediate 
against these threats; the conclusion then summarizes the setup and configuration process for 
Meraki’s Air Marshal WIPS platform in order to achieve the highest security protection possible. 
 
Figure 2: Example of accidental wired LAN compromise in corporate environment
Cisco Systems, Inc.  |  500 Terry A. Francois Blvd, San Francisco, CA 94158  |  (415) 432-1000  |  sales@meraki.com
4
Internet
Corporate network
Unsuspecting employee 
plugs in home AP to create 
wireless access
User gains access to 
corporate LAN resources

Wireless Threats
Understanding the wireless airspace around you can help to take effective measures, both 
preventive and reactive, to ensure that the wireless airspace is secure and interference-free 
from other wireless networks. A number of different threats exist in the modern enterprise 
environment, facilitated by easy access to cheap consumer-grade 802.11 equipment, along with 
open-source hacking tools that can be used to simulate and spoof devices and generate traffic 
floods. Leading enterprise WLAN providers such as Cisco provide built-in WIPS features to ensure 
detection and remediation against these threats. 
Threat classifications
Visibility and classification of potential wireless threats is an important first step in securing the 
wireless network and network infrastructure as a whole. Once classified, remediation can be 
taken against confirmed threats and innocuous alerts can be dismissed. Cisco Meraki Air Marshal 
automatically classifies threats into the following categories to provide the greatest visibility and 
overall protection for your network.
Rogue SSIDs
SSID and AP spoofs: the malicious impersonation of a legitimate AP by either spoofing the SSID 
name or, even worse, the SSID name and the BSSID (the wireless MAC address, which makes it 
indistinguishable from the original AP).
Rogue SSID seen on LAN: SSIDs that are broadcast by rogue APs and seen on wired LAN; this 
could suggest compromise of the wired network.
Other SSIDs
Interfering SSIDs: wireless networks that are broadcasting and could be causing RF interference, 
as well as attracting accidental associations from clients who are supposed to be connecting to 
your own network. 
Ad-Hoc SSIDs: modern smartphones and mobile devices are capable of associating to WiFi 
networks and then re-broadcasting the SSID, essentially acting as a wireless bridge. Devices in 
ad-hoc mode can connect to a client AP and create a gateway for wireless hackers.
Malicious Broadcasts
Denial of Service (DoS) attacks are attempts to prevent clients from associating to the legitimate 
AP by sending an excessive number of broadcast messages to clients. DoS attacks could be from 
malicious clients, APs, or even another WIPS system in the area that considers the corporate 
network a threat and is attempting to remediate.
Packet Floods
Clients or APs that are sending an excessive number of packets to your AP. Packets are 
monitored and classified based on multiple categories including beacon, authentication and 
association frames. An excessive number of any category of packets seen within a short time 
interval will be marked in Air Marshal as a packet flood. 
Cisco Systems, Inc.  |  500 Terry A. Francois Blvd, San Francisco, CA 94158  |  (415) 432-1000  |  sales@meraki.com
5

