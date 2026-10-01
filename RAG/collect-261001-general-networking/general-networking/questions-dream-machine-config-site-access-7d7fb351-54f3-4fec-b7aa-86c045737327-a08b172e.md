---
id: collect-261001-general-networking/general-networking/questions-dream-machine-config-site-access-7d7fb351-54f3-4fec-b7aa-86c045737327-a08b172e
title: "questions-dream-machine-config-site-access-7d7fb351-54f3-4fec-b7aa-86c045737327-a08b172e"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-dream-machine-config-site-access-7d7fb351-54f3-4fec-b7aa-86c045737327-a08b172e.md
source_anchor: ""
source_lines: [1, 17]
sha256: 3a0c1b0a4c1410519a9b90388cd4831e9d9ccc310151f5b9dc37c9102d4de529
---

# questions-dream-machine-config-site-access-7d7fb351-54f3-4fec-b7aa-86c045737327-a08b172e

@UI-Team
Hello,
I'm relatively new to Unifi and we are planning on creating a new network within an existing one and the Dream Machine is planned to be the firewall for this new section.
I already set up the LAN and WAN ports however, I would like to have the option to access the switch from outside vlans, mainly to configure. From searching online I was able to do this by setting up the default network to be the same as the LAN network it connects to and disable most network features. Thing is while I can reach the configuration site, I can only do so from the same VLAN. I would prefer the UDM LAN address and the workstation vlan to be separate, also I would like to access the UDM from other networks too.
From what I noticed this is definietly a setting on the UDM, I tried different vlans and the result is the same, there are full access between vlans other than the UDM. I assume this is a firewall setting, and the guides I followed told me I should make an Any-Any Allow rule to make sure, however I don't think that is fully possible in the latest version, and I would prefer to control who can access the UDM and that the UDM doesn't spam the already existing network trying to take it over.
So my quesion is, where can I change how to access the UDM, which firewall rule determines this, do I need to add the workstation networks to the UDM as well?
Thank you in advance!
You want to be able to access the UDM from networks on the WAN side?
It sounds like you may have disabled NAT? This is not likely to be what you wanted. I'd undo those changes or just factory reset and setup again.
If you enable remote access, you should have no need to change anything, just go to https://unifi.ui.com and you're done.
If you want to do direct access from the LAN networks on the WAN side, you just need a firewall rule to allow traffic from the external zone to the gateway zone. Your rule can limit traffic as tightly or as loosely as you want.
The question about accessing the configuration site on a Dream Machine is useful, especially when trying to manage the device locally. It would be helpful to clarify whether the issue is with reaching the UniFi interface, authentication, or access from a specific network or device.
Currently it's the Unifi interface I'm trying to reach using its LAN interface, I would prefer not using the WAN interface just yet because we are only giving it limited internet access until final launch.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
