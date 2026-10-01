---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-648164-ubiquiti-unifi-wirlesss-access-point-not-obtaining-ip-address-w-4dfaa23f
title: "questions-648164-ubiquiti-unifi-wirlesss-access-point-not-obtaining-ip-address-w-4dfaa23f"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-648164-ubiquiti-unifi-wirlesss-access-point-not-obtaining-ip-address-w-4dfaa23f.md
source_anchor: ""
source_lines: [1, 12]
sha256: 3c94130ff6db8ec7a979cbf8be1480d2e2db8bd1e02deff395f87d836661fae8
---

# questions-648164-ubiquiti-unifi-wirlesss-access-point-not-obtaining-ip-address-w-4dfaa23f

To start with, I was initially a bit confused because to me the way the router presents this configuration is a bit confusing. Primarily the "VLANX" terminology used in the listing down the left side. It finally clicked to me that these labels really has nothing to do with the actual 802.1Q VLANs and could be as easily been listed as "LANX", "NetworkX" or anything else. It is only the VID column that correlates to the actual VLAN in use.
To be clear, I will italicize any name that references the routers labeling throughout the rest of the answer like so: VLANX or LAN X.
In addition, whether a VLAN is tagged or untagged is determined by the "Enable" column (this has nothing to do with whether the VLAN itself is enabled/disabled as I initially thought). So if you want to use a VLAN on two ports as both tagged and untagged, you are required to use two of the VLANX entries.
Again to be clear, in the configuration pictured you only have one line (VLAN2) that indicates it is untagged. The configured VID is actually not used, and it is only the "Subnet" field that is important (in this case LAN 3).
On P4 in your diagram, you only have two VLANs configured on the port, both of which are tagged ("Enable" box is checked). What you still need is the VLAN for the untagged traffic from the AP (i.e. management traffic). This is proven by the fact that it "works" to some degree when plugged into P1 which in the picture has an independent check box that will always allow untagged traffic to reach the router on P1.
Based on what you provided, I see two cases:
- Your AP needs to receive 3 VLANs. This would include two tagged VLANs for the user VLANs (30 and 40) and an untagged VLAN for the management interface.
- You actually intend for the management interface of the AP to be on the same network as one of your two user groups and your AP needs to receive 2 VLANs. This would include one tagged VLAN for one user group and one untagged VLAN for the other user group as well as the management interface.
So in case #1 above, you need to configure another VLANX line for the management VLAN setting the correct "Subnet", leaving the "Enable" check box unchecked, and assign this to P4. (Again, with the "Enable" box unchecked, the VID field means nothing). This will now indicate three VLANs on P4.
In case #2 above, you need to uncheck the "VLAN" check box in the AP configuration; on SSID#1 if the management interface should be on LAN 4 (VLAN 40) and on SSID#2 if LAN 3 (VLAN 30). 
Then on the router, if this is the LAN 3 "Subnet", you should select VLAN2 for P4 and remove VLAN4 from P4. If this is LAN 4, then you need to configure another VLANX with the "Subnet" set to LAN 4, select this VLANX for P4 and remove VLAN3 from P4.
That should at least get you working. Without knowing a bit more about your network, I couldn't say if there might be a more ideal configuration.
