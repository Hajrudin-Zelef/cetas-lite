---
id: collect-250926-servers-hardware/servers-hardware/asus-pro-ws-w890e-sage-se-motherboard-review-2
title: "asus-pro-ws-w890e-sage-se-motherboard-review"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "ethernet", "intel"]
source: docs/RAG/clean4/asus-pro-ws-w890e-sage-se-motherboard-review.md
source_anchor: ""
source_lines: [62, 98]
sha256: 58619b5f7871ce39ed473a611bfc8f8b076faa4a8c08d7f73ea90478e05692ea
---

# asus-pro-ws-w890e-sage-se-motherboard-review

On the long side of the board are a standard 24-pin ATX power socket and a pair of PCIe 8-pin power sockets. The latter sockets are to help power the PCIe slots, particularly when multiple high-powered cards are installed.

Meanwhile, along the top side of the board are no fewer than four more power sockets for the CPU, which come in two different varieties. The white sockets are traditional 12V CPU power connectors, while the gray sockets are PCIe 8-pin connectors that provide additional 12V power to the CPU. Even when not overclocking, all of these 12V connectors need to be wired up in order to adequately power the system

And if you are overclocking, the SAGE SE also has special provisions to support a second PSU. In that case, the left-side power connectors are all attached to the first PSU, while the top power connectors are attached to a second PSU.

Moving on, taking a look at the rear I/O panel of the SAGE SE, we find a collection of ports consistent with a high-end motherboard with BMC support.

Front and center, the motherboard sports a pair of 40Gbps USB-C ports for high-speed I/O.

Notably, while the motherboard supports USB4’s fastest speed grade, it is not Thunderbolt 4-certified. Surprisingly, for all of the other Intel hardware on the board, ASUS went with ASMedia’s ASM4242 controller here rather than one of Intel’s Thunderbolt controllers.

Since the Xeon 600 platform does not have integrated graphics, ASUS has included a pair of miniDP ports for DisplayPort In connectivity so that the USB-C ports can carry DP video.

Six more 10Gbps USB-A ports are located along the bottom edge of the rear I/O panel, including one specifically earmarked for BIOS recovery.

Elsewhere, near the top of the motherboard, there is a trio of RJ45 ports for Ethernet connectivity. For general Ethernet networking, ASUS has installed Intel’s dual-port E610-XAT2 controller on the motherboard, which in turn drives the board’s dual 10GbE connections. Finally, the left-most Ethernet port is a 1GbE connection that hangs off of the board’s ASPEED AST2600 BMC.

For that matter, so is the VGA port on the far left. Certainly quaint for 2026, this is the sole video output option for the BMC.

Finally, in another nod to its enthusiast/overclocking functionality, ASUS has placed CMOS and BIOS buttons on the rear I/O panel as well. One button clears the CMOS, while the other one triggers a BIOS flashback to recover the BIOS if it has become corrupted.

Next, let us get to the block diagram.

“Non angeli, sed angli”, as Pope Gregory the Great didn’t say.

(In other words, a couple of times you call them Angelboost fans, not Angleboost.)

Unfortunately, there are virtually no Xeon 600s available for month. The motherboards are sitting on the shelves at german distributors like a lump of lead.

@James

Whoops! Thank you for pointing that out. It has been fixed.

Xeon processors across the board are just so expensive, at least until you get all the way down to Cascade Lake. The only thing expensive on the AMD side is DDR5 gen threadrippers. I’m not sure why this is the way it is but it’s been like that forever, even though AMD been dominating ever since EPYC Rome and the 3990X. Are there any server or workstation Xeons that are really worth it if your interest is a CPU-based workload performance?

Great review for anyone considering a serious workstation build. The detailed look at the motherboard’s expansion and platform capabilities makes it much easier to understand who this board is really designed for.
