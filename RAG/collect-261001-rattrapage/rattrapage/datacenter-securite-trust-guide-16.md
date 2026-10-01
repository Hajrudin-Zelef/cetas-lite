---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-16
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "incident", "luna"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [1943, 2070]
sha256: c942880e001f8f323ebff7733d58f55f000646bc64aa6451da56ae11f1cdf390
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

| Équipement | Puissance typique | /an (24/7, 0,20 €/kWh) | Commentaire |
|---|---|---|---|
| YubiHSM 2 | **0,1 W** (20 mA sous 5 V) | ~0,20 € | Négligeable |
| Luna PCIe HSM | **14 W** typ. / 18 W max (vérifié) | ~25 € | Dans le serveur hôte |
| Luna Network HSM | **84 W** typ. / 110 W max (vérifié) | ~150 € | + clim (~1:1) |
| Utimaco appliance (indicatif) | ~50-100 W (à vérifier) | ~90-175 € | Selon modèle |
| Serveur PKI/Vault (1U/2U) | 150-350 W | ~260-610 € | Selon charge |
| Carte FPGA Alveo V80 | **190 W** TDP (vérifié) | ~330 € | + airflow serveur |
| Serveur 2U + 2 FPGA | 700-1 100 W | ~1 230-1 930 € | Prévoir la baie en conséquence |
| Baie 42U mixte (exemple) | 5-10 kW | 8 760-17 500 € | Voir section 146 |

Constat : **le HSM lui-même consomme peu** (84 W = une ampoule LED industrielle), mais
**son indisponibilité coûte cher** : c'est un équipement à faible conso et **forte
criticité** — la catégorie qui justifie le plus la redondance électrique.

## 140. Règle d'or : jamais sur onduleur non secouru

Formulation brutale, à afficher dans le local technique :

> **Aucun équipement cryptographique critique (HSM, serveur PKI/AC, Vault, TSA, serveur
> NTP sécurisé) ne doit être branché sur un circuit non secouru par onduleur.**

Pourquoi c'est vital :

- Une coupure franche = **arrêt des signatures et des déchiffrements** (bases TDE qui ne
  redémarrent pas, Vault qui reste scellé, VPN qui ne renégocient plus).
- Pire : des **coupures répétées / micro-coupures** stressent les alimentations et les
  batteries internes des HSM (la batterie SRAM est une sécurité, pas un onduleur).
- Le redémarrage n'est **pas automatique** partout : Vault en Shamir manuel exige des
  humains (section 113), la remise en service d'un HSM après incident exige une
  procédure.

En pratique : circuit **ondulé + groupe électrogène** (autonomie onduleur ≥ temps de
démarrage du groupe + marge, typiquement 15-30 min), **test mensuel** du basculement
(voir le guide onduleurs de Zelef pour les procédures).

## 141. Continuité des services de clés : architecture électrique

```
  ┌─────────┐   ┌──────────┐   ┌───────────────┐   ┌────────────────────┐
  │ Arrivée │──▶│ TGBT +   │──▶│ Onduleur(s)   │──▶│ PDU A (secourue)  │
  │ EDF     │   │ inverseur│   │ N+1, 15-30 min│   └────────┬───────────┘
  └─────────┘   │ de source│   └───────────────┘            │
                │          │   ┌───────────────┐   ┌────────┴───────────┐
                │          │──▶│ Groupe électro│──▶│ PDU B (secourue)  │
                └──────────┘   │ (démarrage    │   └────────┬───────────┘
                               │  < 1 min)      │            │
                               └───────────────┘            ▼
                                              ┌────────────────────────┐
                                              │ HSM 1 → PDU A          │
                                              │ HSM 2 → PDU B          │
                                              │ (jamais les 2 sur la  │
                                              │  même PDU !)           │
                                              │ Serveur Vault → A+B    │
                                              │ Serveur PKI → A+B      │
                                              └────────────────────────┘
```

Règles :

- Les **2 HSM d'un cluster HA sur 2 PDU différentes** (idéalement 2 onduleurs ou 2
  chaînes) : la panne d'une chaîne n'emporte qu'un membre.
