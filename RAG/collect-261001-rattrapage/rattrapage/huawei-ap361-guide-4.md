---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-4
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [450, 605]
sha256: d855d6f9b5b42b544dc5319cd6413ee220aa36a561474652f5e81e72c9d020f2
---

# Huawei eKit AP361 — Guide ultra-complet

- **Onduleur** : le switch PoE **doit** être sur onduleur. 12 AP361 = ~120 W :
  c'est rien pour un onduleur de baie, mais si le switch n'est pas ondulé, tout
  le Wi-Fi tombe à la première microcoupure (et les AP mettent 2-3 min à revenir).
- **Dimensionnement onduleur** : ajoutez la charge PoE réelle (pas le budget max)
  + switch lui-même (~30-50 W) + marge. Voir votre guide onduleurs.
- **Parafoudre** : type 2 en tête de l'alimentation de la baie ; en zone
  tropicale/orages fréquents, envisagez des parafoudres Ethernet sur les liens
  les plus exposés.
- **Disjoncteur dédié** : un départ dédié pour la baie informatique, étiqueté,
  avec consignation possible. Évitez le départ « prises du couloir ».

## 24. Mise en garde : le PoE qui « ne monte pas »

Symptômes : LED éteinte, ou LED allumée puis l'AP reboote en boucle.

| Cause | Diagnostic | Solution |
|---|---|---|
| Budget PoE du switch dépassé | Voyant PoE du switch en défaut, log « power denied » | Délester ou changer de switch |
| Câble trop long / CCA | Testeur : longueur > 90 m, ou résistance anormale | Refaire le lien en cuivre |
| Paire abîmée | Négociation à 100 Mbit/s ou rien | Refaire les connecteurs, retester |
| Injecteur sous-dimensionné | Injecteur 24 V passif ou < 15 W | Injecteur 802.3af 15,4 W |
| Port du switch en défaut | Tester l'AP sur un autre port | Reset du port / RMA switch |
| AP défectueux | Même symptôme sur 2 alimentations différentes | RMA |

> 🔁 **Astuce** : un reboot en boucle à froid le matin puis stable l'après-midi =
> souvent un câble limite (résistance qui varie avec la température) ou un budget
> PoE tout juste. Ne laissez pas traîner.

## 25. Checklist PoE de mise en service

- [ ] Budget PoE calculé et documenté (tableau §19 rempli pour le site)
- [ ] Switch : budget ≥ besoin + 20 %, firmware à jour
- [ ] Chaque lien testé (wiremap + longueur < 90 m câble fixe)
- [ ] Négociation **1000 Mbit/s full** vérifiée sur chaque port (dans l'app eKit ou le switch)
- [ ] Switch sur onduleur, parafoudre en tête
- [ ] Étiquetage : chaque câble repéré aux deux extrémités (AP-01 ↔ port switch 5)
- [ ] Test de coupure : débrancher/rebrancher un port PoE → l'AP doit revenir seul en < 5 min

---
---

# D. PREMIÈRE MISE EN ROUTE

## 26. Reset d'usine : quand et comment

**Quand** : AP d'occasion, AP déjà configuré pour un autre site, onboarding cloud
qui échoue mystérieusement, reprise en main après départ d'un prestataire.

**Comment** (procédure standard Huawei, 🔎 libellé exact du bouton à vérifier) :

1. AP sous tension (PoE branché, LED allumée).
2. Trombone dans le trou **Reset**, appuyer **> 5 secondes** (jusqu'à changement
   d'état de la LED).
3. Relâcher. L'AP redémarre (2 à 4 minutes).
4. L'AP est en configuration d'usine : prêt pour l'onboarding.

**Ce que le reset efface** : SSID, comptes admin, adresse du cloud, certificats.
**Ce qu'il n'efface pas** (en général) : le firmware. Si l'AP a été « briquée »
par un upgrade interrompu, le reset ne suffit pas → voir §118.

## 27. Onboarding par l'app eKit (chemin officiel recommandé)

C'est **le** chemin prévu par Huawei pour l'AP361. Procédure type :

1. Installez **HUAWEI eKit** (App Store / Play Store), créez un compte Huawei.
2. Créez votre **site** (nom du client / bâtiment) dans l'app.
3. **Scannez le QR code / code-barres** de l'AP (sur le carton ou sous l'appareil)
   → l'AP est ajouté au site.
4. Branchez l'AP en PoE sur le réseau (il doit obtenir une IP par DHCP et accéder
   à Internet pour joindre le cloud eKit).
5. Dans l'app : l'AP apparaît « en ligne » → assistant de configuration :
   SSID, mot de passe, bande, etc.
