---
id: collect-261001-cisco/cisco/questions-ubiquiti-unifi-security-gateway-usg-to-cisco-rv042-vpn-a554b2b8-036d-4-cc893eba
title: "questions-ubiquiti-unifi-security-gateway-usg-to-cisco-rv042-vpn-a554b2b8-036d-4-cc893eba"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-ubiquiti-unifi-security-gateway-usg-to-cisco-rv042-vpn-a554b2b8-036d-4-cc893eba.md
source_anchor: ""
source_lines: [1, 25]
sha256: d59328d700f1c9024ce4af53180711407302efb72164bcd7373b7e2e8d667b0b
---

# questions-ubiquiti-unifi-security-gateway-usg-to-cisco-rv042-vpn-a554b2b8-036d-4-cc893eba

@UI-Team
I have about a dozen RV042 routers currently doing site-2-site VPN tunnels. Seeing as many internet connections are now exceeding the bandwidth these devices are capable of combined with them being EOL we are looking to upgrade.
Since I've started using Unifi WiFi APs and like the single dashboard we are considering using the Ubiquiti Unifi Security Gateway (USG) as a replacement. Seeing as we will not be replacing all at once I need to make sure we can integrate the 2 while we do the cutover (which may take some time as in sites where there is no real need to upgrade just yet we will be leaving as-is).
What are thoughts on this? Easy or hard? Can be done through the dashboard or will require CLI customizations?
All sites are a single flat network allowed to communicate to devices across the VPN as well as direct to the Internet. All single WAN.
I had a USG-PRO-4 site-to-site VPN set up to a Cisco RV180W, which is similar to the Cisco RV042. I've since changed out the Cisco RV180W to a USG, but the VPN conneciton was solid to the RV180W and it was set up using the dashboard (no CLI required). At the time I was using Manual IPsec with:
Key Exchange = IKEv1
Encryption = 3DES
Hash/Authentication = SHA-1
DH Group = 2
Should be fairly straightforward and you should be able to do it without any CLI. You can setup Manual IPsec vpns on the controller without messing with any CLI.
Thanks guys, I'm going to order one to try out today.
Jeremy, thanks for the extra details, that will save me some time.
If anyone else has any additional details they feel is relevant please feel free to post them as well.
Its not working for me, could you detail all the setup in the USG and RV042? Many thanks!
The Cisco has been removed from the network, but these are the settings of when it was working (IP addresses and preshared key have been changed).
My on-going issue with this (and I was going to post as a seperate topic) is how to get this to work using FQDN instead of IP addresses. Most of my endpoints are DHCP from their ISPs so putting in static IPs isn't feasible.Also going to post in yet another thread (cause it's not directly related to my other quiries) is how to have multiple USGs in a hub layout. I can get Site-A to talk to Site-B and Site-B connected to Site-C but when I try to add a Site-A to Site-C connection it errors saying it is a duplicate.Once I can get pat the above 2 issues I will start ordering some USGs to replace my RV042s.
Many thanks! Your RV180W is very different from the RV042. I setup as you mention but still the VPN does not work with USG firmware version 4.3.61
One thing I noted is that you use Main mode instead of Aggressive.
Looking forward for your other posts.
Depends on what version of the RV042. The older models looked just like the screenshot in question (the old LInksys interface). Later revisions were updated with newer menus (the Cisco interface). I have a mix of both kinds of RV042s.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
