---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-2
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "CoreWeave", "Lambda", "Nvidia", "TSMC", "vLLM"]
dates: ["2026-09-27"]
keywords: ["gpu", "amd", "blackwell", "compute", "fine-tuning", "fp4", "fp8", "hbm3", "inference", "kv cache", "llama", "mlperf"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [141, 278]
sha256: 4388e8fe714a237384fe01f90e8da8a11098283b0c83c14193426110cfdb95f7
---

# Les GPU datacenter / IA — présent et futur vérifié

- **Référence exacte** : **NVIDIA RTX PRO 6000 Blackwell** — existe en deux éditions
  (vérifié le 27/09/2026) :
  - **Workstation Edition** (ventilateurs, sorties 4× DisplayPort 2.1b), MSRP lancement **8 565 $**.
  - **Server Edition** (passive, sans sorties vidéo utilisables, TDP configurable 400–600 W).
- **Architecture** : Blackwell (GB202 plein, 92,2 milliards de transistors), TSMC 5 nm (4NP).
- **Cœurs** : **24 064 CUDA**, 752 Tensor (5e gén.), 188 RT (4e gén.), 128 Mo de cache L2.
- **Mémoire** : **96 Go GDDR7 avec ECC**, bus 512-bit, **1,79 To/s** — double du L40S.
- **TDP** : **600 W** (Server Edition : configurable 400–600 W), connecteur 1× 16-pin.
- **Format** : PCIe 5.0 x16. Pas de NVLink ; **MIG jusqu'à 4 instances** (4× 24 Go).
- **Précisions IA** : FP32 ~126 TFLOPS ; FP8 Tensor dense ~800 TFLOPS environ (ordre de
  grandeur ; sparse ~1 600) ; **FP4 sparse jusqu'à ~4 000 AI TOPS** (fiche tech-insider 2026).
  Moteurs médias : 1 par instance MIG.
- **Cas d'usage** : inférence LLM 70B quantifié sur 1 carte, fine-tuning, workstation IA,
  Omniverse/rendu pro. La seule carte « unitaire » haut de gamme qu'on achète à l'unité.
- **Prix** : lancement 8 565 $ ; mi-2026 **11 000–13 000 $** (+55 % sur pénurie GDDR7
  constatée par un guide d'achat). À vérifier au jour d'achat.
- **Où l'acheter** : PNY, Dell, HPE, Lenovo, Supermicro (serveurs 4× GPU, ex. Advantech
  SKY-622G4 2U MGX). Disponible à l'unité chez les distributeurs pro.

---

## 6. NVIDIA H100

- **Références exactes** (vérifiées le 27/09/2026) :
  - **H100 SXM5** (HBM3 80 Go) et **H100 NVL** (double GPU, 188 Go HBM3).
  - **H100 PCIe** : déclinaison 80 Go HBM3 (TDP 350 W) et **96 Go HBM3** (TDP 700 W, apparue
    en comparatifs 2025-2026 — la version PCIe 96 Go existe, bande passante ~1 681 Go/s).
- **Architecture** : Hopper (GH100), TSMC 4N, 80 milliards de transistors, 132 SM.
- **Mémoire** : **80 Go HBM3** (SXM : 3,35 To/s ; PCIe 80 Go : ~2 To/s), **141 Go** n'existe
  PAS sur H100 (c'est le H200). NVL : 2× 94 Go = 188 Go.
- **TDP** : **700 W** (SXM5), 350 W (PCIe 80 Go), 700 W (PCIe 96 Go).
- **Format** : SXM5 (baseboard HGX) ou PCIe double-slot FHFL.
- **Précisions IA** : FP64 31 TFLOPS ; FP32 62 TFLOPS ; FP16/BF16 Tensor dense 989 TFLOPS ;
  **FP8 Tensor dense 1 979 TFLOPS** ; INT8 ~3 958 TOPS (sparse ×2). Pas de FP4 (génération Hopper).
- **Cas d'usage** : **training + inférence** — le standard industriel 2023-2025, MLPerf partout.
- **Prix** : carte nue **25 000–31 000 $** ; serveur 8× H100 HGX **250 000–320 000 $**
  (~285 000 $ typique OEM) ; DGX H100 ~290 000 $. Location : 1,49–6,98 $/h (médiane 3,33 $/h).
- **Où l'acheter** : **uniquement en serveurs intégrés** (Dell, HPE, Lenovo, Supermicro, Gigabyte,
  NVIDIA DGX). Pas de vente carte nue au détail.

---

## 7. NVIDIA H200

- **Référence exacte** : NVIDIA H200 (SXM5 et NVL), successeur mémoire du H100
  (vérifiée le 27/09/2026).
- **Architecture** : Hopper (GH100), **même silicium compute que le H100**.
- **Mémoire** : **141 Go HBM3e**, **4,8 To/s** (+76 % de capacité, +43 % de bande passante
  vs H100). NVL : 2× 141 Go = 282 Go.
- **TDP** : **700 W** (SXM5, configurable). NVLink 4 (900 Go/s par GPU).
- **Format** : SXM5 (HGX) ou NVL PCIe double-GPU.
- **Précisions IA** : identiques au H100 (FP8 dense 1 979 TFLOPS, pas de FP4).
- **Cas d'usage** : **inférence gros modèles** (70B en FP16 sur 1 GPU : 141 Go ≥ ~140 Go
  requis), fine-tuning mémoire-intensive. « H100 à mémoire gonflée ».
- **Prix** : carte **30 000–40 000 $** ; serveur 8× H200 HGX **320 000–420 000 $** ;
  4× H200 NVL PCIe + serveur ~160 000–185 000 $ (BOM kinvert). Location : 3,72–10,60 $/h
  (hyperscalers), ~3,70 $/h on-demand médian.
