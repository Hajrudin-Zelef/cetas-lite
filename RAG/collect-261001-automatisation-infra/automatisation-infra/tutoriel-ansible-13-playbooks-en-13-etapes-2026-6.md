---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026-6
title: "Mise à jour des paquets et installation de pipx"
domain: automatisation-infra
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent", "arr", "aws", "distribution", "incident", "mai"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026.md
source_anchor: ""
source_lines: [639, 703]
sha256: 5fded18bbf28b66de581eabc82016c8df8f80922dde5cfad9fb19c89e9eec4a0
---

# Mise à jour des paquets et installation de pipx

| Erreur | Cause probable | Solution | 
|---|---|---|
| UNREACHABLE — ssh: Permission denied | Clé SSH non autorisée ou utilisateur incorrect | ssh-copy-id ou vérifier ansible_user | 
| FAILED — sudo: a password is required | become_ask_pass: false sans NOPASSWD | Configurer /etc/sudoers ou –ask-become-pass | 
| ERROR — vault password is not set | Fichier Vault sans mot de passe fourni | –vault-password-file ou ANSIBLE_VAULT_PASSWORD_FILE | 
| FAILED — ansible_host_key_checking | Hôte inconnu de known_hosts | ssh-keyscan ou host_key_checking = False | 
| ERROR — Failed to import the required Python library | Module Python manquant sur la cible | Installer pip3 install [lib] ou utiliser pip module | 
| FAILED — Could not find or access [file] | Chemin relatif ambigu | Préfixer par {{ playbook_dir }} ou {{ role_path }} | 
| Lent à exécuter (5+ min sur 10 hôtes) | forks=5 par défaut, pas de pipelining | Augmenter forks à 50 + pipelining=True | 
| Variable undefined | Précédence non comprise | ansible-inventory –host hostname pour debug | 
| Handler not triggered | Tâche n’a pas changé l’état | Vérifier la tâche notifie le bon nom | 
| Module deprecated warning | Module legacy < 2.18 | Migrer vers le FQCN ansible.builtin.* | 

## Conseils avancés pour la production

Plusieurs techniques avancées élèvent vos playbooks au niveau enterprise. Les **blocks** regroupent des tâches sous une même condition `when` et permettent un `rescue` en cas d’erreur, à la manière d’un try/catch. Les **tags** vous laissent exécuter une partie ciblée du playbook (`--tags db,backup`) au lieu de tout rejouer. Le module `ansible.builtin.assert` valide les pré-conditions et bloque le déploiement si une variable critique manque ou a une valeur invalide, évitant les états partiellement corrompus.

Pour les déploiements zéro-downtime sur un cluster derrière un load balancer, combinez `serial: 1` (un nœud à la fois) avec le module `community.general.haproxy` qui retire le nœud du backend avant déploiement et le réintègre après health-check. Couplée à `any_errors_fatal: true`, cette stratégie garantit qu’un échec sur un nœud arrête immédiatement la propagation. Pour le monitoring, exposez les métriques Ansible vers Prometheus via le callback `ansible.posix.profile_tasks` et tracez chaque exécution avec un identifiant de run unique stocké dans l’inventaire dynamique.

Enfin, pour les très grandes flottes (1 000+ hôtes), passez à **Ansible Automation Platform 2.6** qui apporte AutomationHub privé, Execution Environments containerisés, RBAC granulaire, ordonnancement, et l’intégration native d’**Event-Driven Ansible**. Côté tarification, TrustRadius situait déjà en mai 2025 les offres payantes de Red Hat AAP entre 5 et 14 dollars par unité, sans essai gratuit ; à l’échelle d’un vrai parc, une comparaison Tech Insider d’avril 2026 chiffre le budget annuel entre 65 000 et 87 500 dollars pour les formules Standard/Premium sur 500 nœuds. Ce niveau d’investissement se justifie largement pour les ESN et grandes entreprises soumises à des audits ANSSI ou ISO 27001, grâce à la valeur ajoutée en audit, rollback et gouvernance.

## Comparatif Ansible vs alternatives en 2026

