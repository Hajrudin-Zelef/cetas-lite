---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-10
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: ["Google"]
dates: ["2026-09-27"]
keywords: ["datacenter", "benchmarks", "incident", "nand", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [1560, 1745]
sha256: ad0ce91359388bb1dfec8c8292a679bf8dad832673b8afd3c268cb9f02587a69
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

## 163. Dell PowerScale (ex-Isilon) : le NAS scale-out (vérifié)

NAS scale-out propriétaire : POSIX, milliards de fichiers, multi-Po.
Le choix « appliance qui marche » pour du fichier d'entreprise mixte
(bureautique, home, analytics) sans équipe SDS. Prix appliance :
**à vérifier**, c'est du haut de gamme.

## 164. Matrice : quand choisir quoi

| Besoin | Choix par défaut | Alternative |
|---|---|---|
| S3 simple, équipe réduite | MinIO / SeaweedFS | Ceph RGW |
| S3 critique multi-sites | Ceph RGW | Scality |
| Bloc + fichier + objet unifiés | Ceph | — |
| POSIX partagé simple | CephFS / JuiceFS | PowerScale |
| Training IA, débit extrême | WEKA / Lustre | VAST |
| Edge / ARM / léger | Garage | — |
| Backup immuable | tout S3 avec Object Lock | Cloudian |

---

# PARTIE K — FIABILITÉ : SMART, ENDURANCE, FIRMWARE, BURN-IN

## 165. La courbe en baignoire

Mortalité des disques : élevée au début (défauts de fabrication) →
plateau (pannes aléatoires) → remontée (usure). D'où le **burn-in**
(section 174) : on fait tomber les morts-nés **avant** la production,
pas pendant.

## 166. SMART : les attributs HDD qui comptent

| Attribut | Nom | Seuil d'action |
|---|---|---|
| 5 | Reallocated Sector Count | > 0 = surveiller, croissance = remplacer |
| 187 | Reported Uncorrectable Errors | > 0 = remplacer |
| 188 | Command Timeout | croissance = câble/contrôleur ou disque |
| 197 | Current Pending Sector | > 0 = surveiller de près |
| 198 | Offline Uncorrectable | > 0 = remplacer |
| 194 | Temperature | > 45 °C = refroidissement |

Automatise : `smartd` + envoi vers Zabbix. Un disque ne meurt presque
jamais sans prévenir — **si on regarde**.

## 167. NVMe Health Log : les champs critiques

`nvme smart-log` : `percentage_used` (usure), `media_errors`,
`num_err_log_entries`, `temperature`, `power_on_hours`,
`unsafe_shutdowns`. Seuil : `percentage_used` > 70 % = planifier le
remplacement ; `media_errors` > 0 répétés = RMA. Les SSD d'entreprise
exposent aussi des logs étendus (OCP 2.0 — vérifié : supporté par le
D5-P5430).

## 168. DWPD : la définition

Drive Writes Per Day : nombre de fois la capacité du disque qu'on peut
écrire par jour pendant la garantie (5 ans typique).
`TBW = DWPD × capacité × 365 × 5`.
Exemple : 7,68 To à 1 DWPD = 7,68 To/jour × 1825 j ≈ **14 Po** de TBW.

## 169. Endurance vérifiée : exemples 2026

| SSD (vérifié 27/09/2026) | Endurance | Classe |
|---|---|---|
| KIOXIA CD9P-R 7,68 To | 1 DWPD | lecture intensive |
| KIOXIA CD8P | 3 DWPD | mixte |
| Solidigm D5-P5430 30,72 To | 31,92 Po (0,58 DWPD) | QLC capacitaire |
| Solidigm D7-PS1010 1,92 To | 3 504 To (1 DWPD) | performance |

Règle d'achat : **1 DWPD** = standard mixte sain ; 3 DWPD = écriture
intensive (journaux, caches) ; < 1 DWPD (QLC) = lecture/froid uniquement.

## 170. Tableau endurance par usage

| Usage | DWPD mini conseillé |
|---|---|
| Boot / OS | 1 |
| VM mixtes | 1 |
| Bases OLTP / journaux | 3 |
| Caches / burst | 3+ (ou Optane-like / SLC) |
| Backup / objet froid | 0,5-1 (QLC OK) |
| block.db Ceph | 3 (écritures intenses) |

Le block.db d'un cluster hybride : **ne jamais mettre du QLC**
dessus — c'est le composant le plus écrit du cluster.

## 171. Over-provisioning (vérifié)

L'over-provisioning = la NAND cachée au-delà de la capacité annoncée.
Vérifié : le D7-PS1010 1,92 To expose **259,9 Go / 14,5 %** d'OP
(TechPowerUp, 27/09/2026). Plus d'OP = meilleure endurance et des
perfs soutenues plus stables. Certains SSD permettent d'augmenter
l'OP en réduisant la capacité exposée (`nvme format` avec namespace
plus petit) — utile pour transformer un SSD lecture en SSD mixte.

