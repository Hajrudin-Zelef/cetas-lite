---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-12
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom", "Samsung"]
dates: ["2026-09-27"]
keywords: ["datacenter", "arr", "capex", "compute", "hbm", "nand", "rack-scale"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [1921, 2095]
sha256: 5a231c8e05a54a776a5b150bcfbdad0778d3a924953bc59ad73c13d3f1b8561e
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

10 nœuds 60 baies = 9,8 kW IT → ~15 kW à climatiser (PUE 1,5).
En eau glacée : ~15 kW / (4,18 × ΔT 6 °C) ≈ 0,6 L/s. En détente directe :
vérifie la capacité des CRAC **par rangée**, pas globale — une rangée
stockage dense peut saturer son CRAC alors que la salle est à 50 %.
Voir ton `proxmox_guide.md` (dimensionnement) et `onduleurs_ups_guide.md`.

## 199. Redondance N+1 électrique

Double alimentation par nœud (A/B), 2 PDU par rack, 2 onduleurs ou
2 chaînes. Règle de capacité : **chaque chaîne doit tenir 100 % de la
charge** (pas 50 %) — sinon la perte d'une chaîne = surcharge de
l'autre = cascade. Teste le basculement en charge réelle 1×/an.

## 200. Le pic de spin-up : le chiffre

60 HDD × ~20 W au spin-up (2× le nominal, ~10-20 s) = **1 200 W**
transitoires par nœud. Sur 10 nœuds qui redémarrent ensemble après
une coupure : 12 kW de pic. Mitigations : spin-up **échelonné**
(staggered spin-up, réglage HBA/BIOS), redémarrage séquentiel des
nœuds, onduleur dimensionné pour le pic pas la moyenne.

## 201. Tableau récapitulatif énergie

| Architecture | W/To utile (ordre de grandeur) | PUE cible | Point chaud |
|---|---|---|---|
| HDD 60 baies EC | ~1-1,5 W/To | 1,4-1,6 | spin-up, airflow |
| Hybride | ~3-5 W/To | 1,4-1,6 | db SSD |
| Full NVMe réplica 3 | ~30-40 W/To | 1,3-1,5 | CPU/NIC, SSD Gen5 |

## 202. Checklist énergie

- [ ] Bilan de puissance par nœud (tableau 184/185) **écrit**
- [ ] Alimentations Titanium sur tout nœud neuf
- [ ] PDU A/B, chaque chaîne à 100 % de la charge
- [ ] Pic de spin-up intégré à l'onduleur
- [ ] Séquence d'arrêt NUT écrite et testée 2×/an
- [ ] Sondes température disques → supervision
- [ ] TCO 5 ans avec énergie, pas CAPEX seul

---

# PARTIE M — À VENIR : ANNONCES VÉRIFIÉES AU 27/09/2026

## 203. CXL : l'état réel (vérifié)

Compute Express Link : mémoire extensible et partageable sur PCIe.
État vérifié le 27/09/2026 (analyses 2026) :
- **CXL 2.0** : en production (ex. Azure), expansion mémoire Type 3.
- **CXL 3.x** : spécifications publiées, silicium et plateformes en
  cours de maturation — **pas de déploiement large** en 2026.
- Produits : Samsung CMM-D (E3.S), SK hynix CMM-DDR5 (E3.S), Micron
  CZ120/CZ122 (E3.S) — 96-256 Go, PCIe Gen5 x8.
Lien stockage : les modules CXL utilisent le **format E3.S** — tes
baies EDSFF de 2026 pourront accueillir de la mémoire CXL demain.
Pour le stockage pur, l'impact 2026-2027 reste indirect (caches
métadonnées géants, déduplication en mémoire).

## 204. Marvell Structera S 30260 (vérifié)

Switch CXL 3.0 : 260 voies, jusqu'à 4 To/s agrégés, **sampling Q3 2026**
(communiqué Marvell, vérifié). Le CXL 2.0 (Structera S 20256) est en
production. À suivre pour le pooling mémoire rack-scale 2027+.

## 205. EDSFF : la trajectoire (vérifié)

E3.S est en 2026 **le format standard du neuf** all-flash (Supermicro,
KIOXIA, Solidigm, Samsung, Micron). Prochaines étapes : E3.L pour les
très hautes capacités, Gen6 (Micron 9650 déjà là), et convergence avec
CXL (même connecteur E3.S). Acheter E3.S en 2026, c'est acheter le
format des 5 prochaines années.

## 206. PLC : l'état (vérifié)

PLC (5 bits/cellule) : Solidigm a démontré un **SSD PLC fonctionnel en
2022** (FMS), mais en 2026 **aucun produit PLC commercial** — Solidigm
déclare publiquement rester focalisée sur la QLC (interview TechRadar,
vérifiée) : « PLC faces major technical hurdles in endurance ».
Conclusion : **PLC = non trouvé en produit au 27/09/2026**. La densité
2026-2027 viendra de la QLC + 200+ couches, pas de la PLC.

## 207. SSD 200+ To : les annonces (vérifié)

- **Samsung BM1773 : 245,76 To**, QLC, PCIe 5.0, annoncé sept. 2026
  (vérifié) — pour datacenters IA haute densité, 0,6 DWPD.
