---
id: etape9-phased-data-protection-raid/00-data-protection-raid/part-9
title: "Step 9 — Phase D: Data Protection & RAID Hardware (part 9)"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [205, 213]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: f35d96daea4e5a972412c70f3ac048b0516ff3b3f16568ad35084693e6ae49bc
---

# Step 9 — Phase D: Data Protection & RAID Hardware (part 9)

- **NVMe SMART (Log 0x02) key fields:** `critical_warning` (bitmask — any non-zero bit is actionable: spare below threshold, temperature out of range, NVM subsystem reliability degraded, read-only mode, volatile backup device failed), `temperature` (composite), `available_spare` + `available_spare_threshold`, `percentage_used` (0 = new, 100 = endurance exhausted), `media_errors` (uncorrectable; should be 0), `data_units_read/written` (×512,000 bytes per unit), `power_cycles`, `unsafe_shutdowns`, error-log entries [secondary].
- **Alert thresholds used in practice:** critical_warning ≠ 0 → critical; media_errors > 0 → investigate; available_spare at/below threshold or < 20% → plan replacement; percentage_used ≥ 90% → endurance nearly exhausted [secondary].
- **ATA/SATA SSD attributes:** ID 5 Reallocated_Sector_Ct, 197 Current_Pending_Sector, 198 Offline_Uncorrectable, 199 CRC_Error_Count (cable/backplane), 177/202/230/231/233 wear-leveling variants (vendor-specific), 194 temperature [secondary].
- **SATA HDD attributes:** the same 5/197/198/199 quartet plus spin-retry and seek-error rates; Backblaze correlation studies link 5/187/188/197/198 to elevated failure probability [independent].
- **NVMe temperature thresholds:** drives expose Warning Composite Temperature Threshold (commonly ~70°C) and Critical (~85°C); exceeding warning sets critical_warning bit 1 and typically engages throttling [secondary].
- **Thermal throttling behavior:** controllers reduce clock/speed or suspend writes to hold temperature; simple implementations use 2–3 stages with visible performance cliffs, while multi-stage designs (e.g. ATP's AceTT: up to 18 stages starting at 85°C) taper gradually to maintain steadier throughput [vendor-reported].
- **Operational impact:** in dense chassis without directed airflow, Gen4/Gen5 NVMe SSDs under sustained load can sit at 70–85°C and throttle; enterprise ratings (e.g. PM9A3: 0–70°C operating) assume data-center airflow, not passive desktop conditions [secondary].
- **Monitoring stack:** `nvme smart-log` / `smartctl -a` for one-shots; node_exporter textfile collectors or smartctl-exporter for fleet telemetry; alert on any critical_warning bit, media_errors > 0, and temperature excursions — not just on FAILED self-assessment [secondary].

