---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-4
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei", "Oracle"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [439, 597]
sha256: 56c230cd6f666478974c15ca9908fa6b93e9f14cb7e27631fb131d49dd79162c
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

1. **Prenez le palier au-dessus** de votre parc réel. Un parc de 180
   équipements avec une croissance de 10 %/an dépasse 200 en un an :
   dimensionnez pour 200-500 dès le départ.
2. **La RAM est le premier levier** : si le budget est serré, privilégiez
   la RAM au CPU. eSight est gourmand en mémoire (services Java + base).
3. **Le disque dépend de la rétention** : le tableau suppose une rétention
   standard. Si vous gardez 1 an de données de performance en pas de
   5 minutes sur 500 équipements, prévoyez large (plusieurs centaines
   de Go) — et un contrôle de croissance (voir cas n°13).
4. **Virtualisez** : eSight tourne très bien en VM (malgré une installation
   jugée plus délicate que sur appliance). Avantages : snapshots avant
   mise à jour, redimensionnement à chaud, PRA simplifié.
5. **Séparez si gros parc** : au-delà de ~1 000 équipements, discutez avec
   le partenaire d'une architecture dédiée (base sur serveur séparé,
   cluster) — **à valider sur la documentation officielle**.

## 18. Check-list de dimensionnement avant achat

- [ ] Inventaire des équipements à superviser (par type : routeurs,
      switches, firewalls, AP, serveurs…) avec **+20 % de marge**.
- [ ] Modules eSight retenus (WLAN ? serveurs ? NTA ?).
- [ ] Durée de rétention souhaitée (alarmes : 6-12 mois ; perfs : 3-12 mois
      selon le pas de collecte).
- [ ] Édition choisie (Standard vs Professional : faut-il la HA ? le NMS
      hiérarchique ?).
- [ ] OS du serveur (Windows vs Linux — la HA/cluster exige Linux).
- [ Bloc de texte : ressources VM réservées (pas de surallocation CPU/RAM),
- [ ] Réseau : le serveur eSight joignable en SNMP/syslog depuis tous les
      équipements (routage + firewall).
- [ ] Sauvegarde : espace de stockage externe pour les backups eSight.

---

# 4. INSTALLATION

## 19. Prérequis d'installation

**Serveur :**
- Ressources selon le tableau de la section 16 (palier supérieur).
- OS supporté par votre version d'eSight (**à vérifier sur la
  documentation officielle** — les versions récentes supportent des
  Windows Server et Linux bien plus récents que le datasheet historique).
- Nom d'hôte fixe, **IP fixe**, résolution DNS correcte (le changement
  d'IP après installation est une source classique de pannes — voir
  cas n°15).
- Synchronisation horaire **NTP** (écarts d'horloge = alarmes incohérentes
  et corrélation faussée).

**Base de données :**
- Installée et accessible (MySQL / SQL Server / Oracle selon l'OS et
  la version — **à vérifier sur la documentation officielle**).
- Compte dédié eSight avec les droits requis, mot de passe robuste
  conservé au coffre.

**Réseau :**
- Connectivité IP entre eSight et tous les équipements à superviser.
- Firewall : ouvrir les flux SNMP (UDP 161/162), syslog (UDP 514),
  NetStream, SSH (TCP 22) **dans le sens équipement ↔ eSight** selon
  le besoin, et HTTPS vers les postes d'exploitation.
- Relais SMTP joignable pour les notifications mail ; passerelle SMS
  si notification SMS prévue.

**Licences :** fichiers/actes de licence disponibles avant l'installation
(voir section 7).

## 20. Check-list pré-installation (à cocher le jour J)

- [ ] Snapshot/clone de la VM (si virtualisé) **avant** de commencer.
- [ ] OS à jour, antivirus configuré avec exclusions sur les répertoires
      eSight et de la base (un antivirus qui verrouille les fichiers
      de base = eSight qui rame ou plante).
- [ ] NTP synchronisé et vérifié (`w32tm /query /status` ou `chronyc tracking`).
- [ ] IP fixe, hostname, DNS : notés dans le dossier d'exploitation.
- [ ] Base de données installée, test de connexion OK.
- [ ] Ports firewall ouverts et testés (telnet/nc vers les équipements
      pilotes).
