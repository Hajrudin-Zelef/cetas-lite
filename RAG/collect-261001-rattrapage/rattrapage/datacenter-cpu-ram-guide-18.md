---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-18
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr", "attention", "datacenter"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [2814, 2988]
sha256: eb9b7e50aef8d020e6398e3a5c331d2c6abfe2ad40e89fb9f993d637126eae81
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

Points critiques :
- **Courbe de disjoncteur** : D (ou K) pour les charges informatiques à
  appel de courant, pas B (déclenchements intempestifs au démarrage).
- **Serrage** : 80 % des échauffements viennent de cosses mal serrées —
  contrôle thermographique annuel (caméra IR).
- **Sélectivité** : le disjoncteur de la PDU doit déclencher avant celui
  de l'arrivée générale (sinon une baie fait tomber la salle).

## 154. Autonomie batterie : calcul détaillé (indicatif)

```
 Energie_utile = P_IT x duree_cible / rendement_onduleur

 Exemple : 13,75 kW x 0,25 h (15 min) / 0,94 = 3,66 kWh utiles
```

Facteurs qui réduisent l'autonomie réelle :
- **Vieillissement** : -20 % de capacité à 3-4 ans (batteries plomb).
- **Température** : chaque +10 °C au-dessus de 20 °C divise la durée de vie
  par 2 (loi d'Arrhenius).
- **Décharge profonde** : les autonomies catalogue sont à 100 % de
  décharge, ce qui tue les batteries — comptez 80 % utilisable en pratique.

Règle : **dimensionnez pour 1,5× la durée cible** (15 min voulues → 22 min
catalogue). Et programmez l'arrêt ordonné (NUT) à 30-40 % de batterie
restante, pas à 5 %.

## 155. Redondance électrique : N+1 vs 2N

| Architecture | Principe | Disponibilité | Coût (indicatif) |
|---|---|---|---|
| N | 1 onduleur, pas de secours | basse | 1× |
| N+1 | 2 onduleurs, 1 suffit (+bypass) | bonne | ~1,6× |
| 2N | 2 chaînes indépendantes A/B | très haute | ~2× |

En pratique serveur : les alimentations redondantes des serveurs (A/B)
exigent **2 sources** : soit 2 onduleurs (2N), soit 1 onduleur + 1 arrivée
directe (compromis courant en PME). **Ne branchez jamais les deux
alimentations d'un serveur sur la même PDU** : c'est la redondance la plus
chèrement payée pour zéro bénéfice.

## 156. Free cooling : chiffrage de l'économie (indicatif)

Principe : utiliser l'air extérieur (ou l'eau de nappe) au lieu du groupe
froid quand la température le permet.

| Climat (France) | Heures de free cooling/an (indicatif) | Économie froid |
|---|---|---|
| Nord / montagne | 5 000-6 500 h | 50-70 % |
| Région parisienne | 4 000-5 000 h | 40-55 % |
| Sud méditerranéen | 2 500-3 500 h | 25-40 % |

Calcul : salle de 11 kW IT, froid = 5,5 kW moyens. Sans free cooling :
5,5 × 8760 × 0,18 = 8 670 €/an. Avec 50 % de free cooling : **~4 300 €/an
d'économie**. Sur 5 ans : 21 500 € — de quoi financer une bonne partie du
système de free cooling lui-même.

## 157. Valorisation de la chaleur fatale

Un datacenter est un **radiateur géant** : 11 kW IT = 11 kW de chaleur
récupérable (plus le froid). Options :
- **Eau tiède (30-45 °C, DLC)** : chauffage de bureaux, préchauffage ECS —
  directement utilisable avec des ventilo-convecteurs.
- **Air chaud (35-45 °C)** : préchauffage d'ateliers, serres (projets
  existants en France).
- **Pompe à chaleur** : remonter à 60-70 °C pour du chauffage urbain
  (rendement 3-4 : 1 kW électrique → 3-4 kW de chaleur).

Ordre de grandeur : 11 kW valorisés à 0,10 €/kWh thermique (prix du gaz
évité) × 8760 h × 80 % de disponibilité = **~7 700 €/an**. La chaleur
n'est un déchet que si on décide de la jeter.

## 158. Suivi de consommation : du wattmètre au tableau de bord

Chaîne de mesure recommandée :
1. **PDU monitorées** : mesure par prise ou par phase, remontée SNMP —
   la source de vérité par serveur.
2. **BMC/IPMI** : puissance instantanée du serveur (précision ±5-10 %,
   suffisante pour le suivi).
3. **Onduleur** : puissance totale + rendement — la source de vérité
   globale.
