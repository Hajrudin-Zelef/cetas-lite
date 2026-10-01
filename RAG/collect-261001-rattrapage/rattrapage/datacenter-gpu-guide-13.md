---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-13
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["datacenter", "gpu", "agents", "amd", "blackwell", "embeddings", "ethernet", "fine-tuning", "fp4", "fp8", "kv cache", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [1761, 1939]
sha256: c27037b61c6902830f33176ec2e3edf2eea886edf926a2b2b2c33a8080b97086
---

# Les GPU datacenter / IA — présent et futur vérifié

| Composant | Choix | Prix |
|---|---|---|
| 2× RTX 5090 32 Go | Retail | 5 200 $ |
| CPU Ryzen 9 9950X | 16c | 550 $ |
| 128 Go DDR5 | 4× 32 Go | 400 $ |
| NVMe 4 To | Gen4 | 250 $ |
| PSU 1 600 W ATX 3.1 | Titanium | 450 $ |
| Boîtier grand tour + airflow | — | 250 $ |
| **Total** | | **~7 100 $** |

Élec : ~1,5 kW → ~1 970 €/an en 24/7 (rarement le cas en lab).
**Limites** : 64 Go VRAM cumulés, pas d'ECC, pas de NVLink.

## 79. Cas pratique n°2 : inférence 70B éco (1× RTX PRO 6000)

**Besoin** : servir un 70B en continu, budget serré, 1 nœud.

| Composant | Choix | Prix |
|---|---|---|
| 1× RTX PRO 6000 Blackwell 96 Go | PNY | 12 000 $ |
| Serveur 2U 1× CPU EPYC 32c | Supermicro | 6 000 $ |
| 256 Go DDR5 ECC | — | 1 200 $ |
| NVMe 2 To | — | 150 $ |
| **Total** | | **~19 500 $** |

70B en FP8 (~90 Go avec KV cache) sur 1 carte : ~50 tok/s, conso ~700 W nœud.
**Le meilleur $/token de l'inférence 70B sur site en 2026.**

## 80. Cas pratique n°3 : plateforme RAG 8× L40S

**Besoin** : embeddings + rerank + LLM support, 10 k requêtes/heure.

| Composant | Choix | Prix |
|---|---|---|
| Serveur 4U 8× L40S 48 Go | Supermicro | ~110 000 $ |
| 2× EPYC 48c, 1 To RAM | Inclus | — |
| 2× ConnectX-7 200G | — | 4 000 $ |
| **Total** | | **~114 000 $** |

384 Go VRAM, 2,8 kW GPU, air classique. ROI vs cloud (8× 0,96 $/h = 67 k$/an) :
**~20 mois**.

## 81. Cas pratique n°4 : training 8× H100 (le standard)

**Besoin** : fine-tuning/full training jusqu'à 70B, équipe de 5.

| Composant | Choix | Prix |
|---|---|---|
| Serveur 8× H100 HGX | Dell/Supermicro | 285 000 $ |
| 2× switch IB 400G + câbles | — | 45 000 $ |
| Stockage NVMe 30 To | — | 8 000 $ |
| Onduleur 20 kVA + install élec | — | 25 000 $ |
| DLC si salle chaude | Option | 30 000 $ |
| **Total** | | **~365 000–395 000 $** |

Ne pas oublier : 8 kW élec + PUE, local ≥ 30 dB d'isolation phonique,
contrat de maintenance NBD.

## 82. Cas pratique n°5 : cluster inférence B200 (scale)

**Besoin** : 405B FP8 en prod, 500 tok/s, multi-tenant.

| Composant | Choix | Prix |
|---|---|---|
| 2 nœuds 8× B200 HGX | OEM | 920 000 $ |
| 4× switch IB 800G | — | 120 000 $ |
| Stockage + réseau mgmt | — | 40 000 $ |
| DLC (CDU + dry cooler 500 kW) | — | 180 000 $ |
| Élec (TGBT, onduleur 2× 100 kVA) | — | 120 000 $ |
| **Total** | | **~1 380 000 $** |

16× B200 = 3 To HBM3e, 576 TFLOPS FP8 pic par nœud. Le DLC et l'élec = 22 % du
budget — **le silicium n'est plus le seul poste**.

## 83. Checklist d'achat (avant de signer)

- [ ] Référence exacte GPU (SXM vs PCIe vs OAM) sur le devis
- [ ] Firmware/vBIOS du constructeur du serveur, pas de la carte
- [ ] NVLink bridges inclus (si NVL) — ligne dédiée
- [ ] RAM système : ≥ 8 Go/GPU (inférence), ≥ 64 Go/GPU (training)
- [ ] Lanes PCIe : block diagram vérifié (x16 électriques par GPU)
- [ ] PSU : N+1, 80 % de charge max, courbe disjoncteur D
- [ ] Refroidissement : inlet max constructeur vs salle réelle en été
- [ ] Réseau : IB/Ethernet dimensionné pour le workload (pas après)
- [ ] Étude électrique lancée (délai 4–9 mois)
- [ ] DCGM/amd-smi prévu dès le jour 1
- [ ] Spare : 1 GPU sur site si ≥ 32 GPU
- [ ] Prix recoupés (3 OEM minimum) + clause de révision si délai > 6 mois

## 84. Checklist de mise en service (jour J)

- [ ] Serrage connecteurs 16-pin vérifié (clic + thermographie à 30 min de charge)
- [ ] `nvidia-smi -q` : ECC enabled, PCIe négocié en x16 gen5/4, P0 atteint en charge
- [ ] `nvidia-smi topo -m` : matrice NUMA documentée
- [ ] Burn-in 24 h (`gpu_burn` ou training synthétique) avant prod
- [ ] DCGM déployé, alertes configurées (temp, Xid, power)
- [ ] Chrony synchronisé sur tous les nœuds (< 10 ms)
- [ ] `ibdiagnet` OK si InfiniBand
- [ ] Image système figée (driver + CUDA + framework versionnés)
- [ ] Procédure de redémarrage documentée (ordre : réseau → stockage → nœuds)
- [ ] Sauvegarde des vBIOS/firmwares d'origine (avant toute MAJ)

## 85. FAQ express

**Q : 4090 ou L40S pour de l'inférence pro ?**
R : L40S : ECC, 48 Go, garantie datacenter, 350 W. La 4090 n'est pertinente qu'en
lab où le budget prime sur la fiabilité.

**Q : H100 d'occasion en 2026, bonne affaire ?**
R : À -40 % du neuf oui, pour de l'inférence. Vérifier les heures (DCGM logs),
le format, et le vendeur. Pas pour du training critique sans garantie.

**Q : AMD ou NVIDIA pour débuter ?**
R : NVIDIA si l'équipe vient de CUDA. AMD si le budget est contraint ET que
quelqu'un maîtrise Linux/ROCm. Le coût du changement d'écosystème se paie en
semaines-homme.

**Q : Combien de GPU pour un chatbot 70B à 100 utilisateurs simultanés ?**
R : 70B FP8 sur 1× H200 ou 1× RTX PRO 6000 (batch 32–64, ~2 000 tok/s agrégés).
Prévoir 2× pour la HA. Au-delà de 500 users : 4–8 GPU.

**Q : Le FP4 est-il « production-ready » ?**
R : Pour chat/RAG/agents : oui avec validation métier. Pour code, maths, juridique :
rester en FP8/BF16 en 2026.

**Q : Faut-il attendre Rubin / MI450 avant d'acheter ?**
R : Non si le besoin est immédiat : Rubin = Q4 2026–2027, allocation hyperscalers
d'abord, prix premium 12–18 mois. Oui si le projet démarre mi-2027+.

## 86. Lexique des unités (pour les BOM)

| Unité | Signification | Exemple |
|---|---|---|
| TFLOPS | 10^12 op/s flottantes | H100 : 1 979 TFLOPS FP8 |
| TOPS | 10^12 op/s entières | 5090 : 3 352 TOPS FP4 sparse |
| PFLOPS | 10^15 op/s | B200 : 9 PFLOPS FP4 |
| To/s, Go/s | Débit mémoire | B200 : 8 To/s |
| kW / MW | Puissance | Nœud B200 : 10,6 kW |
| kWh | Énergie | Nœud B200 : ~93 MWh/an |
| PUE | Ratio énergie totale / IT | 1,1 (liquide) – 1,6 (air) |
| U (rack) | 44,45 mm | Serveur HGX : 8U |
| MT/s | Méga-transferts/s (RAM) | DDR5-4800 |
## 87. Références commerciales exactes (pour les devis)

| GPU | Référence commerciale | Vérifiée le |
|---|---|---|
| NVIDIA L40S 48 Go | 900-2G133-0080-000 (PNY : TCSL40SPCIE-PB) | 27/09/2026 |
| RTX PRO 6000 Blackwell WS | Non communiquée publiquement — demander PNY/Dell | 27/09/2026 |
| RTX PRO 6000 Blackwell Server | Non communiquée publiquement — demander PNY/Dell | 27/09/2026 |
| H100 SXM5 80 Go | 900-21010-0000-000 (module) | 27/09/2026 |
| H100 NVL 188 Go | 900-21010-0030-000 (paire) | 27/09/2026 |
| H200 SXM 141 Go | 900-21010-0080-000 (ordre de grandeur, à confirmer OEM) | 27/09/2026 |
| B200 SXM 192 Go | Référence OEM uniquement (non publique) | 27/09/2026 |
| MI300X OAM | 100-300000010 (AMD, plateforme) | 27/09/2026 |
| Dell L40S | 7WK28 (référence Dell constatée eBay recond.) | 27/09/2026 |

Sans la référence exacte, pas de comparabilité des devis. L'exiger systématiquement.

## 88. B300 / GB300 : la génération intermédiaire (vérifié 27/09/2026)

- **B300** : Blackwell Ultra, 288 Go HBM3e (12-high), 8 To/s, FP4 dense 15 PFLOPS,
  TDP **1 400 W**, NVLink 5.
- **DGX B300** : 8 GPU, 2,1–2,3 To totaux, 11,2 kW, 6× NVSwitch (4,8 To/s).
- **GB300 NVL72** : rack 72 GPU, production de masse depuis mi-2026.
- **Positionnement** : le « vrai » concurrent du MI355X (même TDP 1 400 W).
- **Prix** : location ~9,08 $/h (juillet 2026) ; achat non public — à vérifier.
- **Refroidissement** : liquide fortement recommandé (11,2 kW par nœud).

## 89. MI350P : la carte PCIe AMD pour l'entreprise (vérifié 27/09/2026)

- **Specs** : 144 Go HBM3e, 350 W, PCIe, ~40 % plus rapide que H200 NVL en FP16/FP8
  (chiffre issu d'une base de connaissances tierce — à recouper).
- **Positionnement** : l'équivalent AMD du L40S/RTX PRO 6000 — la carte « raisonnable »
  pour datacenter classique sans UBB.
- **Cas d'usage** : inférence 70B FP8 (144 Go ≥ 90 Go requis), RAG, fine-tuning.
- **Prix** : non public — à vérifier auprès des OEM.

## 90. Comparatif détaillé : H200 vs MI325X vs RTX PRO 6000 (70B)

