---
id: collect-261001-rattrapage/rattrapage/ansible-guide-10
title: "Guide Ansible — Automatisation système en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/ansible_guide.md
source_anchor: ""
source_lines: [2335, 2541]
sha256: 34f0fde6d3b0dfd4a0d3e5d053d4c561fee4bc266991fff5e796e253c973173a
---

# 2. Tester le rendu du template en dry-run avec diff
ansible-playbook site.yml --check --diff -l web-01 --tags config

# 3. Rendre le template tolérant
# {{ ma_variable | default('valeur_par_defaut') }}
# {% if ma_variable is defined %}...{% endif %}
```

Pièges fréquents : variable définie dans `host_vars` d'un autre hôte, faute de
frappe (`http_port` vs `http_ports`), filtre appliqué à `None`, accolades
littérales dans le fichier destination (échappez avec `{{ '{{' }}`).

## 73. Dépannage : cas concret « YAML »

Symptôme : `ERROR! We were unable to read either as JSON nor YAML`.

```yaml
# FAUX : tabulation (YAML interdit les tabulations)
- name: Test
\tansible.builtin.ping:

# FAUX : ":" sans espace dans une valeur
- name: Test
  ansible.builtin.debug:
    msg: "erreur:ici"        # OK car quoté ; sans quotes ce serait une erreur

# FAUX : expression Jinja2 non quotée en début de valeur
  port: {{ http_port }}       # ERREUR
  port: "{{ http_port }}"     # OK

# FAUX : liste mal indentée
  loop:
  - a
   - b                        # indentation incohérente

# JUSTE
  loop:
    - a
    - b
```

Outils : `yamllint` (`pipx install yamllint`), `--syntax-check`, et l'option
`ansible.builtin.debug: var=` pour inspecter.

## 74. Dépannage : cas concret « handler jamais déclenché »

Symptôme : le fichier est déployé mais le service n'est pas rechargé.

Causes et remèdes :

```yaml
# 1. Le notify ne correspond pas EXACTEMENT au nom du handler
notify: Recharger nginx        # doit matcher "name: Recharger nginx"

# 2. La tâche n'est pas "changed" (fichier identique) : normal !
#    Forcer si besoin :
- name: Forcer le rechargement
  ansible.builtin.command: nginx -s reload
  changed_when: true

# 3. Handler défini dans le mauvais play (les handlers sont par play)
# 4. Utiliser listen pour regrouper plusieurs notifies
handlers:
  - name: Recharger le frontal web
    listen: recharger_frontaux
    ansible.builtin.service:
      name: nginx
      state: reloaded
# notify: recharger_frontaux  (insensible au renommage du name)
```

## 75. Dépannage : playbook lent — diagnostic

```bash
# Identifier les tâches lentes
ANSIBLE_CALLBACKS_ENABLED=ansible.posix.profile_tasks ansible-playbook site.yml 2>&1 | tail -20

# Causes typiques et remèdes :
# - gather_facts sur des centaines d'hôtes  -> fact_caching (section 60)
# - apt update_cache à chaque tâche          -> cache_valid_time: 3600
# - forks = 5 par défaut                     -> forks = 30 dans ansible.cfg
# - pas de pipelining                        -> pipelining = True
# - git clone complet à chaque run           -> le module git est idempotent, vérifier version figée
# - handlers redémarrant en boucle           -> un seul notify, handlers en fin de play
```

## 76. Sauvegarde et restauration d'un nœud via Ansible

Exemple : sauvegarder `/etc` et la liste des paquets avant une opération risquée.

```yaml
---
- name: Snapshot pré-intervention
  hosts: all
  become: true
  gather_facts: false

  vars:
    dossier_sauvegarde: /srv/snapshots/{{ inventory_hostname }}

  tasks:
    - name: Créer le dossier de snapshot (sur le control node)
      ansible.builtin.file:
        path: "{{ dossier_sauvegarde }}"
        state: directory
        mode: "0700"
      delegate_to: localhost
      become: false

    - name: Archiver /etc
      community.general.archive:
        path: /etc
        dest: /tmp/etc_backup.tar.gz
        format: gz

    - name: Récupérer l'archive
      ansible.builtin.fetch:
        src: /tmp/etc_backup.tar.gz
        dest: "{{ dossier_sauvegarde }}/etc_backup.tar.gz"
        flat: true

    - name: Sauvegarder la liste des paquets
      ansible.builtin.command: dpkg --get-selections
      register: paquets
      changed_when: false

    - name: Écrire la liste des paquets
      ansible.builtin.copy:
        content: "{{ paquets.stdout }}"
        dest: "{{ dossier_sauvegarde }}/paquets.txt"
        mode: "0600"
      delegate_to: localhost
      become: false
