---
id: collect-261001-general-networking/general-networking/tutoriel-tailscale-1-96-vpn-mesh-en-13-etapes-2026-3
title: "Installation rapide (Debian, Ubuntu, Fedora, Arch, etc.)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-tailscale-1-96-vpn-mesh-en-13-etapes-2026.md
source_anchor: ""
source_lines: [168, 308]
sha256: 09d2e51dca720d5f2c6de1ff316ae0928e68f0724dd1824fd5f15500549016c8
---

# Installation rapide (Debian, Ubuntu, Fedora, Arch, etc.)

```
# Sur le serveur cible
sudo tailscale up --ssh
# Sur le client
tailscale ssh root@ubuntu-server
# Avec un utilisateur précis
tailscale ssh deploy@ubuntu-server
```
Tailscale SSH génère et signe à la volée des certificats SSH éphémères, journalise chaque session avec l’identité IdP de l’utilisateur (audit AuditD, JournalD et KAuditD intégrés depuis Tailscale 1.94) et applique les ACL définies dans la console. Pour exiger une réauthentification IdP avant chaque session sur un serveur sensible, ajoutez la règle *check* dans l’ACL :

```
{
  "ssh": [
    {
      "action": "check",
      "src": ["autogroup:member"],
      "dst": ["tag:prod"],
      "users": ["root", "deploy"],
      "checkPeriod": "1h"
    }
  ]
}
```
Avec `"action": "check"` et `"checkPeriod": "1h"`, l’utilisateur doit revalider son identité IdP toutes les heures. C’est la posture recommandée pour les serveurs de production hébergeant des données personnelles soumises au RGPD.

## Étape 8 : Écrire ses premières ACL JSON

Par défaut, Tailscale autorise tout entre tous les nœuds d’un même tailnet (politique *allow all*). En production, on bascule rapidement vers une matrice de permissions explicite. Voici un exemple complet pour un tailnet de PME segmenté en trois rôles : développeurs, opérateurs et serveurs de production.

```
{
  "groups": {
    "group:devs":  ["[email protected]", "[email protected]"],
    "group:ops":   ["[email protected]"]
  },
  "tagOwners": {
    "tag:prod":    ["group:ops"],
    "tag:staging": ["group:ops", "group:devs"]
  },
  "acls": [
    { "action": "accept", "src": ["group:ops"],  "dst": ["*:*"] },
    { "action": "accept", "src": ["group:devs"], "dst": ["tag:staging:*"] },
    { "action": "accept", "src": ["group:devs"], "dst": ["tag:prod:80,443"] },
    { "action": "accept", "src": ["tag:prod"],   "dst": ["tag:prod:*"] }
  ],
  "ssh": [
    { "action": "accept", "src": ["group:ops"],  "dst": ["tag:prod"], "users": ["root"] },
    { "action": "check",  "src": ["group:devs"], "dst": ["tag:staging"], "users": ["deploy"], "checkPeriod": "8h" }
  ]
}
```
Quelques principes à retenir : les **tags** (préfixés par `tag:`) sont attribués aux machines via `tailscale up --advertise-tags=tag:prod` ; les **groupes** regroupent les utilisateurs ; et les **règles ACL** sont évaluées en mode *default deny* dès qu’au moins une règle est définie. Le plan Personal limite à 3 groupes ACL ; le plan Starter monte à 10 ; le plan Premium les rend illimités.

Validez toujours la politique avec `tailscale netcheck` et le simulateur ACL intégré à la console (*Access controls → Preview*) avant de la déployer. Une ACL mal écrite peut bannir l’administrateur de son propre tailnet.

## Étape 9 : Générer des auth keys pour automatiser

Pour enrôler des nœuds non-interactifs (CI/CD runners, conteneurs Docker, instances Terraform), on utilise des *auth keys*. Le provider Terraform officiel, passé en **version 0.18.0 dès février 2025**, avait déjà posé les bases de cette automatisation, renforcée ensuite par la *Fall Update Week* d’octobre 2025 qui a ciblé toute la branche stable **1.92**, puis, depuis la version **v1.92.5** livrée le 23 janvier 2026, par la prise en charge des identités fédérées de bout en bout sur l’API, le client Go et ce même provider Terraform — ce qui simplifie l’automatisation multi-cloud sans jonglage manuel de clés. Quatre variantes existent dans la console (*Settings → Keys*) :

