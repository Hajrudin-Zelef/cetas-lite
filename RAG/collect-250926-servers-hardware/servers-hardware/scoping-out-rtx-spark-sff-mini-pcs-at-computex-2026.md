---
id: collect-250926-servers-hardware/servers-hardware/scoping-out-rtx-spark-sff-mini-pcs-at-computex-2026
title: "scoping-out-rtx-spark-sff-mini-pcs-at-computex-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Apple", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "blackwell", "ethernet", "gpu", "lpddr5x", "memory", "nvidia", "pricing"]
source: docs/RAG/clean4/scoping-out-rtx-spark-sff-mini-pcs-at-computex-2026.md
source_anchor: ""
source_lines: [1, 47]
sha256: 4c40f01097473c84323ad1840e29587077ffd91ece936a7d2aa233888f69a542
---

# scoping-out-rtx-spark-sff-mini-pcs-at-computex-2026

Although the first systems based around NVIDIA’s new(ish) RTX Spark Arm SoC will not be shipping until the fall, several of the company’s OEM partners already had their hardware on display at this week’s Computex trade show. This includes the 14-to-16-inch laptops for mobile users, as well as some of the small form factor desktop systems that are slated to be released as well. While crawling the NVIDIA venue at the show, our own Patrick Kennedy managed to catch up with a few of the SFF mini-PCs and catch some candid photos of them.

To note, at this point none of NVIDIA’s partners are announcing detailed specifications for their respective systems. Keeping things in line with NVIDIA’s high-level overview from the RTX Spark/N1X reveal at the start of the week, we know that the high-end chip configuration will feature 20 CPU cores, a 48 SM Blackwell integrated GPU, and up to 128GB of LPDDR5X memory. But NVIDIA has not confirmed any details beyond this. The running agreement among the technical press is that this is almost certainly a rebadged version of NVIDIA’s GB10 chip, which has been shipping in systems from NVIDIA and several OEM partners for nearly a year now.

With that said, the press releases for these systems do offer a bit more detail on the I/O options, confirming a few details that were only previously assumed. On the networking side of matters, all of the systems with disclosed specs are listed as offering 10Gb Ethernet, while ASUS’s announcement also notes Wi-Fi 7 + Bluetooth 5.4 for wireless connectivity. Meanwhile the USB-C ports are running at 20Gbps. And finally, storage is backed by a PCIe Gen5 x4 M.2 slot. If this sounds familiar, these are identical to the specs of the NVIDIA DGX Spark and GB10 clone boxes, underscoring the extensive similarities between the GB10 and N1X/RTX Spark hardware families.

By and large, the biggest difference so far between these systems and vendors’ existing GB10 systems is the lack of a ConnectX-7 NIC and QSFP ports, which are a mainstay on GB10 systems. This high-speed networking option was one of the major features of the DGX Spark and its clones, as it allowed for multiple boxes to be linked together to test multi-node scaling, or just to have a very fast link to other systems and storage. But such a powerful NIC is not a cheap thing to add, and Windows considerations aside, NVIDIA’s partners are going to be competing with the likes of the Mac Studio, AMD Ryzen AI Halo, and numerous other dev-focused mini-PCs that don’t have the ConnectX tax. So along with support for Microsoft’s OSes, the lack of a ConnectX NIC would seem to be the other defining feature of RTX Spark mini-PCs.

None of the companies involved have announced pricing, configuration, or an estimated launch window for their systems. Though with all of the systems aimed at AI development and content creation, these are clearly intended to be positioned as premium systems. So expect them to be priced accordingly.

Otherwise, given the outright similarities between all four systems, it would not be surprising to find out that NVIDIA is keeping a heavy hand on the steering wheel here and enforcing certain design requirements, similar to what they have done for GB10-based systems.

### ASUS ProArt GA10 Mini PC

First up we have ASUS’s ProArt GA10 Mini PC. Surprisingly, ASUS has published the dimensions of the system, confirming that it is 150 x 150 x 51mm. Which coincidentally enough, are the dimensions for their GB10 box, the Ascent GX10.

The GA10 is not a doppelganger for the GX10, but it is certainly quite close. Right on down to the iconography and Kensington lock slot. ASUS is offering four USB-C ports (all on the rear of the system0 along with HDMI out and a 10Gb Ethernet jack. They have also published a bit of info about the thermal design, stating that the system is designed to cool up to 140 Watts.

According to the company they are pitching the machine at AI developers as well as creators, essentially making it the Windows counterpart to the GX10.

### Dell XPS RTX Spark Desktop

Next up, we have Dell’s RTX Spark system, which is being released under the XPS brand as the XPS RTX Spark Desktop. This is the RTX Spark counterpart to Dell’s existing Pro Max with GB10.

Notably, the RTX Spark desktop is taller than Dell’s GB10 system. The company has not published official specifications, but based on photos alone it is clear that it stands taller than the NVIDIA-regulated GB10 system.

Dell’s desktop offers four USB-C ports long with HDMI out and a 10Gb Ethernet jack.

### Lenovo SFF RTX Spark

Third up, we have Lenovo’s descriptively named SFF RTX Spark system. This is the RTX counterpart to the company’s ThinkStation PGX, which, despite the similarities, Lenovo seems to have built a new chassis for.

Lenovo’s desktop offers four USB-C ports long with HDMI out and a 10Gb Ethernet jack.

### MSI EdgeMesa N AI+

Last but not least, we have MSI’s EdgeMesa N AI+. The while box certainly stands out compared to its largely gray-and-black competitors. This system will be joining the MSI EdgeXpert as the company’s two NVIDIA-based SFF AI systems.

As with every other RTX Spark mini-PC showcased at Computex, MSI’s desktop offers four USB-C ports long with HDMI out and a 10Gb Ethernet jack.

Overall this system is not quite a clone of the EdgeXpert, though coloring aside it is not a massive departure, either. MSI did publish a press release for this system, outlining that they are pitching the system at AI developers, data scientists, and creators.

It’s hard to see how these are any different from the DGX Spark, which is awesome by the way with its 128gb unified memory. It’s not a rocket ship, but it’s not slow and runs huge LLMs relatively quickly. I guess the only difference is the lack of the connect port, which I know adds a lot of money to the DGX Spark, but seeing that these may be offered with as little as 16gb unified memory boggles my mind. I know RAM prices are nuts, but that’s because you need a lot of it to do the “cool” stuff.

But does the GB10/N1X have enough memory bandwidth and cpu power to do the ‘cool’ stuff? I have a feeling a lot of these unified memory architectures are kind of poorly executed compared to Apple’s M5 Max (followed by the M5 Pro).

So is it just a DGX Spark without the Connectx NIC and possibly with optionally less memory?
