---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/13-4-hot-swap-power-disable-surprise-removal
title: "13.4 Hot-swap, power-disable, surprise removal"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [315, 326]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 81256c75ff857ac4386b9d8a0c01b43f3077bf3d5be3a8cc86e3280b99a8a600
---

# 13.4 Hot-swap, power-disable, surprise removal

- SGPIO (SFF-8485): serial GPIO bit-banging for drive LEDs/activity — the pre-UBM standard, still selectable on UBM chips via pin strapping ("seamless switching between SGPIO and UBM protocols") [vendor-reported](https://www.electronicsforu.com/electronics-projects/reference-design-for-universal-storage-backplane-management).
- SES (SCSI Enclosure Services, 2-wire): the SAS-expander enclosure protocol for slot/LED/thermal management [independent].
- Intel VPP (Virtual Pin Port): per-lane LED/control virtualization for NVMe direct-attach [independent].
- SES/LED/thermal remain the BMC-visible surface: UBM devices expose "I2C communication to Baseboard Management Controller" with up to 8 I2C ports, plus an optional sideband I2C for BMC direct access [vendor-reported](https://www.microsemi.com/product-directory/upcoming-technology/5558-universal-backplane-management-ubm).

### 13.4 Hot-swap, power-disable, surprise removal

- Enterprise bays require: per-slot power control, PCIe reset + reference clock management (UBM's job), and OS/driver surprise-removal handling [secondary](https://www.storagenewsletter.com/2018/06/18/broadcom-co-sponsor-for-sff-ta-1005-universal-backplane-management-spec/).
- NVMe power-disable is a UBM-managed feature on modern backplanes (Microchip EEC1005-UB2 lists it explicitly) [vendor-reported](https://www.microsemi.com/product-directory/upcoming-technology/5558-universal-backplane-management-ubm).
- HighPoint's Gen5 switch adapters advertise "hot-plug and hot-swap NVMe support allows drive replacement while the system remains online" with self-diagnostic LEDs and FRU support [vendor-reported](https://www.scan.co.uk/products/highpoint-rocket-1624a-hba-adapter-2x-mcio-gen5-x8-pcie-50-x16-32-gb-s-hot-swap-broadcom-pex89048-sw).
- Dell 17G E3.S chassis options distinguish "Cold Aisle Supported" vs "Hot Aisle" serviceability — front-serviceable NVMe is a chassis engineering feature, not just a connector [secondary](https://www.dell.com/en-au/shop/servers-storage-and-networking/poweredge-r570-rack-server/spd/poweredge-r570/promo_r570_1?view=configurations).