- [ ] Fichiers d'installation eSight + documentation de la version
      téléchargés et vérifiés (checksum si fourni).
- [ ] Fenêtre de maintenance planifiée et communiquée (même si eSight
      est nouveau, les tests de découverte génèrent du trafic SNMP).
- [ ] Binôme : ne jamais installer seul un outil qui deviendra critique —
      à deux, on documente en même temps.

## 21. Installation pas à pas

Déroulé type (les libellés exacts d'écrans **dépendent de la version** —
**à vérifier sur la documentation officielle**) :

**Étape 1 — Lancer l'installeur.**
Sur appliance Huawei : option **one-click install** (retours d'exploitants :
« très simple »). Sur VM/serveur standard : lancer le setup, qui installe
les services eSight et prépare la base.

**Étape 2 — Paramétrer la connexion à la base de données.**
Saisir hôte, port, nom de base, compte et mot de passe dédiés.
*Point de vigilance :* tester la connexion avant de poursuivre ;
80 % des échecs d'installation viennent d'ici (mauvais port, compte
sans droits, firewall local).

**Étape 3 — Choisir les composants à installer.**
Ne cocher que les modules achetés/licenciés (section 3). Installer un
module non licencié = service en erreur au démarrage.

**Étape 4 — Définir les paramètres système.**
Répertoires d'installation et de données (mettre les données sur un
volume dédié si possible — ça simplifie les sauvegardes et le
redimensionnement), ports d'écoute.

**Étape 5 — Créer le compte administrateur initial.**
Compte `admin` initial (nom exact **à vérifier sur la documentation
officielle**). Lui attribuer immédiatement un mot de passe robuste
et le consigner au coffre. **Ne jamais laisser le mot de passe
d'installation par défaut.**

**Étape 6 — Finaliser et démarrer les services.**
L'installeur démarre les services eSight. Sur Linux, vérifier avec
les outils système (`systemctl status` / scripts fournis) ;
sur Windows, via la console des services.

**Étape 7 — Vérifications post-installation.**
- [ ] Les services eSight sont démarrés et stables (pas de restart en boucle).
- [ ] La console web répond en HTTPS.
- [ ] Login avec le compte admin initial : OK.
- [ ] L'outil **Environmental Health Check** (icône eSight Console →
      Tools) ne remonte aucune anomalie bloquante (il contrôle les
      paramètres de base de l'OS, de la base et du serveur).
- [ ] Date/heure correctes dans l'interface.

## 22. Premier login et tour du propriétaire

1. Ouvrez `https://<IP-eSight>:<port>/` (**port exact à vérifier sur la
   documentation officielle**).
2. Connectez-vous avec le compte administrateur initial.
3. **Changez immédiatement le mot de passe** du compte admin
   (politique : 12+ caractères, coffre-fort d'équipe).
4. Faites le tour des menus : Resource (inventaire), Topology, Fault
   (alarmes), Performance, Configuration, Report, System (administration).
5. Vérifiez la version exacte installée (menu Système / À propos) et
   **notez-la dans le dossier d'exploitation** : toutes les procédures
   futures s'y référeront.

## 23. Activation des licences

1. Récupérez vos fichiers/actes de licence (fournis par le partenaire).
2. Dans la console : menu Système/Licence (**libellé à vérifier sur la
   documentation officielle**) → import/activation.
3. Vérifiez que chaque module apparaît comme **licencié et actif**,
   et que le **nombre d'équipements autorisés** couvre votre parc + marge.
4. Testez immédiatement avec 2-3 équipements pilotes : si la licence
   bloque, mieux vaut le savoir le jour J que pendant la généralisation.
5. Archivez les licences dans le PRA (section 12) : sans elles, une
   réinstallation après sinistre est bloquée.

**Piège classique** : licence importée mais module correspondant non
installé (ou l'inverse) → fonctionnalités grisées ou erreurs obscures.
Alignez toujours *modules installés = modules licenciés*.

## 24. Sécurisation post-installation (à faire dans la semaine)

