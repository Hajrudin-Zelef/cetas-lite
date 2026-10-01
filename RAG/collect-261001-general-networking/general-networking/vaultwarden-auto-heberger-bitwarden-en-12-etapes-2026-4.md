---
id: collect-261001-general-networking/general-networking/vaultwarden-auto-heberger-bitwarden-en-12-etapes-2026-4
title: "Mise à jour complète du système"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["arr", "mai"]
source: docs/RAG/collect-261001-general-networking/vaultwarden-auto-heberger-bitwarden-en-12-etapes-2026.md
source_anchor: ""
source_lines: [276, 365]
sha256: 78a780bb20e59c1ca67908a9ffd6b8c105d93c12f55c3be44e94adbaaf2d2b74
---

# Mise à jour complète du système

- **Extension / application :** sur l’écran de connexion, cliquez sur l’icône d’engrenage (Paramètres régionaux), choisissez « Auto-hébergé » et renseignez`https://coffre.exemple.fr` dans le champ « URL du serveur ».
- **Mobile (iOS / Android) :** même principe, le réglage de l’environnement se trouve sur l’écran de connexion avant la saisie des identifiants.
- **CLI :** configurez le serveur puis authentifiez-vous.

```
# Installer le client en ligne de commande Bitwarden
npm install -g @bitwarden/cli
# Pointer la CLI vers votre serveur auto-hébergé
bw config server https://coffre.exemple.fr
# Se connecter (l'e-mail et le mot de passe maître seront demandés)
bw login
# Déverrouiller et lister un élément pour valider
export BW_SESSION="$(bw unlock --raw)"
bw list items | head
```
Si `bw config server` est oublié, la CLI tentera de joindre le cloud Bitwarden officiel et la connexion échouera avec une erreur d’identifiants : c’est le piège le plus courant côté client. Une fois connecté sur un appareil, créez quelques entrées de test et observez la synchronisation instantanée sur les autres clients – preuve que les WebSockets fonctionnent.

## Étape 10 – Renforcer la sécurité : 2FA, Argon2id et politiques

Le coffre est en ligne ; il faut maintenant le durcir au niveau applicatif. Trois leviers sont prioritaires.

### Activer la double authentification

Dans les paramètres de sécurité de votre compte (via l’interface web), activez la 2FA. Vaultwarden gère gratuitement le TOTP (application d’authentification), les clés matérielles FIDO2/WebAuthn comme une YubiKey, et l’e-mail. Pour un compte administrateur de coffre, la clé matérielle est l’option la plus résistante au hameçonnage. Sauvegardez les codes de récupération hors ligne.

### Passer la dérivation de clé en Argon2id

La robustesse de votre mot de passe maître dépend de la fonction de dérivation de clé (KDF). Par défaut, les clients Bitwarden utilisent PBKDF2-SHA256 avec 600 000 itérations. Vous pouvez basculer sur **Argon2id**, plus résistant aux attaques par matériel spécialisé, depuis Paramètres > Sécurité > Clé de chiffrement. La documentation Bitwarden sur les algorithmes KDF détaille les paramètres recommandés (mémoire, itérations, parallélisme). C’est un réglage côté client : il rechiffre votre coffre lors du changement.

Pour une organisation, définissez aussi des **politiques** via le panneau admin ou l’organisation : longueur minimale du mot de passe maître, 2FA obligatoire pour tous les membres, désactivation de l’export individuel. Accédez au panneau d’administration sur `https://coffre.exemple.fr/admin` et authentifiez-vous avec le mot de passe choisi à l’étape 5. C’est là que vous invitez de nouveaux utilisateurs et consultez l’état du service.

## Étape 11 – Mettre en place des sauvegardes chiffrées automatiques

Sans sauvegarde, l’auto-hébergement est une roulette russe. La base de Vaultwarden est en SQLite, en mode WAL : copier bêtement le fichier pendant que le service tourne risque de produire une sauvegarde corrompue. La méthode correcte est la commande `.backup` de `sqlite3`, qui réalise une copie cohérente à chaud. Le script suivant sauvegarde la base, les pièces jointes et les clés, puis chiffre le tout en AES-256 avec GPG.

