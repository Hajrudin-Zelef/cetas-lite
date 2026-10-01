---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-27
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["arr", "capex", "cpo", "dci", "gpu", "hyperscaler", "incident", "lpo"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [3908, 4025]
sha256: 774e770b438e808ab3cc0b5dc5ebbc2771d350d12d7aa794bfbaec603f0336a4
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

| # | Si | Alors |
|---|---|---|
| 1 | Cluster IA > 128 GPU | Backend dédié 1:1, rail-optimized |
| 2 | Lien < 2 m | DAC passif |
| 3 | Lien 2-5 m | AEC (actif) |
| 4 | Lien 5-30 m | AOC |
| 5 | Lien > 30 m | Transceiver + fibre |
| 6 | Rocade neuve | OS2, pas de débat |
| 7 | Baie ↔ brassage (< 100 m) | OM4 SR si ≤400G, OS2 DR sinon |
| 8 | Budget optique < 3 dB | Revoir le design (retirer un brassage) |
| 9 | BER pré-FEC > 1e-9 | Nettoyer, puis mesurer, puis changer |
| 10 | FEC uncorrectable > 0 | Arrêter la charge, changer le lien |
| 11 | NIC 400G | PCIe Gen5 x16 minimum |
| 12 | NIC 800G | PCIe Gen6 x16 minimum |
| 13 | Pas de PCIe Gen6 | 800G bridé → rester en 400G |
| 14 | RoCEv2 en prod | ECN + PFC par file + MTU 9000 |
| 15 | Pas de lossless | Pas de RDMA |
| 16 | UEC disponible | Exiger UEC-ready, garder RoCEv2 en 2026 |
| 17 | InfiniBand existant | Le garder s'il marche ; UEC pour le neuf |
| 18 | Choix du NOS | Le même 10 ans ; pas de changement pour le prix |
| 19 | Whitebox | Seulement avec une équipe SONiC |
| 20 | Optiques OEM trop chères | Tierces codées, validées 1 réf./switch |
| 21 | Prix non public | Fourchette ⚠️ + 3 devis |
| 22 | Devis unique | Refuser, toujours 3 |
| 23 | Clause CPO absente | La négocier avant signature |
| 24 | FW « latest » proposé | Refuser, version figée validée |
| 25 | Pas de protocole d'acceptation | Ne pas signer |
| 26 | Tests < 72 h | Insuffisant pour un fabric IA |
| 27 | Pas de test ECMP | Polarisation garantie en prod |
| 28 | Pas de test de panne spine | Le faire au lab avant la prod |
| 29 | Supervision sans BER | Aveugle sur la santé des liens |
| 30 | Alertes non testées | En faire sonner une pour de faux |
| 31 | Config non versionnée | Git dès le jour 1 |
| 32 | Changement sans rollback | Interdit |
| 33 | Deux changements à la fois | Jamais |
| 34 | Incident P1/P2 | Rapport 5 pourquoi sous 5 jours |
| 35 | Pas de spares | MTTR ×10 |
| 36 | Spares < 10 % optiques | Augmenter |
| 37 | Switch critique sans spare | 1 par plan critique |
| 38 | MPO sans bouchons | Poussière = BER dans 6 mois |
| 39 | Connecteur tombé | Nettoyer avant de rebrancher |
| 40 | Inspection sans microscope | On ne voit pas la saleté à l'œil |
| 41 | Goulotte > 50 % | Nouvelle goulotte |
| 42 | Fibre pliée à 90° | Remplacer |
| 43 | DAC > portée | AOC ou transceiver |
| 44 | Tx vers Tx | Croiser les fibres |
| 45 | MPO mâle-mâle | Ne jamais forcer |
| 46 | Deux trunks méthode B | Un seul élément croisé par lien |
| 47 | Étiquette manquante | Étiqueter avant de brasser |
| 48 | Pas de photo du brassage | La prendre maintenant |
| 49 | Plan d'adressage absent | NetBox avant le premier câble |
| 50 | IPv6 non prévu | Prévoir le plan, activer plus tard |
| 51 | MTU incohérent | 9000 partout ou 1500 partout |
| 52 | PFC global | Par file RoCE uniquement |
| 53 | PFC storm | Désactiver PFC, vérifier ECN, chercher la boucle |
| 54 | ECN mal réglé | Seuils 30-50 % / 70-80 %, tuner au lab |
| 55 | Deux CC mélangés | Un seul algorithme par fabric |
| 56 | NCCL lent | NCCL_DEBUG=INFO, quelle interface ? |
| 57 | NCCL utilise eth0 | Forcer l'interface RDMA |
| 58 | GDRCopy absent | L'installer pour les petits messages |
| 59 | PCIe bifurcation fausse | x16 ou x8+x8 selon la carte |
| 60 | NUMA ignoré | Pinner les IRQs au bon nœud |
| 61 | IRQ sur CPU 0 | Les répartir (irqbalance ou manuel) |
| 62 | Offloads désactivés | Les activer sauf debug |
| 63 | Buffer leaf trop petit | Leaf shallow + spine deep |
| 64 | Microbursts invisibles | LANZ / streaming, pas le polling |
| 65 | Latence p99 élevée | Chercher la file, pas le lien |
| 66 | Un flux = un lien | Normal en ECMP ; spraying pour l'IA |
| 67 | Oversubscription 3:1 en IA | Non : 1:1 backend |
| 68 | 400G ou 800G en 2026 | 800G (coût/Gb/s inférieur) |
| 69 | 100G en 2026 | Front-end et edge, pas le backend IA |
| 70 | 1,6T en 2026 | Attendre, préparer la fibre |
| 71 | CPO en 2026 | Pilote seulement, pas la prod critique |
| 72 | LPO en 2026 | Valider au cas par cas |
| 73 | DPU pour du web | Non, surcoût inutile |
| 74 | DPU pour l'IA/edge | Oui : offload, isolation, zéro-trust |
| 75 | Chiffrement soft du backend | Non, tue la perf ; HW si exigé |
| 76 | DCI non chiffré | Chiffrer (MACsec/OTN) |
| 77 | Edge < 10 serveurs | Switch simple, pas de fabric |
| 78 | CDN : build ou buy | Buy (sortie) sauf hyperscaler |
| 79 | PUE > 1,5 | Auditer le refroidissement |
| 80 | Optiques = 30 % du thermique réseau | Les compter dans le dimensionnement |
| 81 | Switch 800G tout-optique | Prévoir 1,5-2 kW et l'airflow |
| 82 | Armoire sans allée froide | Ne pas y mettre du 800G |
| 83 | UPS réseau = UPS serveurs | Séparer, réseau prioritaire |
| 84 | Pas de test de coupure | Le faire une fois par an |
| 85 | Canicule | Seuils thermiques vérifiés avant |
| 86 | Équipe d'une personne | Doc + automation + astreinte claire |
| 87 | Formation absente | Lab + 1 jour de casse volontaire |
| 88 | Design non relu | 20 questions du §186 |
| 89 | RFP au seul CAPEX | TCO 5 ans obligatoire |
| 90 | Fournisseur unique | Clause de sortie au contrat |
| 91 | Support J+1 sur l'IA | 4 h 24/7 exigé |
| 92 | TAC sans nom | Exiger un contact nommé |
| 93 | Fin de vie | Wipe + DEEE + inventaire |
| 94 | Revente | Effacer licences et configs |
| 95 | Congés | Checklist §173 |
| 96 | REX mensuel | 15 min, ce qui a coincé |
| 97 | Guide non relu à 6 mois | Les ✅ bougent vite (CPO, UEC, 1,6T) |
| 98 | Donnée incertaine | Marquer ⚠️, ne pas affirmer |
| 99 | Référence introuvable | « Non trouvée au 27/09/2026 » |
| 100 | Doute sur un chiffre | Mesurer avant de changer |
| 101 | Pression du vendeur | 3 devis + lab + refs clients |
| 102 | Hype du moment | Pilote d'abord, prod ensuite |
| 103 | « Ça marche chez nous » | Le prouver au lab avec VOS workloads |
| 104 | Doc du vendeur vs terrain | Le terrain gagne toujours |
| 105 | Réunion sans décision | La suivante avec les chiffres |
| 106 | Chiffre sans source | Exiger la datasheet |
| 107 | Datasheet absente | Pas d'achat |
| 108 | Module sans SN | Refuser le lot |
| 109 | Lot > 2 % rejet | Lot refusé |
| 110 | Le réseau est lent | Arbre de dépannage §143 |

---

*Fin du guide — **187 sections numérotées**, 110 décisions rapides,
80 termes de glossaire, 45 pièges, quiz, 50 FAQ, 100 acronymes,
BOMs, checklists, templates, index. Objectif : ≥ 4000 lignes (`wc -l`).*
