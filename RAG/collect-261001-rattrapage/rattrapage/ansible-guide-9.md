---
id: collect-261001-rattrapage/rattrapage/ansible-guide-9
title: "Guide Ansible — Automatisation système en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/ansible_guide.md
source_anchor: ""
source_lines: [2067, 2334]
sha256: 6d4ccd8b3027481a96e516d27ce40780fd53057199985766fc6abaff457fa717
---

    # --- Services ---
    - name: S'assurer que nginx et PHP-FPM tournent
      ansible.builtin.service:
        name: "{{ item }}"
        state: started
        enabled: true
      loop: [nginx, "php{{ php_version }}-fpm"]
      tags: [service]

  handlers:
    - name: Recharger nginx
      ansible.builtin.service:
        name: nginx
        state: reloaded

  post_tasks:
    # --- Vérification finale ---
    - name: Vérifier que le site répond 200
      ansible.builtin.uri:
        url: "http://localhost/"
        status_code: 200
      register: check_site
      retries: 3
      delay: 5
      until: check_site.status == 200
```

```bash
ansible-playbook playbooks/deployer_web.yml --check --diff -l web-01
ansible-playbook playbooks/deployer_web.yml -l web-01
curl -I http://web-01/
```

## 66. Cas pratique 2 : durcir SSH sur 50 serveurs

Objectif : appliquer le rôle `durcissement_ssh` (section 43) à tout le parc,
par vagues, avec rollback manuel documenté.

```yaml
---
- name: Durcissement SSH du parc
  hosts: all
  become: true
  serial: 5                    # 5 serveurs à la fois : on garde l'accès en cas de pépin
  max_fail_percentage: 20      # stop si > 20 % d'échecs dans une vague

  pre_tasks:
    - name: Sauvegarder sshd_config actuel
      ansible.builtin.copy:
        src: /etc/ssh/sshd_config
        dest: /root/sshd_config.avant_ansible
        remote_src: true
        mode: "0600"

  roles:
    - role: durcissement_ssh
      vars:
        ssh_port: 2222
        ssh_permit_root_login: "no"

  post_tasks:
    - name: Vérifier que le nouveau port répond (avant de couper l'ancien accès)
      ansible.builtin.wait_for:
        host: "{{ ansible_host }}"
        port: 2222
        timeout: 10
      delegate_to: localhost
      become: false
```

Points de vigilance :

- [ ] Testez d'abord sur **1 serveur de staging** avec `--check --diff`.
- [ ] Gardez une session SSH ouverte sur chaque vague pendant l'application.
- [ ] La sauvegarde `/root/sshd_config.avant_ansible` permet un rollback :
      `ansible all -b -m ansible.builtin.command -a "cp /root/sshd_config.avant_ansible /etc/ssh/sshd_config && systemctl restart ssh"`.
- [ ] Mettez à jour l'inventaire (`ansible_port=2222`) **après** validation.

## 67. Cas pratique 3 : gérer les utilisateurs sur un parc

Objectif : créer les comptes d'équipe depuis une variable centrale, déployer les
clés SSH, supprimer les comptes obsolètes.

```yaml
# group_vars/all.yml
utilisateurs_equipe:
  - nom: alice
    groupes: [sudo, docker]
    cle_ssh: "ssh-ed25519 AAAA... alice@poste"
    etat: present
  - nom: bob
    groupes: [docker]
    cle_ssh: "ssh-ed25519 BBBB... bob@poste"
    etat: present
  - nom: charlie          # départ : compte à supprimer
    etat: absent
```

```yaml
---
- name: Gérer les comptes utilisateurs du parc
  hosts: all
  become: true

  tasks:
    - name: Créer / supprimer les comptes
      ansible.builtin.user:
        name: "{{ item.nom }}"
        groups: "{{ item.groupes | default([]) }}"
        append: true
        shell: /bin/bash
        state: "{{ item.etat }}"
        remove: "{{ item.etat == 'absent' }}"   # supprime aussi le home si absent
      loop: "{{ utilisateurs_equipe }}"
      loop_control:
        label: "{{ item.nom }}"

    - name: Déployer les clés SSH (comptes présents uniquement)
      ansible.posix.authorized_key:
        user: "{{ item.nom }}"
        key: "{{ item.cle_ssh }}"
        state: present
      loop: "{{ utilisateurs_equipe }}"
      when: item.etat == 'present' and item.cle_ssh is defined
      loop_control:
        label: "{{ item.nom }}"
```

`loop_control.label` rend la sortie lisible (affiche juste le nom au lieu de tout
le dictionnaire). `remove: true` sur un compte `absent` nettoie le home :
à n'utiliser qu'après validation avec le management.

## 68. Cas pratique 4 : campagne de mises à jour planifiée

Objectif : patching mensuel des serveurs Debian/Ubuntu, par vagues, avec reboot
conditionnel et fenêtre de maintenance.

```yaml
---
- name: Patching mensuel
  hosts: all
  become: true
  serial: "25%"
  max_fail_percentage: 10

  tasks:
    - name: Mise à jour des paquets (dist-upgrade)
      ansible.builtin.apt:
        upgrade: dist
        update_cache: true
        cache_valid_time: 3600
      register: resultat_apt

    - name: Vérifier si un reboot est requis
      ansible.builtin.stat:
        path: /var/run/reboot-required
      register: reboot_requis

    - name: Rebooter si nécessaire (fenêtre de maintenance)
      ansible.builtin.reboot:
        reboot_timeout: 600
        msg: "Reboot planifié après patching Ansible"
      when:
        - reboot_requis.stat.exists
        - autoriser_reboot | default(false)

    - name: Attendre le retour du serveur
      ansible.builtin.wait_for_connection:
        timeout: 300
      when: reboot_requis.stat.exists and (autoriser_reboot | default(false))
```

```bash
# Fenêtre de maintenance : on autorise explicitement le reboot
ansible-playbook playbooks/patching.yml -e "autoriser_reboot=true" -l staging
```

## 69. Dépannage : méthodologie générale

1. **Reproduire en ciblant un seul hôte** : `-l hote-en-panne --check --diff`.
2. **Monter la verbosité** : `-v` (résultats), `-vv` (fichiers transférés),
   `-vvv` (connexions SSH), `-vvvv` (debug connexion brute).
3. **Lire le message d'erreur réel** : Ansible affiche souvent la cause
   (`FAILED! => {"msg": "..."}`) — ne la survolez pas.
4. **Isoler** : `ansible hote -m ansible.builtin.ping` puis le module seul en ad-hoc.
5. **Vérifier les classiques** : YAML valide ? Variables définies ? Droits `become` ?
6. **Consulter** : `ansible-doc <module>`, logs `log_path`, ARA (section 62).

```bash
# Valider la syntaxe YAML de tous les playbooks
for f in playbooks/*.yml; do python3 -c "import yaml,sys; yaml.safe_load(open('$f'))" && echo "OK $f"; done

# Vérifier un playbook sans l'exécuter (syntax check)
ansible-playbook playbooks/site.yml --syntax-check
```

## 70. Dépannage : cas concret « UNREACHABLE » SSH

Symptôme :

```text
fatal: [web-01]: UNREACHABLE! => {"changed": false, "msg": "Failed to connect to the host via ssh: ..."}
```

Checklist de résolution :

```bash
# 1. Le SSH manuel fonctionne-t-il avec la même clé ?
ssh -i ~/.ssh/id_ed25519_ansible deploy@web-01

# 2. Ansible utilise-t-il la bonne clé et le bon utilisateur ?
ansible web-01 -m ansible.builtin.ping -vvvv 2>&1 | grep -E "SSH|identity|user"

# 3. L'hôte est-il bien dans l'inventaire avec la bonne IP ?
ansible-inventory --host web-01

# 4. Le port est-il ouvert ?
ansible web-01 -m ansible.builtin.wait_for -a "port=22 timeout=5"  # via un autre chemin
```

Causes fréquentes : clé non déployée, `ansible_port` oublié après changement de
port, `known_hosts` corrompu (`ssh-keygen -R web-01`), pare-feu, bastion mal configuré.

## 71. Dépannage : cas concret « become / sudo »

Symptôme :

```text
fatal: [db-01]: FAILED! => {"msg": "Missing sudo password"}
```

ou :

```text
"msg": "sudo: a password is required"
```

Résolutions :

```bash
# L'utilisateur a-t-il le droit sudo NOPASSWD ?
ansible db-01 -b -m ansible.builtin.command -a "sudo -n true && echo OK"

# Si mot de passe requis : demander interactivement
ansible-playbook site.yml --ask-become-pass   # ou -K

# Vérifier la config sudoers côté cible
ansible db-01 -m ansible.builtin.command -a "sudo -l"
```

Autre piège : `pipelining = True` + `requiretty` activé dans sudoers →
désactivez `requiretty` (`Defaults !requiretty` dans `/etc/sudoers.d/ansible`)
ou repassez `pipelining = False`.

## 72. Dépannage : cas concret « template / Jinja2 »

Symptôme :

```text
fatal: [web-01]: FAILED! => {"msg": "AnsibleUndefinedVariable: 'dict object' has no attribute 'cle'"}
```

Résolutions :

```bash
# 1. Voir la valeur réelle de la variable
ansible web-01 -m ansible.builtin.debug -a "var=ma_variable"

