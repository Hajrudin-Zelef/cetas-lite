---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-18
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI"]
dates: ["2026-09-16", "2026-09-17", "2026-09-18", "2026-09-23", "2026-09-27"]
keywords: ["agent", "agents", "apache", "claude", "fine-tuning", "luna", "mcp", "open source", "opus 5", "pricing", "sol", "training"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [2884, 3000]
sha256: a103ed902533ef1aaa22ed615231b9faa3d310ced5f9b13f7d93de1301eda51a
---

# Les agents codeurs IA

- Les **tarifs exacts au centime** : ils changent tous les trimestres —
  les pages pricing officielles font foi (liens section 108).
- Les **versions précises** : le secteur publie chaque semaine ; un numéro
  de version imprimé serait faux dans un mois.
- Le **fine-tuning** de modèles pour le code : autre sujet, autre guide.
- Les aspects **juridiques** fins (licence du code généré, responsabilité) :
  à voir avec un juriste pour un usage professionnel encadré — ce guide
  donne les bons réflexes, pas un avis juridique.
- Les outils **« NON TROUVÉ »** : aucun dans cette édition — tous les noms
  de la mission ont été identifiés (OpenClaw pour « open clow »).

## 110. Quiz : 10 questions (réponses en fin de section)

**Q1.** Quelle commande installe OpenCode sur Linux/macOS (méthode officielle) ?
**Q2.** Dans Claude Code, que fait `Shift+Tab` et pourquoi est-ce le réflexe n°1 ?
**Q3.** Cite les trois modes de permission de Codex CLI et celui à utiliser en découverte.
**Q4.** Quelle est la différence de licence entre le CLI de Claude Code et celui de Codex ?
**Q5.** Dans Cursor, à quoi servent respectivement Tab, Composer et l'Agent ?
**Q6.** Pourquoi Cline est-il recommandé comme « école » des agents ?
**Q7.** « Open clow » n'existe pas : quel est le vrai nom de l'outil, et quelle attaque de 2026 impose de le durcir ?
**Q8.** Cite trois choses à ne JAMAIS confier à un agent (section 92).
**Q9.** Un abonnement à 20 $/mois ou du BYOK au token : dans quel cas le BYOK gagne-t-il ?
**Q10.** Qu'est-ce que le « test qui ment » (piège n°3) et comment le détecter ?

---

**Réponses :**

**R1.** `curl -fsSL https://opencode.ai/install | bash` (ou `npm install -g opencode-ai`).
**R2.** `Shift+Tab` bascule en **Plan mode** : l'agent explore et propose sans écrire. C'est le réflexe n°1 car valider un plan coûte des centimes, alors que corriger une mauvaise implémentation coûte des dollars et du temps.
**R3.** **Suggest** (propose, tu valides tout) / **Auto-edit** (édite seul, demande pour exécuter) / **Full-auto** (tout seul). En découverte : **Suggest**.
**R4.** Le CLI de Claude Code est **fermé** (propriétaire) ; Codex CLI est **open source Apache 2.0**.
**R5.** **Tab** = complétion inline pendant la frappe ; **Composer** = applique une description sur plusieurs fichiers (avec diffs à valider) ; **Agent** = exécute une tâche multi-étapes de façon autonome (avec background/cloud).
**R6.** Parce que sa boucle **Plan/Act** montre et fait valider **chaque action** par défaut : en deux semaines de « tout valider », on comprend la boucle agentique mieux qu'en deux mois de pilotage automatique.
**R7.** Le vrai nom est **OpenClaw** (ex-Clawdbot/Moltbot). L'attaque **ClawHavoc** (début 2026 : skills malveillantes, instances exposées, RCE) impose bind local, utilisateur sans sudo, skills auditées, crons lecture seule d'abord.
**R8.** (au choix parmi) : les **secrets** ; la **production sans garde-fou** ; l'**exécution aveugle** (skip-permissions/Full-auto) ; le réseau sans contrôle ; les suppressions auto ; l'installation de dépendances non validées ; les décisions irréversibles.
**R9.** Le BYOK gagne quand : usage **faible ou en rafales** (on paie ce qu'on consomme), ou quand on veut **une seule clé pour plusieurs outils**. L'abonnement gagne pour un usage quotidien régulier et prévisible.
**R10.** L'agent **affaiblit le test** au lieu de corriger le code (seuil baissé, assertion supprimée, cas retiré). Détection : passe 2 de la revue — lire le diff **des tests** aussi, traquer `assert True`, tests vides, mocks suspects.

**Score** : 8/10 → tu peux passer à la pratique (section 107). Moins de 8 → relis les sections indiquées entre parenthèses.

## 111. Historique des vérifications (transparence)

| Date | Vérification |
|---|---|
| 27 sept. 2026 | Existence, installation, prix d'OpenCode, Kilo Code, Claude Code, Codex CLI, Cursor, Cline, OpenClaw, Continue, Supermaven — par recherche web |
| 27 sept. 2026 | Points d'incertitude marqués « à vérifier » : grille Cursor (restructuration 2026), offre OpenCode Zen/Go, tarifs API exacts, syntaxes de config mineures |
| — | Prochaine re-vérification conseillée : **janvier 2027** (ou à chaque changement d'abonnement) |

## 112. Le mot de la fin

Un agent codeur ne remplace ni ton expérience de sysadmin, ni ta
responsabilité de chef de service. Il **multiplie** ce que tu sais déjà
faire — et il multiplie aussi tes erreurs si tu le laisses sans cadre.

Le cadre tient en trois artefacts que ce guide t'a fait construire :

1. **`AGENTS.md`** — ton contrat avec la machine, versionné, relu.
2. **La revue en 4 passes** — ton contrôle qualité, systématique.
3. **Le pense-bête** — tes garde-fous, imprimés au-dessus du poste.

Avec ça, tu n'es plus « un mec qui regarde l'IA coder ». Tu es un
**pilote d'agents** : tu fixes le cap (le plan), tu tiens le manche (les
permissions), et tu restes responsable de l'atterrissage (la revue).

Bon code — et surveille tes tokens. 💰

---

*Guide rédigé le 27 septembre 2026. 110+ sections. Vérifications web : voir section 111.*

## 113. À venir — annonces vérifiées au 27/09/2026

> Méthode : recherche web effectuée le 27/09/2026. Seules les
> annonces officielles ou vérifiables figurent ici ; les rumeurs sont
> étiquetées **RUMEUR non confirmée**. Si rien d'annoncé pour un outil,
> c'est écrit noir sur blanc — rien n'est inventé.

### 113.1 Claude Code — Projects repensés en multi-agents (bêta)

Annonce d'Anthropic le **17/09/2026** : les Projects de Claude Code
sont repensés comme un **coordinateur multi-agents**. Tu décris un
objectif, Claude le découpe en **threads parallèles** (chacun sa session
cloud, sa branche, sa copie du dépôt), suit l'avancement depuis la
conversation principale et partage **mémoire + bibliothèque de fichiers**
entre threads. Les threads peuvent ouvrir des PR, lancer des tests et
se subdiviser en subagents/workflows ; les conflits sont gérés comme
des merge conflicts Git classiques.

- **Statut au 27/09/2026 :** bêta, sélection d'abonnés **Pro et Max** en
  sessions cloud uniquement.
- **Annoncé comme à venir :** exécution sur outils/locaux « très
  bientôt », élargissement à Team et Enterprise ensuite (sans date).
- Conséquence pour toi : l'équivalent « multi-agents coordonnés » que ce
  guide fait faire à la main (sessions parallèles, `tmux`, revue en 4
  passes) devient un produit natif — mais **cloud-only au départ**,
  donc hors air-gap. La revue en 4 passes reste indispensable : le
  coordinateur multiplie la production, pas la relecture.
- Sources : *The Decoder*, 17/09/2026 (« Anthropic keeps pushing Claude
  Code toward autonomous coding with new parallel agent workflows ») ;
  *iTechPost*, 18/09/2026 ; *Complete AI Training*, 18/09/2026.

Anthropic a par ailleurs annoncé le **16/09/2026** (Reuters) la fusion
des interfaces chat + **Cowork**, l'intégration de **Claude Design** et
le lancement en bêta de **Claude Docs** et **Claude Slides** (export
Google Docs / Word / PowerPoint / PDF) — d'abord Pro/Max, puis Team et
Free. Impact indirect : les artefacts de projet (docs, slides) rejoignent
la même interface que l'agent de code.

Sortie du **23/09/2026** : Claude Code **v2.1.281** (correctif,
correctifs de stabilité post-Opus 5.5) — correctifs de réparation de
transcripts (« unexpected tool_use_id »), de timeouts MCP en `http`
après ~5 minutes, et de Write refusés sur des paramètres légèrement
erronés. Le **mode autopilot** est désormais le défaut dans Claude Code.

### 113.2 Codex (OpenAI) — CLI 0.156, Sol/Luna, contexte persistant

