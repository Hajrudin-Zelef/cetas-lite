---
id: collect-261001-meraki/meraki/prod-img-cisco-meraki-whitepaper-air-marshal-pdf-b396bacc-3
title: "prod-img-cisco-meraki-whitepaper-air-marshal-pdf-b396bacc"
domain: meraki
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["containment", "latency"]
source: docs/RAG/collect-261001-meraki/prod-img-cisco-meraki-whitepaper-air-marshal-pdf-b396bacc.md
source_anchor: ""
source_lines: [254, 374]
sha256: 842357efc1350700728d4ff04e8d814b3559ca2d309aaa36c9de65ffcd268098
---

# prod-img-cisco-meraki-whitepaper-air-marshal-pdf-b396bacc

Configuring Meraki’s Air Marshal  
WIPS platform
Dual-radio Meraki APs will run wireless scans opportunistically while also serving clients; this 
means they will scan the channel on which they are serving clients. It is possible to schedule 
‘mandatory’ scans to be run at pre-specified time intervals that can be set as frequently as once 
a day. For users requiring more accurate and real-time wireless threat assessments, it is possible 
to place an AP in Air Marshal mode. While acting as an Air Marshal, an AP will use its radios as 
dedicated scanners to monitor its surrounding environment in real-time. For the dual-radio APs, 
this includes both the 2.4GHz and 5GHz frequencies. Newer Meraki APs include a third radio 
which comes pre-configured for permanent Air Marshal scanning. These APs do not require any 
Air Marshal configuration and will scan and remediate against threats in real-time. 
Air Marshal mode can be switched on by selecting the relevant APs on the Access Points page. 
By clicking the relevant AP and selecting ‘On’ under the Air Marshal scanning section, it is 
possible to designate this AP as a dedicated WIPS scanner. This Air Marshal AP will now be a 
dedicated sensor performing scans of the surrounding environments for threats, the results of 
which will be displayed on the WIPS page in real-time. 
Figure 7 - Configuring Air Marshal APs on the Wireless > Access points page
A note on hybrid vs. dedicated scanners
APs with two radios running in client-serving mode will only scan the airspace opportunisti-
cally; this means they will scan the client-serving channel in real-time, and will scan across 
all channels either once a day or when no clients are being served. Most WLAN vendors 
recommend having dedicated scanning sensors (with no clients being served) in security-
conscious environments, to ensure real-time security alerting and protection. Some vendors 
offer ‘time slicing’ which allows cross-channel scans while serving clients, but this sacrifices 
performance of latency-sensitive applications such as VoIP and is generally not recom-
mended in the industry. For this reason, Meraki recommends placing an AP in dedicated  
Air Marshal mode (or utilizing newer 3-radio APs) for real-time scanning.
Cisco Systems, Inc.  |  500 Terry A. Francois Blvd, San Francisco, CA 94158  |  (415) 432-1000  |  sales@meraki.com
10
iPhone client accidentally 
associating to Rogue AP
iPhone client accidentally 
iPhone client accidentally 
Select an AP and tag it  
‘airmarshal’
3-radio APs come  
pre-equipped with  
Air Marshal radio
Shield icon indicates 
Air Marshal status

Figure 8 - Monitoring Air Marshal Page
On the Wireless > Air Marshal page, a number of manual actions of automated policies can be 
set as a response to the detection of certain types of wireless threats based on administrator 
preferences:
1. Manual rogue containment: when choosing to ‘contain’ a rogue SSID, the Meraki AP will 
perform containment (as described in ‘Threat Remediation’ section of this document) to 
render the rogue AP ineffective. Certain rogue SSIDs known as friendly APs can also be 
whitelisted to avoid confusion in the future. 
2. Automated LAN and Keyword Containment: automated policies can also be set to contain 
rogues seen on the wired network, as well as rogue APs matching a certain keyword. For 
example, if “Acme” is specified as a keyword and a Rogue SSID begins broadcasting an 
SSID named “AcmeCorp”, it will automatically be contained and clients will not be able to 
associate with it. This can be helpful in detecting people who are trying to copy the network 
with similar names and ‘trick’ clients into associating with their own AP.
3. Mandatory scan schedule: set time and days of the week where non-Air Marshal APs should 
scan all channels to ensure daily scanning. 
4. On the Wireless > Group policies page, a special policy can be created to track accidental 
associations by VIP clients; simply select the ‘track clients straying’ policy attribute and save 
the policy. The policy can then be applied to specific clients on the Clients page, and devices 
can also be pre-staged with their MAC addresses to have this policy automatically applied 
upon association by using the ‘Add devices’ function on the page. 
Cisco Systems, Inc.  |  500 Terry A. Francois Blvd, San Francisco, CA 94158  |  (415) 432-1000  |  sales@meraki.com
11
Configure LAN  
containment
Manual Containment/
Whitelist
Specify exact or key-
word matches
Add data on rogues: 
-VLAN 
-Manufacturer 
-Wired/wireless MAC 
-RSSI 
-Encryption type
Outlines # of APs in Air Marshal 
mode and with 3rd scanning radio

Figure 9 - Creating a Client tracking group policy
5. Generic alerts for Rogue APs can be set on the ‘Alerts and Administration’ page, allowing 
administrators to receive automatic alerts when rogue APs are detected that either match 
specified keywords or are seen to be on the wired LAN. 
Creating a WIPS response plan
By configuring alerts and utilizing Meraki’s Air Marshal view to monitor these threats retroactively 
and in real-time, it is possible to build a robust security plan that can be enforced. An example of 
a complete security methodology is as follows:
1. Create a WIPS plan as per your company’s security policies 
(i) Configure mandatory scanning intervals or designate APs to run in Air Marshal mode 
(ii) Configure auto-containment policies for rogue SSID keyword matches or rogues on the 
wired LAN 
(iii) Configure client straying policies to track batches of VIP clients 
(iv) Configure WIPS alerts
2. Proactive monitoring of Air Marshal 
(i) Visit Air Marshal page weekly or quarterly and mark known rogues as ‘whitelisted’, contain 
dangerous rogues 
(ii) Physically contain rogues that may be a threat
3. Reactive monitoring of Air Marshal alerts 
(i) Receive alert and react accordingly (set containment, find and contain rogue, etc). 
Cisco Systems, Inc.  |  500 Terry A. Francois Blvd, San Francisco, CA 94158  |  (415) 432-1000  |  sales@meraki.com
12
Create and apply a group  
policy to track specific clients

Conclusion
By understanding the spectrum of wireless security threats in today’s environment and creating 
a comprehensive response plan, network administrators can preclude the possibility of a serious 
compromise of critical network assets — including access to secure network devices that belong 
to the enterprise. A best-in-class WIPS platform should be capable of delivering intuitive reporting 
and monitoring, along with a robust suite of tools allowing for automatic alerts and security 
enforncement.
Meraki’s Air Marshal system includes real-time detection, remediation and alerting capabilities, 
including the ability to define pre-emptive policies that will intelligently take action to shoot down 
rogue APs using sophisticated containment mechanisms. Meraki’s wireless portfolio contains 
both dual-radio APs which can be converted into full-time sensors running in Air Marshal mode, 
and three-radio APs with dedicated scanning radios permanently running as Air Marshal scanners. 
By utilizing Meraki access points and Meraki’s intuitive web-based Dashboard interface, network 
administrators can create a robust WIPS policy plan, and easily deploy an airtight network to 
deliver enterprise-grade security in a WLAN environment. 
Cisco Systems, Inc.  |  500 Terry A. Francois Blvd, San Francisco, CA 94158  |  (415) 432-1000  |  sales@meraki.com
13