```
#!/usr/bin/env bash
# /opt/vaultwarden/backup.sh
set -euo pipefail
SRC="/opt/vaultwarden/vw-data"
DEST="/opt/vaultwarden/backups"
STAMP="$(date +%F_%H%M)"
mkdir -p "$DEST"
# 1) Copie cohérente de la base SQLite (mode WAL)
sqlite3 "$SRC/db.sqlite3" ".backup '$DEST/db-$STAMP.sqlite3'"
# 2) Archive des fichiers : pièces jointes, sends, clés, config
tar czf "$DEST/files-$STAMP.tar.gz" -C "$SRC" \
  attachments sends config.json rsa_key.pem 2>/dev/null || true
# 3) Chiffrement AES-256 de la base
gpg --batch --yes --symmetric --cipher-algo AES256 \
    --passphrase-file /root/.vw-backup-pass \
    "$DEST/db-$STAMP.sqlite3"
rm -f "$DEST/db-$STAMP.sqlite3"
# 4) Rotation : conserver 14 jours
find "$DEST" -type f -mtime +14 -delete
echo "Sauvegarde $STAMP terminée."
```
Rendez le script exécutable (`chmod +x backup.sh`), stockez la phrase secrète dans `/root/.vw-backup-pass` (permissions `600`), puis planifiez une exécution quotidienne via cron :

```
# Éditer la crontab root
sudo crontab -e
# Ajouter : sauvegarde tous les jours à 3h15
15 3 * * * /opt/vaultwarden/backup.sh >> /var/log/vw-backup.log 2>&1
```
Une sauvegarde n’a de valeur que si la **restauration** est testée. Pour restaurer : déchiffrez l’archive (`gpg --decrypt db-DATE.sqlite3.gpg > db.sqlite3`), arrêtez la pile (`docker compose down`), remplacez les fichiers dans `vw-data`, puis relancez. Copiez régulièrement les sauvegardes chiffrées hors du serveur (stockage objet européen, disque externe) pour survivre à une perte totale du VPS.

## Étape 12 – Fail2ban, mises à jour et supervision

Dernière étape : protéger le coffre contre le bourrage d’identifiants et garder la pile à jour – un impératif rappelé par le CVE-2025-24365, divulgué le 27 janvier 2025, corrigé dans la version 1.33.0 et dont la fiche du NVD (NIST) a encore été mise à jour en juin 2026 selon SentinelOne. L’urgence n’a rien de théorique : une étude d’infrastructure Censys recensait encore, en mai 2026, environ 1 700 serveurs Vaultwarden exposés et vulnérables à ces failles critiques de 2025, preuve qu’une version datée ne doit jamais rester en production. Comme nous avons activé `LOG_FILE` à l’étape 5, Vaultwarden journalise les échecs de connexion, ce que fail2ban peut exploiter. Créez un filtre puis une prison.

```
# /etc/fail2ban/filter.d/vaultwarden.conf
[Definition]
failregex = ^.*Username or password is incorrect\. Try again\. IP: <ADDR>\. Username:.*$
ignoreregex =
```
```
# /etc/fail2ban/jail.d/vaultwarden.local
[vaultwarden]
enabled  = true
filter   = vaultwarden
logpath  = /opt/vaultwarden/vw-data/vaultwarden.log
port     = 80,443
maxretry = 5
bantime  = 14400
findtime = 14400
```
Rechargez fail2ban (`sudo systemctl restart fail2ban`) puis vérifiez la prise en compte avec `sudo fail2ban-client status vaultwarden`. Au bout de 5 tentatives ratées en 4 heures, l’adresse fautive est bannie. Pour la **mise à jour**, modifiez le tag d’image dans `docker-compose.yml` vers une version récente – par exemple la **1.37.2** du 22 août 2026, ou a minima la **1.36.0** du 3 mai 2026, qui embarque le web vault 2026.4.1, mitige plusieurs avis de sécurité selon WinterFlow et a clos à elle seule 6 avis de sécurité en ajoutant l’archivage d’éléments après une série de 13 avis de sécurité publiés entre février et mai 2026 sur les versions 1.35.3 à 1.36.0 selon WZ-IT, sans oublier l’étape intermédiaire de la **1.35.8** du 25 avril 2026, qui a mis à jour la toolchain Rust, corrigé des soucis DNS ainsi que la politique de mot de passe maître et d’authentification, et réglé un bug affectant les codes de récupération et les jetons de rafraîchissement (*refresh tokens*) selon WinterFlow.io – puis :

