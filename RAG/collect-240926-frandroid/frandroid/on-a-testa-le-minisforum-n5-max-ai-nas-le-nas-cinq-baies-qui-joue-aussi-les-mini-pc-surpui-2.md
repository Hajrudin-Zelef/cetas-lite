---
id: collect-240926-frandroid/frandroid/on-a-testa-le-minisforum-n5-max-ai-nas-le-nas-cinq-baies-qui-joue-aussi-les-mini-pc-surpui-2
title: "on-a-testa-le-minisforum-n5-max-ai-nas-le-nas-cinq-baies-qui-joue-aussi-les-mini-pc-surpuissants"
domain: frandroid
role: reference
task: reference
actors: ["AMD", "China", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "benchmark", "compute", "gpu", "license", "nvidia", "throughput"]
source: docs/RAG/clean_en/frandroid/on-a-testa-le-minisforum-n5-max-ai-nas-le-nas-cinq-baies-qui-joue-aussi-les-mini-pc-surpuissants.md
source_anchor: ""
source_lines: [61, 106]
sha256: 1a52967efa9ca80f84eed91869023b3779c0dd0f0cfd894ee0141a84d91f03f7
---

# on-a-testa-le-minisforum-n5-max-ai-nas-le-nas-cinq-baies-qui-joue-aussi-les-mini-pc-surpuissants

Miniscloud sometimes lacks responsiveness and even though Minisforum keeps releasing updates, it still suffers from quite a few crippling bugs. More annoying still, the services and features are limited, despite the presence, for example, of a Docker manager precisely intended to open things up to many possibilities not provided by default. These are only a few examples, but the hardware monitoring portion of the PC/NAS is fairly limited, some options still appear in Chinese, and the interface is not always very "clean."

That said, Minisforum is aware of Miniscloud's limitations and it allows, even encourages, setting up your own OS. To do this, in the manner of certain manufacturers like Terramaster, a USB-A port has been placed internally to simplify the deployment of a live-OS on a USB key.

Of course, it is also possible to access the BIOS of the N5 MAX AI NAS to completely revise the boot order and decide to install the OS of your choice. Minisforum does things rather well since all Windows drivers are available for download from its website.

Better still, since we are talking about an AMD platform (CPU, GPU, NPU, and chipset), many drivers exist, whether we are talking about Windows or Linux. One can therefore very well create one's own configuration or rely on all-in-one solutions such as OpenMediaVault, Rockstor, TrueNAS, or Unraid. We were impressed by the ease with which we were able to deploy this kind of solution and, above all, by the efficiency with which the whole thing then worked.

For our more in-depth tests, however, we favored installing Windows 11 in order to fall back on something more classic in terms of a mini-PC. As we said, the software side poses no problem since all drivers are offered by Minisforum, but also because the AMD platform is an old acquaintance. As on any PC — which this N5 MAX AI NAS ultimately is — it takes only a few dozen minutes to have a Windows 11, perfectly functional... but without a license.

## And in use, how does it run?

We tested the N5 MAX AI NAS in its two configurations: as a NAS and as a mini-PC.

### In NAS use

Although the N5 MAX AI NAS has the advantage of combining the functions of a NAS and a mini-PC within a single machine, we decided to divide our tests into two main categories. We then began with the purely NAS operation of the machine.

An operation that involves inserting several units into the available bays. Unsurprisingly, everything we expect from a 2026 NAS is here: caddies capable of accepting 2.5-inch units (screws provided as needed) or 3.5-inch units (screwless system), RAID capabilities up to RAID 5 and RAID 6 for protection in case of failure, and the choice between ZFS and ext4. No, BTRFS is not on the menu.

We won't go back into that here, but at present, the Miniscloud OS does not do justice to Minisforum's hardware. Admittedly, it contains all the features of a 2026 NAS, but in most cases it remains fairly basic: the multimedia tools (Albums, Movies, Music) are very poor, there is no video surveillance module, the system monitoring has a horrible ergonomics and the tiny application portal highlights these limitations.

