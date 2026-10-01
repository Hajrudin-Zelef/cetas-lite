---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-12
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Samsung", "TSMC"]
dates: ["2026-07-23", "2026-09-27"]
keywords: ["18a", "amd", "benchmarks", "capex", "clearwater forest", "dram", "gpu", "intel", "lpddr5x", "throughput"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [1803, 1973]
sha256: c3e7d749e9c37f5c115a4b8a8296c3e691c763a006525922385dca4f8cd0f6dc
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

| Solution | Capacité typique | Usage |
|---|---|---|
| Climatiseur détente directe | 5-15 kW/unité | petites salles |
| Eau glacée + CTA | 20-200 kW | datacenters |
| Free cooling (air/eau) | selon climat | réduit la facture de froid de 30-70 % |
| Confinement allées | — | +15-25 % d'efficacité sur l'existant |

Repères de conception :
- Température de consigne : **27 °C** (ASHRAE A2) au lieu de 22 °C =
  ~4-5 % d'économie de froid par degré, sans risque pour du matériel
  récent (vérifier la classe ASHRAE des serveurs).
- La chaleur fatale peut être **valorisée** (chauffage de bureaux, eau
  chaude) surtout avec du DLC à eau tiède (section 65) : à 40 °C, c'est
  directement utilisable.
- **Comptez le froid dans le TCO** (section 101) : à PUE 1,5, chaque
  kW IT « coûte » 0,5 kW de froid en énergie.

---

## 93. AMD EPYC 9006 « Venice » : annonces vérifiées au 27/09/2026

Annoncé à l'événement Advancing AI les 22-23/07/2026, production de masse
engagée sur TSMC 2 nm (vérifié le 27/09/2026) :

| Caractéristique | EPYC 9006 « Venice » |
|---|---|
| Architecture | Zen 6 / Zen 6c |
| Gravure | TSMC 2 nm (N2, GAA nanosheets) — 1er CPU HPC en 2 nm |
| Cœurs max | **256** Zen 6c / 512 threads (8 CCD × 32 cœurs) |
| Zen 6 classique | jusqu'à 96 cœurs, boost 5,0 GHz |
| Socket | **SP7** (nouveau, incompatible SP5) |
| Canaux mémoire | **16** × DDR5-8000 |
| MRDIMM | Gen2 jusqu'à **12 800 MT/s** |
| Bande passante crête | ~1,6 To/s (contre 614 Go/s sur Turin) |
| PCIe | **Gen6** (64 GT/s/ligne), 128 lignes |
| CXL | 3.1 |
| Puissance | jusqu'à **600 W** (« Default CPU Power », réf. 9996) |
| Transistors | ~203 milliards |
| 4 gammes | SP7 standard, SP8 (edge, 8-128 c.), 9006X (3D V-Cache, jusqu'à 1 152 Mo de L3), 9006 LP (LPDDR5X, modules SOCAMM2) |
| Disponibilité | production ramp 05/2026, systèmes **T4 2026** (presse, non confirmé par AMD au 27/09/2026) |
| Prix | **non publiés au 27/09/2026** |

Chiffres constructeur : +70 % throughput vs Turin, 2,2× l'inférence vs la
plateforme concurrente citée (chiffres AMD — à relativiser, section 2).
**Stratégie d'achat** : si votre renouvellement est prévu fin 2026/début
2027 et que vous visez le haut de gamme, attendez les prix et les tests
indépendants de Venice avant de signer du Turin au prix fort.

## 94. Intel Diamond Rapids / Xeon 7 : annonces vérifiées au 27/09/2026

Confirmé par Intel au Computex 2026 (vérifié le 27/09/2026) :

| Caractéristique | Diamond Rapids (Xeon 7) |
|---|---|
| Lancement | **2027** (confirmé par Intel) |
| Gravure | Intel 18A-P (version améliorée du 18A) |
| Type de cœurs | P-core nouvelle génération |
| PCIe | **Gen6** |
| Mémoire | **bande passante doublée** vs Xeon 6 (dixit Intel) |
| Hyper-threading | non confirmé au 27/09/2026 |

Ce qu'on ne sait pas (non trouvé au 27/09/2026) : nombre de cœurs final,
TDP, prix, socket. **Ne planifiez pas d'achat 2026 sur des specs Diamond
Rapids** : ce sont des annonces de direction, pas des fiches techniques.

## 95. DDR5 → DDR6 : état de la norme (vérifié le 27/09/2026)

| Point | Statut au 27/09/2026 |
|---|---|
| Spécification JEDEC | **non finalisée** (brouillon fin 2024, détails en cours) |
| Débits visés | 8 400 à 17 600 MT/s selon maturité |
| Production de masse | **2028-2029** (Samsung, SK Hynix, Micron en développement) |
| Tension | < 1,1 V visé (LPDDR6) |
| DDR5 en serveurs | > 80 % des livraisons DRAM serveur (2026) |

Conséquence pratique : **tout achat serveur 2026-2027 se fait en DDR5.**
La DDR6 n'est pas un critère d'achat aujourd'hui, c'est un critère de
**trajectoire** : préférez les plateformes dont le constructeur a annoncé
le support des futurs standards (MRDIMM Gen2/Gen3, CXL) plutôt que des
impasses.

## 96. MRDIMM Gen2/Gen3 : feuille de route (vérifié le 27/09/2026)

| Génération | Débit | Horizon (vérifié) |
|---|---|---|
| Gen1 (8800) | 8 800 MT/s | en production (Xeon 6900P) |
| Gen2 | 12 800 MT/s | systèmes **T1 2027** (Intel), supporté par Venice |
| Gen3 | 17 600 MT/s | ~2030 |

Le MRDIMM est la réponse de l'industrie à l'attente de la DDR6 : **des
débits de classe DDR6 sans changer de slot**. Pour un acheteur 2026 :
si votre workload est gourmand en bande passante et que vous gardez vos
serveurs 5 ans, une plateforme MRDIMM-ready (Xeon 6900, SP7 Venice)
protège votre investissement mémoire mieux qu'une plateforme DDR5
plafonnée à 6400.

## 97. CXL 2.0/3.x : état du déploiement (vérifié le 27/09/2026)

- **CXL 2.0** : en production (Microsoft Azure, premier déploiement cloud
  11/2025). Pooling mémoire entre hôtes : réel, mais réservé aux
  hyperscalers pour l'instant.
- **CXL 3.0/3.1** : silicium de switch en échantillonnage (2025-2026).
  Pas de déploiement de production généralisé au 27/09/2026.
- **CXL 4.0** : attendu avec PCIe 7.0, au stade IP uniquement.

Pour l'acheteur : exigez des CPU/BIOS **CXL-ready** (c'est le cas de Turin,
Xeon 6, Venice), mais n'achetez du matériel CXL (cartes d'extension
mémoire, switchs) que sur besoin avéré et après validation de votre
stack OS/hyperviseur. Le CXL mal supporté par l'OS = de la RAM lente et
chère.

