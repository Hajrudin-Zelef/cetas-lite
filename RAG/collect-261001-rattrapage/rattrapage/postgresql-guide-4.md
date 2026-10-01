---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-4
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [678, 890]
sha256: 18bc016ae2fde4182fe737ba864480e9611b8837d5507d5bdf7123038356114a
---

# PostgreSQL en production — Guide complet pour sysadmin

-- Séquences (indispensable avec les serial/identity)
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO app_ecrivain;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
  GRANT USAGE, SELECT ON SEQUENCES TO app_ecrivain;

-- Révoquer
REVOKE DELETE ON capteurs FROM app_ecrivain;

-- Voir les droits effectifs
\z capteurs        -- ou \dp
SELECT grantee, privilege_type FROM information_schema.role_table_grants
WHERE table_name = 'capteurs';
```

> ⚠️ **Piège n°1 des permissions :** `GRANT ... ON ALL TABLES` ne couvre
> que les tables **existantes**. Sans `ALTER DEFAULT PRIVILEGES`, la
> prochaine table créée sera inaccessible à l'appli.

---

## 20. Schémas : organiser et cloisonner

```sql
CREATE SCHEMA supervision;              -- schéma dédié
CREATE TABLE supervision.alertes (...);

-- Isoler une appli dans son schéma
CREATE SCHEMA appli_energie;
GRANT USAGE ON SCHEMA appli_energie TO app_ecrivain;
ALTER ROLE app_ecrivain SET search_path = appli_energie, public;

-- search_path : ordre de résolution des noms non qualifiés
SHOW search_path;
```

**Bonnes pratiques :**

- Un schéma par application/domaine (`energie`, `reseau`, `supervision`).
- Ne laissez pas les applis écrire dans `public` sans contrôle.
- Depuis PG 15, `public` n'est plus inscriptible par `PUBLIC` par défaut
  (durcissement bienvenu : vérifiez vos migrations).

---

## 21. Row-Level Security (RLS) : introduction

La RLS filtre les lignes visibles **selon le rôle connecté**, au niveau
du serveur : impossible à contourner depuis l'appli.

```sql
-- Activer la RLS sur une table multi-sites
ALTER TABLE mesures ENABLE ROW LEVEL SECURITY;

-- Politique : chaque technicien ne voit que son site
CREATE POLICY site_isolation ON mesures
  FOR ALL
  USING (site_id = current_setting('app.site_id')::int);

-- Le propriétaire contourne par défaut ; forcer aussi pour lui :
ALTER TABLE mesures FORCE ROW LEVEL SECURITY;

-- Côté appli : positionner la variable à chaque connexion
SET app.site_id = '3';
SELECT * FROM mesures;   -- ne retourne que le site 3
```

**Cas d'usage :** SaaS multi-tenant, cloisonnement par site/agence,
données sensibles par profil. **Coût :** une condition ajoutée à chaque
requête — indexez la colonne de filtrage (`site_id`).

---

## 22. Checklist sécurité "mise en production"

- [ ] `listen_addresses` restreint au nécessaire
- [ ] `pg_hba.conf` : `scram-sha-256`, pas de `trust` en `host`
- [ ] SSL `on`, `sslmode=verify-full` côté applis critiques
- [ ] Rôles applicatifs non-superuser, `CONNECTION LIMIT` posée
- [ ] `ALTER DEFAULT PRIVILEGES` en place pour les futurs objets
- [ ] Mots de passe dans le coffre, jamais dans les scripts en clair
- [ ] `log_connections`, `log_disconnections`, `log_min_duration_statement`
- [ ] Compte `supervision` en lecture seule (`pg_monitor`, voir section 60)
- [ ] Sauvegarde testée (section 48) **avant** l'ouverture aux applis
- [ ] Firewall : 5432 ouvert uniquement vers les hôtes applicatifs

---

# Partie C — SQL opérationnel : DDL, types, DML

## 23. Créer une base et s'y connecter

```sql
-- Créer une base (depuis postgres ou template)
CREATE DATABASE inventaire
  OWNER app_admin
  ENCODING 'UTF8'
  LC_COLLATE 'fr_FR.UTF-8'
  LC_CTYPE 'fr_FR.UTF-8'
  TEMPLATE template0;   -- template0 = vierge, sans dépendance locale

