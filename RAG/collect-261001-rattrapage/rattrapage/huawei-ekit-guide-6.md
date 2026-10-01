---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-6
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [391, 476]
sha256: 7fc91d1b5be676a75e55d271ac292b1397e32d909c3da953e8362e33a66a56da
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

- **Remplacer un câble défectueux** ou rebrancher un AP : le cloud ne fait pas de magie physique.
- **Configurer finement** ce que l'interface cloud n'expose pas (ACL avancées, QoS fine, routage dynamique) — bascule en local (CLI/web) si le modèle le permet.
- **Récupérer un équipement totalement planté** (boot en échec) : il faut le reset physique sur site.
- **Contourner une coupure Internet du site** : sans uplink, le cloud ne voit plus rien (le réseau local continue — section 38).
- **Gérer du non-Huawei** : ton vieux switch manageable d'une autre marque reste hors du cloud eKit.
- **Faire de l'analyse forensique poussée** : pas de capture de paquets distante ni d'export NetFlow à ma connaissance (**à vérifier sur la documentation officielle**).

## 38. Panne Internet sur site : que devient le réseau eKit ?

Question que tout client pose : « si Internet tombe, mon réseau local tombe ? » Réponse à donner, honnête :
- **Le data plane est local** : le switch continue de commuter, les AP continuent de servir les SSID, le DHCP local (sur l'AR) continue de distribuer des adresses. Le bureau continue de travailler en local (fichiers, imprimantes).
- **Ce qui tombe** : l'accès Internet (évidemment), la **gestion cloud** (tu ne vois plus le site), les remontées d'alertes, et les fonctions qui dépendent du cloud (portail captif hébergé cloud s'il y en a un — **à vérifier**).
- **Au retour d'Internet** : les équipements se reconnectent seuls au cloud (prévois un délai de quelques minutes, ne redémarre pas tout en panique).
- Conséquence pratique : pour un site isolé (entrepôt au fin fond d'une zone), prévois un **accès de secours** : 4G sur la passerelle si le modèle le permet, ou au minimum un accès local documenté (IP de gestion, identifiants dans le coffre — section 102).

## 39. Planification radio avec le SNC : ce qu'il faut en attendre (ou pas)

Le SNC aide à la **planification et au déploiement** d'après Huawei, mais ne t'attends pas à un outil de survey prédictif complet :
- Utilise le SNC pour **placer tes AP sur plan** et suivre le déploiement (quoi, où, quel état).
- Le **vrai dimensionnement radio** reste manuel : 1 AP pour ~100-150 m² en bureaux cloisonnés placo, 1 AP par 2-3 chambres d'hôtel en cloison légère, 1 AP pour 20-30 utilisateurs actifs en usage bureautique. Ce sont des ordres de grandeur terrain — valide par un **survey post-installation** (tour avec un smartphone et une app de mesure Wi-Fi).
- Laisse l'**allocation automatique des canaux** active au début (les AP eKit gèrent le RRM de base), puis fige les canaux si tu constates des instabilités (**à vérifier** : niveau de contrôle RRM exposé dans l'app/SNC selon version).

## 40. Le portail captif invités : la fonction qui fait vendre

Pour hôtels, restaurants, boutiques : le Wi-Fi invités avec portail captif (page d'accueil avec CGU / code d'accès).
- Les AP eKit supportent l'**authentification PSK et portail** sans WAC ni serveur externe (fiche AP266 : fonctions d'authentification intégrées au cloud).
- L'USG6000F-S propose aussi un **portail local intégré** avec base locale ou RADIUS/AD/LDAP.
- Bonnes pratiques : SSID invités **isolé** (pas d'accès au LAN), débit **limité** par client (ex. fictif : 10 Mbit/s), durée de session limitée, page aux couleurs du client, mentions légales.
- **Piège classique** : le portail captif qui ne s'affiche pas sur certains smartphones → prévois toujours un **PSK invités de secours** communiqué à l'accueil.

## 41. Sauvegarde de la configuration via le cloud

- Le cloud conserve la configuration des sites (c'est le principe du pilotage centralisé). Mais **ne compte pas uniquement dessus** : exporte régulièrement une copie locale (section 89).
- Avant chaque changement majeur : capture d'écran ou export de la config actuelle + note datée (« avant migration VLAN invités »).
- En cas de remplacement d'un équipement défectueux (RMA) : le nouvel équipement onboardé sur le même site **récupère la configuration du site** via le cloud — c'est l'un des gros avantages opérationnels. Vérifie quand même après remplacement (SSID, VLAN, PoE).

## 42. Mises à jour via le cloud : méthode sans stress

1. **Lis la note de version** (release notes) avant tout : quoi de neuf, quoi de corrigé, incompatibilités connues.
2. **Sauvegarde** (section 89).
3. **Site pilote** : 1 site non critique, en heures creuses, avec quelqu'un sur place ou joignable.
4. **48 h d'observation** sur le pilote avant de généraliser.
5. **Par vagues** : jamais plus d'un tiers du parc en une fois.
6. **Fenêtre de maintenance annoncée** au client (même si « ça prend 5 minutes » — quand ça plante, ça prend 2 heures).
7. Note les versions dans le dossier de chaque site.

## 43. Inspection en ligne et « troubleshooting 2.0 »

La fonction d'**inspection en ligne** du SNC (mise en avant par Huawei) sert à :
- vérifier la santé d'un site avant une visite (tu arrives avec un diagnostic, pas avec des questions) ;
- détecter les anomalies récurrentes (AP qui décroche toutes les nuits à 2 h → suspecte le PoE ou l'alimentation, pas le Wi-Fi) ;
- valider après intervention (« tout est vert, je peux partir »).
Construis-toi une **routine d'inspection hebdomadaire** (15 min) : dashboard global → sites orange/rouge → détail → action ou ticket. C'est cette routine qui transforme le cloud en vraie valeur de maintenance.

## 44. Expansion simplifiée : ajouter un site ou étendre un site

- **Nouveau site** : crée le site dans le cloud, applique le modèle de configuration type (section 35), pré-stage les équipements au bureau (section 19), expédie ou va brancher. Un technicien junior peut faire le branchement si tout est pré-configuré.
- **Extension d'un site** : scan du SN du nouvel AP → branchage sur un port PoE → l'AP récupère la config du site automatiquement. C'est le cas d'usage « barcode scanning » de la fiche technique.
- **Règle** : toute extension = mise à jour du **plan** (adressage, ports utilisés, position des AP) le jour même. Le plan obsolète est la première cause de galère 6 mois plus tard.

## 45. Cloud eKit : tableau récapitulatif honnête

| Fonction | Via app/SNC | Remarques |
|---|---|---|
| Onboarding (Wi-Fi / scan SN) | Oui | Le cœur du système |
| Config SSID, VLAN de base, PoE | Oui | Détail selon version — à vérifier |
| Supervision multi-sites, alertes | Oui | Le gros point fort |
| Redémarrage / cycle PoE à distance | Oui | Évite 50 % des déplacements |
| Mises à jour firmware | Oui (HOUP côté switch) | Par vagues, jamais en prod directe |
| Inspection / dépannage simplifié | Oui (« 2.0 ») | Aide au diagnostic, pas de magie |
| Portail captif | Oui (AP ou USG) | Prévoir un PSK de secours |
| Config avancée (ACL fines, QoS, routage dyn.) | Non / limité | Gestion locale (CLI/web) si supportée |
| Équipements non-Huawei | Non | Outil séparé |
| Fonctionnement sans Internet | Gestion : non ; réseau local : oui | À expliquer au client |
| API publique / automatisation | À vérifier sur la documentation officielle | Ne pas promettre |
| Coût de la gestion cloud | Annoncé gratuit / sans licence | Vérifier licences USG avancées auprès du distributeur |

---

## 46. Scénario A — Bureau 20 postes : cadrage

**Besoin type :** PME de 20 personnes, open space + 2 bureaux fermés + 1 salle de réunion, 1 seul étage (~300 m²), fibre opérateur 1 Gbit/s, 25 PC filaires, 40 terminaux Wi-Fi (PC portables, smartphones), 1 imprimante réseau, 2 caméras IP (option), pas de téléphonie IP (ou 5 postes IP).
**Choix d'architecture :** 1 passerelle AR + 1 switch PoE + 3 AP. Simple, évolutif, tout supervisé par l'app.

## 47. Scénario A — Topologie ASCII

