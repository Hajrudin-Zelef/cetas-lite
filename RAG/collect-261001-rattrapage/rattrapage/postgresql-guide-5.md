---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-5
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["valuation"]
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [891, 1155]
sha256: 7d12e8290e5a74bc0e27a0cf9a92cf85364002a95db47ce7388a429513eda32e
---

# PostgreSQL en production — Guide complet pour sysadmin

-- Ajouter une contrainte sans bloquer (PG 12+) : NOT VALID puis VALIDATE
ALTER TABLE mesures ADD CONSTRAINT mesure_positive
  CHECK (valeur >= 0) NOT VALID;          -- rapide, ne scanne pas
ALTER TABLE mesures VALIDATE CONSTRAINT mesure_positive;  -- scan en SHARE UPDATE EXCLUSIVE
```

> 💡 `NOT VALID` + `VALIDATE` : la méthode propre pour ajouter un `CHECK`
> ou une `FOREIGN KEY` sur une table volumineuse **sans** verrou exclusif
> long.

---

## 27. Clés étrangères : comportements ON DELETE / ON UPDATE

```sql
CREATE TABLE interventions (
  id           bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  onduleur_id  bigint NOT NULL REFERENCES onduleurs(id)
                 ON DELETE RESTRICT      -- refuse si des interventions existent
                 ON UPDATE CASCADE,
  technicien_id bigint REFERENCES techniciens(id)
                 ON DELETE SET NULL      -- garde l'historique, annule le lien
);
```

| Action | Effet à la suppression du parent |
|---|---|
| `RESTRICT` / `NO ACTION` | Refuse (défaut prudent) |
| `CASCADE` | Supprime les enfants (dangereux si mal compris) |
| `SET NULL` | Met la colonne à NULL |
| `SET DEFAULT` | Met la valeur par défaut |

**En production :** `RESTRICT` par défaut sur les données métier critiques ;
`CASCADE` uniquement sur des tables techniques liées (lignes de détail).

---

## 28. INSERT / UPDATE / DELETE : DML de base

```sql
-- INSERT simple et multiple
INSERT INTO sites (nom, code) VALUES ('Siège', 'SIEGE');
INSERT INTO sites (nom, code) VALUES
  ('Atelier', 'ATEL'), ('Dépôt', 'DEPOT')
RETURNING id;                                   -- récupère les id générés

-- UPSERT : insère ou met à jour en cas de conflit (idempotent !)
INSERT INTO capteurs (code, site_id)
VALUES ('SONDE-01', 1)
ON CONFLICT (code) DO UPDATE SET site_id = EXCLUDED.site_id;

-- UPDATE avec jointure
UPDATE mesures m
SET valide = false
FROM capteurs c
WHERE m.capteur_id = c.id AND c.actif = false;

-- DELETE avec RETURNING (audit avant suppression réelle)
DELETE FROM _tmp_import WHERE importe_le < now() - INTERVAL '30 days'
RETURNING *;

-- TRUNCATE : vide vite (sans MVCC ligne à ligne), réinitialise les séquences
TRUNCATE TABLE _tmp_import RESTART IDENTITY;
```

> ⚠️ `UPDATE`/`DELETE` sans `WHERE` : commencez toujours par le `SELECT`
> équivalent, ou travaillez en transaction (section 45).

---

## 29. SELECT : jointures et anti-jointures

```sql
-- Jointure interne
SELECT c.code, s.nom AS site
FROM capteurs c
JOIN sites s ON s.id = c.site_id;

-- LEFT JOIN : garder les capteurs même sans mesures
SELECT c.code, count(m.id) AS nb_mesures
FROM capteurs c
LEFT JOIN mesures m ON m.capteur_id = c.id
GROUP BY c.code;

-- Anti-jointure : capteurs SANS mesures (NOT EXISTS > NOT IN avec NULL)
SELECT c.code
FROM capteurs c
WHERE NOT EXISTS (
  SELECT 1 FROM mesures m WHERE m.capteur_id = c.id
);

-- LATERAL : top-N par groupe
SELECT s.nom, t.*
FROM sites s
CROSS JOIN LATERAL (
  SELECT * FROM mesures m
  WHERE m.site_id = s.id
  ORDER BY m.horodatage DESC LIMIT 5
) t;
```

---

## 30. CTE (WITH) : requêtes lisibles et maintenables

```sql
-- CTE simple : dernière mesure par capteur
WITH dernieres AS (
  SELECT capteur_id, max(horodatage) AS h
  FROM mesures GROUP BY capteur_id
)
SELECT c.code, m.valeur, m.horodatage
FROM dernieres d
JOIN capteurs c ON c.id = d.capteur_id
JOIN mesures m ON m.capteur_id = d.capteur_id AND m.horodatage = d.h;

-- CTE récursive : arborescence (ex. sites parent/enfant)
WITH RECURSIVE arborescence AS (
  SELECT id, nom, parent_id, 1 AS niveau
  FROM sites WHERE parent_id IS NULL
  UNION ALL
  SELECT s.id, s.nom, s.parent_id, a.niveau + 1
  FROM sites s JOIN arborescence a ON s.parent_id = a.id
)
SELECT * FROM arborescence ORDER BY niveau, nom;

