---
id: collect-261001-rattrapage/rattrapage/glpi-guide-15
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "arr", "incident", "memory"]
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [2439, 2653]
sha256: b3efd939cad3fd6e52767b94a8df1be649820306eb0291a640c6d24c849c0f5d
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

- [ ] Installation LAMP/nginx + GLPI 10.x OK, assistant terminé.
- [ ] Comptes par défaut sécurisés, dossier `install/` supprimé.
- [ ] URL d'application en HTTPS renseignée.
- [ ] Cron système actif et vérifié (actions automatiques à jour).
- [ ] SMTP testé, modèles de notification relus.
- [ ] Entités, profils, groupes créés ; utilisateurs importés (LDAP).
- [ ] Catégories de tickets + gabarits (champs obligatoires).
- [ ] SLA + calendriers + niveaux d'escalade configurés.
- [ ] Fiches parc importées (échantillon vérifié), contrats liés.
- [ ] Stock cartouches/consommables initialisé (comptage physique).
- [ ] Tickets récurrents (relevés, préventif) programmés.
- [ ] Sauvegarde automatisée + **test de restauration** effectué.
- [ ] Supervision (sondes section 46) en place.
- [ ] Formation techniciens (2 h) et déclarants (1 h) réalisée.
- [ ] Pense-bête affiché à l'atelier.
- [ ] Date d'arrêt de l'ancien système fixée et communiquée.

---

# ANNEXES

## 71. Annexe A — paramètres PHP recommandés

`/etc/php/8.2/apache2/php.ini` (et `cli/php.ini` pour le cron) :

```ini
[PHP]
memory_limit = 256M
max_execution_time = 120
max_input_time = 120
max_input_vars = 5000
upload_max_filesize = 32M
post_max_size = 40M
date.timezone = Europe/Paris
display_errors = Off
log_errors = On
error_log = /var/log/php_errors.log
session.cookie_httponly = 1
session.cookie_secure = 1
session.cookie_samesite = "Lax"
session.gc_maxlifetime = 7200

[opcache]
opcache.enable = 1
opcache.memory_consumption = 128
opcache.interned_strings_buffer = 16
opcache.max_accelerated_files = 10000
opcache.validate_timestamps = 1
opcache.revalidate_freq = 60

[APCu]
apc.enabled = 1
apc.shm_size = 64M
```

---

## 72. Annexe B — configuration Apache de référence

`/etc/apache2/sites-available/glpi.conf` :

```apache
<VirtualHost *:80>
    ServerName glpi.entreprise.lan
    Redirect permanent / https://glpi.entreprise.lan/
</VirtualHost>

<VirtualHost *:443>
    ServerName glpi.entreprise.lan
    DocumentRoot /var/www/glpi/public

    SSLEngine on
    SSLCertificateFile    /etc/ssl/certs/glpi.crt
    SSLCertificateKeyFile /etc/ssl/private/glpi.key

    SetEnv GLPI_CONFIG_DIR /var/lib/glpi/config
    SetEnv GLPI_VAR_DIR /var/lib/glpi/files

    <Directory /var/www/glpi/public>
        Require all granted
        AllowOverride All
        FallbackResource /index.php
    </Directory>

    <Directory /var/www/glpi>
        Require all denied
    </Directory>
    <Directory /var/www/glpi/public>
        Require all granted
    </Directory>

    Header always set Strict-Transport-Security "max-age=31536000; includeSubDomains"
    Header always set X-Content-Type-Options "nosniff"
    Header always set X-Frame-Options "SAMEORIGIN"
    Header always set Referrer-Policy "strict-origin-when-cross-origin"

    ErrorLog  ${APACHE_LOG_DIR}/glpi_error.log
    CustomLog ${APACHE_LOG_DIR}/glpi_access.log combined
</VirtualHost>
```

---

## 73. Annexe C — configuration nginx de référence

`/etc/nginx/sites-available/glpi` (version complète) :

