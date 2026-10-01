---
id: collect-261001-general-networking/general-networking/tutoriel-tailscale-1-96-vpn-mesh-en-13-etapes-2026-4
title: "Installation rapide (Debian, Ubuntu, Fedora, Arch, etc.)"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Intel", "Microsoft", "OpenAI", "Stripe"]
dates: []
keywords: ["intel", "open source", "pricing"]
source: docs/RAG/collect-261001-general-networking/tutoriel-tailscale-1-96-vpn-mesh-en-13-etapes-2026.md
source_anchor: ""
source_lines: [309, 377]
sha256: d461e1369c9fb4b0ca1c75ae47245668127148d4098ea749c5eafb4d8ddc5f6b
---

# Installation rapide (Debian, Ubuntu, Fedora, Arch, etc.)

```
sudo tailscale set --report-posture-identity=all
sudo systemctl edit tailscaled
# Ajouter dans la section [Service]
# Environment="FLAGS=--port=41641 --tun=tailscale0 --metrics-port=9091"
sudo systemctl restart tailscaled
curl http://localhost:9091/metrics | grep tailscaled_
```
Branchez ensuite ce endpoint sur votre instance Prometheus + Grafana. Les indicateurs clés à surveiller : `tailscaled_inbound_bytes_total`, `tailscaled_outbound_dropped_packets_total`, et le nouveau `tailscaled_home_derp_region_id` introduit en v1.96.5 qui révèle la région DERP utilisée comme fallback. Depuis la mise à jour d’avril 2026, la rétention du journal de flux réseau est devenue configurable et son export vers S3 a été enrichi, tandis qu’un accès API-only au tailnet via des clients OAuth a été lancé pour s’intégrer plus proprement aux pipelines d’observabilité sans compte utilisateur dédié.

## Étape 13 : Auto-héberger avec Headscale

Si la dépendance au plan de contrôle SaaS de Tailscale Inc. vous gêne — souveraineté numérique, conformité air-gap, conviction libriste — la communauté maintient **Headscale**, une réimplémentation open source en Go du serveur de coordination Tailscale, distribuée sous licence BSD-3-Clause. Headscale est compatible avec tous les clients officiels Tailscale (Linux, macOS, Windows, iOS, Android) et n’affecte pas les performances réseau (le trafic reste en pair-à-pair WireGuard).

```
# Installation Headscale via paquet .deb
HEADSCALE_VERSION="0.26.1"
wget https://github.com/juanfont/headscale/releases/download/v${HEADSCALE_VERSION}/headscale_${HEADSCALE_VERSION}_linux_amd64.deb
sudo dpkg -i headscale_${HEADSCALE_VERSION}_linux_amd64.deb
# Configuration minimale
sudo nano /etc/headscale/config.yaml
# server_url: https://headscale.exemple.com
# listen_addr: 0.0.0.0:8080
# private_key_path: /var/lib/headscale/private.key
sudo systemctl enable --now headscale
# Créer un utilisateur et une pré-auth key
sudo headscale users create alice
sudo headscale preauthkeys create --user alice --reusable --expiration 24h
```
Côté client, on bascule sur le serveur Headscale en spécifiant son URL :

```
sudo tailscale up \
  --login-server=https://headscale.exemple.com \
  --authkey=<preauthkey>
```
Headscale couvre désormais (version 0.26) les ACL, les routes annoncées, MagicDNS et l’OIDC, mais Tailscale Funnel et Tailscale SSH avec audit AuditD restent l’apanage du SaaS officiel. Pour un homelab personnel, Headscale offre 100 % de l’expérience utile, sans compteur de devices et sans dépendance au cloud.

## Pièges fréquents et résolutions

Voici huit erreurs courantes rencontrées en production avec Tailscale et les manières de les corriger rapidement — sans oublier la prudence de rigueur lors des montées de version majeures, comme le rappelle l’épisode de juillet 2025 où Tailscale avait dû suspendre le déploiement de la v1.86.0, d’abord sur macOS le 25 juillet puis sur l’ensemble des plateformes le 28 juillet, après la découverte de régressions en production.

