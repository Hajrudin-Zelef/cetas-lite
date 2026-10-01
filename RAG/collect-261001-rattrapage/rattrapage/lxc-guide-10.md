---
id: collect-261001-rattrapage/rattrapage/lxc-guide-10
title: "LXC & LXD — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "memory"]
source: docs/RAG/collect-261001-rattrapage/lxc_guide.md
source_anchor: ""
source_lines: [2444, 2685]
sha256: f1df620b8351213013516b820cd24c3c04e78c58b7001f637360c45db438a667
---

# LXC & LXD — Guide ultra-complet

```bash
# Vérifier dnsmasq
lxc network show lxdbr0 | grep ipv4.dhcp
# Voir les baux
sudo cat /var/snap/lxd/common/lxd/networks/lxdbr0/dnsmasq.leases
# Redémarrer le réseau LXD
lxc network set lxdbr0 ipv4.dhcp false && lxc network set lxdbr0 ipv4.dhcp true
lxc restart c1
```

### Erreur 3 : `Failed to run: forklimits: ...` / limites non appliquées

**Cause :** clé mal orthographiée (`limits.memroy`) — LXD ne valide pas tout.
**Solution :** `lxc config show c1 --expanded` et relire chaque clé.

### Erreur 4 : `lxc exec` se fige / pas de sortie

**Cause :** pas de TTY dans un script.
**Solution :** `lxc exec c1 --mode non-interactive -- <cmd>`.

### Erreur 5 : Plus de place sur le pool (`no space left on device`)

```bash
# Diagnostic
lxc storage info default
sudo zfs list -o name,used,avail | head -20
# Nettoyage : vieux snapshots, images inutilisées
lxc image list   # repérer les inutilisées
lxc query /1.0/instances/c1/snapshots -X DELETE  # au cas par cas
```

### Erreur 6 : Conteneur en macvlan injoignable depuis l'hôte

**Normal !** Limitation du noyau (section 26). Solutions : ajouter une
interface sur le bridge, ou passer en `bridged` + proxy.

### Erreur 7 : `security.nesting` oublié → Docker ne démarre pas dans LXC

Symptôme : `docker run` échoue avec des erreurs de montage/cgroups.
**Solution :** section 51-52 (`security.nesting=true` + interceptions).

### Erreur 8 : Droits incohérents sur un dossier partagé hôte ↔ conteneur

**Cause :** oubli du décalage UID (section 49).
**Solution :** `chown 1000000:1000000` côté hôte, ou `shift=true` sur le device.

### Erreur 9 : Snapshot `restore` qui écrase le travail en cours

**Réflexe :** toujours `lxc snapshot c1 avant-restore` avant un restore.
Le restore est destructif pour l'état actuel.

### Erreur 10 : Mise à jour snap qui redémarre les conteneurs

Le refresh snap de LXD peut impacter le daemon. **Solution :**

```bash
sudo snap set lxd refresh.hold="$(date -d '+7 days' +%Y-%m-%dT%H:%M:%S%:z)"
# Planifier les refreshs en fenêtre de maintenance
sudo snap refresh lxd   # manuel, en heure creuse
```

### Erreur 11 : `lxc copy` inter-hôtes très lent

**Cause :** copie sans optimisation entre backends différents.
**Solution :** `--optimized-storage`, ou mieux : même backend ZFS des deux
côtés + `lxc copy --refresh` pour les deltas.

### Erreur 12 : Oubli du `boot.autostart` → services morts après reboot hôte

```bash
# Appliquer à tout le parc prod d'un coup
for ct in web1 web2 pg1 rproxy; do
  lxc config set $ct boot.autostart true
  lxc config set $ct boot.autostart.priority 10
done
```

---

## 81. Dépannage — Méthode en 7 étapes

```
1. DÉFINIR    → Quel est le symptôme exact ? Depuis quand ? Qu'est-ce qui a changé ?
2. PÉRIMÈTRE  → Un conteneur ou tous ? Réseau ? Stockage ? Daemon ?
3. OBSERVER   → lxc list, lxc info, lxc monitor, journalctl -u snap.lxd.daemon
4. ISOLER     → Reproduire sur un conteneur test minimal
5. HYPOTHÈSE  → Une cause à la fois, la plus simple d'abord
6. CORRIGER   → Snapshot avant toute manip destructive !
7. DOCUMENTER → Noter la cause et la solution (ce guide, section 80)
```

**Commandes de premier réflexe :**

```bash
lxc list                              # état global
lxc info c1 --show-log                # log console du conteneur
lxc monitor --type=logging --pretty   # logs temps réel du daemon
sudo journalctl -u snap.lxd.daemon -n 100   # daemon (snap)
sudo journalctl -u lxd -n 100              # daemon (apt)
lxc network show lxdbr0               # état réseau
lxc storage info default              # état stockage
```

---

## 82. Cas de dépannage concrets