```

Restauration paquets : `dpkg --set-selections < paquets.txt && apt-get dselect-upgrade`.

## 77. Les 12 erreurs classiques (et comment les éviter)

| # | Erreur | Symptôme | Remède |
|---|---|---|---|
| 1 | Secrets en clair dans Git | Fuite de mots de passe | Vault (sections 46–48) + `.gitignore` |
| 2 | `shell` au lieu d'un module | Non idempotent, `changed` à chaque run | Utiliser le module dédié + `changed_when` |
| 3 | `notify` qui ne matche pas le handler | Service jamais rechargé | Noms identiques ou `listen` (section 74) |
| 4 | `gather_facts: true` inutile | Playbook lent | `gather_facts: false` ou cache (section 60) |
| 5 | Pas de `--check --diff` avant prod | Changement destructeur surprise | Dry-run systématique + staging d'abord |
| 6 | `serial` oublié sur rolling update | Tout le parc redémarre en même temps | `serial: "25%"` + `max_fail_percentage` |
| 7 | Variables non quotées `{{ }}` | Erreur de parsing YAML | Toujours quoter : `"{{ var }}"` |
| 8 | `become` sans sudoers adapté | `Missing sudo password` | Tester `sudo -n true`, configurer sudoers |
| 9 | Inventaire prod/staging mélangé | Patch appliqué au mauvais env | Un fichier par environnement (section 12) |
| 10 | `ignore_errors: true` abusif | Échecs masqués, état inconnu | `block/rescue` ciblé, jamais global |
| 11 | Pas de `validate` sur les templates critiques | `sshd_config` invalide → lockout | `validate: "sshd -t -f %s"` (sections 22, 28) |
| 12 | `latest` / `upgrade: yes` sans contrôle | Version surprise en production | Épingler les versions, upgrades planifiés |

## 78. Bonnes pratiques production — checklist

- [ ] Chaque playbook commence par `--check --diff` puis passe par le staging.
- [ ] Handlers pour les restarts, jamais de `service restarted` direct dans les tâches
      (évite les redémarrages en boucle à chaque run).
- [ ] `validate` sur tout template critique (`sshd_config`, `nginx.conf`, `sudoers`).
- [ ] Idempotence vérifiée : deux runs successifs → le second affiche `changed=0`.
- [ ] Tags cohérents (`install`, `config`, `service`, `always`, `never`).
- [ ] `serial` + `max_fail_percentage` sur tout changement à impact (reboot, upgrade).
- [ ] Logs conservés 90 jours (`log_path` + rotation logrotate).
- [ ] Code relu (merge request) avant merge sur `main`.
- [ ] Versions épinglées : collections (`requirements.yml`), paquets critiques.
- [ ] `assert` en prérequis OS au début des playbooks sensibles.
- [ ] Pas de `shell`/`command` sans `changed_when`/`failed_when` explicites.
- [ ] Documentation : `README.md` par rôle, en-tête `ansible_managed` dans les templates.

## 79. Sécurité — checklist durcissement

- [ ] Vault pour tous les secrets ; mots de passe Vault distincts prod/staging.
- [ ] Fichier `~/.vault_pass` en `chmod 600`, hors Git, rotation au départ d'un équipier.
- [ ] Clé SSH dédiée à Ansible (`id_ed25519_ansible`), jamais la clé personnelle.
- [ ] `host_key_checking = True` en production.
- [ ] Utilisateur `deploy` avec sudo restreint si possible (pas `NOPASSWD: ALL`
      sans réflexion — voir section 58).
- [ ] `no_log: true` sur les tâches manipulant des secrets en clair transitoire :

```yaml
- name: Appel API avec token (ne pas logger)
  ansible.builtin.uri:
    url: https://api.example.lan/secret
    headers:
      Authorization: "Bearer {{ vault_api_token }}"
  no_log: true
```

- [ ] Audit régulier : `git log -p --all -S "password"` pour traquer les secrets
      ayant fuité dans l'historique.
- [ ] AWX/Controller : credentials en base chiffrée, RBAC, pas de Vault partagé
      par messagerie.

## 80. Mise à jour d'Ansible

```bash
# Via pipx (recommandé)
pipx upgrade ansible
ansible --version

# Épingler une version précise (reproductibilité)
pipx install --force ansible==10.7.0

# Via APT (Debian/Ubuntu)
sudo apt update && sudo apt install --only-upgrade ansible

