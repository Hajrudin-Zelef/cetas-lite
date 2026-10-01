---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-3
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: ["2024-11-05", "2025-03-26", "2026-07-28"]
keywords: ["mcp"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [319, 474]
sha256: cdf993e78556ac197e0ac5cf827f6354793986584324a8bcb369bdf1fb21b213
---

# MCP, LSP, code-server, websearch — Guide pratique

Règles absolues :
1. **Un message = une ligne JSON complète**, sans `\n` interne.
2. **stdout est sacré** : uniquement du JSON-RPC. Aucun `print()`,
   aucun log, aucune bannière de démarrage. Tout le reste → `stderr`.
3. À la fermeture de stdin, le serveur **doit se terminer** (pas de zombie).
4. L'annulation d'un appel en cours passe par `notifications/cancelled`
   avec l'`id` de la requête.
5. Le binaire doit être trouvable et exécutable (`command` + `args` dans
   la config client) ; les variables d'environnement passent par `env`.

💡 Test rapide : `echo '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}' | uv run server.py`
doit répondre une ligne JSON. Si tu vois du texte parasite avant, ton
stdout est pollué.

## 15. Transport Streamable HTTP : le standard distant

Un seul endpoint (par convention `/mcp`). Chaque message JSON-RPC = un
**POST** HTTP :

- **Notification** (pas de réponse attendue) → `202 Accepted`, corps vide.
- **Requête simple** → `200` avec `Content-Type: application/json`.
- **Requête longue** (progress) → `200` avec `text/event-stream` : le serveur
  envoie des notifications de progression puis la réponse finale dans le
  même stream. **Fermer le stream = annuler** la requête.

Propriétés qui changent tout par rapport à l'ancien HTTP+SSE :
- **Sans état** : pas de session obligatoire, pas de sticky sessions —
  chaque requête peut atterrir sur une instance différente derrière un
  load balancer.
- Les serveurs qui ont besoin d'état inter-appels émettent des **handles
  explicites** (tokens passés en arguments d'outils ordinaires), pas de
  cookie magique.
- Anti-timeout : le serveur doit envoyer des lignes de keep-alive SSE
  (commentaires `:`) et, derrière nginx, `X-Accel-Buffering: no` pour
  empêcher le buffering du proxy.

🔒 Côté sécurité : HTTPS obligatoire en production, authentification via
OAuth 2.1 (voir §18) ou clé API / headers personnalisés.

## 16. HTTP+SSE déprécié : pourquoi il faut migrer

L'ancien transport distant (spec 2024-11-05) utilisait **deux endpoints**
(POST `/message` + GET `/sse`) plus un identifiant de session : une
conversation = un état réparti sur deux connexions. Conséquences :
sessions « collantes » obligatoires, reconnexions fragiles, infra complexe.

La spec 2025-03-26 l'a déprécié, la 2026-07-28 l'enterre. Les SDK actuels
marquent `SSEServerTransport` comme legacy. **Tout nouveau serveur distant
doit viser Streamable HTTP.** Si tu hérites d'un vieux serveur en HTTP+SSE,
planifie la migration : les clients récents négocient Streamable HTTP
en priorité.

## 17. Comparatif transports : que choisir

| Critère | stdio | Streamable HTTP |
|---|---|---|
| Cas d'usage | Serveur local, poste de l'utilisateur | Serveur distant, équipe, multi-tenant |
| Lancement | Le client spawn le processus | Processus indépendant (systemd, conteneur) |
| Sécurité | Permissions OS, env du poste | HTTPS + OAuth/clé API |
| Scalabilité | 1 client par serveur | N clients, load balancer OK |
| Complexité | Minimale (un script suffit) | Serveur HTTP, TLS, auth à gérer |
| Debug | Facile (stderr, lancement manuel) | Logs HTTP, curl |
| Cas Zeef | Serveur doc perso sur ton poste | Serveur doc partagé à ton équipe |

💡 Règle simple : **commence en stdio**, passe en Streamable HTTP quand
quelqu'un d'autre que toi doit l'utiliser.

## 18. Authentification : OAuth 2.1, clés API, headers

MCP ne réinvente pas l'authentification :
- **OAuth 2.1** : le mécanisme standard pour les serveurs distants
  multi-utilisateurs. Le serveur expose ses métadonnées « protected resource » ;
  le client découvre l'authorization server et fait le flow classique
  (code + PKCE). Les serveurs hébergés (GitHub, Atlassian, Linear...) utilisent
  ça : une fenêtre navigateur s'ouvre à la première utilisation.
- **Clé API / Bearer** : header `Authorization: Bearer <clé>` — simple,
  suffisant pour un usage équipe restreinte. La clé transite en variable
  d'environnement côté client, jamais en dur dans un fichier commité (📌
  voir §35, expansion `${VAR}`).
- **Headers personnalisés** : `X-Api-Key`, etc. Supportés par les configs
  clients (`headers` dans `.mcp.json`).

🔒 Principe : le serveur **ne fait jamais confiance** au client sur parole.
Chaque `tools/call` sensible re-valide le token et les scopes. Un tool
`query` sur ta BDD prod sans auth = une porte ouverte.

## 19. DNS rebinding : la protection localhost

Un serveur MCP en Streamable HTTP qui écoute sur `127.0.0.1` reste
atteignable... par ton navigateur. Attaque classique : une page web
malveillante fait résoudre un domaine attaquant vers 127.0.0.1 puis appelle
ton serveur local (rebinding DNS). Les SDK récents intègrent une protection :
validation du header `Host`/`Origin` (`enableDnsRebindingProtection`,
`allowedHosts`, `allowedOrigins`).

🔒 Règles :
- En local, bind sur `127.0.0.1`, **jamais** `0.0.0.0` sauf besoin réel.
- Active la protection anti-rebinding si ton SDK la propose.
- Si ton serveur local n'a pas besoin d'être appelé par un navigateur,
  exige un token même en localhost.

## 20. MRTR : quand le serveur a besoin d'une réponse humaine

Avant : le serveur pouvait initier `sampling/createMessage` (redemander au
LLM) ou `elicitation/create` (ouvrir un dialogue). En 2026-07-28, c'est
remplacé par le pattern **MRTR (Multi Round-Trip Requests)** :

1. Le client appelle `tools/call`.
2. Le serveur répond `resultType: "input_required"` + `inputRequests`
   (ex : « il me faut le mot de passe du switch »).
3. Le client (via l'humain) **réémet la même requête** avec `inputResponses`.

Avantage : plus de requêtes serveur→client à travers le transport —
tout reste du request/response classique, compatible avec le mode sans état.
En pratique, peu de serveurs l'implémentent encore fin 2026 ; la plupart
demandent les infos manquantes **en paramètres d'outils** dès le départ.
💡 Design : préfère des tools dont les paramètres sont complets plutôt qu'un
dialogue en plusieurs tours.

## 21. Sampling déprécié : que faire à la place

Le sampling permettait à un serveur de dire « hé, modèle, complète ce texte
pour moi ». Cas d'usage réel : résumer le résultat d'un tool avant de le
renvoyer. Remplacement officiel : le serveur appelle **directement l'API du
fournisseur LLM** (clé API côté serveur). C'est plus explicite (coût visible,
modèle choisi par l'opérateur du serveur) et ça ne dépend plus du client.

⚠️ Si tu vois encore `sampling/createMessage` dans un vieux serveur, c'est
un signal « code non maintenu depuis 2025 ».

## 22. Roots déprécié : que faire à la place

Les « roots » étaient des dossiers que le client déclarait au serveur à
l'initialisation (« voici les racines auxquelles tu as accès »). Remplacement :
- passe les chemins **en paramètres** (`{"path": "/srv/docs"}`),
- ou expose-les en **resources** (`file:///srv/docs/...`).

C'est plus fin : l'accès est décidé **par appel**, pas une fois pour toutes
à la connexion. 🔒 Moindre privilège : un tool `read_file` qui accepte un
chemin arbitraire doit valider qu'il reste dans la racine autorisée
(anti path-traversal, voir §66).

## 23. Logging déprécié → stderr et OpenTelemetry

Fini `logging/setLevel` et les notifications `notifications/message`.
La doctrine 2026 :
- **stdio** : logs sur `stderr`, format libre (texte ou JSON).
- **Streamable HTTP / prod** : **OpenTelemetry** (traces + métriques +
  logs structurés) vers ton collecteur (Tempo, Jaeger, Loki...).

💡 Pour Zeef : un serveur MCP qui loggue chaque `tools/call` (outil, durée,
succès/échec) en JSON sur stderr = directement ingérable par ton Loki
existant (tu as le guide Loki, 3341 lignes — réutilise-le).

## 24. Utilitaires : pagination, complétion, notifications

