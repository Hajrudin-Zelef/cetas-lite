---
id: collect-240926-storagereview/storagereview/fr-review-serial-cables-pcie-gen4-m-2-adapter-review-1c1ebfe3
title: "fr-review-serial-cables-pcie-gen4-m-2-adapter-review-1c1ebfe3"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Samsung"]
dates: []
keywords: ["amd", "consumer"]
source: docs/RAG/clean_en/storagereview/fr-review-serial-cables-pcie-gen4-m-2-adapter-review-1c1ebfe3.md
source_anchor: ""
source_lines: [1, 15]
sha256: 6adc9206408d60c24c2b2ab2f3f6c533606f442831f57294e2128d1dc1b85401
---

# fr-review-serial-cables-pcie-gen4-m-2-adapter-review-1c1ebfe3

<!-- source: https://www.storagereview.com/fr/review/serial-cables-pcie-gen4-m-2-adapter-review -->

The Serial Cables PCIe Gen4 m.2 Adapter (PCI4-AD-x4M2-04-G4) allows test labs and users to add an additional PCIe x4 Gen4 slot to desktop computers and servers with a supported motherboard. Traditionally, m.2 adapters offer companies an economical boot drive alternative for servers (without occupying a front mounting bay); however, this highly customizable adapter can be used for much more than that.
The Serial Cables Gen4 adapter is compatible with virtually all m.2 SSD sizes, as it features plated holes for 2230 (30mm), 2242, 2260, 2280, and full-length 22110 (110mm) form factors. Currently, most Gen4 m.2 SSDs are available in one form factor size (2280), but a range of brackets allows users to add other non-Gen4 cards of different sizes. Although the adapter card is backward compatible with Gen3 drives, you will of course only see Gen3 speeds. This also applies to the host motherboard. If you add this card to a board that does not support Gen4, the drive will be limited to the Gen3 bandwidth of the PCIe slot.
When it comes to swapping drives, this type is unique and versatile. Where most consumer cards use a screw to hold an m.2 SSD in place for more "permanent" means, this card features different forms of quick-release tabs. One of these methods is a small spring-loaded feature that easily seats an m.2 drive. The other is a plastic tab on a pivot point (pick-shaped); simply install the drive into the m.2 slot, press down lightly, then slide the plastic tab over the drive to hold it in place. Eliminating the requirement for a screwdriver will certainly be welcomed by those with use cases that involve constantly swapping drives, as installing m.2 drives on traditional adapter cards can be tedious and annoying.
At the top of the adapter, you will see a range of familiar pins, including CLKREQ # (clock request signal), WAKE # (wake functionality), and PREST # (used to specify power voltage settings). You will also notice a tactile switch (i.e., a button that starts or stops a flow of current along a circuit) located at the bottom left of the PCB. This has no function at the moment, as Serial Cables added this button for a specific customer who was doing SSD development. Look for firmware updates to enable this functionality in the future.
Performance
To demonstrate the performance of the Serial Cables PCIe Gen4 m.2 adapter, we used the Lenovo P620 (a powerful workstation equipped with an AMD Threadripper PRO processor with PCIe Gen4 support) and ran the Blackmagic performance test with the following configurations:
- equipped the adapter with a Samsung 980 Pro PCIe Gen4 SSD and installed it in one of the workstation's P620 PCIe slots
- installed the Samsung 980 Pro directly into the native slot inside the workstation
The goal was to show that the adapter matches and/or exceeds the onboard slot of a tier-1 workstation. Using the Serial Cables adapter inside the P620, the 980 Pro reached 5.29 GB/s read and 4.36 GB/s write.
These results were virtually identical to the speeds recorded when the drive was installed directly on the board, as the Samsung Pro showed 5.28 GB/s read and 4.34 GB/s write.
Conclusion
While we are simply using the Serial Cables PCIe Gen4 m.2 adapter as a pass-through for M.2 NVMe SSDs to a host system that supports Gen4, you can do much more with it in a test lab or engineering lab that works with a lot of drives. As such, the Serial Cables adapter is certainly not for everyone, but it is much better built than the ordinary $15 m.2 adapters you might find at your favorite online retailer.
This is an extremely niche adapter card that can be configured to accomplish a range of highly technical tasks. Coupling all of this with its performance that exceeded onboard speeds, there are not many cards like this one.
