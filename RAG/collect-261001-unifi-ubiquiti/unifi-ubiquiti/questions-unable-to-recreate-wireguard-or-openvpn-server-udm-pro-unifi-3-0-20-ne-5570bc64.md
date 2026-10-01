---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-unable-to-recreate-wireguard-or-openvpn-server-udm-pro-unifi-3-0-20-ne-5570bc64
title: "questions-unable-to-recreate-wireguard-or-openvpn-server-udm-pro-unifi-3-0-20-ne-5570bc64"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-unable-to-recreate-wireguard-or-openvpn-server-udm-pro-unifi-3-0-20-ne-5570bc64.md
source_anchor: ""
source_lines: [1, 24]
sha256: ca3508d3139f7fd5d91e776b31b94e1b2d1538459c1f9b312a6d8d4afda360a2
---

# questions-unable-to-recreate-wireguard-or-openvpn-server-udm-pro-unifi-3-0-20-ne-5570bc64

@UI-Team
I am currently using UDM-Pro with 3.0.20 Unifi OS and on the Network Application 7.4.156.
Initially, when the Wireguard server was made available, I created the server to use my primary WAN Ip address and port 443 and it worked intermittently. When I say intermittently, some of my devices were able to connect, but the rest (mostly Macbooks) had an issue with a handshake (upon research, I see a lot of people are having the same problem).
When Network Application 7.4.156 was released, I wanted to test Open VPN, so I created the server using the same primary WAN but with port 1197. OpenVPN worked like a charm on all devices. That being the case, the plan was that if OpenVPN worked on all devices, then I would delete Wireguard and use OpenVPN as my primary VPN on port tcp/443.
I deleted the Wireguard server and tried to change the OpenVPN port to 443, but it didn't allow it. So I deleted the OpenVPN server and tried recreating the server with port 443, and still, it is not working. I am not able to even recreate WireGuard on port 443.
I get a vague error like below (nothing verbose).
Any thoughts on how I can solve this issue?
Hello @outerlimits,
Port validation was added, we prevent users from using ports that are used/reserved by applications/OS.
Oh ok. Thanks, @UI-Glenn. So there is no way we can run a VPN server on TCP/443 or common (potentially unrestricted by Public Wifi/Companies) ports?
@UI-Glenn Just as an update. Different ports also don't work and getting the same error. I tried Port 8080 and 8443 for both WireGuard and OpenVPN.
Those ports are also reserved (fall under use by application/OS), so try port 8500 or something like that.
Thanks, @Stillabeginner I think that was the issue. I used 8500 for WireGuard and 9443 for OpenVPN, and it worked fine now. Not the ideal solution but its what it is.
@UI-Glenn wrote:
This breaks multiple production sites for me, as the only port that should be listening on the external interface is 443 for WireGuard.
@UI-Glenn - Is there a process to bypass port validation?
Is the list of validated/reserved ports published? I didn't find it with a quick search of the support docs. I kinda understand 443 as that's where the webmin lives, but locking it down forever seems.....less than ideal.
@mausball wrote:
Since the implementation of port validation is new, I doubt there is such a list. Every installation will be different based on ones configuration of the network, severs, services, etc.
This makes wireguard almost useless on most public wifi....
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
