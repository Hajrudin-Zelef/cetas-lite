---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-8
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Alibaba", "Mistral", "Nvidia", "SGLang", "vLLM"]
dates: ["2026-09-27"]
keywords: ["amd", "capex", "fp8", "gpu", "hbm", "hbm3", "kv cache", "llama", "mistral", "nvidia", "nvlink", "open source"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [1258, 1437]
sha256: 2061df40442b48695df9cff7c7ea79ba66fff3ef1595dd2f0524b81596c2a190
---

# Datacenter Builds — Le guide des BOMs

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 2P EPYC, 24 baies | type 2U24E nu | 1 | ≈ 3 970 € |
| EPYC 9655P (96c) | 2× | 2 | 5 208 € × 2 = 10 416 € |
| RAM 24× 64 Go = 1,5 To | DDR5 ECC | 24 | 33 120 € |
| 2× M.2 960 Go RAID 1 (OS) | — | 2 | 360 € |
| 16× NVMe 7,68 To (123 To bruts/nœud) | PM9A3 | 16 | 62 640 € |
| NIC 2× 100 GbE (ingest + réplication) | CX6-DX | 1 | 1 100 € |
| PSU 2× 2 000 W | incluses | — | — |
| **TOTAL / nœud** | | | **≈ 111 605 €** |
| **× 3 nœuds** | | | **≈ 334 815 €** |
| P_max / nœud | | | **≈ 1,6 kW** |

Capacité chaude cluster : 3 × 123 To × 0,85 ≈ 313 To → **≈ 170 j de rétention
chaude à 2 To/j** (avec réplica 2 du SIEM). Confortable.

---

## 75. Stockage tiède/froid : l'objet Ceph

- Au-delà de 90 j, basculez vers un **pool objet Ceph EC 4+2** (section 57) :
  ≈ 310 €/To utile vs 3 100 €/To en NVMe chaud.
- 2 To/j × 275 j (reste de l'année) = 550 To/an → **≈ 170 000 €** en hybride
  vs 1,7 M€ en NVMe. Le tiering n'est pas une option, c'est une obligation.
- Cycle de vie S3 (ILM) : automatisez la bascule chaud → tiède à J+90.

---

## 76. Réseau de capture : SPAN, TAP, NIC

| Besoin | Solution | Prix indicatif |
|---|---|---|
| Capture 1–10G | SPAN sur switch existant | 0 € (mais perte de paquets en surcharge) |
| Capture 10G fiable | TAP optique passif | ≈ 500–1 500 €/lien (à vérifier) |
| Sonde NDR 25G | NIC avec timestamp hardware | ≈ 1 500 € (à vérifier) |
| Agrégation | Packet broker (Garland, Cubro) | 10 000 €+ (à vérifier) |

Règle : **SPAN pour le dépannage, TAP pour la sécurité**. Un attaquant qui
sature le switch fait disparaître les paquets du SPAN — exactement ceux que
vous vouliez voir.

---

## 77. Cas chiffré : SOC 1 To/j complet (collecte + SIEM + froid)

| Poste | Détail | Total |
|---|---|---|
| 3× SIEM S (N+1, 500 Go/j chacun) | 55 895 € × 3 | 167 685 € |
| 2× collecteurs/syslog (1U légers) | 12 000 € × 2 | 24 000 € |
| Stockage froid 1 an (365 To utiles → Ceph hybride) | ≈ 310 €/To | 113 150 € |
| TAP + packet broker | lot | 15 000 € |
| Licences (Wazuh = 0 € ; commercial ≈ 100–200 k€/an) | open source | 0 € |
| **CAPEX** | | **≈ 319 800 €** |
| P_max | 3 × 0,95 + 2 × 0,4 + stockage | **≈ 5 kW** |

Avec un SIEM commercial (Splunk/QRadar), ajoutez **150–300 k€/an de licences**
au volume 1 To/j — souvent plus que le matériel sur 5 ans.

---

## 78. Rétention légale : ce que la loi impose (France)

- Logs de connexion (LCEN) : **1 an** pour les opérateurs/hébergeurs.
- Journaux de sécurité : pas de durée légale unique, mais l'ANSSI recommande
  6–13 mois pour la détection d'intrusions lentes (dwell time médian > 200 j).
- **Chiffrez les archives**, contrôlez l'accès (données personnelles), et
  documentez la politique de rétention (RGPD : minimisation).

---

## 79. Pièges terrain — SOC

1. **Collecter sans parser** : 1 To/j de logs bruts inexploitables = 1 To/j
   de coût disque pour zéro détection. Le parsing est le vrai travail.
2. **Rétention chaude infinie** : le NVMe coûte 10× l'objet. Tiering à J+90,
   sans exception.
3. **EPS pic vs moyen** : un ransomware génère 10× l'EPS normal. Si la file
   d'ingest déborde pendant l'attaque, vous perdez exactement les logs de
   l'attaque. Surdimensionnez l'ingest ×3.
4. **Horloges désynchronisées** : 1 seconde d'écart entre sources = corrélation
   impossible. NTP/Chrony partout, monitoré.
5. **Tester la restauration des archives froides** : un objet Ceph illisible à
   J+300 = rétention légale non conforme.

---

## 80. Ce qu'il faut retenir — SOC

- Dimensionnez à l'EPS et au Go/j, pas « au nombre de serveurs ».
- 90 j chauds NVMe + 1 an froid objet : le tiering fait le budget.
- Ingest ×3 vs moyenne, TAP plutôt que SPAN, NTP partout.
- Open source (Wazuh/ELK) : 0 € de licence ; commercial : 150–300 k€/an à 1 To/j.

---

## 81. Workload 7 — Inférence AMD : pourquoi, quand

Le MI300X (192 Go HBM3, 750 W) est l'arme **mémoire/prix** de 2026 :
- 192 Go sur une carte = un **70B en FP8/FP16 sur 1–2 cartes** au lieu de 4–8.
- Prix standalone ≈ 10 000–12 000 $ (vérifié le 27/09/2026) vs H100 à
  25 000–35 000 $ : **2–3× moins cher au Go de HBM**.
- Location cloud : ≈ 1,99 $/GPU/h (DigitalOcean, vérifié le 27/09/2026).
- **ROCm en 2026** : mature pour vLLM/SGLang sur les modèles courants
  (Llama, Qwen, Mistral). Vérifiez toujours la compatibilité de VOTRE modèle
  avant d'acheter — quelques architectures exotiques restent CUDA-only.

Quand choisir AMD : inférence de gros modèles, budget contraint, pas de besoin
CUDA propriétaire. Quand rester NVIDIA : training, écosystème CUDA verrouillé,
NVLink multi-nœuds.

---

## 82. Serveur 2× MI300X — BOM

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 2× GPU OAM/HGX AMD | plateforme MI300X 2-GPU (à vérifier — devis intégrateur) | 1 | ≈ 25 000 € |
| 2× AMD Instinct MI300X 192 Go | — | 2 | ≈ 10 000 € × 2 = 20 000 € |
| CPU EPYC 9555P (64c) | — | 1–2 | ≈ 4 830 € |
| RAM 12× 64 Go = 768 Go | DDR5 ECC | 12 | 16 560 € |
| 2× NVMe 3,84 To (modèles + cache) | PM9A3 | 2 | 6 070 € |
| NIC 2× 100 GbE (front) | CX6-DX | 1 | 1 100 € |
| PSU 3 000 W redondantes | incluses | — | — |
| **TOTAL** | | | **≈ 77 560 €** |
| P_max / réaliste | 2×750+360+120+60+60+150 | | **≈ 2,3 kW / 1,8 kW** |

384 Go HBM au total : un **70B FP8 (~70 Go) + KV cache confortable** sur les
2 cartes, ou 2× 32B en parallèle.

---

## 83. Serveur 4× MI300X — BOM (« quad »)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 4U 4× MI300X | plateforme 4-GPU (devis intégrateur) | 1 | ≈ 40 000 € |
| 4× MI300X 192 Go | — | 4 | 10 000 € × 4 = 40 000 € |
| 2× EPYC 9555P (64c) | — | 2 | 9 660 € |
| RAM 24× 64 Go = 1,5 To | DDR5 ECC | 24 | 33 120 € |
| 4× NVMe 7,68 To (modèles) | PM9A3 | 4 | 15 660 € |
| 2× NIC 200/400 GbE (scale-out) | — | 2 | ≈ 3 000 € × 2 = 6 000 € |
| PSU 4× 3 000 W (N+1) | incluses | — | — |
| **TOTAL** | | | **≈ 144 440 €** |
| P_max / réaliste | 4×750+720+240+120+120+200 | | **≈ 4,4 kW / 3,5 kW** |

768 Go HBM : **405B FP8 (~405 Go) sur un seul nœud**, ou 4× 70B en parallèle.
4,4 kW dans 4U : **refroidissement liquide ou RDHx obligatoire** (section 145+).

---

## 84. MI325X / MI355X : positionnement (prix non publics)

| GPU | HBM | TDP | Statut prix au 27/09/2026 |
|---|---|---|---|
| MI300X | 192 Go HBM3 | 750 W | 10 000–12 000 $ standalone ✓ |
| MI325X | 256 Go HBM3e | 750 W | non trouvé (prix public) |
| MI355X | 288 Go HBM3e | 1 000 W | non trouvé (prix public) — location ≈ 3 $/h |

Le MI355X vise le segment H200/B200. Sans prix public, **ne le chiffrez pas
dans une BOM ferme** : demandez un devis ou louez à l'heure pour valider.

---

## 85. VRAM vs taille de modèle — tableau de choix

| Modèle | FP16 | FP8 | Config mini recommandée |
|---|---|---|---|
| 8B | 16 Go | 8 Go | 1× L40S / 1× MI300X (large) |
| 32B | 64 Go | 32 Go | 1× L40S 48 Go (FP8) / 1× MI300X |
| 70B | 140 Go | 70 Go | 1× MI300X / 2× L40S / 1× RTX PRO 6000 (FP8) |
| 405B | 810 Go | 405 Go | 3× MI300X (FP8) / 6× H100 / 8× B200 |

Règle : **poids + 30 % (KV cache, overhead)** ≤ VRAM totale. Le FP8 divise par
2 vs FP16 avec une perte qualité faible sur les modèles récents — c'est le
format d'inférence par défaut en 2026.

---

## 86. vLLM / SGLang : sizing pratique

