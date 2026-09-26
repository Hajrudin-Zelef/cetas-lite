---
id: collect-240926-storagereview/storagereview/fr-review-ubiquiti-enterprise-nas-review-16-bays-dual-25gbe-and-zfs-for-3999-25cfd5e9-2
title: "fr-review-ubiquiti-enterprise-nas-review-16-bays-dual-25gbe-and-zfs-for-3999-25cfd5e9"
domain: storagereview
role: reference
task: reference
actors: ["AWS", "Google"]
dates: ["2025-11"]
keywords: ["cost", "ethernet", "licenses", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-ubiquiti-enterprise-nas-review-16-bays-dual-25gbe-and-zfs-for-3999-25cfd5e9.md
source_anchor: ""
source_lines: [3, 61]
sha256: dc1c955161640e9ee196b82989d1c8456bb2498782493477568a3d47bda76c7e
---

# fr-review-ubiquiti-enterprise-nas-review-16-bays-dual-25gbe-and-zfs-for-3999-25cfd5e9

The Enterprise NAS (ENAS) is Ubiquiti's first attempt at offering a lighter enterprise storage solution. This 3U chassis includes 16 hot-swappable drive bays, an eight-core Arm Neoverse N2 processor, 64 GB of ECC memory, two SFP28 25 GbE ports, and two 550 W hot-swappable redundant power supplies, all for a suggested retail price of $3,999. Note: the Ubiquiti store is currently applying a memory surcharge at checkout due to rising memory prices. The final price is therefore higher than the list price until component costs stabilize. The file system used is ZFS, the bays accept all drives, and all software is included. No licenses, no feature unlocks, and no support contracts are required.
The price of a bare NAS may seem high at first glance, but it offers many advantages. Dual 25 GbE connectivity as standard is a major asset that the competition cannot match: Synology's 16-bay RS4021xs+ ships with two 10 GbE RJ45 ports and incorporates 25 GbE connectivity via a PCIe expansion card. The observation is similar for rackmount NAS units of equivalent capacity, where network connectivity above 10 GbE is a paid option. Add redundant power and ECC memory to that, and Ubiquiti meets the criteria that traditionally distinguished SMB NAS from entry-level enterprise arrays. ENAS is available now directly from Ubiquiti (affiliate link).
The ENAS crowns a lineup expansion carried out at a remarkable pace. Ubiquiti has built a solid reputation with its access points and gateways, and its networking offering continues to grow. In the past year alone, we tested the UniFi E7 and E7 Campus WiFi 7 access points, the Cloud Gateway Fiber gateway, and the Dream Router 7 router. Storage is a more recent area. The lineup began with the UNAS Pro, a 2U seven-bay enclosure sold for $499, aimed primarily at loyal UniFi customers. It was then expanded with the compact UNAS 2 and the UNAS Pro 8, which added redundant power and NVMe cache. The ENAS is the first model in the family to include ZFS, ECC memory, native iSCSI, and SAS expansion ports. It is important to note that the file systems differ enough that Ubiquiti specifies that UNAS data cannot be restored directly onto an ENAS; only users, groups, and settings are transferred.
The segment ENAS is entering is not empty. Synology and QNAP, among others, have dominated the SMB NAS market for twenty years, and their platforms offer far more comprehensive software access: application ecosystems, containers, surveillance suites, and high-performance backup tools that UniFi Drive does not currently offer. However, the established players have paved the way for new competitors. Synology spent most of 2025 enforcing a compatibility policy that effectively limited its 2025 Plus series units to only validated drives, primarily Synology-branded, before reversing its decision with DSM 7.3 in the fall. In this context, a 16-bay ZFS system advocating drive openness and with no software licensing represents a perfectly positioned product.
The ideal buyer profile is easy to identify: a small or medium-sized business, or the managed service provider operating it, that already has a medium-to-large UniFi deployment. In these environments, ENAS integrates into the same console as switches, gateways, access points, and cameras, inherits Site Manager for multi-site visibility, and integrates into the same identity model, with UniFi Endpoint managing remote file access. Storage is no longer a separate vendor with a separate user interface and renewal. For businesses that need simple, reliable file and block storage and are already familiar with the UniFi interface, operational continuity relies as much on the product as on the hardware.
A clarification regarding our test unit: we have had the ENAS in the lab since November 2025, and the production hardware revision changed the rear I/O configuration, moving from 10 GbE to the dual 25 GbE interface available today. Everything presented and tested here corresponds to the hardware currently being sold, and the software has evolved rapidly since we received our unit; we will revisit this.
Ubiquiti Enterprise NAS Specifications
| Specifications | Ubiquiti Enterprise NAS | 
|---|---|
| Market |  | 
| Dimensions | 481.4 × 480 × 132 mm (19 × 18.9 × 5.2 inches) | 
| Storage capacity | 16 × 2.5/3.5-inch drive bays 2 × M.2 NVMe bays | 
| Network interface | 2 × 25G SFP28 (25G/10G/1G) 1 × 10GbE RJ45 (10G/5G/2.5G/1G/100M) | 
| Expansion port | 2 × SFF-8644 (24G) | 
| Power redundancy | Yes | 
| Form factor | 3U rackmount | 
| Hardware |  | 
| Drive support | 16 × 2.5/3.5-inch HDDs/SSDs 2 × M.2 NVMe SSDs 2 expansion ports | 
| Maximum power budget for drives | 450W | 
| Maximum power consumption | 550W | 
| Power method | Hot-swappable power supply modules, dual AC input | 
| Power supply | 2 × 550 W hot-swappable AC/DC power supply modules | 
| Processor | Eight-core ARM N2 processor clocked at 2.4 GHz | 
| Memory | 64GB | 
| Management | Ethernet | 
| RF interface | Bluetooth 4.1 | 
| Weight | 16.1 kg (35.5 lb) | 
| Enclosure material | SGCC steel | 
| Mounting hardware | SGCC steel | 
| Supported rack depth | Rails support four-post racks of 600 mm (23.6 inches) with square holes (9.5 × 9.5 mm). Post depth of 600 to 1066 mm (23.6 to 42 inches) | 
| Front panel | 3U Bezel (4.7-inch touchscreen, RGBW LEDs) 3U Bezel Lite (blank) Both are optional | 
| LEDs |  | 
| Status LEDs | Ethernet, SFP28, hard drive, system, expansion port, CRPS | 
| Environmental compliance |  | 
| Operating temperature | -5 °C to 40 °C (23 °F to 104 °F) | 
| Operating humidity | 5% to 95% non-condensing | 
| NDAA compliant | Yes | 
| Certifications | FCC, CE, IC | 
| Software |  | 
| Supported file protocols | NFS, SMB | 
| Supported block protocols | iSCSI | 
| RAID types | Mirror, RAID-Z1, RAID-Z2, RAID-Z3 | 
| RAID groups | Multiple | 
| Spare drive support | Yes | 
| Personal and shared drives | Yes | 
| SSD cache | Yes | 
| Maximum NVMe SSD capacity | 8 TiB | 
| File encryption | Yes | 
| Backup support | Remote UNAS server, CIFS/SMB server, cloud services | 
| Cloud backup services | Google Drive, OneDrive, Dropbox, Amazon S3, Backblaze B2, Wasabi | 
| Snapshots | Yes | 
| Share links | Yes | 
| Time Machine backup | Yes | 
| Client application support | Yes | 
| User groups | Yes | 
Build and design
The ENAS is designed to different standards than the rest of the UniFi storage lineup, resembling the rackmount servers used in labs more than the UNAS Pro from which it is derived. The differences begin at the drive bays. The ENAS bays use a true eject button with an integrated activity LED, replacing the push-to-eject system of the UNAS Pro. This small change greatly simplifies hot-swapping when the rack is plunged into darkness and a drive needs to be removed. A cover in the upper right of the front panel conceals a connector that we will discuss later in this section.
The aluminum handles on the rackmount brackets serve a dual function: they allow the mounting latches to be unlocked and hold the optional side panel in place. Ubiquiti ships the ENAS with King Slide rails, identical to those used on servers from major manufacturers, and their installation in the rack is tool-free. Installing the inner rail on the chassis is tricky and requires a screwdriver and the screws located underneath the unit. Once this is done, the ENAS slides into place and locks using the latches at the top of each handle.
It is also at the drive trays that the cost savings are apparent. The release lever is plastic, unlike that of the UNAS Pro, which appears to be aluminum; a less judicious choice for a part handled every time a drive is moved. Installing 3.5-inch drives is tool-free, with an optional locking screw. However, 2.5-inch drives still require four screws through the bottom of the tray. If you install SSDs in all 16 bays, that is 64 screws that will prevent you from installing all your drives in the enclosure.