-- CTE chaînées pour un rapport
WITH par_site AS (
  SELECT site_id, count(*) AS nb FROM capteurs GROUP BY site_id
),
alertes AS (
  SELECT site_id, count(*) AS nb FROM supervision.alertes
  WHERE resolue = false GROUP BY site_id
)
SELECT s.nom, coalesce(p.nb,0) AS capteurs, coalesce(a.nb,0) AS alertes
FROM sites s
LEFT JOIN par_site p ON p.site_id = s.id
LEFT JOIN alertes a ON a.site_id = s.id;
```

> 💡 Depuis PG 12, les CTE non récursives sont **inlinées** par le
> planificateur quand c'est rentable : plus de barrière d'optimisation
> systématique. Ajoutez `MATERIALIZED` pour forcer l'évaluation unique.

---

## 31. Fonctions de fenêtrage : le top-N, les rangs, les écarts

```sql
-- Dernière mesure par capteur (alternative propre au DISTINCT ON)
SELECT DISTINCT ON (capteur_id) capteur_id, valeur, horodatage
FROM mesures ORDER BY capteur_id, horodatage DESC;

-- Rang des onduleurs par puissance et par site
SELECT reference, site_id, puissance_kva,
       rank() OVER (PARTITION BY site_id ORDER BY puissance_kva DESC) AS rang_site
FROM onduleurs;

-- Écart entre deux mesures successives (dérive d'un capteur)
SELECT capteur_id, horodatage, valeur,
       valeur - lag(valeur) OVER (PARTITION BY capteur_id ORDER BY horodatage)
         AS delta
FROM mesures;

-- Moyenne glissante sur 1 heure (utile : lissage de courbes)
SELECT horodatage, valeur,
       avg(valeur) OVER (
         ORDER BY horodatage
         RANGE BETWEEN INTERVAL '1 hour' PRECEDING AND CURRENT ROW
       ) AS moyenne_1h
FROM mesures WHERE capteur_id = 42;
```

---

## 32. JSONB : requêtes et index

```sql
-- Extraire
SELECT config->>'seuil_alerte' AS seuil,          -- texte
       (config->>'seuil_alerte')::numeric AS seuil_num,
       config #> '{reseau,ip}' AS ip             -- chemin imbriqué
FROM capteurs;

-- Filtrer
SELECT * FROM capteurs WHERE config @> '{"actif": true}';      -- contient
SELECT * FROM capteurs WHERE config ? 'seuil_alerte';          -- clé existe
SELECT * FROM capteurs WHERE config->>'site' = 'SIEGE';

-- Modifier (sans réécrire tout le document à la main)
UPDATE capteurs
SET config = jsonb_set(config, '{seuil_alerte}', '75')
WHERE code = 'SONDE-01';

-- Index GIN pour les recherches dans le JSON
CREATE INDEX idx_capteurs_config ON capteurs USING gin (config);
```

---

## 33. Types réseau inet/cidr : requêtes IP natives

```sql
CREATE TABLE equipements (
  id  bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  nom text NOT NULL,
  ip  inet NOT NULL
);
CREATE INDEX idx_equipements_ip ON equipements USING gist (ip inet_ops);

-- Tous les équipements du /24
SELECT * FROM equipements WHERE ip << inet '10.0.1.0/24';

-- Le réseau contenant une IP
SELECT * FROM equipements WHERE inet '10.0.1.0/24' >> ip;

-- Trier par adresse IP (ordre numérique, pas lexical !)
SELECT nom, ip FROM equipements ORDER BY ip;
```

---

## 34. Dates et heures : timestamptz sans piège

```sql
-- Toujours timestamptz : stocké en UTC
SHOW timezone;                 -- ex. 'Europe/Paris'
SET timezone = 'UTC';

SELECT now(),                    -- timestamptz : instant absolu
       CURRENT_DATE,             -- date selon TimeZone
       now() AT TIME ZONE 'Europe/Paris' AS heure_paris;

-- Plages et intervalles
SELECT * FROM interventions
WHERE debut >= now() - INTERVAL '7 days';

-- Tronquer à l'heure / au jour (agrégations de supervision)
SELECT date_trunc('hour', horodatage) AS heure,
       avg(valeur) AS moyenne
FROM mesures
GROUP BY 1 ORDER BY 1;

-- Générer une série temporelle (boucher les trous d'un graphe)
SELECT g.h AS heure, count(m.id) AS nb_mesures
FROM generate_series(
       date_trunc('day', now() - INTERVAL '1 day'),
       date_trunc('day', now()),
       INTERVAL '1 hour'
     ) g(h)
LEFT JOIN mesures m
  ON date_trunc('hour', m.horodatage) = g.h
GROUP BY 1 ORDER BY 1;
```

---

## 35. Import/export CSV : COPY et \copy

```bash
# Export côté serveur (superuser, fichier SUR le serveur)
sudo -u postgres psql -d inventaire -c \
  "COPY (SELECT * FROM mesures WHERE horodatage > now() - INTERVAL '1 day')
   TO '/var/lib/postgresql/export/mesures.csv' WITH (FORMAT csv, HEADER);"

