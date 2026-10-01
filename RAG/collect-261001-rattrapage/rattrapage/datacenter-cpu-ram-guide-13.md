---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-13
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Oracle"]
dates: []
keywords: ["arr", "benchmark", "benchmarks", "capex", "gpu", "inference", "mlperf"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [1974, 2139]
sha256: ef62aeb48c29ee0f7b16deef768b76db503384ded296d55ba64d9029b699f439
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

## 102. Exemple TCO : 1 serveur 2S sur 5 ans (indicatif)

Serveur 2U 2S EPYC 9655, 768 Go RAM, 8 NVMe (section 31) :

| Poste | Montant 5 ans (indicatif) |
|---|---|
| Serveur complet (CAPEX) | ~25 000 € |
| Énergie (1100 W max, 30 % charge, PUE 1,5, 0,18 €/kWh) | ~7 200 € |
| Licences (ex. virtualisation, par cœur) | 10 000-40 000 € (selon éditeur !) |
| Maintenance (contrat J+1) | ~4 000 € |
| Quote-part onduleur + PDU | ~1 500 € |
| **Total** | **~48 000 à 78 000 €** |

Le serveur à 25 000 € coûte en réalité **2 à 3× son prix** sur 5 ans.
Présentez toujours le TCO, jamais le CAPEX seul : c'est ce chiffre qui
fait choisir entre deux offres à 3 000 € d'écart.

## 103. Benchmarks : comment les lire sans se faire piéger

| Benchmark | Mesure | Pertinence |
|---|---|---|
| SPEC CPU 2017 | calcul généraliste | bonne base comparative |
| SPECvirt | virtualisation | proche du réel (VM) |
| STREAM / HPCG | bande passante mémoire | HPC, DB analytique |
| TPC-C / TPC-H | bases de données | OLTP / décisionnel |
| MLPerf Inference | inférence IA | choisir CPU vs GPU |

Règles de lecture :
1. **Même benchmark, même version, même OS** : sinon la comparaison est
   nulle.
2. Les chiffres constructeur sont des **maxima** sur config optimale :
   retirez 10-20 % pour votre réalité (moins de RAM, autre OS).
3. Un benchmark ne remplace pas un **POC** : 2 semaines de test avec votre
   workload valent mieux que 200 slides.
4. Méfiez-vous des comparaisons « cœur à cœur » entre architectures
   différentes (SMT vs non-SMT, P-core vs E-core) : comparez au **prix**
   et au **watt**, pas au cœur.

## 104. Firmwares et microcodes : l'hygiène invisible

- **BIOS/BMC à jour** avant mise en production : les premières versions
  ont des courbes de ventilation et des tables de puissance perfectibles.
  Les mises à jour corrigent aussi des failles (BMC exposé = porte
  d'entrée).
- **Microcode CPU** : livré via BIOS ou OS. Les failles type Spectre/
  Meltdown se mitigent par microcode **au prix de quelques % de
  performance** : mesurez avant/après sur vos workloads sensibles.
- **Upgrade CPU sur carte existante** (Genoa → Turin) : BIOS minimum
  requis, parfois nouveau dissipateur (TDP supérieur), parfois nouveau
  profil de ventilation. Vérifiez la **matrice de compatibilité OEM**
  ligne par ligne.
- **Ne mettez jamais à jour tous vos nœuds en même temps** : vague 1 =
  1 nœud pilote, validation 1 semaine, puis le reste. Un BIOS défectueux
  sur tout un cluster est une panne totale évitable.

---

## 105. Piège n°1 : acheter des cœurs au lieu d'acheter du débit

Le réflexe « plus de cœurs = plus rapide » est faux dès que la bande
passante mémoire sature. Un EPYC 9965 (192 cœurs, 3,2 Go/s/cœur) sur un
workload HPC affamé en mémoire sera **battu** par un 9755 (128 cœurs,
4,8 Go/s/cœur) moins cher. **Dimensionnez toujours le ratio BP/cœur
(section 30) avant de signer.** Le bon CPU est celui qui nourrit ses
cœurs, pas celui qui en a le plus.

## 106. Piège n°2 : oublier les licences dans le choix du CPU

Une base Oracle sur 2× 64 cœurs coûte **2× plus cher en licences** que
sur 2× 32 cœurs (facteur 0,5 x86), pour un gain de performance souvent
inférieur à 50 %. Faites chiffrer les licences **avant** de choisir le
CPU, pas après. Dans 30 % des projets, le « petit » CPU fréquence (F)
est le choix le moins cher au TCO (section 32).

## 107. Piège n°3 : laisser des canaux mémoire vides

6× 128 Go au lieu de 12× 64 Go = même capacité, **moitié moins de bande
passante** (section 46). Le commercial vous vend de la capacité ; le
workload a besoin de débit. Règle absolue : **tous les canaux remplis,
barrettes identiques**, quitte à prendre une capacité unitaire plus petite.

## 108. Piège n°4 : mélanger les barrettes

Ajouter 6 mois plus tard des barrettes « presque identiques » (même
capacité, autre rang, autre fabricant de puces) : le contrôleur s'aligne
sur le **plus petit dénominateur commun** (fréquence la plus basse, timings
les plus lents) et peut désactiver l'entrelacement optimal. **Achetez toute
la RAM d'un serveur en une fois**, même référence, même lot si possible.
L'extension ultérieure se fait en remplaçant, pas en complétant.

## 109. Piège n°5 : le 2 DPC qui tue la fréquence

Passer de 12 à 24 barrettes pour doubler la capacité fait chuter la
fréquence (ex. 6400 → 5200 MT/s, section 47). Sur un workload sensible à
la bande passante, vous payez plus cher pour aller moins vite. **Comparez
toujours 12× grosse barrette en 1 DPC vs 24× petite en 2 DPC** au prix
**et** au débit effectif avant de choisir.

## 110. Piège n°6 : le serveur sans redondance ventilateur

Certaines configurations 1U extrêmes (2× 500 W) n'ont pas de redondance
N+1 (section 59). Un module qui lâche un vendredi soir = throttling tout
le week-end ou arrêt thermique. **Exigez la confirmation écrite de la
redondance N+1 sur votre configuration exacte**, pas sur la fiche générique
du châssis.

## 111. Piège n°7 : le sens du flux d'air inversé

Après une maintenance, un module ventilateur remonté à l'envers ou un
serveur installé face arrière devant (ça arrive !) inverse le flux :
l'air chaud est réaspiré, les sondes ne voient rien d'anormal au début,
puis les CPU throttlent. **Après toute intervention physique, vérifiez le
sens du flux** (feuille de papier devant la grille : elle doit être
aspirée à l'avant, repoussée à l'arrière).

## 112. Piège n°8 : les caches de baie oubliés

Chaque U vide sans cache (blanking panel) est un court-circuit aéraulique :
l'air frais contourne les serveurs (section 63). Dans une baie à moitié
vide sans caches, les serveurs peuvent prendre **10 °C de plus** à l'entrée.
Des caches coûtent quelques euros pièce : c'est le meilleur investissement
thermique qui existe.

## 113. Piège n°9 : dimensionner l'onduleur sur la moyenne

Un onduleur dimensionné sur la consommation moyenne disjoncte au
**démarrage simultané** (appel de courant) ou au pic de charge. Résultat :
la protection censée vous sauver provoque la panne. **Dimensionnez sur le
pic + 25 %** (section 87), et testez le démarrage simultané avant la mise
en production.

## 114. Piège n°10 : le PUE flatteur

Un PUE de 1,2 mesuré en hiver, à 30 % de charge, sans compter l'éclairage,
ne vaut rien (section 82). Exigez le **PUE annualisé à charge nominale,
périmètre complet**. Et rappelez-vous : un PUE excellent avec des serveurs
à vide allumés 24/7 reste un gâchis — **éteignez ce qui ne sert pas**
(300 W × 8760 h = 710 €/an par serveur, section 85).

---

## 115. Piège n°11 : la RAM « compatible » hors QVL

Une barrette DDR5 qui « respecte la norme » mais ne figure pas sur la QVL
(Qualified Vendor List) de l'OEM peut fonctionner… à fréquence réduite, avec
des erreurs intermittentes, ou pas du tout après une mise à jour BIOS.
**La QVL n'est pas une suggestion commerciale, c'est une matrice de
validation électrique.** En cas de panne, le support OEM commence par
vérifier la QVL : hors QVL = hors support.

## 116. Piège n°12 : le BIOS d'usine jamais mis à jour

Serveur livré, branché, jamais mis à jour : courbes de ventilation d'une
version bêta, tables de puissance incomplètes, failles BMC connues. Les
premiers mois de la vie d'une plateforme concentrent les correctifs
critiques. **Mettez à jour BIOS/BMC/microcode avant la mise en production**,
puis planifiez une fenêtre annuelle (section 104).

## 117. Piège n°13 : le RAID du pauvre sur NVMe

Mettre 8 NVMe en RAID logiciel sans vérifier l'alignement NUMA et sans
provisionner le CPU pour les interruptions : les latences s'envolent et un
socket fait tout le travail pendant que l'autre dort. **Attachez les NVMe
au socket qui les consomme** (placement NUMA, section 24) et mesurez avec
fio avant/après.

