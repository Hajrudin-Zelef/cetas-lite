---
id: collect-240926-tomshardware/tomshardware/intel-s-proposed-orbital-data-centers-would-manage-thousands-of-simple-leo-satel
title: "intel-s-proposed-orbital-data-centers-would-manage-thousands-of-simple-leo-satel"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "AWS", "Google", "Intel", "Nvidia", "SpaceX", "United States"]
dates: []
keywords: ["intel", "amd", "compute", "cost", "energy", "gpus", "latency", "nvidia", "robotics"]
source: docs/RAG/clean_en/tomshardware/intel-s-proposed-orbital-data-centers-would-manage-thousands-of-simple-leo-satel.md
source_anchor: ""
source_lines: [1, 53]
sha256: 58e9ed7cdd7c7eb9e82ecbbebf65614710cf183ae35deb70c73cf8018aae351e
---

# intel-s-proposed-orbital-data-centers-would-manage-thousands-of-simple-leo-satel

<!-- source: https://www.tomshardware.com/tech-industry/space/intels-proposed-orbital-data-centers-would-manage-thousands-of-simple-leo-satellites-two-tier-network-puts-the-brains-of-satellite-constellations-in-higher-orbit -->

An Intel patent application published on August 6, spotted by __Patentlyze__, describes an orbital data center architecture that moves some of the computing used to operate massive satellite constellations off the ground and into space. The architecture proposes a two-tier satellite network in which a small number of more powerful satellites in higher orbits manage large constellations of relatively simple satellites in low-Earth orbit, handling much of the computing and constellation coordination normally performed by data centers and network operations centers on the ground.

The application, US 2026/0230175 A1, is a continuation of an earlier Intel filing that was granted as US 12,542,604 B2 in February.

Intel’s proposed architecture is a different proposition from the orbital AI data centers now being pursued by companies such as SpaceX and Google, which aim to move AI compute itself into low-Earth orbit. SpaceX’s planned AI1 satellite and Google’s Project Suncatcher both envision running large-scale computing workloads in space, with the resulting data beamed back to Earth over high-bandwidth optical links. Intel’s orbital data centers, on the other hand, are designed primarily to serve the satellite network itself, acting as higher-orbit compute and control hubs for the much larger constellations operating below them.

In large LEO constellations such as Starlink, Telesat Lightspeed, and Amazon’s Project Kuiper, thousands of satellites are constantly moving relative to one another and the Earth. While the satellites perform their individual tasks, the network itself still has to determine how traffic is routed, which satellites and links should communicate, how spectrum is allocated, and how the constellation responds to failures, interference, weather, and other changing conditions. Much of that network planning and control processing is traditionally handled by computers on the ground, with routing and operational instructions calculated at terrestrial network operations centers and then transmitted back up to the satellites.

Intel says that this constant dependency on terrestrial infrastructure delays time-sensitive decisions, increases reliance on ground stations, and makes management harder as constellations grow into the thousands of satellites. Intel’s solution is to move part of that control and compute layer into orbit. Its architecture places more powerful satellites — which contain much more compute and storage capability than the individual LEO satellites — in Medium Earth Orbit (MEO), Geosynchronous Earth Orbit (GEO), or highly elliptical orbits, where they can maintain a broader and more persistent view of the LEO constellation below and take over tasks such as routing, mission planning, scheduling and network coordination without continually sending those workloads back to Earth.

The proposed setup does not eliminate the need for ground stations. It just keeps satellite network control processing in space. Intel specifically describes moving mission planning and scheduling operations into orbit. The company also argues that offloading heavier network-management tasks to a smaller number of powerful satellites could allow operators to build simpler, cheaper LEO spacecraft. Under current architectures, individual satellites still have to actively participate in network-control functions, requiring additional onboard compute and communications hardware. There’s currently no indication that Intel is actively building the satellites.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

Etiido Uko is a news contributor for Tom's Hardware covering the latest updates in big tech and the PC industry. He is a mechanical engineer and senior technical writer with over nine years of experience in documentation and reporting. He is deeply passionate about all things engineering and technology, and is an expert in gadgets, manufacturing, robotics, automotive, and aerospace.

