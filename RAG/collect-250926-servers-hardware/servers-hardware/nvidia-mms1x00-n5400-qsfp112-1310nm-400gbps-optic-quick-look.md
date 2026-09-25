---
id: collect-250926-servers-hardware/servers-hardware/nvidia-mms1x00-n5400-qsfp112-1310nm-400gbps-optic-quick-look
title: "nvidia-mms1x00-n5400-qsfp112-1310nm-400gbps-optic-quick-look"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "cost", "decode", "optics"]
source: docs/RAG/clean4/nvidia-mms1x00-n5400-qsfp112-1310nm-400gbps-optic-quick-look.md
source_anchor: ""
source_lines: [1, 37]
sha256: 5a8f3e05c8ac710bc041716fbeb5f8e1c60634efe471902f2529cbe5866ba3be
---

# nvidia-mms1x00-n5400-qsfp112-1310nm-400gbps-optic-quick-look

Today we are taking a quick look at the NVIDIA MMS1X00-N5400 QSFP112 1310NM 400Gbps modules. We bought these to connect two NVIDIA GB300 AI Station systems to an NVIDIA Spectrum-X switch. Since we had them, we figured we would at least show you what they are and why they matter. Folks liked seeing the massive NVIDIA 800G OSFP to 2x 400G QSFP112 DAC cable so we figured why not show these.

## NVIDIA MMS1X00-N5400 QSFP112 1310NM 400Gbps Optic Hardware

The module is a QSFP112 form factor. We went into the QSFP Versus QSFP-DD Key Differences if you want to learn a bit more about this module, versus some of the ones we have looked at previously.

A super quick decode here: QSFP means we have a quad-channel module, and the 112 tells us we have 112G PAM4 signaling.

That means we get a standard pluggable form factor that is more compact than the OSFP form factor we see often at 800Gbps and 1.6Tbps these days.

The connector is an MPO-12. As you may have surmised at this point, this is a 4x 112G design, so we are only using eight of the twelve fibers.

Something else neat is that NVIDIA adds small heatsinks to these modules outside the pluggable cage. Often 400Gbps ports have cooling internally on the device’s cages.

Here though, NVIDIA is using a small heatsink on the outside of the device so they stick out a bit from the NIC and switch.

The connector here is important since it is QSFP112, not QSFP56-DD. If you saw our recent Cheap Desktop 400GbE Switch MikroTik CRS804-4DDQ-hRM Review, that uses lower-cost PAM4 56Gbps lanes, so it needs eight lanes to achieve 400Gbps. As a result, QSFP56-DD is a different connector from QSFP112 because it needs eight channels instead of four back to the device.

We tend to run a lot of DR4/ LR4 optics, so the advantage here is that, unlike using a QSFP56-DD module, we do not need a gearbox in the module to connect a 400Gbps 500m DR4 link from our switch.

Next, let us plug it in to check the speed.

## Connecting at 400GbE Speeds NVIDIA ConnectX-8

We knew the answer to this, but here is an NVIDIA GB300 AI Station that we will review soon, connected to a high-end switch at 400Gbps speeds.

This autonegotiated almost instantaneously on the NVIDIA ConnectX-8 NIC that is part of that platform.

We can see here that it is reporting as a 10W maximum device.

We also saw this running in the systems in the mid 40C range, which is far below the threshold.

## Final Words

We have many DR4 optics for different devices in our lab, so getting these modules, despite the fact that they cost more than a DAC was worthwhile. We can swap switches, or do a direct GB300 AI Station to GB300 AI Station connection using the ConnectX-8 NICs and a 12-fiber MPO/MTP-12 singlemode cable. The other reason we purchased these is that, in large AI clusters, NVIDIA optics tend to run well at higher temperatures on the hot-aisle side than some popular generic modules. Since there is not a ton of active cooling for optical modules in the liquid-cooled GB300 AI Station machines, we thought it would be wise to just get ones we would not worry about.

Hopefully you enjoyed this look at these modules. They were a big investment, but having these to connect to our lab’s fiber infrastructure made life much easier from an integration standpoint, so we figured we would at least show them to you.
