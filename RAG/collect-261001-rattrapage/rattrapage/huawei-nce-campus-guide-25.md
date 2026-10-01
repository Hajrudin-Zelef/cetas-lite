---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-25
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [2224, 2340]
sha256: fff77c137563b419cc37a6975f388258b0ad21d88760f98ce7f43f0f54144a08
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

Bonnes pratiques de script : compte API dédié (section 89), gestion du token (re-login sur 401), journalisation, versionnement (Git), exécution planifiée (cron) avec alerte sur échec du script lui-même.

## 169. Plan de formation — qui apprend quoi

| Public | Contenu | Durée indicative | Format |
|---|---|---|---|
| Référent NCE (+ suppléant) | Architecture, templates, ZTP, 802.1X, API, dépannage avancé | 5-10 jours + 3 mois de pratique | Formation officielle + maquette |
| Administrateurs réseau | Onboarding, supervision, maintenance, rollback | 3-5 jours | Formation + compagnonnage |
| Opérateurs NOC | Dashboards, alarmes, diagnostic de niveau 1, escalade | 1-2 jours | Atelier pratique |
| Techniciens sites | Fiche ZTP, branchement, tests de base, escalade | 0,5 jour | Fiche + démo |
| Direction / métiers | Ce que NCE change (et ne change pas) pour eux | 1 heure | Présentation |

Budgéter la formation **dans** le projet (pas après), et prévoir le **recyclage** (nouveautés de version, 1 jour/an). Le suppléant du référent n'est pas un luxe : c'est la continuité de service.

## 170. RACI d'exploitation NCE — qui fait quoi

Exemple de matrice RACI (à adapter) :

| Activité | Référent NCE | Admin réseau | NOC | Technicien site | Direction |
|---|---|---|---|---|---|
| Créer/modifier un template | R/A | C | I | — | — |
| Déployer un template (vague) | A | R | I | I | I |
| Onboarding ZTP (site) | A | C | I | R | — |
| Acquitter une alarme | I | C | R/A | — | — |
| Upgrade firmware (vague) | A | R | I | I | I |
| Gérer les licences | R/A | I | — | — | C (budget) |
| Renouveler un certificat | A | R | I | — | — |
| Restaurer NCE (sinistre) | R/A | C | I | — | I |
| Rapport mensuel | R | C | I | — | A (destinataire) |

(R = réalise, A = approuve/responsable, C = consulté, I = informé.) Sans RACI, tout le monde attend que « quelqu'un » s'occupe des licences — et personne ne le fait.

## 171. KPI d'exploitation — piloter le réseau comme un service

Indicateurs à suivre mensuellement (tableau de bord du chef de service) :

- **Disponibilité** : % par site (objectif ex : 99,5 % — à définir contractuellement).
- **MTTR** : temps moyen de résolution des incidents (par sévérité).
- **Conformité** : % d'équipements sans écart de configuration (objectif : 100 %).
- **Onboarding** : délai moyen de mise en service d'un équipement (ZTP vs manuel).
- **Licences** : % de consommation du pool device-days.
- **Changements** : nombre de déploiements, taux d'échec, taux de rollback.
- **Sécurité** : échecs 802.1X, équipements non conformes, délai de patch des vulnérabilités.
- **Expérience** : plaintes Wi-Fi / 100 utilisateurs, temps moyen d'association.

Un KPI sans objectif ni action associée est une statistique décorative. Revoir les objectifs **annuellement**.

## 172. Fiche réflexe incident — une page pour l'astreinte

Modèle de fiche (à afficher / garder sous la main) :

```
INCIDENT RÉSEAU — FICHE RÉFLEXE
1. QUALIFIER : quel site ? quels services impactés ? depuis quand ? (NCE : dashboard)
2. ISOLER : le problème est-il local (1 équipement) ou global (site) ?
3. VÉRIFIER : y a-t-il eu un CHANGEMENT récent ? (template, firmware, énergie ?)
4. ÉNERGIE : l'onduleur du site est-il OK ? (corrélation Zabbix — cas 25)
5. DÉGRADER PROPREMENT : si 802.1X bloque tout → mode ouvert temporaire (procédure)
6. ESCALADER : au-delà de 30 min sans diagnostic → référent NCE + direction si critique
7. COMMUNIQUER : message aux utilisateurs (canal prévu), mise à jour toutes les 30 min
8. TRACER : tout noter (heures, actions) pour le REX
9. REX : sous 72 h, fiche incident + actions correctives (jamais de REX = incident qui reviendra)
NE JAMAIS : multiplier les changements à l'aveugle / travailler sans filet / cacher un incident.
```

