---
id: collect-261001-rattrapage/rattrapage/ansible-guide-5
title: "Guide Ansible — Automatisation système en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "distribution"]
source: docs/RAG/collect-261001-rattrapage/ansible_guide.md
source_anchor: ""
source_lines: [1007, 1244]
sha256: 8e2371df8c3214e971dd365c397c0f40b3b2dc9f6517f386bf7f8365cf0ef1ab
---

# Guide Ansible — Automatisation système en production

Du **moins** prioritaire au **plus** prioritaire (mémorisez les 5 qui comptent
vraiment au quotidien : ils sont marqués ★) :

| # | Niveau | Exemple |
|---|---|---|
| 1 | Arguments de ligne de commande `--extra-vars` ★ | `-e "http_port=8080"` |
| 2 | `set_facts` / variables enregistrées | `register: resultat` |
| 3 | Variables du play (`vars:`) | bloc `vars` du play |
| 4 | Variables de tâche (`vars:` sur la tâche) | bloc `vars` de la tâche |
| 5 | Variables de bloc | `vars` sur un `block` |
| 6 | Variables de rôle (`vars/main.yml`) | `roles/x/vars/main.yml` |
| 7 | Variables d'hôte (`host_vars/`) ★ | `host_vars/db-01.yml` |
| 8 | Variables de groupe (`group_vars/`) ★ | `group_vars/webservers.yml` |
| 9 | Facts de l'hôte (`ansible_facts`) | `ansible_facts['os_family']` |
| 10 | Inventaire : `host_vars` inline | `web-01 ansible_host=...` |
| 11 | Inventaire : `group_vars` inline | `[groupe:vars]` |
| 12 | `defaults/main.yml` du rôle ★ | valeurs par défaut surchargeables |
| 13 | Paramètres de connexion | `ansible_user`, `ansible_port`… |

> Les niveaux 14 à 21 concernent des cas avancés (vars plugins, `include_vars`
> dynamiques, variables de `import_role`…). En pratique : **defaults < groupe <
> hôte < extra-vars**. Le reste se consulte dans `ansible-doc` au besoin.

Règle pratique : mettez les valeurs par défaut dans `defaults/`, les valeurs
d'environnement dans `group_vars/`, les exceptions dans `host_vars/`, et les
surcharges ponctuelles en `-e`.

## 35. Facts Ansible : connaître le système cible

Les facts sont collectés automatiquement (`gather_facts: true`) au début du play.
Exemples de facts indispensables :

```yaml
- name: Utiliser les facts pour adapter la configuration
  ansible.builtin.template:
    src: app.conf.j2
    dest: /etc/monapp/app.conf
  vars:
    memoire_totale: "{{ ansible_facts['memtotal_mb'] }}"
    nb_cpu: "{{ ansible_facts['processor_vcpus'] }}"
```

| Fact | Contenu |
|---|---|
| `ansible_facts['distribution']` | `Debian`, `Ubuntu`… |
| `ansible_facts['distribution_major_version']` | `12`, `22`… |
| `ansible_facts['os_family']` | `Debian`, `RedHat`… |
| `ansible_facts['default_ipv4']['address']` | IP principale |
| `ansible_facts['memtotal_mb']` | RAM totale en Mo |
| `ansible_facts['processor_vcpus']` | Nombre de vCPU |
| `ansible_facts['mounts']` | Points de montage |
| `ansible_facts['hostname']` | Nom court de la machine |

Désactiver la collecte pour gagner du temps quand elle est inutile :

```yaml
- hosts: all
  gather_facts: false
  tasks:
    - name: Redémarrage simple sans facts
      ansible.builtin.reboot:
```

## 36. Variables magiques : inventory_hostname, groups, hostvars

```yaml
- name: Générer la liste de tous les serveurs web
  ansible.builtin.template:
    src: upstreams.conf.j2
    dest: /etc/nginx/upstreams.conf
```

```jinja2
{# templates/upstreams.conf.j2 — utilise les variables magiques #}
upstream backend {
{% for hote in groups['webservers'] %}
    server {{ hostvars[hote]['ansible_host'] }}:8080;
{% endfor %}
}
```

| Variable magique | Contenu |
|---|---|
| `inventory_hostname` | Nom de l'hôte courant dans l'inventaire |
| `inventory_hostname_short` | Nom court (avant le premier point) |
| `groups` | Dictionnaire groupe → liste d'hôtes |
| `group_names` | Groupes auxquels appartient l'hôte courant |
| `hostvars` | Variables de **tous** les hôtes |
| `play_hosts` / `ansible_play_hosts` | Hôtes du play en cours |
| `ansible_check_mode` | `true` si `--check` actif |

## 37. Handlers : réagir aux changements, une seule fois

