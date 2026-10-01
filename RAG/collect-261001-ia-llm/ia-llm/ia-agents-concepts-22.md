---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-22
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Meta", "Microsoft", "OpenAI", "Stripe"]
dates: ["2025-12-18", "2026-03-21", "2026-12-11"]
keywords: ["agent", "agentic", "agents", "chatgpt", "claude", "copilot", "gemini", "license", "mcp", "model context protocol", "muse"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [3073, 3229]
sha256: fbe9ffa0faa16c044c93302fa8b27b76e5846b8f6c17f9ea774a400eed7afa12
---

# Concepts : agents IA, agentic, autonomie

- **Préparer une réunion** : « résume mes e-mails non lus sur le projet X
  et propose un ordre du jour » (connecteur mail).
- **Logistique déplacements/astreintes** : réservation hôtel + vol, créneaux
  d'intervention synchronisés au calendrier.
- **Veille** : « surveille les CVE critiques sur les produits Y cette semaine
  et résume » — puis tu valides dans ton RAG.
- **Rédaction** : comptes rendus d'intervention, notes de calcul mises en
  forme, réponses aux tickets récurrents (brouillons à relire, jamais d'envoi
  auto).
- **Achats** : comparer devis onduleurs/pièces en ligne, remplir les
  formulaires fournisseurs — avec validation avant tout paiement.

Règle d'or : Muse = **exosquelette administratif**, ton infra reste sur tes
outils (scripts, RAG, supervision). Ne mélange pas les périmètres.

## 135. Skill (les compétences packagées des agents)

### 135.1. Le concept : une skill, c'est quoi ?

Une **skill** = un paquet de savoir-faire réutilisable pour un agent :
**instructions + outils + ressources**, rangés dans un simple dossier versionné
en git. Pas de serveur, pas de build, pas de SDK obligatoire : le filesystem
fait office d'API. Analogie d'Anthropic (conférence AI Engineer, avril 2026) :
*les modèles sont les processeurs, les runtimes d'agents sont les OS, les
skills sont les applications* — et n'importe qui peut en construire une en
mettant des fichiers dans un dossier.

**Format** (standard ouvert Agent Skills, Anthropic, 18/12/2025) :

```
ma-skill/
├── SKILL.md            # OBLIGATOIRE : frontmatter YAML + corps markdown
├── references/         # docs chargées à la demande (optionnel)
├── scripts/           # scripts exécutables par l'agent (optionnel)
└── assets/            # données, templates, images (optionnel)
```

`SKILL.md` minimal :

```markdown
---
name: diagnostic-onduleur
description: Diagnostiquer une alarme d'onduleur/UPS à partir des symptômes
  et des relevés (tensions, autonomie, logs). À utiliser quand un onduleur
  signale un défaut ou quand on prépare une intervention de maintenance.
---

# Diagnostic onduleur

Procédure à suivre quand cette skill est active...
```

Le **frontmatter** (`name` + `description`, obligatoires) est la clé du
système. Le **corps** markdown = la procédure que le modèle lit quand la
skill s'active.

### 135.2. La divulgation progressive (l'idée géniale)

Charger toutes les skills dans le contexte à chaque tour brûlerait des
milliers de tokens pour rien. Le standard résout ça en **3 niveaux** :

1. **Métadonnées** (~100 tokens par skill : `name` + `description`) compilées
   dans le prompt système au démarrage — assez pour savoir *quand* une skill
   est pertinente, rien de plus.
2. **Corps** : quand le modèle juge la skill pertinente, il charge lui-même
   le `SKILL.md` complet (via un outil de lecture de fichier).
3. **Fichiers groupés** : `references/` et `scripts/` lus/exécutés à la
   demande, uniquement si le corps les cite.

Décision de chargement = **le modèle**, via un appel d'outil — jamais un
simple matching de mots-clés dans le harness. Conséquence pratique : tu peux
avoir **des centaines de skills** sans exploser ta fenêtre de contexte.
L'anti-pattern connu : l'injection eager du corps complet à chaque tour.

Deux drapeaux utiles dans le frontmatter (convention Claude Code) :
`disable-model-invocation` (skill à usage humain uniquement — procédures à
effets de bord) et `user-invocable` (savoir de fond, jamais auto-chargé
comme procédure). Une skill peut aussi servir de **commande slash**.

### 135.3. État du standard en sept. 2026 (vérifié)

- **Annonce** : 18 décembre 2025 par Anthropic ; **standard ouvert**,
  gouverné par l'**Agentic AI Foundation** (Linux Foundation).
- **Adoption** : 30+ outils — Codex CLI, Gemini CLI, VS Code/Copilot,
  Cursor, Goose, Kiro, et d'autres (sources : notes de design clio-agent,
  oct. 2026 ; article timetobuildbob, 21/03/2026).
