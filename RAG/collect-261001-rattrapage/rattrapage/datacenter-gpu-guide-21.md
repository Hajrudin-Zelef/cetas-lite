---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-21
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["datacenter", "gpu", "arr", "embeddings", "fine-tuning", "helios", "incident", "lora", "nvidia", "qlora", "rubin", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [3092, 3270]
sha256: 23779a13d416558ff1ffba16e4636319377e2e9b328fb48673441b0c28610499
---

# Les GPU datacenter / IA — présent et futur vérifié

**Paradoxe** : le serverless néocloud est moins cher que le self-hosted B200
amorti — parce que leur utilisation est > 90 % (mutualisation). Le self-hosted
ne gagne que si l'utilisation dépasse 70 % **et** que les données doivent rester
sur site.

## 166. Souveraineté : l'argument non-financier

- Données de santé, défense, secteur public : l'on-premise n'est pas un choix
  économique mais réglementaire (HDS, SecNumCloud, ITAR…).
- Dans ce cas, le TCO se compare au **coût de la non-conformité**, pas au cloud.
- Dimensionner au besoin réel + 30 % (les projets souverains grossissent vite).

## 167. Le marché de l'occasion : guide d'achat

| Source | Fiabilité | Prix vs neuf | Vérifications |
|---|---|---|---|
| OEM reconditionné (Dell Outlet) | Haute | -20 à -30 % | Garantie incluse |
| eBay pro (vendeurs 99 %+ ) | Moyenne | -30 à -50 % | Heures, Device ID, photos |
| Jawa / forums | Moyenne | -30 à -40 % | Idem + protection acheteur |
| Lots « faillite datacenter » | Variable | -50 % et plus | Burn-in obligatoire |

**Checklist occasion** : `nvidia-smi -q` (heures via DCGM si dispo), test 24 h,
facture d'origine, pas de vBIOS modifié (comparer le Device ID).

## 168. Garantie et heures de fonctionnement

- Un GPU n'a pas de « compteur horaire » officiel accessible (contrairement aux
  disques). Les logs DCGM peuvent révéler l'historique thermique.
- **Signes d'usure** : ventilateurs bruyants (paliers), pâte thermique sèche
  (> 4 ans), connecteurs oxydés (ambiance humide).
- En occasion datacenter, exiger le **rapport de burn-in** du vendeur.

## 169. Tendances prix 2027 (projection prudente)