- 
The patents have some merit but they can't really restrict the natural evolution of satellite placement or the natural progression of trying to make low latency or orbital efficiency of data centers better. Companies like SpaceX would likely develop something similar without running afoul of these patents if the logistics of lower and higher orbit synergy made sense. Ultimately a lot of this is just the likely evolution of orbital data center progression especially once we connect AI to space station and moon base operations. I'm not sure how much value it has for terrestrial AI data center operations over low earth orbit networks that likely interconnect with laser speed systems anyway. Adding more latency with higher orbit operations has some trade offs when coordinating back on Earth.Reply
- 
Reply
 If you mean this one, it's not even started construction yet, and also it's going to use/be an intel FAB.re3eyul said:the 3l0n will be there before Intel figures out how to make Leo , already got a factory doing Starship payload sized orbital compute modules .
https://www.tomshardware.com/tech-industry/semiconductors/terafab-starts-to-take-shape-100-million-square-feet-of-manufacturing-space-and-usd16-8b-initial-capital-investment
- 
A patent does not a "proposed orbital data center" make.Reply
 
The economics of the whole idea are simply ludicrous. Datacenters are big, heavy, power hungry machines that need cooling, regular maintenance, and upgrades to remain competitive. Literally none of that gets easier or cheaper in space.
- 
Reply
 Even though Nvidia and AMD want to release new accelerators on a near yearly cadence (these roadmaps are slipping in practice), I think the rate of improvement could be slowing down by a lot. The move to support lower precision formats was like a free lunch that couldn't last forever. Yearly upgrades won't make sense for most users. Big users that are buying whatever they can will gradually phase out the old as they buy the new ones.chaos215bar2 said:Datacenters are big, heavy, power hungry machines that need cooling, regular maintenance, and upgrades to remain competitive. Literally none of that gets easier or cheaper in space.
 
 You can't do any maintenance, so you have to make it robust and failure tolerant in the first place. It will last a few years before being deorbited to burn up, talk about planned obsolescence.
 
 Cooling in space is simply balancing an equation. It has to radiate away the energy that it uses, around 100-150 kW. So the satellites have to be relatively large, and fit in a payload bay using an unfolding design. It's a solvable engineering problem of a scale already accomplished by the International Space Station. They will put as many GPUs/accelerators on there as can be supported.
 
 If Starship becomes operational in a fully reusable configuration, it will be able to lift >100 metric tons to orbit at a relatively low cost, much faster than datacenters can be constructed on the ground. The supply chain to build a lot of these and lower costs is already there, since the satellites will be based on Starlink. The sun-synchronous orbits targeted will power the satellites nearly continuously.
 
I won't claim the economics of this make sense, because it's entirely possible the numbers won't add up, or Starship will experience another string of failures and not be ready in time, or the AI bubble will pop and the AI compute demand won't materialize. I don't think it's impossible for this to make sense, but it's obviously risky. On the Starship front, they have gotten the ship suborbital (by choice), and released Starlink v3 payloads. The booster is experiencing a lot of engine failures. Both need to be reliable and reflown regularly to lower the launch costs.
- 
Reply
Right. The distance from LEO to earth is much less than the distance from LEO to MEO much less to GEO, so I don’t understand why communicating with terrestrial is so horrible, at least as explained by this article.Zaranthos said:Adding more latency with higher orbit operations has some trade offs when coordinating back on Earth.
- 
And what happens 5 years later when we realize there's a hardware flaw and all of them can be hacked? If not that, the hardware will be outdated before they even get them all in orbit. How long before there's so much space junk up there? How long before burnt up rocket fuel becomes out next environmental disaster. It's all just so ugh.Reply
- 
Reply
 LEO, with minimal re-boosting capabilities. They will come down in a few years when they run out of fuel. Something like 1600 Starlink have already come down.Trake_17 said:And what happens 5 years later when we realize there's a hardware flaw and all of them can be hacked? If not that, the hardware will be outdated before they even get them all in orbit. How long before there's so much space junk up there? How long before burnt up rocket fuel becomes out next environmental disaster. It's all just so ugh.
 
 These are basically still inside the outermost atmosphere. The ISS has been getting re-fueled and re-boosted every few years.
 
I still don't think it is a good idea, though. I think the engineering solutions on the ground are easier to deal with. I would like to see long term benefits of something like a cheaper closed loop cooling solution for datacenters and sensible power grid expansion. Just slapping massive generators next to neighborhoods and zoos seems counter productive.
