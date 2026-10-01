---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-12
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "AWS", "CoreWeave", "Lambda", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["gpu", "amd", "arr", "aws", "compute", "fine-tuning", "fp8", "hbm3", "hyperscaler", "nvidia", "qlora", "rubin"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [1608, 1760]
sha256: de44a28395d1dd7479690ebda59d39fcdd1ffe3aefc043b4302a1646af607390
---

# Les GPU datacenter / IA — présent et futur vérifié

NCCL supporte l'hétérogène, mais le training va à la vitesse du plus lent **et**
la mémoire utile est celle du plus petit (80 Go). Un nœud mixte = un nœud H100
au prix d'un nœud H200. Séparer les pools par génération.

### Piège n°29 : oublier la licence logicielle

NVIDIA AI Enterprise (support + vGPU + Base Command) : ~3 500–4 500 $/GPU/an.
VMware vSphere + vGPU : licence par GPU. Ces coûts n'apparaissent pas dans le
BOM matériel mais pèsent 10–15 % du TCO 3 ans. Les chiffrer dès l'étude.

### Piège n°30 : le « on va mettre ça dans le cloud » sans calcul

8× H100 à 3,33 $/GPU/h en 24/7 = **233 k$/an**. Le serveur acheté = 285 k$ une fois
+ 15 k$/an d'élec. Break-even : **~15 mois**. Au-delà, le cloud coûte 2× plus cher.
En deçà de 60 % d'utilisation : le cloud gagne. **Faire le calcul à chaque fois.**

## 68. Fiches d'identité rapide (pense-bête)

```
RTX 4090      : 24 Go GDDR6X | 450 W | PCIe | lab/dev | ~2 400 $
RTX 5090      : 32 Go GDDR7  | 575 W | PCIe | lab/dev | ~2 600 $
L40S          : 48 Go GDDR6  | 350 W | PCIe | inférence RAG | ~10 000 $
RTX PRO 6000  : 96 Go GDDR7  | 600 W | PCIe | 70B Q/FP8 1 GPU | ~12 000 $
H100          : 80 Go HBM3   | 700 W | SXM  | training std | ~30 000 $
H200          : 141 Go HBM3e | 700 W | SXM  | 70B FP16 1 GPU | ~35 000 $
B200          : 192 Go HBM3e |1000 W | SXM  | training/inférence max | ~40 000 $
MI300X        : 192 Go HBM3  | 750 W | OAM  | alternative H100 | devis OEM
MI325X        : 256 Go HBM3e |1000 W | OAM  | alternative H200 | devis OEM
MI350X        : 288 Go HBM3e |1000 W | OAM  | alternative B200 air | devis OEM
MI355X        : 288 Go HBM3e |1400 W | OAM  | alternative B200 liquide | devis OEM
MI300A        : 128 Go unif. | 760 W | APU  | HPC | devis HPC
```

## 69. Compatibilité CPU / carte mère par format

| Format GPU | CPU conseillés | Lanes PCIe requis | Note |
|---|---|---|---|
| 4× PCIe DW (L40S/PRO) | 1–2× EPYC 9004/9005 ou Xeon 6 | 64+ (4× x16) | PCIe 5.0 pour PRO 6000 |
| 8× PCIe DW | 2× EPYC (128 lanes) | 128 (8× x16) | Bifurcation x16/x16 par slot |
| HGX 8× SXM | 2× EPYC/Xeon (plateforme OEM) | 8× x16 vers baseboard | Design OEM uniquement |
| 8× OAM (UBB AMD) | 2× EPYC Turin | PCIe 5.0 x16 mgmt | Design de référence AMD |
| 2× RTX 5090 (lab) | 1× Ryzen 9 / Threadripper | 2× x16 (ou x8/x8) | x8 acceptable en inférence |

**Bifurcation** : une carte mère 8× GPU PCIe a besoin de slots câblés en x16
électriques avec bifurcation — vérifier le block diagram du constructeur,
pas juste le nombre de slots physiques.

## 70. Connecteurs et câbles : l'inventaire

| GPU | Connecteur alim | Câble | PSU min (système 1 GPU) |
|---|---|---|---|
| RTX 4090 | 1× 16-pin (12VHPWR) | 4× 8-pin → 16-pin (fourni) | 850 W (1 000 W conseillé) |
| RTX 5090 | 1× 16-pin (12V-2×6) | Natif ATX 3.1 | 1 000 W |
| L40S | 1× 16-pin | 2–3× 8-pin → 16-pin | 750 W |
| RTX PRO 6000 | 1× 16-pin | Selon OEM | 1 000 W |
| H100 PCIe | 1× 8-pin EPS | EPS 8-pin | 800 W |
| Serveurs HGX/UBB | Barres bus / connecteurs propriétaires | — | 6–8× 3 000 W |

**Règle** : 1 rail 12 V dédié par connecteur 16-pin, jamais de « daisy chain »
sur un seul câble 8-pin pour 450 W+.

## 71. Firmware et outillage constructeur

| Constructeur | Outil firmware | Particularité |
|---|---|---|
| NVIDIA | `nvflash` (réservé), DCGM | vBIOS signé — pas de flash sauvage |
| Dell | iDRAC + catalogues | Firmware GPU via iDRAC Lifecycle |
| HPE | iLO + SPP | SPP trimestriels, tester avant prod |
| Supermicro | SMCIPMITool / Redfish | Le plus permissif, MAJ manuelles |
| AMD | `amd-smi`, ROCm | Firmware via package ROCm |

**Politique** : figer un « firmware train » par trimestre, tester sur 1 nœud
pilote, déployer par vagues. Un flash GPU raté = carte brickée (RMA).

## 72. RMA et support : ce qu'il faut savoir

- **Délais typiques** : NBD (next business day) chez Dell/HPE en contrat ProSupport ;
  2–4 semaines en garantie de base distributeur.
- **Stock tampon** : pour un cluster ≥ 32 GPU, prévoir **1 GPU de spare** sur site
  (un nœud à 7/8 GPU en training = run dégradé ou arrêté).
- **Preuve de panne** : DCGM logs (température, Xid errors) — sans logs, le support
  fait traîner. `nvidia-smi -q -x` + `dmesg | grep -i xid` en pièce jointe au ticket.
- **Xid 79** (GPU has fallen off the bus) : dans 80 % des cas, c'est l'alim ou le
  connecteur, pas le GPU. Vérifier avant de demander un RMA.

## 73. Assurance et risques

- Un nœud 8× B200 (500 k$) s'assure en « tous risques matériel » avec valeur
  agréée — l'assurance standard du bâtiment ne couvre pas ce type d'actif.
- **Clauses** : vol, dégât des eaux (DLC !), surtension (vérifier la coordination
  des parafoudres avec l'onduleur).
- Valoriser aussi le **manque à gagner** : 1 semaine d'arrêt = 5 600 $ de compute
  perdu (8× H100 à 3,33 $/h) + pénalités clients éventuelles.

## 74. Supply chain : délais réels constatés 2026

| Produit | Délai typique | Tendance |
|---|---|---|
| RTX 4090/5090 retail | Immédiat–2 sem. | Volatil (prix +145 % constatés) |
| L40S / RTX PRO 6000 | 2–6 sem. | Tendu sur GDDR7 (PRO 6000 +55 %) |
| Serveur 8× H100 | 8–12 sem. | Détendu (fin de cycle) |
| Serveur 8× H200 | 10–16 sem. | Stable |
| Serveur 8× B200 | 12–20 sem. | Tendu, allocation OEM |
| Plateforme 8× MI300X | 10–16 sem. | Disponible |
| Plateforme 8× MI355X | 12–20 sem. | Allocation |
| Rack Rubin NVL72 | Q4 2026 → 2027 | Réservé hyperscalers d'abord |

## 75. Négociation : 7 leviers

1. **Fin de cycle** : H100 = -15 à -25 % négociables fin 2026.
2. **Volume** : à partir de 32 GPU, les OEM ouvrent des remises « projet ».
3. **Concurrence** : faire chiffrer Dell + HPE + Supermicro sur le même BOM.
4. **Alternative AMD** : même sans acheter AMD, le devis MI300X fait baisser le H100.
5. **Services** : négocier l'installation et 1 an de support avancé offerts.
6. **Reprise** : certains OEM reprennent l'ancien cluster (A100/H100).
7. **Financement** : leasing 36 mois = OPEX, souvent moins cher que l'achat cash
   avec actualisation (le GPU se déprécie vite — le leasing transfère le risque).

## 76. TCO 3 ans comparé : 8× H100 vs 8× B200 vs 8× MI300X

Hypothèses : achat serveur, PUE 1,4, 0,15 €/kWh, utilisation 80 %, maintenance 5 %/an.

| Poste (3 ans) | 8× H100 | 8× B200 | 8× MI300X |
|---|---|---|---|
| Serveur | 285 000 $ | 460 000 $ | ~350 000 $ (est.) |
| Électricité (80 %) | 31 000 € | 41 000 € | 33 000 € |
| Maintenance/support | 43 000 $ | 69 000 $ | 52 000 $ |
| Réseau (amorti) | 40 000 $ | 50 000 $ | 40 000 $ |
| **TCO 3 ans** | **~400 000 $** | **~625 000 $** | **~480 000 $** |
| PFLOPS FP8 | 15,8 | 36,0 | 20,9 |
| **$/PFLOPS/3 ans** | **~25 300** | **~17 400** | **~23 000** |

Le B200 gagne au TCO par unité de compute malgré son prix — **si** l'utilisation
dépasse 60 %. En dessous, le H100 d'occasion ou la location gagnent.
## 77. Location cloud : comparatif providers (27/09/2026)

| Provider | H100 $/h | H200 $/h | B200 $/h | 4090 $/h | Note |
|---|---|---|---|---|---|
| Spheron (marketplace) | 2,64 | 5,55 | 4,38 (spot) | — | Prix live au 27/09/2026 |
| Runpod | ~2,5–3 | ~4–5 | ~5–6 | 0,34 (community) | Large dispo |
| Lambda Labs | ~2,9 | — | — | 0,50 | Pas de L40S |
| CoreWeave | ~3,5–6 | 6,30 | ~6 | — | Réservé entreprise |
| AWS (p5/g6e) | jusqu'à 12,29 | jusqu'à 10,60 | — | — | Hyperscaler premium |
| DigitalOcean | 3,39 (H100) | — | — | — | MI300X à 1,99 |
| Vast.ai (marketplace) | 1,49–3 | 3,72+ | 2,12 (spot) | 0,25–0,40 | Particuliers, variable |

**Stratégie** : dev/test sur Vast.ai/Runpod spot, prod sur réservé 12–36 mois
(-40 à -60 % vs on-demand), burst chez l'hyperscaler.

## 78. Cas pratique n°1 : lab IA 2× RTX 5090 (budget < 10 k$)

**Besoin** : dev, tests, fine-tuning QLoRA, inférence 32B.

