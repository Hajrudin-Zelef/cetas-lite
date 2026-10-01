---
id: collect-261001-general-networking/general-networking/tutoriel-mcp-server-13-etapes-typescript-et-python-2026-3
title: "Forcer une version et lancer l'inspector en mode HTTP"
domain: general-networking
role: reference
task: reference
actors: ["IREN"]
dates: ["2025-06-18", "2025-11-25"]
keywords: ["attention", "mcp"]
source: docs/RAG/collect-261001-general-networking/tutoriel-mcp-server-13-etapes-typescript-et-python-2026.md
source_anchor: ""
source_lines: [210, 335]
sha256: 0e3654711b5b94ed6a55b1c3d172ac5ec96fbd8dfa17ff9937245a86efb90c56
---

# Forcer une version et lancer l'inspector en mode HTTP

Pour un serveur distant consommé par plusieurs hosts, stdio ne suffit pas. Le transport **Streamable HTTP**, introduit dans la spec 2025-06-18 et stabilisé en 2025-11-25, remplace l’ancien SSE. Il utilise une seule URL `/mcp` gérant POST (requêtes) et GET (flux SSE pour notifications serveur).

```
import express from "express";
import { randomUUID } from "node:crypto";
import { StreamableHTTPServerTransport } from "@modelcontextprotocol/sdk/server/streamableHttp.js";
const app = express();
app.use(express.json());
const transports: Record<string, StreamableHTTPServerTransport> = {};
app.post("/mcp", async (req, res) => {
  const sessionId = req.headers["mcp-session-id"] as string | undefined;
  let transport = sessionId ? transports[sessionId] : undefined;
  if (!transport) {
    transport = new StreamableHTTPServerTransport({
      sessionIdGenerator: () => randomUUID(),
    });
    await server.connect(transport);
    transports[transport.sessionId!] = transport;
  }
  await transport.handleRequest(req, res, req.body);
});
app.get("/mcp", async (req, res) => {
  const sessionId = req.headers["mcp-session-id"] as string;
  const transport = transports[sessionId];
  if (!transport) return res.status(400).send("Session inconnue");
  await transport.handleRequest(req, res);
});
app.listen(3000, () => console.error("MCP HTTP sur http://localhost:3000/mcp"));
```
Trois points d’attention : (1) l’en-tête `Mcp-Session-Id` est obligatoire après le premier `initialize` ; (2) activez CORS explicitement pour les clients navigateur (`cors({ exposedHeaders: ["Mcp-Session-Id"] })`) ; (3) prévoyez un timeout de connexion SSE à 30 secondes minimum, sinon les notifications longues sont coupées. Testez depuis Inspector en mode « HTTP » avec l’URL `http://localhost:3000/mcp`.

## Étape 6 – Construire un serveur MCP en Python avec FastMCP

Le SDK Python officiel `mcp` expose la classe `FastMCP`, inspirée de FastAPI, qui génère automatiquement les schémas JSON à partir des annotations de type et des docstrings. Créez un nouveau projet avec `uv` (gestionnaire recommandé) :

```
mkdir mcp-py-server && cd mcp-py-server
uv init --python 3.12
uv add "mcp[cli]>=1.15"
uv add httpx
```
Créez `server.py` avec un serveur qui interroge l’API gouvernementale française `entreprise.data.gouv.fr` pour récupérer un SIREN. La fonction est typée et documentée en français : FastMCP extrait la description et les types pour le LLM.

```
import httpx
from mcp.server.fastmcp import FastMCP
mcp = FastMCP("ti-siren-server")
@mcp.tool()
async def rechercher_siren(siren: str) -> dict:
    """Recherche les informations publiques d'une entreprise française par SIREN.
    Args:
        siren: Le numéro SIREN à 9 chiffres.
    """
    url = f"https://recherche-entreprises.api.gouv.fr/search?q={siren}"
    async with httpx.AsyncClient(timeout=10) as client:
        r = await client.get(url)
        r.raise_for_status()
        data = r.json()
        if not data.get("results"):
            return {"erreur": "SIREN inconnu"}
        e = data["results"][0]
        return {
            "nom": e.get("nom_complet"),
            "siren": e.get("siren"),
            "naf": e.get("activite_principale"),
            "effectif": e.get("tranche_effectif_salarie"),
            "siege": e.get("siege", {}).get("adresse"),
        }
@mcp.resource("sirets://{siren}")
def profil_entreprise(siren: str) -> str:
    return f"Profil synthétique pour SIREN {siren}."
if __name__ == "__main__":
    mcp.run(transport="stdio")
```
Testez via `uv run mcp dev server.py`, qui lance automatiquement MCP Inspector sur le serveur Python. Pour un transport HTTP, appelez `mcp.run(transport="streamable-http", host="0.0.0.0", port=8000)`. Les décorateurs `@mcp.tool()`, `@mcp.resource()` et `@mcp.prompt()` suivent la même ergonomie : typage, docstring, retour JSON-sérialisable.

