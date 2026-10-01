---
id: collect-261001-rattrapage/rattrapage/glpi-guide-3
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "apache"]
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [99, 277]
sha256: b3a113d65c1c93ad67fa3b9050bf34b7b9a0a14506d813b2617219b527db48f5
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

1. **Helpdesk / ITSM** : réception, qualification, assignation, suivi et clôture des
   tickets d'incidents et de demandes. C'est le guichet unique entre les utilisateurs
   (ou les clients, dans le cas d'une société de maintenance) et les techniciens.
2. **Gestion de parc / CMDB** : inventaire exhaustif du matériel (ordinateurs, imprimantes,
   copieurs, équipements réseau, onduleurs...), des logiciels, des licences, des contrats,
   des consommables et des budgets.

Pour un **chef de service systèmes & énergies** qui dirige aussi la maintenance des
copieurs, GLPI devient :

- le **registre unique** de chaque copieur client : marque, modèle, numéro de série,
  compteur, contrat de maintenance, historique des interventions ;
- le **canal officiel** des demandes d'intervention : fini les appels « au technicien
  que je connais » qui ne laissent aucune trace ;
- le **tableau de bord** du service : tickets ouverts par technicien, respect des SLA,
  taux de résolution au premier passage, consommation de toner par client ;
- la **mémoire** du service : base de connaissances alimentée par les retours terrain
  (codes erreur copieurs, procédures de maintenance préventive).

> **Règle d'or :** si ce n'est pas dans GLPI, ça n'existe pas. Chaque intervention,
> chaque compteur relevé, chaque cartouche posée doit laisser une trace.

**Ce que GLPI n'est pas :** ce n'est pas un outil de supervision temps réel
(remplacé par Zabbix, Nagios...), ni un ERP comptable, ni un GMAO pure
(même s'il s'en rapproche beaucoup avec les tickets récurrents et les projets).

---

## 2. Concepts fondamentaux : ITSM, CMDB, entités, profils

### ITSM (IT Service Management)

L'ITSM désigne l'ensemble des bonnes pratiques de gestion des services informatiques,
structurées notamment par le référentiel **ITIL**. GLPI implémente les processus ITIL
essentiels :

| Processus ITIL | Équivalent GLPI |
|---|---|
| Gestion des incidents | Tickets (statut, urgence, impact, priorité) |
| Gestion des demandes | Tickets de type « Demande » |
| Gestion des problèmes | Tickets liés, base de connaissances |
| Gestion des changements | Projets + tickets (plugin Change possible) |
| Gestion de la configuration (CMDB) | Parc : matériels, logiciels, liaisons |
| Gestion des niveaux de service | SLA (temps de prise en compte / résolution) |
| Gestion des connaissances | Base de connaissances |

### CMDB (Configuration Management Database)

La CMDB de GLPI, c'est l'onglet **Parc** : chaque élément (ordinateur, imprimante,
copieur, switch...) est un **élément de configuration (CI)** avec ses caractéristiques,
ses liaisons (branché à, installé sur, géré par) et son historique. L'intérêt : quand un
copieur tombe en panne, le technicien voit d'un coup d'œil le modèle, le contrat, les
tickets passés et les consommables compatibles.

### Les trois piliers du modèle GLPI

1. **Entités** : découpage organisationnel (société → sites → services). Tout objet
   GLPI appartient à une entité (ou est « récursif » sur plusieurs).
2. **Profils** : ensembles d'habilitations (ce qu'un utilisateur a le droit de voir/faire).
3. **Utilisateurs** : rattachés à une ou plusieurs entités avec un ou plusieurs profils.

> **Analogie métier :** l'entité, c'est l'agence ou le site client ; le profil, c'est la
> fiche de poste (technicien, superviseur, simple déclarant) ; l'utilisateur, c'est la personne.

---

## 3. Architecture technique de GLPI 10.x

GLPI 10.x est une application **PHP 8** adossée à **MariaDB/MySQL**, servie par
**Apache** ou **nginx**. Schéma simplifié :

```
Navigateur ──HTTPS──▶ Apache/nginx ──▶ PHP-FPM / mod_php ──▶ GLPI (/var/www/glpi)
                                                        │
                                                        ▼
                                              MariaDB (base glpidb)
                                                        │
Agent FusionInventory ──HTTPS──▶ GLPI ◀── cron système ──┘
        (inventaire)              (tâches planifiées : mail, SLA, alertes)
```

### Répertoires importants (installation standard)