- **Où l'acheter** : serveurs intégrés OEM (Dell/HPE/Lenovo/Supermicro), DGX H200,
  CoreWeave/lambda en location (~6,30 $/GPU/h chez CoreWeave).

---

## 8. NVIDIA B200

- **Référence exacte** : NVIDIA B200 (GPU Blackwell B200, double-die GB100+GB200 sur
  interposeur), déclinaisons **HGX B200** (air) et **GB200 NVL72** (rack).
  (Vérifiée le 27/09/2026.)
- **Architecture** : Blackwell, 208 milliards de transistors.
- **Mémoire** : **192 Go HBM3e** (8-high), **8 To/s** par GPU.
- **TDP** : **1 000 W** par GPU (HGX). Le saut thermique de la génération.
- **Format** : SXM (HGX 8 GPU) ou rack NVL72 (72 GPU + 36 CPU Grace).
- **Précisions IA** : FP16/BF16 dense 4 500 TFLOPS ; **FP8 dense 4 500 TFLOPS** ;
  **FP4 dense 9 000 TFLOPS** (sparse ~18 000). NVLink 5 : **1,8 To/s** par GPU.
- **Cas d'usage** : training frontier + inférence très gros modèles. MLPerf Training :
  jusqu'à **2,2× le H100**.
- **Prix** : carte **35 000–45 000 $** (estimation) ; serveur 8× B200 HGX **400 000–520 000 $**
  ; DGX B200 ~515 000 $ ; rack GB200 NVL72 ~3 000 000 $. Location : 4,99–6,02 $/h on-demand,
  spot dès 2,12 $/h, réservé 36 mois ~2,25 $/h.
- **Où l'acheter** : OEM (Dell/HPE/Supermicro/Lenovo), NVIDIA DGX, cloud (CoreWeave, Lambda).
  **Refroidissement liquide quasi-obligatoire en dense** (rack > 50 kW).

---

## 9. AMD Instinct MI300X

- **Référence exacte** : AMD Instinct MI300X (OAM), gfx942 / CDNA 3 (vérifiée le 27/09/2026).
- **Architecture** : CDNA 3, chiplets : 8 XCD (TSMC N5) + 4 IOD (TSMC N6), empilement 3D SoIC,
  304 unités de calcul.
- **Mémoire** : **192 Go HBM3**, **5,3 To/s** — 2,4× la capacité d'un H100.
- **TDP** : **750 W** (pic), format OAM passif, bus PCIe 5.0 x16 (hôte).
- **Précisions IA** : FP16 dense 1 307 TFLOPS ; **FP8 dense 2 615 TFLOPS (format FNUZ)** ;
  pas de FP4/FP6 natif sur CDNA 3. 256 Mo d'Infinity Cache.
- **Cas d'usage** : inférence gros modèles (le 192 Go avale un 70B FP16 entier),
  alternative crédible au H100/H200 sur vLLM/ROCm.
- **Prix** : pas de prix public carte nue (vente en plateformes 8 GPU). Cloud :
  **1,99 $/h** (DigitalOcean, juin 2025 — prix choc qui a fait école).
- **Où l'acheter** : plateformes OEM (Dell, HPE, Supermicro, Gigabyte) en serveurs 8× OAM ;
  cloud (DigitalOcean, Hot Aisle, Azure).

---

## 10. AMD Instinct MI325X

- **Référence exacte** : AMD Instinct MI325X (OAM), CDNA 3 refresh mémoire
  (vérifiée le 27/09/2026).
- **Architecture** : CDNA 3 (gfx942), même compute que MI300X.
- **Mémoire** : **256 Go HBM3e**, **6 To/s**.
- **TDP** : **1 000 W** (pic), OAM passif.
- **Précisions IA** : identiques MI300X (FP8 FNUZ 2 615 TFLOPS dense). Pas de FP4.
- **Cas d'usage** : inférence très gros modèles / KV cache géant, pont vers CDNA 4.
- **Prix** : non public à l'unité (plateformes 8 GPU). À vérifier.
- **Où l'acheter** : mêmes canaux OEM que MI300X. Positionnement : concurrent direct du H200
  (+115 Go de mémoire).

---

## 11. AMD Instinct MI350X et MI355X

- **Références exactes** : AMD Instinct **MI350X** (air) et **MI355X** (liquide),
  gfx950 / **CDNA 4** (vérifiées le 27/09/2026).
- **Architecture** : CDNA 4, TSMC N3 (XCD), 256 CU, refonte des matrix cores.
- **Mémoire** : **288 Go HBM3e**, **8 To/s** pour les deux.
- **TDP** : MI350X **1 000 W** (air) ; MI355X **1 400 W** (liquide, pleine perf).
- **Format** : OAM (passif pour 350X ; passif+actif pour 355X selon AMD).
- **Précisions IA** : FP16 dense 2,3 PFLOPS (350X) / 2,5 PFLOPS (355X) ;
  **MXFP4 9,2 / 10,1 PFLOPS** ; MXFP6 10,1 PFLOPS (355X) ; FP8 format OCP. Premier support
  natif FP4/FP6 chez AMD.
- **Cas d'usage** : training + inférence, concurrent direct du B200. MLPerf Inference v6.0
  (avril 2026) : MI355X à **92–104 % du B300** sur Llama 2 70B single-node (103 480 tok/s
  offline) — chiffres AMD, à prendre avec recul.
- **Prix** : non publics (plateformes OEM). À vérifier.
- **Où l'acheter** : Dell, HPE, Supermicro, Gigabyte (serveurs 8× OAM). Existe aussi en
  **MI350P** : version PCIe 144 Go HBM3e, 350 W, pour l'entreprise classique.

---

## 12. AMD Instinct MI300A

