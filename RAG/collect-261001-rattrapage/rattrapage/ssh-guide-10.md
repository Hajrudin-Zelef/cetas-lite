---
id: collect-261001-rattrapage/rattrapage/ssh-guide-10
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr"]
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [1906, 2113]
sha256: d24fc81b10893cbc3f89634b1bbde0b395860421bc4b37063ed9a59e0a6ebe60
---

# Guide SSH approfondi

1. **La bonne clé est-elle proposée ?** `ssh -v` → `Offering public key:`.
   Si ta clé n'apparaît pas : `IdentityFile` manquant, `IdentitiesOnly yes`
   sans `IdentityFile`, ou agent vide (`ssh-add -l`).
2. **La clé publique est-elle déployée ?** Compare l'empreinte locale
   (`ssh-keygen -l -f ~/.ssh/id_ed25519.pub`) avec `authorized_keys` distant
   (via un autre accès).
3. **Permissions côté serveur** : `~/.ssh` en `700`, `authorized_keys` en
   `600`, **home en 755 max** (un home en 777 fait refuser la clé avec
   `StrictModes yes`). Le log serveur dit `bad ownership or modes`.
4. **La clé est-elle restreinte ?** Options `from=` (tu n'es pas sur le bon
   réseau), `command=` (normal si c'est une clé de service).
5. **Serveur** : `PubkeyAuthentication yes` ? L'utilisateur est-il dans
   `AllowUsers`/`AllowGroups` ? Le compte est-il verrouillé (`passwd -S`) ?
6. **Trop de clés proposées** avant la bonne → voir cas 6 (section 67).

Réparation express des permissions :

```bash
serveur$ chmod 700 ~/.ssh && chmod 600 ~/.ssh/authorized_keys && chmod 755 ~
```

## 63. Cas 2 : `Host key verification failed`

```
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@    WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
```

Significations possibles, par probabilité :

1. **Le serveur a été réinstallé** (clés d'hôte régénérées) — cas le plus
   fréquent en labo/homelab.
2. **Tu tombes sur une autre machine** (IP réattribuée en DHCP, NAT qui a
   changé, DNS menteur).
3. **Attaque homme-du-milieu** — rare, mais c'est exactement ce que ce message
   est censé détecter : ne le contourne pas à l'aveugle sur un réseau hostile.

Conduite à tenir :

```bash
# 1. Vérifier la cause (réinstall ? nouvelle IP ?)
# 2. Comparer l'empreinte proposée avec l'inventaire (section 8)
# 3. Si légitime, purger l'ancienne entrée puis reconnecter :
client$ ssh-keygen -R web1
client$ ssh-keygen -R 192.0.2.10
client$ ssh web1   # vérifier l'empreinte affichée avant "yes"
```

**Ne fais jamais** : `StrictHostKeyChecking=no` en global, ni supprimer tout
`known_hosts` « pour aller vite ». En automatisation, utilise
`StrictHostKeyChecking=accept-new` (section 7), jamais `no`.

## 64. Cas 3 : la connexion freeze / timeout

Symptômes : `ssh web1` reste bloqué, ou la session se fige après quelques
minutes d'inactivité.

**À la connexion** :

```bash
client$ ssh -v web1
# Bloqué sur "Connecting to 192.0.2.10 [192.0.2.10] port 22." -> réseau/firewall
```

- `telnet 192.0.2.10 22` / `nc -vz 192.0.2.10 22` : si ça ne répond pas, le
  problème est **réseau** (firewall, routage, sshd arrêté), pas SSH.
- Bloqué après `SSH-2.0-OpenSSH...` : filtrage applicatif ou MTU (rare).
- Avec `-J` : c'est le **bastion** qui est injoignable, pas la cible.

**Session qui freeze en inactivité** : un NAT/firewall coupe les connexions
TCP idle. Remède client (section 13) :

```
Host *
    ServerAliveInterval 60
    ServerAliveCountMax 3
```

Et côté serveur : `ClientAliveInterval 300` (section 44). Pour les tunnels
persistants : `autossh` (section 35).

**Freeze sur gros transfert** : suspecte le MTU (VPN/PPPoE) → teste avec des
petits paquets, ou `ping -M do -s 1472`.

## 65. Cas 4 : l'agent ne transmet pas la clé (forwarding)

Symptôme : `ssh -A bastion` puis, sur le bastion, `ssh cible` demande un mot
de passe ou échoue alors que ça marche en direct depuis ton poste.

Diagnostic :

```bash
# Sur TON poste :
client$ ssh-add -l          # l'agent a-t-il la clé ? Sinon : ssh-add
client$ echo $SSH_AUTH_SOCK # le socket est-il défini ?
# Sur le BASTION :
bastion$ echo $SSH_AUTH_SOCK # vide = pas de forwarding arrivé
bastion$ ssh-add -l          # doit lister TES clés (via le socket transmis)
```

Causes fréquentes :

1. `ForwardAgent yes` absent (et pas de `-A`). Vérifie `ssh -G bastion | grep forwardagent`.
2. **Le serveur l'interdit** : `AllowAgentForwarding no` dans `sshd_config`.
3. `sudo -s` sur le bastion : `sudo` **ne conserve pas** `SSH_AUTH_SOCK` par
   défaut → l'agent « disparaît » en root. (Et c'est tant mieux : voir les
   dangers section 20.)
4. Multiplexage : une connexion maîtresse ouverte **sans** `-A` est réutilisée
   par les suivantes → le forwarding ne s'active pas. `ssh -O stop bastion`
   puis reconnecte avec `-A`.

Rappel : la bonne réponse est souvent **ne pas utiliser le forwarding** mais
`ProxyJump` (section 26).

## 66. Cas 5 : connexion lente à s'établir (DNS, GSSAPI)

Symptôme : 10 à 30 secondes entre `ssh` et le prompt mot de passe/shell, alors
que le réseau est sain.

Coupables usuels, dans l'ordre :

1. **Résolution DNS inverse côté serveur** : avec `UseDNS yes` (défaut
   historique), sshd résout l'IP du client avant d'accepter. Si le DNS est
   lent/absent → attente. Remède :

```
# /etc/ssh/sshd_config.d/30-limits.conf
UseDNS no
```

   (Sauf si tu utilises `from=` avec des noms DNS ou `AllowUsers user@host` —
   dans ce cas, soigne ton DNS plutôt que de couper.)

2. **GSSAPI côté client** : le client tente une auth Kerberos qui timeout.
   Remède client :

```
Host *
    GSSAPIAuthentication no
```

3. **IPv6 qui timeout** avant fallback IPv4 : `AddressFamily inet` côté client
   pour forcer IPv4 (diagnostic, pas une solution durable).

Mesure : `time ssh -o GSSAPIAuthentication=no web1 'exit'` avant/après pour
objectiver le gain.

## 67. Cas 6 : `Too many authentication failures`

```
Received disconnect from 192.0.2.10 port 22:2: Too many authentication failures
```

Le client a proposé **plus de clés** que `MaxAuthTries` (défaut 6) n'en
autorise : typique quand l'agent contient 10 clés et que la bonne est la 8e.
Le serveur coupe avant même d'essayer la bonne.

Remèdes :

```bash
# Ponctuel : forcer LA clé
client$ ssh -o IdentitiesOnly=yes -i ~/.ssh/id_ed25519_web web1
```

```
# Durable : dans ~/.ssh/config
Host web1
    IdentityFile ~/.ssh/id_ed25519_web
    IdentitiesOnly yes
```

- `IdentitiesOnly yes` limite aux `IdentityFile` déclarés (+ ceux de la ligne
  de commande), sans piocher dans l'agent.
- Alternative : vider l'agent des clés inutiles (`ssh-add -d`), ou augmenter
  `MaxAuthTries` côté serveur (mauvaise idée : ça aide aussi les attaquants).

## 68. Cas 7 : `WARNING: UNPROTECTED PRIVATE KEY FILE`

```
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@         WARNING: UNPROTECTED PRIVATE KEY FILE!          @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
Permissions 0644 for '/home/zelef/.ssh/id_ed25519' are too open.
```

La clé privée est lisible par d'autres → SSH **refuse de l'utiliser**.
Souvent après une copie via clé USB, un `umask` exotique, ou une restauration
de sauvegarde.

```bash
client$ chmod 600 ~/.ssh/id_ed25519
client$ chmod 700 ~/.ssh
client$ chmod 644 ~/.ssh/id_ed25519.pub ~/.ssh/known_hosts
client$ chmod 600 ~/.ssh/config
```

Règle mémo : `700` le répertoire, `600` les privés et la config, `644` les
publiques et `known_hosts`. Si le fichier vient d'un support non fiable
(USB trouvée, pièce jointe), **considère la clé comme compromise** : régénère
(section 58), ne te contente pas de `chmod`.

## 69. Cas 8 : `ssh-copy-id` échoue

```
ERROR: failed to open ID file '/home/zelef/.ssh/id_ed25519.pub': No such file
```

