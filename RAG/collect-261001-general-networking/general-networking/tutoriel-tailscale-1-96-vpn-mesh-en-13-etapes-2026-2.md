---
id: collect-261001-general-networking/general-networking/tutoriel-tailscale-1-96-vpn-mesh-en-13-etapes-2026-2
title: "Installation rapide (Debian, Ubuntu, Fedora, Arch, etc.)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "distribution"]
source: docs/RAG/collect-261001-general-networking/tutoriel-tailscale-1-96-vpn-mesh-en-13-etapes-2026.md
source_anchor: ""
source_lines: [50, 167]
sha256: 75898f6111c9525af38e7dbdaca5760a7ceb7847089cca80e77bf455a7445e60
---

# Installation rapide (Debian, Ubuntu, Fedora, Arch, etc.)

Lors de la première connexion, votre tailnet reçoit un nom DNS unique de la forme `tail-xxxx.ts.net`. Vous pouvez ensuite le renommer dans *Settings → General* en quelque chose de plus mémorable, par exemple `monequipe.ts.net`. Ce suffixe servira plus tard à MagicDNS pour résoudre les noms d’hôtes du tailnet sans configuration DNS additionnelle.

Activez immédiatement deux options dans *Settings → User management* : **2-Step verification** (obligatoire pour toute organisation) et **Device approval**, qui empêche l’enrôlement automatique de nouveaux nœuds tant qu’un administrateur n’a pas validé leur arrivée. Cela neutralise la principale classe d’attaque sur Tailscale, à savoir le détournement d’un compte SSO.

## Étape 2 : Installer Tailscale sur Debian et Ubuntu

Le moyen le plus rapide consiste à passer par le script officiel, qui détecte la distribution et configure le dépôt apt automatiquement. C’est l’approche recommandée pour les machines de développement et le homelab.

```
# Installation rapide (Debian, Ubuntu, Fedora, Arch, etc.)
curl -fsSL https://tailscale.com/install.sh | sh
# Vérifier la version installée
tailscale version
# Attendu : 1.96.5
```
Pour la production, on préfère ajouter le dépôt manuellement afin de pouvoir épingler une version précise et auditer les mises à jour. La séquence ci-dessous installe la dernière build stable sur Ubuntu 24.04 :

```
# Importer la clé GPG signée par Tailscale Inc.
curl -fsSL https://pkgs.tailscale.com/stable/ubuntu/noble.noarmor.gpg | \
  sudo tee /usr/share/keyrings/tailscale-archive-keyring.gpg >/dev/null
# Ajouter le dépôt apt
echo "deb [signed-by=/usr/share/keyrings/tailscale-archive-keyring.gpg] \
  https://pkgs.tailscale.com/stable/ubuntu noble main" | \
  sudo tee /etc/apt/sources.list.d/tailscale.list
# Installer
sudo apt update
sudo apt install tailscale
# Démarrer le démon
sudo systemctl enable --now tailscaled
sudo systemctl status tailscaled
```
Sur Debian 12, remplacez `noble` par `bookworm`. Pour Fedora 40+, le dépôt est `https://pkgs.tailscale.com/stable/fedora/tailscale.repo`. Sur Synology DSM 7.2, le client est disponible directement depuis le centre de paquets.

## Étape 3 : Joindre le tailnet avec tailscale up

Une fois le démon `tailscaled` lancé, on rejoint le tailnet avec une seule commande, qui ouvre une URL d’authentification dans votre navigateur :

```
sudo tailscale up
# Sortie attendue
To authenticate, visit:
   https://login.tailscale.com/a/abcd1234ef56
Success.
```
Cliquez sur le lien, validez sur la page web (vous serez automatiquement connecté à votre IdP), puis revenez au terminal. La machine reçoit immédiatement une adresse IP dans la plage `100.64.0.0/10` (CGNAT), réservée par l’IETF aux opérateurs et donc utilisable sans risque de collision avec votre LAN. Vérifiez avec :

```
tailscale ip -4
# Exemple : 100.115.92.7
tailscale status
# 100.115.92.7    debian-laptop      vous@      linux   -
# 100.83.45.12    ubuntu-server      vous@      linux   active; direct 89.x.x.x:41641, tx 4.2k rx 8.1k
```
Répétez les étapes 2 et 3 sur la deuxième machine. En moins de cinq minutes, vous avez désormais deux nœuds maillés, joignables directement par leur IP `100.x.y.z`. Testez avec `ping 100.83.45.12` ou `tailscale ping ubuntu-server` : la deuxième forme affiche également le chemin emprunté (direct ou via DERP).

