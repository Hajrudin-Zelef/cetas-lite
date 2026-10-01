---
id: collect-261001-general-networking/general-networking/tutoriel-mcp-server-13-etapes-typescript-et-python-2026-2
title: "Forcer une version et lancer l'inspector en mode HTTP"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Microsoft"]
dates: []
keywords: ["agent", "agents", "claude", "copilot", "dpo", "mcp"]
source: docs/RAG/collect-261001-general-networking/tutoriel-mcp-server-13-etapes-typescript-et-python-2026.md
source_anchor: ""
source_lines: [48, 209]
sha256: 0de0a80f140f98627a25d9026822fcd63e8ed76d66b85d5c7f5827fef296bf31
---

# Forcer une version et lancer l'inspector en mode HTTP

Créez un dossier `mcp-ts-server`, générez un `package.json` en ESM strict et installez le SDK officiel ainsi que Zod pour la validation des schémas. ESM est obligatoire : le SDK TypeScript n’exporte plus de CommonJS depuis la v1.10.

```
mkdir mcp-ts-server && cd mcp-ts-server
npm init -y
npm pkg set type="module"
npm install @modelcontextprotocol/[email protected] [email protected]
npm install -D [email protected] @types/node@22 [email protected]
```
Ajoutez un fichier `tsconfig.json` minimal. Le `moduleResolution` doit être `NodeNext` sinon les imports `.js` du SDK échoueront :

```
{
  "compilerOptions": {
    "target": "ES2022",
    "module": "NodeNext",
    "moduleResolution": "NodeNext",
    "outDir": "dist",
    "rootDir": "src",
    "strict": true,
    "esModuleInterop": true,
    "skipLibCheck": true
  },
  "include": ["src/**/*"]
}
```
Complétez les scripts npm pour le build, le run en développement (via `tsx`) et le lancement via MCP Inspector :

```
npm pkg set scripts.build="tsc"
npm pkg set scripts.dev="tsx watch src/index.ts"
npm pkg set scripts.start="node dist/index.js"
npm pkg set scripts.inspector="npx @modelcontextprotocol/inspector node dist/index.js"
```
Créez `src/index.ts` avec un squelette `McpServer`. Le constructeur prend un objet `{ name, version }` et déclare les capabilities. Sans `tools: {}`, la méthode `tools/list` renverra une erreur `-32601 (Method not found)`.

```
import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
const server = new McpServer(
  { name: "ti-demo-server", version: "1.0.0" },
  { capabilities: { tools: {}, resources: {}, prompts: {} } }
);
async function main() {
  const transport = new StdioServerTransport();
  await server.connect(transport);
  console.error("MCP server ready on stdio");
}
main().catch((err) => { console.error(err); process.exit(1); });
```
Compilez avec `npm run build` et lancez `npm run inspector`. Une interface web s’ouvre sur `http://localhost:5173` et montre le handshake réussi. Si rien ne s’affiche, vérifiez que vous écrivez bien vos logs sur `stderr` (`console.error`) et jamais sur `stdout` – ce dernier est réservé au protocole JSON-RPC.

## Étape 2 – Créer votre premier Tool avec Zod et schémas JSON

Un tool MCP est une fonction exposée au LLM. Elle est décrite par un *input schema* Zod, converti automatiquement en JSON Schema par le SDK, et un *handler* asynchrone qui reçoit les arguments typés et renvoie un tableau de `content`. Ajoutons un outil `calculer_tva` qui applique les taux français de TVA.

```
import { z } from "zod";
server.tool(
  "calculer_tva",
  "Calcule la TVA française pour un montant HT donné",
  {
    montant_ht: z.number().positive().describe("Montant hors taxes en euros"),
    taux: z.enum(["20", "10", "5.5", "2.1"]).describe("Taux de TVA applicable"),
  },
  async ({ montant_ht, taux }) => {
    const t = parseFloat(taux) / 100;
    const ttc = montant_ht * (1 + t);
    const tva = montant_ht * t;
    return {
      content: [
        {
          type: "text",
          text: `HT: ${montant_ht.toFixed(2)} € · TVA (${taux}%) : ${tva.toFixed(2)} € · TTC: ${ttc.toFixed(2)} €`,
        },
      ],
    };
  }
);
```
La signature `server.tool(name, description, inputShape, handler)` est la forme canonique du SDK 1.29. Les trois erreurs les plus courantes : (1) passer un schéma Zod au lieu du *shape* brut (le SDK wrap lui-même) ; (2) retourner un objet au lieu de `{ content: [...] }` ; (3) oublier de décrire les paramètres avec `.describe()`, ce qui réduit nettement la qualité d’appel du LLM.

