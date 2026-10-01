---
id: collect-261001-rattrapage/rattrapage/ansible-guide-11
title: "Guide Ansible — Automatisation système en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention", "distribution"]
source: docs/RAG/collect-261001-rattrapage/ansible_guide.md
source_anchor: ""
source_lines: [2542, 2694]
sha256: d4a3029c451f255afdad7cfaaab37d69433e87a1d8ad4d5dffd3f746918a257d
---

# Après mise à jour : vérifier les dépréciations
ansible-playbook playbooks/site.yml --check 2>&1 | grep -i deprecat
```

Points d'attention entre versions majeures :

- Lire le **porting guide** (`docs.ansible.com` → Porting Guides) : suppressions
  de modules, renommages (`ansible.builtin.` stabilisé depuis la 2.10).
- Tester en staging avec `--check --diff` avant de basculer le control node de prod.
- Les collections évoluent indépendamment : `ansible-galaxy collection install -r
  collections/requirements.yml --upgrade`.

## 81. Pense-bête de poche

```bash
# --- Inventaire ---
ansible-inventory -i prod.ini --graph
ansible-inventory -i prod.ini --list | python3 -m json.tool | less

# --- Ad-hoc ---
ansible all -m ansible.builtin.ping
ansible web -b -m ansible.builtin.service -a "name=nginx state=reloaded"
ansible all -b -m ansible.builtin.apt -a "upgrade=dist update_cache=yes" --check --diff

# --- Playbooks ---
ansible-playbook site.yml --syntax-check
ansible-playbook site.yml --check --diff
ansible-playbook site.yml -l staging --tags config
ansible-playbook site.yml --vault-password-file ~/.vault_pass -f 30
ansible-playbook site.yml --step          # confirmation tâche par tâche
ansible-playbook site.yml --start-at-task "Déployer le vhost"

# --- Vault ---
ansible-vault create group_vars/all/vault.yml
ansible-vault edit group_vars/all/vault.yml
ansible-vault encrypt_string 'secret' --name 'token'

# --- Debug ---
ansible-playbook site.yml -vvv -l hote-en-panne
ansible hote -m ansible.builtin.setup -a "filter=ansible_distribution*"
ansible-doc ansible.builtin.template

