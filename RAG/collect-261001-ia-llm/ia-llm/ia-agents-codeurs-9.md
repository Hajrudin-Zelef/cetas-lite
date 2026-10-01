---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-9
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft", "OpenAI", "OpenRouter"]
dates: []
keywords: ["agent", "agents", "claude", "copilot", "mcp", "open source"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [1440, 1604]
sha256: 7265fbf0f3d0101d2cb589270292bbdac40a74efb2d2b32c2792cf45eb9dd2e0
---

# Les agents codeurs IA

|  | OpenCode | Claude Code | Codex CLI | Cursor | Kilo Code |
|---|---|---|---|---|---|
| Nature | CLI/TUI open source | CLI officiel Anthropic | CLI officiel OpenAI (OSS) | Éditeur IA | Agent multi-surfaces (OSS) |
| Modèles | Tous (BYOK) | Claude only | GPT only | Multi (dont maison) | Tous (BYOK) |
| Install | `curl .../install \| bash` | `npm i -g @anthropic-ai/claude-code` | `npm i -g @openai/codex` | cursor.com/download | marketplace / `npm i -g @kilocode/cli` |
| Mémoire projet | `AGENTS.md` | `CLAUDE.md` + `AGENTS.md` | `AGENTS.md` | `.cursor/rules/` | `AGENTS.md` |
| Plan mode | `Tab` (Plan/Build) | `Shift+Tab` | `/plan` | périmètre validé | Architect |
| Prix d'entrée | 0 € + ta clé | 20 $/mois ou API | 20 $/mois (Plus, inclus) ou API | ~20 $/mois (à vérifier) | 0 € + ta clé / Pass ~19 $ |
| Scriptable CI | `opencode run` | `claude "…"` | **`codex exec`** | Cloud Agents | CLI `kilo` |
| Verrouillage | Aucun | Anthropic | OpenAI | Anysphere | Aucun (mais jeune) |

## 51. Cline : c'est quoi, pour qui

**Cline** (anciennement « Claude Dev », extension VS Code `saoudrizwan.claude-dev`)
est **l'ancêtre open source (MIT)** de toute la famille : Cline → Roo Code →
Kilo Code. C'est un **agent autonome dans la barre latérale de VS Code** :
il lit/écrit des fichiers, exécute des commandes, utilise même un navigateur,
**toujours avec ta validation** par étape.

Particularités 2026 :

- **Boucle Plan/Act explicite** : rien ne se fait sans ton feu vert par
  défaut — l'agent le plus « transparent » sur ce qu'il fait.
- **BYOK + 200+ modèles via OpenRouter** : agnostique, comme OpenCode.
- **MCP** : oui, config manuelle.
- **CLI 2.0** : mode headless pour CI/CD (les extensions agentiques n'ont
  généralement pas ça).
- Statut août 2026 : **actif et maintenu**, la référence historique — même
  si l'innovation s'est déplacée vers Kilo Code.

Pour qui : tu veux un **agent open source DANS VS Code** (pas un nouvel
éditeur), avec un contrôle fin à chaque étape, et tu acceptes de configurer
toi-même (clés, MCP). Le choix « prudent et transparent ».

## 52. Cline : installation pas à pas

1. VS Code → Extensions (`Ctrl+Shift+X`) → cherche **« Cline »**
   (éditeur : `saoudrizwan.claude-dev`).
2. Installe → l'icône Cline apparaît dans la barre d'activité.
3. Ouvre la vue Cline → **choisis ton fournisseur** :
   - **Clé API directe** (Anthropic, OpenAI, Google…) : colle ta clé.
   - **OpenRouter** : une seule clé pour 200+ modèles (recommandé pour
     tester sans multiplier les comptes).
   - **Modèle local** (Ollama) : gratuit, privé, plus lent/faible.
4. Règle les **auto-approvals** (section 53) avant la première tâche.

> 🔒 Comme toujours : la clé API vit dans le gestionnaire de Cline
> (stockage sécurisé de VS Code), jamais dans un fichier du projet.

## 53. Cline : modes Plan/Act et approbations

Le cœur de Cline, c'est son **contrat de confiance explicite** :

```
┌─────────┐     ┌─────────┐     ┌──────────────┐
│  PLAN   │ ──▶ │ TON AVIS│ ──▶ │     ACT      │
│ lit,    │     │ tu valides│    │ édite, exécute│
│ propose │     │ ou corriges│   │ par étapes   │
└─────────┘     └─────────┘     └──────────────┘
       ▲                               │
       └─────── tu interromps/corriges ─┘
```

**Auto-approve** : par type d'action, tu règles ce qui passe sans demander :

- ✅ Lecture de fichiers : auto (sans risque).
- ⚠️ Écriture/édition : **demander** au début, auto quand tu fais confiance
  au périmètre.
- 🔴 Exécution de commandes : **toujours demander** au début (c'est là que
  vivent les `rm -rf` et les `pip install` sauvages).
- 🌐 Accès navigateur : demander.

Checklist de démarrage (à cocher) :

- [ ] Lecture : auto-approve ON
- [ ] Écriture : OFF (demander) les 2 premières semaines
- [ ] Commandes : OFF, avec limite « ne jamais auto-approuver `sudo`,
      `rm -rf`, `pip install`, `curl | bash` »
- [ ] Chaque session : relire le résumé d'actions avant de fermer

## 54. Cline : premier workflow commenté

Scénario : ajouter du **logging** propre à un script, en gardant la main.

```
1. Ouvre /tmp/lab-agent dans VS Code. Panneau Cline.
2. Prompt (mode Plan implicite) :
   « Analyse check_ups.py et propose un plan pour remplacer les print()
     par du logging structuré (niveau configurable via --log-level).
     Ne modifie rien pour l'instant. »
   → Cline répond avec un plan + les fichiers qu'il lirait.
3. Clique « Act » (ou valide) : Cline exécute étape par étape.
   À CHAQUE action sensible (édition, commande), une carte d'approbation
   apparaît : [Approve] [Reject] (+ « toujours approuver ce type »).
4. Quand il propose `pytest -q`, approuve : tu vois la sortie dans le panneau.
5. À la fin : « Review » du diff intégré, puis commit git manuel.
```

Le point pédagogique : **Cline te montre tout**. C'est plus lent que Cursor,
mais après 2 semaines tu as compris la boucle agentique dans les moindres
détails — une excellente école avant de passer à des agents plus autonomes.

**Mode headless (CLI 2.0)** pour CI :

```bash
# Exemple indicatif (syntaxe exacte : doc Cline en vigueur)
cline --headless "exécute pytest -q et résume les échecs" --approve read-only
```

## 55. Cline : MCP — configuration manuelle

Cline a son **propre fichier de settings MCP**, séparé de celui de VS Code :

| OS | Chemin |
|---|---|
| Linux | `~/.config/Code/User/globalStorage/saoudrizwan.claude-dev/settings/cline_mcp_settings.json` |
| macOS | `~/Library/Application Support/Code/User/globalStorage/saoudrizwan.claude-dev/settings/cline_mcp_settings.json` |
| Windows | `%APPDATA%\Code\User\globalStorage\saoudrizwan.claude-dev\settings\cline_mcp_settings.json` |

Le plus simple : panneau Cline → icône **MCP Servers** → **Configure MCP
Servers** (ouvre le fichier), puis :

```json
{
  "mcpServers": {
    "docs": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-filesystem", "/home/zelef/docs"],
      "disabled": false,
      "autoApprove": []
    }
  }
}
```

- `autoApprove: []` = l'agent **demandera** avant d'utiliser les outils de ce
  serveur. Mets des outils en auto-approve uniquement quand ils sont en
  lecture seule et de confiance.
- Redémarre la vue Cline après modification.

## 56. Cline : prix, forces, faiblesses, verdict

**Prix** : extension **gratuite et open source (MIT)**. Tu paies uniquement
l'usage modèle (ta clé API / OpenRouter / local). Le moins cher du marché à
fonctionnalité égale — et le plus transparent sur ce que tu consommes
(puisque tu vois la facture du fournisseur).

**Points forts** : transparence totale (chaque action est montrée et
validée) ; open source MIT ; BYOK multi-fournisseurs ; MCP ; CLI headless
pour CI ; historique et communauté énormes (tutos, modes personnalisés
partagés) ; excellent pour **apprendre** l'agenticité.

**Faiblesses honnêtes** : **pas d'autocomplétion inline** (ce n'est pas son
métier — couple-le avec Supermaven/Copilot si tu veux du Tab) ; config MCP
manuelle (fichier caché, pas de marketplace intégré comme Kilo) ; un seul
agent généraliste (pas de modes Architect/Code/Debug natifs — des modes
personnalisés JSON existent mais c'est du bricolage) ; rythme d'innovation
plus lent que Kilo Code depuis 2026 ; l'innovation s'est déplacée.

**Verdict Zelef** : **installe-le pour apprendre**, même si tu choisis un
autre agent au quotidien. Deux semaines avec Cline en mode « tout valider »
t'apprennent plus sur les agents que deux mois de Cursor en pilote
automatique. Et pour un usage sobre (BYOK + modèle pas cher), ça reste
l'option la plus économique.

