---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-unifi-dream-machine-udm-home-small-business-configuration-guide-with-s-3cd5e628
title: "questions-unifi-dream-machine-udm-home-small-business-configuration-guide-with-s-3cd5e628"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-unifi-dream-machine-udm-home-small-business-configuration-guide-with-s-3cd5e628.md
source_anchor: ""
source_lines: [1, 40]
sha256: be17b6f94b85e27f96064080f6d12e07b903b779fecb2fa684d871d9cdf4b238
---

# questions-unifi-dream-machine-udm-home-small-business-configuration-guide-with-s-3cd5e628

@UI-Team
I was looking for a guide that shows how to use the UDM features to securely configure a home/small office but most information is around describe each feature rather than how to use them to enable connectivity but is done in such a way that it maintains a reasonable good security baseline.
Does anyone have a summary or the steps they followed to configure the UDM once they had gone through the initially setup using default UDM settings? Better still it would be great benefit to work through the below use case on how you would go about achieving it.
Use Case: Home Network/Small Business Office that is connected direct to the Internet and has fixed and WiFi connected devices
Devices:
TVs connected via fixed eth and Wifi
Google Home connected via WiFi
Solar Inverter connected via WiFi
Computers connected via fixed eth
Laptops connected via WiFi
Mobiles connected via WiFi
Printer connected via fixed eth
Network Storage device via fixed eth
Network:
UDM connected to Internet via ISP WAN connection
Network LAN switch that supports VLAN configuration
Fixed devices connect to UDM or switch
WiFi devices connect via UDM WiFi
Security
From a security perspective we want to allow normal home/office applications and devices to continue to operate but limit internal and external security related risks by enabling security controls within the UDM.
I assume you would start at the network layer and start separating out the different device traffic types then focus on enabling threat management and firewall rules to prevent common threats now that the UDM has been initially setup.
What steps and in what order would you go about tackling this?
Steps required
Regards Nomadz
Hi,
Good KB regarding IPS/IDS settings below:
If it's SOHO, then you would want to have two VLANs for Corp and Guest traffic (as a bare minimum).
Before introducing any feature or complexity, ensure you have base config working and all devices can connect and say connected.
Thanks,
Myky
| PM for consultancy | CWSP | GMT time zone
Thanks Myky, that link is very useful for describing what each feature does but what is needed next is how you go about securing your network using these features. What would you protect first by using that feature:
Intrusion Detection and PreventionBy implementing IDS/IPS at different Mode levels, does that automatically configure the other options below or is that in addition to? if so what order would you enable and config them? Firewall RestrictionsSystem Sensitivity Levels
Does the above settings have to be configured to work on specific vlans so that we can put IOT devices into its own vlan and monitor those and do a different configs for computers that are fixed and or WiFI connected?
Thanks in Advance Nomadz
Any more info on this? I just set up a UDM base for my home network, and could use some basic advice for security, especially with regard to configuring the firewall.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
