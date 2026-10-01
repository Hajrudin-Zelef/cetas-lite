---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-4
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox", "arr", "exploit", "memory"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [609, 844]
sha256: 29f5233e557cc8d4085e7811b9ff3e513ace4c694d8e75fe6db66e7bc9eee71e
---

# Guide complet du sandboxing sous Linux

| Profil | Programme | Protections notables |
|---|---|---|
| `firefox.profile` | Firefox | `private-bin`, `seccomp`, `noroot`, netfilter |
| `chromium.profile` | Chromium | idem + `disable-mnt` |
| `thunderbird.profile` | Thunderbird | `private-tmp`, restriction `~/` |
| `evince.profile` | Lecteur PDF GNOME | `private-bin`, pas de réseau (`net none`) |
| `okular.profile` | Lecteur PDF KDE | idem |
| `vlc.profile` | VLC | `nosound` optionnel, `seccomp` |
| `libreoffice.profile` | LibreOffice | `private-tmp`, `dbus-user none` |
| `wget.profile` / `curl.profile` | Téléchargements | `private-bin`, écriture limitée |

Consulter un profil :

```bash
cat /etc/firejail/firefox.profile
```

## 25. Utilisation quotidienne : les commandes essentielles

```bash
# Lancer avec le profil par défaut (détection automatique)
firejail firefox

# Lancer SANS profil (sandbox générique)
firejail --noprofile evince doc.pdf

# Lancer avec un profil personnalisé
firejail --profile=/home/zelef/.config/firejail/navigateur.profile firefox

# Confinement maximal ad hoc : pas de réseau, /home masqué
firejail --noprofile --net=none --private=/tmp/test --seccomp \
  --caps.drop=all --noroot evince suspect.pdf

# Lister les sandboxes actifs
firejail --list
# 1234:toto:firejail firefox:/usr/bin/firefox

# Voir l'arbre des processus d'un sandbox
firejail --tree

# Arrêter un sandbox
firejail --shutdown=1234

# Top des sandboxes (ressources)
firejail --top
```

Intégration au bureau : Firejail peut remplacer les lanceurs via `firecfg` :

```bash
# Crée des liens symboliques /usr/local/bin/firefox -> /usr/bin/firejail
sudo firecfg
# Désormais, cliquer sur l'icône Firefox lance firejail automatiquement
```

Vérifiez ensuite : `which firefox` doit pointer vers `/usr/local/bin/firefox`.

## 26. Anatomie d'un profil Firejail

Un profil est un fichier texte avec une directive par ligne. Exemple commenté :

```ini
# ~/.config/firejail/navigateur.profile
include firefox.local

# --- Filesystem ---
private-dev                 # /dev minimal (pas d'accès disques bruts)
private-tmp                 # /tmp isolé par namespace
private-bin firefox,sh,bash # seuls ces binaires visibles dans PATH
private-etc passwd,group,resolv.conf,hosts
whitelist ${HOME}/Téléchargements
whitelist ${HOME}/.mozilla
# tout le reste de ${HOME} est masqué (blacklist implicite via whitelist)

# --- Réseau ---
netfilter                   # applique les règles iptables du profil
# net none                  # décommentez pour couper tout réseau

# --- Noyau / syscalls ---
seccomp                     # filtre seccomp par défaut
caps.drop all               # retire toutes les capabilities
noroot                      # même root dedans = nobody dehors
nonewprivs                  # NoNewPrivileges (bloque setuid/sudo)
protocol unix,inet,inet6    # familles de sockets autorisées

# --- Divers ---
nodbus                      # coupe D-Bus (ou dbus-user none)
notv                        # pas d'accès TV/DVB
nosound                     # pas d'audio (à adapter)
name navigateur-durci
hostname sandbox-navigateur
```

Ordre de priorité : la ligne de commande écrase le profil utilisateur (`~/.config/firejail/`), qui écrase le profil système (`/etc/firejail/`).

## 27. Créer un profil : navigateur durci pas à pas

Objectif : un profil Firefox pour la navigation à risque (liens reçus par mail, sites douteux).

```bash
mkdir -p ~/.config/firejail
cp /etc/firejail/firefox.profile ~/.config/firejail/firefox.profile
```

Ajoutez à la fin de `~/.config/firejail/firefox.profile` :

