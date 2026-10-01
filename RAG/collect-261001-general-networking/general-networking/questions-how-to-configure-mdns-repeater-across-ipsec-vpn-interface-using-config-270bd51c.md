---
id: collect-261001-general-networking/general-networking/questions-how-to-configure-mdns-repeater-across-ipsec-vpn-interface-using-config-270bd51c
title: "questions-how-to-configure-mdns-repeater-across-ipsec-vpn-interface-using-config-270bd51c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-how-to-configure-mdns-repeater-across-ipsec-vpn-interface-using-config-270bd51c.md
source_anchor: ""
source_lines: [1, 27]
sha256: 6f57ea4381f800278268f4c6a1a42fd02a3bdd109c96c0931258dd1a5b659f99
---

# questions-how-to-configure-mdns-repeater-across-ipsec-vpn-interface-using-config-270bd51c

@UI-Team
I have software for a device (USB TCP/IP Server) that uses UDP multicast for discovery. The client software is on a machine at site A (on subnet 10.10.10.0) and the device is at site B (on subnet 10.20.10.0). I need to bridge the subnets so a Site A broadcast will be forwarded to Site B subnet. I read it is not advised to enable mDNS via GUI (as it also forwards broadcast packets to all networks to include the WAN side). So, I found a gateway.config.json statement that should just repeat specifically what I need. My question is does anyone know what "interface" I would put so mDNS repeater will go across the IPSEC VPN established between the USGs?
Here json for site A, I assume "eth1" as that is my untagged VLAN (10.10.10.) but what do I put as the site B interface? I did a show interface and found vti0 (which is the IPSEC VPN). So, assuming if I put that then it will repeat all broadcasts from Site A subnet (on eth1) to all of the Site B subnets (VTI0)?
{
	"service": {
		"mdns": {
			"reflector": "''",
			"repeater": {
				"interface": [
					"eth1",
					"vti0"
				]
			}
		}
	}
}
Anyone see any issues with what I'm trying to accomplish?
I haven't tried this before, and I'm not sure it will actually work across an IPSEC tunnel, but it looks like you're on the right track here. The mDNS repeater is what you'd want, and I think that vti0 interface would be correct.
You would also need to add essentially the same config on the site B side to repeat the mDNS traffic again from the IPSEC "network" to the local site VLAN there.
You may need to do some tcpdump's from the USG interfaces to verify the traffic is actually being repeated to the IPSEC interfaces as expected.
@virtualfreak have you got this working?
I'm trying to accomplish the same but no luck so far. Repeater setup on both sides, but the traffic doesn't get repeated for VTI interface.
I have it working without passing mDNS across the site-site link - just added services at SiteA I wanted visible at SiteB to the mDNS servers at SiteB and it all works fine. Details here: https://community.ui.com/questions/USG-VPN-and-mDNS/71b8da66-9d0c-4463-b46b-46d3b32cb37e#answer/d2a29628-e7fc-4248-9fcc-13e3c24a0fbb
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
