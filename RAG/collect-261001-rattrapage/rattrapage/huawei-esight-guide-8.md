---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-8
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [1106, 1280]
sha256: fb60f8ea4a1c954902a5d57ecb1b6f34ba49d512c5a8d9f4e3be0d134113fdbe
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

- **Dashboard NOC** : topologie + alarmes courantes critiques/majeures +
  top 10 des liens les plus chargés + équipements injoignables.
- **Dashboard chef de service** : nb d'alarmes par sévérité (semaine),
  disponibilité par site (%), top 5 des équipements les plus alarmés,
  tickets ouverts liés réseau.
- **Dashboard énergie** (votre casquette !) : état des alimentations,
  onduleurs supervisés, température des locaux techniques.

Rituel : le dashboard chef de service s'ouvre **tous les matins** en
5 minutes. Si un chiffre déraille, on creuse avant la réunion d'équipe.

## 54. NTA / NetStream : qui consomme la bande passante

Le composant **NTA (Network Traffic Analyzer)**, en édition Standard,
exploite les flux **NetStream** exportés par les équipements Huawei
(équivalent NetFlow/IPFIX) pour répondre à : *qui parle à qui, avec quel
protocole, quel volume ?*

Usages terrain :
- Identifier le **top talker** qui sature un lien (sauvegarde qui part
  en pleine journée, poste infecté, réplication mal planifiée).
- Vérifier qu'une **QoS** est respectée (le trafic prioritaire passe-t-il
  vraiment en priorité ?).
- Enquêtes « le réseau est lent » : prouver par les chiffres où part
  la bande passante.

Prérequis : activer l'export NetStream côté équipements vers eSight
(charge CPU non négligeable sur les petits modèles — à tester) et
prévoir l'espace disque (les données de flux sont volumineuses).

## 55. Rapports : la preuve du travail accompli

eSight fournit des **rapports prédéfinis** et un **concepteur de rapports**
personnalisés (d'après le constructeur : *« predefined reports and an
easy-to-use report design function »*).

Rapports à mettre en place d'office :
- **Disponibilité mensuelle par site/équipement** (le % qui va au COPIL).
- **Top des alarmes** (qu'est-ce qui nous pourrit la vie ? → plan d'action).
- **Capacité** : liens > 70 %, équipements CPU/RAM tendus (anticiper les
  investissements).
- **Inventaire** : parc supervisé, versions logicielles (préparer les
  montées de version).
