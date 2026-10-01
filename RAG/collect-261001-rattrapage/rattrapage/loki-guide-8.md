---
id: collect-261001-rattrapage/rattrapage/loki-guide-8
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "incident"]
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [1821, 2110]
sha256: 8ac633a0bffccefc49419dbef0ba2b91ffef2b21b35d676a307458180aa85501
---

# Grafana Loki — Le guide complet

1. **Query type** : `Range` (graphique + logs) ou `Instant` (table).
2. **Label browser** : cliquez sur `{job="..."}` pour découvrir les
   labels et valeurs disponibles sans connaître LogQL.
3. **Filtres ad hoc** : `|=` / `|~` / parsers s'ajoutent en un clic sur
   une ligne de log (« Filter for » / « Filter out »).
4. **Live tailing** : le bouton ▶ suit les logs en temps réel.
5. **Plage** : commencez par 1 h, élargissez si besoin.

Raccourcis qui changent la vie :

| Action | Comment |
|---|---|
| Voir le contexte d'une ligne | « Show context » (± 20 lignes autour) |
| Exclure un motif bruyant | clic droit → « Filter out » (`!=`) |
| Passer en plein écran | `e` puis `f` |
| Copier la requête | menu du panneau |

> 💡 En incident, la séquence gagnante : sélecteur resserré (`{job, host,
> level}`) → plage 1 h → `|=` sur le symptôme → « Show context » sur les
> lignes intéressantes.

---

## 48. Dashboards logs : exemples de panneaux

Panneau 1 — **Taux d'erreurs 5xx nginx** (Time series) :

```logql
sum(rate({job="nginx"} | regexp `(?P<status>\S+)` | status=~"5.." [5m]))
```

Panneau 2 — **Logs bruts filtrés** (Logs panel) :

```logql
{job="nginx", level="error"} |~ " 5\d\d "
```

Panneau 3 — **Top 10 des URLs en erreur** (Table) :

```logql
topk(10, sum by (path) (count_over_time(
  {job="nginx"} | regexp `"[A-Z]+ (?P<path>\S+)" \d{3}` | status=~"5.." [1h]
)))
```

Panneau 4 — **Volume d'ingestion par job** (Time series) :

```logql
sum by (job) (bytes_rate({job=~".+"}[5m]))
```

Panneau 5 — **Carte de chaleur des erreurs par host** (Heatmap) :

```logql
sum by (host) (count_over_time({job="syslog"} |= "error" [1m]))
```

---

## 49. Variables de dashboard LogQL

Rendez les dashboards réutilisables avec des variables :

| Variable | Type | Requête |
|---|---|---|
| `$job` | Query | `label_values({job=~".+"}, job)` |
| `$host` | Query | `label_values({job="$job"}, host)` |
| `$level` | Custom | `debug,info,warn,error` |
| `$intervalle` | Interval | `1m,5m,15m,1h` |

Utilisation dans les panneaux :

```logql
sum(rate({job="$job", host=~"$host", level=~"$level"}[$intervalle]))
```

> ⚠️ `label_values()` sur un gros périmètre peut être lent : restreignez
> avec un sélecteur (`{job="$job"}`) et activez le **refresh on dashboard
> load** plutôt que « on time range change ».

---

## 50. Cas pratique 1 : centraliser le syslog de 20 serveurs

**Objectif** : tous les `/var/log/syslog` (+ journald) de 20 serveurs
Debian dans un Loki central, avec le hostname en label.

### Côté Loki central (`loki.local`)

`loki.yaml` monolithique de la section 13 + pare-feu :

```bash
# N'autoriser que le réseau des serveurs sur le port 3100
sudo ufw allow from 10.0.0.0/24 to any port 3100 proto tcp
```

### Côté chaque serveur : Promtail

`/etc/promtail/promtail.yaml` (identique partout sauf `host`) :

```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /var/lib/promtail/positions.yaml

clients:
  - url: http://loki.local:3100/loki/api/v1/push
    batchwait: 1s
    batchsize: 1048576

scrape_configs:
  - job_name: syslog
    static_configs:
      - targets: [localhost]
        labels:
          job: syslog
          host: web01          # ← à personnaliser par serveur
          env: prod
          __path__: /var/log/syslog
  - job_name: journal
    journal:
      path: /var/log/journal
      max_age: 12h
      labels:
        job: systemd-journal
        host: web01
        env: prod
    relabel_configs:
      - source_labels: ['__journal__systemd_unit']
        target_label: unit
```

### Déploiement automatisé (Ansible, extrait)

