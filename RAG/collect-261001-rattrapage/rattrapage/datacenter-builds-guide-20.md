---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-20
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Broadcom", "Intel", "Samsung"]
dates: []
keywords: ["arr", "capex", "compute", "gpu", "intel"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [3538, 3749]
sha256: b6f502eddcc243ca455c3a72dbed7aef294ead512a5a6e7bd5e79b292acad72c
---

# Datacenter Builds — Le guide des BOMs

| Modèle | Format | Endurance | Lecture | Prix/To indicatif |
|---|---|---|---|---|
| Samsung PM9A3 | U.2 | 1 DWPD | 6 900 Mo/s | ≈ 555 $/To (7,68 To) |
| Samsung PM9A5 | U.2 | 3 DWPD | 6 900 Mo/s | ≈ 800 $/To (à vérifier) |
| Micron 7450 PRO | U.2/E1.S | 1–3 DWPD | 6 800 Mo/s | ≈ 600–900 $/To (à vérifier) |
| Kioxia CM7 | U.2/E3.S | 1–3 DWPD | 14 000 Mo/s (PCIe 5) | ≈ 700–1 000 $/To (à vérifier) |
| Solidigm D7-P5520 | U.2 | 1–3 DWPD | 7 100 Mo/s | ≈ 600 $/To (à vérifier) |

Le PM9A3 reste la référence prix/perf en read-intensive. Pour du PCIe 5.0
plein débit (14 Go/s), Kioxia CM7 et équivalents — mais vérifiez que le
châssis et le CPU négocient bien en Gen5 (sinon vous payez pour rien).

---

## 215. Comparatif NIC 25/100G

| NIC | Ports | Débit | RDMA | Prix indicatif |
|---|---|---|---|---|
| Intel E810-XXVDA2 | 2× 25G | 25G | RoCE | ≈ 450 € |
| Broadcom P225P | 2× 25G | 25G | RoCE | ≈ 350 € (à vérifier) |
| Mellanox CX6-DX | 2× 100G | 100G | RoCE v2 | ≈ 1 100 € |
| Mellanox CX7 | 2× 400G | 400G | RoCE v2 | ≈ 3 000 € (à vérifier) |

Pour Ceph : CX6-DX en 100G = le standard. L'E810 suffit en 25G front.
**Vérifiez la compatibilité OCP vs PCIe** avec votre châssis avant commande.

---

## 216. Cas chiffré : Ceph NVMe 1 Po utile, réplica 3

Besoin brut : 1 000 / 0,85 × 3 = 3 530 To → 29 nœuds M (29 × 123 To).

| Poste | Détail | Total |
|---|---|---|
| 29× nœuds OSD M | 96 180 € × 29 | 2 789 220 € |
| 5× MON/MGR (cluster large) | 13 485 € × 5 | 67 425 € |
| 2× switch 100G 64p (back) | 25 000 € × 2 | 50 000 € |
| 2× switch 25G 48p (front) | 8 000 € × 2 | 16 000 € |
| Câblage/optiques | lot | 30 000 € |
| **CAPEX** | | **≈ 2 953 000 €** |
| P_max | 29 × 1,49 + 5 × 0,45 + 2 | **≈ 47 kW** |
| **Coût/To utile** | 2 953 000 / 1 000 | **≈ 2 950 €/To** |

47 kW = **2 racks à 24 kW** (RDHx) ou 1 rangée DLC. À cette échelle,
négociez les SSD par lot de 500 : −10 à 15 % réaliste.

---

## 217. Cas chiffré : VDI 1 000 postes bureautiques

| Poste | Détail | Total |
|---|---|---|
| 10× serveurs VDI S (N+1 : 9+1) | 64 205 € × 10 | 642 050 € |
| Licences VDI (1 000 users) | ≈ 200 €/user/an | 200 000 €/an |
| Stockage profils (CephFS 60 To) | — | 80 000 € |
| 1 000× thin clients | 400 € × 1 000 | 400 000 € |
| Réseau (quote-part) | — | 40 000 € |
| **CAPEX année 1** | | **≈ 1 162 000 € + 200 k€/an** |
| P_max | 10 × 1,35 kW | **≈ 13,5 kW** |
| **Coût/poste/an (5 ans)** | (1 162 000 + 1 000 000) / 1000 / 5 | **≈ 432 €** |

---

## 218. Cas chiffré : PostgreSQL 10 To actifs, HA + backup

| Poste | Détail | Total |
|---|---|---|
| 2× PG Warehouse (adapté 10 To) | 144 125 € × 2 | 288 250 € |
| 1× témoin | — | 8 000 € |
| Repo backup (30 To utiles) | serveur S adapté | 30 000 € |
| **CAPEX** | | **≈ 326 000 €** |
| P_max | 2 × 1,54 + 0,5 | **≈ 3,6 kW** |

---

## 219. Dimensionnement disjoncteurs — exemple TGBT complet

Salle §181 (60 kW IT) :

| Départ | Calibre | Courbe | Section câble |
|---|---|---|---|
| 8× PDU rack (tri 32 A) | 8× 32 A tétra | C | 5G6 mm² |
| 2× onduleurs 80 kW | 2× 160 A tétra | D | 4× 70 mm² |
| Groupe 200 kVA | 250 A tétra | D | 4× 120 mm² |
| Climatisation | 3× 63 A tétra | D | 5G16 mm² |
| Éclairage + prises | 2× 20 A | C + 30 mA | 3G2,5 mm² |
| Arrivée générale | 400 A tétra | — | TGBT 400 A |

**Étude par un électricien qualifié obligatoire** — ce tableau est une base
de discussion, pas un schéma d'exécution.

---

## 220. Tableau — temps de transfert utiles

| Volume | 1 Gb/s | 10 Gb/s | 25 Gb/s | 100 Gb/s |
|---|---|---|---|---|
| 1 To | 2 h 30 | 15 min | 6 min | 1 min 30 |
| 10 To | 25 h | 2 h 30 | 1 h | 15 min |
| 100 To | 10 j | 25 h | 10 h | 2 h 30 |
| 1 Po | 104 j | 10 j | 4 j | 25 h |

À coller au mur de la salle : c'est ce tableau qui fait comprendre pourquoi
le 25G est le minimum en 2026 pour tout ce qui brasse de la donnée.

---

## 221. Erreurs de devis les plus coûteuses (vécu)

| Erreur | Surcoût typique |
|---|---|
| RAM commandée en 2 DPC au lieu d'1 (perte 20 % bande passante) | refonte partielle : 5–10 k€ |
| NIC PCIe au lieu d'OCP (ne rentre pas) | retour + port : 1–2 k€ |
| SSD SATA commandés au lieu de NVMe U.2 (mauvais backplane) | incompatibles : 10–30 k€ |
| Onduleur sans bypass externe | arrêt pour maintenance : coût d'indispo |
| PDU 16 A pour un rack à 9 kW | déclenchements : 2 k€ de PDU |
| Garantie oubliée (restée à 1 an) | extension après-coup : +30 % |

**Relisez chaque devis ligne par ligne avec la check-list §122.**
Le copier-coller d'un commercial pressé coûte plus cher que votre temps.

---

## 222. Garantie et pièces détachées — la politique

| Parc | Politique recommandée |
|---|---|
| < 10 serveurs | Garantie constructeur 5 ans NBD |
| 10–50 serveurs | 3 ans NBD + stock pièces (2 disques, 1 PSU, 1 ventilo) |
| > 50 serveurs | 3 ans + stock 5 % + contrat 4 h sur le critique |

Un disque de rechange **sur étagère** vaut mieux qu'une garantie NBD quand
le rebuild prend 28 h (§59) : le NBD ne raccourcit pas le rebuild, le spare
si.

---

## 223. Câbles d'alimentation : C13 vs C19 vs C21

| Connecteur | Intensité | Usage |
|---|---|---|
| C13/C14 | 10 A | serveurs 1U/2U standard |
| C19/C20 | 16 A | serveurs GPU, gros 2U/4U |
| C21/C22 | 20 A | très haute densité (à vérifier) |

**Vérifiez le connecteur du serveur AVANT de commander les PDU** : une PDU
full C13 ne branche pas un serveur GPU en C19. Les PDU modernes mixtes
(C13/C19/C21, §129) évitent le problème.

---

## 224. Températures de fonctionnement — limites réelles

| Composant | Limite conseillée | Throttling/danger |
|---|---|---|
| CPU EPYC | < 75 °C | > 95 °C (Tjmax) |
| GPU L40S/MI300X | < 80 °C | > 90–95 °C |
| NVMe | < 60 °C | > 70–75 °C |
| HDD | 25–40 °C | > 50 °C |
| Entrée rack (ASHRAE A1) | 18–27 °C | > 32 °C (classe A2+) |

Le NVMe est le plus fragile : à 75 °C il throttle et vos IOPS s'effondrent.
**Surveillez les températures SSD** (smartctl), surtout dans les 2U denses.

---

## 225. Check-list annuelle d'exploitation

- [ ] Revue thermique caméra IR (connexions électriques chaudes ?)
- [ ] Resserrage des bornes TGBT/PDU (dilatation thermique)
- [ ] Test groupe en charge 1 h + analyse fioul
- [ ] Test autonomie onduleur (décharge réelle mesurée)
- [ ] Exercice restore backup complet (1 appli critique)
- [ ] Failover firewall réel
- [ ] Revue des accès (badges, comptes BMC)
- [ ] Mise à jour as-built (tout ce qui a changé dans l'année)
- [ ] Renégociation contrats maintenance
- [ ] Plan de capacité : à 70 % d'une ressource, on commande

---

## 226. Plan de branchement PDU — exemple par serveur

```
Serveur srv-03 (Compute M, 2 PSU) :
  PSU-A (C19) → PDU-A prise 12 (phase L1) — cordon ROUGE
  PSU-B (C19) → PDU-B prise 12 (phase L1) — cordon NOIR

Règle : même numéro de prise sur A et B pour un serveur donné.
Le technicien trouve la prise en 10 secondes, même à 3 h du matin.
Documenté dans : plan de câblage (§203) + étiquette sur le cordon.
```

---

## 227. Chemins de câbles — dimensionnement

| Câbles | Largeur chemin | Charge |
|---|---|---|
| 50 DAC + 20 fibres | 300 mm | ~15 kg |
| 150 câbles mixtes | 600 mm | ~40 kg |
| Fibre inter-salles | goulotte fermée 100 mm | — |

- Charge admissible : vérifiez la fiche (un chemin plein de cuivre pèse lourd).
- **Séparation** : cuivre / fibre / électrique dans 3 chemins ou 3
  compartiments. Le 230 V à côté du DAC 25G = diaphonie et danger.
- Mise à la terre du chemin métallique (obligatoire).

---

## 228. Consignation électrique (LOTO) — procédure