## Étape 7 – Sécuriser avec OAuth 2.1 et PKCE (spec 2025-11-25)

La spec 2025-11-25 mandate **OAuth 2.1 avec PKCE** pour tout serveur MCP distant. Le serveur est classifié comme OAuth Resource Server (RFC 9728), et expose un `/.well-known/oauth-protected-resource` qui pointe vers l’Authorization Server. Les clients utilisent la Dynamic Client Registration (RFC 7591) ou les nouveaux OAuth Client ID metadata documents pour s’enregistrer sans interaction humaine.

```
import { requireBearerAuth } from "@modelcontextprotocol/sdk/server/auth/middleware/bearerAuth.js";
const authMiddleware = requireBearerAuth({
  verifier: {
    verifyAccessToken: async (token) => {
      const response = await fetch("https://auth.exemple.fr/introspect", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: new URLSearchParams({ token }),
      });
      const data = await response.json();
      if (!data.active) throw new Error("Token invalide");
      return {
        token,
        clientId: data.client_id,
        scopes: data.scope?.split(" ") ?? [],
      };
    },
  },
  requiredScopes: ["mcp:read", "mcp:call"],
});
app.post("/mcp", authMiddleware, async (req, res) => { /* ... */ });
```
Côté client, le flow est : (1) le client découvre `/.well-known/oauth-protected-resource` ; (2) il récupère le `authorization_server` ; (3) il lance un Authorization Code flow avec PKCE et `resource` indicator (RFC 8707) pointant sur le serveur MCP ; (4) l’access token obtenu est présenté en `Authorization: Bearer` à chaque requête. Le consentement incrémental de scopes, nouveauté de la spec 2025-11-25, évite de tout demander d’un coup : le serveur peut renvoyer un `403` avec un `insufficient_scope`, déclenchant une nouvelle phase d’autorisation granulaire.

Trois impératifs de sécurité : utilisez TLS 1.3 en production, stockez les tokens chiffrés (AES-GCM, jamais en clair dans une base), et n’acceptez jamais `audience` absent ou vide – c’est le vecteur classique d’exfiltration de jetons entre serveurs MCP.

## Étape 8 – Déboguer avec MCP Inspector

MCP Inspector est l’outil officiel de débogage. Il se lance via `npx @modelcontextprotocol/inspector <commande>`, par exemple `npx @modelcontextprotocol/inspector node dist/index.js` ou `npx @modelcontextprotocol/inspector uv run server.py`. L’interface web expose cinq onglets : *Connection*, *Tools*, *Resources*, *Prompts*, *Sampling*.

Dans *Connection*, vérifiez la version de protocole négociée et les capabilities. Un écart de version entre client et serveur (par exemple 2025-06-18 vs 2025-11-25) provoque un fallback vers la plus ancienne supportée ; si aucune ne matche, la connexion échoue. Pour forcer une version, passez `--protocol-version 2025-11-25` en ligne de commande.

Dans *Tools*, chaque outil est listé avec son input schema en JSON. Un bouton *Call* permet d’exécuter l’outil avec des valeurs de test et d’afficher la réponse brute ainsi que les erreurs JSON-RPC. Le panneau *Request Log* à droite montre toute la trace JSON-RPC en temps réel, incluant les `notifications/message` pour le logging structuré.

```
# Forcer une version et lancer l'inspector en mode HTTP
npx @modelcontextprotocol/inspector \
  --protocol-version 2025-11-25 \
  --transport streamable-http \
  --server-url http://localhost:3000/mcp
```
Astuce : activez le logging côté serveur avec `server.server.sendLoggingMessage({ level: "info", data: "..." })`. Chaque message apparaît dans l’Inspector sous *Server Logs*. Si rien ne s’affiche, vérifiez que la capability `logging: {}` est déclarée et que le client a appelé `logging/setLevel`.

## Étape 9 – Connecter à Cursor, Windsurf et VS Code

