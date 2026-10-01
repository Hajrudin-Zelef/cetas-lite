---
id: collect-261001-rattrapage/rattrapage/ssh-guide-4
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [618, 844]
sha256: 9a1649de53670ce765b6b08b67dcce5e74efe9ad715071919aa5fca1f2b9288e
---

# Guide SSH approfondi

**GNOME / Ubuntu desktop** : le trousseau GNOME fournit un agent compatible
(`SSH_AUTH_SOCK` déjà positionné) ; `ssh-add` suffit.

**macOS** : l'agent est lancé par le système ; persistance dans le trousseau :

```
Host *
    AddKeysToAgent yes
    UseKeychain yes
    IdentityFile ~/.ssh/id_ed25519
```

**Windows (OpenSSH natif)** : service `ssh-agent` à passer en automatique :

```powershell
Set-Service -Name ssh-agent -StartupType Automatic
Start-Service ssh-agent
ssh-add $env:USERPROFILE\.ssh\id_ed25519
```

## 20. Agent forwarding : fonctionnement et **dangers**

`ssh -A bastion` (ou `ForwardAgent yes`) expose ton agent local **au serveur
distant** via un socket. Depuis ce serveur, `ssh autre-machine` utilise tes
clés locales sans les copier. Pratique pour les chaînes de sauts… et dangereux :

- **Quiconque a root sur le serveur distant peut utiliser ton agent** tant que
  tu es connecté (il ne peut pas *extraire* les clés, mais il peut *signer*
  avec : se connecter partout où tes clés sont autorisées).
- Le risque persiste même après déconnexion si le socket n'est pas nettoyé.

Règles d'équipe :

1. **Jamais** de forwarding vers une machine que tu n'administres pas
   entièrement (mutualisé, client, prestataire).
2. Préfère **`ProxyJump`** (section 26) : l'authentification vers la cible se
   fait **depuis ton poste**, le bastion ne voit que du trafic chiffré.
3. Si tu dois vraiment forwarder (vieux workflow) : `ForwardAgent yes`
   **uniquement** sur le bloc `Host` concerné, jamais dans `Host *`, et
   déconnecte-toi dès que c'est fini.
4. Alternative moderne : les **certificats SSH** (section 51) suppriment le
   besoin de forwarding dans la plupart des cas.

Vérifier qu'un forwarding est actif (et le couper si surprise) :

```bash
serveur$ echo "$SSH_AUTH_SOCK"
# vide = pas de forwarding ; /tmp/ssh-XXX/agent.PID = forwarding actif
```

## 21. `ssh-copy-id` et déploiement des clés

```bash
client$ ssh-copy-id -i ~/.ssh/id_ed25519.pub zelef@web1
```

`ssh-copy-id` ajoute ta clé publique à `~/.ssh/authorized_keys` distant en
créant le répertoire avec les bonnes permissions. Options utiles :

```bash
client$ ssh-copy-id -i ~/.ssh/id_ed25519.pub -p 2222 zelef@web1   # port custom
client$ ssh-copy-id -i ~/.ssh/id_ed25519.pub -o ProxyJump=bastion zelef@web1
```

À l'échelle d'un parc, **n'utilise pas `ssh-copy-id` à la main** : déploie
`authorized_keys` via Ansible (template + `authorized_key`), Puppet ou ton
provisionnement cloud (cloud-init). `ssh-copy-id` reste parfait pour les
machines isolées et le dépannage.

Vérification post-déploiement :

```bash
client$ ssh -o PreferredAuthentications=publickey -o PubkeyAuthentication=yes web1 'echo OK'
```

## 22. `authorized_keys` : format

Fichier `~/.ssh/authorized_keys` (permissions `600`, répertoire `~/.ssh` en
`700`) côté serveur. Une clé par ligne :

```
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIEXEMPLEdefictifNePasUtiliser zelef@poste-prod-2026-09-26
```

Format : `<type> <clé en base64> <commentaire>`. Les lignes vides et celles
commençant par `#` sont ignorées. Les **options** se placent en début de ligne,
séparées par des virgules (sections 23-25) :

```
from="10.0.0.0/8",command="/usr/local/bin/backup.sh",no-pty ssh-ed25519 AAAAC3... zelef@backup
```

Erreurs fréquentes : clé coupée/collée avec un saut de ligne au milieu,
commentaire contenant une virgule non échappée dans les options, permissions
`~/.ssh` ou home trop permissives (le serveur refuse alors silencieusement la
clé si `StrictModes yes`, défaut).

## 23. Restreindre par origine : `from=`

```
from="192.168.10.0/24,2001:db8::/32" ssh-ed25519 AAAAC3... zelef@poste-admin
from="bastion.example.com" ssh-ed25519 AAAAC3... deploy@ci
```

