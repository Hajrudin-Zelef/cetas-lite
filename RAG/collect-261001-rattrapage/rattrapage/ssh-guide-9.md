---
id: collect-261001-rattrapage/rattrapage/ssh-guide-9
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["agent", "distribution"]
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [1718, 1905]
sha256: 974f29fee539621d04ea342b5ad32b17a289ab41ebc4515d2173e2ffb52e9443
---

# Guide SSH approfondi

```bash
#!/bin/bash
# /usr/local/bin/ssh-alert.sh — alerte sur connexion root réussie
if journalctl -u ssh --since "5 minutes ago" | grep -q "Accepted.*for root"; then
  echo "ALERTE: connexion root SSH détectée sur $(hostname) à $(date)" \
    | mail -s "[SEC] SSH root $(hostname)" astreinte@example.com
fi
```

Mieux : forwarder `auth.log` vers ta stack (rsyslog → Loki/ELK/Wazuh) et écrire
les règles là-bas. Le script ci-dessus dépanne en attendant.

## 57. Sauvegarder ses clés (critique)

Tes clés privées **sont** ton identité d'admin. Les perdre = perdre l'accès ;
les exposer = compromettre le parc. Les deux se traitent par la sauvegarde
**chiffrée**.

Stratégie recommandée :

1. **Chiffrement** : archive chiffrée (pas de clé en clair sur un disque
   externe, un NAS ou un cloud).

```bash
client$ tar -czf - ~/.ssh | gpg --symmetric --cipher-algo AES256 -o ssh-backup-2026-09-26.tar.gz.gpg
# Restauration :
client$ gpg -d ssh-backup-2026-09-26.tar.gz.gpg | tar -xzf - -C /tmp/restauration
```

2. **3-2-1** : 3 copies, 2 supports différents, 1 hors site (coffre, cloud
   chiffré). La passphrase GPG **sur papier au coffre**, pas dans le même
   cloud.
3. **Quoi sauvegarder** : `~/.ssh/` entier (privées + publiques + config +
   known_hosts), la passphrase dans un gestionnaire de mots de passe
   d'équipe (Bitwarden/Vaultwarden), et **la liste des serveurs où chaque clé
   est déployée** (sinon la restauration est aveugle).
4. **Clés de CA** (section 52) : sauvegarde **offline** dédiée (clé USB
   chiffrée au coffre), jamais sur le réseau.

Ce qu'il ne faut **jamais** faire :

- ❌ Commiter une clé privée dans Git (même « temporairement », même en privé).
- ❌ L'envoyer par mail, Slack, ticket.
- ❌ La laisser dans `/tmp`, un partage réseau, une image Docker/VM.
- ❌ La stocker en clair dans un cloud « parce que c'est pratique ».

Test de restauration **une fois par an** : restaure sur une machine jetable et
vérifie qu'une connexion aboutit. Une sauvegarde non testée n'est pas une
sauvegarde.

## 58. Rotation des clés

Politique simple et tenable :

| Événement | Action |
|---|---|
| Départ d'un collaborateur | Supprimer ses clés de tous les `authorized_keys` **le jour J** (+ révoquer certificats) |
| Compromission suspectée (poste volé, malware) | Nouvelle clé **immédiatement**, ancienne révoquée partout |
| Rotation préventive | Tous les 12-24 mois pour les clés humaines, 6-12 mois pour les clés de service |
| Changement de périmètre | Nouvelle clé dédiée plutôt que réutiliser l'ancienne |

Procédure de rotation sans coupure :

1. Générer la nouvelle clé (nouveau commentaire daté).
2. La déployer **en plus** de l'ancienne (Ansible/`ssh-copy-id`).
3. Vérifier la connexion avec la nouvelle (`ssh -i nouvelle`).
4. Retirer l'ancienne de tous les `authorized_keys`.
5. Archiver l'ancienne clé (au cas où un serveur aurait été oublié, 30 jours),
   puis la détruire (`shred -u`).

Avec les certificats (section 51), la rotation devient triviale : courte durée
de vie (`-V :+8h`) = rotation **automatique** quotidienne.

## 59. Inventaire des clés autorisées sur le parc

Le risque silencieux : des `authorized_keys` qui accumulent les clés
d'anciens collègues, de prestataires, de tests oubliés. Audit trimestriel :

```bash
# Sur chaque serveur : lister toutes les clés autorisées avec leur commentaire
serveur# for f in /home/*/.ssh/authorized_keys /root/.ssh/authorized_keys; do
  [ -f "$f" ] && echo "== $f" && ssh-keygen -l -f "$f"
done
```

