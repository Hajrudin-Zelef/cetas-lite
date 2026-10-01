---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-9
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [724, 824]
sha256: ee788f4ce7a368f2becd99016165c0d9216f4674ae9915306e4c824d6f99b9f5
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

**Phase 1 — Inventaire (1/2 journée) :** liste chaque équipement existant (marque, modèle, âge, rôle, PoE ou non), le plan d'adressage actuel, les SSID et PSK en service, les points de douleur (« le Wi-Fi coupe dans la salle de réunion »).
**Phase 2 — Coexistence (jour J) :** installe le cœur eKit (AR + switch) **en parallèle** de l'existant. Nouvelle plage IP eKit, SSID eKit en plus des anciens (ex. fictif : `BUREAU-NEW` à côté de `BUREAU`).
**Phase 3 — Bascule pilote :** migre 2-3 utilisateurs volontaires + 1 imprimante. 48 h d'observation.
**Phase 4 — Bascule générale :** en heures creuses, bascule les clients (change le PSK de l'ancien SSID pour forcer la migration, ou éteins les vieux AP un par un). Garde l'ancien matériel sous la main 1 semaine.
**Phase 5 — Décommissionnement :** une fois le client satisfait (2 semaines), retire l'ancien matériel, mets à jour le plan, archive l'ancien adressage.
**Jamais** de big-bang un vendredi à 17 h. La migration en douceur, c'est 90 % de la réussite.

## 63. Réutiliser l'existant : quoi garder, quoi jeter

- **Câblage Cat5e/Cat6 en bon état** : à garder (teste-le au testeur, surtout les prises qui « marchaient avant »).
- **Vieux switch 100 Mbit/s** : à jeter (goulot d'étranglement garanti avec du Wi-Fi 6).
- **Vieux switch Gbit non-PoE en bon état** : réutilisable en cascade pour des PC filaires non critiques (pas pour alimenter des AP).
- **Vieux AP Wi-Fi 4/5** : à éteindre une fois les eKit en place (un vieux AP qui diffuse encore crée des interférences et attire les clients lents qui plombent la cellule radio).
- **Onduleur existant** : à garder si les batteries ont moins de 3 ans et la puissance suffit (sinon : remplacement — cf. guide onduleurs).

## 64. Adressage IP : le plan type PME eKit

Utilise du **privé RFC 1918** avec une logique par site. Modèle recommandé (exemple fictif, à adapter) :

| Site | Réseau de base | VLAN staff | VLAN invités | VLAN caméras | VLAN voix |
|---|---|---|---|---|---|
| Site 1 (siège) | 192.168.10.0/23 | .10.0/24 (VLAN 10) | .11.0/24 (VLAN 20) | .12.0/24 (VLAN 40) | .13.0/24 (VLAN 60) |
| Site 2 (boutique) | 192.168.20.0/24 | .20.0/24 (VLAN 10) | .21.0/24 (VLAN 20) | — | — |

**Règles :**
- Le **3e octet = le site**, le VLAN affine. En dépannage à 2 h du matin, tu sais où tu es rien qu'en voyant l'IP.
- **Équipements d'infrastructure en IP fixes basses** : passerelle en .1, switchs en .2, .3…, AP en .10-.49 (ou DHCP avec réservation — les deux se défendent ; le fixe documenté est plus robuste si le DHCP tombe).
- **Jamais** de 192.168.1.0/24 ou 192.168.0.0/24 sur un site pro : ce sont les plages par défaut des box grand public → conflits garantis dès qu'on branche une box ou un VPN.
- Documente tout dans une **fiche site** (section 66).

## 65. Fiche site type (à remplir pour chaque déploiement)

```
FICHE SITE - [Nom client - Ville - Batiment]          Date : __________
Contact client : __________ Tel : __________
Compte eKit / tenant : __________ Role : __________
Acces Internet : operateur __________, type __________, debit __________
Passerelle : modele __________, SN __________, IP LAN __________, FW __________
Switchs : modele __________, SN __________, IP __________, FW __________
AP : modele __________, SN __________, emplacement __________, IP __________
Adressage : (tableau section 64)
VLAN : (tableau)
SSID : nom __________, PSK (dans le coffre, pas ici), VLAN __________
Onduleur : modele __________, puissance __________, autonomie estimee __________
Plan de cablage : [photo/schema joint]
Remarques : __________________________________________
```

## 66. VLAN : le découpage standard PME

| VLAN ID | Nom | Usage | Inter-VLAN |
|---|---|---|---|
| 1 | Défaut | À éviter pour le trafic (risque de fuite) | — |
| 10 | STAFF | Utilisateurs, PC, imprimantes | Vers Internet : oui ; vers autres VLAN : selon besoin |
| 20 | GUESTS | Invités / clients | Internet uniquement, **isolation totale** du LAN |
| 30 | CHAMBRES (hôtel) | Clients hôtel | Comme GUESTS |
| 40 | CAMERAS | Vidéosurveillance | Vers NVR uniquement, pas d'Internet (sauf accès distant via VPN) |
| 50 | CAISSES/TPV | Paiement, caisses | Internet + serveur de caisse uniquement |
| 60 | VOIX | Téléphonie IP | Prioritaire (QoS), vers IPBX |
| 70 | TV-IP | Télévision sur IP | Multicast local |
| 99 | MGMT | Gestion des équipements | Accès restreint aux admins |

**Erreurs à éviter :** mettre les caméras sur le VLAN staff « parce que c'est plus simple » (une caméra compromise = porte d'entrée sur le LAN) ; oublier le VLAN natif cohérent sur les trunks ; utiliser le VLAN 1 pour de la prod.

## 67. SSID : la doctrine (moins c'est mieux)

- **3 SSID maximum** par site : staff, invités, + 1 technique (douchettes, IoT) si indispensable. Chaque SSID supplémentaire = overhead radio (beacons) qui dégrade tout le monde.
- **Nommage** : explicite, sans caractères spéciaux ni espaces si possible (ex. fictifs : `HOTELPALM-Staff`, `HOTELPALM-Invites`). Évite les noms génériques (`WiFi`, `Linksys`) qui créent des confusions.
- **Bandes** : diffuse en 2,4 + 5 GHz par défaut ; si les clients sont récents, envisage un SSID **5 GHz uniquement** pour le staff (moins d'interférences) — à tester.
- **Sécurité** : WPA2-AES minimum, **WPA3** si tous les clients le supportent (les AP eKit récents le supportent d'après les fiches — **à vérifier** par modèle). PSK long (20+ caractères) généré, pas un mot du dictionnaire.
- **SSID invités** : portail captif ou PSK simple tourné régulièrement, **isolation client** activée, débit limité.
- **Ne diffuse pas le SSID staff** si tu veux réduire le bruit (sécurité par l'obscurité = faible, mais ça évite les tentatives des voisins). Le SSID invités, lui, doit être visible.

## 68. PoE : dimensionner sans se tromper (la méthode)

**Étape 1 — Lister les consommateurs :**

| Équipement | Standard PoE | Conso typique (terrain) |
|---|---|---|
| AP361 | 802.3af | ~9 W |
| AP761 | 802.3at (recommandé) | ~17,7 W max |
| AP Wi-Fi 7 (AP572/673) | 802.3at/bt selon modèle | 20-30 W (à vérifier par fiche) |
| Caméra IP fixe | 802.3af | 5-10 W |
| Caméra PTZ / IR puissant | 802.3at | 20-30 W |
| Téléphone IP | 802.3af | 3-7 W |

**Étape 2 — Additionner et comparer au budget du switch** (ex. : S310-48P4S = 380 W ; S220-8P4S = 125 W).
**Étape 3 — Marge de 20-30 %** : un budget PoE utilisé à 100 % ne laisse aucune place à l'ajout d'une caméra dans 6 mois.
**Étape 4 — Vérifier le budget PAR PORT** : 30 W max en 802.3at. Une caméra PTZ à 25 W sur un port af (15,4 W) = caméra qui redémarre en boucle.
**Étape 5 — Câble** : PoE = courant dans le câble → **Cat5e minimum, Cat6 recommandé**, 90 m max par brin, pas de « rallonge bricolée ».

## 69. PoE : les fonctions qui sauvent des déplacements

D'après les fiches (S220 : Fast PoE, Perpetual PoE) :
- **Perpetual PoE** : le switch continue d'alimenter les ports PoE pendant son redémarrage (mise à jour sans couper les caméras/AP). Vérifie que c'est activé.
- **Fast PoE** : réalimentation rapide des ports après une coupure électrique (les AP reviennent vite).
- **Cycle PoE à distance** : via l'app/cloud, coupe/rallume le PoE d'un port pour redémarrer un AP ou une caméra plantée — **le geste n°1 du dépannage à distance** (section 36).
- **Planification PoE** (si supportée — à vérifier) : couper le PoE des AP la nuit dans un bureau vide = économies d'énergie.

## 70. Ordre de mise en service : la séquence qui marche à tous les coups

