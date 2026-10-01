---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026-7
title: "Mise à jour des paquets et installation de pipx"
domain: automatisation-infra
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "mai"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026.md
source_anchor: ""
source_lines: [704, 712]
sha256: 6ce429b9e3ce87bfe80a92658f60e7e103e1d58c67041f638ef2b0929ee8e811
---

# Mise à jour des paquets et installation de pipx

Trois étapes : 1) Vérifier que tous les hôtes contrôlés ont Python 3.9+ et que le contrôleur a Python 3.12+ — le guide de portage d’ansible-core 2.17, mis à jour en juillet 2026, rappelle d’ailleurs que le plancher historique pour l’exécution distante des playbooks reste Python 3.7+, un minimum désormais largement dépassé en pratique ; cette branche 2.17, qui a vu se succéder les correctifs 2.17.8 le 27 janvier 2025 puis 2.17.9 le 24 février 2025, a reçu sa dernière révision stable, la 2.17.14, le 8 septembre 2025 selon CompatHub, avant de céder la place aux lignes 2.19 puis 2.20 ; 2) Lancer `ansible-lint --profile production` pour identifier les modules dépréciés (notamment l’usage des noms courts sans FQCN) ; 3) Mettre à jour les collections avec `ansible-galaxy collection install -r requirements.yml --upgrade`. Testez le tout en staging avec `--check --diff` avant la bascule production. Comptez en moyenne 2 jours pour un projet de taille moyenne.

### Related Coverage

## Conclusion : Ansible reste le couteau suisse de l’automatisation 2026

Au terme de ce tutoriel, vous maîtrisez la chaîne complète : installation, structure de projet, inventaire dynamique, playbooks idempotents, rôles réutilisables, gestion des secrets avec Vault, intégration CI/CD et optimisations à grande échelle. Ansible 13.6.0 et ansible-core 2.20.5 constituent la base la plus stable jamais publiée, avec un support garanti jusqu’au 31 mai 2027 sur la 2.20. La feuille de route 2026 confirme l’investissement de Red Hat dans Event-Driven Ansible et l’Automation Platform 2.6, qui transforment Ansible en plateforme opérationnelle réactive.

Pour aller plus loin, explorez la galaxie Ansible Galaxy pour découvrir les collections communautaires, lisez les tips officiels et inscrivez-vous au forum Ansible où la communauté française est très active. Le projet final déployé en 4 minutes dans ce tutoriel sert de point de départ idéal pour automatiser vos propres infrastructures, qu’elles tournent sur OVHcloud, Scaleway, AWS, Azure ou Proxmox auto-hébergé. La prochaine version **ansible-core 2.21**, en bêta depuis le 13 avril 2026, introduira encore plus d’optimisations et arrivera en stable courant mai 2026.
