---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/t5-unifi-routing-switching-dpi-impact-on-usg-pro-inter-vlan-routing-td-p-2393802-c50e47f9
title: "t5-unifi-routing-switching-dpi-impact-on-usg-pro-inter-vlan-routing-td-p-2393802-c50e47f9"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/t5-unifi-routing-switching-dpi-impact-on-usg-pro-inter-vlan-routing-td-p-2393802-c50e47f9.md
source_anchor: ""
source_lines: [1, 17]
sha256: 8893a85cb1136acaabf36fcad83174ae8db5ab34a08da7d0a3b13bf67e959869
---

# t5-unifi-routing-switching-dpi-impact-on-usg-pro-inter-vlan-routing-td-p-2393802-c50e47f9

@UI-Team
I have deployed a USG Pro and 2 Unifi switches (16 Port POE and 24 Port) for a client. The network has 3 VLANs, Guest, Hosts and Servers. I have DPI turned on. On installation I have been backing up a lot of files to the server and speeds have been around 25Mb/s. When the server was tested before installation in our office (on different switches) circa 110 Mb/s was achieved, as measured by windows. I am aware of the speed restrictions for the USGs when DPI is enabled.
Is DPI inspecting all traffic passing through the router or just the traffic destined for the WAN port?
My impression is that DPI is active on all routed traffic through the USG, and that yes DPI will count packets going inter-vlan. However that doesn't explain your issue, as DPI functions with hardware offload and should only show a negligible bandwidth hit. 25mbits/s is really slow (do you means 25 mega-bytes/sec?), so its not DPI causing this.
Disabling hardware offload would reduce a USG-Pro to around 250-300mbits (close to 25MBytes/sec), so I suspect this is the issue. Can you confirm that hardware offload is not disabled in the USG? Features that disable offload include IPS/IDS, some custom routing through CLI, bridged USG ports, etc. Look under the advanced tab for the USG under devices in the controller, or SSH and enter "show ubnt offload"
Hello,
DPI checks every single package on the WAN interface.
Regards,
Glenn R.
So the speed of inter VLAN routing should be unaffected by DPI as this traffic does not pass through the WAN interface?
Shouldn't be..
It only reads WAN packages that go through the WAN
It was IPS. I have disabled this for now and speeds for inter VLAN routing are at around line speed now. Thanks for the help.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