## 173. Lexique des erreurs à ne pas commettre — top 10

1. Déployer NCE sans **TCO** ni pilote — signer sur une promesse commerciale.
2. Sous-dimensionner le serveur — économiser 2 000 € pour perdre des semaines.
3. Oublier les **flux réseau** (pare-feu) avant l'onboarding — 80 % des échecs ZTP.
4. Pousser un template **non testé** en production — le big bang du vendredi soir.
5. Faire du 802.1X sans **mode monitor** — le lundi matin noir.
6. Négliger les **certificats** — l'incident le plus bête (cas 15).
7. Ne pas **tester la restauration** — un backup non testé n'existe pas.
8. Laisser les techniciens configurer en **CLI sauvage** — la conformité s'effondre.
9. Oublier la **licence** au budget annuel — la grâce de 30 jours n'est pas une stratégie.
10. Ne pas **documenter** — quand le référent part, NCE devient une boîte noire.

---

# PARTIE 20 — GOUVERNANCE, RECETTE ET CADRAGE FOURNISSEUR

## 174. Interconnexion avec l'USG6000 — où s'arrête NCE, où commence le firewall

Dans le parc de Zelef, l'USG6000 protège le périmètre. Articulation avec NCE-Campus :

- **NCE-Campus** pilote le campus (switches, AP/WAC, AR, politiques d'accès utilisateurs, QoS).
- **L'USG6000** applique les politiques de sécurité inter-zones (filtrage, NAT, VPN IPsec/SSL, prévention d'intrusion selon la licence).
- Les deux se **complètent** : NCE authentifie et segmente (qui accède à quoi, VLAN/VN), l'USG filtre les flux entre zones et vers Internet.
- Vérifier dans la **matrice de compatibilité** ce que NCE peut faire sur l'USG (supervision ? déploiement de certaines politiques ? activation de licences via NCE — la fonction existe dans la documentation Monitoring and O&M).
- Règle d'architecture : **ne jamais** gérer les règles firewall uniquement depuis NCE si l'équipe sécurité a son processus sur l'USG — définir qui administre quoi (RACI, section 170), sinon les politiques se contredisent.

## 175. Plan d'adressage — méthode

Un plan d'adressage propre conditionne tout (ZTP, supervision, dépannage) :

1. **Blocs par usage** : management (/24 ou plus par région), utilisateurs (par site/bâtiment), voix, IoT, invités, serveurs, interconnexions, loopbacks.
2. **Hiérarchie** : un supernet par site (ex : 10.<site>.0.0/16) avec des /24 par usage — la lecture d'une IP révèle le site et l'usage.
3. **Réservations** : adresses de management en début de plage (passerelle, NCE, WAC, serveurs), DHCP pools documentés (début/fin, exclusions).
4. **DNS** : nommage systématique (sw-acc-01.site.domaine, ap-101.site.domaine) — un équipement sans nom DNS est un équipement qu'on ne retrouvera pas à 3 h du matin.
5. **Documentation** : le plan d'adressage est un **document versionné** (pas un tableau blanc), relu à chaque nouveau site.
6. **IPv6** : même si non déployé aujourd'hui, réserver la réflexion (les équipements récents le supportent — éviter de devoir tout renuméroter dans 5 ans).

## 176. AAA et mots de passe des équipements — politique

Même pilotés par NCE, les équipements gardent des accès locaux (console, SSH de secours) — à sécuriser :

- **Comptes locaux** : un compte admin local par équipement (mot de passe **unique par équipement**, stocké en coffre — jamais le même partout).
- **AAA centralisé** : authentifier les accès SSH/TACACS+/RADIUS sur l'annuaire quand possible (traçabilité nominative).
- **Rotation** : politique de changement (ex : annuelle, et à chaque départ d'un admin).
- **Console** : protéger l'accès physique (locaux fermés — un accès console = un contournement total).
- **NCE comme relais** : l'accès SSH via NCE (jump) centralise les logs — à privilégier sur l'accès direct.

## 177. Procédure de recette — valider avant de payer

Chaque jalon (installation, pilote, vague) se termine par une **recette** formelle :