```yaml
- name: Déployer Promtail
  hosts: all
  become: true
  vars:
    loki_url: http://loki.local:3100/loki/api/v1/push
  tasks:
    - name: Copier le binaire
      ansible.builtin.copy:
        src: files/promtail-linux-amd64
        dest: /usr/local/bin/promtail
        mode: '0755'
    - name: Configurer (host dynamique)
      ansible.builtin.template:
        src: templates/promtail.yaml.j2   # {{ inventory_hostname }} pour host
        dest: /etc/promtail/promtail.yaml
      notify: Redémarrer promtail
  handlers:
    - name: Redémarrer promtail
      ansible.builtin.systemd:
        name: promtail
        state: restarted
        enabled: true
```

### Vérification

```logql
# Les 20 hosts envoient-ils tous ?
count(count_over_time({job="syslog"}[5m]) by (host))
# → doit retourner 20
```

---

## 51. Cas pratique 2 : suivre les erreurs nginx

**Objectif** : dashboard + alerte sur les erreurs HTTP 5xx de 3 reverse
proxies nginx.

### Promtail (sur chaque proxy)

```yaml
scrape_configs:
  - job_name: nginx
    static_configs:
      - targets: [localhost]
        labels:
          job: nginx
          host: proxy01
          env: prod
          __path__: /var/log/nginx/access.log
    pipeline_stages:
      - regex:
          expression: '^(?P<ip>\S+) \S+ \S+ \[(?P<time>[^\]]+)\] "(?P<method>\S+) (?P<path>\S+) \S+" (?P<status>\d{3}) (?P<size>\d+|-) "[^"]*" "(?P<ua>[^"]*)"'
      - labels:
          method:
          status:
      - timestamp:
          source: time
          format: "02/Jan/2006:15:04:05 -0700"
```

Format nginx correspondant (`log_format` dans `nginx.conf`) :

```nginx
log_format main '$remote_addr - $remote_user [$time_local] "$request" '
                '$status $body_bytes_sent "$http_referer" '
                '"$http_user_agent"';
access_log /var/log/nginx/access.log main;
```

### Requêtes du dashboard

```logql
# Taux de 5xx par proxy
sum by (host) (rate({job="nginx", status=~"5.."}[5m]))

# Top URLs en erreur
topk(10, sum by (path) (count_over_time({job="nginx", status=~"5.."}[1h])))

# Taux d'erreur global en %
sum(rate({job="nginx", status=~"5.."}[5m]))
/ sum(rate({job="nginx"}[5m])) * 100
```

### Alerte (ruler)

```yaml
      - alert: NginxTropDErreurs500
        expr: sum(rate({job="nginx", status=~"5.."}[5m])) > 0.5
        for: 5m
        labels:
          severity: critical
          service: nginx
        annotations:
          summary: "Plus de 0,5 erreur 500/s sur nginx"
```

---

## 52. Cas pratique 3 : détecter les attaques SSH (avec fail2ban)

**Objectif** : repérer les scans/brute-force SSH dans Loki, et faire le
lien avec **fail2ban** qui bloque déjà les attaquants.

### Ce que fail2ban fait (rappel)

fail2ban lit `/var/log/auth.log`, détecte les `Failed password` répétés
et bannit l'IP via iptables/nftables. Loki ne le remplace pas : il apporte
la **vision globale** (quels serveurs sont attaqués ? quelles IP reviennent ?).

### Promtail : parser auth.log

```yaml
scrape_configs:
  - job_name: auth
    static_configs:
      - targets: [localhost]
        labels:
          job: sshd
          host: srv01
          __path__: /var/log/auth.log
    pipeline_stages:
      - regex:
          expression: 'Failed password for (invalid user )?(?P<user>\S+) from (?P<ip>\S+) port \d+'
      - labels:
          user:
      - structured_metadata:
          ip:
```

> ⚠️ `ip` en structured metadata (pas en label !) : la cardinalité des IP
> attaquantes est trop forte pour un label.

### Requêtes de détection

```logql
# Tentatives par minute (tous serveurs)
sum(count_over_time({job="sshd"} |= "Failed password" [1m]))

# Top IP attaquantes sur 24h
topk(20, sum by (ip) (count_over_time({job="sshd"} |= "Failed password" [24h])))

# Quels serveurs sont visés ?
sum by (host) (count_over_time({job="sshd"} |= "Failed password" [1h]))

# Utilisateurs ciblés (root ?)
topk(10, sum by (user) (count_over_time({job="sshd"} |= "Failed password" [24h])))
```

### Alerte brute-force