- **Reusable** : la clé peut enrôler plusieurs nœuds, idéale pour un parc homogène.
- **One-off** : usage unique, à privilégier pour un nœud sensible.
- **Pre-authorized** : le nœud est immédiatement autorisé sans approbation manuelle.
- **Ephemeral** : le nœud disparaît du tailnet à la déconnexion (parfait pour CI/CD).

```
# Exemple d'enrôlement non-interactif (CI/CD)
sudo tailscale up \
  --authkey=tskey-auth-kxxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxx \
  --hostname=ci-runner-${CI_JOB_ID} \
  --advertise-tags=tag:ci \
  --ephemeral=true \
  --ssh=false
```
Les clés ont une durée de vie maximale de 90 jours, ramenée à 30 jours pour les nœuds tagués comme `tag:prod`. Combinez-les avec le secret manager de votre pipeline (GitHub Actions secrets, GitLab CI variables protégées, HashiCorp Vault) — ne les commitez jamais en clair.

## Étape 10 : Déployer Tailscale dans Docker

L’image officielle `tailscale/tailscale` permet de containeriser un nœud, par exemple pour exposer un service interne à votre tailnet sans modifier l’hôte. Voici un fichier `compose.yaml` minimal qui place un serveur web Caddy derrière Tailscale :

```
services:
  tailscale:
    image: tailscale/tailscale:v1.96.5
    container_name: ts-caddy
    hostname: caddy
    environment:
      - TS_AUTHKEY=${TS_AUTHKEY}
      - TS_EXTRA_ARGS=--advertise-tags=tag:web
      - TS_STATE_DIR=/var/lib/tailscale
      - TS_SERVE_CONFIG=/config/serve.json
    volumes:
      - ./tailscale-state:/var/lib/tailscale
      - ./serve.json:/config/serve.json:ro
      - /dev/net/tun:/dev/net/tun
    cap_add:
      - net_admin
      - sys_module
    restart: unless-stopped
  caddy:
    image: caddy:2.8-alpine
    network_mode: service:tailscale
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile:ro
    restart: unless-stopped
```
Le pattern `network_mode: service:tailscale` partage la pile réseau du conteneur Tailscale avec celui de Caddy. Aucun port n’est publié sur l’hôte : le service est uniquement joignable à l’intérieur du tailnet, sous l’hostname `caddy.monequipe.ts.net`. Lancez avec :

```
echo "TS_AUTHKEY=tskey-auth-..." > .env
docker compose up -d
docker compose logs -f tailscale
```
## Étape 11 : Exposer un service public avec Tailscale Funnel

**Tailscale Funnel** permet de publier un service local sur Internet en HTTPS, sans ouvrir de port sur votre routeur ni configurer Cloudflare Tunnel ou Ngrok. Tailscale termine le TLS sur ses propres serveurs DERP et tunnelise le trafic chiffré jusqu’à votre nœud. La fonctionnalité est gratuite jusqu’à trois services concurrents par tailnet.

Activez d’abord Funnel pour votre nœud dans la console (*Funnel → Edit nodes attribute*), puis exécutez :

```
# Servir un site local sur le port 8080 en HTTPS public
sudo tailscale serve https / http://localhost:8080
# Activer Funnel (exposition publique)
sudo tailscale funnel 443 on
# Statut
tailscale funnel status
# https://caddy.monequipe.ts.net (Funnel on)
# |-- proxy http://127.0.0.1:8080
```
Le certificat TLS est délivré automatiquement par Let’s Encrypt via le service Tailscale et renouvelé en arrière-plan. Pour désactiver l’exposition publique tout en gardant le service joignable à l’intérieur du tailnet : `sudo tailscale funnel 443 off`. Pour aller plus loin sur la terminaison HTTPS, voir notre tutoriel Traefik reverse proxy.

## Étape 12 : Audit, monitoring et logs Tailscale

Tout déploiement sérieux exige un volet observabilité. Tailscale offre quatre canaux complémentaires pour superviser un tailnet :

- **Configuration audit log** : trace toutes les modifications ACL, ajouts/retraits de nœuds (gratuit, plan Personal).
- **Network flow log** : journal détaillé de toutes les connexions tailnet, exporté vers S3 ou syslog (plan Premium).
- **SSH session recording** : enregistrement intégral des sessions Tailscale SSH (plan Premium).
- **Métriques Prometheus** : exposées en local par`tailscaled` sur l’endpoint`/metrics` .

Pour activer les métriques Prometheus :

