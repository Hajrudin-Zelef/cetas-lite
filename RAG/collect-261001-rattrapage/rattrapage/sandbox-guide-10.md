---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-10
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [1966, 2161]
sha256: 62e4876d8271bddef5a89f764c839b3576505bc26fb71fd9087c526a3a35dc78
---

# Guide complet du sandboxing sous Linux

- **Firejail** : `sudo firecfg` après chaque mise à jour (régénère les liens).
- **Profils AppArmor** : `aa-logprof` après mise à jour majeure d'une appli confinée.
- **Images Docker** : rebuild régulier, scan Trivy en CI, politique "pas d'image > 90 jours".
- **Chroots de build** : reconstruction mensuelle depuis `debootstrap` frais.
- **Profils Firejail personnalisés** : revue trimestrielle (un `private-bin` peut manquer un binaire ajouté par une mise à jour).

---

# PARTIE IX — CAS D'USAGE COMMENTÉS

## 79. Cas 1 : ouvrir un PDF suspect reçu par mail (procédure complète)

```bash
#!/bin/bash
# pdf-suspect.sh — procédure d'ouverture sécurisée d'un PDF douteux
set -euo pipefail
PDF="$1"
QUARANTAINE=/tmp/pdfs-suspects
mkdir -p "$QUARANTAINE"

# 1. Copie en quarantaine (on ne touche pas à l'original)
cp -- "$PDF" "$QUARANTAINE/in.pdf"
chmod 400 "$QUARANTAINE/in.pdf"

# 2. Analyse statique rapide (sans l'ouvrir)
pdfinfo "$QUARANTAINE/in.pdf" | head -20        # métadonnées
strings "$QUARANTAINE/in.pdf" | grep -iE "javascript|/JS|/Launch|/EmbeddedFile" | head
# Si /JS ou /Launch présent : méfiance renforcée

# 3. Ouverture sandboxée : pas de réseau, home masqué, seccomp
firejail --profile="$HOME/.config/firejail/pdf-suspect.profile" \
  evince "$QUARANTAINE/in.pdf"

# 4. Nettoyage
shred -u "$QUARANTAINE/in.pdf"
echo "Terminé. Si le document est légitime, demandez une version saine à l'expéditeur."
```

Alternative sans Firejail (pur systemd-run) :

```bash
systemd-run --user --scope -p PrivateTmp=yes -p ProtectSystem=strict \
  -p ProtectHome=tmpfs -p SystemCallFilter=@system-service \
  -p RestrictAddressFamilies=AF_UNIX -p NoNewPrivileges=yes \
  /usr/bin/evince /tmp/pdfs-suspects/in.pdf
```

## 80. Cas 2 : navigateur durci pour sites à risque

Procédure pour l'équipe : un profil Firefox "quarantaine" lancé via un script :

```bash
#!/bin/bash
# navigateur-quarantaine.sh
# Usage : navigateur-quarantaine.sh https://site-douteux.example/
URL="${1:-about:blank}"
firejail \
  --profile="$HOME/.config/firejail/firefox.profile" \
  --private-cache \
  --name=quarantaine \
  firefox --no-remote --new-instance "$URL"
```

Règles d'usage à afficher aux utilisateurs :

1. Jamais de mot de passe / session bancaire dans le navigateur quarantaine.
2. Téléchargements désactivés (ou vers tmpfs jeté à la fermeture).
3. Fermer le sandbox après usage (`firejail --shutdown=quarantaine`).
4. En cas de comportement bizarre (popups, ralentissement) : tuer le sandbox, signaler au support.

## 81. Cas 3 : tester un script d'installation douteux (`curl | bash`)

Le classique `curl https://site/install.sh | sudo bash` est une roulette russe. Procédure :

```bash
# 1. Télécharger SANS exécuter
curl -fsSL https://site.example/install.sh -o /tmp/install.sh
chmod 400 /tmp/install.sh

# 2. Lire le script (au moins en diagonale : curl, wget, chmod +x, /etc, systemd ?)
less /tmp/install.sh
grep -nE "curl|wget|chmod \+x|/etc/(cron|systemd)|useradd|passwd|ssh" /tmp/install.sh

# 3. Exécution à blanc dans un conteneur jetable (jamais sur l'hôte)
docker run --rm -it --read-only --cap-drop=ALL --network none \
  --user 10000:10000 -v /tmp/install.sh:/tmp/install.sh:ro \
  debian:bookworm-slim bash -c 'bash -n /tmp/install.sh && echo "syntaxe OK"'
# bash -n = parse seul, n'exécute rien

# 4. Si le script semble sain ET nécessaire : exécution supervisée
#    dans une VM de test d'abord, avec snapshot avant/après (diff du filesystem)
```

