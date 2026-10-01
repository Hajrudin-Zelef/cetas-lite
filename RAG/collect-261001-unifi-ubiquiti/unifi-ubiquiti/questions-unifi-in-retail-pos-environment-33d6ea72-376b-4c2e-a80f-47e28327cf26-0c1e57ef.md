---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-unifi-in-retail-pos-environment-33d6ea72-376b-4c2e-a80f-47e28327cf26-0c1e57ef
title: "questions-unifi-in-retail-pos-environment-33d6ea72-376b-4c2e-a80f-47e28327cf26-0c1e57ef"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-unifi-in-retail-pos-environment-33d6ea72-376b-4c2e-a80f-47e28327cf26-0c1e57ef.md
source_anchor: ""
source_lines: [1, 29]
sha256: ffd40c89a4043317ceac7e0fe1ed65b54ca58e57bdea9710558dd25f25572629
---

# questions-unifi-in-retail-pos-environment-33d6ea72-376b-4c2e-a80f-47e28327cf26-0c1e57ef

@UI-Team
We are currently looking into alternatives to our current product line in our retail POS environments.
We currently use higher end but still soho un-managed access points and are looking for a product that would be easier to support, more reliable, longer range and cost effective.
Our company was (8+ years ago) installing high end Cisco APs but the decision was made to switch to soho mostly because the cost was prohibitive to our customers. To me this product line seems like a no brainier for our requirements.
As far as network load we are talking about 1-3 devices using very minimal amounts of bandwidth however we need to be able to cover large retail spaces, warehouses, backrooms etc. We also have the challenges of dealing with noise around refrigeration/utility/etc.
We also have the challenge of supporting the occasional customer who would like to be able to use their access points for their own devices or to allow public wireless. Currently the access points we use support multiple SSIDs/vlaning/etc.
I have seen a lot of praise and a lot of caution regarding this UniFi system but what I haven't been able to find are case studies. I'm interested to know what others have found and if they have been used in low traffic but high availability environments where roaming will be a concern.
Also has anyone dealt directly with Ubiquiti Works?
www.ubiquitiworks.com/
Are they a reputable dealer? After I put the request through our parts department to buy a few units for evaluation they reached out to Ubiquiti Works who basically told them the UniFi Product line is junk and what we want are the AirMAX products... As far as I am concerned the salesperson is trying to up-sell us to a more expensive product since as far as I can tell the AirMAX is for an entirely different application but they have our parts guy spooked.
I just got my 33 UAP-LR from Streakwave, Christie was my sales person and was Amazing! Tell her Charles sent you
😉
We are using these in a hospital deployment and are excited to replace the older cysco system.
I work for a chain of 6 retail stores. I switched earlier this year to UniFi APs throughout and couldn't be happier.
Some details on my setup:
Each building is around 20,000 Sq.Ft., 17,000 of that is retail space, the rest is warehouse/office space. Each building has 15ft ceilings and only wood/drywall separating the store/warehouse area. The store area is filled with your standard retail shelving (which if you look at it is mostly particle board and therefore mostly radio transparent)
I mounted a single Unifi AP-LR at ceiling level in the very center of the building. That single AP covers the entire building including the warehouse area better than the 2 SOHO APs I used before. The lowest signal I measured at any extreme is around -79. After seeing that I decided a single AP would work, and so far have run into no problems.
The wireless network is used primarily by a handful of Symbol/Motorola PDAs and Zebra QL series label printers. There are a handful of roaming salespeople that use laptops connected to it as well.
I have 2 SSIDs setup company wide. One for staff devices and one for guest access. The staff SSID runs untagged, but the guest SSID is VLAN tagged. I use Guest services in the UniFi software to speed limit the guest SSID. The firewall at each branch handles access control.
The controller software I have running on a VM at a central site we use as sort of a "data center". Each branch has a VPN tunnel back to this central site, and thus all the APs can see the controller.
Keep in mind that the UniFi APs do NOT handle roaming. If you have multiple access points setup in a given area, its up to the client which one they connect to. Though a upcoming feature in v3.0 called Minimum RSSI should help with this!
I work for a chain of 6 retail stores. I switched earlier this year to UniFi APs throughout and couldn't be happier.Some details on my setup:Each building is around 20,000 Sq.Ft., 17,000 of that is retail space, the rest is warehouse/office space. Each building has 15ft ceilings and only wood/drywall separating the store/warehouse area. The store area is filled with your standard retail shelving (which if you look at it is mostly particle board and therefore mostly radio transparent)I mounted a single Unifi AP-LR at ceiling level in the very center of the building. That single AP covers the entire building including the warehouse area better than the 2 SOHO APs I used before. The lowest signal I measured at any extreme is around -79. After seeing that I decided a single AP would work, and so far have run into no problems.The wireless network is used primarily by a handful of Symbol/Motorola PDAs and Zebra QL series label printers. There are a handful of roaming salespeople that use laptops connected to it as well.I have 2 SSIDs setup company wide. One for staff devices and one for guest access. The staff SSID runs untagged, but the guest SSID is VLAN tagged. I use Guest services in the UniFi software to speed limit the guest SSID. The firewall at each branch handles access control.The controller software I have running on a VM at a central site we use as sort of a "data center". Each branch has a VPN tunnel back to this central site, and thus all the APs can see the controller.Keep in mind that the UniFi APs do NOT handle roaming. If you have multiple access points setup in a given area, its up to the client which one they connect to. Though a upcoming feature in v3.0 called Minimum RSSI should help with this! 
Thanks for you reply!
It sounds like you guys are basically doing the same stuff we are going to do with it.
Our test unit is on the way.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