```ini
# --- Durcissement navigation à risque ---
private-cache               # cache isolé, jeté à la fermeture
#whitelist ${HOME}/Téléchargements
# -> commentez whitelist pour interdire tout téléchargement persistant
read-only ${HOME}
tmpfs ${HOME}/.cache
noexec /tmp
noexec ${HOME}
seccomp.block-secondary     # bloque les ABI 32 bits (réduit la surface)
restrict-namespaces         # interdit de créer de nouveaux namespaces
memory-deny-write-execute   # bloque les pages W+X (anti-exploit JIT limité)
```

Testez :

```bash
firejail --profile=~/.config/firejail/firefox.profile firefox
# Dans Firefox : about:support -> vérifier le fonctionnement
# Tenter un téléchargement -> doit échouer ou aller en RAM
```

**Limite connue :** `memory-deny-write-execute` peut casser le JIT JavaScript de Firefox sur certains sites (baisse de perf ou plantage d'onglet). Si c'est gênant, retirez la ligne pour la navigation courante et gardez-la pour le profil "à risque".

## 28. Créer un profil : lecteur PDF pour documents suspects

Scénario : ouvrir un PDF d'origine inconnue. Le lecteur n'a besoin ni de réseau, ni de vos fichiers.

```ini
# ~/.config/firejail/pdf-suspect.profile
# Usage : firejail --profile=pdf-suspect.profile evince /tmp/suspect.pdf

include disable-common.inc
include disable-devel.inc
include disable-programs.inc

caps.drop all
net none                    # AUCUN réseau : un PDF n'en a pas besoin
nonewprivs
noroot
seccomp
protocol unix               # sockets unix locales uniquement (X11/Wayland)
private-dev
private-tmp
private-bin evince,sh
private-etc passwd,group,fonts
dbus-user none
dbus-system none

# Le PDF est copié dans un répertoire dédié, seul visible
whitelist /tmp/pdfs-suspects
read-only /tmp/pdfs-suspects
```

Procédure d'ouverture d'un PDF suspect :

```bash
mkdir -p /tmp/pdfs-suspects
cp ~/Téléchargements/suspect.pdf /tmp/pdfs-suspects/
firejail --profile=~/.config/firejail/pdf-suspect.profile \
  evince /tmp/pdfs-suspects/suspect.pdf
# Après lecture : rm /tmp/pdfs-suspects/suspect.pdf
```

Même si le PDF exploite Evince, l'attaquant n'a ni réseau pour exfiltrer, ni accès à vos fichiers, ni `ptrace`, ni `mount`.

## 29. Créer un profil : Thunderbird

```ini
# ~/.config/firejail/thunderbird-durci.profile
include thunderbird.profile

# Durcissement supplémentaire
seccomp.block-secondary
restrict-namespaces
# Thunderbird a besoin du réseau (IMAP/SMTP) : on ne coupe pas net,
# mais on restreint les protocoles
protocol unix,inet,inet6,netlink
private-cache
```

Pièces jointes : configurez Thunderbird pour enregistrer les pièces jointes dans `/tmp/pdfs-suspects` (ou `~/Téléchargements`), puis ouvrez-les avec le profil PDF ci-dessus. Ne jamais ouvrir une pièce jointe directement depuis le client mail sans sandbox.

## 30. Application non fiable : patron générique

Pour n'importe quel binaire douteux, ce patron couvre 90 % des besoins :

```bash
firejail --noprofile \
  --private \                    # home temporaire vide
  --net=none \                   # pas de réseau (retirez si nécessaire)
  --seccomp \                    # filtre appels système
  --caps.drop=all \              # aucune capability
  --noroot \                     # pas de root même factice
  --nonewprivs \                 # bloque setuid
  --private-dev \                # /dev minimal
  --private-tmp \                # /tmp isolé
  --read-only=/ \                # système en lecture seule...
  --tmpfs=/tmp --tmpfs=/var/tmp \ # ...sauf tmpfs explicites
  --hostname=sandbox \
  /chemin/vers/binaire --args
```

Variante avec accès réseau contrôlé mais sans accès à votre home :

```bash
firejail --noprofile \
  --private=/tmp/faux-home \
  --dns=9.9.9.9 \
  --seccomp --caps.drop=all --noroot \
  /chemin/vers/binaire
```

## 31. Firejail et X11 : le point faible à connaître

Sous X11, **tout client X peut espionner le clavier et capturer l'écran** (keylogging, screenshots). Firejail propose :

```ini
x11 xephyr     # serveur X imbriqué : l'appli ne voit que sa fenêtre
# ou
x11 xpra       # similaire, avec persistance
# ou (radical)
x11 none       # pas d'accès X du tout (applis console uniquement)
```

```bash
sudo apt install -y xephyr
firejail --x11=xephyr firefox
```

