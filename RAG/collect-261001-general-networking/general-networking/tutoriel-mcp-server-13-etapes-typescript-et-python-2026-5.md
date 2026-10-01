---
id: collect-261001-general-networking/general-networking/tutoriel-mcp-server-13-etapes-typescript-et-python-2026-5
title: "Forcer une version et lancer l'inspector en mode HTTP"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: ["2025-06-18", "2025-11-25"]
keywords: ["benchmark", "claude", "mcp", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-mcp-server-13-etapes-typescript-et-python-2026.md
source_anchor: ""
source_lines: [452, 548]
sha256: eb78e0e61567fa00946eadd9e818aec9e716317714376f9c81806469301381b8
---

# Forcer une version et lancer l'inspector en mode HTTP

```
import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { z } from "zod";
import { readFile, readdir } from "node:fs/promises";
import path from "node:path";
const ALLOWED_ROOTS = (process.env.MCP_ALLOWED_ROOTS ?? "").split(",").filter(Boolean);
function checkPath(p: string) {
  const abs = path.resolve(p);
  if (!ALLOWED_ROOTS.some((r) => abs.startsWith(path.resolve(r)))) {
    throw new Error(`Chemin refusé : ${abs}`);
  }
  return abs;
}
const server = new McpServer(
  { name: "ti-filesystem-mcp", version: "1.0.0" },
  { capabilities: { tools: {}, resources: {}, prompts: {} } }
);
server.tool(
  "lire_fichier",
  "Lit un fichier texte dans un dossier autorisé",
  { chemin: z.string() },
  async ({ chemin }) => {
    const abs = checkPath(chemin);
    const contenu = await readFile(abs, "utf-8");
    return { content: [{ type: "text", text: contenu }] };
  }
);
server.tool(
  "lister_dossier",
  "Liste les entrées d'un dossier autorisé",
  { chemin: z.string() },
  async ({ chemin }) => {
    const abs = checkPath(chemin);
    const entries = await readdir(abs, { withFileTypes: true });
    const lines = entries.map((e) => `${e.isDirectory() ? "DIR " : "FILE"} ${e.name}`);
    return { content: [{ type: "text", text: lines.join("\n") }] };
  }
);
server.prompt(
  "resume_code",
  "Résume un fichier de code",
  { fichier: z.string() },
  async ({ fichier }) => ({
    messages: [{ role: "user", content: { type: "text", text: `Résume clairement le rôle, les dépendances et les risques du fichier ${fichier}.` } }],
  })
);
async function main() {
  const transport = new StdioServerTransport();
  await server.connect(transport);
  console.error("ti-filesystem-mcp ready");
}
main();
```
Lancez avec `MCP_ALLOWED_ROOTS="/Users/vous/projets" node dist/index.js`. Le serveur refuse toute requête hors de l’allowlist avec une erreur explicite. Testez-le dans Claude Desktop ou Cursor : demandez « Liste le contenu du dossier projets et résume le fichier `main.ts` », le modèle chaîne automatiquement `lister_dossier`, `lire_fichier` puis `resume_code`.

## 8 pièges courants et comment les éviter

La communauté MCP a identifié huit pièges récurrents qui cassent les intégrations débutantes. Les connaître évite des heures de débogage stérile.

| Piège | Symptôme | Solution | 
|---|---|---|
| Écriture sur stdout | Handshake échoue, Claude Desktop déconnecte | Logger sur stderr uniquement ( `console.error` ) | 
| Versions protocole désynchronisées | Erreur « unsupported version » | Mettre à jour SDK et spec à 2025-11-25 | 
| Session HTTP sans header | 400 Bad Request sur 2e appel | Renvoyer l’en-tête `Mcp-Session-Id` | 
| Schémas Zod nested non supportés | Appel outil ignoré | Aplatir les schémas, éviter les unions complexes | 
| Chemins relatifs sous Claude Desktop | ENOENT au démarrage | Utiliser des chemins absolus | 
| Token OAuth sans audience | Risque d’exfiltration silencieuse | Valider le `aud` claim à chaque requête | 
| Timeouts longs non gérés | Connexion SSE fermée au bout de 10 s | Heartbeats toutes les 25 s | 
| Permissions Linux 644 | `Permission denied` | `chmod +x dist/index.js` post-build | 

Deux pièges avancés méritent un mot supplémentaire. D’abord, la **fuite de sessions** côté serveur HTTP : si vous stockez les transports dans un objet global sans purge, la mémoire grimpe linéairement. Ajoutez un `transport.onclose = () => delete transports[sid]` et un TTL de 15 minutes par défaut. Ensuite, la **réentrance des tools** : un tool qui appelle un autre tool via le client interne peut boucler si le LLM enchaîne des requêtes. Protégez avec un sémaphore par session.

## Astuces avancées pour serveurs MCP en production

Au-delà de la base, quelques pratiques distinguent les serveurs MCP sérieux. Première astuce : utilisez **structured content** (spec 2025-06-18) plutôt que du texte brut. Un tool peut retourner `{ content: [{ type: "text", text }], structuredContent: { champs... } }`, qui est consommé directement par le LLM sans parsing. Deuxième astuce : exposez un endpoint `/.well-known/mcp-metadata` pour que les hosts découvrent automatiquement vos capabilities sans ouvrir une session complète – pratique pour les marketplaces de serveurs MCP.

Troisième astuce : implémentez le **sampling**. Un serveur peut, au lieu de répondre lui-même, demander au client d’appeler son propre LLM (c’est la primitive inverse). Utile pour des tools qui nécessitent une génération (traduction, résumé) : le client paie l’inférence, pas le serveur. Quatrièmement, exploitez l’**elicitation** : le serveur peut requérir un input utilisateur en cours d’exécution d’un tool (confirmation, valeur manquante), via une URL ou un prompt dédié – très puissant pour les workflows humains-dans-la-boucle.

Cinquième astuce : profilez avec `0x` (Node) ou `py-spy` (Python) avant toute optimisation. Un tool qui paraît lent est souvent bloqué sur un appel HTTP externe, pas sur son code – un simple cache LRU de 30 secondes sur les appels API gouv.fr ou GitHub réduit la latence médiane de 400 ms à moins de 10 ms.

Sixième astuce : exploitez les **notifications/cancelled** pour annuler proprement un tool long. Le client peut envoyer un message pour interrompre une requête ; côté serveur, écoutez l’événement via `signal.onabort` et propagez l’annulation à vos appels `fetch` ou requêtes SQL. Sans cela, un utilisateur qui ferme son IDE laisse tourner des handlers fantômes pendant plusieurs minutes. Septième astuce : paginez systématiquement les ressources volumineuses. Le SDK prend en charge `nextCursor` dans les réponses `tools/list` et `resources/list` ; implémentez-le dès que votre serveur dépasse 50 entrées pour éviter les payloads de plusieurs mégaoctets qui saturent le contexte LLM.

Huitième astuce : adoptez une convention de nommage claire pour vos tools (`verbe_objet` en snake_case, descriptions commençant par un verbe d’action, e.g. « Calcule », « Liste », « Crée »). Les LLM modernes sélectionnent mieux un outil dont le nom et la description reflètent exactement son intention métier. Testez empiriquement la reconnaissance avec le framework `mcp-eval` (en open source depuis fin 2025), qui mesure le taux d’appel correct d’un tool sur un benchmark de 500 requêtes utilisateur réelles.

## FAQ sur les serveurs MCP en 2026

### MCP remplace-t-il OpenAI Function Calling ?

Non, ils sont complémentaires. Function Calling est un format d’API propriétaire OpenAI (et Anthropic), destiné à invoquer une fonction dans un prompt unique. MCP est un protocole de session persistante entre hôte et serveur, supportant tools, resources, prompts, sampling et elicitation. En pratique, un serveur MCP traduit ses tools en Function Calls vers n’importe quel modèle.

### Faut-il apprendre TypeScript ou Python pour MCP ?

Les deux SDK sont au même niveau de maturité en avril 2026. TypeScript bénéficie du plus large écosystème de serveurs publics (plus de 2 000 paquets npm dans la communauté), Python convient mieux aux serveurs qui s’appuient sur scikit-learn, PyTorch ou pandas. Pour Java, C# et Kotlin, des SDK officiels sont disponibles depuis 2025.

### Quels sont les serveurs MCP publics les plus populaires ?

