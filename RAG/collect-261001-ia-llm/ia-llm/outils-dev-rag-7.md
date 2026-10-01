---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-7
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: []
dates: ["2026-09-19"]
keywords: ["apache", "open source"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [1214, 1420]
sha256: 3182dd4579f03a66abaa738635c6d17606512ecbc81efaf463320bf0cd1ad478
---

# Outils dev + ingénierie RAG (chunk & corpus)

```bash
# Un serveur local tourne sur le port 8000 (ex. : uvicorn, python -m http.server)
ngrok http 8000
```

Sortie typique (vérifiée) :

```text
Forwarding  https://abcd1234.ngrok-free.app -> http://localhost:8000
```

- `https://abcd1234.ngrok-free.app` est accessible **depuis Internet**.
- Dashboard local d'inspection du trafic : http://127.0.0.1:4040
  (rejoue les requêtes, vois les headers — idéal pour débugger un webhook).
- `curl http://127.0.0.1:4040/api/tunnels` : liste les tunnels en JSON.

Autres protocoles :

```bash
ngrok http 3000 --domain=demo-perso.ngrok-free.app  # domaine réservé (payant pour custom)
ngrok tcp 22                                        # SSH — ⚠ CB requise en free (voir §38)
ngrok http http://192.168.1.50:8080                 # service sur une autre machine du LAN
```

Config persistante `~/.config/ngrok/ngrok.yml` :

```yaml
version: "2"
tunnels:
  api-rag:
    proto: http
    addr: 8000
```

```bash
ngrok start api-rag        # démarre le tunnel nommé
ngrok start --all          # tous les tunnels du fichier
```

## 37. Service systemd : tunnel permanent (optionnel)

`/etc/systemd/system/ngrok.service` :

```ini
[Unit]
Description=Tunnel ngrok API RAG
After=network-online.target
Wants=network-online.target

[Service]
ExecStart=/usr/local/bin/ngrok start api-rag --config /home/<utilisateur>/.config/ngrok/ngrok.yml
Restart=on-failure
RestartSec=10
User=<utilisateur>  # remplace par ton login Linux

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now ngrok.service
systemctl status ngrok.service
```

> ⚠ **Réfléchis avant de rendre un tunnel permanent** : chaque service
> exposé est une surface d'attaque. Voir §39.

## 38. Limites de l'offre gratuite ngrok (vérifié sept 2026)

Chiffres vérifiés sur plusieurs sources à jour (19/09/2026) :

| | Free ($0) | Hobbyist ($8/mois annuel, $10 mensuel) | Pay-as-you-go ($20/mois + usage) |
|---|---|---|---|
| Endpoints en ligne | **3** | 3 | illimités |
| Transfert | **1 Go/mois** | 5 Go | facturé ($0.10/Go) |
| Requêtes HTTP | **20 000/mois** | 100 000 | facturé |
| Page d'avertissement | **oui** (interstitiel) | non | non |
| Domaine | `*.ngrok-free.app` **persistant** (ne change plus au restart depuis 01/2026) | sous-domaines ngrok | domaine custom ($0.01/heure active) |
| Timeout de session | **aucun** (tourne indéfiniment) | aucun | aucun |
| TCP (ex. : SSH, Postgres) | ⚠ **vérification CB requise** | inclus | inclus |
| UDP | ❌ **non supporté** (aucun plan) | ❌ | ❌ |

**Traduction concrète :**
- Tester un webhook ou partager une démo ponctuelle : le Free suffit.
- La page d'avertissement fait fuir les non-techniques (« ça ressemble à
  du phishing ») : prévois le Hobbyist pour une vraie démo client.
- 1 Go / 20k requêtes : un dashboard qui poll toutes les 5 s les épuise
  en quelques jours. Surveille sur le dashboard ngrok.
- **Pas d'UDP** : inutile pour WireGuard, jeux, VoIP, CoAP.

## 39. Sécurité : ne pas exposer n'importe quoi

Checklist avant `ngrok http` :

- [ ] **Authentification** sur le service exposé (ne JAMAIS exposer un
      service sans mot de passe : pas de `psql`, pas de dashboard sans auth).
- [ ] **Principe du moindre temps** : tunnel levé pour la démo, coupé après
      (`Ctrl+C`). Pas de tunnel permanent « au cas où ».
- [ ] **Ne pas exposer** : bases de données, interfaces d'admin, Jupyter
      sans token, dossiers de fichiers personnels.
- [ ] **Logs** : le dashboard 127.0.0.1:4040 montre tout le trafic —
      vérifie qu'aucun secret ne transite en clair.