## Étape 4 : Activer MagicDNS pour les noms d’hôtes

Mémoriser des IP en `100.x.y.z` n’est pas tenable au-delà de quelques nœuds. Tailscale propose **MagicDNS**, qui résout automatiquement le nom d’hôte de chaque machine du tailnet. Activez-le dans la console : *DNS → Enable MagicDNS*. La fonctionnalité est gratuite à partir du plan Personal.

Ensuite, sur chaque machine, vous pouvez désormais faire :

```
# SSH sur le serveur Ubuntu sans connaître son IP
ssh user@ubuntu-server
# Avec FQDN complet
ssh [email protected]
# Vérification DNS
tailscale dns status
# magicdns enabled, resolvers: 100.100.100.100
```
L’adresse spéciale `100.100.100.100` est le résolveur DNS interne de Tailscale, écouté localement par le démon. Vous pouvez aussi pousser des serveurs DNS personnalisés via la section *DNS → Nameservers*, ce qui permet par exemple de combiner Tailscale avec un AdGuard Home auto-hébergé pour bloquer la publicité sur tout le tailnet.

## Étape 5 : Configurer un nœud de sortie (exit node)

Un *exit node* est un nœud du tailnet qui peut router tout le trafic Internet d’autres nœuds. Cas d’usage typiques : se connecter en wifi public d’aéroport en ressortant chez soi, contourner un blocage géographique en passant par un VPS européen, ou auditer les connexions sortantes d’un poste de développement.

Sur le nœud destiné à servir de sortie (par exemple une VM dans un datacenter européen ou un Raspberry Pi à la maison), exécutez :

```
# Activer l'IP forwarding au niveau noyau
echo 'net.ipv4.ip_forward = 1' | sudo tee -a /etc/sysctl.d/99-tailscale.conf
echo 'net.ipv6.conf.all.forwarding = 1' | sudo tee -a /etc/sysctl.d/99-tailscale.conf
sudo sysctl -p /etc/sysctl.d/99-tailscale.conf
# Annoncer le rôle d'exit node
sudo tailscale up --advertise-exit-node
# Approuver dans la console : Machines → ... → Edit route settings
```
Côté client, pour utiliser cet exit node :

```
sudo tailscale up --exit-node=ubuntu-server --exit-node-allow-lan-access=true
# Vérifier que tout le trafic sort par l'exit node
curl ifconfig.me
# Doit afficher l'IP publique de l'exit node, pas la vôtre
```
L’option `--exit-node-allow-lan-access=true` est cruciale : sans elle, votre imprimante ou votre NAS local devient injoignable. Pour repasser en routage normal, exécutez `sudo tailscale up --exit-node=` (chaîne vide).

## Étape 6 : Annoncer un sous-réseau (subnet routes)

Là où l’exit node fait sortir tout le trafic Internet, le *subnet router* expose un sous-réseau interne aux autres nœuds du tailnet. C’est la fonctionnalité reine pour accéder à un homelab depuis la route, ou pour qu’un tailnet d’entreprise atteigne le LAN d’un site distant sans installer Tailscale sur chaque appareil (caméra IP, imprimante réseau, automate industriel).

```
# Sur la machine du LAN à exposer (par ex. 192.168.1.0/24)
sudo tailscale up --advertise-routes=192.168.1.0/24
# Multi-routes (LAN + sous-réseau Docker)
sudo tailscale up --advertise-routes=192.168.1.0/24,172.17.0.0/16
# Vérifier que les routes sont bien annoncées
tailscale status --self=true
```
Les routes annoncées doivent ensuite être **approuvées dans la console d’administration** : *Machines → … → Edit route settings → Approve all*. Sans cette validation, les autres nœuds ne reçoivent pas l’information de routage. Ce double consentement est volontaire : il évite qu’une machine compromise ne s’auto-déclare passerelle vers un sous-réseau sensible.

Astuce production : couplez ce subnet router avec un pare-feu OPNsense ou pfSense en amont, qui filtre granulairement le trafic en provenance du tailnet vers le LAN. Cette segmentation supplémentaire est imposée par l’ANSSI dans son guide d’hygiène informatique pour les opérateurs d’importance vitale.

## Étape 7 : Activer Tailscale SSH

Tailscale SSH remplace la gestion classique des clés SSH (autorisations dans `~/.ssh/authorized_keys`) par l’identité Tailscale du tailnet. Concrètement, vos administrateurs n’ont plus à distribuer ou révoquer manuellement des clés publiques : tout est piloté par les ACL JSON.

