---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-4
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom", "Samsung"]
dates: ["2026-09-27"]
keywords: ["datacenter", "nand"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [492, 668]
sha256: 2bcc9f04bcf0308ceca603e1230866358682aff7e9a71cc31b7729df0c74a95a
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

Le pendant externe du 24G, présent sur les 9600-16e (vérifié : fiches
produit). Remplace le 8644 pour les liaisons HBA → JBOD en 24G.
Compatibilité descendante : un JBOD 12G (8644) se branche sur un port
8674 **avec le bon câble adaptateur**, négociation à 12G.

## 49. Longueurs et limites physiques

| Type | Distance indicative | Note |
|---|---|---|
| Cuivre passif SAS HD | 1-4 m | dépend de l'AWG et du débit |
| Cuivre actif | jusqu'à ~10 m | à vérifier par référence |
| Optique AOC | 10-100 m | cher, pour baies distantes |

Règle terrain : **le plus court possible**, jamais de câble enroulé en
boucle (diaphonie), et un câble de rechange de chaque type en stock.
Un câble SAS qui vieillit = erreurs CRC intermittentes = enfers de diagnostic.

## 50. Expanders SAS : principe

Un expander est un « switch SAS » : 1 port amont vers la HBA, N ports vers
les disques. Il permet 60 disques derrière une HBA 16 voies. Coût : la
bande passante amont est **partagée** (voir section 80), et l'expander
ajoute ~quelques µs de latence. Deux expanders en cascade = bande passante
encore divisée.

## 51. DataBolt2 (vérifié le 27/09/2026)

Technologie Broadcom (série 9600) : le contrôleur parle **24G vers l'amont**
tout en servant des disques/backplanes **12G/6G** en aval, avec buffering
intelligent. Intérêt : upgrader la HBA sans changer les disques ni le
backplane, et gagner quand même sur l'agrégat. Ce n'est pas de la magie :
le débit d'un disque 12G reste 12G.

## 52. T-10 EEDP / DIF (vérifié)

End-to-End Data Protection : 8 octets de protection (guard tag, app tag,
ref tag) ajoutés à chaque bloc 512 B → blocs 520 B sur les disques
entreprise. Vérifié : supporté par les 9600. Avec ZFS/Ceph qui font leurs
propres checksums, l'EEDP est redondant mais inoffensif ; en RAID matériel
classique, il reste utile. Les disques « 512e » grand public ne l'ont pas.

## 53. SATA sur contrôleur SAS : les règles

Le SAS accepte le SATA (tunneling STP), l'inverse est faux. Mais :
- pas de double-domaine en SATA (un seul chemin),
- pas de commandes SCSI avancées (pas d'EEDP),
- les SATA derrière un expander SAS partagent la file plus brutalement,
- **SATA + expander + RAID = la combinaison la plus fragile** : privilégie
  le branchement direct ou des expanders de qualité.
Pour Ceph sur SATA : HBA IT + backplane correct = OK, c'est courant.

## 54. Power management SAS (vérifié)

La norme SAS 2.1 (supportée par les 9600, vérifié) inclut le power management
des liens. En pratique datacenter : on le laisse souvent désactivé sur les
liens HBA↔JBOD (la renégociation coûte de la latence), et on gère l'énergie
au niveau disque (spin-down, section 188) plutôt qu'au niveau lien.

## 55. Négociation de vitesse

Chaque lien négocie au plus haut débit commun : 24G↔24G = 22,5G effectifs,
24G↔12G = 12G, etc. (vérifié : 22,5/12/6 Gb/s sur les 9600). Diagnostic :
`storcli` / `sas2ircu` affichent la vitesse négociée par phy. **Un lien qui
négocie en dessous de son nominal = câble, connecteur ou firmware à
suspecter en premier.**

## 56. Compatibilité descendante

Le SAS est historiquement exemplaire : un contrôleur 24G parle à un disque
6G de 2012. C'est ce qui rend les migrations douces possibles (HBA neuve,
disques anciens). L'exception : certains très vieux expanders 3G/6G avec
des disques 12G+ peuvent causer des resets — à tester en burn-in.

## 57. Limites : nombre de périphériques

| Élément | Limite indicative |
|---|---|
| HBA 9500/9600 | 1024 SAS/SATA, 32 NVMe (vérifié) |
| Domaine SAS (adressage) | 16 384 adresses théoriques |
| Pratique par HBA | quelques centaines avant que le firmware ne rame |

En pratique, on ne met jamais 1000 disques sur une HBA : le goulot de
bande passante arrive bien avant (section 80).

## 58. Quand le SAS a encore du sens en 2026

- Capacitaire HDD : le 12G est le standard, pas cher, mature.
- SSD SAS 12G/24G : dual-port natif (deux chemins sans expander spécial),
  utile en SAN classique.
- Transition : DataBolt2 permet de monter en 24G sans tout changer.
- En revanche : pour du NVMe neuf performant, **le SAS est un détour** —
  vise le PCIe direct ou le tri-mode.

---

# PARTIE E — U.2 / U.3 / EDSFF (NVMe)

## 59. U.2 (SFF-8639) : le standard installé

Le 2,5" 15 mm avec connecteur SFF-8639 : PCIe x4 (jusqu'à Gen5), hot-swap,
présent dans 90 % des serveurs NVMe 2020-2025. Limites : épais, bloque
l'air en façade, 25 W max pratiques par baie sans refroidissement spécial.
Reste un excellent choix « sûr » en 2026 pour du NVMe Gen4.

## 60. U.3 : la différence avec U.2

U.3 = **même connecteur physique** que U.2, mais avec négociation **tri-mode**
(SAS/SATA/NVMe) sur les mêmes broches, pilotée par le backplane (UBM).
Un SSD U.3 fonctionne dans une baie U.3 ; dans une baie U.2 pure, un SSD
U.3 **peut ne pas être reconnu** (pas de négociation tri-mode côté hôte).
Règle d'achat : **baie U.3 + SSD U.3 = le couple sûr et évolutif**.

## 61. UBM / SFF-TA-1005 (vérifié)

Universal Bay Management : le standard qui permet au backplane de dire au
SSD quel protocole parler. Vérifié : les HBA Broadcom 9500 sont « UBM ready ».
Sans UBM, le tri-mode ne fonctionne pas : la baie doit être câblée et
firmwarée pour. À vérifier à la commande du châssis : « backplane UBM ? ».

## 62. Tableau de compatibilité U.2 / U.3

| SSD \ Baie | U.2 pure | U.3 (UBM) |
|---|---|---|
| SSD U.2 NVMe | OK | OK |
| SSD U.3 tri-mode | **non garanti** | OK |
| SSD SAS 12G | OK (baie SAS) | OK |

En cas de doute : mets tout en U.3 des deux côtés. Le surcoût est faible,
le regret d'une baie incompatible est grand.

## 63. EDSFF : pourquoi un nouveau format

L'U.2 n'a jamais été pensé pour la flash : boîtier épais, connecteur en bout,
air bloqué. L'EDSFF (Enterprise & Datacenter Standard Form Factor, SNIA)
redessine le SSD autour de la NAND : **règles fines verticales, air qui
circule entre elles, plus de surface de PCB, plus de puissance admissible**.
Résultat vérifié : 0,5 Po en 1U16 E3.S, 1 Po en 2U32 (Supermicro, 27/09/2026).

## 64. E1.S (vérifié)

Le format « 1U » : règle courte (~111 mm), épaisseurs 5,9 / 9,5 / 15 / 25 mm.
Le remplaçant désigné du M.2 en datacenter, hot-swap en façade. Vérifié :
Samsung PM9A3 (Gen4), KIOXIA NX1 (E1.S, refroidissement liquide direct —
27/09/2026), Micron 9650 (Gen6, E1.S + E3.S). Densité typique : 24 baies
E1.S en 1U.

## 65. E1.L (vérifié)

La « règle longue » (~318 mm) : maximise la surface NAND → **capacité par
baie**. Moins courant que E1.S/E3.S dans les catalogues généralistes 2026,
plutôt chez les hyperscalers. Si ton besoin est « max de To par U en 1U »,
c'est lui ; sinon E1.S suffit.

## 66. E3.S (vérifié)

Le format « 2U » généraliste : carte ~76 × 112,75 mm, 7,5 mm d'épaisseur,
**le successeur désigné de l'U.2**. Vérifié : KIOXIA CD9P-R 7,68 To E3.S
(14 800 Mo/s, 2 600K IOPS), Solidigm D7-PS1010 E3.S (1,92-15,36 To),
Samsung PM9D3 128 To E3.S en production, Micron 9650 E3.S Gen6. **Pour tout
projet NVMe neuf 2026+ : E3.S par défaut.**

## 67. E3.L (vérifié)

Version longue de E3 pour très haute capacité. Mentionné dans la
documentation EDSFF/SNIA ; moins déployé que E3.S dans l'offre publique
2026. À considérer si les roadmaps 245 To+ (section 208) arrivent en E3.L.

## 68. EDSFF et PCIe Gen5 / Gen6

L'EDSFF est agnostique de la génération PCIe : les baies E3.S actuelles
sont câblées Gen5 x4 (16 voies possibles sur E3). Le **Micron 9650**
(vérifié, août 2026) est le premier SSD **PCIe Gen6** E1.S/E3.S du marché,
décliné en versions Pro (lecture intensive) et Max (mixte), air ou liquide.
Gen6 = ~28 Go/s par x4 : le réseau (même 400G) devient le goulot avant le SSD.

## 69. Refroidissement EDSFF : air et liquide (vérifié)

