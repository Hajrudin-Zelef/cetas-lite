---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-6
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["datacenter", "gpu", "capex", "compute", "containment", "distribution", "hbm", "helios", "mi455x", "nvidia", "nvlink", "rubin"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [693, 843]
sha256: 92bda87b4038c56bf56d85c63c8097803c5935d4d5a5adfb46f67e911171140f
---

# Les GPU datacenter / IA — présent et futur vérifié

| GPU | TDP | Refroidissement en serveur | Rack 8 GPU (GPU seuls) |
|---|---|---|---|
| RTX 4090/5090 | 450/575 W | Air (workstation) | ~4,6 kW — air OK |
| L40S | 350 W | Air passif (flux châssis) | 2,8 kW — air OK |
| RTX PRO 6000 | 600 W | Air (actif WS / passif serveur) | 4,8 kW — air OK |
| H100 SXM | 700 W | Air (DGX/HGX) | 5,6 kW — air OK |
| H200 | 700 W | Air (HGX) | 5,6 kW — air OK |
| B200 | 1 000 W | Air limite / **liquide recommandé** | 8 kW — **DLC recommandé** |
| B300 | 1 400 W | **Liquide** | 11,2 kW — DLC obligatoire |
| MI300X | 750 W | Air (OAM passif) | 6 kW — air OK |
| MI350X | 1 000 W | Air | 8 kW — air limite, DLC conseillé |
| MI355X | 1 400 W | **Liquide** | 11,2 kW — DLC obligatoire |
| MI455X (Helios) | ~1 200–1 500 W | **Liquide** | Rack 72 GPU : **246 kW** (design Schneider) |
| Rubin (NVL72) | n.c. | **Liquide** (dry cooling, sans ventilateurs) | Rack : ordre de 120–200 kW (est.) |

### 26.3. Ce que le liquide change pour l'exploitant

- **PUE** : 1,4–1,6 (air) → **~1,1** (liquide). Sur 1 MW IT, c'est 300–500 kW
  d'économie de refroidissement en continu.
- **Température d'eau** : les designs modernes acceptent 32–45 °C en entrée
  (NVL72 Rubin : 45 °C d'inlet) → **free cooling** possible une grande partie
  de l'année en Europe, sans groupe froid.
- **Contraintes** : boucles eau/glycol, CDU (Coolant Distribution Unit) par rangée,
  détecteurs de fuite, maintenance qualifiée, coût capex +30–50 % sur la salle.

## 27. TDP réels : spec sheet vs terrain

Le TDP est un **plafond de design**, pas la conso moyenne. Mesures typiques relevées
(ordres de grandeur, charges IA réelles) :

| GPU | TDP spec | Inférence (typ.) | Training (typ.) | Idle |
|---|---|---|---|---|
| RTX 4090 | 450 W | 250–350 W | 380–430 W | 20–30 W |
| RTX 5090 | 575 W | 350–480 W | 500–560 W | 25–35 W |
| L40S | 350 W | 180–280 W | 300–340 W | 30–40 W |
| RTX PRO 6000 | 600 W | 350–500 W | 520–590 W | 40–50 W |
| H100 SXM | 700 W | 400–600 W | 620–700 W | 60–80 W |
| H200 | 700 W | 450–650 W | 620–700 W | 60–80 W |
| B200 | 1 000 W | 600–850 W | 900–1 000 W | 100–150 W |
| MI300X | 750 W | 450–650 W | 680–750 W | 70–90 W |

**Règle de dimensionnement électrique** : dimensionner sur **TDP × 1,1** (pics +
CPU + réseau + stockage), jamais sur la moyenne. Un disjoncteur calibré sur la
moyenne disjoncte au premier all-reduce.

## 28. Consommation par config 8×GPU (tout compris)

Hypothèses : GPU au TDP, 2× CPU 350 W, 8× NIC 400G 30 W, ventilateurs/stockage 500 W,
PSU rendement 94 % (pertes ~6 %).

| Config | GPU seuls | Nœud complet (DC) | Au mur (AC) | /an à 0,15 €/kWh (24/7) |
|---|---|---|---|---|
| 8× L40S | 2,8 kW | ~4,5 kW | ~4,8 kW | ~6 300 € |
| 8× RTX PRO 6000 | 4,8 kW | ~6,5 kW | ~6,9 kW | ~9 100 € |
| 8× H100 | 5,6 kW | ~7,5 kW | ~8,0 kW | ~10 500 € |
| 8× H200 | 5,6 kW | ~7,5 kW | ~8,0 kW | ~10 500 € |
| 8× B200 | 8,0 kW | ~10,0 kW | ~10,6 kW | ~13 900 € |
| 8× B300 | 11,2 kW | ~13,5 kW | ~14,4 kW | ~18 900 € |
| 8× MI300X | 6,0 kW | ~8,0 kW | ~8,5 kW | ~11 200 € |
| 8× MI355X | 11,2 kW | ~13,5 kW | ~14,4 kW | ~18 900 € |

Un rack de 4 nœuds B200 = **~42 kW** → hors de portée de l'air classique.
C'est le chiffre à mettre devant la direction avant de signer le bon de commande.

## 29. Énergie, PUE et coût électrique : le lien avec le métier

### 29.1. Formule du coût annuel

```
Coût/an = P_IT(kW) × PUE × 8 760 h × prix_kWh
```

Exemple : nœud 8× B200 (10 kW IT), PUE 1,5 (air), 0,15 €/kWh :
10 × 1,5 × 8 760 × 0,15 = **19 710 €/an** d'électricité par nœud.
En PUE 1,1 (liquide) : **14 454 €/an** — 5 256 € économisés par nœud et par an.

### 29.2. Tableau : PUE par type de refroidissement

| Refroidissement | PUE typique | Inlet air/eau | Densité max |
|---|---|---|---|
| CRAC détente directe | 1,5–1,8 | 22–27 °C air | 10–15 kW/rack |
| CRAH eau glacée | 1,3–1,5 | 22–27 °C air | 15–25 kW/rack |
| Containment + CRAH | 1,25–1,4 | 27–32 °C air | 25–40 kW/rack |
| DLC (plaques froides) | 1,08–1,2 | 32–45 °C eau | 60–120 kW/rack |
| Immersion | 1,03–1,1 | 40–50 °C huile | 100–250 kW/rack |

### 29.3. Alimentation : ce que 8× B200 exigent

- **6× PSU 3 000 W** en N+1 minimum (18 kW installés pour 10,6 kW utiles).
- Arrivées **triphasées 400 V** par rack ; un rack 42 kW = ~63 A par phase.
- **Onduleur** : un nœud 8× B200 + réseau = ~12 kVA → onduleur 20 kVA par nœud
  en N+1, autonomie 10 min pour bascule groupe. (Lien direct avec le guide onduleurs.)
- **Groupe électrogène** : 1,5–2× la puissance IT pour le démarrage et les non-linéaires.
- **Disjoncteurs courbe D** : les PSU à découpage appellent des courants d'enclenchement
  qui font disjoncter les courbes C sous-dimensionnées — classique vu sur site.

### 29.4. Schéma ASCII : chaîne électrique d'un nœud 8× B200

```
 HT/BT ──► TGBT ──► Onduleur 20 kVA ──► PDU rack triphasé 63 A
                                              │
                        ┌─────────────────────┼─────────────────────┐
                        ▼                     ▼                     ▼
                   6× PSU 3 kW           Ventilateurs           Réseau
                   (N+1, 18 kW)         + stockage             (switch IB)
                        │
                        ▼
              Baseboard 8× B200 (8 kW) + 2× CPU (700 W)
                        │
              Total nœud : ~10,6 kW AC au mur
              + refroidissement (PUE) → ~11,7 kW (PUE 1,1) à ~15,9 kW (PUE 1,5)
```
## 30. Pièges terrain (1/3) — 18 erreurs réelles

### Piège n°1 : brancher un GPU SXM dans un slot PCIe (ou l'inverse)

Un H100 SXM **n'est pas** une carte PCIe. Il n'a pas de connecteur PCIe x16 de données,
pas de bracket, pas de sorties. Il se monte sur une baseboard HGX avec un dissipateur
spécifique et un torque précis. Acheter « un H100 » sans préciser SXM vs PCIe, c'est
acheter un presse-papier à 30 000 $. **Toujours exiger la référence exacte sur le devis.**

### Piège n°2 : le 12VHPWR / 12V-2×6 mal enfiché

RTX 4090/5090, L40S, RTX PRO 6000 : un connecteur 16-pin **pas à fond** = résistance de
contact = échauffement = connecteur fondu (cas documentés sur 4090 dès 2022). En
datacenter : vérifier l'enclenchement **à chaque intervention**, utiliser les câbles
d'origine du PSU (pas d'adaptateurs tiers), et préférer les PSU ATX 3.1 natifs 12V-2×6.

### Piège n°3 : PCIe 4.0 vs 5.0 — le goulot invisible

Une RTX PRO 6000 (PCIe 5.0) dans un slot PCIe 4.0 x16 : 64 Go/s au lieu de 128 Go/s.
Pour de l'inférence 1 GPU, indolore. Pour du **chargement de modèle fréquent** ou du
multi-GPU en PCIe sans NVLink, c'est -30 à -50 % de débit effectif. Vérifier la
négociation avec `nvidia-smi -q | grep -i pcie`.

### Piège n°4 : CPU trop faible — famine de données

8× H100 avec 2× CPU 16 cœurs et 256 Go RAM : les GPU attendent les données. Règle :
**8–16 Go RAM système par GPU minimum** pour l'inférence, **64 Go+ par GPU** pour le
training sérieux, et assez de cœurs pour les dataloaders (2–4 cœurs/GPU). Le GPU n'est
que le tiers du BOM.

### Piège n°5 : l'ECC désactivé ou absent

Les GeForce (4090/5090) **n'ont pas d'ECC**. Un bitflip silencieux en training long =
run à jeter après 3 semaines et 50 000 € de compute. En production : L40S minimum
(ECC GDDR6), idéalement HBM ECC (H100+). `nvidia-smi -q` : vérifier « ECC Mode :
Enabled ».

### Piège n°6 : firmware / vBIOS non datacenter