| Répertoire | Contenu |
|---|---|
| `/var/www/glpi` (ou `/var/www/html/glpi`) | Code applicatif |
| `/var/www/glpi/config` | `config_db.php` (accès BDD) — **à protéger** |
| `/var/www/glpi/files` | Documents joints, dumps, exports — **hors web si possible** |
| `/var/www/glpi/marketplace` | Plugins installés via le marketplace |
| `/var/www/glpi/plugins` | Plugins manuels (legacy) |
| `/var/lib/glpi` (recommandé) | Données (`files/`, `config/`) déplacées hors du DocumentRoot |

> **Bonne pratique GLPI 10 :** déplacer `config/` et `files/` hors de la racine web
> (variables `GLPI_CONFIG_DIR` et `GLPI_VAR_DIR`). Voir section 45.

### Composants et flux

- **Tâches planifiées (cron)** : envoi des notifications, calcul des SLA, purge,
  remontée des collecteurs mail, inventaire. Sans cron, GLPI est « aveugle ».
- **Collecteurs mail** : GLPI lit une boîte (IMAP) et transforme chaque courriel en ticket.
- **API REST** : `/apirest.php` — création de tickets et relevés de compteurs depuis
  des scripts externes.
- **Agent FusionInventory** : remonte l'inventaire (matériel, logiciels, SNMP réseau).

---

## 4. Prérequis serveur détaillés (Debian/Ubuntu)

**Cible :** Debian 12 (Bookworm) ou Ubuntu 22.04/24.04 LTS, GLPI **10.0.x**.

### Matériel minimal / recommandé

| Profil | vCPU | RAM | Disque |
|---|---|---|---|
| Test / < 50 tickets/mois | 2 | 4 Go | 40 Go |
| Production PME (5-15 techniciens) | 4 | 8 Go | 100 Go SSD |
| Production + inventaire SNMP large | 4-8 | 16 Go | 200 Go SSD |

### Extensions PHP obligatoires (GLPI 10.x)

```
mysqli, curl, gd, intl, ldap, apcu, xml, mbstring, zip, bz2,
openssl, json, session, simplexml, domxml (dom), fileinfo,
sodium (recommandé), exif (recommandé), zend-opcache (recommandé)
```

### Base de données

- **MariaDB ≥ 10.6** (recommandé) ou MySQL ≥ 8.0.
- Jeu de caractères : `utf8mb4`, collation `utf8mb4_unicode_ci`.
- Privilèges : un utilisateur dédié `glpi` avec droits complets **sur la seule base** `glpidb`.

### Réseau et système

- Nom DNS interne (ex. `glpi.entreprise.lan`) + certificat TLS (section 44).
- NTP synchronisé (les SLA et les horodatages en dépendent !).
- Accès SMTP sortant (relais interne ou fournisseur) pour les notifications.
- Accès IMAP si collecteur mail utilisé.
- Pare-feu : ouvrir 443 (et 80 pour la redirection) depuis le LAN ; restreindre
  l'administration (SSH) aux postes d'exploitation.

> **Avertissement de version :** GLPI 10.x exige **PHP 8.1 minimum**
> (8.2/8.3 supportés selon la sous-version). PHP 7.4 n'est **plus** supporté.
> Vérifiez toujours la matrice de compatibilité de la version exacte téléchargée.

---

## 5. Installation pas à pas — pile LAMP (Apache + PHP 8 + MariaDB)

Procédure complète sur **Debian 12**. Adaptez les noms de paquets pour Ubuntu
(`php8.2-*` / `php8.3-*` selon la version distribuée).

### Étape 1 — mise à jour du système

```bash
sudo apt update && sudo apt full-upgrade -y
sudo apt install -y curl wget unzip ca-certificates gnupg lsb-release
```

### Étape 2 — installation d'Apache et MariaDB

```bash
sudo apt install -y apache2 mariadb-server
sudo systemctl enable --now apache2 mariadb
sudo mysql_secure_installation   # définir le mot de passe root, réponses Y
```

### Étape 3 — installation de PHP 8 et des extensions

```bash
sudo apt install -y php php-mysql php-curl php-gd php-intl php-ldap \
  php-apcu php-xml php-mbstring php-zip php-bz2 php-sodium php-exif \
  php-opcache libapache2-mod-php
php -v        # vérifier : PHP 8.2 ou 8.3
php -m | grep -E 'mysqli|curl|gd|intl|ldap|apcu'
```

### Étape 4 — création de la base de données

