---
id: collect-250926-servers-hardware/servers-hardware/supermicro-sys-110a-16c-rn10sp-1u-25gbe-intel-snow-ridge-server-review-1
title: "supermicro-sys-110a-16c-rn10sp-1u-25gbe-intel-snow-ridge-server-review"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "United States"]
dates: ["2021-07-10"]
keywords: ["intel", "cost", "distribution", "ethernet"]
source: docs/RAG/clean4/supermicro-sys-110a-16c-rn10sp-1u-25gbe-intel-snow-ridge-server-review.md
source_anchor: ""
source_lines: [1, 76]
sha256: 4dc074ebe6f9db7e58a8efd534bbe953dc65a74dfbaca58a8f603f684ea4d5c7
---

# supermicro-sys-110a-16c-rn10sp-1u-25gbe-intel-snow-ridge-server-review

Today we have one of the most fun systems we have tested recently. The Supermicro SYS-110A-16C-RN10SP is an Intel “Snow Ridge” platform designed with a lot of connectivity. Putting this into perspective, we get a 16-core processor, eight 1GbE ports, and two SFP28 25GbE ports built-in. This compact server somehow ended up resulting in a longer-than-we-expected review, indicating that a lot is going on in this short-depth server. Let us get to it.

## Supermicro SYS-110A-16C-RN10SP Hardware Overview

The Supermicro SYS-110A-16C-RN10SP is a short-depth 1U server. Indeed, it is only 9.8″ deep which means it will fit in most racks.

On the right side, we get the Supermicro logo and a sticker with the BMC MAC address as well as the randomized password.

The center of the server has fans, the power button, LED indicator, and a pull-out service tag.

On the left side, we get two 2.5″ SATA drive bays. As we are going to discuss later, we wish that there was a way these could be NVMe instead. At the same time, for a network appliance, SATA is usually fine and tends to use less power.

On the rear of the server, we have a lot going on.

First, we have redundant 300W 80Plus Gold level efficiency power supplies.

In the middle, we have standard I/O including the out-of-band IPMI management port, two USB 3 ports, and a VGA port. Then things get more exciting with eight 1GbE ports. Four are powered by the Intel i350-am4 onboard. The other four are powered by the Intel P5342 SoC and use a Marvell Alaska 88E1543 PHY. Supermicro’s specs will highlight this PHY as it is very well known in the industry, but as we showed in the motherboard review, the OS uses the Intel SoC NIC driver, not a Marvell driver.

There are also two ports labeled “SFP”. These are actually SFP28 ports for dual 25GbE networking built-in. This 25GbE is also part of the Intel Snow Ridge SoC.

On the right rear we see space for a single expansion card.

In the PCIe Gen3 x8 riser, we have a dual 10Gbase-T NIC.

With that, let us get inside and see how this server is built.

For storage use, how about install a PCIe 3.0 or higher SATA/SAS card in the riser slot, assuming it will fit. Use the onboard NMvE slots as the OS/cache drives.

If such a card has external connectors then you can connect to external storage chassis via appropriate cabling.

Init7 users in Switzerland will probably snap this up as it finally looks like a homelab type box that can route 25Gbs

Interesting they’re still using wires for all the power distribution. All the Dell servers I’ve seen for years now route power on the PCB itself, and the PSUs, fans and drive bays all slot into connectors directly on the motherboard so there’s no need to route any power cables through tight spaces.

Not that there’s anything wrong with using standard connectors, it just seems like it would require more manufacturing labour and thus increased cost.

I love reading about all the cool new hardware, but how much is this is this thing? I’m guessing in 2000 – 3000 dollar range – which waaay over a normal “serving the home” budget I think. I really think these high end products should be separated to another website or something, because so many or the articles (while cool) are not serving the home anymore.

@owen I belive Patrick has answered that question many times in the future, the serve the “home” part is now way in the past. I understand that thr name servethehome.com has a brand recognition thing attached to it and this is probably why they still use it, but I would have expected a rebrand of this website at this point in time to be honest. It would take some work but maybe they can refer to the old website name for a year or two while thr new rebranded name is getting traction.

The past. Not future:-)

Malvineous that’s why Supermicro has so many SKUs and Dell has so few. They’ve got one motherboard in many form factors with wired power inputs. They’ll even keep the same case for generations. FlexATX you can get C3000, P5000, D-1500, D-2100, and there’s older C2000 I’d think. Then you can use mITX in FlexATX cases too. It’s like 100 different motherboard SKUs that can be fit into the same chassis. If you’re selling these as an OEM appliance you can get the same look and make models with different perf levels.

Owen I’ve been reading STH for years. I don’t understand that comment at all. They do lots of inexpensive gear. Lots of very expensive gear. I’d wager most here have homelabs because we use servers at work. That work gear in 2-3 years becomes homelab gear.

This is also a machine that homes with high-end internet like Korev is mentioning can use. If you’re in silicon valley there’s people spending $11,000 USD for 1m^2 on homes not even in the nicer areas or downtown San Francisco. If they’ve got fiber at a few hundred a month because US broadband prices are very high, then a $2000 machine as a firewall is not even a year’s worth of broadband service.

I’d also understand if you’re coming from a place like Albania, Rawanda, Alabama, Argentina homes are cheap, but then they’re covering cheap gear too that’s more applicable. Homes aren’t just in those cheap places.

Owen the barebones cost is $1800 retail for this edge server. By no means is that high end. If you go and get a Cisco 4200 series firewall (only the 9000 series is above that by them) you are looking at $200,000 or more for one device. For a home this would be high end but that’s it.

@owen, I believe due to the nature of readers here, it serves home IT professionals. For many, this price suffices for homelab, small business and high end home clients. I will probably end up buying a D-1700 series myself….

For provisioning a homelab, it would be interesting what the minimum spend would be for a system that can effectively route 25 gigabit Ethernet. A further study on the effects of firewall rules and operating systems at that level of performance would also be useful.

The pictures and measured time to build a Python interpreter from source are also nice.

How does a Raspberry Pi with a 2.5 gigabit USB Ethernet dongle compare? What about those router style mini PCs with multiple built-in Ethernet ports? How does increasing the complexity of the firewall rules affect these cheaper solutions?

@Owen has a reasonable point, in that this is not what most would think of as a home server; however some of my friends do have racks in their basement, and for many the office has migrated to the home as well.

Personally I’m still satisfied with the HP MicroServer Gen 8 that I got for <£200 shipped, and upgraded to 12GB RAM and an i5-3470T once it was cheaper to do so. But I find this site useful for reviews of equipment I may lease in 3-6 years' time, or possibly get second-hand.

As for a new name, may I suggest ServeTheOffice?

You know, it’d be interesting to see a series on 10+ Gbps office/lab/home routing options with performance tests. What can you really push through a Snow Ridge CPU running pfsense? What about VyOS? TNSR? How do N100s/Xeon D/Ryzen 7xxx CPUs compare?

I’m seeing it just today and felt that missing opportunity on design, if they rotated 180° the board facing all the back ports at the front (only redundant power at the back) that could have been a pretty cool all around 1U rackable multirole appliance, and with a couple of 2.5″ U.2 compatible bays, oooofff.

Eric Olson: For provisioning a homelab, it would be interesting what the minimum spend would be for a system that can effectively route 25 gigabit Ethernet.

See https://michael.stapelberg.ch/posts/2021-07-10-linux-25gbit-internet-router-pc-build/ for a good example – there are other interesting 25Gb internet posts there too

I love these short-depth supermicro servers. I’ve got a stack of them.

