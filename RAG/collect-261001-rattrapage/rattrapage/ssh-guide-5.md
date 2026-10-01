---
id: collect-261001-rattrapage/rattrapage/ssh-guide-5
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [845, 1078]
sha256: 870bef7fa98c7f0576252a299581bbbf521c7b5f38461a946168b6245e7f9b71
---

# Guide SSH approfondi

```
Host srv-via-vieux-bastion
    HostName 10.40.0.9
    ProxyCommand ssh -W %h:%p vieux-bastion
```

`ssh -W %h:%p` ouvre un simple relais TCP via le bastion : c'est exactement ce
que `ProxyJump` fait en interne, en plus simple. `ProxyCommand` avec `nc` ou
des wrappers est à éviter sauf besoin réel (chaque maillon ajoute de la
fragilité).

Conseil d'architecture : **un bastion par zone réseau**, pas un bastion
mondial. Chaque saut supplémentaire = latence + point de panne. Au-delà de
deux sauts, interroge le design réseau.

## 28. Multiplexage : `ControlMaster`, `ControlPath`, `ControlPersist`

Le multiplexage réutilise **une seule connexion TCP** pour plusieurs sessions
vers le même hôte : le 2e `ssh`, le `scp`, le `rsync` passent par le tunnel
déjà établi. Résultat : connexion quasi instantanée, une seule authentification.

```
Host *
    ControlMaster auto
    ControlPath ~/.ssh/sockets/%r@%h:%p
    ControlPersist 10m
```

- `ControlMaster auto` : la première connexion devient maîtresse, les
  suivantes s'y rattachent automatiquement.
- `ControlPath` : chemin du socket de contrôle. `%r` = user, `%h` = hôte,
  `%p` = port. **Crée le répertoire** (`mkdir -p ~/.ssh/sockets`) sinon rien
  ne fonctionne, silencieusement.
- `ControlPersist 10m` : la connexion maîtresse survit 10 minutes après la
  fermeture du dernier client → les connexions suivantes restent instantanées
  même si tu as tout fermé.

Gestion manuelle :

```bash
client$ ssh -O check web1        # la connexion maîtresse est-elle active ?
client$ ssh -O stop web1         # fermer proprement la maîtresse
client$ ssh -O exit web1         # variante : ferme aussi les sessions actives
client$ ssh -S ~/.ssh/sockets/zelef@web1:22 -O check web1
```

Points d'attention :

- Ne pas multiplexer vers des hôtes **non fiables** : toutes les sessions
  partagent le même niveau de confiance.
- `ControlPersist` long sur un portable = session qui survit au changement de
  réseau et freeze (section 64). `10m` est un bon compromis ; `yes` (illimité)
  seulement sur poste fixe.
- Avec `ProxyJump`, le multiplexage s'applique aussi au saut : gros gain sur
  les chaînes de bastions.

## 29. Mesurer le gain du multiplexage

```bash
client$ time ssh web1 'exit'          # sans multiplexage (1re connexion)
real    0m1.842s
client$ time ssh web1 'exit'          # avec multiplexage (maîtresse active)
real    0m0.041s
```

Ordre de grandeur typique : **10 à 40× plus rapide** sur les connexions
suivantes, surtout à travers un bastion (2 handshakes évités) ou sur un lien
avec de la latence (chaque aller-retour du handshake coûte).

Test propre :

```bash
client$ for i in 1 2 3; do /usr/bin/time -f "%es" ssh -o ControlMaster=no web1 'exit'; done
client$ ssh -fN web1   # ouvre une maîtresse en fond (avec ControlPersist)
client$ for i in 1 2 3; do /usr/bin/time -f "%es" ssh web1 'exit'; done
```

En pratique d'équipe : Ansible, scripts de déploiement et `for` loops sur 50
machines passent de plusieurs minutes à quelques secondes. C'est le réglage
client au **meilleur ratio effort/gain** de tout ce guide.

## 30. Tunnels locaux (`-L`)

`-L [bind:]port_local:hôte_distant:port_distant` : un port ouvert **sur ton
poste** est relayé vers une destination **vue depuis le serveur SSH**.

```bash
# Le port 8080 local -> interface web d'admin du serveur (vue depuis le serveur)
client$ ssh -L 8080:localhost:9090 web1
# Puis http://localhost:8080 dans le navigateur
```

```
[ton poste]:8080 ---(chiffré via SSH)---> [web1] ---> localhost:9090 (vu par web1)
```

Variantes :

```bash
# Atteindre une autre machine du réseau privé via le serveur SSH
client$ ssh -L 3306:db-interne:3306 bastion
# Le tunnel n'écoute que sur 127.0.0.1 par défaut (bien). Pour l'ouvrir au LAN :
client$ ssh -L 0.0.0.0:8080:localhost:9090 web1     # nécessite GatewayPorts=yes côté serveur
```

