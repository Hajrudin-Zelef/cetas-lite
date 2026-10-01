---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-15
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["arr", "benchmark", "capex", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [2361, 2524]
sha256: c6b3ca113dc902fe91eb05394e2137acc1f83506bfcd97897192253c616edd57
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

Hypothèses : 30 To utiles VM, 3 nœuds Proxmox+Ceph convergés (lien
avec ton `proxmox_guide.md`). Disques : 6 NVMe 7,68 To par nœud =
18 OSD, réplica 3. Brut = 18 × 6,99 = 125,8 Tio → utiles ≈ 125,8 /
3 × 0,8 ≈ 33 To. RAM : 6 × 3 + 16 + 32 (VM) = **128 Go/nœud**.
CPU : 2× 12c. Réseau : 2×25G (client+cluster sur VLAN séparés —
compromis acceptable à cette échelle). PG : (100×18)/3 = 600 →
**512**. Coût ordre de grandeur : 3 × ~15 k€ = **~45 k€** (à vérifier).
C'est le sweet spot PME : simple, résilient, évolutif par ajout de
nœuds.

## 242. Cas n°4 : objet S3 2 Po (froid, QLC)

Hypothèses : 2 Po utiles S3, RGW, EC 8+3. Disques : SSD QLC E3.S
30,72 To (D5-P5430, vérifié : 0,58 DWPD — OK pour du WORM froid).
Brut = 2000 × 1,375 × 1,15 ≈ 3 162 To → **104 SSD de 30,72 To**.
Nœuds : 13 nœuds × 8 SSD (13 ≥ 11 domains). DB : RGW → 4 % :
104 × 30,72 × 4 % = 128 To de db → les SSD étant homogènes, db
colocalisé (full flash). RAM : RGW gourmand en omap → **256 Go/nœud**.
Conso : 13 × 500 W ≈ 6,5 kW IT. CAPEX disques : 104 × 30,72 To ×
~250 €/To (QLC moins cher — à vérifier) ≈ **800 k€**. Le QLC fait
du 2 Po « performant-froid » là où le HDD demanderait 3× plus de
baies.

## 243. Comparaison des 4 cas (€/To utile/an, TCO 5 ans)

| Cas | €/To utile (TCO 5 ans) | W/To | Latence |
|---|---|---|---|
| Backup 1 Po HDD EC | ~170 € | ~1 W | ~10 ms |
| IA training NVMe | à vérifier (élevé) | ~20 W | < 0,5 ms |
| PME 30 To HCI | ~300 € (ordre de grandeur) | ~15 W | < 2 ms |
| S3 2 Po QLC EC | ~500 € (ordre de grandeur) | ~4 W | ~2 ms |

Tous **à vérifier** sur devis — mais les **rapports** entre cas sont
robustes : le froid HDD coûte ~10× moins cher au To que le NVMe
performant.

## 244. Le dimensionnement PCIe : l'oublié

Un nœud 24 NVMe Gen5 x4 = 96 voies PCIe rien que pour les disques +
2×100G (32 voies) + HBA. Total > 128 voies : il faut du **double
socket** ou des CPU à 128+ voies (EPYC). Vérifie le **bilan PCIe**
de la carte mère **avant** d'acheter les disques : un SSD branché
sur un lien x2 au lieu de x4 perd la moitié de son débit (même
famille de piège que la section 219).

## 245. Le dimensionnement mémoire du MON/MGR

Sur un gros cluster (500+ OSD), les MON deviennent gourmands :
**32-64 Go** par MON, SSD dédié pour la base (les MON écrivent peu
mais lisent beaucoup). 5 MON pour > 200 OSD. Ne colle jamais un MON
sur un nœud OSD saturé en production — ou accepte des élections
lentes en cas de panne.

## 246. RGW : le dimensionnement des gateways

1 gateway RGW ≈ 10-20 Gb/s de S3 (selon TLS et taille d'objets).
Pour 100 Gb/s de S3 : **6-10 gateways** derrière un load-balancer,
sur des nœuds dédiés ou colocalisés (colocalisé = moins cher, dédié
= plus prévisible). Le TLS termine **sur les gateways** (pas sur le
LB) si tu veux du bout-en-bout — avec le coût CPU associé (~1 cœur
/ 5 Gb/s TLS, ordre de grandeur à vérifier).

## 247. Erasure coding et CPU : le chiffrage

Encodage EC 8+3 à 1 Go/s ≈ 2-4 cœurs modernes (isa-l). Sur un nœud
qui ingère 5 Go/s en EC : **10-20 cœurs** rien que pour l'encodage.
C'est pour ça que la section 112 prend la fourchette haute en RGW/EC.
Si le CPU est le goulot : passe en réplica ou ajoute des nœuds.

## 248. Le « petit » cluster : 3 nœuds, 36 HDD

Le classique qui échoue : 3 nœuds × 12 HDD 20 To, réplica 3, PG auto.
Problèmes : 36 OSD seulement → PG/OSD faibles si mal réglés ;
perte d'un nœud = 33 % de capacité en moins = **nearfull immédiat**
si le cluster était à 70 %. Règle : sur 3 nœuds, ne dépasse jamais
**50 % de remplissage** (sinon la perte d'un nœud tue le cluster).

## 249. Extension : ajouter des nœuds

Ajouter 2 nœuds à un cluster de 6 : +33 % de capacité, mais le
rebalance déplace ~1/4 des données → **trafic cluster massif**
pendant des heures/jours. Procédure : `ceph osd set norebalance`
pendant les heures ouvrées, ajout la nuit, `unset` ensuite, et
surveille le `misplaced`. Préviens les utilisateurs : la latence
p99 bouge pendant un rebalance.

## 250. Retrait : sortir un nœud

`ceph osd out` nœud par nœud, **jamais** `ceph osd purge` brutalement
en production. Le drain d'un nœud 12×20 To = 240 To à recopier :
à 5 Go/s = 13 h. Pendant ce temps, le cluster est dégradé : **un
seul nœud à la fois**, et pas pendant la fenêtre de backup.

## 251. Le multi-site : stretch cluster

Ceph stretch (2 sites + tiebreaker) : survit à la perte d'un site.
Coût : latence d'écriture = RTT inter-sites (2 ms à 50 km = OK pour
du backup, KO pour de l'OLTP synchrone). Alternative : 2 clusters +
réplication asynchrone (RGW multisite). Pour ton métier : le stretch
impose aussi le **lien énergie** des deux sites (2 onduleurs, 2 groupes).

## 252. Benchmark : la méthode honnête

1. Préconditionne les SSD (2× écriture complète).
2. `rados bench` sur pool **dédié** (jamais prod — supprime les objets
   `bench*` après, cause documentée de OSD_FULL).
3. Teste 4K aléatoire QD32, 128K séquentiel, 70/30 mixte.
4. Mesure **p99**, pas la moyenne.
5. Refais après remplissage à 50 % (les perfs changent avec le
   remplissage sur QLC et sur HDD zone externe/interne).

## 253. HDD : zones et débit variable

Un HDD 20 To ne fait pas 280 Mo/s partout : ~280 Mo/s en zone externe,
~150 Mo/s en zone interne. En début de vie (vide), tout est rapide ;
à 80 % plein, les écritures tombent en zone lente. Pour les SLA de
backup : dimensionne sur le **débit zone interne**, pas le pic.

## 254. SSD : le steady state

Un SSD neuf sort de boîte 2× plus vite qu'après préconditionnement.
Tous les chiffres constructeur sont en **burst**. Pour dimensionner :
utilise les chiffres **steady state** (tests avec préconditionnement —
vérifié : méthodologie TweakTown). Écart typique : 30-50 % de moins
en écriture soutenue sur TLC, 70 % sur QLC.

## 255. Le test de panne : game day

1×/trimestre : tue un OSD (en prod, en heures creuses), mesure le
temps de recovery et l'impact p99. 1×/an : tue un nœud entier. 1×/an :
coupure électrique simulée (NUT). **Documente** les résultats : c'est
ce document qui justifie le budget d'extension quand le DSI demande
« pourquoi il faut 30 % de marge ».

## 256. Capacité vs performance : le point de bascule

Formule : si `capacité_utile / débit_cible` < 1 h, tu es **limité en
capacité** (ajoute des disques) ; si > 10 h, tu es **limité en
performance** (ajoute du réseau/CPU). Entre les deux : les deux.
Exemple : 100 To utiles à 10 Go/s = 2,8 h → équilibré. À 100 Go/s =
0,28 h → le réseau/CPU domine, les disques sont presque « gratuits ».

## 257. Le tiering manuel : politique de déplacement

Sans cache tiering (section 137), le tiering = **politique** : données
< 30 jours sur NVMe (device class), > 30 jours migrées sur HDD/EC par
script applicatif (RGW lifecycle, rbd migration, cron + rsync).
C'est moins magique que du tiering automatique, mais c'est **prévisible**
et ça ne corrompt rien. Documente la politique, alerte sur son échec.

## 258. La doc d'exploitation minimale

Pour chaque cluster : schéma réseau + électrique, plan de câblage SAS
(photo), versions firmware (HBA/expander/SSD), seuils d'alerte,
procédures (ajout/retrait nœud, RMA disque, upgrade Ceph, arrêt NUT),
contacts support avec numéros de contrat. **Si ce n'est pas écrit,
ça n'existe pas** — surtout à 3 h du matin en astreinte.

---

# PARTIE R — RUNBOOKS D'EXPLOITATION

## 259. Runbook : remplacement d'un disque Ceph