### Cas A : Un conteneur ne démarre plus après MAJ de l'hôte

```bash
# 1. Message d'erreur exact
lxc start c1
# Error: Failed to run: /snap/lxd/.../lxd forklimits ...

# 2. Souvent : module noyau ou AppArmor après reboot
sudo dmesg | grep -i apparmor | tail -20
sudo systemctl restart snap.lxd.daemon

# 3. Si le rootfs semble corrompu : monter le volume à la main (ZFS)
sudo zfs list | grep c1
sudo mkdir -p /mnt/rescue && sudo mount -t zfs default/containers/c1 /mnt/rescue
# → récupérer les données, puis recréer le conteneur
```

### Cas B : Lenteurs disque soudaines sur tous les conteneurs

```bash
# 1. Le pool est-il plein à > 80 % ? (ZFS s'effondre quand il est plein)
sudo zpool list -o name,size,alloc,free,cap

# 2. Scrub en cours ?
sudo zpool status | grep -A2 scrub

# 3. Un conteneur sature l'I/O ?
lxc info c1 | grep -A5 Disk
# → appliquer limits.read/write (section 43) au fautif

# 4. Santé des disques
sudo smartctl -a /dev/sdb | grep -i "reallocated\|pending"
```

### Cas C : Perte réseau d'un seul conteneur

```bash
# 1. Le veth existe-t-il côté hôte ?
ip link show | grep -i c1

# 2. Le conteneur voit-il son interface ?
lxc exec c1 -- ip addr show eth0

# 3. DHCP : le conteneur demande-t-il ?
lxc exec c1 -- dhclient -v eth0   # ou systemctl restart systemd-networkd

# 4. Règles firewall de l'hôte qui bloquent le forward ?
sudo nft list ruleset | grep -i forward
sudo iptables -L FORWARD -n -v | head

# 5. Solution rapide : recréer le device réseau
lxc config device remove c1 eth0
lxc config device add c1 eth0 nic network=lxdbr0 name=eth0
lxc restart c1
```

### Cas D : `lxc exec` impossible (conteneur gelé)

```bash
# Le daemon répond-il encore ?
lxc list --fast

# Forcer l'arrêt puis redémarrer
lxc stop c1 --force
lxc start c1

# Si le daemon lui-même est bloqué :
sudo systemctl restart snap.lxd.daemon
# Les conteneurs SURVIVENT au redémarrage du daemon (ils sont indépendants)
```

> 💡 Point clé : **le daemon LXD n'est pas dans le chemin d'exécution** des
> conteneurs. Le redémarrer ne tue pas les conteneurs — rassurant en prod.

---

## 83. Pense-bête de poche — Cheat sheet

```
=== CYCLE DE VIE ===
lxc launch ubuntu:24.04 c1            créer + démarrer
lxc init ubuntu:24.04 c1              créer sans démarrer
lxc start|stop|restart|freeze c1      contrôler
lxc delete c1 [--force]               supprimer
lxc list [-c n,s,4]                   lister

=== ACCÈS ===
lxc exec c1 -- bash                   shell root
lxc exec c1 --user 1000 -- whoami     en tant qu'autre user
lxc shell c1                          shell rapide
lxc file push f c1/path               envoyer fichier
lxc file pull c1/path f               récupérer fichier
lxc file edit c1/etc/hosts            éditer

=== CONFIG ===
lxc config show c1 --expanded         config complète
lxc config set c1 limits.memory 4GiB  définir
lxc config unset c1 limits.memory     hériter du profil
lxc config edit c1                    éditer YAML

=== PROFILS ===
lxc profile create|edit|list|show p   gérer
lxc profile add c1 p                  appliquer
lxc launch img c1 -p p1 -p p2         appliquer au lancement

=== RÉSEAU ===
lxc network list|show lxdbr0          réseaux
lxc network create dmz0 ipv4.address=192.168.50.1/24 ipv4.nat=false
lxc config device set c1 eth0 ipv4.address 10.10.10.50
lxc config device add c1 http proxy listen=tcp:0.0.0.0:8080 connect=tcp:127.0.0.1:80

=== STOCKAGE ===
lxc storage list|info default         pools
lxc storage volume create default v1 size=50GiB
lxc config device add c1 d disk pool=default source=v1 path=/data

=== SNAPSHOTS ===
lxc snapshot c1 [nom]                 créer
lxc restore c1 nom                    restaurer (destructif !)
lxc delete c1/nom                     supprimer
lxc publish c1 --alias golden-v1      publier en image

=== LIMITES ===
lxc config set c1 limits.cpu 2
lxc config set c1 limits.memory 4GiB
lxc config set c1 limits.processes 512

=== SÉCURITÉ ===
lxc config set c1 security.nesting true      # Docker dans LXC
lxc config get c1 security.privileged        # doit être false

=== BACKUP ===
lxc export c1 /backups/c1.tar.gz
lxc import /backups/c1.tar.gz [nouveau-nom]