4. **Supervision** (Zabbix, Prometheus) : historisation, alertes sur seuils
   (ex. serveur > 120 % de sa conso nominale = anomalie).

Tableaux de bord minimaux : kW instantané par baie, kWh/jour/semaine,
PUE glissant, top 5 des serveurs consommateurs. **On n'optimise que ce
qu'on mesure** — et la première optimisation est toujours un serveur
oublié allumé.

## 159. Green IT : au-delà du PUE

Le PUE ne fait pas une politique environnementale. Complétez avec :
- **CUE** (carbone) : gCO2/kWh selon le mix électrique du fournisseur —
  choisir un fournisseur bas-carbone bat tous les PUE.
- **WUE** (eau) : les tours de refroidissement consomment de l'eau ;
  le DLC à eau tiède en boucle fermée n'en consomme presque pas.
- **Durée de vie** : garder un serveur 5-7 ans au lieu de 3-4 divise
  l'empreinte fabrication (qui représente 20-30 % du total sur la vie).
- **Réemploi** : reconditionné (section 138), don/revente en fin de vie.

Pour un chef de service : intégrez ces 4 indicateurs au reporting annuel.
C'est ce qui transforme « on fait attention » en pilotage.

## 160. Plan électrique type d'une petite salle (récapitulatif)

```
 Arrivee EDF -> TGBT -> [Onduleur 20 kVA] -> Tbl. ondulé -> PDU A/B -> serveurs
                          | (bypass)          |
                          +-> froid, éclairage (non ondulé, délestageable)

 Protections : courbe D, sélectivité amont/aval, parafoudre type 1+2
 Mesure : PDU monitorées + supervision (section 158)
 Tests : mensuel (onduleur), annuel (groupe, thermographie)
```

Règle d'or : **tout ce qui est ondulé doit être délestable par priorité**
(serveurs critiques > serveurs secondaires > bureautique). En cas de
coupure longue, on éteint proprement par vagues (NUT) au lieu de tout
perdre d'un coup.

---

## 161. Supervision BMC : ce qu'il faut remonter

Via IPMI/Redfish, chaque serveur expose : températures (CPU, DIMM,
entrée/sortie), vitesses ventilateurs, tensions, puissance instantanée,
compteurs ECC, état alimentations, journaux SEL.

Seuils d'alerte recommandés (à adapter) :
- Température entrée d'air > 30 °C → alerte (dérive climatique).
- Ventilateur en panne → critique immédiate (redondance consommée).
- Erreurs ECC corrigées > 10/jour sur une barrette → ticket préventif.
- Alimentation en défaut → critique (redondance consommée).
- Puissance > 120 % de la nominale → anomalie workload ou thermique.

**Un serveur non supervisé est un serveur dont on apprend la panne par
l'utilisateur.** Le BMC est gratuit (déjà payé) : l'excuse du « on n'a pas
le temps » coûte plus cher que la mise en place.

## 162. Zabbix/Prometheus : modèles de supervision serveur

Points de collecte par serveur (templates à créer une fois) :
- SNMP vers PDU : puissance par prise (section 158).
- Redfish/IPMI vers BMC : sondes, fans, ECC, SEL.
- Agent OS : charge CPU, mémoire, I/O disques, réseau.
- Onduleur (NUT/snmp) : charge %, batterie %, autonomie restante.

Règles d'alerte (exemples) :
- `temp_inlet > 30°C` pendant 15 min → warning.
- `fan_status != ok` → critical.
- `ecc_corrected_delta > 10/h` → warning avec nom de la barrette.
- `ups_battery < 40%` et secteur absent → déclencher arrêt ordonné.

## 163. Maintenance préventive : planning annuel type

| Fréquence | Actions |
|---|---|
| Mensuelle | alertes BMC/PDU, onduleur (test), sauvegardes supervisées |
| Trimestrielle | dépoussiérage grilles, écoute ventilateurs, revue des ECC |
| Semestrielle | firmware (vague pilote puis général), test arrêt ordonné NUT |
| Annuelle | thermographie armoire électrique, test groupe, revue capacitaire |

Documentez chaque intervention (date, serveur, action, résultat) : c'est
l'historique qui permet de passer du curatif au prédictif — et de justifier
les budgets de renouvellement avec des faits.

## 164. Gestion des firmwares en parc : la méthode des vagues

1. **Veille** : alertes constructeur (failles, correctifs critiques).
2. **Vague 0** : 1 serveur pilote (non critique), 1-2 semaines d'observation.
3. **Vague 1** : 25 % du parc, en dehors des heures critiques.
4. **Vague 2** : le reste, avec rollback planifié (firmware précédent
   conservé).
5. **Jamais** : tout le parc en même temps, ni pendant un pic d'activité.

