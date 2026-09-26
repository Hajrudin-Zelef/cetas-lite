---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d10-hbas-tri-mode-and-expanders
title: "D10 — HBAs, tri-mode, and expanders"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Broadcom"]
dates: []
keywords: ["latency", "serdes", "throughput"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [334, 368]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: deea16239dced145c286e08a35b48e055fde5078f74c1313b753cfdc01edf23e
---

# D10 — HBAs, tri-mode, and expanders

## D10 — HBAs, tri-mode, and expanders

- **HBA vs RAID card:** an HBA (Host Bus Adapter) presents disks directly to the OS with no RAID logic; an eHBA (enhanced HBA, Broadcom 9600 terminology) adds enterprise management features (secure boot, SPDM attestation, SED key management, UBM) while still passing disks through [secondary].
- **Broadcom HBA 9600 series:** PCIe Gen4 x8, 24G SAS tri-mode; 9600-16i (05-50111-00, internal, 16 ports, 15 W, MTBF 5M hours per reseller spec sheet); 9600-24i (05-50111-01, 24 ports); 9600W-16e (external, SFF-8674); up to 240 SAS/SATA or 32 NVMe devices per controller [secondary].
- **Broadcom HBA 9500 series:** PCIe Gen4 tri-mode, 12G SAS generation (e.g. 9500-16e 05-50075-00 external, 9500-8i 05-50077-03); connects up to 1024 SAS/SATA or 32 NVMe [secondary].
- **Tri-mode SerDes:** one adapter concurrently serves NVMe, SAS, and SATA devices in the same backplane, auto-negotiating protocol and speed (SAS 22.5/12/6 Gb/s lanes, SATA 6 Gb/s, PCIe up to 16 GT/s per lane); enables non-disruptive migration from SAS/SATA to NVMe [secondary].
- **SAS expanders:** fan out one HBA port to dozens of drives; Broadcom SAS4x24 expander (24G SAS-4) doubles per-lane bandwidth vs 12G; AIC launched a SAS4 JBOD family (4U 108-bay, 78-bay, 2U 24-bay) on Broadcom SAS4x expanders at SC24 (Nov 2024) with dual-hub expander modules for dual-path redundancy [secondary].
- **Enclosure services (SES/SES-2):** expanders and backplanes report drive presence, temperature, voltage, and locator/fault LEDs; UBM (Universal Bay Management, SFF-TA-1005) standardizes tri-mode bay management [secondary].
- **Cabling notes:** SFF-8654 (SlimSAS) internal and SFF-8674 external on 24G parts; existing SFF-8639 (U.2) backplanes are compatible; DataBolt2 lets 24G controllers use 12/6G drives and backplanes at aggregated bandwidth [secondary].
- **Crossflashing caution:** many HBAs ship as RAID cards flashed to IT/HBA firmware (or vice versa). For ZFS/Ceph, use genuine HBA firmware or properly crossflashed IT-mode firmware; a RAID card in "JBOD mode" is not equivalent (see D12) [secondary].
- **9500 vs 9600 at a glance:** 9500 = PCIe Gen4, 12G SAS tri-mode (mature, cheaper, plenty for HDD/SATA-SSD estates); 9600 = PCIe Gen4, 24G SAS tri-mode, eHBA security features (secure boot, SPDM attestation), higher NVMe device counts. Buy 9600 for new NVMe-heavy builds; 9500 remains rational for bulk HDD [secondary].
- **Port/cable map:** internal SFF-8654 (SlimSAS x4/x8), external SFF-8674; U.2 (SFF-8639) backplanes compatible via tri-mode. One x8 SlimSAS cable carries 8× 24G SAS lanes ≈ 192 Gb/s raw to an expander/backplane [secondary].
- **Expander topology rules:** SAS4x24-class expanders fan one HBA port to dozens of bays; cascade expanders enclosure-to-enclosure (AIC's SAS4 JBODs support this); keep cascade depth shallow (1–2) to bound latency and failure blast radius [secondary].
- **Dual-domain multipath:** dual expander modules + dual-port SAS drives give two independent paths; OS multipath (Linux dm-multipath, Windows MPIO) survives a cable/expander/HBA-port failure. SATA drives are single-port — no true dual-path for SATA [secondary].
- **Zoning:** expander zoning carves one JBOD into isolated groups for multiple hosts — useful for shared-JBOD clusters without a full SAN [secondary].
- **SES/SES-2 and UBM:** enclosure services report per-slot presence, fault/locate LEDs, temperature, and voltage; UBM (SFF-TA-1005) unifies this across NVMe/SAS/SATA tri-mode bays so one backplane design serves all three [secondary].
- **Power/thermal:** 9600-16i eHBA is rated ~15 W (reseller spec) — plan PCIe slot power and airflow; 24-port and RAID cards with cache run hotter [secondary].
- **Bandwidth math for planning:** 24G SAS = 24 Gb/s ≈ 2,400 MB/s per lane after encoding; a x4 SFF-8654 link to an expander ≈ 9,600 MB/s raw. 12 HDDs at 250 MB/s = 3,000 MB/s — one x4 24G link feeds them with headroom; 24 SATA SSDs at 550 MB/s = 13,200 MB/s — needs dual x4 links or fewer drives per expander uplink [secondary].
- **Oversubscription rule:** expander uplinks are routinely oversubscribed for HDD (disks rarely stream simultaneously); for all-flash NVMe/SAS-SSD shelves, provision uplinks near 1:1 with aggregate drive bandwidth or accept contention [secondary].
- **SATA tunneling caveat:** SAS expanders tunnel SATA via STP; SATA drives don't support dual-porting, so a "dual-domain" shelf with SATA drives still has single-path drives — dual-path requires SAS drives [secondary].
- **Interop checklist before buying:** HBA firmware ↔ expander firmware ↔ backplane (SES/UBM) ↔ drive firmware. Tri-mode multiplies the matrix — get the vendor's tested-configuration list in writing for NVMe+SATA mixed shelves [secondary].
- **HBA queue depth:** enterprise HBAs expose deep queues (1K+ per device); a shallow-queue HBA starves NVMe SSDs capable of 64K+ queue depth — check queue-depth specs when the workload is high-IOPS NVMe [secondary].
- **Boot support:** HBAs (and eHBAs) support boot from attached devices via UEFI driver; RAID cards additionally offer bootable virtual disks. For ZFS-root or mdadm-root, HBA boot is sufficient [secondary].
- **Driver maturity:** `mpt3sas` (12G) and `mpi3mr`/`mpt3sas`-successor drivers for 24G parts are in-tree on modern Linux; verify the exact kernel's driver supports the card's PCI ID before deploying a new distro release [secondary].
- **LED/locator integration:** SES-managed fault/locate LEDs are the physical-safety layer for drive pulls — confirm the HBA↔backplane SES path works (test a locate blink) before you need it at 2 a.m. [secondary].
- **Cable length and signal integrity:** 24G SAS passive cables are length-limited (typically ≤1 m internal, check vendor spec); longer runs need active cables or retimers. CRC errors (SATA SMART 199 / SAS PHY error counters) that rise with cable length are a cabling problem, not a drive problem [secondary].
- **PHY error counters:** `smartctl -l sasphy` exposes per-PHY invalid DWORDs, disparity errors, and link resets — the diagnostic layer below SMART for flaky links [secondary].
- **Expander firmware:** expanders run firmware too (SES processor) — include it in the firmware inventory (D3); a hung expander looks like many simultaneous drive failures [secondary].
- **External SAS for JBOD shelves:** 9600-16e/9500-16e external ports (SFF-8674) connect rack JBOD shelves; keep external cable runs short and latching — a bumped external SAS cable is a classic multi-disk "failure" [secondary].
- **Enclosure numbering:** record enclosure:slot topology (`sas3ircu`/`sas3flash` topology pages); after expander replacement, re-verify the map — SES addresses can shift [secondary].
- **RAID 0 stripe width:** wider stripes help sequential throughput; RAID 0's risk scales with member count — 8-disk RAID 0 has ~8× the failure exposure of one disk [secondary].
- **Nested RAID naming:** RAID 10 = stripe of mirrors (1+0); RAID 01 = mirror of stripes (obsolete, weaker) — the order matters [secondary].

## D11 — NVMe hardware RAID: GRAID SupremeRAID and the software alternative

