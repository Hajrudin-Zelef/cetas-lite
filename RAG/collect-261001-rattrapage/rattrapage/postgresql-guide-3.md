---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-3
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [442, 677]
sha256: 081a7b12c3644d32423fde5a27230a8397e72bc38a9d118ea95b1f6b87405cd9
---

# Connexion TCP avec mot de passe (variable d'env, pas d'option -W en script)
PGPASSWORD='...' psql -h 10.0.0.5 -U app_lecteur -d inventaire

# Connexion avec fichier .pgpass (recommandé pour les scripts)
# ~/.pgpass : host:port:db:user:password, chmod 600
psql -h 10.0.0.5 -U app_lecteur -d inventaire -c "SELECT now();"

# Exécuter un fichier SQL
psql -h srv-bdd -U deploy -d inventaire -f migration_042.sql

# Mode non interactif pour scripts (erreurs -> exit 1)
psql -v ON_ERROR_STOP=1 -f migration.sql
```

**Options utiles :**

| Option | Effet |
|---|---|
| `-c "SQL"` | Exécute une commande puis quitte |
| `-f fichier` | Exécute un script |
| `-v ON_ERROR_STOP=1` | Stoppe au premier échec |
| `-X` | Ignore `~/.psqlrc` (scripts reproductibles) |
| `-q -t -A` | Sortie silencieuse/terse, pratique pour `$(...)` |
| `-x` | Affichage étendu (colonnes en lignes) |

---

## 14. Méta-commandes psql : l'essentiel

Les méta-commandes commencent par `\` et ne vont **pas** au serveur :
c'est `psql` qui les interprète.

| Commande | Affiche |
|---|---|
| `\l` / `\l+` | Bases de données |
| `\du` / `\du+` | Rôles et attributs |
| `\dn` | Schémas |
| `\dt` / `\dt+` | Tables (+ tailles) |
| `\di` | Index |
| `\dv` | Vues |
| `\df` | Fonctions |
| `\d nom` | Description d'une table (colonnes, index, contraintes) |
| `\d+ nom` | + tailles et descriptions |
| `\sf nom_fonction` | Code source d'une fonction |
| `\x` | Bascule affichage étendu |
| `\timing` | Affiche le temps d'exécution |
| `\pset null '[NULL]'` | Représentation des NULL |
| `\copy` | Import/export CSV côté **client** |
| `\i fichier.sql` | Exécute un fichier |
| `\o fichier` | Redirige la sortie |
| `\e` | Édite la dernière requête dans `$EDITOR` |
| `\set` / `\unset` | Variables psql |
| `\q` | Quitter |

```sql
-- Exemples en session
\dt+                  -- tables avec tailles
\d+ capteurs          -- détail d'une table
\x on                 -- lecture d'une ligne large
SELECT * FROM capteurs WHERE id = 42;
\timing on
SELECT count(*) FROM mesures;
```

> 💡 `\copy` (client) ≠ `COPY` (serveur) : `\copy` lit/écrit les fichiers
> **de votre poste**, `COPY` ceux **du serveur** (réservé au superuser).

---

## 15. Rôles, utilisateurs, groupes : le modèle PostgreSQL

PostgreSQL ne distingue pas utilisateurs et groupes : tout est un
**rôle** (`ROLE`). Un rôle avec `LOGIN` peut se connecter.

```sql
-- Créer un rôle applicatif (sans superuser !)
CREATE ROLE app_lecteur WITH LOGIN PASSWORD '...' ;
CREATE ROLE app_ecrivain WITH LOGIN PASSWORD '...';

-- Rôle groupe puis membres (bonne pratique)
CREATE ROLE equipe_reseau;                       -- NOLOGIN par défaut
GRANT equipe_reseau TO alice, bob;                -- membres
ALTER ROLE alice INHERIT;                         -- défaut : hérite des droits

-- Attributs importants
ALTER ROLE app_ecrivain WITH CONNECTION LIMIT 50; -- limite de connexions
ALTER ROLE batch SET statement_timeout = '5min';  -- paramètre par rôle

