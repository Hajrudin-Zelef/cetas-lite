---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-7
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: ["sandbox"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [1324, 1575]
sha256: b0e7ad3ba8856812628c7496061c9977ce8f5ba89ca9f2e742b271d62a363f9c
---

# Guide complet du sandboxing sous Linux

```bash
systemctl --user daemon-reload
systemctl --user enable --now syncthing-perso.service
systemd-analyze security --user syncthing-perso.service
```

---

# PARTIE V — APPARMOR : CONFINEMENT PAR PROFIL (DEBIAN/UBUNTU)

## 51. AppArmor : principe

AppArmor est un système de **contrôle d'accès obligatoire (MAC)** : chaque programme confiné suit un profil qui liste ce qu'il a le droit de faire (fichiers, réseau, capabilities). Tout le reste est refusé — même si les permissions Unix l'autoriseraient. C'est le MAC par défaut sur Debian/Ubuntu (contrairement à SELinux, défaut sur RHEL).

```bash
# Vérifier qu'AppArmor est actif
sudo aa-status
# apparmor module is loaded.
# 45 profiles are loaded.

# Profils chargés et leur mode
sudo aa-status | head -30
```

Deux modes :

| Mode | Comportement |
|---|---|
| `enforce` | Les violations sont **bloquées** et journalisées |
| `complain` | Les violations sont **journalisées uniquement** (apprentissage) |

## 52. Lire un profil AppArmor

Les profils sont dans `/etc/apparmor.d/` :

```bash
ls /etc/apparmor.d/
cat /etc/apparmor.d/usr.sbin.cupsd | head -60
```

Anatomie :

```apparmor
# /etc/apparmor.d/usr.bin.monapp (exemple simplifié)
#include <tunables/global>

/usr/bin/monapp {
  #include <abstractions/base>      # règles de base (libc, /proc, etc.)
  #include <abstractions/nameservice> # DNS : /etc/resolv.conf, nss

  capability net_bind_service,       # capabilities autorisées
  capability sys_nice,

  network inet tcp,                  # réseau autorisé
  network inet6 tcp,

  /usr/bin/monapp mr,                # m = mmap exécutable, r = lecture
  /etc/monapp/** r,                  # lecture config
  /var/lib/monapp/** rwk,            # rwx + k (verrouillage)
  /var/log/monapp/** w,              # écriture logs
  /tmp/** rw,

  deny /home/** r,                   # refus explicite (loggé)
}
```

Syntaxe des droits fichiers : `r` (lecture), `w` (écriture), `x` (exécution), `m` (mmap PROT_EXEC), `k` (lock), `l` (lien). Les règles s'accumulent (union).

## 53. Passer un profil en complain / enforce

```bash
# Passer en mode apprentissage (ne bloque plus, loggue)
sudo aa-complain /usr/bin/monapp

# Revenir en enforce
sudo aa-enforce /usr/bin/monapp

# Désactiver totalement un profil
sudo aa-disable /usr/bin/monapp

# Recharger tous les profils
sudo systemctl reload apparmor
```

Cycle de travail recommandé pour un nouveau profil : `complain` → exercer le programme normalement → générer le profil → `enforce` → surveiller les logs.

## 54. aa-genprof : générer un profil par apprentissage

```bash
# 1. Lancer l'assistant (terminal 1)
sudo aa-genprof /usr/bin/monapp

# 2. Dans un autre terminal, utilisez le programme normalement :
#    ouvrez des fichiers, cliquez partout, testez toutes les fonctions
monapp

# 3. Revenez au terminal 1, appuyez sur "S" (Scan) :
#    aa-genprof propose chaque accès détecté : (A)llow, (D)eny, (G)lob...
# 4. "F" (Finish) quand vous avez tout couvert, puis "S" (Save)
```

Bonnes pratiques pendant l'apprentissage :

- Exercez **tous** les chemins : ouverture, sauvegarde, impression, préférences, mise à jour.
- Faites-le sur une machine de test, pas en production.
- Relisez le profil généré avant de l'enforcer : `aa-genprof` a tendance à autoriser large (`/tmp/** rw`).

## 55. aa-logprof : affiner depuis les logs

Quand un profil `enforce` bloque quelque chose de légitime, les refus sont dans les logs :

```bash
# Voir les refus AppArmor
sudo dmesg | grep -i apparmor | grep DENIED
sudo journalctl -k | grep "apparmor.*DENIED"
# Exemple :
# audit(1234.567:89): apparmor="DENIED" operation="open"
#   profile="/usr/bin/monapp" name="/etc/monapp/extra.conf" comm="monapp"
```

```bash
# Assistant interactif : propose d'ajouter les accès refusés au profil
sudo aa-logprof
```

## 56. Exemple : profil AppArmor pour un lecteur PDF

```apparmor
# /etc/apparmor.d/usr.bin.evince-sandbox
#include <tunables/global>

/usr/bin/evince {
  #include <abstractions/base>
  #include <abstractions/fonts>
  #include <abstractions/gnome>

  capability dac_override,

  network none,          # aucun réseau (redondant avec Firejail, défense en profondeur)

  /usr/bin/evince mr,
  /usr/share/evince/** r,
  /usr/share/poppler/** r,
  /etc/fonts/** r,
  /usr/share/fonts/** r,

  # Le répertoire des PDF suspects : lecture seule
  /tmp/pdfs-suspects/** r,

  # Cache et config GNOME minimaux
  owner @{HOME}/.cache/evince/** rw,
  owner @{HOME}/.config/evince/** rw,

  # Tout le reste du home : refusé
  deny owner @{HOME}/** rwklx,
}
```

```bash
sudo apparmor_parser -r /etc/apparmor.d/usr.bin.evince-sandbox
sudo aa-enforce /usr/bin/evince
```

## 57. AppArmor et systemd : combiner les deux

Les directives systemd (`ProtectSystem`, etc.) et AppArmor se cumulent : si l'un bloque, l'accès est refusé. Pour un service critique, utilisez les deux :

```ini
[Service]
# systemd confine les namespaces/syscalls...
ProtectSystem=strict
ProtectHome=tmpfs
NoNewPrivileges=true
SystemCallFilter=@system-service
# ...et AppArmor confine les chemins fins
AppArmorProfile=monapp-confine
```

```bash
# Vérifier que le profil AppArmor est bien appliqué au service
sudo aa-status | grep monapp
cat /proc/$(pidof monapp)/attr/current
# -> monapp-confine (enforce)
```

## 58. Dépannage AppArmor : cas concrets

**Cas 1 — Le programme ne démarre plus après `aa-enforce`, aucun message clair**
Regardez les DENIED :

```bash
sudo journalctl -k --since "5 min ago" | grep DENIED
```

Souvent : une bibliothèque dans un chemin non prévu (`/opt/...`), ou une abstraction manquante (`#include <abstractions/X>`).

**Cas 2 — `aa-genprof` ne détecte rien**
Le programme tourne peut-être déjà sous un autre profil, ou les logs audit sont désactivés. Vérifiez `sudo aa-status` et relancez le programme après `aa-complain`.

**Cas 3 — Conflit avec un profil existant du paquet**
Debian livre des profils pour cupsd, dhclient, etc. Si vous écrivez le vôtre pour le même binaire, le paquet peut l'écraser à la mise à jour. Placez vos surcharges dans `/etc/apparmor.d/local/usr.bin.monapp` (inclus via `#include <local/...>` si le profil le prévoit).

**Cas 4 — Performances**
AppArmor a un coût négligeable (< 1 %). Si un doute : `sudo sysctl -w kernel.apparmor_restrict_unprivileged_userns=1` n'a rien à voir avec les perfs, c'est un durcissement.

---

# PARTIE VI — SELINUX : TOUR D'HORIZON (CULTURE GÉNÉRALE)

## 59. SELinux : notions essentielles

SELinux est le MAC historique (NSA, défaut sur RHEL/Fedora). Plus puissant et plus complexe qu'AppArmor : il étiquette **chaque objet** (fichier, processus, port) avec un **contexte**, et n'autorise que les interactions prévues par la politique.

```bash
# Sur un système RHEL/Fedora :
getenforce        # Enforcing | Permissive | Disabled
sestatus
```

Modes :

| Mode | Effet |
|---|---|
| Enforcing | Bloque + journalise |
| Permissive | Journalise seulement (apprentissage) |
| Disabled | Désactivé (nécessite reboot) |

## 60. Contextes SELinux : lire les étiquettes

Format : `utilisateur:rôle:type:niveau` — ex. `system_u:system_r:httpd_t:s0`.

```bash
# Voir le contexte des fichiers
ls -Z /var/www/html/
# system_u:object_r:httpd_sys_content_t:s0 index.html

# Voir le contexte des processus
ps -eZ | grep httpd
# system_u:system_r:httpd_t:s0 1234 ? /usr/sbin/httpd

# Voir le contexte des ports
semanage port -l | grep http
```

Règle d'or : un processus `httpd_t` ne peut lire que les fichiers étiquetés `httpd_sys_content_t`. Un fichier créé ailleurs garde son ancien type → accès refusé même avec `chmod 777`. D'où 90 % des "bugs" SELinux : **mauvais contexte après déplacement de fichiers**.

## 61. SELinux au quotidien : les 5 commandes qui sauvent

```bash
# 1. Le fichier a le mauvais type ? Restaurez le type par défaut du chemin
restorecon -Rv /var/www/html/