Rebuildez, relancez l’Inspector, et dans l’onglet *Tools*, testez avec `{ "montant_ht": 1000, "taux": "20" }`. Vous obtiendrez `HT: 1000,00 € · TVA (20%) : 200,00 € · TTC: 1200,00 €`. Le LLM reçoit exactement ce texte comme contenu d’appel outil, formaté prêt à être cité.

## Étape 3 – Exposer Resources et Prompts pour l’agent

Les *resources* exposent des données en lecture, identifiées par une URI. Contrairement aux tools, elles ne sont pas « appelées » : le client les référence et le host décide quand les injecter dans le contexte du LLM. Deux formes existent : statique (une URI fixe) ou dynamique via `ResourceTemplate` avec paramètres.

```
import { ResourceTemplate } from "@modelcontextprotocol/sdk/server/mcp.js";
// Ressource statique
server.registerResource(
  "config-entreprise",
  "config://entreprise",
  { title: "Configuration entreprise", mimeType: "application/json" },
  async () => ({
    contents: [{ uri: "config://entreprise", text: JSON.stringify({ pays: "FR", devise: "EUR" }) }],
  })
);
// Ressource templatée
server.registerResource(
  "facture",
  new ResourceTemplate("factures://{id}", { list: undefined }),
  { title: "Facture par ID" },
  async (uri, { id }) => ({
    contents: [{ uri: uri.href, text: `Facture ${id} – statut : payée` }],
  })
);
```
Les *prompts* sont des gabarits réutilisables avec des arguments. Ils apparaissent dans le menu *slash commands* de Claude Desktop (`/`). Un prompt bien conçu évite à l’utilisateur de réécrire une amorce complexe et sert aussi aux agents automatisés.

```
server.prompt(
  "audit_rgpd",
  "Génère un audit RGPD pour un fichier de traitement",
  { fichier: z.string().describe("Chemin du fichier à auditer") },
  async ({ fichier }) => ({
    messages: [
      {
        role: "user",
        content: {
          type: "text",
          text: `Réalise un audit RGPD complet du fichier ${fichier}. Liste : finalité, base légale, DPO, durée de conservation, risques identifiés, actions correctives.`,
        },
      },
    ],
  })
);
```
Dans MCP Inspector, l’onglet *Resources* liste `config://entreprise` et permet de tester `factures://INV-2026-001`. L’onglet *Prompts* affiche `audit_rgpd` et accepte l’argument. Si une ressource ne s’affiche pas, vérifiez que `capabilities.resources = {}` est bien déclaré dans le constructeur du serveur.

## Étape 4 – Configurer le transport stdio pour Claude Desktop

Le transport `stdio` reste le plus simple pour un serveur MCP local. Claude Desktop le lance comme sous-processus, dialogue en JSON-RPC via stdin/stdout, et ferme le flux à la déconnexion. C’est aussi le transport par défaut pour Cursor, Windsurf et VS Code avec Copilot MCP.

Le fichier de configuration Claude Desktop dépend du système d’exploitation :

| Système | Chemin du fichier de configuration | 
|---|---|
| macOS | `~/Library/Application Support/Claude/claude_desktop_config.json` | 
| Windows | `%APPDATA%\Claude\claude_desktop_config.json` | 
| Linux | `~/.config/Claude/claude_desktop_config.json` | 

Ajoutez votre serveur dans la clé `mcpServers`. Le chemin du binaire `node` doit être absolu sous macOS (pas d’alias `nvm` car Claude Desktop ne charge pas votre shell) et les arguments doivent pointer sur le `dist/index.js` compilé :

```
{
  "mcpServers": {
    "ti-demo": {
      "command": "/usr/local/bin/node",
      "args": ["/Users/vous/mcp-lab/mcp-ts-server/dist/index.js"],
      "env": {
        "NODE_ENV": "production"
      }
    }
  }
}
```
Redémarrez Claude Desktop (`Cmd+Q` puis relance). L’icône 🔌 dans la barre d’entrée indique que le serveur est connecté. Cliquez dessus pour lister les tools exposés. En cas d’échec, consultez les logs dans `~/Library/Logs/Claude/mcp-server-ti-demo.log` sous macOS. Les erreurs les plus fréquentes : chemin Node inexistant, permissions d’exécution manquantes (`chmod +x dist/index.js`), ou écriture parasite sur stdout (`console.log` au lieu de `console.error`).

## Étape 5 – Passer au Streamable HTTP avec Express et SSE

