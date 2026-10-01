---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-site-to-site-vpn-between-pfsense-and-usg-possible-in-gui-83b426d1-3914-efd75acc
title: "questions-site-to-site-vpn-between-pfsense-and-usg-possible-in-gui-83b426d1-3914-efd75acc"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Intel", "United States"]
dates: []
keywords: ["cost", "intel"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-site-to-site-vpn-between-pfsense-and-usg-possible-in-gui-83b426d1-3914-efd75acc.md
source_anchor: ""
source_lines: [1, 24]
sha256: 8634a79de650ff1cc1badbf702dedd0409799813311c456c73030e7c6dfad09e
---

# questions-site-to-site-vpn-between-pfsense-and-usg-possible-in-gui-83b426d1-3914-efd75acc

@UI-Team
I’m responsible for IT in our schools. For our secundary school we chose a couple of years ago for Rukus together with pfSense, no bad words, its working without any problems. At home I have pfsense with an US-8-150W switch and a UAP-AC-LR-1T AP. Between my home and the secundary school I have site to site vpn over OpenVPN with pfSense and in VLAN10 at home I have the same SSID as in my secondary school. Our primary schools don’t have the money for Rukus so last year I decided to go for Ubiqiti and I am as happy with it is as I am with Ruckus for a fraction of the Rukus cost. I have two primary school sites, a small one with a dynamic IP, only one AP and an USG-3P and a larger one with a Static Public IP, an USG-3P, an US 24p and 8p POE, a cloud key and 7 AP’s. Between my two locations I have an automatic VPN between the USG’s. My three Ubiqiti sites (home and primary schools) are on the cloudkey. I put the AP of the small school in the controller of the large school so I gets the same settings as the bigger school. Now my problem: I want to set up a site to site VPN between my pfSense at home and the USG. With two pfSense appliances it is easy, the one with static IP (secundary school) is server and my pfSense at home (dynamic IP) is client. But how do I do it with the USG (with the static IP) in the GUI, I don’t want to mess with CLI or putting files in my controller like the post in the unifi stories. Is it possible in the GUI with version 5.4.11? Or do I have to make my home pfSense (with the dynamic IP) OpenVPN server and use dynDNS to let the USG connect to my home instead?
I contacted support and it is not possible, it's a known bug in USG OpenVPN, they are working on it.I succedeed to connect my pfSense to USG using the GUI with IPsec IKEv2
Here are some screenshots on how to configure
Key exchange only works with IKEv2
First I had a strong PW with 16 characters which did not work, 14 characters was no problem...
FQDN is not working as peer IP, hopefully they will fix this as I have a dynamic IP at home...
unselect PFS and dynamic routing
Unifi side:
pfSense side:
greetings from Belgium ;-)
Yup, I was about to say im doing this right now with Pfsense and a USG but not using OpenVPN - just normal IPSec IKEv1 .
Pretty stable and getting about 20Mbps on it using my Intel Atom D510 on PfSense. When doing site to site with a Cisco RV082 and the USG3 I got a big higher speeds but I haven't tweaked my PFsense yet to see if I can get more out of it.
Strange as pfSense can handle higher speeds and more concurrent sessions as RV082, as for the tunnel I only got results with IKEv2, IKEv1 was not working with my setup
in phase 2 proposal (PFsense) i dont see you setting mine are now as follow.
care to share :-)
I'm seeking a similar configuration.
My USG is the home router. I run pfsense as a VPN server on a VPS in the cloud.
I would like to route all traffic from a particular network on the USG over the VPN. What settings should I change from your example config to accomplish this?
This vpn system on the USG is litterly the worst. Best practice is to install it in the bottom bottom of a landfill. Then set the landfill on fire.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
