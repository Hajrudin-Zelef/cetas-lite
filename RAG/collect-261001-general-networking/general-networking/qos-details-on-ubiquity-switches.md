---
id: collect-261001-general-networking/general-networking/qos-details-on-ubiquity-switches
title: "qos-details-on-ubiquity-switches"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["licenses"]
source: docs/RAG/collect-261001-general-networking/qos-details-on-ubiquity-switches.md
source_anchor: ""
source_lines: [1, 13]
sha256: b7c623cc3499970e32d1760d12e1dca542bcecada4a877429a16019806d98bf8
---

# qos-details-on-ubiquity-switches

Yes those Netgear are expensive for the small switches, doesn’t make sense to buy less than the one I have at work. The other important thing for the Cisco is that if you need help from Audinate, they are familiar with the SG300/350 switches.

I’d give some of the Mikrotik switches a look, I’m having good luck with my CRS3091g-8s+in (using both as a Top of Rack right now) It’s a 10gbps  sfp+ switch, but you can use 1gbps modules and change the settings so that most modules work. But no POE, which means the little AVio devices are going to be a problem. I’ve been thinking about this switch for home MikroTik Routers and Wireless - Products: CRS328-24P-4S+RM

Here is the Serve the Home article and video MikroTik CRS328-24P-4S+RM Review 24-port PoE and 4x 10GbE Switch

If you only need a few ports for moving audio, it should be fine. If you are moving 64 channels each port and a bunch of multicast, it might still be OK and possibly worth a try. Support has been helpful when I contacted them about one of my 309’s and the feature set seems to be very much everything you might want. Processor and buffer RAM might be the only things I’d like to have more from. I wouldn’t run too many other features, like not running Suricata or other scanners on your audio network, just use it as a plain old switch (probably in Switch OS). There is a good GUI in both the Sw OS and Router OS, and you can freely boot between the two.

And all of that said, while I am a Dante Level 3, we are still just getting started on our big Dante roll out. Been waiting for over a year for parts on our other switches, several buildings on campus will be on my network so we can move Dante and NDI from place to place for recordings or live events (we teach audio and video production). The day to day troubleshooting of Dante problems is still just theory to me, same for NDI (though I haven’t taken are certifications for the NDI stuff yet, been going through the Newtek courses).

What is your workflow for this switch? Might be worth sitting down to calculate what you really need per port. And add in things like, are you working directly to a storage server and does that server have a 10gbps port. How many devices do you use or is it mostly a big stage box like a Millenia Preamp box. Any other network connected things like the mixer, etc.

We are using Midas M32 mixers, a couple of Millenia preamp stage boxes, we have a bunch of AVio but none have been used yet, and a few other things that mostly haven’t been used yet, and a couple of DVS licenses for computers.
