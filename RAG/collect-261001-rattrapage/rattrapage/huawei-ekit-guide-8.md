---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-8
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Huawei"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [606, 723]
sha256: 75ccfe2c4e1fff15841e67d9c6eb2cd6faa870bcaf8f19c5b2d099693628fa38
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

1. **Portail captif** sur VLAN 30 : code chambre ou e-mail + CGU, session 24 h, débit 15 Mbit/s/client (exemple fictif), isolation inter-clients.
2. **QoS voix** : VLAN 60 prioritaire (les téléphones de la réception ne doivent jamais hacher).
3. **Isolement** : VLAN 30 et 50 sans accès inter-VLAN (un client ne voit ni les autres clients ni le LAN staff).
4. **Plages DHCP** : /24 par VLAN avec baux adaptés (chambres : 12 h ; restaurant : 4 h).
5. **Nommage AP** : `HOTEL-Etage1-AP01` … — quand l'alerte « AP hors ligne » arrive, tu sais où envoyer le technicien sans réfléchir.
6. **Tests** : chambre la plus éloignée (débit), piscine (AP761), bascule (débranche un switch → vérifie que l'autre moitié tient), portail captif sur iPhone + Android (les deux comportements diffèrent).

## 55. Scénario C — Commerce / boutique : cadrage

**Besoin type :** boutique 150 m² + réserve, 8 employés, 2 caisses (TPE), Wi-Fi clients gratuit, 4 caméras, 1 enseigne avec 3 boutiques (multi-sites).
**Architecture :** par boutique : 1 AR180 (suffit : ~100 terminaux) + 1 S220-8P4S (8 ports PoE+ 125 W : 2 AP + 4 caméras + marge) + 2 AP361 (surface de vente + réserve). Le tout supervisé en **multi-sites** dans le cloud.

## 56. Scénario C — Topologie ASCII (par boutique)

```
   [Internet Box operateur]
              |
   [AR180] (192.168.1.1, DHCP, Wi-Fi integre en secours)
              |
   [S220-8P4S] (125 W PoE+)
     |    |    |    |
  [AP361] [AP361] [Cameras x4] [Caisses x2 (VLAN 10)]
  Surface  Reserve  (VLAN 40)   TPE (VLAN 10, prioritaire)
  (VLAN 10/20)

VLAN:
  VLAN 10  CAISSES+STAFF  192.168.1.0/24
  VLAN 20  CLIENTS         192.168.2.0/24 (portail captif, isole, 5 Mbit/s)
  VLAN 40  CAMERAS         192.168.4.0/24

Cloud : 1 tenant prestataire, 3 sites (Boutique Centre, Boutique Nord, Boutique Sud),
        modele de config "boutique standard" applique aux 3.
```

**Spécificités commerce :**
- Les **TPE (terminaux de paiement)** : jamais sur le Wi-Fi invités, VLAN dédié ou a minima VLAN staff, et surtout **ne jamais couper leur connectivité** pendant une mise à jour en journée. Fenêtre de maintenance : avant ouverture.
- **Portail captif** = outil marketing (récupération d'e-mails opt-in, page promo). Cadre-le juridiquement (mentions, consentement).
- Multi-sites : le changement du PSK invités se fait **une fois pour les 3 boutiques** depuis le cloud.

## 57. Scénario C — Liste matériel et configs clés

| Qté / boutique | Équipement |
|---|---|
| 1 | AR180 |
| 1 | S220-8P4S (125 W PoE+) |
| 2 | AP361 |
| 4 | Caméras IP |
| 1 | Onduleur 1000 VA |

**Budget PoE :** 2× 9 + 4× 10 = 58 W < 125 W : OK.
**Configs clés :** modèle « boutique standard » dans le cloud (VLAN, SSID `BOUTIQUE-Clients` + `BOUTIQUE-Staff`, portail captif identique, plages DHCP identiques — attention : **mêmes plages IP sur des sites différents, c'est OK** car chaque site est un LAN indépendant ; ça simplifie même la maintenance). Alertes « équipement hors ligne » par boutique avec le nom du responsable local.

## 58. Scénario D — Petit entrepôt : cadrage

**Besoin type :** entrepôt 800 m², 10 employés, terminaux mobiles (douchettes Wi-Fi), 6 caméras, portail d'accès, bureau mezzanine (5 PC), chariots élévateurs (pas de Wi-Fi temps réel critique, mais connectivité stable).
**Défis spécifiques :** hauteur sous plafond (6-8 m), rayonnages métalliques (réflexions, zones d'ombre), poussière, température.
**Architecture :** 1 AR280 + 1 S310-24P4S + 4-6 AP (AP361 en intérieur, ou AP761 si zones semi-ouvertes / quais). AP fixés **sous plafond mais pas trop haut** : en entrepôt, on descend les AP à 4-5 m sur des mats ou chemins de câbles plutôt que de les coller à 8 m (à 8 m, le signal s'étale et les douchettes en bas captent mal).

## 59. Scénario D — Topologie ASCII

```
   [Internet Fibre]
          |
   [AR280] (bureau mezzanine, baie)
          |
   [S310-24P4S] (baie mezzanine, 380 W PoE+)
     |        |         |          |
  [AP x2]  [AP x2]   [Cameras]  [PC bureau]
  Allee A  Allee B   (VLAN 40)  (VLAN 10)
  (mats 4-5 m, VLAN 10/20)

VLAN:
  VLAN 10  BUREAU+DOUCHETTES  192.168.10.0/24 (les douchettes = critiques, SSID dedie)
  VLAN 20  INVITES            192.168.20.0/24 (transporteurs au bureau, isole)
  VLAN 40  CAMERAS            192.168.40.0/24

Regle radio entrepot : 1 AP / ~400-500 m² en rayonnages hauts (ordre de grandeur),
canaux 5 GHz uniquement pour les douchettes si elles le supportent (moins d'interferences),
puissance moderee, survey OBLIGATOIRE avec un chariot (test en mouvement).
```

## 60. Scénario D — Liste matériel et points de vigilance

| Qté | Équipement |
|---|---|
| 1 | AR280 |
| 1 | S310-24P4S |
| 4-6 | AP361 (intérieur) — valider le nombre par survey |
| 6 | Caméras IP (quais + allées) |
| 1 | Onduleur 1500-2000 VA |

**Vigilances terrain :**
- **Douchettes** : SSID dédié, 5 GHz, pas de portail captif (une douchette ne sait pas cliquer sur « J'accepte »), PSK robuste, roaming testé **en marchant** avec le chariot.
- **Poussière** : baie fermée avec filtre, AP IP adaptés si zone très poussiéreuse (sinon AP361 standard + nettoyage annuel).
- **Câbles** : chemin de câbles métallique, pas de câble volant au-dessus des allées (chariots + câbles = catastrophe).
- **Coupure électrique** : l'entrepôt a souvent des coupures — onduleur dimensionné (cf. guide onduleurs) + redémarrage automatique des équipements (ils reviennent seuls, vérifier quand même via le cloud).

---

## 61. Interopérabilité : eKit avec du non-eKit, est-ce que ça marche ?

Oui, **aux standards ouverts**, non aux protocoles propriétaires. Tableau de vérité :

| Fonction | eKit ↔ non-Huawei | Détails |
|---|---|---|
| Commutation L2 / VLAN 802.1Q | Oui | Un trunk 802.1Q entre un S310 et un switch d'une autre marque fonctionne (tagging standard). Vérifie le VLAN natif des deux côtés — cause n°1 d'incompatibilité. |
| Routage statique / DHCP / DNS | Oui | Standards IP universels. |
| Wi-Fi (clients) | Oui | 802.11ax/be : n'importe quel client Wi-Fi 6/7 se connecte. |
| Roaming entre AP eKit et AP tiers | **Non / dégradé** | Le roaming rapide (802.11r/k/v) inter-constructeurs est aléatoire. En pratique : les clients se reconnectent, avec une micro-coupure. Ne mélange pas les marques sur une même zone de couverture. |
| Gestion unifiée (app eKit) | **Non** | L'app ne voit que du eKit. Le reste = son outil d'origine. |
| PoE 802.3af/at | Oui | Standard : un AP eKit sur un switch PoE tiers fonctionne, et inversement. |
| Agrégation de liens (LACP) | Oui si 802.3ad des deux côtés | À vérifier par modèle côté eKit (documentation officielle). |
| Stacking | **Non** | Le stacking est propriétaire : on ne stacke pas un S310 avec un switch d'une autre marque. |

**Règle d'or :** homogène par **fonction** (tous les AP d'un site = même gamme ; tous les switchs d'accès = même gamme), hétérogène autorisé aux **frontières** (un vieux switch en cascade derrière un S310, ça passe).

## 62. Migration depuis un parc existant : la méthode en 5 phases

