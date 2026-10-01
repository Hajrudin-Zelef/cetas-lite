---
id: collect-261001-rattrapage/rattrapage/lxc-guide-5
title: "LXC & LXD — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "memory"]
source: docs/RAG/collect-261001-rattrapage/lxc_guide.md
source_anchor: ""
source_lines: [1104, 1396]
sha256: 5d4691c2ae206d9ce886a4844faa1434bc9fff29129588fc1e08dab97191d2dd
---

# LXC & LXD — Guide ultra-complet

Alternative COW quand ZFS n'est pas disponible (licence, noyau) :

```bash
lxc storage create btrfspool btrfs source=/dev/sdc
```

| ✅ | ❌ |
|---|---|
| Snapshots natifs, subvolumes | Moins mature que ZFS pour cet usage |
| Intégré au noyau | Fragmentation sur charges DB intenses |
| | Outils moins riches |

**Note :** btrfs convient bien pour des conteneurs légers et du dev.
Pour des bases de données en production, ZFS reste le choix sûr.

---

## 36. Stockage — Pool LVM

```bash
# Sur un groupe de volumes existant
sudo vgcreate vg-lxd /dev/sdd
lxc storage create lvmpool lvm source=vg-lxd
lxc storage set lvmpool lvm.thinpool_name=LXDThinPool
```

| ✅ | ❌ |
|---|---|
| Snapshots possibles (thin) | Snapshots moins efficaces que COW |
| Standard entreprise bien connu | Configuration plus manuelle |

**Usage :** parcs déjà standardisés sur LVM, ou contraintes de politique interne.

---

## 37. Volumes — Créer, attacher, détacher

```bash
# Créer un volume custom de 100 Go
lxc storage volume create default data-db --type=custom size=100GiB

# L'attacher à un conteneur (monté en /var/lib/postgresql)
lxc config device add c-db1 pgdata disk \
  pool=default \
  source=data-db \
  path=/var/lib/postgresql

# Détacher (le volume survit, les données sont conservées)
lxc config device remove c-db1 pgdata

# Le réattacher à un autre conteneur (migration de données !)
lxc config device add c-db2 pgdata disk \
  pool=default source=data-db path=/var/lib/postgresql

# Supprimer le volume (définitif)
lxc storage volume delete default data-db
```

**Bonnes pratiques :**

- Un volume par rôle de données (`pgdata`, `uploads`, `backups`).
- Nommez explicitement : `data-<service>-<env>`.
- Documentez les attachements (`lxc config device list`).

---

## 38. Snapshots — Créer, lister, restaurer

```bash
# Snapshot manuel
lxc snapshot c1
lxc snapshot c1 avant-maj           # avec un nom explicite

# Lister
lxc info c1 | grep -A20 Snapshots
lxc query /1.0/instances/c1/snapshots | jq

# Restaurer (état du snapshot ; le conteneur est arrêté brièvement)
lxc restore c1 avant-maj

# Restaurer SANS état (stateful=false par défaut ; les snapshots stateful
# capturent aussi la RAM — utile avant une manip risquée)
lxc snapshot c1 --stateful
lxc restore c1 snap-stateful

# Supprimer
lxc delete c1/avant-maj
```

> ⚠️ `lxc restore` **écrase** l'état actuel. Faites un snapshot "avant-restore"
> si le moindre doute subsiste.

---

## 39. Snapshots — Planification automatique

```bash
# Un snapshot par jour à 3h, conservés 7 jours
lxc config set c1 snapshots.schedule "0 3 * * *"
lxc config set c1 snapshots.schedule.stopped true   # même si arrêté
lxc config set c1 snapshots.expiry 7d
lxc config set c1 snapshots.pattern "auto-%Y%m%d-%H%M"

# Appliquer via un profil à tout un parc
lxc profile set backup-daily snapshots.schedule "0 3 * * *"
lxc profile set backup-daily snapshots.expiry 7d
```

**Stratégie type :**

| Fréquence | Rétention | Usage |
|---|---|---|
| `@hourly` | 24h | conteneurs critiques, pré-change |
| `@daily` | 7-14j | standard production |
| `@weekly` | 4-8 sem | conteneurs stables |

Les snapshots ne remplacent **pas** une sauvegarde externalisée (section 59) :
ils protègent des erreurs humaines, pas de la perte du disque hôte.

---

## 40. Publier un conteneur en image (golden image)

Transformez un conteneur configuré en **modèle réutilisable** :

```bash
# 1. Préparer le conteneur modèle (nettoyage)
lxc exec c-web-template -- apt clean
lxc exec c-web-template -- rm -rf /tmp/* /var/tmp/*
lxc exec c-web-template -- cloud-init clean 2>/dev/null

# 2. Arrêter
lxc stop c-web-template

# 3. Publier
lxc publish c-web-template --alias web-golden-v1

# 4. Utiliser
lxc launch web-golden-v1 web-prod-01
lxc launch web-golden-v1 web-prod-02

# Lister vos images publiées
lxc image list
```

