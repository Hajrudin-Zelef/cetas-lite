---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-11
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["gpu", "amd", "capex", "compute", "distribution", "hbm", "incident", "nvidia", "rubin", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [1438, 1607]
sha256: 63258d781ff618b5778487494f1e3f812450147bc5917dd83257844ba386dabb
---

# Les GPU datacenter / IA — présent et futur vérifié

- **Dimensionnement** : 1,5–2× la puissance IT pour absorber les appels de courant
  des PSU et la non-linéarité (harmoniques des alimentations à découpage).
- **Exemple** : 37 kW IT → groupe **75–100 kVA**.
- **Démarrage** : les GPU ne redémarrent pas seuls proprement après coupure —
  prévoir une **procédure de remise en service** (ordre : réseau → stockage →
  nœuds, vérification nvidia-smi, rechargement des modèles).
- **Test mensuel** : essai en charge du groupe (banc de charge si pas de délestage
  possible) — un groupe qui ne démarre pas le jour J est un classique.

## 60. Refroidissement DLC : design pratique

### 60.1. Architecture type

```
Refroidisseur sec (dry cooler) / groupe froid
        │ eau glycolée 32–40 °C
        ▼
CDU (Coolant Distribution Unit) — 1 par rangée, 500–1 000 kW
        │ boucle technologique (eau traitée)
        ▼
Manifold du rack → plaques froides GPU/CPU → retour chaud 40–50 °C
```

### 60.2. Chiffres clés

| Paramètre | Valeur typique |
|---|---|
| ΔT eau entrée/sortie | 8–12 °C |
| Débit par rack 100 kW | ~150–200 L/min |
| Pression boucle techno | 2–4 bar |
| Qualité d'eau | Déminéralisée, inhibiteurs anti-corrosion |
| Détection fuite | Câble sensible par rack + vannes d'isolement auto |

### 60.3. Coûts (ordres de grandeur 2026)

- CDU 500 kW : ~80 000–120 000 $.
- Dry cooler 500 kW : ~60 000–100 000 $.
- Rétrofit d'une salle existante : +200–400 $/kW IT.
- Neuf avec DLC natif : +10–20 % vs air sur le capex salle.

## 61. Monitoring énergétique et thermique

### 61.1. Ce qu'il faut mesurer

| Mesure | Outil | Seuil d'alerte |
|---|---|---|
| Puissance GPU | `nvidia-smi dmon` / DCGM | > 95 % TDP sustained |
| Température GPU | DCGM `dcgm_field` | > 85 °C (HBM throttle ~95–100 °C) |
| Température eau DLC | Sonde CDU | > 45 °C entrée |
| PUE instantané | Compteurs TGBT vs IT | Dérive > 0,1 vs baseline |
| THD / facteur de puissance | Analyseur réseau | THD > 8 % |

### 61.2. DCGM (Data Center GPU Manager)

- Le standard NVIDIA : collecte 300+ métriques par GPU, export Prometheus.
- **À déployer dès le jour 1** sur tout nœud H100+ : c'est le seul moyen de prouver
  un RMA (courbes de température/puissance avant panne).
- AMD : `amd-smi` + exporters ROCm (moins riche, mais suffisant).

## 62. Plan de maintenance préventive (salle GPU)

| Fréquence | Action |
|---|---|
| Hebdomadaire | Relecture alertes DCGM/amd-smi, contrôle visuel voyants |
| Mensuelle | Filtres à air, test groupe (si applicable), revue PUE |
| Trimestrielle | Serrage connecteurs 16-pin (thermographie), firmware/drivers |
| Semestrielle | Thermographie complète (VRM, connecteurs, PDU), test onduleur en charge |
| Annuelle | Remplacement préventif ventilateurs > 3 ans, analyse batteries, exercice PRA |

## 63. PRA/PCA pour cluster IA

- **RPO** : checkpoints toutes les 2 h → RPO 2 h de training (~1 800 $ de compute
  perdus par incident sur 8× H100).
- **RTO** : 4 h (redémarrage nœuds + rechargement dataset + reprise checkpoint).
- **Sauvegarde** : checkpoints sur stockage répliqué (pas sur le NVMe local seul —
  un nœud mort = run mort).
- **Test** : exercice semestriel de reprise sur 1 nœud (pas besoin de tout casser).

## 64. Sécurité physique et incendie

- **Batteries Li-ion** (onduleurs) : local coupe-feu 2 h, détection gaz, pas d'eau
  en extinction (poudre/classe D ou gaz).
- **DLC** : bac de rétention sous CDU, vanne d'isolement automatique sur détection.
- **Accès** : un nœud 8× B200 = 500 k$ — contrôle d'accès biométrique, vidéosurveillance,
  traçabilité des interventions (qui a ouvert quel capot, quand).

## 65. Bruit : le paramètre oublié

- Un nœud 8× H100 air : **~85–95 dB** à 1 m (ventilateurs 80 mm à fond).
- Impossible en open space / bureau — local technique dédié obligatoire.
- Le DLC divise le bruit par 2 à 3 (plus de ventilateurs GPU, ventilation réduite).
- Réglementation : 85 dB = protection auditive obligatoire pour les intervenants.

## 66. Tableau récapitulatif énergie par GPU (fiche Zelef)

| GPU | TDP | Nœud 8× (AC) | PUE 1,4 → total | Coût/an (0,15 €/kWh) | Refroidissement |
|---|---|---|---|---|---|
| L40S | 350 W | 4,8 kW | 6,7 kW | 8 800 € | Air |
| RTX PRO 6000 | 600 W | 6,9 kW | 9,7 kW | 12 700 € | Air |
| H100 | 700 W | 8,0 kW | 11,2 kW | 14 700 € | Air |
| H200 | 700 W | 8,0 kW | 11,2 kW | 14 700 € | Air |
| B200 | 1 000 W | 10,6 kW | 14,8 kW | 19 500 € | DLC conseillé |
| MI300X | 750 W | 8,5 kW | 11,9 kW | 15 600 € | Air |
| MI355X | 1 400 W | 14,4 kW | 20,2 kW | 26 500 € | DLC obligatoire |

**Message à la direction** : sur 5 ans, l'électricité d'un nœud 8× B200 (97 k€)
représente ~20 % du prix du serveur. Le choix du refroidissement (PUE 1,1 vs 1,5)
pèse ±35 k€ par nœud sur 5 ans — plus que l'écart de prix entre deux OEM.
## 67. Pièges terrain (4/4) — suite

### Piège n°19 : le driver « latest » en production

Un driver NVIDIA flambant neuf peut casser NCCL ou un kernel custom.
**Règle** : driver validé = celui de la matrice NGC/CUDA du framework déployé,
pas le dernier du site. Figer les versions (driver + CUDA + framework) dans
l'image et ne mettre à jour que sur fenêtre de maintenance testée.

### Piège n°20 : cgroup et partage GPU non isolé

Deux conteneurs sur le même GPU sans MPS/MIG : un OOM tue les deux workloads.
En multi-tenant : **MIG** (isolation matérielle, RTX PRO/H100+) ou **MPS**
(partage logiciel, perf > isolation). Ne jamais faire du « best effort » partagé
en prod facturée.

### Piège n°21 : le Secure Boot qui bloque le driver

Driver NVIDIA DKMS + Secure Boot sans clé MOK enrôlée = driver refusé au boot =
nœud sans GPU après un reboot planifié un dimanche à 3 h. Enrôler la clé MOK
**avant** la mise en production, ou désactiver Secure Boot (documenté, assumé).

### Piège n°22 : horloges et NCCL — le NTP oublié

Un décalage > 500 ms entre nœuds : les timeouts NCCL deviennent aléatoires,
les logs sont inexploitables. **Chrony sur tous les nœuds**, même source de temps,
vérifié par monitoring. Ça coûte 0 € et ça évite des semaines de debug.

### Piège n°23 : câbles IB cuivre trop longs

DAC cuivre 400G : 2–3 m max. Au-delà : fibre ou AEC (Active Electrical Cable).
Un lien IB qui flappe = training qui passe de 40 % à 5 % MFU sans erreur claire.
Toujours valider avec `ibdiagnet` après câblage.

### Piège n°24 : le « B200 air » dans une salle à 30 °C

Le B200 HGX air-cooled existe, mais il exige ≤ 27 °C en entrée et un débit d'air
massif. Dans une salle tropicale ou un shelter : **throttling garanti**.
Vérifier la spec d'inlet du constructeur **et** la température réelle de la salle
en été avant de signer « air ».

### Piège n°25 : acheter des GPU avant le réseau électrique

Délai typique : GPU livrés en 8–12 semaines, **raccordement électrique renforcé :
4–9 mois** (Enedis/GRD + TGBT + onduleur). Commander le GPU sans avoir lancé
l'étude électrique, c'est payer un serveur qui dort. Ordre : étude élec → commande
parallèle GPU + élec → réception.

### Piège n°26 : la revente — l'obsolescence programmée du marché

Un H100 acheté 30 k$ en 2024 vaut ~15–18 k$ fin 2026 (arrivée B200/Rubin).
Amortissement comptable : **3 ans max** pour du GPU IA. Ne pas mettre du GPU
dans un plan d'amortissement 5 ans « serveurs généralistes ».

### Piège n°27 : le support logiciel a une fin

CUDA abandonne les vieilles architectures : Kepler/Maxwell déjà sortis du support
récent. Un cluster acheté aujourd'hui doit être amorti **avant** la fin du support
driver (~8 ans NVIDIA, mais le support *optimisé* dure ~4–5 ans).

### Piège n°28 : mélanger H100 et H200 dans un même run

