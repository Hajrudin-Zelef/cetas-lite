---
id: collect-261001-rattrapage/rattrapage/grafana-guide-11
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["attribution"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [2162, 2400]
sha256: 4015446817656204df98fcb42395dc742a29fe899592029928d93aca9d1872da
---

# Guide Grafana — Dashboards, visualisation et alerting

| Capacité | Viewer | Editor | Admin (orga) | Server Admin |
|---|---|---|---|---|
| Voir les dashboards | ✅ | ✅ | ✅ | ✅ |
| Utiliser Explore | ✅ | ✅ | ✅ | ✅ |
| Créer/modifier ses dashboards | ❌ | ✅ | ✅ | ✅ |
| Créer des datasources | ❌ | ❌ | ✅ | ✅ |
| Gérer équipes/dossiers de l'orga | ❌ | ❌ | ✅ | ✅ |
| Gérer utilisateurs, orgas, plugins | ❌ | ❌ | ❌ | ✅ |
| Voir/modifier l'alerting* | lecture | selon conf | ✅ | ✅ |

\* L'alerting a ses propres permissions fines (section 56).

Règles d'attribution :

- **Viewer** par défaut pour tout nouvel utilisateur.
- **Editor** : aux exploitants qui construisent des dashboards.
- **Admin** : 1-2 personnes par orga (vous, votre adjoint).
- **Server Admin** : vous seul (et un suppléant documenté dans le coffre).

> **Le compte `admin` générique** : changez son mot de passe, stockez-le dans
> le coffre, et **ne l'utilisez pas au quotidien**. Créez-vous un compte
> nominatif Server Admin.

---

## 56. RBAC fin (permissions granulaires)

Depuis Grafana 9, le **RBAC** permet des permissions plus fines que les
rôles de base, via des **rôles personnalisés**.

Cas d'usage typiques :

- **Éditeurs d'alertes sans droit datasource** : rôle avec
  `alert.rules:read/write`, `alert.notifications:read/write` mais sans
  `datasources:write`.
- **Lecteurs d'un seul dossier** : Viewer global + permission `dashboards:read`
  limitée au dossier « Metier ».
- **Gestionnaires d'équipes** : `teams:write` sans être Admin de l'orga.

Création (Administration → Users and access → Roles), exemple en API :

```bash
curl -s -X POST -H "Content-Type: application/json" \
  -d '{
    "name": "editeur-alertes",
    "description": "Gère les règles et notifications, sans toucher aux datasources",
    "permissions": [
      {"action": "alert.rules:read", "scope": "folders:*"},
      {"action": "alert.rules:write", "scope": "folders:*"},
      {"action": "alert.notifications:read", "scope": "*"},
      {"action": "alert.notifications:write", "scope": "*"}
    ]
  }' \
  http://admin:<A_COMPLETER>@localhost:3000/api/access-control/roles | python3 -m json.tool
```

> **Principe du moindre privilège** : commencez restrictif, élargissez sur
> demande justifiée. C'est plus facile que l'inverse.

---

## 57. Authentification : basique, LDAP, OAuth (aperçu)

**Basique (base interne)** : suffisant pour une petite équipe. Activez le
renforcement :

```ini
[auth.basic]
enabled = true

[security]
# Verrouiller après échecs répétés (anti brute-force)
disable_brute_force_login_protection = false
```

**LDAP / Active Directory** (entreprise avec annuaire) :

```ini
[auth.ldap]
enabled = true
config_file = /etc/grafana/ldap.toml
allow_sign_up = true
```

```toml
# /etc/grafana/ldap.toml (extrait)
[[servers]]
host = "ad.mondomaine.fr"
port = 636
use_ssl = true
bind_dn = "CN=grafana_svc,OU=Comptes de service,DC=mondomaine,DC=fr"
bind_password = '<A_COMPLETER>'
search_base_dns = ["OU=Utilisateurs,DC=mondomaine,DC=fr"]

[servers.group_mappings]
group_dn = "CN=G_Grafana_Admins,OU=Groupes,DC=mondomaine,DC=fr"
org_role = "Admin"
[servers.group_mappings]
group_dn = "CN=G_Grafana_Editeurs,OU=Groupes,DC=mondomaine,DC=fr"
org_role = "Editor"
# Tout le reste tombe en Viewer via :
# [[servers]] ... group_dn = "*" → org_role = "Viewer"
```

**OAuth générique** (Keycloak, Azure AD, Google…) : via `[auth.generic_oauth]`
(client_id, client_secret, auth_url, token_url, api_url). Le `root_url`
(section 10) doit être exact, sinon la redirection échoue.

---

## 58. Service accounts et tokens d'API

Les **service accounts** remplacent les anciennes API keys (dépréciées) pour
l'automatisation : provisioning, CI, scripts d'export.

Création : **Administration → Users and access → Service accounts → Add**.

```bash
# Créer un token pour le compte de service (id 2, rôle Viewer)
curl -s -X POST -H "Content-Type: application/json" \
  -d '{"name":"token-ci","secondsToLive":0}' \
  http://admin:<A_COMPLETER>@localhost:3000/api/serviceaccounts/2/tokens \
  | python3 -m json.tool
# → {"key": "glsa_<A_COMPLETER>..."}  # affiché UNE SEULE FOIS
```

Usage :

```bash
# Lister les dashboards via le token
curl -s -H "Authorization: Bearer glsa_<A_COMPLETER>" \
  http://localhost:3000/api/search?query=supervision | python3 -m json.tool

# Exporter un dashboard en JSON (pour Git)
curl -s -H "Authorization: Bearer glsa_<A_COMPLETER>" \
  http://localhost:3000/api/dashboards/uid/<UID> | python3 -m json.tool > dashboard.json
```

Bonnes pratiques :

- Un service account **par usage** (CI, export, script de health check), avec
  le rôle minimal.
- `secondsToLive` : mettez une **expiration** (ex. 90 jours) sauf besoin
  permanent documenté.
- Stockez les tokens dans le coffre, jamais dans les scripts en clair.
- Révoquez les tokens des comptes de service inutilisés (revue trimestrielle).

---

## 59. HTTPS en production via reverse proxy Nginx

Architecture recommandée : Nginx termine TLS, Grafana écoute en HTTP sur
`127.0.0.1:3000`.

```ini
# /etc/grafana/grafana.ini
[server]
protocol = http
http_addr = 127.0.0.1
http_port = 3000
root_url = https://grafana.mondomaine.fr/
```

```nginx
# /etc/nginx/sites-available/grafana
upstream grafana {
    server 127.0.0.1:3000;
}

server {
    listen 443 ssl http2;
    server_name grafana.mondomaine.fr;

    ssl_certificate     /etc/letsencrypt/live/grafana.mondomaine.fr/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/grafana.mondomaine.fr/privkey.pem;

    # En-têtes de sécurité
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
    add_header X-Content-Type-Options nosniff always;
    add_header X-Frame-Options SAMEORIGIN always;
    add_header Referrer-Policy no-referrer always;

    # WebSocket (nécessaire pour les live dashboards)
    location / {
        proxy_pass http://grafana;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }

    # API : timeouts un peu plus généreux pour les grosses requêtes
    location /api/ds/query {
        proxy_pass http://grafana;
        proxy_read_timeout 120s;
        proxy_set_header Host $host;
    }
}

server {
    listen 80;
    server_name grafana.mondomaine.fr;
    return 301 https://$host$request_uri;
}
```

```bash
sudo nginx -t && sudo systemctl reload nginx
# Certificat Let's Encrypt :
sudo certbot --nginx -d grafana.mondomaine.fr
```

> Si Grafana est servi sous un **sous-chemin** (`https://intranet/grafana/`),
> renseignez `root_url = https://intranet/grafana/` **et**
> `serve_from_sub_path = true`. Mais préférez un sous-domaine dédié : moins
> de surprises (cookies, OAuth, websockets).

---

## 60. Durcissement sécurité — checklist

- [ ] `allow_sign_up = false`, pas d'anonyme sauf kiosque dédié (section 72).
- [ ] `admin_password` fort, stocké au coffre ; compte `admin` non utilisé
      au quotidien.
- [ ] `secret_key` générée aléatoirement, sauvegardée avec la base.
- [ ] HTTPS partout (Nginx + HSTS) ; `cookie_secure = true`.
- [ ] Grafana **non exposé directement** sur Internet sans nécessité ;
      derrière VPN/SSO si possible.
- [ ] Comptes nominatifs ; revue trimestrielle des comptes et tokens.
- [ ] Datasources en mode `Server`, comptes lecture seule côté backends.
- [ ] `hide_version = true` (hygiène).
- [ ] Pare-feu : n'ouvrir que 443 (et 22 admin) ; 3000 en localhost uniquement.
- [ ] Mises à jour suivies (section 66) ; alertes CVE Grafana surveillées.
- [ ] Logs d'audit : qui a modifié quoi (table `annotation` / logs en `info`).
- [ ] Sauvegardes chiffrées, testées (sections 62-65).

