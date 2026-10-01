---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-431206-what-does-the-ubiquiti-bullet-do-5ebb2efd
title: "questions-431206-what-does-the-ubiquiti-bullet-do-5ebb2efd"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "distribution", "ethernet", "memory", "research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-431206-what-does-the-ubiquiti-bullet-do-5ebb2efd.md
source_anchor: ""
source_lines: [1, 31]
sha256: 3ed1c247660fa619e29cb23322f81b8bcca92d82da7288213477c0d2ec7daf7e
---

# questions-431206-what-does-the-ubiquiti-bullet-do-5ebb2efd

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
4
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I was briefly shown the networking setup for a rural hospital campus. Some of the things I was told or shown are:
There is an 802.11b signal, broadcast over an omnidirectional antenna on a tower. This network is password-protected. 802.11b was chosen for the sake of longer range, even though its bandwidth is not as great as 802.11g, because the ISP bandwidth isn't huge anyway.
One individual building I was shown has a directional antenna, pointed at the previously mentioned omnidirectional one. Its coaxial cable is run through a device called the Ubiquiti Bullet.
Out of the Bullet comes an Ethernet cable, which goes to a consumer-grade router, which (at least in this case) has an open in-building WiFi network.
This kind of networking is way beyond my experience level, such that I don't even know the terminology. I'm curious, though.
What is the Bullet - what is the name for this type of device?
I see from its specs that it has a processor and memory. Is it able to join the wide-area encrypted WiFi network and pass an unencrypted connection to the consumer-grade router?
Why might this kind of setup be chosen? Is it that a laptop WiFi card can't produce a strong enough WiFi signal to join the campus-wide network?
Basically I'm just trying to understand how this setup works and what the pieces do.
Thanks to anyone who will humor the naive questions of a programmer who lacks networking chops.
What is the Bullet - what is the name for this type of device?
I call it a radio, but it's basically the non-antenna part of the equation. What makes it unique is that it's really small compared with a regular access point (the smoke detector style, the larger panels, etc.) that would have an integrated internal antenna or a "rubber duck" external antenna like you see on home routers. This small size is convenient for mounting discretely.
I see from its specs that it has a processor and memory. Is it able to join the wide-area encrypted WiFi network and pass an unencrypted connection to the consumer-grade router?
Why might this kind of setup be chosen? Is it that a laptop WiFi card can't produce a strong enough WiFi signal to join the campus-wide network?
Sort of. The antenna on the bullet is going to be a higher gain than a laptop's antenna (and directional), so this creates a more stable uplink (to the router) than many laptops roaming around (most likely inside the buildings where the signal would be weaker) with non-directional antennas trying to maintain a connection. They also likely wanted to partition the network by placing a router behind the bullet bridge. Typically you'll see multiple access points (radio + antenna) repeating the signal in a wireless distribution system (WDS) or mesh network.
A bullet is a wifi radio. It's just a different form factor to some of their all in one units.
You mount these on a rooftop or similar to provide a long range point to point link
Long range meaning several Kilometres.
It works like a client bridge which means you can wirelessly connect an entire network to another network.
They are powered over PoE and you need to connect them to an antenna.
In terms of encryption, yes they support the normal protocols wpa etc. They also support PPPoE authentication if you need that.
I had one of these paired with a patch antenna on my house at one point. It was what was providing my internet access. The ethernet cable came from from it down through the ceiling and was connected to the WAN side of my normal ethernet wifi router. The access point that it was paired with was 8km away.
The Bullet is a wireless radio. It can be used as an access point, a client, or a bridge. In functionality, it's not unlike a Linksys device that one would find in most homes, but is enterprise grade. The radio is more powerful than one you'd find in a home device so it can operate over a much longer range (several kilometers in a point-to-point configuration). It also doesn't have an antenna. You would connect it directly to the antenna of your choice depending upon your need. Since the bullet can attach directly to the antenna, you reduce the signal loss that might otherwise be incurred from using RF cable. You can then run the twisted pair cable from the back of the Bullet to wherever you need without worrying about loss (at least within the limits of the twisted pair cable).
Yes, the Bullet will be able to use encryption and pass an unencrypted connection to the Ethernet port. It can do all of that which you might be familiar with from using home AP devices, but with more options and some additional features. The manual for the OS has lots of screenshots if you are interested in what else it can do. There's also a datasheet that provides more information and pictures than the product page.
The reason someone might use this is to establish a wireless point-to-point connection between two sites, such as for the building in which you saw the device. Or there may be a need to place an antenna in a location far away from where the network equipment will reside, such as up on a tower.