-- Lister, puis s'y connecter
\l
\c inventaire
```

```bash
# Équivalents shell
createdb -O app_admin inventaire
dropdb inventaire   # destructif, demande confirmation si données
```

> ⚠️ `CREATE DATABASE` copie un template : impossible si des connexions
> sont actives sur la base source. `TEMPLATE template0` évite les
> surprises de locale.

---

## 24. Types de données : le tableau de référence

| Famille | Types | Quand l'utiliser |
|---|---|---|
| Entiers | `smallint`, `integer`, `bigint` | Compteurs, clés ; `bigint` pour les gros volumes |
| Séries | `GENERATED ... AS IDENTITY` | Clés primaires auto-incrémentées (moderne, remplace `serial`) |
| Numériques exacts | `numeric(p,s)`, `money` | Monnaie, énergie (kWh facturés) : **jamais** `float` pour l'argent |
| Flottants | `real`, `double precision` | Mesures physiques, calculs scientifiques |
| Texte | `text`, `varchar(n)`, `char(n)` | `text` par défaut ; `varchar(n)` = contrainte de longueur |
| Booléen | `boolean` | `TRUE/FALSE/NULL` |
| Dates/heures | `date`, `time`, `timestamp`, `timestamptz` | **Toujours `timestamptz`** en prod (fuseau géré) |
| Durées | `interval` | `INTERVAL '90 minutes'` |
| Réseau | `inet`, `cidr`, `macaddr` | Adresses IP : requêtes `>>` (contenu dans) natives |
| JSON | `jsonb` | Configs, payloads : indexable (GIN), **préférer à `json`** |
| UUID | `uuid` | Clés externes, identifiants publics (`gen_random_uuid()`) |
| Tableaux | `text[]`, `int[]` | Listes simples ; pas de jointure possible |
| Géométrie | `point`, extensions PostGIS | SIG (extension `postgis`) |
| Binaire | `bytea` | Petits blobs ; gros fichiers → stockage objet externe |

```sql
-- Exemples de création avec bons types
CREATE TABLE capteurs (
  id          bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  code        text NOT NULL UNIQUE,
  ip          inet,
  site_id     integer NOT NULL REFERENCES sites(id),
  config      jsonb NOT NULL DEFAULT '{}',
  installe_le timestamptz NOT NULL DEFAULT now()
);
```

**Règles pratiques :**

- `timestamptz` partout : stocke en UTC, affiche selon `TimeZone` session.
- `numeric` pour la facturation, `double precision` pour les mesures.
- `inet` > `varchar` pour les IP (validation + opérateurs réseau).

---

## 25. DDL : CREATE / ALTER / DROP sans se blesser

```sql
-- CREATE TABLE IF NOT EXISTS : idempotent pour les migrations
CREATE TABLE IF NOT EXISTS sites (
  id   integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  nom  text NOT NULL,
  code text NOT NULL UNIQUE
);

-- ALTER : ajouter une colonne avec défaut (rapide depuis PG 11,
-- le défaut est stocké en méta, pas réécrit)
ALTER TABLE capteurs ADD COLUMN IF NOT EXISTS actif boolean NOT NULL DEFAULT true;

-- Renommer (attention aux applis !)
ALTER TABLE capteurs RENAME TO sondes;
ALTER TABLE sondes RENAME COLUMN code TO reference;

-- Changer un type (USING si conversion non implicite)
ALTER TABLE mesures ALTER COLUMN valeur TYPE numeric USING valeur::numeric;

-- DROP avec garde-fou
DROP TABLE IF EXISTS _tmp_import;                 -- simple
DROP TABLE mesures CASCADE;                       -- + objets dépendants (dangereux)
```

**Migrations en production :**

1. Toujours `IF NOT EXISTS` / `IF EXISTS` dans les scripts rejouables.
2. `ALTER TABLE ... ADD COLUMN` prend un verrou `ACCESS EXCLUSIVE`
   bref ; sur une table énorme, préférez les heures creuses.
3. Testez chaque migration sur une copie (restauration, section 51).

---

## 26. Contraintes : la qualité des données côté serveur

| Contrainte | Rôle | Exemple |
|---|---|---|
| `PRIMARY KEY` | Unicité + NOT NULL (index B-tree implicite) | `id bigint PRIMARY KEY` |
| `FOREIGN KEY` | Intégrité référentielle | `REFERENCES sites(id)` |
| `UNIQUE` | Unicité (NULL multiples autorisés) | `UNIQUE(code)` |
| `NOT NULL` | Valeur obligatoire | `nom text NOT NULL` |
| `CHECK` | Règle métier | `CHECK (puissance_kva > 0)` |
| `EXCLUDE` | Non-chevauchement (plages) | Pas de réservations qui se chevauchent |
| `DEFAULT` | Valeur par défaut | `DEFAULT now()` |

```sql
CREATE TABLE onduleurs (
  id            bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  reference     text NOT NULL UNIQUE,
  puissance_kva numeric(8,2) NOT NULL CHECK (puissance_kva > 0),
  site_id       integer NOT NULL REFERENCES sites(id) ON DELETE RESTRICT,
  mise_en_service date NOT NULL DEFAULT CURRENT_DATE,
  CONSTRAINT puissance_coherente CHECK (puissance_kva <= 2000)
);