| Critère | Ansible 13 | Terraform 1.10 | Puppet 8 | Chef 18 | SaltStack 3007 | 
|---|---|---|---|---|---|
| Modèle | Push, agentless | Push, agentless | Pull, agent | Pull, agent | Push/Pull, agent | 
| Langage | YAML + Jinja2 | HCL | Puppet DSL | Ruby DSL | YAML + Python | 
| Cas d’usage primaire | Configuration + déploiement | Provisionnement cloud | Configuration | Configuration | Configuration + orchestration | 
| Courbe d’apprentissage | Faible | Moyenne | Forte | Forte | Moyenne | 
| Communauté GitHub stars | 68k | 43k | 7,5k | 7,7k | 14k | 
| Adoption France 2025 | ~62% | ~58% | ~14% | ~9% | ~7% | 
| Plateforme entreprise | AAP 2.6 | Terraform Cloud/HCP | PE 2025 | Chef 360 | Saltstack Enterprise | 

## FAQ — Questions fréquentes sur les playbooks Ansible

### Quelle est la différence entre ansible-core et le paquet Ansible community ?

`ansible-core` est le moteur minimal qui n’embarque qu’une vingtaine de modules essentiels (`ansible.builtin.*`). Le paquet **Ansible 13.6.0**, lui, est une distribution qui empaquette ansible-core 2.20 plus une centaine de collections vérifiées (community.general, ansible.posix, amazon.aws, etc.). Pour un projet typique, installez le paquet community ; pour un Execution Environment minimaliste, partez d’ansible-core et ajoutez vos collections via `requirements.yml`.

### Ansible peut-il remplacer Terraform ?

Non, ils sont complémentaires. Terraform excelle dans le **provisionnement déclaratif d’infrastructure** (créer des VM, des VPC, des bases RDS) avec un state file qui suit l’existant. Ansible excelle dans la **configuration et le déploiement applicatif** sur des machines existantes. La pratique standard en 2026 consiste à provisionner avec Terraform puis configurer avec Ansible, idéalement via le provider `ansible.cloud` qui crée automatiquement l’inventaire dynamique depuis l’état Terraform.

### Comment gérer les serveurs Windows avec Ansible ?

La collection `ansible.windows` et `community.windows` fournissent les modules `win_*` qui pilotent Windows Server 2019/2022/2025 via WinRM ou PSRP. Activez WinRM côté serveur avec `winrm quickconfig`, puis configurez les variables `ansible_connection: winrm`, `ansible_winrm_transport: kerberos` et l’authentification par compte de service AD. Tous les modules YAML restent identiques côté playbook.

### Comment déboguer un playbook qui échoue silencieusement ?

Augmentez la verbosité avec `-v`, `-vv`, `-vvv` ou `-vvvv` (ce dernier affiche tout le trafic SSH). Utilisez `register` pour capturer la sortie d’une tâche puis `debug: var=ma_variable` pour l’inspecter. Le module `ansible.builtin.pause` avec `prompt` arrête l’exécution pour vous laisser inspecter la cible. Enfin, `ANSIBLE_DEBUG=1` en variable d’environnement active des logs très détaillés sur le contrôleur.

### Combien d’hôtes Ansible peut-il gérer simultanément ?

En pratique, un poste de contrôle bien configuré (`forks=100`, pipelining, ControlPersist, fact_caching) gère sans difficulté 1 000 à 5 000 hôtes en parallèle. Au-delà, l’architecture distribuée d’**Ansible Automation Platform 2.6** avec plusieurs Execution Nodes et un mesh de Receptor permet de cibler des dizaines de milliers de nœuds. Red Hat publie des cas d’usage à 50 000 hôtes chez de grandes banques européennes.

### Faut-il certifier RHCE pour utiliser Ansible en production ?

La certification **Red Hat Certified Engineer (RHCE)** EX294 valide les compétences Ansible et reste reconnue par les grands comptes français. Elle n’est pas obligatoire pour utiliser l’outil mais valorise nettement les CV : un RHCE confirmé en France 2026 facture 15 à 20 % de plus qu’un profil non certifié, selon les baromètres LeHibou et Free-Work. L’examen coûte environ 500 euros et se passe à distance.

### Event-Driven Ansible est-il prêt pour la production ?

Oui, Event-Driven Ansible est intégré à **Ansible Automation Platform 2.6** en version supportée. Il s’appuie sur le moteur `ansible-rulebook` qui consomme des sources d’événements (Kafka, webhook, syslog, alertmanager) et déclenche des actions définies en YAML. Cas typiques : redémarrer un service après alerte Prometheus, ajouter un nœud au load balancer après scaling AWS, créer un ticket ServiceNow sur incident syslog.

### Comment migrer d’Ansible 2.18 vers Ansible 2.20 ?