- La clé n'est acceptée que si la connexion vient de ces adresses/noms.
- **Oblige** à passer par le bastion ou le réseau d'administration : même si la
  clé privée fuit, elle est inutilisable depuis Internet.
- Le motif DNS est résolu à chaque connexion : avec `UseDNS no` côté serveur
  (recommandé, section 65), préfère les IP/CIDR.
- Combine avec `AllowUsers`/`AllowGroups` côté serveur pour la défense en
  profondeur.

## 24. Clés à usage unique : `command=`, `no-pty`

```
command="/usr/local/bin/deploy.sh",no-pty,no-agent-forwarding,no-X11-forwarding,no-port-forwarding ssh-ed25519 AAAAC3... ci@runner
```

- `command="..."` : **ignore** la commande demandée par le client et exécute
  toujours celle-ci. Idéal pour : scripts de backup, déploiement CI, tunnels
  dédiés, `rrsync`.
- Le script reçoit la commande originale dans `$SSH_ORIGINAL_COMMAND` et peut
  décider de l'autoriser ou non (pattern classique : wrapper qui valide).
- `no-pty` : pas de terminal interactif → la clé ne sert qu'en non-interactif.
- Toujours combiner `command=` avec les `no-*` (section 25) : sinon la clé
  « restreinte » peut quand même ouvrir des tunnels.

Exemple de wrapper :

```bash
#!/bin/bash
# /usr/local/bin/deploy.sh — n'autorise que le script de déploiement attendu
case "$SSH_ORIGINAL_COMMAND" in
  "/opt/app/deploy.sh production") exec /opt/app/deploy.sh production ;;
  *) echo "Commande non autorisée : $SSH_ORIGINAL_COMMAND" >&2; exit 1 ;;
esac
```

## 25. `no-agent-forwarding`, `no-X11-forwarding`, `no-port-forwarding`

Tableau des options de restriction par clé :

| Option | Effet |
|---|---|
| `no-agent-forwarding` | Interdit le forwarding d'agent avec cette clé |
| `no-X11-forwarding` | Interdit le X11 forwarding |
| `no-port-forwarding` | Interdit `-L`, `-R`, `-D` |
| `no-pty` | Pas d'allocation de pseudo-terminal |
| `no-user-rc` | Ignore `~/.ssh/rc` |
| `permitopen="host:port"` | Avec `no-port-forwarding` absent : limite les tunnels `-L` à ces destinations |
| `permitlisten="port"` | Limite les tunnels `-R` à ces ports |

Recette « clé de service minimale » :

```
from="10.0.5.0/24",command="/usr/local/bin/backup.sh",no-pty,no-agent-forwarding,no-X11-forwarding,no-port-forwarding ssh-ed25519 AAAAC3... backup@client
```

Principe : **moindre privilège**. Une clé de backup n'a pas besoin d'un shell,
ni de tunnels, ni de forwarding d'agent. Si elle fuit, l'attaquant ne peut
exécuter que le script de backup depuis le réseau de backup.

## 26. `ProxyJump` et bastions

Le bastion (jump host) est le point d'entrée unique vers un réseau privé.
`ProxyJump` (option `-J`) fait transiter la connexion **chiffrée de bout en
bout** : le bastion relaie des octets, il ne voit ni tes clés ni le contenu.

```bash
client$ ssh -J bastion web1-interne
client$ ssh -J zelef@bastion:2222 web1-interne
```

En config (la forme à privilégier) :

```
Host bastion
    HostName bastion.example.com
    User zelef

Host web1-interne
    HostName 10.10.1.10
    User zelef
    ProxyJump bastion
```

Puis simplement `ssh web1-interne`. Avantages sur le forwarding d'agent :

- L'authentification vers la cible part **de ton poste** (l'agent local signe).
- Le bastion ne peut pas réutiliser tes clés.
- Fonctionne avec `scp`, `sftp`, `rsync`, `git`, `ansible` sans option spéciale.

Copie de fichiers à travers un bastion :

```bash
client$ scp -J bastion rapport.pdf web1-interne:/tmp/
client$ rsync -avz -e "ssh -J bastion" ./app/ web1-interne:/opt/app/
```

## 27. Chaînes de sauts : plusieurs bastions

`ProxyJump` accepte une liste séparée par des virgules (évaluée de gauche à
droite, du plus proche au plus lointain) :

```bash
client$ ssh -J bastion-paris,bastion-dmz srv-dmz-interne
```

```
Host bastion-paris
    HostName bastion.paris.example.com
    User zelef

Host bastion-dmz
    HostName 10.20.0.5
    User zelef
    ProxyJump bastion-paris

Host srv-dmz-interne
    HostName 10.30.0.8
    User zelef
    ProxyJump bastion-dmz
```

Pour les cas exotiques (bastion qui n'est pas SSH, ou besoin d'une commande
spécifique), l'ancêtre `ProxyCommand` reste disponible :

