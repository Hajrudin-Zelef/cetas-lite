---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-16
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "AWS", "Intel"]
dates: ["2026-09-27"]
keywords: ["amd", "aws", "exploit", "gpu", "graviton", "intel", "latency"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [2457, 2632]
sha256: 88fb195ddcea9a5282068ba682bdb272e2ec962b64a65cb07d9ca65571b00e48
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

| Critère | Siena (SP6) | Turin (SP5) |
|---|---|---|
| Sockets | 1 uniquement | 1 ou 2 |
| Canaux mémoire | 6 | 12 |
| PCIe | 96 Gen5 | 128 Gen5 |
| TDP max | ~225 W | 500 W |

Quand la choisir : armoire edge, site distant, faible consommation exigée.
**Jamais** comme alternative « moins chère » à un besoin SP5 : 6 canaux et
1 socket plafonnent vite tout workload sérieux.

## 132. Xeon 6500/6300 et variantes SoC : le bas de la gamme

Intel décline le Xeon 6 vers le bas (vérifié le 27/09/2026 dans leurs
principes) :
- **Xeon 6500 P-core** : entrée de gamme serveur, 8 canaux, TDP contenus —
  le concurrent des petits EPYC pour la PME.
- **Variantes SoC** : Xeon 6 en version système-sur-puce pour le réseau
  et l'edge (fonctions réseau intégrées, enveloppe réduite).
- **Xeon 6300** : bas de gamme, petits serveurs et appliances.

Règle : en dessous de ~16-24 cœurs utiles, comparez aussi les **Core/EPYC
4000** et les offres reconditionnées : le « petit Xeon » neuf est rarement
le moins cher au TCO sur ce segment.

## 133. cTDP en pratique : le levier d'efficacité gratuit

La plupart des EPYC et Xeon acceptent un TDP configurable (cTDP). Exemple
de l'effet (ordres de grandeur, à mesurer sur votre workload) :

| Réglage | Perf. relative (indicatif) | Conso relative |
|---|---|---|
| TDP nominal (400 W) | 100 % | 100 % |
| cTDP -15 % (340 W) | 95-98 % | ~85 % |
| cTDP -25 % (300 W) | 90-95 % | ~75 % |
| cTDP -40 % (240 W) | 80-88 % | ~60 % |

Sur des workloads virtualisation à 30-40 % de charge, baisser le cTDP de
20 % est **quasi invisible en performance** et économise ~15 % d'énergie
24/7 — plus le froid associé. C'est un réglage BIOS gratuit : faites-en un
standard de votre build, avec une exception documentée pour le HPC.

## 134. AVX-512 : pourquoi ça change tout en HPC/IA

AVX-512 traite 8 doubles (ou 16 floats) par cycle et par cœur, contre 4/8
en AVX2. Sur Turin et Xeon 6, le chemin de données est **512 bits complet**
(parité enfin atteinte côté AMD, qui émulait en 2×256 bits sur Zen 4).

| Workload | Gain typique AVX-512 vs AVX2 (indicatif) |
|---|---|
| CFD / FEA | 1,3-1,8× |
| Encodage vidéo | 1,2-1,5× |
| Inférence INT8 (VNNI) | 1,5-2× |
| Code scalaire | 1,0× (aucun) |

Conditions : le logiciel doit être compilé avec les bons flags (-mavx512…,
librairies MKL/OpenBLAS). **Un code non vectorisé ne voit rien.** Avant
d'acheter « pour l'AVX-512 », vérifiez que votre stack l'exploite —
sinon vous payez du silicium inutile.

## 135. AMX vs AVX-512/VNNI : le duel de l'inférence CPU

| Critère | Intel AMX | AMD AVX-512 + VNNI |
|---|---|---|
| Principe | tuiles matricielles dédiées | vecteurs larges généralistes |
| Débit INT8/BF16 | très élevé (accélérateur) | élevé (vectoriel) |
| Flexibilité | opérations matricielles | tout code vectorisable |
| Support logiciel | frameworks optimisés Intel | standard, large |

Verdict 2026 (vérifié dans les principes au 27/09/2026) : pour de
l'inférence de modèles quantifiés **sur CPU**, AMX donne l'avantage à
Intel — parfois 2× ou plus à cœurs comparables (chiffres à relativiser,
section 2). Mais dès qu'un GPU entre dans l'équation, le débat CPU
devient secondaire : un GPU modeste bat les deux.

## 136. Prix par cœur et par thread : lecture économique

À partir des tarifs publics au lancement (vérifié le 27/09/2026) :

| CPU | Prix | $/cœur | $/thread |
|---|---|---|---|
| EPYC 9965 (192 c./384 th.) | 14 813 $ | ~77 $ | ~39 $ |
| EPYC 9755 (128 c./256 th.) | 12 984 $ | ~101 $ | ~51 $ |
| EPYC 9175F (16 c./32 th.) | 4 256 $ | ~266 $ | ~133 $ |
| Xeon 6990E+ (288 c./288 th.) | 14 995 $ | ~52 $ | ~52 $ |
| Xeon 6980P (128 c./256 th.) | 17 800 $ | ~139 $ | ~70 $ |

Lecture : le **$/cœur** favorise les denses (E-core, Zen 5c) ; le
**$/thread** favorise le SMT d'AMD. Mais ni l'un ni l'autre ne prédit la
performance : un cœur « F » à 266 $/cœur peut être 3× plus rentable qu'un
cœur dense à 52 $/cœur sur un workload sous licence au cœur (section 25).
**Le seul ratio qui compte est $/unité de votre workload.**

## 137. Serveurs ARM : le contexte à connaître (bref)

Hors x86, les CPU ARM serveur (Ampere, Graviton d'AWS) gagnent du terrain
dans le cloud sur le ratio perf/watt des workloads scale-out. En 2026,
ils restent **marginaux dans l'entreprise européenne** (écosystème
logiciel, habitudes d'achat).

Positionnement : ne les évaluez que si (1) vous êtes 100 % Linux
conteneurisé, (2) votre fournisseur cloud les propose moins cher, et
(3) vous avez mesuré le portage. Pour tout le reste (Windows, VMware,
bases propriétaires), x86 reste la réponse. Ce guide ne détaille pas
l'ARM : périmètre assumé.

## 138. Reconditionné : la 4e option d'achat

Un serveur Genoa (EPYC 9004) reconditionné avec garantie 1-3 ans coûte
**40 à 60 % moins cher** qu'un Turin neuf (indicatif) pour 70-80 % de la
performance. Cas où c'est pertinent :
- Environnements de dev/test, PRA, edge.
- Budgets contraints avec besoin de cœurs immédiat.
- Extension d'un parc Genoa existant (homogénéité).

Vérifications : heures de fonctionnement (BMC), état des batteries RAID
et des ventilateurs, **mise à jour BIOS/BMC incluse**, garantie écrite
avec SLA. Évitez le reconditionné pour : la production critique sans
redondance, tout ce qui exige les dernières instructions (AMX récent,
CXL).

## 139. Négociation OEM : ce qui se négocie vraiment

Le prix catalogue CPU (sections 6, 12) n'est qu'un point de départ. Leviers
de négociation réels (indicatif, varie selon volume et OEM) :
- **Remises volume** : 20-50 % sur CPU et RAM dès quelques dizaines de
  serveurs.
- **RAM et disques** : marges OEM élevées — faites chiffrer la RAM
  séparément et comparez au prix des barrettes QVL seules.
- **Garantie** : passer de 3 à 5 ans coûte moins cher à l'achat qu'en
  extension ultérieure.
- **Services** : installation, mise en rack, reprise de l'ancien parc —
  négociables en pack.
- **Fin de trimestre** : les commerciaux ont des objectifs ; un devis signé
  fin de trimestre vaut 5-15 % de plus (pratique de terrain).

Toujours demander : **3 devis minimum** (2 OEM + 1 intégrateur), à
configuration strictement identique, TCO 5 ans inclus (section 101).

## 140. Erreurs de sizing vues en audit (retours terrain)

1. **Le « CPU trop gros, RAM trop petite »** : 2× 96 cœurs avec 256 Go de
   RAM — les cœurs s'ennuient, les VM manquent de mémoire. Ratio ignoré
   (section 48).
2. **Le « tout en 2 DPC »** : 1,5 To installés à 4800 MT/s alors que le
   workload voulait du débit — 12× 128 Go en 1 DPC auraient tout changé.
3. **Le « 1U par principe »** : 2× 400 W dans 1U, throttling l'été,
   remplacement des ventilateurs tous les ans.
4. **Le « pas de PDU monitorée »** : 3 ans de factures estimées au doigt
   mouillé, impossible d'optimiser.
5. **Le « onduleur hérité »** : baie doublée, onduleur inchangé → surcharge
   à 95 %, batteries mortes en 2 ans.
6. **Le « BIOS d'origine »** : 4 ans sans mise à jour, failles BMC
   exploitées, courbes de ventilation d'un autre âge.

---

## 141. Timings et latence CAS : ce qui compte vraiment

On lit « CL40 » sur une DDR5-5600 et « CL46 » sur une DDR5-6400 : la latence
**en cycles** augmente avec la fréquence. La vraie latence (en ns) :

```
 latence_ns = CL x 2000 / frequence_MT/s

 DDR5-5600 CL46 : 46 x 2000 / 5600 = 16,4 ns
 DDR5-6400 CL52 : 52 x 2000 / 6400 = 16,25 ns
```

Conclusion : **à génération égale, la latence réelle varie peu** avec la
fréquence. Le choix 5600 vs 6400 se fait sur la **bande passante**, pas sur
la latence. Et les timings agressifs (« low latency ») du monde desktop
n'existent pas en serveur : la stabilité et l'ECC priment.

## 142. Bande passante : calculs détaillés par plateforme

