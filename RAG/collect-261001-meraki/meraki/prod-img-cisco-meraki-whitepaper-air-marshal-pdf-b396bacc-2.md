---
id: collect-261001-meraki/meraki/prod-img-cisco-meraki-whitepaper-air-marshal-pdf-b396bacc-2
title: "prod-img-cisco-meraki-whitepaper-air-marshal-pdf-b396bacc"
domain: meraki
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["containment", "incident", "parameters"]
source: docs/RAG/collect-261001-meraki/prod-img-cisco-meraki-whitepaper-air-marshal-pdf-b396bacc.md
source_anchor: ""
source_lines: [124, 253]
sha256: c8fc982ad67eda1920829a91e2b8fb7297dbf4f90348a65610d7f5d7890de7a5
---

# prod-img-cisco-meraki-whitepaper-air-marshal-pdf-b396bacc

Client Straying Threats
Accidental associations: client devices that belong to your infrastructure associating to a wireless 
network in your airspace that has not been sanctioned by your corporation. Straying clients could 
accidentally connect to Rogue SSIDs or spoofed SSIDs if proper action is not taken to protect the 
wireless airspace.  
PCI Compliance
Understanding and remediating against wireless threats is also a requirement under the Payment 
Card Industry Data Security Standard (PCI DSS), a standard required for retailers to follow when 
processing credit card data over WLAN networks. Examples of WIPS requirements under PCI DSS 
include:
Section 9.1.3 Physical Security: Restrict physical access to known wireless devices.
Section 10.5.4 Wireless Logs: Archive wireless access centrally using a WIPS for 1 year.
Section 11.1 Quarterly Wireless Scan: Scan all sites with card dataholder environments (CDE) 
whether or not they have known WLAN APs in the CDE. Sampling of sites is not allowed. A WIPS 
is recommended for large organizations since it is not possible to manually scan or conduct a 
walk-around wireless security audit of all sites on a quarterly basis
Section 11.4 Monitor Alerts: Enable automatic WIPS alerts to instantly notify personnel of rogue 
devices and unauthorized wireless connections into the CDE.
Section 12.9 Eliminate Threats: Prepare an incident response plan to monitor and respond to 
alerts from the WIPS. Enable automatic containment mechanism on WIPS to block rogues and 
unauthorized wireless connections.
Figure 3: Security threats in an enterprise environment
Cisco Systems, Inc.  |  500 Terry A. Francois Blvd, San Francisco, CA 94158  |  (415) 432-1000  |  sales@meraki.com
6
Site A
Client AP
Rogue AP Client Laptop 
‘Accidental association’  
High-value asset associations 
to wrong network
Client AP
Ad Hoc 
Devices in ad-hoc mode can 
connect to a client AP and 
create a gateway for wireless 
hackers
Wireless Hacker
Rogue AP 
Rogue APs on  
the wired LAN can  
compromise your 
entire wired and 
wireless network
Site B

Threat Remediation using Meraki’s  
Air Marshal WIPS platform
A careful study of the common wireless security threats has led to the development of Meraki’s 
Air Marshal platform, which allows access points to be turned into dedicated WIPS sensors called 
‘Air Marshal’ APs. Air Marshal is a WIPS platform which comes equipped with security alerting and 
threat remediation mechanisms. This includes the following:
a. Monitoring and alerting: a robust and intuitive display of all of the threats for a particular 
network, including auto-alerting based on the network administrator’s preferences. Monitoring 
techniques include:
i. Rogue AP monitoring: Meraki APs scan across all 2.4 GHz and 5 GHz channels to build 
a list of rogue access points in the nearby vicinity. In addition, further mechanisms are in 
place to track APs on the wired LAN network by inspecting traffic on the wired port of the 
Meraki AP, and using this to build a list of rogue APs that may be on the wired LAN. E-mail 
alerts will be triggered and sent based on parameters predefined by the network admin. 
ii. Tracking ‘client straying’ of VIP clients: Air Marshal allows tagging of VIP clients and 
an alert is sent if those clients connect to a unsanctioned SSID. Air Marshal does this 
by monitoring traffic with the source MAC address of the VIP clients. Wireless devices 
communicate with three types of 802.11 frames: management frames are used during the 
probing and association process. Control and data frames are used when the client is 
actually connected. If Air Marshal sees data frames originating from VIP clients which are 
not connected to the corporate wireless network, an alert can be sent to administrators for 
remediation.  
b. Remediation mechanisms: Air Marshal APs come equipped with the ability to automatically 
‘contain’ rogue APs and alert on rogue APs and accidental associations, allowing for 
administrators to take physical action to remove rogue APs and recover straying devices. 
Figure 4: Tracking accidental associations 
Cisco Systems, Inc.  |  500 Terry A. Francois Blvd, San Francisco, CA 94158  |  (415) 432-1000  |  sales@meraki.com
7
Client accidentally  
associates to Rogue AP
Air Marshal AP detects 
data frames exchanged
Email alert sent to  
network administrator

What is containment?
‘Containment’ is a common mechanism that calls for the Air Marshal AP to impersonate or spoof 
the rogue AP in order to render it ineffective. Air Marshal does this by generating a large number 
of 802.11 packets and using the BSSID of the rogue AP as the the source MAC address. Air 
Marshal APs also provides more sophisticated containment methods including spoofing clients 
attempting to associate to the rogue by generating packets with the source MAC of the clients; 
this allows for a ‘two-way’ spoof and ensures a fool-proof shutdown of the rogue AP. 
Packet types generated by WIPS during containment: 
1. 802.11 Broadcast deauthorizations with source = Rogue AP, destination = broadcast 
2. 802.11 Deauthorization messages with source = Rogue AP, destination MAC = client  
3. 802.11 Deauthorization and disassociate messages with source = client, destination = Rogue AP 
#3 ensures that more sophisticated 802.11 clients with battery-saving capabilities are also unable 
to connect to the rogue, as they may ignore deauthorization messages from the Rogue AP if they 
are ‘sleeping’ in order to save battery life. 
2As containment renders any standard 802.11 network completely ineffective, containment 
measures should taken in your airspace. Extreme caution should be taken to ensure that 
containment is not being performed on a legitimate network nearby and, action should only be 
taken as a last resort. Unauthorized containment is prosecutable by law (subject to the FCC’s 
Communications Act of 1934, Section 333, ‘Willful or Malicious Interference’).  
http://transition.fcc.gov/Reports/1934new.pdf
Figure 5. Rogue AP ‘Containment’ explained 
Cisco Systems, Inc.  |  500 Terry A. Francois Blvd, San Francisco, CA 94158  |  (415) 432-1000  |  sales@meraki.com
8
802.11 packets
Packet Types
1. Rogue AP sending broadcast 
deauthorizations
2. Rogue AP deauthorizing client
3. Client deauthorizing rogue AP
Air Marshal AP
Rogue AP thinks client is 
requesting deauthorization
Client thinks rogue AP is  
forcing deauthorizing

Figure 6 - Example of containment technique - 802.11 Deauthentication packet
Cisco Systems, Inc.  |  500 Terry A. Francois Blvd, San Francisco, CA 94158  |  (415) 432-1000  |  sales@meraki.com
9
iPhone client accidentally 
associating to Rogue AP
802.11n packet type is 
deauthentication, source is 
Rogue AP and destination 
is client
Repeated deauthentications  
sent within short time period
Dummy packet is generated 
by our WIPS sensor, looks 
like it came from Rogue AP