- **Connexion bloquée en relay DERP** : si`tailscale status` affiche`relay "fra"` au lieu de`direct` , votre pare-feu bloque UDP/41641. Ouvrez ce port en sortant pour récupérer le débit P2P.
- **MagicDNS ne résout pas les noms** : les containers Docker ignorent souvent`/etc/resolv.conf` de l’hôte. Forcez`--dns 100.100.100.100` dans la config Docker ou systemd-resolved.
- **Subnet route non visible** : sans approbation manuelle dans la console (*Edit route settings* ), la route reste invisible. Approuvez systématiquement et activez l’option*Auto-approve routes* via les ACL pour les sous-réseaux de confiance.
- **Conflit IP avec le LAN** : Tailscale utilise`100.64.0.0/10` (CGNAT). Si votre opérateur vous a déjà attribué une IP CGNAT, désactivez`--accept-routes` ou changez la plage côté FAI.
- **Erreur “tailscale: command not found”** après installation : le démon est peut-être installé mais le PATH n’a pas été rafraîchi. Faites`hash -r` ou ouvrez un nouveau shell.
- **Exit node trop lent** : un Raspberry Pi 4 plafonne autour de 250 Mbit/s en exit node à cause de l’absence d’AES-NI matériel. Préférez un Pi 5 (1 Gbit/s atteint) ou un mini-PC x86 type Intel N100.
- **Tailscale up bloque l’IPv6** : sur certains hôtes,`tailscaled` capture`::1` . Désactivez avec`sudo tailscale up --netfilter-mode=off` si votre stack dépend d’IPv6 local.
- **Auth key expirée silencieusement** : les clés non-éphémères durent 90 jours max. Programmez une rotation automatique via cron ou Ansible avec notre tutoriel Ansible.

## Tarification et limites des plans Tailscale en 2026

Tailscale a simplifié sa grille tarifaire en début 2025, avant de la faire évoluer à nouveau courant 2026 : le plan Starter a été renommé Standard et repositionné à **8 $/utilisateur/mois**, tandis que le plan Premium reste à **18 $/utilisateur/mois** avec **10 000 minutes/mois** de ressources éphémères. Voici la structure constatée en août 2026 selon ModSignal : le plan Personal demeure gratuit pour 6 utilisateurs maximum, avec **50 ressources taguées** et **1 000 minutes/mois** de quota éphémère, un socle validé également sur la page *tailscale.com/pricing*.

| Plan | Prix | Utilisateurs | Devices | ACL groups | Ephemeral min/mois | 
|---|---|---|---|---|---|
| Personal | 0 $ | 6 max | Illimité | 3 | 1 000 | 
| Standard | 8 $/user/mois | 3 inclus | Illimité | 10 | 3 000 | 
| Premium | 18 $/user/mois | Illimité | Illimité | Illimité | 10 000 | 
| Enterprise | Sur devis | Illimité | Illimité | Illimité | Illimité | 
| Headscale (DIY) | 0 € | Illimité | Illimité | Illimité | N/A | 

Le seuil de bascule économique se situe vers 8 utilisateurs payants : à 8 $ × 8 = 64 $/mois sur le nouveau tarif Standard d’août 2026, l’écart avec un VPS Hetzner à 6 €/mois pour héberger Headscale se resserre, ce qui rend l’arbitrage plus sensible à la valeur du temps administrateur qu’auparavant. Au-delà de 30 utilisateurs, le passage au plan Premium devient nécessaire pour bénéficier des audits SSH et du flow log exhaustif. Pour les grands comptes, le plan Enterprise a évolué sur deux points en 2026 : le provisionnement SCIM avec les principaux fournisseurs d’identité (Okta, Microsoft Entra ID) est disponible depuis le 19 juin 2026, et la facturation est passée à un modèle à l’usage indexé sur le nombre d’utilisateurs actifs mensuels (MAU) depuis juillet 2026, remplaçant les forfaits fixes historiques.

## Astuces avancées pour environnements production

Quelques bonnes pratiques observées chez les clients Tailscale les plus matures (Stripe, OpenAI, GitLab et plusieurs administrations européennes) :