Even so, it remains usable. Data sharing poses no problem and the power of the configuration allows several people to work on the NAS without difficulty. Below, we measured the NAS's performance via CrystalDiskMark and we were a little surprised by the results. Not that the throughput is bad, but we are quite far from saturating the 10 GbE controller.

To settle the matter, we carried out the same CrystalDiskMark test after changing the operating system and swapping Miniscloud for Windows 11. As you can see, the results are nothing alike, proof that our storage units (Kingston DC600M SSDs) are not to blame: the network controller loses itself a little along the way, without it being scandalous.

Before taking fuller advantage of the N5 MAX AI NAS under Windows 11, we continued our NAS tests for a little "routine" check. We deliberately broke the RAID 6 array that we had created to assess the time needed to rebuild it. Well, you should know that the Ryzen AI MAX+ 395 works wonders: quite simply, we had never seen a NAS rebuild our array so quickly, made up of 100 GB spread across eight large files and 10 GB spread across more than 4,000 small files. A record, that's all.

### In mini-PC use

The other side of the N5 MAX AI NAS is its use as a mini-PC. There, you can install quite a few things, but we settled for Windows 11. As we said, the installation procedure posed no problem and within a few tens of minutes, we found ourselves on the desktop of Microsoft's OS. A useful point to note right away: it is possible to manage up to four screens thanks to the three USB4 ports that complement the HDMI port.

There is actually not much more to say: the N5 MAX AI NAS is then a Windows PC like any other Windows PC, knowing that the hardware configuration makes things extremely comfortable with its minimum of 64 GB of RAM, this 16-core Zen 5 processor and its 40 RDNA 3.5 compute units. Rather than dwelling on the machine's use, we instead suggest you discover its performance.

On Cinebench 2026, first of all, only the CPU part is tested. We indeed left out the GPU test. The Zen 5 architecture of the Ryzen AI Max+ 395 allows for honest single-thread/single-core scores, but it is obviously in multi-threads that things become more interesting: with more than 7,100 points, the N5 MAX AI NAS is clearly ahead of the GMKtec EVO-T2S that we tested recently, but also ahead of the MS-S1 Max by Minisforum. The hardware and, above all, software optimizations are not unrelated to this fine performance.

To drive the point home, we conducted measurements on Blender benchmark. There, the Ryzen AI MAX+ 395 again performs very well, but note above all that performance is much higher once the GPU part of the AMD processor is brought into play. Generally speaking, Radeon chips do far less than GeForce ones on this test, but between a GPU and a CPU, there is still no comparison.

We finish these "general-purpose" measurements with PCMark, the synthetic tool par excellence. By simulating countless usage scenarios for a machine (office work, video conferencing, photo retouching, video editing, 3D modeling, etc.), it offers a broad overview of the machine's potential and, it must be said, it is rather good. Nearly 10,000 points overall and, above all, more than 16,000 on both the productivity test and the digital content creation test, that's unprecedented on a mini-PC without a dedicated graphics card!

The power and therefore the impact of the Radeon 8060S solution needed to be verified more precisely. To do this, we set things straight with 3DMark and our three scenes of choice: Fire Strike, Time Spy Extreme and Steel Nomad. Each time, the results are very interesting. We remain far from a real graphics card, but the scores are high and suggest fine results in the face of real video games, even demanding ones.

*Shadow of the Tomb Raider* first. It runs rather very well, even without enabling any "assistance" like FSR. In 1920 x 1080, with details at minimum, we exceed 200 frames per second. So, naturally, we set the details to maximum and there, it's better than any mini-PC ever tested at Frandroid: 138 frames per second and the assurance of never dropping below 100 fps.

We move on to something much tougher, *Cyberpunk 2077* by the Poles at CD Projekt RED. A technological showcase widely used by NVIDIA, the game now also runs very well on AMD hardware and the N5 MAX AI NAS is there to bear witness. We stayed in 1920 x 1080 and enabled FSR 3.1 with frame generation, but what a slap. With details on "low," we flirt with 260 fps and in "ultra" we are still at nearly 170 fps, so why not enable ray tracing? It drops, of course, but FSR works wonders: in low ray tracing we are at 150 fps and in ultra ray tracing, we come close to 90 fps. Impressive.

