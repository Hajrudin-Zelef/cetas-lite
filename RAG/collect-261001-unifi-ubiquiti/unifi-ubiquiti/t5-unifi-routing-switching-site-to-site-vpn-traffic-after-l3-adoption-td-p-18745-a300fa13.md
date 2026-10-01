---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/t5-unifi-routing-switching-site-to-site-vpn-traffic-after-l3-adoption-td-p-18745-a300fa13
title: "t5-unifi-routing-switching-site-to-site-vpn-traffic-after-l3-adoption-td-p-18745-a300fa13"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/t5-unifi-routing-switching-site-to-site-vpn-traffic-after-l3-adoption-td-p-18745-a300fa13.md
source_anchor: ""
source_lines: [1, 10]
sha256: 73d54052670ac621d75e0881cc78929c449b1e3efef52f0ebea3f6794c143c33
---

# t5-unifi-routing-switching-site-to-site-vpn-traffic-after-l3-adoption-td-p-18745-a300fa13

@UI-Team
I have two sites under one controller. site A (controller) and site B with a bunch of APs.
I setup a site-to-site vpn between the sites. My question is, will all the traffic/stats from site B (where controller is) will be sent to site A via the tunnel such as DPI data, traffic, and stuff that is in the dashboard?
Or none of it is encrypted since it uses ip:8080/inform
I think it depends on how you set the inform address. If you used an internal IP, then it would go over the VPN, if not, then the external.
Be careful if you use the internal, because if the VPN goes down, and the device can't reach back using the inform address, then you can't reconfigure it using the controller. This would be extra bad if you had a USG at the remote side using the internal IP as the inform address.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
