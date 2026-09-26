---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/12-2-used-enterprise-ssds-for-homelab
title: "12.2 Used enterprise SSDs for homelab"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: ["2026-08", "2026-08-23"]
keywords: ["datacenter", "nand", "pricing"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [270, 289]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: 95dca09aa215b6609c1f810614c2d05aec4a9513ea9d1e5238f971a6629d9303
---

# 12.2 Used enterprise SSDs for homelab

- Market band (2026-08-23): new enterprise NVMe U.2 $300–$1,172/TB; nearline SATA flat $25/TB; used enterprise SATA/SAS 3.84 TB at $78–$127/TB [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).
- Samsung PM1743 1.92 TB (Lenovo SKU 4XB7A82609): $8,968 ($4,671/TB) — OEM channel markup, far above NAND-market pricing [secondary](https://www.shi.com/product/45962451/Samsung-PM1743-SSD).
- Samsung PM9D3a 7.68 TB (Lenovo 4XB7A93069): $17,362 ($2,261/TB; MSRP $16,899) [secondary](https://www.shi.com/product/48742687/Samsung-PM9D3a-SSD).
- Kioxia CD8P-R 30.72 TB SED (Lenovo KCD8DPUG30T7): $21,685.99 ($706/TB) — large-capacity SKUs price much better per TB [secondary](https://corgitech.us/products/lenovo-thinksystem-read-intensive-30-72tb-2-5-pcie-5-0-sed-kioxia-cd8p-r-ssd).
- Micron 9550 PRO 7.68 TB: €1,597 excl. VAT (≈€208/TB) [secondary](https://www.computeruniverse.net/en/p/3314-02V).
- WD SN861 3.2 TB (3 DWPD class): £572.81 excl. VAT (≈£179/TB) [secondary](https://www.convergetp.co.uk/wd-ultrastar-dc-sn861-wus6ca232psp9x1-ssd-data-centre-3-2-tb-u-2-pcie-5-0-x4-nvme-1tstosto-041686/).
- Phison X200E 6.4 TB at review time: $1,309.99 (≈$205/TB) — older figure, pre-2026 price surge [independent](https://www.techpowerup.com:443/review/phison-pascari-x200e/single-page.html).
- VDURA August 2026: 30 TB TLC drive at $22,600 (≈$753/TB) as a vendor-published reference point [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).

### 12.2 Used enterprise SSDs for homelab

- The used channel prices off decommissioned-hardware pools, not current NAND contracts — that's why used enterprise SATA/SAS SSDs (3.84 TB, $78–$127/TB) undercut new enterprise NVMe by 4–15x in 2026 [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).
- Popular homelab picks (secondary market): Samsung PM9A3 (Gen4, V6 TLC, 1 DWPD, PLP) — verify remaining endurance; older Gen3/Gen4 Kioxia/Samsung/SK hynix datacenter drives similarly.
- Pre-purchase checks (practice, independent guidance): read `nvme smart-log` — `percentage_used` (media wear), `power_on_hours`, `media_errors`, `available_spare` vs `available_spare_threshold`, `data_units_written` vs rated TBW. Avoid drives with non-zero `media_errors` or depleted spare. Prefer 1 DWPD+ TLC for write-heavy pools (ZFS SLOG/ZIL, Ceph journals); QLC fine for media libraries and backup targets.
- Vendor lock-in caution: some OEM-branded drives (Dell, HPE, Lenovo SKUs) ship custom firmware — firmware updates may require the OEM's tools; sector size may be 520/528 bytes (reformat to 512/4096 with `nvme format` where supported). Plain "Samsung"/"Kioxia" retail-channel SKUs avoid most of this [independent guidance].
- Cooling: Gen5 U.2 drives (19–25 W typical) need real airflow; E3.S/E1.S need backplane airflow specs (WD SN861 E1.S needs 1.5 m/s airflow at 30 °C for max performance) — passive M.2-style cooling is insufficient [secondary](https://cdn.multitronic.fi/media/c/d/mmo_133086493_1752741630_7228_401777.pdf).
- Endurance math for homelab: a 7.68 TB 1-DWPD drive rated 14,016 TBW survives ~3.8 TB/day of writes for 10 years — far beyond typical homelab write rates, so even 50%-used drives are usually fine [independent guidance].

---