```nginx
server {
    listen 80;
    server_name glpi.entreprise.lan;
    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl;
    server_name glpi.entreprise.lan;
    root /var/www/glpi/public;
    index index.php;

    ssl_certificate     /etc/ssl/certs/glpi.crt;
    ssl_certificate_key /etc/ssl/private/glpi.key;
    ssl_protocols TLSv1.2 TLSv1.3;

    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-Frame-Options "SAMEORIGIN" always;

    client_max_body_size 40M;

    location / {
        try_files $uri /index.php$is_args$args;
    }

    location ~ ^/index\.php(/|$) {
        fastcgi_pass unix:/run/php/glpi.sock;
        fastcgi_split_path_info ^(.+\.php)(/.*)$;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param PATH_INFO $fastcgi_path_info;
        fastcgi_param GLPI_CONFIG_DIR /var/lib/glpi/config;
        fastcgi_param GLPI_VAR_DIR /var/lib/glpi/files;
    }

    location ~ ^/(config|files|inc|install|tests|marketplace)/ { deny all; }
    location ~ /\. { deny all; }
}
```

---

## 74. Annexe D — modèle de gabarit « Intervention copieur »

À reproduire dans `Configuration > Gabarits > Gabarits de tickets` :

```
Gabarit : INTERVENTION COPIEUR
Catégorie par défaut : Panne matérielle > Code erreur affiché (modifiable)
Type : Incident (verrouillé)

Champs :
  [OBLIGATOIRE] Éléments associés  → aide : « Sélectionnez LE copieur concerné »
  [OBLIGATOIRE] Urgence            → aide : « Très haute = client totalement bloqué »
  [OBLIGATOIRE] Compteur N&B       → champ personnalisé (nombre)
  [OBLIGATOIRE] Compteur couleur   → champ personnalisé (nombre)
  [FACULTATIF ] Code erreur affiché → champ personnalisé (texte, ex. C-0202, J-0511)
  [OBLIGATOIRE] Titre              → aide : « Ex. : Bourrage récurrent bac 2 »
  [OBLIGATOIRE] Description        → aide : « Symptômes, fréquence, depuis quand,
                                      ce qui a déjà été tenté »
  [MASQUÉ     ] Impact             → fixé par la hotline à la qualification
  [PRÉ-REMPLI ] Source             → « Téléphone » / « Collecteur mail »

Notifications : déclarant + technicien assigné + groupe assigné
SLA suggéré   : selon l'entité (règle métier)
```

---

## 75. Annexe E — modèle de rapport hebdomadaire SAV

```markdown
# Rapport SAV — semaine 39 (22 → 26 sept. 2026)

## 1. Activité
- Tickets créés : 47 | résolus : 44 | en cours : 18 (dont 3 > 15 j)
- Répartition : pannes 28, consommables 9, demandes 7, préventif 3

## 2. SLA
- TTO tenu : 96 % (objectif 95 %) | TTR tenu : 91 % (objectif 90 %)
- Dépassements : #4821 (attente pièce), #4833 (client injoignable) → actions

## 3. Qualité
- Satisfaction : 4,6/5 (12 réponses) — 1 note 2/5 (#4815, rappel technicien fait)
- Tickets rouverts : 2

## 4. Parc & préventif
- Compteurs relevés : 118/120 (2 relances)
- Visites préventives réalisées : 3/3

## 5. Stock & contrats
- Alertes : toner TK-8345K sous seuil (commande passée)
- Contrats à J-90 : Clinique Saint-Roch (renégociation en cours)

## 6. Budget
- Pièces : 3 240 € (cumul 41 k€ / 60 k€) | Consommables : 1 120 €

## 7. Actions
- [ ] Relancer fournisseur pièces (M. Diarra, avant vendredi)
- [ ] Former N1 au diagnostic J-0511 (article KB n°42)
```

---

*Fin du guide — GLPI 10.x. Document vivant : revoyez-le à chaque mise à jour
majeure et après chaque revue annuelle des processus SAV.*
