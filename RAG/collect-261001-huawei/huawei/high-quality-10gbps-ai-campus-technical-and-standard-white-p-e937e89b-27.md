---
id: collect-261001-huawei/huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b-27
title: "high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["energy", "latency"]
source: docs/RAG/collect-261001-huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b.md
source_anchor: ""
source_lines: [991, 1055]
sha256: 674430b47b4b921b1cb7a6a80f19fe2b53aacba9d4506fccba2ad2cbc2abb260
---

# high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b

                   Resolution of audio and video services                                                       2.6.2.1 Device-Level Energy Saving                         2.6.2.2 Network-Level Energy Saving
                   without downgrading or video freezing                4K                    720p
                                                                                                                In device-level energy saving, energy saving modes         N e t w o r k- l e v e l d y n a m i c e n e r g y s a v i n g i s
                            upon link congestion
                                                                                                                are defined on single devices based on different           implemented on demand through intelligent SDN
  Application    Time taken to insert a 10 MB image in the                                                      loads. For example, a device enters the sleep/             measures. Specifically, shutdown can be performed
                                                                        6s                     20s
   assurance            collaborative office scenario                                                           standby mode when it is unloaded, and enters               by group at a scheduled time and by role for each
                 Performance for 15-channel video playing                                                       the low power consumption mode when lightly                region, so that the network intelligently shuts
                                                                 1080p, latency <     720p, latency < 1000      loaded.                                                    down some devices by time segment and device
                  and 15-channel PPT slide switching on
                                                                     100 ms                    ms
                             cloud desktops                                                                                                                                role to achieve network-level energy saving.
                                                                                                                APs are the most numerous nodes on a WLAN,
                                                                  M-LAG, 50 ms
                               Hitless upgrade                                           Not supported          and their working status becomes the focus of an           AI analyzes network tidal characteristics and
                                                                service switchover
                                                                                                                energy-saving solution. APs are usually powered            configures energy-saving policies for each region.
                                                                 10 ms to 20 ms                                 by PoE switches. In addition to entering the               The following figure shows the network traffic
   Industrial                                                                         > 50 ms switchover
                 High-performance industrial ring network       switchover upon a                               sleep mode or low power consumption mode,                  statistics of a teaching building in a university over
   assurance                                                                             upon a fault
                                                                      fault
                                                                                                                APs can achieve energy saving through PoE OFF              five days. It can be seen that the network usage
                     Dual fed and selective receiving for                                                       operations.                                                increases significantly at 07:00, reaches the peak
                                                                    Supported            Not supported
                                redundancy                                                                                                                                 in the morning and afternoon, and drops to near
                                                                                                                . PoE OFF mode: The PoE switch disables the port           zero after 00:00.
                   Latency when the air interface channel                                                         from supplying power to the AP. In this case, the
                                                                     < 50 ms                < 200 ms
                       utilization is greater than 80%                                                            AP does not consume any power. When
                                                                                                                  the AP is started, the port provides power to the
                    Less than –68 dBm downlink signal
      User                                                                                                        AP again.
                     strength of a STA (or about 10 m               Bandwidth            No bandwidth
   assurance                                                                                                    . Sleep mode: Key hardware such as the CPU runs
                 horizontal distance between a STA and an       increased by 30%           increase
                           AP, without blocking)                                                                  at extremely low power, and other components
                                                                                                                  are shut down. The sleep mode saves 80% to
                            Fault alarm function                    Supported            Not supported
                                                                                                                  90% energy for a single AP, and the AP can
                                                                                                                  resume operation within one minute.




2.6 Energy-Saving

 2.6.1          Definition

Campus network energy saving includes multiple           collaborates with the building management system
dimensions, including device-level energy saving,        to significantly reduce the energy consumption
network-level energy saving, and smart building          of facilities such as air conditioners and lighting,
energy saving. Smart building energy saving uses         thereby reducing energy waste in campuses and
the CSI sensing capability of WLANs to monitor           achieving the goal of green buildings.
and identify environment changes. CSI sensing


                                                                                                                                Figure 2-13 Network traffic statistics of a teaching building in a university




                                                    34                                                                                                                35