6. Répétez pour chaque AP, puis vérifiez la topologie dans l'app.

**Prérequis réseau pour l'onboarding cloud** (à préparer **avant** le jour J) :

- [ ] DHCP actif sur le VLAN d'accueil des AP (ou IP statique prévue)
- [ ] Accès Internet sortant (l'AP doit joindre le cloud eKit : **ne pas**
      bloquer les flux HTTPS sortants vers les domaines Huawei)
- [ ] DNS fonctionnel
- [ ] Heure correcte (NTP) — un décalage > quelques minutes peut faire échouer
      l'établissement TLS vers le cloud

## 28. Interface web locale (mode autonome/Fat)

Si l'AP est en mode autonome (Fat) ou pour un accès local de secours :

1. Branchez un PC sur le même réseau que l'AP (ou directement sur son port via
   un switch/injecteur).
2. Trouvez son IP : via le DHCP du site (cherchez l'entrée correspondant à la
   MAC de l'étiquette) ou via un scan réseau.
3. `https://<IP-AP>` — identifiants par défaut : 🔎 **à vérifier sur le guide de
   démarrage du lot** (les valeurs par défaut changent selon les versions ; les
   firmwares récents **forcent** la création d'un mot de passe à la première
   connexion — ne cherchez pas un « admin/admin » universel).
4. Première action : **changer le mot de passe** (voir §144).

> ⚠️ Si la page web ne répond pas : vérifiez que vous êtes bien sur le même
> sous-réseau, que le firewall du PC ne bloque pas, et essayez en HTTP si le
> firmware est ancien (puis basculez en HTTPS — voir §146).

## 29. Accès SSH/CLI : état des lieux honnête

- Sur les AP Huawei récents gérés par eKit, l'accès SSH/Telnet **peut être
  désactivé par défaut** ou réservé au mode Fat. 🔎 Vérifiez dans la documentation
  de votre version logicielle.
- Si disponible : activez SSH (jamais Telnet en production), restreignez-le au
  VLAN de management (voir §145), et utilisez des clés ou des mots de passe
  robustes (voir §144).
- Les exemples CLI des sections suivantes sont donnés pour le mode autonome /
  pour comprendre la logique VRP ; en mode cloud eKit, **l'app fait foi** et
  certaines commandes peuvent ne pas exister.

## 30. Mise à niveau logicielle initiale (upgrade « jour 1 »)

**Règle** : ne laissez jamais un AP en production sur le firmware d'usine.

1. Dans l'app eKit / le cloud : vérifiez la version actuelle de chaque AP.
2. Consultez les **release notes** de la version cible (correctifs de sécurité,
   compatibilité WPA3, stabilité).
3. Planifiez l'upgrade **hors heures ouvrées** (un upgrade = 5 à 15 min/AP avec
   reboot).
4. Upgrader **un AP pilote** d'abord, validez 24 h (SSID, roaming, débit), puis
   généralisez par vagues (jamais 100 % du parc d'un coup).
5. Vérifiez que tous les AP sont sur la **même version** à la fin (un parc
   hétérogène = diagnostics infernaux).

> ⛔ **Interdit** : couper le PoE pendant l'upgrade. C'est le meilleur moyen de
> corrompre la flash (voir §118).

## 31. Vérifications post-boot (les 5 minutes qui évitent les retours)

Pour chaque AP, après mise en service :

- [ ] LED verte fixe (état normal)
- [ ] Visible « en ligne » dans l'app eKit / le cloud
- [ ] IP correcte (bon VLAN, bon sous-réseau)
- [ ] Port négocié à **1000 Mbit/s**
- [ ] SSID diffusés visibles depuis un téléphone (les 2 bandes si activées)
- [ ] Association test : un client se connecte, obtient une IP, ping la passerelle
- [ ] Débit test : speedtest local ou iperf vers un serveur du LAN (notez la valeur)
- [ ] Version firmware = version cible du site

## 32. Erreurs classiques de la mise en route

| Erreur | Conséquence | Prévention |
|---|---|---|
| Onboarding sans DHCP | L'AP ne joint pas le cloud, reste « hors ligne » | Préparer le DHCP **avant** le chantier |
| Firewall qui bloque le cloud | AP « en ligne » puis « hors ligne » en boucle | Ouvrir HTTPS sortant vers Huawei eKit dès la maquette |
| Tous les AP upgradés d'un coup | Panne généralisée si la version a un bug | Vague pilote + généralisation |
| Pas d'étiquetage SN/MAC | Impossible de retrouver quel AP est où | Photo de l'étiquette + plan numéroté |
| Mot de passe par défaut conservé | Porte ouverte (voir §144) | Changement systématique jour 1 |

---
---