- **Partenaires de lancement** : Canva, Notion, Figma, Atlassian,
  Cloudflare, Stripe, Zapier (source : SKILL.md du projet octomind-tap,
  vérifié 2026).
- **Portable** : une même skill voyage entre Claude.ai, Claude Code,
  l'Agent SDK et la plateforme développeur Anthropic — et, le standard
  étant ouvert, vers les outils tiers compatibles.
- **Doc officielle** : `docs.claude.com` (Claude Code → skills ;
  agents-and-tools → agent-skills/best-practices) — à vérifier au jour
  de ta visite, les chemins bougent.

### 135.4. Les équivalents chez les autres

- **OpenAI — Plugins + Skills** (annoncé sept. 2026, cf. section 136.1) :
  les Custom GPTs sont retirés le **11/12/2026** ; leurs instructions
  deviennent une **Skill**, leurs fichiers de connaissances des **fichiers
  de référence**, leurs applis connectées un **Plugin**. ChatGPT sélectionne
  automatiquement une Skill quand la demande correspond à sa description —
  même logique de routage par description que le standard Anthropic.
  OpenAI avait aussi annoncé un « Skills Editor » pour exporter vers le
  format standard (source : recherche octomind, 2026 — à vérifier).
- **AGENTS.md** (convention `agents.md/`) : le « README pour agents » —
  contexte projet toujours chargé (Codex, Factory AI, Builder.io...).
  Différence : AGENTS.md = contexte global permanent ; **skill = procédure
  chargée à la demande**. Les deux se combinent.
- **MCP** (Model Context Protocol) : complément, pas concurrent.
  **Skill = le savoir-faire** (comment faire) ; **MCP = l'accès aux outils**
  (avec quoi agir). Une skill peut documenter l'usage d'un serveur MCP.
- **gptme / leçons** : précurseurs (fichiers auto-inclus, puis chargement
  par mots-clés mi-2025) — la skill standardise ce que ces outils faisaient
  déjà de façon ad hoc.

### 135.5. Écrire sa propre skill : exemple complet et fonctionnel

Prenons le cas de Zelef : une skill **`diagnostic-onduleur`** qui guide un
agent (ou un technicien assisté par IA) face à une alarme UPS.

**Arborescence :**

```
skills/diagnostic-onduleur/
├── SKILL.md
├── references/
│   ├── alarmes.md          # dictionnaire alarmes → causes probables
│   └── seuils.md           # seuils de décision (tensions, impédances)
├── scripts/
│   └── check_batteries.sh  # calcul d'autonomie à partir des relevés
└── assets/
    └── gabarit_cr.md       # template de compte rendu d'intervention
```

**`SKILL.md` :**

```markdown
---
name: diagnostic-onduleur
description: Diagnostiquer une alarme d'onduleur/UPS (VFI, VI ou VFD)
  à partir des symptômes, des relevés électriques et des logs.
  À utiliser quand un onduleur signale un défaut, une alarme batterie,
  un transfert intempestif sur bypass, ou avant une intervention de
  maintenance. Ne pas utiliser pour le dimensionnement d'une
  installation neuve (voir la skill 'dimensionnement-ups').
license: MIT
---

# Diagnostic onduleur

Tu es un assistant de diagnostic pour onduleurs. Tu ne répares rien :
tu analyses, tu proposes un plan d'action, un humain valide.

## Procédure