## 172. SLC cache : principe et limite

Beaucoup de SSD (y compris entreprise QLC) absorbent les écritures
dans un cache pseudo-SLC avant de les replier en QLC. Tant que le
cache absorbe : perfs excellentes. Quand il est plein (écriture
soutenue) : chute brutale au débit QLC natif. Pour du workload
soutenu : regarde les perfs **après** saturation du cache (les bons
tests le mesurent — vérifié : méthodologie TweakTown avec
préconditionnement).

## 173. Firmware : les pièges

- Un SSD neuf n'a pas toujours le dernier firmware **stable**
  (le « dernier » n'est pas toujours le « stable »).
- Certains bugs firmware tuent par lots (même modèle, même heure) :
  **ne jamais flasher tout le parc le même jour** ; commence par un
  nœud pilote.
- Note les versions par numéro de série dans l'inventaire.
- Les SSD grand public dans un rôle entreprise : firmwares sans
  power-loss protection → corruption en cas de coupure (voir 174).

## 174. Burn-in : la procédure

Tout disque neuf subit avant production :
1. `smartctl -t long` (HDD) ou écriture complète + lecture (SSD),
2. 2-4 cycles d'écriture/lecture complète (`badblocks -w` ou `fio`),
3. surveillance température et erreurs pendant 48-72 h,
4. `percentage_used` / SMART de référence archivé.
Un disque qui survit au burn-in sans erreur a passé le début de la
courbe en baignoire. **Ne jamais sauter le burn-in** pour « gagner
une semaine » — c'est la semaine la plus rentable du projet.

## 175. URE et BER : les ordres de grandeur

| Classe de disque | BER typique | Sens |
|---|---|---|
| HDD entreprise | 10⁻¹⁵ | 1 erreur / 125 To lus |
| HDD grand public | 10⁻¹⁴ | 1 erreur / 12,5 To lus |
| SSD entreprise | 10⁻¹⁶ - 10⁻¹⁷ | bien meilleur |

C'est ce BER qui condamne le RAID 5 sur gros disques (section 91) :
lire 20 To avec un BER de 10⁻¹⁴ pendant un rebuild, c'est jouer à la
roulette. En Ceph/ZFS, les checksums transforment l'URE en « relecture
sur une autre réplique » : incident, pas catastrophe.

## 176. MTBF et AFR : lire entre les lignes

MTBF Exos : **2 500 000 h** (vérifié). Traduction honnête : AFR
(annualized failure rate) ≈ 0,35 %/an en conditions nominales.
Sur 60 disques : 60 × 0,35 % ≈ **0,2 panne/an** — soit une panne tous
les ~5 ans par nœud. Sur 600 disques : ~2 pannes/an : il faut un
process de RMA rodé, pas de l'improvisation. La MTBF ne prédit pas
*ton* disque, elle dimensionne *ton stock de rechange*.

## 177. AFR terrain : les retours

Les études publiées (Backblaze, Google — ordres de grandeur publics) :
AFR HDD datacenter **0,5-1,5 %/an** selon modèle et température, avec
des lots à problème qui montent à 3-5 %. Règle : suis l'AFR **par
modèle et par lot d'achat** dans ton parc, et bannis les modèles qui
dévient. Un tableau de bord « pannes par modèle » vaut tous les
benchmarks.

## 178. Scrub et patrol read

Le scrub (ZFS/Ceph) lit périodiquement toutes les données et vérifie
les checksums : c'est ce qui transforme la corruption silencieuse en
incident réparable. Patrol read (RAID matériel) : l'équivalent côté
carte. **Planifie les scrubs** (hebdo sur HDD, mensuel sur SSD) et
surveille leur durée : un scrub qui s'allonge = disques qui rament.

## 179. SMART long test vs short

- Short : ~2 min, électronique + petite zone.
- Long : lecture complète de la surface (heures sur 20 To).
En burn-in : long. En production : short hebdo + long trimestriel,
**jamais pendant les fenêtres de backup**. Un long test qui trouve
des secteurs pending = remplace avant la panne.

## 180. Remplacement proactif : la politique

Seuils de remplacement **écrits** dans la doc d'exploitation :
- SMART 5/187/197/198 > 0 et croissant,
- `percentage_used` SSD > 70 %,
- température > 50 °C de façon durable,
- 2 incidents « disque lent » (slow ops) en 30 jours.
Un disque remplacé à 70 % d'usure coûte un RMA planifié ; à 100 % en
pleine nuit, il coûte une astreinte + un rebuild.

## 181. Garantie 5 ans : le standard entreprise (vérifié)