En config :

```
Host tunnel-monitoring
    HostName web1
    LocalForward 8080 localhost:9090
    # -N : pas de shell, -f : en arrière-plan
```

```bash
client$ ssh -N -f tunnel-monitoring
```

Cas d'usage rois : accéder à une UI web interne (section 33), à une base de
données (MySQL/PostgreSQL), à un Elasticsearch/Kibana sans les exposer.

## 31. Tunnels distants (`-R`)

`-R [bind:]port_distant:localhost:port_local` : l'inverse — un port ouvert
**sur le serveur SSH** est relayé vers **ton poste** (ou ton réseau).

```bash
# Exposer ton serveur web local (port 8000) sur le port 9000 du serveur distant
client$ ssh -R 9000:localhost:8000 web1
# Depuis web1 : curl http://localhost:9000 -> atteint ton poste
```

```
[web1]:9000 ---(chiffré via SSH)---> [ton poste] ---> localhost:8000
```

Côté serveur, `sshd_config` contrôle ce qui est autorisé :

```
# Autoriser les -R mais seulement sur loopback (défaut sûr)
GatewayPorts no
# Pour exposer sur toutes les interfaces (à manier avec précaution) :
# GatewayPorts yes
# Pour autoriser seulement certains ports :
PermitListen 9000 9001
```

**Sécurité** : un `-R` avec `GatewayPorts yes` expose ton service local à tout
le réseau du serveur distant (voire Internet). Vérifie toujours sur quelle
interface le port distant écoute (`ss -tlnp` côté serveur). En équipe,
préfère `PermitListen` avec une liste blanche de ports.

Cas d'usage : donner à un collègue/serveur distant l'accès temporaire à un
service de ton poste, webhook de test, démo.

## 32. Tunnels dynamiques (`-D`, SOCKS)

`-D port_local` : ouvre un **proxy SOCKS** sur ton poste. Tout le trafic que
tu y envoies est relayé par le serveur SSH, avec résolution DNS côté serveur
si le client le demande.

```bash
client$ ssh -D 1080 bastion
# Firefox : proxy SOCKS5 manuel -> 127.0.0.1:1080, "Proxy DNS via SOCKS5" coché
# curl :
client$ curl --socks5-hostname 127.0.0.1:1080 http://intranet.local/
```

En config :

```
Host socks-bastion
    HostName bastion
    DynamicForward 1080
```

Différences avec `-L` :

| | `-L` (local) | `-D` (dynamique) |
|---|---|---|
| Destination | fixée à l'ouverture | choisie par connexion |
| Protocole | TCP brut | SOCKS4/5 |
| Usage | 1 service précis | navigation / multi-services |

Attention : le proxy SOCKS ne chiffre que **jusqu'au serveur SSH** ; au-delà,
c'est le protocole applicatif qui compte (HTTPS reste HTTPS, HTTP reste
clair sur le LAN distant). Et tout ton trafic navigateur passe par le serveur
distant : à réserver aux besoins pro, pas au surf perso.

## 33. Cas concret : accéder à une UI web interne

Situation : Proxmox sur `pve1` (10.10.1.5:8006), Grafana sur `mon1`
(10.10.1.20:3000), accessibles uniquement depuis le LAN, bastion en entrée.

```
Host bastion
    HostName bastion.example.com
    User zelef

Host pve1
    HostName 10.10.1.5
    User root
    ProxyJump bastion
    LocalForward 8006 localhost:8006

Host mon1
    HostName 10.10.1.20
    User zelef
    ProxyJump bastion
    LocalForward 3000 localhost:3000
```

```bash
client$ ssh -N -f pve1 && ssh -N -f mon1
# https://localhost:8006 -> Proxmox (accepte le certificat autosigné, normal)
# http://localhost:3000  -> Grafana
```

Pourquoi c'est mieux que d'exposer les UI : zéro port ouvert sur Internet,
authentification SSH (clé + éventuellement 2FA) **avant** même d'atteindre
l'applicatif, chiffrement de bout en bout. C'est le pattern standard pour
iDRAC/iLO, interfaces d'onduleurs (lien avec ton guide UPS), switchs web.

Astuce : `-N -f` + `ExitOnForwardFailure yes` dans la config pour que SSH
échoue franchement si le port local est déjà pris, plutôt que de tourner sans
tunnel (section 70).

## 34. Cas concret : exposer un service local vers l'extérieur

Situation : tu développes un webhook sur `localhost:5000`, un partenaire doit
le tester depuis Internet, via ton serveur `web1` public.

Côté serveur (`/etc/ssh/sshd_config.d/tunnels.conf`) :

