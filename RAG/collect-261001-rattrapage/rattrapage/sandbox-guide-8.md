---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-8
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox", "memory"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [1576, 1795]
sha256: 200fa9192ced9c6a5f82b56b333acda373197df61b3178f7f0a516cba5d4397d
---

# 2. Changer durablement le type d'un chemin personnalisé
semanage fcontext -a -t httpd_sys_content_t "/srv/web(/.*)?"
restorecon -Rv /srv/web

# 3. Activer/désactiver une booléenne (ex. : httpd peut envoyer des mails)
setsebool -P httpd_can_sendmail on
getsebool -a | grep httpd

# 4. Comprendre un refus : lire l'audit
ausearch -m avc -ts recent | audit2why
# audit2why explique : "il manque la booléenne X" ou "contexte incorrect"

# 5. Générer un module de politique pour autoriser un refus légitime
ausearch -m avc -ts recent | audit2allow -M monmodule
semodule -i monmodule.pp
```

**Avertissement :** `audit2allow` génère des autorisations parfois trop larges. Relisez le `.te` avant `semodule -i`. Et ne désactivez jamais SELinux pour "résoudre" un problème : passez en Permissive, corrigez, repassez en Enforcing.

## 62. AppArmor vs SELinux : que choisir ?

| Critère | AppArmor | SELinux |
|---|---|---|
| Défaut sur | Debian, Ubuntu, openSUSE | RHEL, Fedora, CentOS |
| Modèle | Profils par programme (chemins) | Politique globale (types/étiquettes) |
| Courbe d'apprentissage | Modérée | Raide |
| Finesse | Bonne (fichiers, réseau, caps) | Excellente (tout objet) |
| Outils d'aide | aa-genprof, aa-logprof | audit2allow, setroubleshoot |

**Pour Zelef (Debian/Ubuntu) :** investissez sur AppArmor. Connaissez SELinux pour la culture et les environnements RHEL croisés en entreprise.

## 63. Checklist MAC

- [ ] `aa-status` : profils en enforce pour les services exposés (web, mail, DNS, impression)
- [ ] Nouveau binaire sensible → `aa-genprof` en recette avant production
- [ ] `aa-logprof` passé après chaque mise à jour majeure d'une appli confinée
- [ ] Sur RHEL : `getenforce` = Enforcing, `restorecon` connu par l'équipe
- [ ] Les refus MAC remontent dans la supervision (voir §73)

## 64. Erreurs classiques MAC

1. **Désactiver AppArmor/SELinux "pour que ça marche"** : on perd le filet sans corriger la cause. Toujours diagnostiquer le DENIED/AVC.
2. **`audit2allow` aveugle** : autorise parfois `dontaudit` massifs. Relire le `.te`.
3. **Profils AppArmor écrasés par les mises à jour** : personnalisations dans `/etc/apparmor.d/local/`, jamais dans le fichier du paquet.
4. **Oublier `restorecon` après `mv`** : `mv` conserve le contexte d'origine, `cp` applique celui de la destination. En cas de doute : `restorecon`.

---

# PARTIE VII — DOCKER COMME SANDBOX : LIMITES ET DURCISSEMENT

## 65. Docker isole, mais n'est pas une VM : les limites

Un conteneur partage le **noyau de l'hôte**. Conséquences :

| Menace | Conteneur Docker par défaut | VM KVM |
|---|---|---|
| Faille applicative | Contenue (namespaces) | Contenue |
| Faille noyau | **Compromet l'hôte** | Contenue (noyau séparé) |
| `--privileged` | Équivaut à root sur l'hôte | Sans objet |
| Montage `/var/run/docker.sock` | Contrôle total de l'hôte | Sans objet |
| Évasion par capabilities | Possible si mal configuré | Non |

**Règle :** ne faites jamais tourner du code *activement hostile* (malware à analyser) dans Docker sur un hôte de production. Pour l'analyse de malware : VM KVM isolée, snapshots, pas de réseau ou réseau simulé (INetSim).

## 66. Durcir un conteneur : les options indispensables

```bash
docker run -d --name monapp-durci \
  --read-only \                    # filesystem racine en lecture seule
  --tmpfs /tmp --tmpfs /run \      # tmpfs inscriptibles explicites
  --cap-drop=ALL \                 # retire TOUTES les capabilities...
  --cap-add=NET_BIND_SERVICE \     # ...sauf celle(s) nécessaire(s)
  --security-opt no-new-privileges \ # bloque setuid/setgid
  --pids-limit 100 \               # anti fork-bomb (cgroup pids)
  --memory 512m --memory-swap 512m \ # limite RAM (swap inclus)
  --cpus 1.0 \                     # limite CPU
  --network mon-reseau-isole \     # réseau dédié, pas le bridge par défaut
  --user 10000:10000 \            # ne PAS tourner en root dans le conteneur
  --restart on-failure:3 \
  monimage:1.2.3