Version Ansible (à mettre dans ton dépôt d'exploitation) :

```yaml
- name: Inventaire des authorized_keys
  hosts: all
  tasks:
    - name: Lister les clés autorisées
      ansible.builtin.shell: |
        for f in /home/*/.ssh/authorized_keys /root/.ssh/authorized_keys; do
          [ -f "$f" ] && echo "== $f" && ssh-keygen -l -f "$f"
        done
      register: cles
      changed_when: false
    - ansible.builtin.debug: var=cles.stdout_lines
```

Règles d'hygiène :

- Chaque ligne doit avoir un **commentaire identifiable** (`user@machine-date`).
  Ligne sans commentaire = à clarifier ou à supprimer.
- Zéro clé d'ancien salarié, zéro clé « test », zéro doublon.
- Croise avec les logs `Accepted ... ED25519 SHA256:...` (section 54) : une clé
  autorisée mais **jamais utilisée depuis 6 mois** est une candidate à la
  suppression.

## 60. Mettre à jour OpenSSH

OpenSSH évolue vite (nouvelles KEX, retraits d'algorithmes faibles, CVE). Reste
sur les versions de ta distribution **à jour** plutôt que de compiler à la main
(sauf besoin précis).

**Debian / Ubuntu :**

```bash
serveur# apt update && apt install --only-upgrade openssh-server openssh-client
serveur# sshd -t && systemctl reload ssh
serveur# ssh -V   # vérifier
```

**Après une mise à jour majeure**, vérifie :

1. `sshd -T` : des options **dépréciées** ? (ex. `ChallengeResponseAuthentication`
   renommé `KbdInteractiveAuthentication` dans les versions récentes —
   l'ancien nom reste accepté un temps, mais migre.)
2. Les algorithmes par défaut : une montée de version peut **désactiver**
   d'anciens algos (ex. RSA-SHA1) et couper de vieux clients/équipements.
   Teste tes cas limites (équipements réseau anciens, scripts).
3. `man sshd_config` de la version installée : la doc embarquée fait foi, pas
   un tuto de 2019.

**Politique** : applique les mises à jour de sécurité sous 48h sur les serveurs
exposés (unattended-upgrades bien configuré, cf. ton guide Debian/Ubuntu),
fenêtre planifiée pour les montées de version majeures, avec test de
non-régression SSH (une connexion clé + un `rsync` + un tunnel si tu en as).

---

# DÉPANNAGE

## 61. Dépannage : la méthode (`-v`, `-vvv`)

Devant une panne SSH, toujours dans cet ordre :

1. **Reproduire avec verbosité** : `ssh -v cible` (souvent suffisant),
   `ssh -vvv cible` pour les cas tordus. Lis **les 20 dernières lignes** :
   la cause y est presque toujours (`Offering public key`, `Authentications
   that can continue`, `No more authentication methods to try`).
2. **Vérifier la config effective** : `ssh -G cible` (section 9). 50 % des
   « pannes » sont une option du `~/.ssh/config` qu'on avait oubliée.
3. **Regarder côté serveur** en parallèle : `journalctl -u ssh -f` ou
   `/var/log/auth.log`. Le serveur dit souvent exactement pourquoi il refuse
   (`Authentication refused: bad ownership or modes`, `User zelef not allowed
   because not listed in AllowUsers`).
4. **Isoler** : tester sans la config (`ssh -F /dev/null`), sans l'agent
   (`ssh -o IdentitiesOnly=yes -i clé`), depuis un autre réseau.

```bash
client$ ssh -v web1 2>&1 | tail -30
client$ ssh -F /dev/null -i ~/.ssh/id_ed25519 zelef@192.0.2.10
```

Et surtout : **garde une session ouverte** quand tu modifies `sshd_config`
(section 41). Le dépannage commence par ne pas s'enfermer dehors.

## 62. Cas 1 : `Permission denied (publickey)`

Le classique. Message complet typique :

```
zelef@web1: Permission denied (publickey).
```

Diagnostic avec `ssh -v` : repère ces lignes :

```
debug1: Offering public key: /home/zelef/.ssh/id_ed25519 ED25519 SHA256:...
debug1: Authentications that can continue: publickey
debug1: No more authentication methods to try.
```

Le serveur a **refusé toutes les clés proposées** (ou aucune n'a été proposée).
Checklist, dans l'ordre :