- **Activité d'exploitation** : alarmes traitées, temps d'acquittement
  (le reporting de l'équipe).

**Automatiser** : génération et envoi par e-mail en début de mois.
Un rapport que personne ne lit ne sert à rien : 1 page de synthèse pour
la direction, le détail pour l'équipe.

## 56. SLA réseau : le diagnostic périodique automatique

Le composant **network SLA management** diagnostique automatiquement et
périodiquement les chemins réseau (type ping/traceroute supervisés,
mesures de latence/perte).

- Définir les **chemins critiques** : siège ↔ datacenter, siège ↔ sites
  distants, vers les applications SaaS critiques.
- Seuils : latence et perte de paquets (ex. : alerte si perte > 1 % ou
  latence > 2× la normale sur 15 min).
- En cas de plainte « l'application X est lente », le graphe SLA dit
  immédiatement si le réseau est en cause — **fini les guerres de
  tranchées entre équipes réseau et applicatives**.

---

# 9. GESTION DES CONFIGURATIONS

## 57. Sauvegarde automatique des configurations

C'est l'une des fonctions les plus rentables d'eSight : la **sauvegarde
périodique et automatique des fichiers de configuration** des équipements
(d'après le constructeur : sauvegarde immédiate, périodique, et déclenchée
par changement de configuration).

Mise en place :
1. Vérifier l'accès **SSH/Telnet** (ou SNMP selon version) depuis eSight
   vers chaque équipement avec un compte dédié.
2. Planifier la sauvegarde : **quotidienne** (nuit) + **à chaque changement
   détecté** si la version le permet.
3. Conserver N générations (ex. : 30 jours glissants) — vérifier l'espace
   disque alloué.
4. **Tester la restauration** sur un équipement de labo : une sauvegarde
   non testée n'est pas une sauvegarde.

**Le jour où un switch meurt à 3h du matin**, pouvoir pousser la dernière
config connue sur le matériel de remplacement en 10 minutes au lieu de
la reconstruire de mémoire en 3 heures : c'est ça, le ROI de cette fonction.

## 58. Comparaison de configurations : voir ce qui a changé

eSight permet de **comparer** deux fichiers de configuration (ex. : la
version d'hier vs celle d'aujourd'hui, ou la config d'un site vs le
modèle de référence).

Usages :
- « Ça marchait hier, ça ne marche plus » → diff des configs, la cause
  est dedans 9 fois sur 10.
- **Audit de conformité** : comparer les configs du parc à un
  template de référence (mot de passe, SNMP, syslog, NTP, banners).
- Avant/après une intervention : figer la config avant, comparer après.

## 59. Restauration : remettre une configuration

- Restauration **manuelle** : choisir une génération sauvegardée et la
  pousser vers l'équipement (via SSH/Telnet).
- **Procédure encadrée** : toute restauration se fait en fenêtre de
  maintenance, avec config actuelle sauvegardée juste avant (filet de
  sécurité), et test de non-régression après.
- **Ne jamais restaurer « à l'aveugle »** une config de plus de quelques
  jours : entre-temps, des changements légitimes ont pu intervenir
  (nouvelles VLAN, routes). Comparer d'abord (section 58).

## 60. Smart Configuration Tool : déployer en masse

Le **Smart Configuration Tool** (inclus dès l'édition Compact) permet
d'appliquer des configurations ou des scripts à **plusieurs équipements**
en une fois.

Cas d'usage :
- Déployer un changement standard sur 50 switch d'accès (ex. : ajouter
  le serveur syslog, changer la community SNMP).
- Montées de version logicielles planifiées (firmware) par lots.

**Garde-fous :**
- Toujours tester sur **2-3 équipements pilotes** avant généralisation.
- Prévoir le **plan de rollback** (config sauvegardée avant, section 57).
- Tracer : qui a lancé quoi, quand, sur quels équipements (audit,
  section 11).

## 61. Inventaire, ressources physiques et étiquettes électroniques

Au-delà des configs, eSight gère l'**inventaire** :
- **Ressources physiques** : châssis, cartes, alimentations, SFP —
  savoir ce qui est *dans* chaque équipement sans se déplacer.
- **Étiquettes électroniques** (electronic labels) : les infos d'identification
  lues dans l'équipement (numéros de série, modèles) — précieuses pour
  le suivi de garantie et les RMA.
- **IP topology / gestion des liens** : inventaire des interconnexions.

**Exploitation** : export mensuel de l'inventaire (CSV) vers votre CMDB
ou votre tableau de suivi de parc. C'est la base du plan de renouvellement
et des déclarations de garantie.

---

# 10. GESTION WLAN

## 62. Supervision des AC et des AP

Le module **WLAN management** (édition Standard minimum) supervise les
**contrôleurs (AC)** et les **points d'accès (AP)** :

- État des AP (en ligne/hors ligne), nombre d'utilisateurs associés
  par AP et par SSID.
- Santé des AC (CPU, mémoire, licences AP).
- **Topologie unifiée filaire + sans fil** : voir d'un coup d'œil le
  chemin complet d'un utilisateur Wi-Fi jusqu'au cœur de réseau.

**Seuil d'alerte à configurer d'office** : pourcentage d'AP hors ligne
par site (ex. : > 10 % des AP d'un site down = alerte major — un AP
isolé qui tombe, c'est du minor).

## 63. Diagnostic des pannes sans fil

Le module aide au **diagnostic des pannes du réseau sans fil** :

- Un utilisateur se plaint : retrouver son **historique de connexion**
  (quel AP, quel signal, quelles déconnexions) au lieu de lui demander
  « vous êtes où exactement ? ».
- Interférences / bruit radio : repérer les AP qui souffrent et
  envisager un ajustement de canaux/puissance.
- Saturation : repérer les AP avec trop d'utilisateurs simultanés →
  densifier ou équilibrer.

