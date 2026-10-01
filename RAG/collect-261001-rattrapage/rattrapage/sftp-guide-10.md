---
id: collect-261001-rattrapage/rattrapage/sftp-guide-10
title: "Guide SFTP complet — OpenSSH, chroot, supervision et automatisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/sftp_guide.md
source_anchor: ""
source_lines: [1926, 2066]
sha256: f82d31700fbdd4927ca080c35dfeb214d25f1979e550da4b6f987db3eb577661
---

# Guide SFTP complet — OpenSSH, chroot, supervision et automatisation

| Paramètre | Effet | Valeur de départ |
|---|---|---|
| `MaxSessions` | Sessions multiplexées par connexion | 5–10 |
| `MaxStartups` | Connexions simultanées non authentifiées (`10:30:60`) | défaut OK, durcir si flood |
| `ClientAliveInterval` | Détection des sessions mortes | 300 s |
| `LoginGraceTime` | Temps max pour s'authentifier | 30 s |
| Buffers SFTP (`-B`, `-R`) | Débit sur liens à forte latence | `-B 262144 -R 64` |

Autres leviers :

- **Stockage** : `/srv` sur RAID ou volume rapide ; surveillez `iostat -x` pendant les pics.
- **Parallélisme** : `lftp --parallel` ou plusieurs flux plutôt qu'un seul gros fichier.
- **Fenêtres de transfert** : étalez les batchs partenaires (cron à heures décalées)
  pour éviter le pic de 2h du matin.
- **Limites système** : `nofile`/`nproc` suffisants si des dizaines de partenaires
  transfèrent en parallèle (`/etc/security/limits.conf`, systemd `LimitNOFILE`).

---

## 89. Annexe — sshd_config complet commenté (modèle)

```
# ============ /etc/ssh/sshd_config — modèle SFTP entreprise ============
Port 22
Protocol 2

# Clés hôtes (Ed25519 en priorité)
HostKey /etc/ssh/ssh_host_ed25519_key
HostKey /etc/ssh/ssh_host_rsa_key

# --- Authentification : clés uniquement ---
PermitRootLogin no
PasswordAuthentication no
PubkeyAuthentication yes
ChallengeResponseAuthentication no
UsePAM yes

# --- Crypto moderne ---
KexAlgorithms curve25519-sha256,curve25519-sha256@libssh.org,ecdh-sha2-nistp521,ecdh-sha2-nistp384,ecdh-sha2-nistp256,diffie-hellman-group18-sha512
Ciphers aes128-gcm@openssh.com,aes256-gcm@openssh.com,chacha20-poly1305@openssh.com,aes256-ctr,aes128-ctr
MACs hmac-sha2-512-etm@openssh.com,hmac-sha2-256-etm@openssh.com,umac-128-etm@openssh.com

# --- Robustesse ---
LoginGraceTime 30
MaxAuthTries 3
MaxSessions 10
ClientAliveInterval 300
ClientAliveCountMax 2

# --- Divers ---
X11Forwarding no
AllowAgentForwarding no
AllowTcpForwarding no
PermitTunnel no
PrintMotd no

# --- SFTP : serveur intégré + logs détaillés ---
Subsystem sftp internal-sftp -f AUTH -l VERBOSE

# --- Clés publiques centralisées hors chroot ---
AuthorizedKeysFile /etc/ssh/authorized_keys/%u

# --- Restriction d'accès globale (exemple) ---
# AllowGroups sftpusers sftp_compta admins

# ============ BLOCS MATCH : TOUJOURS EN FIN DE FICHIER ============
# Comptes de transfert : chroot, pas de shell, pas de mot de passe, pas de tunnel
Match Group sftpusers
    ChrootDirectory /srv/sftp/%u
    ForceCommand internal-sftp
    PasswordAuthentication no
    AllowTcpForwarding no
    AllowAgentForwarding no
    X11Forwarding no
    PermitTunnel no

# Comptabilité : chroot commun au groupe
Match Group sftp_compta
    ChrootDirectory /srv/sftp/compta
    ForceCommand internal-sftp
    PasswordAuthentication no
    AllowTcpForwarding no
    AllowAgentForwarding no
    X11Forwarding no
    PermitTunnel no

# Exception ciblée : vieux copieur sans support des clés (dette à résorber)
# Match User svc_copieur_ancien
#     ChrootDirectory /srv/sftp/scans_copieurs
#     ForceCommand internal-sftp
#     PasswordAuthentication yes
```

Après toute modification : `sudo sshd -t && sudo systemctl reload ssh`.

---

## 90. Annexe — script d'onboarding complet

```bash
#!/bin/bash
# /usr/local/sbin/sftp_onboard.sh — crée un espace client de A à Z
# Usage : sudo sftp_onboard.sh sftp_acme client_acme "Dépôt SFTP client ACME"
set -euo pipefail
U=${1:?usage: $0 <user> <perimetre> "<commentaire>"}
PERIM=${2:?}
COMMENT=${3:-"Compte SFTP $U"}
BASE=/srv/sftp

# 1. Utilisateur sans shell, mot de passe verrouillé
useradd -m -d "$BASE/$PERIM" -s /usr/sbin/nologin -G sftpusers -c "$COMMENT" "$U"
passwd -l "$U"

# 2. Arborescence : chroot root:root 0755, dépôt inscriptible par l'utilisateur
mkdir -p "$BASE/$PERIM"/{depot,restitutions}
chown root:root "$BASE/$PERIM"; chmod 0755 "$BASE/$PERIM"
chown "$U":sftpusers "$BASE/$PERIM/depot"; chmod 0750 "$BASE/$PERIM/depot"
chown root:sftpusers "$BASE/$PERIM/restitutions"; chmod 0750 "$BASE/$PERIM/restitutions"

# 3. Clé publique (collée par l'exploitant)
echo "Collez la clé PUBLIQUE de $U puis Ctrl+D :"
sh -c "cat > /etc/ssh/authorized_keys/$U"
chmod 0644 "/etc/ssh/authorized_keys/$U"

# 4. Contrôles
echo "--- Contrôle chroot ---"; namei -l "$BASE/$PERIM" | head -8
echo "--- Contrôle sshd ---"; sshd -t && echo "sshd_config OK"

# 5. Rappel des actions manuelles restantes
cat <<EOF
Actions restantes :
  - [ ] setquota -u $U <soft> <hard> 0 0 /srv   (quota)
  - [ ] bloc Match si chroot != /srv/sftp/$U, puis systemctl reload ssh
  - [ ] test : sftp -i <cle_test> $U@<hote> (put/get + tentative d'évasion)
  - [ ] fiche d'exploitation § 83 remplie
EOF
logger -t sftp_onboard "compte $U créé (périmètre $PERIM)"
```

> Ce script standardise les créations : même arborescence, mêmes permissions,
> mêmes contrôles à chaque fois. Adaptez les valeurs (quota, groupe) à votre contexte.
