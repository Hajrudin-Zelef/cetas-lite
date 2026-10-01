---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-15
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia", "vLLM"]
dates: []
keywords: ["datacenter", "gpu", "amd", "benchmarks", "fine-tuning", "hbm", "helios", "kv cache", "lora", "memory", "nvidia", "nvlink"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [2120, 2273]
sha256: 7adc874cc14ffe8e59aed989780b51253374fea0d53068f0cb96fae32561aa1c
---

# Les GPU datacenter / IA — présent et futur vérifié

| Semaine | Sujet | Livrable |
|---|---|---|
| 1 | Bases GPU : archi, mémoire, TDP, formats | Fiche mémo équipe |
| 2 | Linux GPU : drivers, nvidia-smi, DCGM | Nœud pilote monitoré |
| 3 | Inférence : vLLM, quantification, KV cache | 8B servi en local |
| 4 | Dimensionnement : calculs VRAM, batch | Dimensionnement validé d'un cas réel |
| 5 | Réseau : NCCL, IB, topologies | Test NCCL all-reduce documenté |
| 6 | Énergie : PUE, onduleur, DLC | Bilan électrique de la salle |
| 7 | Training : DP/TP, checkpoints, MFU | Fine-tuning LoRA complet |
| 8 | Prod : HA, PRA, RMA, runbook | Runbook d'exploitation v1 |

## 101. Free cooling : le gisement d'économies

- **Principe** : quand l'air extérieur < température de consigne, on refroidit sans
  groupe froid (économiseur à air ou à eau).
- **Avec DLC à 40 °C d'inlet** (standard Rubin/Helios) : free cooling possible
  **> 8 000 h/an** en climat tempéré européen — le groupe froid ne tourne presque jamais.
- **Calcul** : 100 kW IT, PUE 1,5 → 1,1 = 40 kW économisés × 8 000 h × 0,15 € =
  **48 000 €/an** par tranche de 100 kW.
- **Condition** : dry cooler dimensionné pour le ΔT le plus défavorable (été canicule),
  sinon le groupe froid de secours doit prendre le relais — le surdimensionner de 20 %.

## 102. Récupération de chaleur : le DLC comme source

- Un rack 100 kW en DLC sort de l'eau à **40–50 °C** — utilisable directement pour :
  - préchauffage ECS (bureaux),
  - chauffage basse température (plancher chauffant, 35 °C),
  - serres / piscines municipales (partenariats territoriaux).
- **Ordre de grandeur** : 100 kW IT × 8 760 h × 80 % récupérable = **700 MWh/an**
  thermiques — de quoi chauffer ~50 logements.
- **Modèle économique** : revente à un réseau de chaleur (20–40 €/MWh) = 14–28 k€/an
  par 100 kW — finance une partie du surcoût DLC.
- **Contrainte** : la chaleur fatale est disponible quand le datacenter tourne —
  besoin d'un stockage tampon (ballon 5 000–20 000 L) pour lisser jour/nuit.

## 103. Eau et DLC : qualité et traitement

| Paramètre | Boucle technologique | Boucle primaire (dry cooler) |
|---|---|---|
| Fluide | Eau déminéralisée + inhibiteurs | Eau + glycol 30 % |
| Conductivité | < 10 µS/cm | < 100 µS/cm |
| pH | 7–8,5 | 7–9 |
| Traitement | Filtration 50 µm, anti-légionelles inutile (circuit fermé) | Anti-gel, anti-corrosion |
| Appoint | Faible (circuit fermé) | Faible |

**Légionellose** : circuit fermé = pas de risque (pas d'aérosols). Le risque existe
sur les tours ouvertes — à éviter en 2026, préférer dry cooler ou adiabatique.

## 104. Bruit et implantation : règles pratiques

| Distance | Niveau (nœud 8× H100 air) | Usage possible |
|---|---|---|
| 1 m | 90 dB | Local technique uniquement |
| 5 m (porte fermée) | 65 dB | Couloir technique |
| 10 m (mur + porte) | 50 dB | Bureaux (limite) |
| Dry cooler extérieur | 70 dB à 3 m | Étude acoustique si voisinage < 50 m |

Prévoir l'étude acoustique **avant** le permis — un dry cooler de 500 kW à 20 m
d'habitations = recours des voisins garanti sans écran acoustique.
## 105. Cheat sheet nvidia-smi (les 20 commandes utiles)

```bash
nvidia-smi                          # état rapide : util %, temp, mémoire, process
nvidia-smi -q                       # tout le détail (ECC, PCIe, clocks)
nvidia-smi -q -x                    # idem en XML (pour les tickets support)
nvidia-smi topo -m                  # matrice NUMA / NVLink / PCIe
nvidia-smi dmon -s pucvmet -d 5     # monitoring temps réel (power, util, clocks, temp)
nvidia-smi --query-gpu=index,name,temperature.gpu,power.draw,utilization.gpu,memory.used \
  --format=csv -l 5                 # CSV pour scripts
nvidia-smi -pm 1                     # persistence mode (obligatoire en serveur)
nvidia-smi -lgc 1500,2100           # verrouille les clocks (reproductibilité bench)
nvidia-smi -rgc                     # déverrouille
nvidia-smi -pl 500                  # limite la puissance à 500 W (cap)
nvidia-smi --gpu-reset -i 0         # reset GPU 0 (après Xid, si possible)
nvidia-smi mig -lgip                # profils MIG disponibles
nvidia-smi mig -cgi 19 -C           # crée une instance MIG (profil 19)
nvidia-smi -q | grep -i "ecc mode"  # vérifie l'ECC
nvidia-smi -q | grep -i "pcie"      # vérifie la négociation PCIe
dmesg | grep -i xid                 # erreurs GPU côté kernel
nvidia-debugdump -d                 # dump complet pour le support
```

## 106. Cheat sheet amd-smi / ROCm

```bash
amd-smi metric --all               # métriques : temp, power, util, mémoire
amd-smi topology                   # topologie xGMI / PCIe
rocminfo                           # détail devices + gfx target (gfx942/gfx950)
rocm-smi --showmeminfo vram        # mémoire HBM utilisée
rocm-smi --setpoweroverdrive 700    # cap de puissance (W)
```

## 107. DCGM : les compteurs à surveiller

| Field ID | Métrique | Alerte |
|---|---|---|
| 155 | Température GPU | > 85 °C |
| 156 | Power draw | > 95 % TDP |
| 203 | SM utilization | < 20 % en training = anomalie |
| 204 | Memory utilization | — |
| 210 | PCIe throughput | Saturé = goulot host |
| 1001–1008 | NVLink throughput par lien | Déséquilibre = mauvais placement TP |
| Xid errors | Erreurs matérielles | Tout Xid 79/119 = investiguer |

Exporter vers Prometheus (`dcgm-exporter`) + Grafana : le dashboard de base de
toute salle GPU sérieuse.

## 108. Diagnostic express : l'arbre de décision

```
GPU lent ?
├── nvidia-smi : util GPU < 50 % en training ?
│   ├── OUI → problème d'alimentation données : vérifier I/O disque, réseau, dataloader
│   └── NON → vérifier clocks (throttling thermique ? power cap ?)
├── Temp > 85 °C ?
│   ├── OUI → airflow, filtres, pâte thermique (si > 3 ans), DLC
│   └── NON → clocks verrouillés ? driver ?
├── Xid dans dmesg ?
│   ├── 79 (fallen off the bus) → alim / connecteur 16-pin / reseat
│   ├── 119 (GSP error) → driver / firmware
│   └── 13 (graphics fault) → souvent logiciel (kernel CUDA)
└── Multi-GPU lent ?
    ├── NCCL timeout → réseau (ibdiagnet), NTP, firewall
    └── Un GPU plus lent → isoler (test 1 GPU), RMA si confirmé
```

## 109. Reproductibilité des benchmarks

- Verrouiller les clocks (`-lgc`), fixer `CUDA_VISIBLE_DEVICES`, seed PyTorch,
  batch fixe, 3 runs (écarter le 1er = warmup).
- Noter : driver, CUDA, cuDNN, PyTorch/vLLM, version du modèle **et** de la
  quantification. Un bench sans ces 6 infos est inutilisable dans 3 mois.
- Template de fiche bench : GPU / driver / framework / modèle / précision /
  batch / ctx in-out / tok/s / W mesurés / date.

## 110. Sécurité des modèles et des données

- Les poids chargés en VRAM survivent au process (pas au reset GPU).
- En multi-tenant : MIG + reset GPU entre clients (`nvidia-smi --gpu-reset`).
- Les prompts clients transitent par le KV cache — même règle : isolation par
  instance, pas de partage best-effort sur données sensibles.
- Chiffrement : modèles au repos (LUKS sur le NVMe), TLS pour l'API d'inférence,
  pas de logs de prompts en clair (ou avec purge ≤ 30 j et consentement).

## 111. Contrats de maintenance : quoi exiger

- **Garantie** : 3 ans min, NBD sur site pour la prod (4 h si cluster critique).
- **SLO pièces** : GPU de remplacement sous 5 jours ouvrés (contractualisé).
- **Firmware** : catalogue validé OEM, pas de MAJ forcée.
- **Support logiciel** : accès au support NVIDIA AI Enterprise ou équivalent AMD.
- **Pénalités** : à négocier au-delà de 32 GPU (leverage réel).

## 112. Financement : comparatif rapide

