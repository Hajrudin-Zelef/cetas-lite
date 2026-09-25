---
id: collect-250926-servers-hardware/servers-hardware/gigabyte-nvidia-vera-rubin-and-more-at-nvidia-gtc-2026
title: "gigabyte-nvidia-vera-rubin-and-more-at-nvidia-gtc-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "rubin", "agentic", "compute", "gpu", "liquid cooling", "nvlink", "optics", "vera rubin"]
source: docs/RAG/clean4/gigabyte-nvidia-vera-rubin-and-more-at-nvidia-gtc-2026.md
source_anchor: ""
source_lines: [1, 115]
sha256: 02fba6f63f281e9e0f9814a5c80b6fa3aba2ad192b26fa3f5b8f9f98caba15de
---

# gigabyte-nvidia-vera-rubin-and-more-at-nvidia-gtc-2026

At NVIDIA GTC 2026, we got to tour the Gigabyte booth and saw so many AI systems. We also saw the next-generation NVIDIA networking components, Vera CPUs, Rubin GPU platforms, and even the company’s NVIDIA GB300 AI Station, which is going to be the hot thing in local agentic AI soon.

We also have a quick 2-minute video on everything we saw at the booth if you want the quick overview.

## Gigabyte NVIDIA Vera Rubin NVL72 Rack

First, and like many companies, Gigabyte had an NVIDIA Vera Rubin NVL72 rack on display. Unlike other companies, Gigabyte had many of the components on display beyond just the compute tray.

If you have not seen them yet, the new Vera Rubin NVL72 compute tray is really different. The front faceplate shows components with a new optimized layout for liquid cooling and easier service.

In the middle, there are still the NVLink compute trays.

Here are the additional compute nodes at the bottom.

This is a great shot of something that is different. Here we get four ConnectX-9 ports and two E1.S SSDs. Unlike previous generations, the horizontal and more spread-out storage is easier to cool.

If it seems like there is a pattern with the networking cages and SSDs, there is and we will show that later because we found the internal component at the Gigabyte booth. Next, though, we wanted to get to the NVLink switch tray.

## NVIDIA NVLink Switch Tray

The NVIDIA NVLink switch tray is something that was on display in Gigabyte’s booth.

Immediately apparent is that the new switch tray is designed for liquid-cooling, and that involves just about everything in the chassis so that fans can be removed.

Here is the bottom part of the NVLink switch tray where the actual switches are found and the NVLink spine connectors are covered in an orange cap for the show.

While there are still a few internal cables, another very noticeable feature is the leak detection built into the system. This was something we did not see as much of in the B200 generation, but it seems like the integrated leak detection has been a major focus point.

Here is a look at the right angle quick disconnect fittings inside the chassis.

Next, let us get to some of the neat networking bits we saw in this rack, including a co-packaged optics switch, BlueField-4 DPU, and ConnectX-9 SuperNIC.

On the last system, what is that “USB” port, really? Displayport? But it looks too small. It’s clearly not USB-C.

@justsomeguy

The W775-V10-L01 Block Diagram does not mention the “USB” port.

https://www.gigabyte.com/Enterprise/Tower-Server/W775-V10-L01#Overview

But the Specifications page mentions 1 x Micro USB port (USB-to-UART) on the Rear I/O panel

https://www.gigabyte.com/Enterprise/Tower-Server/W775-V10-L01#Specifications

MiniDP

@justsomeguy

The W775-V10-L01 Block Diagram does not mention the “USB” port.

https://www.gigabyte.com/Enterprise/Tower-Server/W775-V10-L01#Overview

But the Specifications page mentions 1 x Micro USB port (USB-to-UART) on the Rear I/O panel

https://www.gigabyte.com/Enterprise/Tower-Server/W775-V10-L01#Specifications

Rear I/O

4 x USB 3.2 Gen2 ports (Type-A)

1 x Micro USB port (USB-to-UART)

1 x Mini-DP

2 x QSFP ports

1 x RJ45 port

1 x MLAN port

2 x Wi-Fi antenna ports

3 x Audio jacks (Audio in/Audio out/Mic)

@Jerry

justsomeguy was referring to the “USB” port, not the MiniDP.

Both are used by the BMC. “In addition, there is a BMC USB port and display output”.

@justsomeguy

The W775-V10-L01 Block Diagram does not mention the “USB” port.

https://www.gigabyte.com/Enterprise/Tower-Server/W775-V10-L01#Overview

But the Specifications page mentions 1 x Micro USB port (USB-to-UART) on the Rear I/O panel

https://www.gigabyte.com/Enterprise/Tower-Server/W775-V10-L01#Specifications

Rear I/O

4 x USB 3.2 Gen2 ports (Type-A)

1 x Micro USB port (USB-to-UART)

1 x Mini-DP

2 x QSFP ports

1 x RJ45 port

1 x MLAN port

2 x Wi-Fi antenna ports

3 x Audio jacks (Audio in/Audio out/Mic)

@Jerry

justsomeguy was referring to the “USB” port, not the MiniDP.

Both are used by the BMC. Text on Page 5: “In addition, there is a BMC USB port and display output”

The only thing I’m excited for is hot swappable cx9 cards and end of row cdus.