- [ ] **Token ngrok** : révocable depuis le dashboard si compromis.
- [ ] Alternative plus sûre pour un accès perso : **Tailscale** (réseau
      privé, pas d'URL publique) plutôt qu'un tunnel public.

## 40. Alternative n°1 : Cloudflare Tunnel (cloudflared) — gratuit

**Vérifié sept 2026 :** le meilleur rapport qualité/prix pour un usage
perso sérieux.

```bash
# Tunnel éphémère : ZÉRO compte, ZÉRO config — parfait pour un test
cloudflared tunnel --url http://localhost:8000
# → https://quelque-chose.trycloudflare.com  (aléatoire, change à chaque restart)
```

Limites du Quick Tunnel (vérifié) : usage test/dev uniquement, ~200 requêtes
simultanées max (429 au-delà), pas de Server-Sent Events, sous-domaine
aléatoire temporaire.

**Tunnel nommé** (stable, toujours gratuit, mais nécessite un domaine chez
Cloudflare) :

```bash
cloudflared tunnel login            # une fois (navigateur)
cloudflared tunnel create rag-api   # crée le tunnel + credentials
cloudflared tunnel route dns rag-api api-perso.ton-domaine.fr
cloudflared tunnel run rag-api
```

Fichier `~/.cloudflared/config.yml` :

```yaml
tunnel: rag-api
credentials-file: /home/<utilisateur>/.cloudflared/<uuid>.json
ingress:
  - hostname: api-perso.ton-domaine.fr
    service: http://localhost:8000
  - service: http_status:404   # règle de repli obligatoire
```

`cloudflared` installe son propre service systemd (`cloudflared service install`).

## 41. Alternative n°2 : Tailscale Funnel — gratuit, sans domaine

Si tu utilises déjà Tailscale (ou comptes l'utiliser) :

```bash
tailscale serve 8000    # HTTPS visible UNIQUEMENT de ton tailnet (tes appareils)
tailscale funnel 8000   # HTTPS visible de TOUT Internet
# → https://<machine>.<ton-tailnet>.ts.net  (stable, certificat auto)
```

**Vérifié sept 2026 :** inclus dans le plan **Personal gratuit** ; ports
publics limités à **443, 8443, 10000** ; pas de domaine custom ; bande
passante plafonnée (non publiée). Idéal : « exposer un service de ma
maison sans acheter de domaine ».

## 42. Autres alternatives (vérifié : existent en 2026)

| Outil | Modèle | Points forts | Limites (vérifiées) |
|---|---|---|---|
| **bore** (`rapiz1/bore`) | open source, self-hosted | simple, rapide ; serveur sur ton VPS | install via cargo — **à vérifier** sur le dépôt |
| **frp** (`fatedier/frp`) | open source, self-hosted | très complet (TCP/UDP/HTTP) | config des deux côtés (client + serveur VPS) |
| **zrok** | open source (Apache 2.0) + offre hébergée | moderne, zero-trust | jeune — **à vérifier** la maturité pour ton usage |
| **pinggy** | freemium (Pro ~$2.50/mois annuel) | TCP **et** UDP, URL persistante pas chère | limites du free — **à vérifier** |
| **localxpose** | freemium (Pro ~$8/mois) | UDP, nombreux tunnels | idem — **à vérifier** |

> **Ma recommandation pour toi :** Quick Tunnel `cloudflared` pour les tests
> jetables (gratuit, sans compte) ; Tailscale Funnel si tu veux du stable
> sans domaine ; ngrok Free si tu as besoin de l'inspection de trafic
> (dashboard 4040) ou de TCP.

## 43. Cas d'usage : exposer ton API RAG en démo

```bash
# Terminal 1 (session tmux "rag") : ton API
uvicorn api:app --host 127.0.0.1 --port 8000

# Terminal 2 : tunnel éphémère pour la démo
cloudflared tunnel --url http://localhost:8000
# → partage l'URL trycloudflare.com à ton collègue

# Si tu veux rejouer/inspecter les requêtes : plutôt ngrok
ngrok http 8000   # puis http://127.0.0.1:4040 pour l'inspection
```

## 44. Pense-bête tunnels

```bash
ngrok http 8000                              # tunnel ngrok basique
ngrok config add-authtoken "TOKEN"           # auth (une fois)
cloudflared tunnel --url http://localhost:8000  # tunnel jetable gratuit
tailscale funnel 8000                        # stable, sans domaine
```

- [ ] Token ngrok stocké hors git
- [ ] Service exposé = authentifié
- [ ] Tunnel coupé après usage (sauf besoin permanent documenté)
- [ ] Free ngrok : surveiller 1 Go / 20k req

---

## 45. Playwright : pourquoi c'est ton arme de scraping

