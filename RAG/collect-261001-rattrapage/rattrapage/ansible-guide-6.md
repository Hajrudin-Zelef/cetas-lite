---
id: collect-261001-rattrapage/rattrapage/ansible-guide-6
title: "Guide Ansible — Automatisation système en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["aws", "packaging"]
source: docs/RAG/collect-261001-rattrapage/ansible_guide.md
source_anchor: ""
source_lines: [1245, 1542]
sha256: 1ba34859fb6b61e3f34f31b062dbec7520c3f844d425d5394503936cf6c0e7af
---

# Guide Ansible — Automatisation système en production

Un rôle regroupe tâches, handlers, templates, fichiers et variables par fonction.
C'est l'unité de réutilisation d'Ansible.

```text
roles/nginx/
├── tasks/
│   ├── main.yml          # point d'entrée : inclut les autres
│   ├── install.yml
│   └── configure.yml
├── handlers/
│   └── main.yml
├── templates/
│   ├── nginx.conf.j2
│   └── vhost.conf.j2
├── files/
│   └── dhparam.pem
├── vars/
│   └── main.yml          # variables du rôle (forte précédence)
├── defaults/
│   └── main.yml          # valeurs par défaut (faible précédence ★)
├── meta/
│   └── main.yml          # dépendances, plateformes supportées
└── README.md
```

Créer le squelette :

```bash
ansible-galaxy role init roles/nginx
```

`tasks/main.yml` typique :

```yaml
---
- name: Inclure l'installation
  ansible.builtin.include_tasks: install.yml

- name: Inclure la configuration
  ansible.builtin.include_tasks: configure.yml
```

## 43. Rôle complet : exemple `durcissement_ssh`

```yaml
# roles/durcissement_ssh/defaults/main.yml
---
ssh_port: 22
ssh_permit_root_login: "no"
ssh_password_authentication: "no"
ssh_max_auth_tries: 3
```

```yaml
# roles/durcissement_ssh/tasks/main.yml
---
- name: Déployer sshd_config durci
  ansible.builtin.template:
    src: sshd_config.j2
    dest: /etc/ssh/sshd_config
    owner: root
    group: root
    mode: "0600"
    backup: true
    validate: "sshd -t -f %s"
  notify: Redémarrer sshd

- name: S'assurer que sshd tourne
  ansible.builtin.service:
    name: ssh
    state: started
    enabled: true
```

```yaml
# roles/durcissement_ssh/handlers/main.yml
---
- name: Redémarrer sshd
  ansible.builtin.service:
    name: ssh
    state: restarted
```

```jinja2
{# roles/durcissement_ssh/templates/sshd_config.j2 #}
# {{ ansible_managed }}
Port {{ ssh_port }}
PermitRootLogin {{ ssh_permit_root_login }}
PasswordAuthentication {{ ssh_password_authentication }}
ChallengeResponseAuthentication no
MaxAuthTries {{ ssh_max_auth_tries }}
X11Forwarding no
AllowAgentForwarding no
```

Utilisation :

```yaml
- hosts: all
  become: true
  roles:
    - role: durcissement_ssh
      vars:
        ssh_port: 2222
```

## 44. Ansible Galaxy : partager et réutiliser des rôles

```bash
# Initialiser un rôle partageable
ansible-galaxy role init mon_role --init-path roles/

# Installer un rôle depuis Galaxy
ansible-galaxy role install geerlingguy.nginx

# Figer les versions dans requirements.yml (recommandé)
cat > roles/requirements.yml <<'EOF'
---
roles:
  - name: geerlingguy.nginx
    version: 3.2.0
    src: https://galaxy.ansible.com
EOF
ansible-galaxy role install -r roles/requirements.yml
```

> **Prudence production** : un rôle Galaxy est du code tiers exécuté en root.
> Épinglez la version, lisez le code (`roles/geerlingguy.nginx/tasks/`), et
> testez en staging avant la production.

## 45. Collections : le packaging moderne

Les collections regroupent modules, plugins, rôles et playbooks par domaine
(`community.general`, `amazon.aws`, `ansible.posix`…).

```yaml
# collections/requirements.yml
---
collections:
  - name: community.general
    version: ">=9.0.0"
  - name: amazon.aws
    version: ">=8.0.0"
  - name: ansible.posix
    version: ">=1.5.0"
```

```bash
ansible-galaxy collection install -r collections/requirements.yml
ansible-galaxy collection list
```

Dans les playbooks, utilisez les **FQCN** (noms pleinement qualifiés) :

```yaml
# Bien : explicite, pas d'ambiguïté
- ansible.posix.firewalld:
    service: https
    state: enabled

# Éviter : dépend de la résolution implicite
- firewalld:
    service: https
```

## 46. Ansible Vault : créer un coffre-fort

