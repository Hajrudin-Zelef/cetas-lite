---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d8-hardware-raid-controllers-2026-landscape
title: "D8 — Hardware RAID controllers: 2026 landscape"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: []
dates: []
keywords: ["memory", "nand"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [214, 246]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 959e1e2be2106eede472ef947786698e660c82b66792bb6560065c37e9a18f32
---

# D8 — Hardware RAID controllers: 2026 landscape

---
- **Backblaze's five failure-correlated HDD attributes (independent fleet data):** SMART 5 (Reallocated_Sector_Count), 187 (Reported_Uncorrectable_Errors), 188 (Command_Timeout), 197 (Current_Pending_Sector), 198 (Offline_Uncorrectable). In over 76% of observed failures, at least one of the five had a non-zero raw value before the drive died [independent].
- **Magnitude of correlation:** drives with non-zero SMART 5 show roughly 14× the failure rate of clean drives; non-zero 187/198 roughly 7.5×. Attribute 188 (Command_Timeout) is noisier — low counts are common from power management, but sustained high counts indicate controller/interconnect degradation [independent].
- **Rate matters more than count:** a drive jumping 0→20 on SMART 187 in one day is a far worse omen than a drive accumulating 60 over five years. Alert on *deltas*, not just thresholds [independent].
- **NVMe critical_warning bitmask (byte 0 of SMART log):** bit 0 = available spare below threshold; bit 1 = temperature above threshold or below lower threshold; bit 2 = NVM subsystem reliability degraded; bit 3 = media in read-only mode; bit 4 = volatile memory backup device failed (PLP/capacitor path — direct D1 relevance); bit 5 = persistent memory region read-only. Any set bit is a replace/mitigate-now signal [secondary].
- **SAS HDD deep telemetry:** `smartctl -l background` / SAS log pages expose grown-defect lists, IOEDC (I/O error detection code) counters, and background media scan results — richer than ATA SMART for pinpointing whether errors are media, head, or interconnect [secondary].
- **Monitoring stack (practical):** `smartctl -a` / `nvme smart-log` for spot checks; smartmontools + Prometheus `smartctl_exporter` or node_exporter textfile for fleets; Zabbix templates exist for both. Retention of SMART history enables the delta-based alerting Backblaze's data justifies [secondary].
- **Temperature policy:** NVMe exposes Warning and Critical Composite Temperature Thresholds (commonly ~70 °C / ~85 °C); exceeding warning sets critical_warning bit 1 and engages the controller's thermal policy. Enterprise ratings (PM9A3: 0–70 °C operating) assume data-center airflow — a Gen4/Gen5 drive in a passive desktop slot will throttle [secondary].
- **Throttling design spectrum:** simple controllers use 2–3 coarse stages (visible performance cliffs); ATP's AceTT uses up to 18 stages from 85 °C for gradual taper. Either way, sustained thermal throttling is a capacity-planning signal — add airflow or spread load, don't just accept it [vendor-reported].
- **Spot-check commands:** `smartctl -a /dev/sdX` (SATA/SAS), `smartctl -a /dev/nvme0` (NVMe via smartmontools), `nvme smart-log /dev/nvme0` (native NVMe log page 0x02), `nvme error-log /dev/nvme0` (error entries — any growth is significant) [secondary].
- **What "healthy" looks like:** `critical_warning: 0`, `media_errors: 0`, `num_err_log_entries` stable, `percentage_used` well under 100, `available_spare` at 100%, temperature under WCT, `unsafe_shutdowns` not climbing. Anything else is a ticket, not a shrug [secondary].
- **SSD-specific early warnings beyond SMART:** sudden sustained write-performance drop (GC distress / full drive), rising `percentage_used` faster than host writes imply (high WAF), and thermal throttling under previously-fine loads (heatsink/airflow degradation) [secondary].
- **HDD acoustic/mechanical tells:** SMART won't catch everything — new clicking, spin-up retries (SMART 10), and rising seek-error rates warrant proactive replacement even with SMART 5/197/198 at zero [secondary].
- **Retention policy for SMART history:** keep ≥12 months of per-drive SMART telemetry; Backblaze's rate-of-change finding (0→20 in a day vs gradual) is only actionable with history [independent].
- **Thermal design power reality:** a Gen4 NVMe SSD can draw 8–12 W under sustained load; 24 of them in a 2U shelf = ~250 W of heat in a small volume. Thermal throttling is a *cooling* problem first, a drive problem second [secondary].
- **Heatsinks and directed airflow:** motherboard M.2 slots with no heatsink routinely push client NVMe past 70 °C; enterprise U.2/E1.S carriers assume front-to-back chassis airflow. Match the carrier to the thermal spec, not just the connector [secondary].
- **Temperature vs retention:** high operating temperature accelerates NAND charge leakage (retention) and capacitor aging (PLP) simultaneously — the hot drive is attacked on two axes [secondary].
- **Alerting thresholds (sane defaults):** warning at WCT − 5 °C, critical at WCT; page on critical_warning bit 1 or any excursion past CCT. Log temperature histograms — a slow upward drift predicts cooling failure before throttling starts [secondary].
- **`nvme list-subsys` and topology:** map namespaces → controllers → PCIe addresses before alerting; a "failed drive" alert is useless without the slot/server mapping (combine with SES locate, D10) [secondary].
- **Error log triage:** `nvme error-log` entries with status codes like namespace-not-ready or LBA-out-of-range are often software bugs, not media; media/data-integrity errors are the ones that page [secondary].
- **Dashboarding:** Grafana + Prometheus smartctl_exporter gives per-drive temperature, percentage_used, and media_errors history — the three panels that predict most SSD incidents [secondary].
- **HDD-specific monitors:** load/unload cycle count (SMART 193) on laptop-class drives in 24/7 service; spin-retry (10); and helium-level attributes on He-filled drives where exposed [secondary].
- **Burn-in procedure:** before trusting a new drive, run full-device write + read-verify + SMART review; infant-mortality failures are cheapest when the drive holds no data [secondary].
- **Drive self-tests:** `smartctl -t long` (SATA) / NVMe device self-test log — schedule quarterly on HDDs; a self-test failure is a replace order even if all attributes look fine [secondary].
- **Log persistence across power loss:** SMART counters persist; temperature *history* may not on all drives — export telemetry continuously rather than relying on the drive's memory [secondary].
- **False positives:** a single CRC error (SMART 199) after reseating a cable is noise; a climbing 199 is a cable/backplane fault. Context distinguishes [secondary].
- **Over-provisioning via namespaces:** NVMe allows creating a smaller namespace than the physical capacity, donating the remainder as dynamic OP — a host-side lever when the drive's factory OP is insufficient [secondary].
- **TRIM and RAID passthrough matrix:** mdadm passes discards; ZFS issues them per vdev; hardware RAID usually doesn't — know your layer's behavior before assuming TRIM works [secondary].
- **Reporting hygiene:** when comparing drives, normalize to AFR at stated temperature/workload; "2M hr MTBF" without conditions is marketing, not engineering [secondary].
- **Infant mortality screening:** the first 90 days dominate early failures — burn-in plus close SMART watch in quarter one pays for itself [independent].

## D8 — Hardware RAID controllers: 2026 landscape