- **Solidigm : roadmap 245 To+ avant fin 2026** (interview vérifiée).
- **KIOXIA LC9 : 245,76 To** E3.S (tableau sectoriel — **à vérifier**
  la disponibilité).
- **Micron 6600 ION : 122-245 To** E3.S, en qualification hyperscalers
  (tableau sectoriel — **à vérifier**).
- **SK hynix PS1012 : 244 To** (roadmap — **à vérifier**).
Avec des 245 To, un 2U32 E3.S = **7,8 Po bruts**. Le pétaoctet par U
devient le pétaoctet par **demi-U**. L'enjeu se déplace : rebuild,
blast radius, et surtout **le réseau**.

## 208. PCIe Gen6 / Gen7

Gen6 : 64 GT/s (~28 Go/s en x4) — le Micron 9650 est le premier SSD
Gen6 (vérifié, août 2026). Gen7 : spécification en cours, produits
**non trouvés au 27/09/2026**. Note : à partir de Gen6, un seul SSD
sature presque un lien 200G — le design réseau doit suivre (400G/800G).

## 209. HDD : HAMR et 30 To+ (vérifié)

- **Seagate Exos M 30 To** (HAMR/Mozaic 3+, 10 plateaux de 3 To) :
  ~565-600 $, soit **~19 $/To** (vérifié, sept. 2026) — le meilleur
  €/To du moment en entreprise, CMR, 7200 tr/min, 300 Mo/s, 5 ans
  de garantie, MTBF 2,5 M h.
- WD : gamme Ultrastar 24-28 To (à vérifier les références exactes).
- Trajectoire : 40-50 To HAMR annoncés en roadmap (à vérifier).
Le HDD n'est pas mort : à ~19 $/To vs ~350-600 $/To le NVMe, il reste
**20× moins cher** au To pour le froid. Le tiering HDD/QLC/NVMe a
encore de beaux jours.

## 210. QLC : la position officielle Solidigm (vérifié)

Interview Solidigm (vérifiée, 2026) : 120+ exaoctets de QLC livrés,
D5-P5336 122,88 To en production, focus QLC (pas PLC). La QLC est
devenue **la NAND du capacitaire flash** : 30,72 To en E3.S à ~0,6
DWPD. Pour ton tiering : QLC = le nouveau « disque froid performant »
entre HDD et TLC.

## 211. NVMe FDP — Flexible Data Placement (vérifié)

Supporté par le KIOXIA NX1 (vérifié) et la spec OCP Datacenter NVMe
SSD 2.6 : l'hôte indique au SSD où placer les données (par flux),
réduisant l'amplification d'écriture. Intérêt direct : **endurance
et perfs soutenues** sur workloads mixtes (Ceph !). À suivre côté
support noyau Linux / Ceph.

## 212. Computational storage (mention)

SSD avec calcul embarqué (compression, filtrage) : standard en
gestation, produits de niche. **Non trouvé en déploiement large au
27/09/2026** — à surveiller pour le traitement près des données IA.

## 213. Marché : allocation et prix (vérifié, sept. 2026)

Fusion Worldwide Greensheet (vérifié) : **allocation contrainte, prix
en hausse** sur : SSD entreprise PCIe 5.0 / haute capacité, HBM,
DDR5 haute densité. Conséquences achats : devis datés, commandes
fermes tôt, double sourcing, stocks tampons. Le « meilleur prix »
d'aujourd'hui peut être indisponible demain.

## 214. Ce qui n'existe pas (encore) au 27/09/2026

- SSD PLC commercial : **non trouvé**.
- PCIe Gen7 en produit : **non trouvé**.
- CXL 3.x en déploiement large : **non trouvé** (sampling).
- JBOD 90+ baies en catalogue public généraliste : **non trouvé**
  (vérifié partiellement — à vérifier chez intégrateurs).
- Disques SAS 24G grand public/grand volume : rares et chers.

## 215. Veille : les sources à suivre

- SNIA / EDSFF (specs formats), CXL Consortium (specs),
- communiqués Samsung / Solidigm / KIOXIA / Micron (roadmaps NAND),
- docs.ceph.com (releases Reef/Squid/Tentacle),
- Broadcom / Microchip (fiches HBA),
- Fusion Worldwide Greensheet (tension marché, mensuel).

## 216. Recommandations d'achat 2026-2027

1. **Neuf performant** : E3.S Gen5 (CD9P-R, D7-PS1010, PM9D3) — vérifié.
2. **Neuf capacitaire** : HDD 20-30 To CMR (Exos M 30 To — vérifié).
3. **HBA** : 9500 tri-mode (12G) ; 9600 si besoin 24G réel.
4. **Châssis** : E3.S 1U16/2U32 pour l'all-flash ; 4U60 pour le froid.
5. **Évite** : U.2 pour du neuf dense (sauf contrainte existante),
   cache tiering Ceph, PLC (n'existe pas), RAID 5 sur gros disques.

---

# PARTIE N — PIÈGES TERRAIN (17 PIÈGES VÉCUS OU DOCUMENTÉS)

## 217. Piège n°1 : la HBA livrée en IR au lieu d'IT