-- Voir les rôles
\du
SELECT rolname, rolsuper, rolcanlogin, rolconnlimit
FROM pg_roles ORDER BY rolname;
```

**Principes de production :**

1. **Jamais** l'application en superuser (`postgres`).
2. Un rôle par application/fonction : `app_lecteur`, `app_ecrivain`,
   `batch_nuit`, `supervision`.
3. Les humains passent par des rôles nominatifs membres de rôles groupes.
4. `CONNECTION LIMIT` pour éviter qu'un batch ne sature `max_connections`.

---

## 16. pg_hba.conf : qui se connecte, comment, d'où

`pg_hba.conf` (Host-Based Authentication) décide, **dans l'ordre des
lignes**, quelle méthode d'authentification s'applique.

```
# TYPE  DATABASE  USER       ADDRESS        METHOD
local   all       postgres                 peer
local   all       all                      peer
host    all       all        127.0.0.1/32   scram-sha-256
host    all       all        ::1/128        scram-sha-256
host    inventaire app_lecteur 10.0.1.0/24  scram-sha-256
host    all       all        0.0.0.0/0      reject
```

**Lecture :** première ligne qui correspond (type + base + user + adresse)
→ la méthode s'applique, les suivantes sont ignorées.

### 16.1. Méthodes d'authentification

| Méthode | Usage | Sécurité |
|---|---|---|
| `trust` | ⚠️ Aucune vérif. Dev/test local uniquement | Nulle |
| `peer` | Identité OS via socket Unix | Bonne (local) |
| `scram-sha-256` | Mot de passe chiffré (défaut moderne) | Bonne |
| `md5` | Obsolète, à migrer vers SCRAM | Faible |
| `cert` | Certificat client TLS | Très bonne |
| `ldap` / `radius` | Annuaire externe | Selon infra |
| `reject` | Refuse explicitement | — |

```bash
# Après modification : recharger (pas besoin de redémarrer)
sudo pg_ctlcluster 17 main reload
# ou : SELECT pg_reload_conf();

# Tester une règle depuis un hôte distant
psql "host=10.0.0.5 dbname=inventaire user=app_lecteur sslmode=require"
```

**Checklist pg_hba.conf :**

- [ ] Pas de `trust` sur `host` (jamais en prod)
- [ ] `scram-sha-256` partout où un mot de passe est utilisé
- [ ] Règles du plus spécifique au plus général, `reject` final explicite
- [ ] `peer` mappe l'OS user `postgres` au rôle `postgres` (pratique + sûr)

---

## 17. Migrer les mots de passe vers SCRAM-SHA-256

```sql
-- 1. Vérifier la méthode de stockage actuelle
SELECT rolname,
       CASE WHEN rolpassword LIKE 'SCRAM-SHA-256%' THEN 'scram'
            WHEN rolpassword LIKE 'md5%' THEN 'md5'
            ELSE 'autre/absent' END AS methode
FROM pg_authid WHERE rolcanlogin;

-- 2. Forcer SCRAM pour les nouveaux mots de passe
-- postgresql.conf :
-- password_encryption = 'scram-sha-256'   (défaut depuis la v14)

-- 3. Régénérer les mots de passe md5 (chaque rôle doit le refaire)
ALTER ROLE app_lecteur PASSWORD 'nouveau-mot-de-passe-du-coffre';

-- 4. Recharger et tester les connexions applicatives
```

> ⚠️ Changer `password_encryption` ne convertit **pas** les hash existants :
> il faut ré-émettre chaque mot de passe.

---

## 18. SSL/TLS : chiffrer les connexions

Sur Debian, le cluster génère un certificat auto-signé (`snakeoil`).
En production, utilisez un certificat de votre CA interne.

```ini
# postgresql.conf
ssl = on
ssl_cert_file = '/etc/postgresql/17/main/server.crt'
ssl_key_file  = '/etc/postgresql/17/main/server.key'  # chmod 600, owner postgres
ssl_ca_file   = '/etc/postgresql/17/main/ca.crt'      # si auth par certificat
```

```bash
# Déployer un certificat interne (exemple)
sudo cp srv-bdd.crt /etc/postgresql/17/main/server.crt
sudo cp srv-bdd.key /etc/postgresql/17/main/server.key
sudo chown postgres:postgres /etc/postgresql/17/main/server.*
sudo chmod 600 /etc/postgresql/17/main/server.key
sudo pg_ctlcluster 17 main restart   # ssl = postmaster -> restart

# Vérifier côté client
psql "host=srv-bdd dbname=inventaire user=app_lecteur sslmode=verify-full \
      sslrootcert=/etc/ssl/certs/ca-interne.crt" -c "SHOW ssl;"
```

**Modes `sslmode` côté client :**

| Mode | Chiffrement | Vérifie le certificat |
|---|---|---|
| `disable` | Non | Non |
| `require` | Oui | Non (vulnérable au MITM) |
| `verify-ca` | Oui | Signé par la CA |
| `verify-full` | Oui | CA + nom d'hôte ✅ |

**En production : `verify-full`** pour les applis, `require` minimum.

---

## 19. Droits fins : GRANT / REVOKE en détail

```sql
-- Droits sur une base (connexion)
GRANT CONNECT ON DATABASE inventaire TO app_lecteur;

-- Droits sur un schéma (USAGE pour voir les objets)
GRANT USAGE ON SCHEMA public TO app_lecteur;

-- Droits sur tables existantes
GRANT SELECT ON ALL TABLES IN SCHEMA public TO app_lecteur;
GRANT SELECT, INSERT, UPDATE ON capteurs TO app_ecrivain;

-- Droits sur les FUTURES tables (piège classique !)
ALTER DEFAULT PRIVILEGES IN SCHEMA public
  GRANT SELECT ON TABLES TO app_lecteur;
-- À exécuter pour chaque rôle créateur de tables :
ALTER DEFAULT PRIVILEGES FOR ROLE deploy IN SCHEMA public
  GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO app_ecrivain;