```

Checklist de revue d'un `docker run` / compose :

- [ ] Pas de `--privileged`
- [ ] Pas de montage de `/var/run/docker.sock` (sauf cas justifié et audité)
- [ ] `--read-only` + tmpfs ciblés
- [ ] `--cap-drop=ALL` + `--cap-add` minimal
- [ ] `--user` non-root (ou `USER` dans le Dockerfile)
- [ ] Limites mémoire/CPU/pids
- [ ] Image épinglée par digest, pas `:latest`
- [ ] Réseau dédié, ports publiés minimaux

## 67. Seccomp et AppArmor sous Docker

Docker applique par défaut :

1. **Un profil seccomp** (`/usr/share/docker/seccomp/default.json`) qui bloque ~44 syscalls : `mount`, `ptrace`, `reboot`, `swapon`, `kexec_load`, `open_by_handle_at`, `init_module`, etc.
2. **Un profil AppArmor** `docker-default` (sur Ubuntu/Debian si AppArmor actif).

```bash
# Voir le profil AppArmor d'un conteneur
docker inspect monapp-durci --format '{{.AppArmorProfile}}'
# docker-default

# Désactiver seccomp (À NE FAIRE QU'EN DEBUG)
docker run --security-opt seccomp=unconfined ...

# Utiliser un profil seccomp personnalisé
docker run --security-opt seccomp=/opt/seccomp/monapp.json ...
```

Vérifier depuis l'hôte :

```bash
grep Seccomp /proc/$(docker inspect -f '{{.State.Pid}}' monapp-durci)/status
# Seccomp: 2  -> filtre actif
```

## 68. Dockerfile : construire des images sandbox-friendly

```dockerfile
# 1. Base minimale et épinglée par digest
FROM debian:bookworm-slim@sha256:<digest>

# 2. Un seul RUN pour limiter les couches, nettoyage à la fin
RUN apt-get update && apt-get install -y --no-install-recommends \
      ca-certificates tini \
 && rm -rf /var/lib/apt/lists/*

# 3. Utilisateur non-root DÉDIÉ (uid fixe pour les volumes)
RUN groupadd -r app --gid 10000 && useradd -r -g app --uid 10000 app
USER 10000:10000
WORKDIR /app

# 4. Copie avec bons propriétaires, pas de secrets dans l'image
COPY --chown=10000:10000 ./bin/monapp /app/monapp

# 5. init léger pour réaper les zombies (PID 1)
ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["/app/monapp"]

# 6. Healthcheck
HEALTHCHECK --interval=30s --timeout=3s CMD ["/app/monapp", "health"]
```

Interdits dans un Dockerfile sérieux :

- `ADD http://...` (utilisez `curl` + vérification de checksum)
- Secrets en `ENV` ou en couche (utilisez BuildKit `--secret`)
- `:latest` en production
- `RUN chmod 777` / `USER root` final

## 69. docker-compose durci : exemple complet

```yaml
# compose.durci.yaml
services:
  web:
    image: "monregistry/monapp:1.2.3@sha256:<digest>"
    read_only: true
    tmpfs:
      - /tmp:noexec,nosuid,size=100m
      - /run:noexec,nosuid,size=50m
    cap_drop:
      - ALL
    cap_add:
      - NET_BIND_SERVICE
    security_opt:
      - no-new-privileges:true
    user: "10000:10000"
    pids_limit: 100
    mem_limit: 512m
    memswap_limit: 512m
    cpus: 1.0
    networks:
      - front
    ports:
      - "8080:8080"   # n'exposez que le nécessaire
    volumes:
      - app-data:/var/lib/monapp:rw
      - ./config.yaml:/etc/monapp/config.yaml:ro
    restart: on-failure:3
    healthcheck:
      test: ["CMD", "/app/monapp", "health"]
      interval: 30s
      timeout: 3s
      retries: 3

networks:
  front:
    driver: bridge
    internal: false   # true si pas d'accès Internet nécessaire

volumes:
  app-data:
```

```bash
docker compose -f compose.durci.yaml up -d
docker compose -f compose.durci.yaml exec web id   # uid=10000
```

## 70. Rootless Docker / Podman : ne plus tourner en root

Le daemon Docker classique tourne en **root** : une faille du daemon = root sur l'hôte. Deux alternatives :

**Docker rootless :**

```bash
# Installation (utilisateur normal, pas de sudo)
dockerd-rootless-setuptool.sh install
systemctl --user enable --now docker

# Vérifier : le daemon tourne avec votre UID
ps aux | grep dockerd | grep -v grep
```

**Podman (recommandé sur serveurs, daemonless) :**

