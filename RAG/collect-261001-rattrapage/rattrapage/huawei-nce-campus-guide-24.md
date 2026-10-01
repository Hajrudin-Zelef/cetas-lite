---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-24
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Huawei"]
dates: []
keywords: ["attention", "memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [2114, 2223]
sha256: ff3bec0527f927cc83e53f8b65a8c301cdcbf60904479c3eaddf5e5c1352aee3
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

Fonctions WAN pilotées :
- **Provisionnement des liens** : Internet, MPLS, 4G/5G — avec bascule automatique (le lien de secours prend le relais sans intervention).
- **VPN site à site** : tunnels chiffrés vers le siège / entre branches, déployés par politique (pas de configuration crypto manuelle par routeur).
- **Sélection de chemin par application** : le trafic critique (voix, applicatif métier) passe par le meilleur lien disponible.
- **ZTP du routeur de branche** : l'AR720 d'une nouvelle agence se configure seul (section 59) — c'est souvent **le** cas d'usage qui vend NCE à la direction (ouvrir une agence en jours, pas en semaines).

Point d'honnêteté : le SD-WAN Huawei est mature sur les scénarios standard (dual-link, VPN) ; les architectures exotiques (multi-opérateurs avec SLA différenciés fins) demandent une validation en maquette. Et la qualité du SD-WAN ne compensera jamais un lien opérateur pourri — mesurer les liens (NQA, section 99) avant d'accuser le contrôleur.

## 162. Atelier de conception NAC — méthode pas à pas

Déployer le 802.1X/NAC est un projet dans le projet. Méthode recommandée :

1. **Inventaire des populations** : qui se connecte ? (employés, prestataires, invités, IoT, équipements techniques — ne pas oublier les copieurs !).
2. **Matrice d'autorisation** : pour chaque population → méthode d'authentification (802.1X/MAB/portail), VLAN, ACL, QoS. Tenir sur **une page** — si ça ne tient pas, simplifier.
3. **Choix des méthodes EAP** : PEAP-MSCHAPv2 pour démarrer (AD), trajectoire EAP-TLS (section 67).
4. **Infrastructure** : RADIUS (composant NCE ou externe), AD/LDAP, PKI si EAP-TLS, supplicants (GPO pour les postes Windows).
5. **Mode monitor d'abord** : déployer en **observation** (on authentifie mais on ne bloque pas) pendant 2-4 semaines pour découvrir tous les cas particuliers (la vieille imprimante du comptable, le badgeuse...).
6. **Bascule progressive** : par VLAN/bâtiment, avec fenêtre et rollback.
7. **Filets** : MAB pour le non-802.1X (liste blanche), portail pour les invités, procédure d'urgence (mode ouvert temporaire — cas pratique 10).

Le mode monitor n'est pas une option : c'est lui qui évite le « lundi matin noir ».

## 163. Dimensionner le RADIUS — ne pas l'oublier

Le RADIUS est le cœur battant du 802.1X : s'il ne répond pas (ou lentement), tout le monde est bloqué ou ralenti.

- **Capacité** : estimer les pics (le matin à 8 h, tout le monde s'authentifie en même temps — prévoir la charge de pointe, pas la moyenne).
- **Redondance** : au moins 2 serveurs RADIUS (ou le composant NCE + un secondaire), sur des sites/VM différents.
- **Timeouts** : régler les délais côté switch (trop court = rejets intempestifs, trop long = ouverture de port interminable pour l'utilisateur).
- **Supervision** : sonder le RADIUS en permanence (une sonde d'authentification test toutes les 5 min — si elle échoue, alerte critique **avant** les utilisateurs).
- **Composant local** : sur les branches à WAN fragile, le composant d'authentification NCE local (section 21) évite la dépendance totale au central.

## 164. Portail captif invités — conception détaillée

Au-delà des principes (section 79), la conception d'un portail invité propre :

- **Parcours utilisateur** : association au SSID → redirection → page (logo, CGU, choix : voucher / SMS / sponsor) → authentification → accès. Chaque étape doit marcher sur **tous** les OS (iOS, Android, Windows ont des comportements de détection de portail différents — tester les trois).
- **Vouchers** : générés en masse (événements), à durée limitée, à usage unique si possible ; traçabilité (quel voucher → qui → quand).
- **SMS** : passerelle SMS avec contrat opérateur ; prévoir le cas « pas de réseau mobile » (voucher de secours).
- **Sponsor** : un employé valide l'invité (workflow mail) — adapté aux visiteurs réguliers.
- **Sécurité** : HTTPS sur le portail, isolation des invités, débit plafonné, durée de session, journalisation (obligations locales — **à vérifier**).
- **Test** : le test d'acceptation = une personne non-technique qui se connecte en moins de 2 minutes, sans aide.

## 165. Commandes de vérification côté équipement — mémo

NCE est la source de vérité, mais le dépannage passe parfois par le CLI des équipements. Mémo des vérifications utiles (syntaxe indicative Huawei VRP — adapter à la version) :

- État général : `display version`, `display device`, `display cpu-usage`, `display memory-usage`.
- Interfaces : `display interface brief`, `display interface <nom>` (erreurs CRC, collisions).
- VLAN : `display vlan`, `display port vlan`.
- 802.1X : `display dot1x` (sessions, statistiques d'authentification).
- WLAN (WAC) : `display ap all` (état des AP), `display radio all`, `display station all` (clients).
- PoE : `display poe power-state` / informations de puissance par port.
- LLDP : `display lldp neighbor brief` (confronter à la topologie NCE).
- Logs : `display logbuffer`, `display trapbuffer`.

Règle : toute modification faite en CLI pendant un dépannage est **retranscrite ensuite dans le template** (sinon l'écart de conformité la signalera — et c'est normal).

## 166. Exemple de template — squelette commenté

Squelette indicatif d'un template « switch d'accès » (pseudo-syntaxe — la syntaxe exacte dépend de l'éditeur de templates de la version NCE) :

```
# Template: SW-ACC-STD v1.2
# Variables: {{NOM_SITE}}, {{VLAN_GESTION}}, {{VLAN_USERS}}, {{VLAN_VOIP}}, {{NTP_SRVR}}

system-view
 sysname {{NOM_SITE}}-{{NOM_EQPT}}
vlan {{VLAN_GESTION}} / description GESTION
vlan {{VLAN_USERS}}  / description USERS-802.1X
vlan {{VLAN_VOIP}}   / description VOIP
# Management
interface Vlanif{{VLAN_GESTION}} / ip address {{IP_GESTION}} {{MASQUE}}
# NTP, DNS, SNMPv3 (variables)
ntp-service unicast-server {{NTP_SRVR}}
# 802.1X sur les ports d'accès (boucle sur {{PORTS_ACCES}})
dot1x enable (par port)
# Uplink en trunk avec les VLAN autorisés
```

Principes du squelette : **variables en haut** (documentées), sections **commentées**, pas de valeur en dur spécifique à un site, version dans l'en-tête. Ce squelette se teste en maquette avant industrialisation (section 83).

## 167. Seuils d'alerte recommandés — point de départ

Point de départ à ajuster selon l'historique (jamais de seuils « sortis du chapeau » en production sans période d'observation) :

| Métrique | Seuil d'attention | Seuil critique |
|---|---|---|
| CPU équipement | > 70 % pendant 15 min | > 90 % pendant 5 min |
| Mémoire équipement | > 80 % | > 90 % |
| Utilisation de lien | > 70 % (planifier l'upgrade) | > 90 % pendant 15 min |
| Erreurs CRC sur port | > 0,1 % des paquets | augmentation brutale |
| Température | > 50 °C (selon spec) | > 60 °C / seuil constructeur |
| AP hors ligne | 1 AP (info) | > 10 % des AP d'un site |
| Échecs 802.1X | pic inhabituel vs baseline | > 20 % des tentatives sur 15 min |
| Espace disque NCE | > 70 % | > 85 % |
| Certificats | J-90 / J-60 | J-30 / J-7 |

Après 3 mois d'historique : remplacer ces seuils génériques par des **baselines** (le ML de CampusInsight aide — section 95).

## 168. Scripts d'exploitation — exemples d'usage API

Exemples de scripts à développer avec l'API NBI (logique — endpoints à valider sur la version) :

1. **Export d'inventaire quotidien** : GET équipements → CSV/CMDB (réconciliation avec GLPI).
2. **Rapport de conformité hebdo** : équipements en écart → mail au référent NCE.
3. **Création de tickets** : alarmes critiques → tickets GLPI automatiques (avec l'ID d'alarme NCE pour traçabilité).
4. **Check de certificats** : liste des certificats + expirations → alerte à J-60.
5. **Nettoyage** : lister les équipements « non enregistré » depuis > 7 jours → relance ou suppression.

