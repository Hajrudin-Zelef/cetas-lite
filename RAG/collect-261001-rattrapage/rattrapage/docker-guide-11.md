---
id: collect-261001-rattrapage/rattrapage/docker-guide-11
title: "Guide Docker complet — Production & Sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/docker_guide.md
source_anchor: ""
source_lines: [2477, 2525]
sha256: f79e7861a21dbdd4400d28629d364761f14c9f59d53dbe7e53d6e8a9757c91db
---

# Guide Docker complet — Production & Sysadmin

1. L'image est un modèle immuable (lecture seule) ; le conteneur est une
   instance en cours d'exécution de cette image, avec une couche inscriptible
   éphémère.
2. Le paquet `docker.io` est souvent en retard de versions (CVE non corrigées)
   et le snap pose des problèmes de permissions, volumes et réseau. Le dépôt
   officiel `download.docker.com` garantit des versions à jour et supportées.
3. Les fichiers `*-json.log` grossissent indéfiniment jusqu'à remplir
   `/var/lib/docker` à 100 %, ce qui bloque le daemon et peut planter l'hôte.
4. `depends_on` ne garantit que l'**ordre de démarrage**, pas que le service
   soit prêt à accepter des connexions. Il faut `condition: service_healthy`
   + un healthcheck sur la base + une logique de retry dans l'application.
5. (a) Docker initialise le volume avec le contenu/permissions de l'image
   (pas de problème d'UID comme avec les bind mounts) ; (b) le chemin est
   géré par Docker, portable entre hôtes ; (c) cycle de vie découplé du
   conteneur et adapté aux drivers distants (NFS…).
6. Quiconque contrôle le socket peut lancer un conteneur `--privileged`
   montant `/` de l'hôte → **root sur l'hôte**. Réservé aux outils d'admin
   légitimes, en `:ro` si possible.
7. `ARG` n'existe qu'au build (non persisté dans l'image finale) ; `ENV`
   persiste dans l'image et le conteneur (visible via `docker inspect`).
   Aucun des deux ne doit contenir de secret.
8. `latest` est un tag mouvant : on ne sait pas quelle version exacte tourne,
   impossible de faire un rollback fiable ni un déploiement reproductible.
9. Arrête la stack **et supprime les volumes nommés** → perte des données
   persistantes. Dangereux en production.
10. Au choix : utilisateur non-root (`USER`/`user:`), `read_only: true`,
    `cap_drop: [ALL]` + `cap_add` minimal, `no-new-privileges:true`,
    `pids_limit`, pas de `--privileged`, ports restreints, healthcheck.

## 84. Pour aller plus loin

- **Documentation officielle** : https://docs.docker.com (référence, guides,
  bonnes pratiques de sécurité)
- **Docker Bench for Security** : script d'audit CIS Docker
  (`docker run --rm -v /var/run/docker.sock:/var/run/docker.sock:ro ...`)
- **Rootless Docker** : faire tourner le daemon sans root
  (https://docs.docker.com/engine/security/rootless/)
- **Hadolint** : linter de Dockerfiles (à intégrer en CI)
- **Trivy** : scan de vulnérabilités images + IaC (https://aquasecurity.github.io/trivy/)
- **Harbor** : registre privé d'entreprise (RBAC, scan, réplication)
- **Kubernetes** : quand Swarm ne suffit plus (k3s pour commencer léger)
- **Livre** : *Docker in Practice* / *Docker Deep Dive* (Nigel Poulton)
- **Pratique** : refaites ce guide en labo — une stack LAMP, un Traefik, un
  Nextcloud — puis cassez-les volontairement (kill -9, disque plein,
  `down -v` sur une copie) pour apprendre le dépannage sans risque.

---

*Fin du guide — bon courage, et sauvegardez vos volumes.* 🐳
