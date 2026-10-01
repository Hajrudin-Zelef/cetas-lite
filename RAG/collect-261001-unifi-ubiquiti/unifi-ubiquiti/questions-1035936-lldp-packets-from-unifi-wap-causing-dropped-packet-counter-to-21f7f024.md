---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1035936-lldp-packets-from-unifi-wap-causing-dropped-packet-counter-to-21f7f024
title: "questions-1035936-lldp-packets-from-unifi-wap-causing-dropped-packet-counter-to--21f7f024"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google", "Intel"]
dates: []
keywords: ["ethernet", "intel", "latency", "memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1035936-lldp-packets-from-unifi-wap-causing-dropped-packet-counter-to--21f7f024.md
source_anchor: ""
source_lines: [1, 113]
sha256: af16c7b7146f27f45fb76b0b457a71bc04584819e88789f25ea6f0601bcd56e4
---

# questions-1035936-lldp-packets-from-unifi-wap-causing-dropped-packet-counter-to--21f7f024

I have two UniFi UAP-AC-Pro WAPs on my network that are generating LLDP packets that are causing the received frame dropped counter to increase on my Proxmox server.
I've confirmed that it is lldpd on the WAPs that are generating the "bad" traffic, as it stops when I stop and disable the service on the WAP.
I've attempted to block LLDP via ebtables, to no success. I've also installed the LLDP tools on the Proxmox server to attempt to "catch" the traffic, and while that worked in the sense that I'm able to view LLDP stats on my proxmox server, the error counter is still incrementing.
Possibly of note: since this is a virtualization server, there are several standard Linux bridges attached to this interface.
Here's the output of various commands that I've ran when troubleshooting:
ifconfig:
enp6s0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        ether 68:05:ca:bd:3b:7d  txqueuelen 1000  (Ethernet)
        RX packets 469667  bytes 523243591 (499.0 MiB)
        RX errors 0  dropped 738  overruns 0  frame 0
        TX packets 129160  bytes 18630382 (17.7 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
        device interrupt 35  memory 0xf72c0000-f72e0000
ethertool -S
NIC statistics:
     rx_packets: 470347
     tx_packets: 129354
     rx_bytes: 523549801
     tx_bytes: 18673005
     rx_broadcast: 4137
     tx_broadcast: 846
     rx_multicast: 19087
     tx_multicast: 5599
     rx_errors: 0
     tx_errors: 0
     tx_dropped: 0
     multicast: 19087
     collisions: 0
     rx_length_errors: 0
     rx_over_errors: 0
     rx_crc_errors: 0
     rx_frame_errors: 0
     rx_no_buffer_count: 0
     rx_missed_errors: 0
     tx_aborted_errors: 0
     tx_carrier_errors: 0
     tx_fifo_errors: 0
     tx_heartbeat_errors: 0
     tx_window_errors: 0
     tx_abort_late_coll: 0
     tx_deferred_ok: 0
     tx_single_coll_ok: 0
     tx_multi_coll_ok: 0
     tx_timeout_count: 0
     tx_restart_queue: 0
     rx_long_length_errors: 0
     rx_short_length_errors: 0
     rx_align_errors: 0
     tx_tcp_seg_good: 604
     tx_tcp_seg_failed: 0
     rx_flow_control_xon: 0
     rx_flow_control_xoff: 0
     tx_flow_control_xon: 0
     tx_flow_control_xoff: 0
     rx_csum_offload_good: 450825
     rx_csum_offload_errors: 0
     rx_header_split: 0
     alloc_rx_buff_failed: 0
     tx_smbus: 0
     rx_smbus: 0
     dropped_smbus: 0
     rx_dma_failed: 0
     tx_dma_failed: 0
     rx_hwtstamp_cleared: 0
     uncorr_ecc_errors: 0
     corr_ecc_errors: 0
     tx_hwtstamp_timeouts: 0
     tx_hwtstamp_skipped: 0
lspci -vv
06:00.0 Ethernet controller: Intel Corporation 82574L Gigabit Network Connection
        Subsystem: Intel Corporation Gigabit CT Desktop Adapter
        Control: I/O- Mem+ BusMaster+ SpecCycle- MemWINV- VGASnoop- ParErr- Stepping- SERR- FastB2B- DisINTx+
        Status: Cap+ 66MHz- UDF- FastB2B- ParErr- DEVSEL=fast >TAbort- <TAbort- <MAbort- >SERR- <PERR- INTx-
        Latency: 0, Cache Line Size: 64 bytes
        Interrupt: pin A routed to IRQ 35
        Region 0: Memory at f72c0000 (32-bit, non-prefetchable) [size=128K]
        Region 1: Memory at f7200000 (32-bit, non-prefetchable) [size=512K]
        Region 2: I/O ports at c000 [disabled] [size=32]
        Region 3: Memory at f72e0000 (32-bit, non-prefetchable) [size=16K]
        Expansion ROM at f7280000 [disabled] [size=256K]
        Capabilities: [c8] Power Management version 2
                Flags: PMEClk- DSI+ D1- D2- AuxCurrent=0mA PME(D0+,D1-,D2-,D3hot+,D3cold+)
                Status: D0 NoSoftRst- PME-Enable- DSel=0 DScale=1 PME-
        Capabilities: [d0] MSI: Enable- Count=1/1 Maskable- 64bit+
                Address: 0000000000000000  Data: 0000
        Capabilities: [e0] Express (v1) Endpoint, MSI 00
                DevCap: MaxPayload 256 bytes, PhantFunc 0, Latency L0s <512ns, L1 <64us
                        ExtTag- AttnBtn- AttnInd- PwrInd- RBE+ FLReset- SlotPowerLimit 0.000W
                DevCtl: Report errors: Correctable+ Non-Fatal+ Fatal+ Unsupported+
                        RlxdOrd+ ExtTag- PhantFunc- AuxPwr- NoSnoop+
                        MaxPayload 128 bytes, MaxReadReq 512 bytes
                DevSta: CorrErr+ UncorrErr- FatalErr- UnsuppReq+ AuxPwr+ TransPend-
                LnkCap: Port #4, Speed 2.5GT/s, Width x1, ASPM L0s L1, Exit Latency L0s <128ns, L1 <64us
                        ClockPM- Surprise- LLActRep- BwNot- ASPMOptComp-
                LnkCtl: ASPM Disabled; RCB 64 bytes Disabled- CommClk+
                        ExtSynch- ClockPM- AutWidDis- BWInt- AutBWInt-
                LnkSta: Speed 2.5GT/s, Width x1, TrErr- Train- SlotClk+ DLActive- BWMgmt- ABWMgmt-
        Capabilities: [a0] MSI-X: Enable+ Count=5 Masked-
                Vector table: BAR=3 offset=00000000
                PBA: BAR=3 offset=00002000
        Capabilities: [100 v1] Advanced Error Reporting
                UESta:  DLP- SDES- TLP- FCP- CmpltTO- CmpltAbrt- UnxCmplt- RxOF- MalfTLP- ECRC- UnsupReq- ACSViol-
                UEMsk:  DLP- SDES- TLP- FCP- CmpltTO- CmpltAbrt- UnxCmplt- RxOF- MalfTLP- ECRC- UnsupReq- ACSViol-
                UESvrt: DLP+ SDES- TLP- FCP+ CmpltTO- CmpltAbrt- UnxCmplt- RxOF+ MalfTLP+ ECRC- UnsupReq- ACSViol-
                CESta:  RxErr- BadTLP- BadDLLP- Rollover- Timeout- NonFatalErr-
                CEMsk:  RxErr- BadTLP- BadDLLP- Rollover- Timeout- NonFatalErr+
                AERCap: First Error Pointer: 00, GenCap- CGenEn- ChkCap- ChkEn-
        Capabilities: [140 v1] Device Serial Number 68-05-ca-ff-ff-bd-3b-7d
        Kernel driver in use: e1000e
        Kernel modules: e1000e
I also have a tcpdump of the LLDP packets, if that could be of any help.
In addition to the above, I hacked a bit and got dropwatch to run on the Debian stable release Proxmox is based off of, and that did not display the LLDP packets, so there was not any info on why the packets were being dropped.
Unfortunately, I seem to have reached the end of Google when it comes to my specific issue, so hopefully someone here will be able to provide some resources to help me along.