## 98. PCIe 6.0 et 7.0 : repères

| Génération | Débit/ligne | Disponible sur CPU serveur (vérifié) |
|---|---|---|
| PCIe 5.0 | 32 GT/s | Turin, Xeon 6, Xeon 6+ (standard actuel) |
| PCIe 6.0 | 64 GT/s | **Venice (fin 2026)**, Diamond Rapids (2027) |
| PCIe 7.0 | 128 GT/s | spécification en cours, pas de CPU annoncé |

Impact pratique : le PCIe 6.0 double la bande passante vers les GPU et
les NIC 400/800 GbE. Si vous achetez des serveurs hôtes GPU fin 2026,
**Venice SP7 (PCIe 6.0 natif) mérite d'être évalué** contre une plateforme
PCIe 5.0 qui bridera les GPU de prochaine génération.

## 99. Ce qui n'a PAS été trouvé au 27/09/2026

Par honnêteté, voici ce que la recherche n'a pas permis de vérifier :
- Prix des EPYC 9006 Venice (non publiés par AMD).
- Spécifications détaillées des SKU Venice (noms, TDP par référence).
- Specs finales de Diamond Rapids (cœurs, TDP, socket).
- Prix publics des barrettes MRDIMM Gen2 et des DIMM DDR5-8000.
- Benchmarks indépendants du Xeon 6+ Clearwater Forest (tests en cours).
- Date de finalisation JEDEC de la DDR6.

**Si un commercial vous cite ces chiffres, demandez la source écrite.**

## 100. Stratégie d'achat 2026-2027 : conseils

| Situation | Recommandation |
|---|---|
| Renouvellement immédiat (T4 2026) | Turin SP5 ou Xeon 6 : plateformes matures, prix négociables |
| Haut de gamme, pas d'urgence | Attendre prix/tests Venice (T4 2026/T1 2027) |
| Scale-out cloud | Xeon 6+ Clearwater Forest à évaluer (tests indépendants à suivre) |
| HPC/IA | Xeon 6900P + MRDIMM aujourd'hui ; Venice + MRDIMM-12800 en 2027 |
| Budget serré | Génération précédente reconditionnée (Genoa) : -40 à -60 % (indicatif) |
| Salle existante saturée | Ne pas acheter sans le calcul de la section 78 |

Règle anti-obsolescence : sur un cycle de 5 ans, **la plateforme (socket,
canaux, PCIe) compte plus que le CPU du jour**. Un SP5 acheté en 2024
accepte un Turin en upgrade ; un SP7 acheté fin 2026 suivra la même logique
pour la génération suivante.

## 101. TCO : méthode de calcul complète (5 ans)

```
TCO_5ans = CAPEX + OPEX_5ans

CAPEX = serveurs + RAM + stockage + reseau + rails/PDU + onduleur (quote-part)
OPEX_annuel = energie + froid (via PUE) + licences + maintenance + admin

Detail energie : P_moy x 8760 x PUE x prix_kWh  (section 83)
```

Checklist des postes (ceux qu'on oublie en gras) :
- **Licences logicielles** (souvent > matériel sur 5 ans, section 25).
- **Froid** (via PUE — pas gratuit).
- **Batteries onduleur** (remplacement tous les 4-5 ans).
- **Administration** (temps humain, supervision, sauvegardes).
- **Fin de vie** (revente, destruction sécurisée des disques).
- **Coût de la panne** (indisponibilité × coût horaire métier).

