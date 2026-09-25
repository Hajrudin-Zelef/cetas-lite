---
id: collect-250926-servers-hardware/servers-hardware/broadcom-shows-the-pex90144-144-lane-pcie-gen6-switch-at-sc25
title: "broadcom-shows-the-pex90144-144-lane-pcie-gen6-switch-at-sc25"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Broadcom", "Nvidia", "OpenAI"]
dates: []
keywords: ["accelerator", "blackwell", "chatgpt", "gpus", "nvidia"]
source: docs/RAG/clean4/broadcom-shows-the-pex90144-144-lane-pcie-gen6-switch-at-sc25.md
source_anchor: ""
source_lines: [1, 35]
sha256: 1b4bef00d3aa1285cd573e3ef01ec6dfd035bef570af5f15d5624569bde0cd0e
---

# broadcom-shows-the-pex90144-144-lane-pcie-gen6-switch-at-sc25

At SC25, Broadcom showed off its PCIe Gen6 switch chips for the 2026 generation of servers. Not only was the PEX90144, a big 144 lane/ 72 port PCIe Gen6 on display, but it was also integrated into a Dell PowerEdge XE7745 style PCIe board.

## Broadcom Shows the PEX90144 144 lane PCIe Gen6 Switch at SC25

This was one easy to miss and walk by at SC25, but Broadcom had a PEX90144 switch platform on display. You can see that this is actually a Dell PowerEdge XE7745 platform outfitted with a red board that usually indicates a pre-release board, although PCIe Gen6 would be more focused on the next-generation of NVIDIA Blackwell Ultra GPUs and ConnectX-8 NICs.

We can see inside the chassis so we can focus on the PCIe switch board. Here we can see two PCIe switches. 144 lanes is “only” 9x PCIe x16 slots, and this chassis has eight accelerators.

Beyond the accelerator slots (the PowerEdge XE7745 is a bit of a funky design since they are all internal) the accelerators also need NICs, often in a 1:1 or 1:2 ratio of NICs to GPUs. The switches also need a x16 link to the CPU, so 144 lanes goes very quickly.

Here are the rear NIC I/O slots as another example of all the I/O required in even CEM slot AI systems.

Broadcom’s Avago and PLX switch chip legacy has become the dominant force in PCIe switches today. Broadcom also has the PLX91144 (144 lane PCIe Gen7) on tap for 2027 and the PLX 92144 (144 lane PCIe Gen8) set for 2029.

The roadmap here is strong.

## Final Words

Broadcom is pushing back with an aggressive roadmap just as competitors such as Astera Labs, XConn, and others are pushing into the space as Microchip’s line of switches have become far less seen in servers. This is important for future generations of AI and storage servers (as well as CXL switches, which come from this line as well.) It is important that Broadcom executes on this roadmap because NVIDIA is pushing Broadcom out of both high-speed NIC and PCIe switch opportunities in NVIDIA B300 servers, as we discussed in our Substack.

Nothing of that is of any use for mere mortals.

Even for the stuff that we could technically put to great use, prices are beyond insane.

Microchip’s PCIe5 switch chips costs MORE THAN A FRIGGIN EPYC THAT CONTAINS _BIGGER_PCIE5 SWITCH.

WTF?!?

@Not_Really_Me,

You might not be able imagine it yet, but there are AI workloads that could crush that system, regardless of the accelerators you put into it. ;)

Understanding how ChatGPT works under the hood (all the different layers between your chat conversation and the frozen weights of the neutral network) will probably help, ask it about that, and it’ll blow your mind on how insanely leveraged these systems can get. :D

Drkrieger, this is servetheHOME, not servetheBILLIONDoLLArComPAnY