```yaml
tasks:
  - name: Déployer nginx.conf
    ansible.builtin.template:
      src: nginx.conf.j2
      dest: /etc/nginx/nginx.conf
      validate: "nginx -t -c %s"
    notify: Recharger nginx

  - name: Déployer le vhost
    ansible.builtin.template:
      src: vhost.conf.j2
      dest: /etc/nginx/sites-enabled/app.conf
      validate: "nginx -t -c %s"
    notify: Recharger nginx

handlers:
  - name: Recharger nginx
    ansible.builtin.service:
      name: nginx
      state: reloaded
```

Même si les deux tâches changent quelque chose, **le handler ne s'exécute qu'une fois**
en fin de play. Règles :

- Un handler se déclenche uniquement si la tâche notificatrice est `changed`.
- `notify` peut appeler plusieurs handlers (liste) ou utiliser `listen` pour regrouper.
- Forcer l'exécution immédiate : `meta: flush_handlers`.

```yaml
- name: Forcer les handlers notifiés jusqu'ici
  ansible.builtin.meta: flush_handlers
```

## 38. Jinja2 : les bases du templating

Jinja2 est le moteur de templates d'Ansible. Trois délimiteurs :

| Délimiteur | Usage |
|---|---|
| `{{ ... }}` | Afficher une variable / expression |
| `{% ... %}` | Instruction (boucle, condition) |
| `{# ... #}` | Commentaire (n'apparaît pas dans le rendu) |

```jinja2
# {{ ansible_managed }} — en-tête standard : fichier géré par Ansible, ne pas éditer
serveur={{ inventory_hostname }}
port={{ http_port }}
tls={{ tls_active | ternary('on', 'off') }}
```

Bonnes pratiques :

- Quotez toujours les expressions qui commencent par `{{` en YAML :
  `port: "{{ http_port }}"` (sinon erreur de parsing YAML).
- `{{ ansible_managed }}` en en-tête de chaque template : traçabilité.
- Testez un template avec `--check --diff` avant application.

## 39. Jinja2 : boucles et conditions

```jinja2
{# Boucle simple #}
{% for dns in serveurs_dns %}
nameserver {{ dns }}
{% endfor %}

{# Boucle avec index et condition #}
{% for hote in groups['dbservers'] %}
{% if hostvars[hote]['ansible_host'] is defined %}
# {{ hote }} ({{ loop.index }}/{{ loop.length }})
db_host={{ hostvars[hote]['ansible_host'] }}
{% endif %}
{% endfor %}

{# Conditions multiples #}
{% if tls_active and cert_path is defined %}
ssl_certificate {{ cert_path }};
{% elif tls_active %}
# ATTENTION : TLS demandé mais aucun certificat défini
{% else %}
# TLS désactivé
{% endif %}
```

## 40. Jinja2 : les filtres les plus utiles

| Filtre | Exemple | Résultat |
|---|---|---|
| `default` | `{{ var \| default('vide') }}` | Valeur ou `vide` |
| `ternary` | `{{ ok \| ternary('oui','non') }}` | `oui` / `non` |
| `upper` / `lower` | `{{ nom \| upper }}` | `MONAPP` |
| `join` | `{{ liste \| join(',') }}` | `a,b,c` |
| `split` | `{{ 'a:b' \| split(':') }}` | `['a','b']` |
| `length` | `{{ liste \| length }}` | `3` |
| `sort` / `unique` | `{{ liste \| unique \| sort }}` | triée, dédupliquée |
| `dict2items` | `{{ dico \| dict2items }}` | liste de `{key, value}` |
| `to_json` / `to_yaml` | `{{ var \| to_json }}` | sérialisation |
| `password_hash` | `{{ mdp \| password_hash('yescrypt') }}` | hash yescrypt |
| `ipaddr` | `{{ '192.168.1.5/24' \| ipaddr('network') }}` | `192.168.1.0/24` |
| `b64encode` / `b64decode` | `{{ secret \| b64encode }}` | base64 |
| `regex_replace` | `{{ v \| regex_replace('^v','') }}` | substitution |

Exemple combiné :

```jinja2
serveurs_autorises={{ ips_admin | unique | sort | join(',') }}
motd_hash={{ motd_admin | password_hash('yescrypt') }}
```

## 41. Template complet : vhost nginx avec Jinja2

```jinja2
# {{ ansible_managed }}
# Rôle nginx — vhost {{ nom_app }}

upstream {{ nom_app }}_backend {
{% for hote in groups['webservers'] %}
    server {{ hostvars[hote]['ansible_host'] }}:{{ port_backend }} max_fails=3 fail_timeout=30s;
{% endfor %}
}

server {
    listen {{ http_port }}{% if tls_active %} ssl{% endif %};
    server_name {{ nom_domaine }};

{% if tls_active %}
    ssl_certificate {{ cert_path }};
    ssl_certificate_key {{ cle_path }};
    ssl_protocols TLSv1.2 TLSv1.3;
{% endif %}

    access_log /var/log/nginx/{{ nom_app }}_access.log;
    error_log  /var/log/nginx/{{ nom_app }}_error.log;

    location / {
        proxy_pass http://{{ nom_app }}_backend;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

## 42. Rôles : l'arborescence standard

