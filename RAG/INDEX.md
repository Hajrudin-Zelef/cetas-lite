# INDEX — RAG

Corpus RAG de référence pour Cetas. Chaque corpus est une partition exacte de sa source, avec index et manifest.

Un corpus `delta` ne contient que les faits nouveaux/corrigés d'une source déjà couverte par un corpus `base` ; il ne la remplace pas. `delta_of` indique la base visée.

| corpus | titre | relation | fichiers | source | index |
|---|---|---|---|---|---|
| `briefing-ia-2026` | AI News 2026 — Reference Dossier |  | 126 | `docs/RAG/briefing-ia-2026-en.md` | [INDEX](briefing-ia-2026/INDEX.md) |
| `briefing-general-tech-2026` | General Tech News 2026 — Hardware, Infrastructure & Consumer Tech |  | 112 | `docs/RAG/briefing-general-tech-2026-en.md` | [INDEX](briefing-general-tech-2026/INDEX.md) |
| `ai-industry-kb-2026` | AI Industry Knowledge Base 2026 | base | 123 | `docs/RAG/ai-industry-knowledge-base-2026.md` | [INDEX](ai-industry-kb-2026/INDEX.md) |
| `ai-industry-kb-2026-wave6` | AI Industry Knowledge Base 2026 — Wave 6 Consolidation | delta de `ai-industry-kb-2026` | 136 | `docs/RAG/ai-industry-knowledge-base-2026-wave6.md` | [INDEX](ai-industry-kb-2026-wave6/INDEX.md) |
| `frontier-models-2026` | Frontier AI Models 2026 — Vague 1 (EN) |  | 12 | `docs/RAG/Grands titres IA modèlesEN.md` | [INDEX](frontier-models-2026/INDEX.md) |
| `labs-grok-platforms-2026` | Labs, Grok, Tools & Platforms 2026 — Vague 2 (EN) |  | 17 | `docs/RAG/Labos, Grok, outils & plateformes_EN.md` | [INDEX](labs-grok-platforms-2026/INDEX.md) |
| `open-local-models-2026` | Open / Local AI Models 2026 (EN) |  | 26 | `docs/RAG/Modèles IA open  locauxEN.md` | [INDEX](open-local-models-2026/INDEX.md) |
| `tools-platforms-2026` | AI Tools & Platforms 2026 (Step 2) |  | 8 | `docs/RAG/Outils & plateformes IAEN.md` | [INDEX](tools-platforms-2026/INDEX.md) |
| `etape4-trackd-unsloth-training` | Step 4 — Track D : Unsloth & outillage d'entraînement/fine-tuning 2026 |  | 3 | `docs/RAG/etape4_trackD_unsloth_training.md` | [INDEX](etape4-trackd-unsloth-training/INDEX.md) |
| `etape5-tracka-nvidia` | Step 5 — Track A : Nvidia (2026) |  | 4 | `docs/RAG/etape5_trackA_nvidia.md` | [INDEX](etape5-tracka-nvidia/INDEX.md) |
| `etape5-trackc-huawei-intel` | Step 5 — Track C : Huawei & Intel (2026) |  | 3 | `docs/RAG/etape5_trackC_huawei_intel.md` | [INDEX](etape5-trackc-huawei-intel/INDEX.md) |
| `etape4-trackc-cuda-rocm-pytorch` | Step 4 — Track C : CUDA / ROCm / PyTorch (2026) |  | 3 | `docs/RAG/etape4_trackC_cuda_rocm_pytorch.md` | [INDEX](etape4-trackc-cuda-rocm-pytorch/INDEX.md) |
| `etape5-trackb-amd` | Step 5 — Track B : AMD (2026) |  | 4 | `docs/RAG/etape5_trackB_amd.md` | [INDEX](etape5-trackb-amd/INDEX.md) |
| `etape4-tracka-vllm-sglang` | Step 4 — Track A : vLLM + SGLang (2026) |  | 20 | `docs/RAG/etape4_trackA_vllm_sglang.md` | [INDEX](etape4-tracka-vllm-sglang/INDEX.md) |
| `etape4-trackb-local-inference` | Step 4 — Track B : pile d'inférence locale (llama.cpp, Ollama, LM Studio) |  | 17 | `docs/RAG/etape4_trackB_local_inference.md` | [INDEX](etape4-trackb-local-inference/INDEX.md) |
| `etape5-trackd-servers` | Step 5 — Track D : serveurs IA, marché et réseau datacenter (2026) |  | 24 | `docs/RAG/etape5_trackD_servers.md` | [INDEX](etape5-trackd-servers/INDEX.md) |
| `labs-hyperscalers-2026` | Step 3 — Labs & Hyperscalers (2026) |  | 47 | `docs/RAG/Labos  hyperscalersEN.md` | [INDEX](labs-hyperscalers-2026/INDEX.md) |
| `etape6-trackc-huawei-mikrotik` | Step 6 — Track C : Huawei + MikroTik (matériel réseau) |  | 3 | `docs/RAG/etape6_trackC_huawei_mikrotik.md` | [INDEX](etape6-trackc-huawei-mikrotik/INDEX.md) |
| `etape6-trackd-firewalls` | Step 6 — Track D : FortiGate + pfSense + OPNsense (pare-feu & sécurité réseau) |  | 4 | `docs/RAG/etape6_trackD_firewalls.md` | [INDEX](etape6-trackd-firewalls/INDEX.md) |
| `etape6-trackb-arista-sonic` | Step 6 — Track B : Arista + SONiC + Cumulus (fabric datacenter & réseau ouvert) |  | 5 | `docs/RAG/etape6_trackB_arista_sonic.md` | [INDEX](etape6-trackb-arista-sonic/INDEX.md) |
| `etape6-tracka-cisco-juniper` | Step 6 — Track A : Cisco + Juniper (réseau datacenter entreprise) |  | 10 | `docs/RAG/etape6_trackA_cisco_juniper.md` | [INDEX](etape6-tracka-cisco-juniper/INDEX.md) |
| `etape10-phasea-news-tech` | Step 10A — Tech News: AI, Chips, Cloud, Cybersecurity (February → 22 September 2026) |  | 11 | `docs/RAG/etape10_phaseA_news_tech.md` | [INDEX](etape10-phasea-news-tech/INDEX.md) |
| `etape10-phaseb-news-mobile` | Step 10 — Phase B: Mobile & Devices News (February → September 2026) |  | 13 | `docs/RAG/etape10_phaseB_news_mobile.md` | [INDEX](etape10-phaseb-news-mobile/INDEX.md) |
| `etape10-phasec-news-pc-mac` | Step 10 Phase C — PC & Mac News (February → September 2026) |  | 14 | `docs/RAG/etape10_phaseC_news_pc_mac.md` | [INDEX](etape10-phasec-news-pc-mac/INDEX.md) |
| `etape10-phased-news-sport` | Step 10D — Sports News (February 2026 → 22 September 2026) |  | 13 | `docs/RAG/etape10_phaseD_news_sport.md` | [INDEX](etape10-phased-news-sport/INDEX.md) |
| `etape10-phasee-news-world` | Etape 10E — World News: Wars, Geopolitics, Elections, Economy, Fortunes (February 2026 – 22 September 2026) |  | 13 | `docs/RAG/etape10_phaseE_news_world.md` | [INDEX](etape10-phasee-news-world/INDEX.md) |
| `etape10-phasef-news-buzz` | STEP 10 — Phase F: News, Buzz & Society (2026-02-01 → 2026-09-22) |  | 11 | `docs/RAG/etape10_phaseF_news_buzz.md` | [INDEX](etape10-phasef-news-buzz/INDEX.md) |
| `etape10-phaseg-news-science` | Step 10 Phase G — Science & Planet News (February → September 2026) |  | 12 | `docs/RAG/etape10_phaseG_news_science.md` | [INDEX](etape10-phaseg-news-science/INDEX.md) |
| `etape6-phasea-vendors-dc` | Step 6 — Phase A: Enterprise Data-Center Switching Vendors |  | 38 | `docs/RAG/etape6_phaseA_vendors_dc.md` | [INDEX](etape6-phasea-vendors-dc/INDEX.md) |
| `etape6-phaseb-smb-networking` | Step 6 Extension — Phase B: SMB / Prosumer Networking |  | 42 | `docs/RAG/etape6_phaseB_smb_networking.md` | [INDEX](etape6-phaseb-smb-networking/INDEX.md) |
| `etape6-phasec-optics-cabling` | Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure |  | 52 | `docs/RAG/etape6_phaseC_optics_cabling.md` | [INDEX](etape6-phasec-optics-cabling/INDEX.md) |
| `etape6-phased1-fabrics-spine-leaf` | Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs |  | 13 | `docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md` | [INDEX](etape6-phased1-fabrics-spine-leaf/INDEX.md) |
| `etape6-phased2-evpn-vxlan` | Step 6 — Phase D wave 2: EVPN-VXLAN overlay |  | 16 | `docs/RAG/etape6_phaseD2_evpn_vxlan.md` | [INDEX](etape6-phased2-evpn-vxlan/INDEX.md) |
| `etape6-phased3-bgp-ha` | Phase D3 — BGP underlay and high availability in the data center |  | 13 | `docs/RAG/etape6_phaseD3_bgp_ha.md` | [INDEX](etape6-phased3-bgp-ha/INDEX.md) |
| `etape6-phased4-segmentation-qos-multicast` | Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services |  | 14 | `docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md` | [INDEX](etape6-phased4-segmentation-qos-multicast/INDEX.md) |
| `etape6-phasee1-netbox-nautobot` | Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File) |  | 11 | `docs/RAG/etape6_phaseE1_netbox_nautobot.md` | [INDEX](etape6-phasee1-netbox-nautobot/INDEX.md) |
| `etape6-phasee2-ansible-nornir-terraform` | Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries |  | 16 | `docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md` | [INDEX](etape6-phasee2-ansible-nornir-terraform/INDEX.md) |
| `etape6-phasee3-gnmi-openconfig-telemetry` | Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability |  | 14 | `docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md` | [INDEX](etape6-phasee3-gnmi-openconfig-telemetry/INDEX.md) |
| `etape6-phasee4-validation-cicd` | Phase E4 — Network Validation, Observability & CI/CD |  | 14 | `docs/RAG/etape6_phaseE4_validation_cicd.md` | [INDEX](etape6-phasee4-validation-cicd/INDEX.md) |
| `etape6-phasef1-nic-dpu-smartnic` | Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing) |  | 12 | `docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md` | [INDEX](etape6-phasef1-nic-dpu-smartnic/INDEX.md) |
| `etape6-phasef2-fpga` | Phase F2 — FPGA (Field-Programmable Gate Arrays) |  | 14 | `docs/RAG/etape6_phaseF2_fpga.md` | [INDEX](etape6-phasef2-fpga/INDEX.md) |
| `etape6-phasef3-platform-security` | Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust |  | 14 | `docs/RAG/etape6_phaseF3_platform_security.md` | [INDEX](etape6-phasef3-platform-security/INDEX.md) |
| `etape6-phasef4-virtualization-io` | Phase F4 — I/O Virtualization & CPU Acceleration Extensions |  | 12 | `docs/RAG/etape6_phaseF4_virtualization_io.md` | [INDEX](etape6-phasef4-virtualization-io/INDEX.md) |
| `etape7-phasea-os` | Step 7 — Phase A: Server Operating Systems (Linux OS Layer) |  | 14 | `docs/RAG/etape7_phaseA_os.md` | [INDEX](etape7-phasea-os/INDEX.md) |
| `etape7-phaseb-containers` | Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes |  | 12 | `docs/RAG/etape7_phaseB_containers.md` | [INDEX](etape7-phaseb-containers/INDEX.md) |
| `etape7-phasec-iac` | Step 7 Phase C — IaC & Platform Automation |  | 12 | `docs/RAG/etape7_phaseC_iac.md` | [INDEX](etape7-phasec-iac/INDEX.md) |
| `etape7-phased-proxmox-backup` | Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research |  | 13 | `docs/RAG/etape7_phaseD_proxmox_backup.md` | [INDEX](etape7-phased-proxmox-backup/INDEX.md) |
| `etape7-phasee-storage-software` | Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS) |  | 11 | `docs/RAG/etape7_phaseE_storage_software.md` | [INDEX](etape7-phasee-storage-software/INDEX.md) |
| `etape7-phasef-databases` | Step 7 — Phase F: Databases & Cache (Operations Angle) |  | 14 | `docs/RAG/etape7_phaseF_databases.md` | [INDEX](etape7-phasef-databases/INDEX.md) |
| `etape7-phaseg-linuxnet-vpn` | Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring |  | 12 | `docs/RAG/etape7_phaseG_linuxnet_vpn.md` | [INDEX](etape7-phaseg-linuxnet-vpn/INDEX.md) |
| `etape7-phaseh-webproxy` | Step 7 — Phase H: Web Servers, Reverse Proxies & API Gateways |  | 13 | `docs/RAG/etape7_phaseH_webproxy.md` | [INDEX](etape7-phaseh-webproxy/INDEX.md) |
| `etape7-phasei-media` | Step 7 — Phase I: Media Servers, Transcoding and Upscaling |  | 16 | `docs/RAG/etape7_phaseI_media.md` | [INDEX](etape7-phasei-media/INDEX.md) |
| `etape8-phasea-system-languages` | Step 8 — Phase A: Systems Languages |  | 12 | `docs/RAG/etape8_phaseA_system_languages.md` | [INDEX](etape8-phasea-system-languages/INDEX.md) |
| `etape8-phaseb-web-mobile-languages` | Step 8 — Phase B — Web & Mobile Languages, Runtimes and Toolchains |  | 12 | `docs/RAG/etape8_phaseB_web_mobile_languages.md` | [INDEX](etape8-phaseb-web-mobile-languages/INDEX.md) |
| `etape8-phasec-backend-frameworks` | Step 8 — Phase C: Backend Frameworks & APIs |  | 10 | `docs/RAG/etape8_phaseC_backend_frameworks.md` | [INDEX](etape8-phasec-backend-frameworks/INDEX.md) |
| `etape8-phased-frontend` | Step 8 — Phase D: Frontend & Web Platform (Dev Angle) |  | 1 | `docs/RAG/etape8_phaseD_frontend.md` | [INDEX](etape8-phased-frontend/INDEX.md) |
| `etape8-phasee-aiml-stacks` | Step 8, Phase E — AI/ML Software Stacks (Developer-Facing) |  | 13 | `docs/RAG/etape8_phaseE_aiml_stacks.md` | [INDEX](etape8-phasee-aiml-stacks/INDEX.md) |
| `etape8-phasef-data-messaging` | Step 8 Phase F — Data & Messaging: Developer Angle |  | 15 | `docs/RAG/etape8_phaseF_data_messaging.md` | [INDEX](etape8-phasef-data-messaging/INDEX.md) |
| `etape9-phasea-enterprise-ssd` | Step 9 — Enterprise SSD Hardware (Phase A) |  | 13 | `docs/RAG/etape9_phaseA_enterprise_ssd.md` | [INDEX](etape9-phasea-enterprise-ssd/INDEX.md) |
| `etape9-phaseb-form-factors-interfaces` | Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle) |  | 12 | `docs/RAG/etape9_phaseB_form_factors_interfaces.md` | [INDEX](etape9-phaseb-form-factors-interfaces/INDEX.md) |
| `etape9-phasec-memory` | Step 9 — Phase C: Server & Accelerator Memory Hardware |  | 12 | `docs/RAG/etape9_phaseC_memory.md` | [INDEX](etape9-phasec-memory/INDEX.md) |
| `etape9-phased-data-protection-raid` | Step 9 — Phase D: Data Protection & RAID Hardware |  | 11 | `docs/RAG/etape9_phaseD_data_protection_raid.md` | [INDEX](etape9-phased-data-protection-raid/INDEX.md) |
| `etape9-phasee-storage-market` | Step 9 Phase E — Storage & Memory Market 2026 |  | 8 | `docs/RAG/etape9_phaseE_storage_market.md` | [INDEX](etape9-phasee-storage-market/INDEX.md) |
| `collect-korben` | Korben.info — veille IA & tech (2026) |  | 18 | `docs/RAG/Collect RAG/01_korben` | [INDEX](collect-korben/INDEX.md) |
| `collect-mindstudio` | MindStudio — IA locale, modèles open-weight & agents (2026) |  | 170 | `docs/RAG/Collect RAG/02_mindstudio` | [INDEX](collect-mindstudio/INDEX.md) |
| `collect-huggingface` | Hugging Face — fiches de modèles (2026) |  | 125 | `docs/RAG/Collect RAG/03_huggingface` | [INDEX](collect-huggingface/INDEX.md) |
| `collect-opencode-docs` | opencode — documentation officielle |  | 4 | `docs/RAG/Collect RAG/04_opencode_docs` | [INDEX](collect-opencode-docs/INDEX.md) |
| `collect-presse-fr` | Presse FR — IA & tech (2026) |  | 8 | `docs/RAG/Collect RAG/05_presse_fr` | [INDEX](collect-presse-fr/INDEX.md) |
| `collect-benchmarks` | Benchmarks — modèles 2026 |  | 5 | `docs/RAG/Collect RAG/06_benchmarks` | [INDEX](collect-benchmarks/INDEX.md) |
| `collect-tutoriels` | Tutoriels & reviews — IA (2026) |  | 6 | `docs/RAG/Collect RAG/07_tutoriels` | [INDEX](collect-tutoriels/INDEX.md) |
| `vague2-briefia` | Briefia — veille IA (Collect Vague 2) |  | 1 | `docs/RAG/Collect RAG Vague 2/01_briefia` | [INDEX](vague2-briefia/INDEX.md) |
| `vague2-datacamp` | DataCamp — articles IA, data & dev (Collect Vague 2) |  | 68 | `docs/RAG/Collect RAG Vague 2/02_datacamp` | [INDEX](vague2-datacamp/INDEX.md) |
| `vague2-nerdykings` | NerdyKings — guides IA & dev (Collect Vague 2) |  | 34 | `docs/RAG/Collect RAG Vague 2/03_nerdykings` | [INDEX](vague2-nerdykings/INDEX.md) |
| `vague2-vision-ia` | Vision-IA — newsletter IA (Collect Vague 2) |  | 55 | `docs/RAG/Collect RAG Vague 2/04_vision_ia` | [INDEX](vague2-vision-ia/INDEX.md) |
| `collect-240926-mindstudio` | MindStudio — agents, modèles & routing (Collect 240926) |  | 170 | `docs/RAG/clean_en/mindstudio` | [INDEX](collect-240926-mindstudio/INDEX.md) |
| `collect-240926-huggingface` | Hugging Face — modèles & fiches (Collect 240926) |  | 125 | `docs/RAG/clean_en/huggingface` | [INDEX](collect-240926-huggingface/INDEX.md) |
| `collect-240926-storagereview` | StorageReview — stockage & serveurs (Collect 240926) |  | 74 | `docs/RAG/clean_en/storagereview` | [INDEX](collect-240926-storagereview/INDEX.md) |
| `collect-240926-datacamp` | DataCamp — IA, data & dev (Collect 240926) |  | 69 | `docs/RAG/clean_en/datacamp` | [INDEX](collect-240926-datacamp/INDEX.md) |
| `collect-240926-tomshardware` | Tom's Hardware — CPU, GPU & PC (Collect 240926) |  | 68 | `docs/RAG/clean_en/tomshardware` | [INDEX](collect-240926-tomshardware/INDEX.md) |
| `collect-240926-vision-ia` | Vision-IA — newsletter IA (Collect 240926) |  | 55 | `docs/RAG/clean_en/vision-ia` | [INDEX](collect-240926-vision-ia/INDEX.md) |
| `collect-240926-nerdykings` | NerdyKings — guides IA & dev (Collect 240926) |  | 34 | `docs/RAG/clean_en/nerdykings` | [INDEX](collect-240926-nerdykings/INDEX.md) |
| `collect-240926-korben` | Korben.info — veille IA & tech (Collect 240926) |  | 18 | `docs/RAG/clean_en/korben` | [INDEX](collect-240926-korben/INDEX.md) |
| `collect-240926-frandroid` | FrAndroid — mobile & tech (Collect 240926) |  | 14 | `docs/RAG/clean_en/frandroid` | [INDEX](collect-240926-frandroid/INDEX.md) |
| `collect-240926-hardwarecooking` | HardwareCooking — PC & matériel (Collect 240926) |  | 12 | `docs/RAG/clean_en/hardwarecooking` | [INDEX](collect-240926-hardwarecooking/INDEX.md) |
| `collect-240926-misc` | Collect 240926 — longue traîne (14 sites) |  | 29 | `docs/RAG/clean_en/misc` | [INDEX](collect-240926-misc/INDEX.md) |
| `collect-250926-servers-hardware` | Clean 4 — serveurs, GPU & stockage 2026 |  | 196 | `docs/RAG/clean4` | [INDEX](collect-250926-servers-hardware/INDEX.md) |

Voir aussi [README](README.md) et `manifest.json`.