- Les serveurs Vault/PKI en **double alimentation** (A+B).
- **Ordre de redémarrage** après coupure longue, documenté et testé : réseau → HSM →
  KMS/Vault (auto-unseal) → bases/applications. Chaque étape **vérifiée** avant la
  suivante (section 138).
- Le **coffre-fort** contenant le HSM backup n'a pas besoin d'être secouru — mais son
  **accès** (contrôle d'accès, éclairage) doit l'être pour une intervention de nuit.

## 142. Refroidissement : HSM et FPGA en baie

- **HSM appliance** (84-110 W) : charge thermique modeste, mais en **1U dense** — ne pas
  le coincer entre deux serveurs à 800 W sans airflow. Plage constructeur : **0-35 °C**
  (Luna, vérifié) — en salle à 27 °C ça passe, en local technique à 38 °C l'été, non.
- **FPGA** (190 W, refroidissement **passif**) : critique — la carte compte sur le flux
  d'air du serveur (avant → arrière, ~20-30 CFM). **Ne jamais** laisser un slot adjacent
  obstrué, ne jamais faire tourner le serveur capot ouvert en production (l'airflow est
  calculé capot fermé).
- **Baie** : panneaux obturateurs sur les U vides (évite le recyclage d'air chaud),
  brosses passe-câbles, sondes haut/bas de baie avec alertes à 30 °C / 32 °C.
- **Lien sécu** : une surchauffe qui coupe les serveurs = même effet qu'une coupure
  électrique sur les services de clés. La clim est un **organe de sécurité**.

## 143. Dimensionnement électrique : méthode

Pour la zone « crypto » d'une baie (exemple) :

1. **Inventaire** : 2× HSM réseau (2× 110 W max) + 1 serveur Vault (350 W) + 1 serveur
   PKI (300 W) + switch management (50 W) = **~920 W**.
2. **Marge** : ×1,5 (pics, vieillissement, évolutions) → **~1 400 W**.
3. **Onduleur** : la zone crypto doit tenir **30 min** sur batteries → ~700 Wh utiles
   (avec rendement 0,9 : ~780 Wh de batteries).
4. **Disjoncteurs** : 1 400 W / 230 V ≈ 6,1 A → disjoncteur **10 A** par PDU (courbe C),
   câbles 1,5 mm² mini (2,5 mm² recommandé).
5. **Clim** : ~1 400 W thermiques à évacuer pour cette zone (+ les serveurs voisins).

Documenter ce calcul dans le **dossier électrique de la baie** — c'est aussi une pièce
d'audit (disponibilité = contrôle ISO 27001 A.8.x).

## 144. Monitoring énergie des équipements de sécurité

Superviser, comme la température :

- **Puissance instantanée** par PDU (les PDU metered le font nativement) — dérive
  anormale = alerte (ventilateur HS, composant en surcharge).
- **État onduleur** : charge %, autonomie restante, état batteries, dernier test.
- **Température** : HSM (via SNMP), FPGA (via driver), baie (sondes).
- **Batterie interne HSM** : état remonté en SNMP — alerte **préventive** 6 mois avant
  fin de vie estimée.
- **Corrélation** : pic de conso + alerte température + échec d'opération crypto =
  incident unique à traiter comme tel (pas 3 tickets séparés).

## 145. Groupe électrogène et HSM : séquence de redémarrage

Après une coupure longue (groupe électrogène en service) :

1. **Stabiliser** : attendre la tension/fréquence stables du groupe (2-5 min), vérifier
   l'onduleur en mode normal (plus sur batteries).
2. **Réseau** : switches, firewall du VLAN crypto.
3. **HSM** : allumage, vérification des scellés/journaux (une coupure brutale peut
   générer des alertes tamper — les **lire**, pas les ignorer), vérification des
   partitions.
4. **KMS** : Vault en auto-unseal (vérifier le descellement), sinon procédure Shamir
   avec les custodians.
5. **Applications** : bases TDE, PKI, TSA — dans l'ordre des dépendances.
6. **Vérifications** : une signature test, un déchiffrement test, les CRL/OCSP à jour.
7. **PV** : durées, écarts, anomalies — archivé.

