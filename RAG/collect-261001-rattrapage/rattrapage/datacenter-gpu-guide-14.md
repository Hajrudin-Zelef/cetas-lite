---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-14
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Broadcom", "Nvidia", "SGLang", "vLLM"]
dates: ["2026-09-27"]
keywords: ["gpu", "amd", "blackwell", "diffusion", "distribution", "embedding", "embeddings", "ethernet", "fp8", "hbm4", "helios", "mistral"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [1940, 2119]
sha256: b11253575feeed009c0ee0af0306810fc9df7887a742378e66fb7a2d4e809046
---

# Les GPU datacenter / IA — présent et futur vérifié

| Critère | H200 | MI325X | RTX PRO 6000 |
|---|---|---|---|
| VRAM | 141 Go | 256 Go | 96 Go |
| 70B FP16 1 GPU | Oui (161 Go : non — 141 < 161 !) | Oui | Non |
| 70B FP8 1 GPU | Oui | Oui | Oui (90 Go) |
| Contexte 128K (70B FP8) | Non (240 Go requis) | Oui (256 Go) | Non |
| NVLink | Oui (NVL) | xGMI (plateforme) | Non |
| Écosystème | CUDA mature | ROCm (bon en 2026) | CUDA mature |
| Prix indicatif | ~35 000 $ | Devis OEM | ~12 000 $ |
| Verdict | Le standard | Le max VRAM/1 GPU | Le $/Go imbattable |

Note : le 70B FP16 + overhead (161 Go) ne tient PAS sur un H200 seul (141 Go) —
il faut le NVL (282 Go) ou passer en FP8. Détail qui piège beaucoup de dimensionnements.

## 91. Dimensionnement multi-nœuds : le réseau d'abord

### 91.1. Débits requis par workload

| Workload | Inter-nœuds requis | Solution |
|---|---|---|
| Inférence DP (requêtes indépendantes) | 25–100 GbE | Ethernet standard |
| Training 8–64 GPU | 400G IB ou Ethernet | InfiniBand NDR / UEC |
| Training 64+ GPU | 800G IB | InfiniBand XDR / Spectrum-6 |
| MoE géant (EP inter-nœuds) | 800G+ | IB XDR, faible latence |

### 91.2. Topologies

- **Fat-tree** : le standard (2 niveaux pour ≤ 128 GPU, 3 niveaux au-delà).
- **Dragonfly** : hyperscalers, latence optimisée.
- **Règle** : oversubscription ≤ 3:1 en training, 1:1 (non-bloquant) pour les
  collectives critiques.

### 91.3. Coût réseau par nœud (2026)

| Composant | Prix unitaire | Par nœud 8 GPU |
|---|---|---|
| NIC ConnectX-7 400G | ~1 500 $ | 12 000 $ |
| NIC ConnectX-8 800G | ~2 500 $ (est.) | 20 000 $ |
| Switch IB 64× 400G | ~40 000 $ | ~10 000 $ (amorti/4 nœuds) |
| Câbles DAC/fibre | 100–500 $ | 4 000–8 000 $ |
| **Total** | | **25 000–40 000 $** |

## 92. Ethernet vs InfiniBand en 2026 : le match se resserre

| Critère | InfiniBand (NDR/XDR) | Ethernet (UEC / Spectrum-X) |
|---|---|---|
| Latence | ~1 µs | ~2–3 µs (en baisse) |
| RDMA | Natif | RoCEv2 / UEC |
| Écosystème IA | NCCL optimisé depuis 10 ans | Rattrapage rapide (2025-2026) |
| Prix | Premium +20–30 % | Moins cher, équipes déjà formées |
| Choix AMD Helios | — | **UEC natif** (pari stratégique) |
| Choix NVIDIA | Natif (ConnectX) | Spectrum-X en alternative |

**Tendance** : l'Ultra Ethernet Consortium (AMD, Broadcom, Cisco…) rend l'Ethernet
crédible pour le training. Pour un nouveau cluster 2026-2027 sans legacy IB,
l'Ethernet UEC mérite une vraie évaluation — les équipes réseau le maîtrisent déjà.

## 93. Sécurité : firmware, supply chain, isolation

- **Firmware signé** : NVIDIA signe ses vBIOS (pas de flash custom) ; AMD via ROCm.
  Vérifier les signatures après chaque MAJ (`nvidia-smi --query` + catalogue OEM).
- **Supply chain** : acheter via distribution agréée — les lots gris peuvent contenir
  des cartes reconditionnées vendues comme neuves (heures DCGM non remises à zéro :
  vérifier `nvidia-smi -q | grep -i "retired"`).
- **Isolation multi-tenant** : MIG (matériel) > vGPU (hyperviseur) > MPS (logiciel) >
  partage best-effort (à proscrire en facturation).
- **Effacement avant revente** : les poids de modèles propriétaires restent en VRAM
  après usage — `nvidia-smi --gpu-reset` + réécriture mémoire avant toute sortie
  de carte (RMA, revente, recyclage).

## 94. Fin de vie et recyclage

- **Durée de vie utile** : 3–4 ans en training intensif (obsolescence économique),
  5–6 ans en inférence légère.
- **Seconde vie** : H100 → inférence (encore excellent), A100 → labs/universités,
  cartes gaming → marché secondaire.
- **Recyclage** : les GPU contiennent or, cuivre, terres rares (aimants ventilateurs).
  Filières DEEE pro — ne pas mettre un HGX à la benne.
- **Données** : effacement certifié avant recyclage (cf. §93).

## 95. Jalons 2024-2026 : comment on en est arrivé là

| Date | Événement |
|---|---|
| 03/2024 | NVIDIA annonce Blackwell (B200) au GTC |
| 06/2024 | AMD annonce MI325X (256 Go) |
| 10/2024 | AMD annonce MI350X/MI355X (CDNA 4, 288 Go) |
| 01/2025 | RTX 5090 (32 Go GDDR7) et RTX PRO 6000 annoncées au CES |
| 03/2025 | NVIDIA : Rubin détaillé (roadmap GTC 2025) |
| 06/2025 | DigitalOcean : MI300X à 1,99 $/h (choc tarifaire) |
| 02/2026 | MLPerf v6.0 : MI355X à 92–104 % du B300 (chiffres AMD) |
| 03/2026 | GTC 2026 : Vera Rubin en pleine production (7 puces, 5 racks) |
| 05/2026 | GTC Taipei : Rubin ramp-up confirmé, Supermicro en production |
| 07/2026 | GB300 NVL72 en production de masse |
| 08/2026 | AMD Advancing AI : MI450/Helios lancés, livraisons fin Q3 |
| 09/2026 | Pénurie GDDR7 : RTX PRO 6000 +55 %, RTX 5090 +145 % vs MSRP |

## 96. Ce que 2027 devrait apporter (officialisé uniquement)

- **NVIDIA** : ramp-up Rubin (Q4 2026 → 2027), puis **Feynman** (nom officialisé,
  specs non publiées au 27/09/2026).
- **AMD** : ramp-up MI450/Helios (fin 2026 → 2027), **MI500 en 2027** (annoncé,
  specs non publiées), EPYC Venice H2 2026.
- **Tendances structurelles** : HBM4 généralisé, TDP 1 000 W+ devenant la norme,
  DLC obligatoire au-delà de 60 kW/rack, Ethernet UEC en alternative crédible à IB.
## 97. Dimensionnement par workload (hors LLM texte)

### 97.1. Embeddings

| Modèle | VRAM (batch 512) | Débit/L40S | GPU conseillé |
|---|---|---|---|
| BGE-large (335M) | ~4 Go | ~3 000 req/s | L40S ou 4090 |
| E5-mistral 7B | ~16 Go | ~800 req/s | L40S |
| BGE-M3 (multilingue) | ~6 Go | ~2 000 req/s | L40S |

L'embedding est le workload le plus rentable : un seul L40S absorbe des milliers
de requêtes/s. Ne jamais mettre un H100 là-dessus (gaspillage ×3).

### 97.2. Vision (détection, segmentation)

| Modèle | VRAM | GPU conseillé |
|---|---|---|
| YOLOv8-x (inférence) | ~4 Go | L40S / 4090 |
| SAM 2 (segmentation) | ~8 Go | L40S |
| DINOv2-giant (features) | ~6 Go | L40S |
| Diffusion (SDXL, batch 4) | ~12–16 Go | 4090/5090 (rapide) ou L40S (stable) |

### 97.3. Audio / parole

| Modèle | VRAM | Note |
|---|---|---|
| Whisper large-v3 | ~6 Go | Temps réel sur L40S (×10) |
| TTS (XTTS, Bark) | ~4–8 Go | 4090 suffisant |
| Diarisation (pyannote) | ~4 Go | CPU possible |

### 97.4. Règle générale

```
Si modèle < 10B params et batchable → L40S (ou 4090 en lab).
Si 10–70B → RTX PRO 6000 / H200 / MI300X selon budget.
Si > 70B ou training → H100/H200/B200 / MI355X.
```

## 98. Template de devis type (à envoyer aux OEM)

```
OBJET : Demande de devis — Serveur IA 8× GPU
Date : __/__/____

1. GPU : 8× [référence exacte + SXM/PCIe/OAM], [VRAM] Go, TDP [___] W
2. Baseboard : [HGX / UBB / PCIe] + [NVSwitch / bridges NVLink si applicable]
3. CPU : 2× [modèle], RAM : [___] Go DDR5 ECC
4. Stockage : [___] To NVMe
5. Réseau : [___]× [ConnectX-7/8 ___G] + câbles [DAC/fibre, longueur]
6. Alimentation : [___]× [___] W, redondance N+1
7. Refroidissement : [air / DLC] — inlet max garanti : [___] °C
8. Firmware : version [___], catalogue OEM [___]
9. Garantie : [___] ans, [NBD / 4h], sur site
10. Services : installation, mise en service, burn-in 24 h
11. Délai : [___] semaines, clause de révision si > [___] mois
12. Prix : unitaire détaillé par ligne (pas de forfait global)
```

Un devis sans les lignes 1, 5 et 12 détaillées = devis à rejeter.

## 99. Organisation : qui fait quoi autour d'un cluster GPU

| Rôle | Responsabilités | Charge (cluster 32 GPU) |
|---|---|---|
| Admin système GPU | Drivers, DCGM, firmware, burn-in | 0,5 ETP |
| Ingénieur réseau | IB/Ethernet, NCCL, stockage | 0,3 ETP |
| Électricien / énergéticien | TGBT, onduleur, DLC, PUE | 0,2 ETP (+ astreinte) |
| MLOps | vLLM/SGLang, modèles, monitoring métier | 1 ETP |
| Data engineer | Datasets, checkpoints, stockage | 0,5 ETP |

**Total** : ~2,5 ETP pour 32 GPU. En dessous, c'est le mode « pompier ».
Le poste énergie (0,2 ETP) est celui qu'on oublie — jusqu'à la première disjonction.

## 100. Plan de montée en compétence (8 semaines)