**Règle d'équipe :** aucun `curl|bash` en production sans relecture par un pair et test en VM.

## 82. Cas 4 : pièce jointe Thunderbird → extraction sandboxée

```bash
#!/bin/bash
# pj-sandbox.sh — ouvre une pièce jointe dans le bon sandbox selon son type
set -euo pipefail
PJ="$1"
MIME=$(file -b --mime-type "$PJ")

case "$MIME" in
  application/pdf)
    firejail --profile="$HOME/.config/firejail/pdf-suspect.profile" evince "$PJ" ;;
  application/zip|application/x-tar)
    # Lister d'abord, sans extraire
    firejail --noprofile --net=none --private-tmp 7z l "$PJ" | less ;;
  application/vnd.openxmlformats*|application/msword)
    firejail --profile=/etc/firejail/libreoffice.profile libreoffice "$PJ" ;;
  *)
    echo "Type $MIME : ouverture refusée par politique. Demandez au support." ;;
esac
```

## 83. Cas 5 : compilation d'un projet GitHub inconnu (récapitulatif outillé)

Combine les §17 et §72 : pipeline "clone → build → artefact" jamais exécuté sur l'hôte.

```bash
#!/bin/bash
# build-tiers.sh <url-git> <branche>
set -euo pipefail
URL="$1"; BRANCHE="${2:-main}"
WORK=$(mktemp -d); trap 'rm -rf "$WORK"' EXIT

# 1. Clone SANS exécuter de hooks
git clone --depth 1 --branch "$BRANCHE" --config core.hooksPath=/dev/null "$URL" "$WORK/src"

# 2. Inspection : que va exécuter le build ?
grep -rE "curl|wget|pip install|npm install" "$WORK/src/Makefile" "$WORK/src/configure" 2>/dev/null | head

# 3. Build dans nspawn sans réseau (voir §17)
sudo systemd-nspawn -D /srv/chroot/build --private-users=pick \
  --private-network --bind-ro="$WORK/src:/src" --tmpfs=/build --as-pid2 \
  bash -c 'cp -a /src/. /build/ && cd /build && (./configure && make -j$(nproc))'

# 4. Les binaires restent suspects : testez-les DANS le sandbox avant déploiement
echo "Artefacts dans $WORK (non nettoyé pour inspection)"
```

---

# PARTIE X — SUPERVISION ET LOGS D'AUDIT

## 84. Centraliser les signaux de sandboxing

Sources à collecter :

| Source | Ce qu'elle révèle | Commande / chemin |
|---|---|---|
| Journal noyau (audit) | Refus AppArmor, kills seccomp | `journalctl -k \| grep -E "apparmor.*DENIED\|seccomp"` |
| `auditd` | AVC SELinux, appels surveillés | `/var/log/audit/audit.log` |
| Journal systemd | Services tués (SIGSYS=31) | `journalctl -u <service>` + `status=31/SYS` |
| Docker events | OOM, die, restart loops | `docker events` |
| Firejail | `--audit`, `--tree` | `firejail --list` |

```bash
# Installer auditd pour une traçabilité fine (serveurs sensibles)
sudo apt install -y auditd audispd-plugins
sudo systemctl enable --now auditd

# Surveiller les exécutions de binaires sensibles
sudo auditctl -w /usr/bin/docker -p x -k conteneurs
sudo auditctl -w /etc/apparmor.d/ -p wa -k apparmor-conf

# Rechercher
sudo ausearch -k conteneurs -ts today
```

## 85. Alertes : ce qui doit réveiller l'astreinte

| Alerte | Seuil indicatif | Action |
|---|---|---|
| Refus AppArmor sur service prod | > 5 / heure | `aa-logprof`, vérifier mise à jour |
| Processus tué par seccomp (SIGSYS) | ≥ 1 | Filtre trop strict OU exploitation tentée |
| OOM kill conteneur | > 3 / jour | Fuite mémoire ou `MemoryMax` trop bas |
| Restart loop Docker | > 5 restarts / 10 min | Config cassée, rollback image |
| Nouveau profil AppArmor en complain | rappel 7 j | Repasser en enforce après apprentissage |
| Image Docker > 90 jours | hebdo | Rebuild + scan |

Exemple d'alerte simple avec un timer systemd :

```bash
#!/bin/bash
# /usr/local/bin/check-sandbox-alerts.sh
SEUIL=5
N=$(journalctl -k --since "1 hour ago" | grep -c "apparmor.*DENIED")
[ "$N" -gt "$SEUIL" ] && echo "ALERTE: $N refus AppArmor sur $(hostname) la dernière heure" \
  | mail -s "[sandbox] refus AppArmor" astreinte@example.com
```

## 86. Tableau de bord minimal (sans usine à gaz)