Vault chiffre les variables sensibles (mots de passe, clés API, certificats).
Le contenu chiffré reste versionnable en Git en toute sécurité.

```bash
# Créer un fichier chiffré
ansible-vault create group_vars/all/vault.yml

# Le fichier contient ensuite (éditable uniquement avec le mot de passe) :
# vault_pve_api_password: "SUPER_SECRET_A_CHANGER"
# vault_smtp_password: "AUTRE_SECRET"

# Chiffrer un fichier existant
ansible-vault encrypt host_vars/db-01/secrets.yml

# Déchiffrer (rare, à éviter : préférez edit)
ansible-vault decrypt /tmp/fichier.yml

# Changer le mot de passe du coffre
ansible-vault rekey group_vars/all/vault.yml
```

Fichier de mot de passe pour l'automatisation (CI/AWX) :

```bash
# ~/.vault_pass : chmod 600, hors Git !
echo "mot-de-passe-du-coffre" > ~/.vault_pass
chmod 600 ~/.vault_pass
# ansible.cfg :
# [defaults]
# vault_password_file = ~/.vault_pass
```

> **Sécurité** : le fichier `~/.vault_pass` a les droits `600`, n'est jamais
> commité, et idéalement est fourni par un gestionnaire de secrets (HashiCorp
> Vault, Bitwarden, variable d'environnement CI masquée).

## 47. Vault : chiffrer, éditer, utiliser au quotidien

```bash
# Éditer un coffre existant
ansible-vault edit group_vars/all/vault.yml

# Voir sans éditer
ansible-vault view group_vars/all/vault.yml

# Chiffrer une seule valeur (pratique pour une variable isolée)
ansible-vault encrypt_string 'mon_secret' --name 'api_token'
# Résultat à coller dans un fichier vars :
# api_token: !vault |
#   $ANSIBLE_VAULT;1.1;AES256
#   3362663138623839...

# Exécuter un playbook avec le coffre
ansible-playbook site.yml --ask-vault-pass
ansible-playbook site.yml --vault-password-file ~/.vault_pass
```

Utilisation dans un playbook :

```yaml
- name: Créer l'utilisateur de sauvegarde
  ansible.builtin.user:
    name: backup
    password: "{{ vault_backup_password_hash }}"  # hash yescrypt chiffré par Vault
```

## 48. Vault : bonnes pratiques et erreurs à éviter

- [ ] **Un coffre par environnement** : `vault_prod.yml` / `vault_staging.yml` avec des
      mots de passe différents (`--vault-id prod@~/.vault_pass_prod`).
- [ ] Ne chiffrez que les secrets, pas tout le fichier : le diff Git reste lisible.
- [ ] Nommez les variables `vault_*` pour repérer d'un coup d'œil ce qui est sensible.
- [ ] Ne commitez **jamais** `~/.vault_pass`, même par accident : ajoutez-le au `.gitignore`
      global et vérifiez avec `git log --all --full-history -- "*vault_pass*"`.
- [ ] Rotation : `ansible-vault rekey` à chaque départ d'un membre de l'équipe.
- [ ] En cas de fuite du mot de passe Vault, considérez **tous** les secrets compromis.

```bash
# Multi-vault-id : prod et staging séparés
ansible-playbook site.yml \
  --vault-id prod@~/.vault_pass_prod \
  --vault-id staging@~/.vault_pass_staging
```

## 49. Stratégies d'exécution

La stratégie définit **comment** les tâches sont réparties sur les hôtes.

```yaml
- hosts: webservers
  strategy: linear        # défaut : tâche N finie partout avant tâche N+1
  tasks: [...]

- hosts: webservers
  strategy: free          # chaque hôte avance à son rythme (plus rapide, logs mélangés)

- hosts: webservers
  strategy: host_pinned   # comme free mais un hôte reste sur le même worker
```

| Stratégie | Cas d'usage |
|---|---|
| `linear` (défaut) | Déploiements ordonnés, la plupart des cas |
| `free` | Tâches longues indépendantes (mises à jour, sauvegardes) |
| `host_pinned` | `free` + affinité hôte/worker (gros parcs) |
| `debug` | Pas à pas interactif (débogage) |

## 50. Parallélisme : forks, serial, throttle

```ini
# ansible.cfg — forks globaux
[defaults]
forks = 30
```

```yaml
# serial : vagues successives (rolling update)
- hosts: webservers
  serial: 2                 # 2 hôtes à la fois
  # serial: "25%"            # ou en pourcentage
  # serial: [1, 5, 10]       # 1 puis 5 puis 10 (montée en charge prudente)
  tasks:
    - name: Déployer la nouvelle version
      ...

# throttle : limite globale même avec beaucoup de forks
- name: Appel API limité à 3 simultanés
  ansible.builtin.uri:
    url: https://api.example.lan/deploy
  throttle: 3

