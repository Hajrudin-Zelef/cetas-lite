---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-2
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: ["2026-09-27"]
keywords: ["amd", "clearwater forest", "datacenter", "inference", "intel"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [148, 289]
sha256: 3b236f6595bfd4bd9203e071732b2ccd0fd24a4079b182df6039578e52677d55
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

| Référence | Cœurs/threads | Variante | Base/Boost | TDP | L3 | Prix 1KU |
|---|---|---|---|---|---|---|
| 9965 | 192/384 | Zen 5c | 2,25/3,7 | 500 W | 384 Mo | 14 813 $ |
| 9845 | 160/320 | Zen 5c | 2,1/3,7 | 390 W | 320 Mo | 13 564 $ |
| 9825 | 144/288 | Zen 5c | 2,2/3,7 | 390 W | 384 Mo | 13 006 $ |
| 9755 | 128/256 | Zen 5 | 2,7/4,1 | 500 W | 512 Mo | 12 984 $ |
| 9745 | 128/256 | Zen 5c | 2,4/3,7 | 400 W | 256 Mo | 12 141 $ |
| 9655 | 96/192 | Zen 5 | 2,6/4,5 | 400 W | 384 Mo | 11 852 $ |
| 9655P | 96/192 | Zen 5 | 2,6/4,5 | 400 W | 384 Mo | 10 811 $ |
| 9645 | 96/192 | Zen 5c | 2,3/3,7 | 320 W | 384 Mo | 11 048 $ |
| 9565 | 72/144 | Zen 5 | 3,15/4,3 | 400 W | 384 Mo | 10 486 $ |
| 9575F | 64/128 | Zen 5 | 3,3/5,0 | 400 W | 256 Mo | 11 791 $ |
| 9555 | 64/128 | Zen 5 | 2,4/4,3 | 300 W | 256 Mo | 7 983 $ |
| 9555P | 64/128 | Zen 5 | 3,2/4,4 | 360 W | 256 Mo | 9 826 $ |
| 9475F | 48/96 | Zen 5 | 3,65/4,8 | 400 W | 256 Mo | 7 592 $ |
| 9455 | 48/96 | Zen 5 | 3,15/4,4 | 300 W | 192 Mo | 4 819 $ |
| 9375F | 32/64 | Zen 5 | 3,8/4,8 | 320 W | 256 Mo | 5 306 $ |
| 9355 | 32/64 | Zen 5 | 3,55/4,4 | 280 W | 256 Mo | 2 998 $ |
| 9335 | 32/64 | Zen 5 | 3,0/4,4 | 210 W | 256 Mo | 3 178 $ |
| 9275F | 24/48 | Zen 5 | 4,1/4,8 | 320 W | 256 Mo | 3 439 $ |
| 9255 | 24/48 | Zen 5 | 3,25/4,3 | 200 W | 128 Mo | 2 495 $ |
| 9175F | 16/32 | Zen 5 | 4,2/5,0 | 320 W | 512 Mo | 4 256 $ |
| 9135 | 16/32 | Zen 5 | 3,65/4,3 | 200 W | 64 Mo | 1 214 $ |
| 9115 | 16/32 | Zen 5 | 2,6/4,1 | 125 W | 64 Mo | 726 $ |
| 9015 | 8/16 | Zen 5 | 3,6/4,1 | 125 W | 64 Mo | 527 $ |

Lecture du tableau : le prix au cœur chute avec la densité (le 9965 revient
à ~77 $/cœur contre ~66 $/cœur pour le 9015… non : 527/8 = 66 $/cœur pour le
petit, 14813/192 = 77 $/cœur pour le gros). **Le gros CPU n'est pas moins cher
au cœur ; il est moins cher au rack** (moins de serveurs, moins de licences
par socket, moins de ports réseau — voir sections 25-26).

## 7. Suffixes EPYC : F, P, et les gammes spéciales

- **F (Frequency)** : 9175F, 9275F, 9375F, 9475F, 9575F. Fréquences boost
  maximales (jusqu'à 5,0 GHz), TDP élevés (320-400 W) pour peu de cœurs.
  Cible : workloads mono-thread critiques (bases de données, trading, CAO).
  Le 9175F (16 cœurs, 512 Mo de L3, 5,0 GHz) est l'arme anti-latence du
  catalogue.
- **P (mono-socket)** : 9655P, 9555P, 9455P, 9355P. Fonctionnent **uniquement
  en 1P** (pas de liens inter-socket activés), prix réduit de 8 à 20 % par
  rapport à l'équivalent 2P. Si vous êtes sûr de rester en mono-socket,
  c'est de l'argent gratuit — voir section 32.
- **Séries X (3D V-Cache)** : existent sur les générations précédentes
  (ex. 9684X Genoa-X, 1,15 Go de L3). Aucune référence X annoncée pour
  Turin au lancement (vérifié le 27/09/2026) ; Venice-X est annoncé pour
  2027 avec jusqu'à 1 152 Mo de L3 (voir section 93).
- **Série 8004 « Siena » (SP6)** : entrée de gamme 1P, jusqu'à 64 cœurs
  Zen 4c, TDP contenus. Pour l'edge et le stockage, pas pour le datacenter
  dense. Non détaillée ici (hors périmètre, gamme 2023).

Règle d'achat : **ne payez jamais le prix fort d'une référence 2P si votre
architecture est figée en 1P** — la version P fait le même travail pour moins
cher, et les licences par socket (section 25) ne changent pas.

## 8. Sockets SP5 et SP6 : différences physiques et d'usage

| Caractéristique | SP5 | SP6 |
|---|---|---|
| Format | LGA 6096 contacts | LGA 4844 contacts |
| Gammes | EPYC 9004/9005 (Genoa, Bergamo, Turin) | EPYC 8004 (Siena) |
| Canaux mémoire | 12 × DDR5 | 6 × DDR5 |
| PCIe | 128 lignes Gen5 | 96 lignes Gen5 |
| TDP max supporté | 500 W (Turin) | ~225 W |
| Sockets par carte mère | 1 ou 2 | 1 uniquement |
| Cible | Datacenter, HPC, cloud | Edge, telco, stockage froid |

Conséquences pratiques :
- SP5 2P = 24 canaux mémoire et 2×192 cœurs max par serveur. C'est la
  plateforme de référence de ce guide.
- Le passage Genoa → Turin sur carte SP5 existante est possible **si** l'OEM
  fournit le BIOS/microcode adéquat et si le circuit d'alimentation (VRM) et
  le refroidissement supportent le TDP cible (500 W !). Vérifiez la QVL de
  l'OEM avant d'acheter des CPU seuls (section 104).
- SP6 ne concerne pas les workloads datacenter denses : 6 canaux mémoire
  seulement, pas de 2P. Ne comparez jamais un prix « EPYC » SP6 avec un
  besoin SP5.

## 9. TDP, cTDP, PPT : le vocabulaire puissance d'AMD

- **TDP (Thermal Design Power)** : enveloppe thermique de référence. Un EPYC
  9965 « 500 W » est conçu pour dissiper 500 W en charge soutenue. C'est la
  valeur à utiliser pour dimensionner le refroidissement (partie F).
- **cTDP (configurable TDP)** : la plupart des EPYC acceptent une plage de
  réglage (ex. 320-500 W). Baisser le cTDP de 20 % ne fait perdre que ~5-10 %
  de performance sur beaucoup de workloads : c'est un levier d'efficacité
  énergétique direct (voir section 83).
- **PPT (Package Power Tracking)** : limite réelle mesurée au socket, pilotée
  par le firmware. Le CPU peut dépasser brièvement le TDP en boost (jusqu'à
  la limite PPT), d'où l'importance des alimentations avec marge (section 84).
- **« Default CPU Power »** : nouvelle terminologie AMD introduite avec
  Venice (le 9996 est donné pour 600 W — vérifié le 27/09/2026). Ne confondez
  pas les générations dans un comparatif.

Règle d'or énergie : **dimensionnez le refroidissement et l'électrique sur
le PPT max / TDP haut de plage, pas sur la consommation moyenne.** La moyenne
sert à la facture, le max sert à ne pas disjoncter.

## 10. Choisir son EPYC par workload : table de décision

| Workload | Variante | Références types (indicatif) | Critère n°1 |
|---|---|---|---|
| Virtualisation dense (VMware/Proxmox) | Zen 5c | 9655, 9745, 9555 | cœurs/euro, RAM/cœur |
| Base de données OLTP | Zen 5 / F | 9175F, 9375F, 9575F | fréquence + L3 |
| Base de données analytique | Zen 5 | 9755, 9655 | bande passante mémoire |
| HPC / simulation | Zen 5 | 9755, 9655 | AVX-512, BP mémoire |
| Cloud / conteneurs | Zen 5c | 9965, 9845, 9825 | cœurs/watt |
| Inference IA sur CPU | Zen 5c | 9965, 9745 | débit (AMX absent côté AMD) |
| VDI | Zen 5c | 9535, 9455 | cœurs, coût/VM |
| Stockage / SDS | Zen 5c entrée | 9335, 9255 | PCIe, prix |

Note « AMX absent côté AMD » : l'accélération matricielle IA d'Intel (AMX)
n'a pas d'équivalent direct sur EPYC ; AMD mise sur AVX-512/VNNI. Pour de
l'inférence CPU intensive, comparez toujours les deux plateformes sur
**votre** modèle (section 22).

---

## 11. Intel Xeon 6 : la logique de gamme (P-core vs E-core)

Intel a scindé sa gamme serveur en deux familles de cœurs, avec une
nomenclature unifiée « Xeon 6 » (vérifié le 27/09/2026) :

- **P-core (Performance)** : cœurs Redwood Cove, avec hyper-threading (2
  threads/cœur), fréquences élevées, AVX-512 et **AMX** (accélération
  matricielle IA). Cible : HPC, bases de données, IA, charges généralistes
  exigeantes. Noms de code : Granite Rapids (6900P, 6700P/6500P).
- **E-core (Efficient)** : cœurs Crestmont puis Darkmont, **sans**
  hyper-threading (1 thread/cœur), densité maximale, efficacité énergétique.
  Cible : cloud, scale-out, telco, microservices. Noms de code : Sierra
  Forest (6700E, 6900E), puis Clearwater Forest / Xeon 6+ (Darkmont).

Règle de lecture des références : le chiffre des milliers = la plateforme
(6700 = milieu de gamme, 6900 = haut de gamme), la lettre = le type de cœur
(P ou E). Un 6980P et un 6980E n'ont **rien** en commun architecturalement
malgré un numéro proche : l'un a 128 P-cores avec HT, l'autre 288 E-cores
sans HT (génération 6+).

## 12. Granite Rapids : Xeon 6900P (P-core haut de gamme)

Le fer de lance P-core d'Intel (lancé 09/2024, vérifié le 27/09/2026) :