# --- Galaxy ---
ansible-galaxy collection install -r collections/requirements.yml
ansible-galaxy role install -r roles/requirements.yml
```

## 82. Glossaire

| Terme | Définition |
|---|---|
| **Ad-hoc** | Commande Ansible unique, sans playbook |
| **Become** | Élévation de privilèges (sudo, su…) |
| **Callback** | Plugin qui traite la sortie / les événements d'exécution |
| **Check mode** | Dry-run (`--check`) : simule sans modifier |
| **Collection** | Paquet distribuable : modules, plugins, rôles |
| **Control node** | Machine où Ansible est installé et lancé |
| **Diff** | Affichage des différences (`--diff`) avant/après |
| **Facts** | Informations collectées sur les nœuds gérés |
| **Forks** | Processus parallèles d'exécution (défaut : 5) |
| **FQCN** | Fully Qualified Collection Name (`ansible.builtin.apt`) |
| **Galaxy** | Hub communautaire de rôles et collections |
| **Gather facts** | Phase de collecte des facts en début de play |
| **Handler** | Tâche exécutée une fois en fin de play si notifiée |
| **Host vars** | Variables spécifiques à un hôte |
| **Group vars** | Variables spécifiques à un groupe |
| **Idempotence** | Rejouer sans effet de bord si l'état est déjà correct |
| **Inventory** | Inventaire : hôtes, groupes, variables de connexion |
| **Jinja2** | Moteur de templates (`{{ }}`, `{% %}`) |
| **Managed node** | Machine gérée via SSH/WinRM, sans agent |
| **Module** | Unité d'action idempotente |
| **Notify** | Déclenchement d'un handler depuis une tâche |
| **Play** | Ensemble de tâches appliquées à des hôtes |
| **Playbook** | Fichier YAML contenant un ou plusieurs plays |
| **Role** | Regroupement réutilisable : tasks, handlers, templates… |
| **Serial** | Exécution par vagues successives d'hôtes |
| **Tags** | Étiquettes pour exécuter un sous-ensemble de tâches |
| **Task** | Appel d'un module avec ses paramètres |
| **Throttle** | Limite de concurrence sur une tâche |
| **Vault** | Chiffrement des variables sensibles |
| **AWX / Controller** | Interface web, ordonnanceur et RBAC au-dessus d'Ansible |

## 83. Quiz — 10 questions

1. Quelle est la différence entre `ansible` et `ansible-core` ?
2. Pourquoi Ansible est-il qualifié d'« agentless » et quel protocole utilise-t-il
   par défaut sur Linux ?
3. Citez les 4 niveaux de précédence des variables à connaître par cœur, du moins
   au plus prioritaire.
4. Que se passe-t-il si deux tâches différentes notifient le même handler et que
   les deux sont `changed` ?
5. À quoi sert le paramètre `validate` du module `template`, et sur quel fichier
   critique l'utiliseriez-vous en priorité ?
6. Expliquez la différence entre `serial: "25%"` et `throttle: 3`.
7. Votre playbook échoue avec `AnsibleUndefinedVariable`. Citez 3 vérifications
   dans l'ordre.
8. Pourquoi faut-il éviter `ignore_errors: true` sur un play entier, et que
   proposer à la place ?
9. Où stocker le mot de passe Vault pour une exécution automatisée (CI), et avec
   quels droits ?
10. Votre second run d'un playbook affiche encore des `changed`. Est-ce normal ?
    Que devez-vous vérifier ?

## 84. Quiz — réponses

1. `ansible-core` est le moteur seul (playbooks, modules `ansible.builtin`) ;
   `ansible` (paquet communautaire) ajoute des dizaines de collections. En
   production, partez du core et n'ajoutez que le nécessaire.
2. Aucun agent à installer sur les cibles : Ansible pousse des scripts Python via
   **SSH** (SFTP/SCP) et récupère le résultat en JSON. Sur Windows : WinRM.
3. `defaults/` du rôle < `group_vars/` < `host_vars/` < `--extra-vars` (`-e`).
4. Le handler ne s'exécute **qu'une seule fois** en fin de play, quel que soit le
   nombre de notifications.
5. `validate` exécute une commande (ex. `sshd -t -f %s`, `nginx -t -c %s`) sur le
   fichier généré **avant** de remplacer la destination : si la validation échoue,
   le fichier en production n'est pas touché. Priorité : `sshd_config` (un fichier
   invalide = lockout SSH).
6. `serial` découpe les **hôtes** en vagues successives (rolling update) ;
   `throttle` limite le nombre d'exécutions **simultanées d'une tâche** même avec
   beaucoup de forks (ex. : ne pas saturer une API).
7. (a) `ansible -m debug -a "var=ma_variable"` pour voir la valeur réelle ;
   (b) vérifier l'orthographe et la portée (host_vars vs group_vars) ;
   (c) ajouter `| default(...)` ou `is defined` si la variable est optionnelle.
8. Parce qu'il masque tous les échecs, y compris les vrais problèmes, et laisse le
   parc dans un état inconnu. À la place : `block/rescue/always` ciblé sur les
   tâches à risque, avec alerte explicite dans le `rescue`.
9. Dans un fichier dédié (ex. `~/.vault_pass`) en `chmod 600`, hors Git, de
   préférence fourni par un gestionnaire de secrets ou une variable CI masquée —
   jamais dans le dépôt ni en clair dans un script.
10. Non : un playbook idempotent doit afficher `changed=0` au second run (hors
    tâches explicitement non idempotentes comme `command` sans `changed_when`).
    Vérifiez les tâches `changed` : utilisez un module dédié, ajoutez
    `changed_when`/`creates`, ou corrigez le template qui génère un diff à chaque fois.

## 85. Pour aller plus loin

- **Documentation officielle** : `docs.ansible.com` — le porting guide à chaque
  montée de version, et `ansible-doc` en local (plus rapide que le web).
- **Molecule** : tester vos rôles dans des conteneurs (`pipx install molecule
  molecule-plugins[docker]`) avant de toucher au staging.
- **Ansible Lint** : `pipx install ansible-lint` puis `ansible-lint playbooks/` —
  intégrez-le en CI pour uniformiser le style d'équipe.
- **AWX** : quand plusieurs opérateurs et de la planification entrent en jeu
  (section 64).
- **Terraform + Ansible** : Terraform provisionne (VM, réseau), Ansible configure —
  le duo standard, pas des concurrents.
- **Communauté** : Galaxy pour les rôles/collections, forum `forum.ansible.com`.
- **Prochaine étape concrète** : prenez le cas pratique 1 (section 65), adaptez-le
  à votre intranet, versionnez-le en Git, et planifiez le patching (section 68)
  via un cron sur le control node ou un job AWX.

---