- H100 : poursuite de la baisse (-10 à -20 %/an) — fin de vie logicielle ~2028.
- H200 : stable (la mémoire reste demandée pour l'inférence).
- B200 : -15 % avec l'arrivée massive de Rubin.
- RTX PRO 6000 : baisse si la pénurie GDDR7 se résorbe.
- MI300X : pression à la baisse (MI355X/MI450 disponibles).
- **Incertitude majeure** : géopolitique (restrictions d'export) et énergie
  (prix de l'électricité) — les deux font plus bouger les prix que la techno.
## 170. Stack de monitoring complète (référence)

```
GPU ──► DCGM-exporter ──┐
Système ──► node-exporter ─► Prometheus ──► Grafana ──► Alertmanager ──► PagerDuty/mail
Réseau ──► snmp-exporter ──┘         │
Elec ──► compteurs Modbus ───────────┘
```

| Couche | Métriques clés | Rétention |
|---|---|---|
| GPU | Temp, power, util, Xid, clocks | 90 j (1 an pour RMA) |
| Système | CPU, RAM, disque, NTP offset | 30 j |
| Réseau | Débit IB, erreurs, latence | 30 j |
| Élec | kW, PUE, THD | 3 ans (facturation/optimisation) |
| Métier | tok/s, latence p99, erreurs 5xx | 1 an |

**Alertes critiques** : temp GPU > 87 °C, Xid quelconque, PUE dérive > 0,15,
NTP offset > 100 ms, onduleur sur batterie.

## 171. Runbook modèle : incident GPU (1 page)

```
INCIDENT : GPU en erreur / perfs dégradées
1. Identifier : nvidia-smi → quel GPU, dmesg | grep Xid
2. Isoler : drainer Slurm (scontrol update nodename state=drain)
3. Diagnostiquer (arbre §108) : thermique ? alim ? réseau ? logiciel ?
4. Si Xid 79 → vérifier connecteur 16-pin + reseat (fenêtre de maintenance)
5. Si thermique → vérifier airflow/filtres/DLC, DCGM history
6. Si logiciel → rollback driver/framework (image versionnée)
7. Si matériel confirmé → ticket OEM (joindre nvidia-smi -q -x + dmesg)
8. Spare : remplacer par le GPU de stock si ≥ 24 h d'immobilisation
9. Post-mortem : cause racine + action corrective dans le runbook
ESCALADE : > 4 h d'arrêt cluster → astreinte N+1 → direction
```

## 172. Tableau de bord direction (mensuel, 1 page)

| Indicateur | Cible | Ce mois |
|---|---|---|
| Disponibilité GPU | > 99,5 % | ___ |
| Utilisation moyenne | 60–80 % | ___ |
| PUE | < 1,3 | ___ |
| Coût/M tokens (si inférence) | < objectif | ___ |
| Incidents / MTTR | < 2 / < 4 h | ___ |
| Électricité (€) | < budget | ___ |
| RMA en cours | 0 | ___ |

Un tableau que la direction lit en 2 minutes — c'est lui qui débloque les
budgets d'extension.

## 173. Décisions rapides : l'arbre en 1 page

```
Besoin ?
├── Lab / dev perso ──► RTX 5090 (32 Go, ~2 600 $)
├── Inférence < 32B ──► L40S (48 Go ECC, ~10 000 $)
├── Inférence 70B éco ──► RTX PRO 6000 (96 Go, ~12 000 $)
├── Inférence 70B perf ──► H200 (141 Go) ou MI300X (192 Go)
├── Training ≤ 70B ──► 8× H100 (le standard éprouvé)
├── Training 100B+ ──► 8× B200 + IB 800G + DLC
├── Alternative ouverte ──► 8× MI355X (POC ROCm d'abord)
├── Budget illimité + 2027 ──► Attendre Rubin / Helios
└── Usage < 60 % ──► LOUER (cloud/néocloud), ne pas acheter
```

## 174. Les erreurs de ce guide (honnêteté)

- Les prix sont une photographie au 27/09/2026 — le marché bouge vite.
- Les perfs tok/s sont des ordres de grandeur (±30 %), pas des mesures labo.
- Les chiffres Rubin/Helios sont des claims constructeurs (GTC 2026, Advancing AI
  2026) — à valider en POC.
- Les TDP « configurables » (RTX PRO 6000 400–600 W) dépendent du vBIOS OEM.
- Si une donnée contredit votre devis OEM : **le devis a raison** (c'est lui
  qui est contractuel).

## 175. Feuille de route suggérée pour Zelef

| Horizon | Action |
|---|---|
| Immédiat | Lab 1–2× RTX 5090 pour la R&D RAG (ce guide + le guide Proxmox) |
| 3 mois | POC inférence 70B : 1× RTX PRO 6000 vs location H200 |
| 6 mois | Si prod : 1 nœud 8× L40S (RAG/embeddings) + étude élec salle |
| 12 mois | Extension training : 8× H100 (prix en baisse) ou 8× B200 |
| 18–24 mois | Évaluer Rubin/Helios sur le renouvellement — pas avant |

Chaque étape se rentabilise avant la suivante : pas de « big bang » à 1 M$.

## 176. Conclusion : les 5 règles d'or

1. **La mémoire d'abord** : 80 % des dimensionnements se jouent sur la VRAM,
   pas sur les TFLOPS.
2. **Le système, pas la carte** : un GPU sans serveur, réseau, élec et froid
   adaptés est un presse-papier cher.
3. **L'énergie est un poste** : 20 % du TCO 5 ans — la mesurer, l'optimiser,
   la facturer.
4. **Le logiciel décide** : CUDA vs ROCm se tranche en POC, pas en réunion.
5. **Acheter au bon moment** : ni au pic (2023), ni trop tard (fin de support) —
   et jamais sans 3 devis.

---

*Guide rédigé le 27/09/2026. Données vérifiées par recherche web le même jour.
Re-vérifier prix et disponibilités avant tout achat. Bon courage — et que vos
PUE soient bas et vos tok/s élevés.*
## 177. Fine-tuning : QLoRA chiffré (le cas le plus courant)

| Modèle | Base Q4 (VRAM) | Adapteurs + optimiseur | Total | GPU mini |
|---|---|---|---|---|
| 8B | 5 Go | ~8 Go | ~13 Go | RTX 4090 |
| 32B | 18 Go | ~20 Go | ~38 Go | L40S (48 Go) |
| 70B | 40 Go | ~35 Go | ~75 Go | RTX PRO 6000 (96 Go) |

QLoRA (base gelée en 4 bits + adapteurs LoRA en BF16) divise la VRAM par ~3
vs full fine-tuning. **C'est la méthode par défaut en 2026** pour adapter un
modèle existant — le full FT est réservé au pre-training et aux gros labs.

## 178. Full fine-tuning : le calcul complet (70B, Adam)

| Poste | Octets/param | 70B |
|---|---|---|
| Poids (BF16) | 2 | 140 Go |
| Gradients (FP32) | 4 | 280 Go |
| États Adam (m+v, FP32) | 8 | 560 Go |
| Activations (checkpointées) | ~2 | ~140 Go |
| **Total** | **~16** | **~1 120 Go** |

→ **8× H200 (1,1 To)** minimum, ou 16× H100. D'où l'intérêt du parallélisme
3D (TP+PP+DP) et du sharding d'optimiseur (ZeRO-3 / FSDP).

## 179. ZeRO / FSDP : le sharding qui sauve

| Niveau | Ce qui est shardé | VRAM/GPU (70B, 8 GPU) |
|---|---|---|
| ZeRO-1 | États optimiseur | ~840 Go → ~105 Go/GPU |
| ZeRO-2 | + Gradients | ~560 Go → ~70 Go/GPU |
| ZeRO-3 | + Poids | ~140 Go → ~18 Go/GPU (+ comm) |