**Workflow d'équipe :** `template` → `publish --alias web-golden-v<N>` →
déploiements. Versionnez les alias, documentez les changements dans la
description (`--prop description="..."`).

---

## 41. Limites de ressources — CPU

```bash
# Nombre de vCPU
lxc config set c1 limits.cpu 2

# Épinglage sur des cœurs précis (isolation NUMA / charges sensibles)
lxc config set c1 limits.cpu 0-3
lxc config set c1 limits.cpu 0,2,4,6

# Priorité relative (shares, défaut 1024)
lxc config set c1 limits.cpu.priority 512

# Temps CPU dur (microsecondes par période) — plafond strict
lxc config set c1 limits.cpu.allowance 50%
```

**Vérification dans le conteneur :**

```bash
lxc exec c1 -- nproc
lxc exec c1 -- cat /sys/fs/cgroup/cpu.max 2>/dev/null
```

> 📝 Avec cgroups v2, le conteneur voit ses limites via `/sys/fs/cgroup`.
> `nproc` reflète correctement les CPUs alloués sur les images récentes.

---

## 42. Limites de ressources — RAM

```bash
# Limite mémoire (dur)
lxc config set c1 limits.memory 4GiB

# Interdire le swap pour le conteneur (recommandé pour les DB)
lxc config set c1 limits.memory.swap false

# Pourcentage de la RAM hôte
lxc config set c1 limits.memory 25%
```

**Dimensionnement :** prévoyez toujours une marge hôte (~10-20 % pour le
système + ZFS ARC si applicable). La somme des limites peut dépasser la RAM
physique (overcommit), mais surveillez la pression mémoire réelle
(`lxc info c1` → Memory).

---

## 43. Limites de ressources — I/O disque

```bash
# Débit max (lecture/écriture)
lxc config set c1 limits.disk.priority 5
lxc config device set c1 root limits.read 100MB
lxc config device set c1 root limits.write 50MB

# IOPS max
lxc config device set c1 root limits.max 1000
```

**Cas d'usage :** empêcher un conteneur de monopoliser le disque (build qui
sature, logs qui s'emballent). Particulièrement utile sur pool partagé.

**Vérifier :**

```bash
lxc config device show c1 root
```

---

## 44. Limites — Processus, fichiers ouverts, mémoire verrouillée

```bash
# Nombre max de processus/threads
lxc config set c1 limits.processes 512

# Fichiers ouverts (nofile)
lxc config set c1 limits.kernel.nofile 65536

# Mémoire verrouillée (memlock, pour certaines DB)
lxc config set c1 limits.kernel.memlock unlimited

# Nombre max de conteneurs imbriqués / clés
lxc config set c1 limits.kernel.nproc 4096
```

**Exemple concret — PostgreSQL :**

```bash
lxc config set c-db1 limits.kernel.nofile 100000
lxc config set c-db1 limits.memory 8GiB
lxc config set c-db1 limits.memory.swap false
```

---

## 45. Limites — Ce que voit le conteneur (cgroups v2)

Sur les systèmes récents, les limites sont **visibles et opposables** dans
le conteneur :

```bash
lxc exec c1 -- cat /sys/fs/cgroup/memory.max
lxc exec c1 -- cat /sys/fs/cgroup/cpu.max
lxc exec c1 -- cat /sys/fs/cgroup/pids.max
```

**Pourquoi c'est important :** la JVM, Go, ou `systemd` adaptent leur
comportement (heap, threads) à la mémoire/CPU **vue**. Des limites correctes
= des applis qui s'auto-dimensionnent bien.

**Piège :** une image très ancienne peut ne pas exposer `/sys/fs/cgroup`
correctement → préférez des images récentes (Ubuntu 22.04+/Debian 12+).

---

## 46. Conteneurs privilégiés vs non privilégiés — Sécurité

|  | Non privilégié (défaut LXD) | Privilégié |
|---|---|---|
| UID 0 du conteneur | Mappé sur un UID non-root de l'hôte (ex. 1000000) | = UID 0 réel de l'hôte |
| Évasion → impact | Limitée par le user namespace | **Root complet sur l'hôte** |
| Cas d'usage | 99 % des cas | Pilotes/noyau spécifiques, rarement justifié |

```bash
# Vérifier le mode
lxc config get c1 security.privileged   # vide/false = non privilégié

# (DÉCONSEILLÉ) passer en privilégié
lxc config set c1 security.privileged true
```

**Mapping des UIDs (idmap) :**

```bash
# Voir la plage allouée à un conteneur
lxc config get c1 volatile.idmap.next
cat /etc/subuid /etc/subgid   # plages de l'hôte
```

