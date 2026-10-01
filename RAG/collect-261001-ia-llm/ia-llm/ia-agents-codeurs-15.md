---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-15
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "arr", "claude", "cost", "mcp"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [2435, 2617]
sha256: 7394390e2169025c57edb1c837f7b12f195a6d8f8616ea0b4fe4700c6efe2388
---

# Les agents codeurs IA

1. **Modèle adapté à la tâche** : petit modèle pour explorer/relire, gros
   pour coder. Le levier n°1 (OpenCode/Kilo le font nativement).
2. **Plan mode** : 1 session de plan à 0,50 $ évite 3 sessions de
   réécriture à 5 $.
3. **Prompts avec critères d'arrêt** : « arrête-toi quand les tests passent ».
4. **Donne les chemins exacts** : `@app/main.py:120` plutôt que « trouve où… »
   (chaque recherche exploratoire = des dizaines de k tokens).
5. **Sessions courtes** : `/compact` ou nouvelle session — un contexte
   obèse coûte cher ET dégrade la qualité.
6. **Limite le périmètre** : « ne touche qu'à `app/` » bat « améliore le projet ».
7. **Désactive les MCP inutiles** : chaque serveur = des outils décrits à
   chaque tour = des tokens.
8. **Sois concis dans les instructions système** (« réponds brièvement » dans
   AGENTS.md fait baisser l'output).
9. **Surveille** : `/cost` (Claude Code), `/status` (Codex), dashboard
   fournisseur, **alertes de facturation à 50/80/100 %**.
10. **Abonnement vs BYOK** : recalcule tous les 3 mois avec TES chiffres
    (section 88) — le bon choix change avec ton usage.

**Budget mensuel type** (usage régulier, 1 personne) :

| Profil | Coût mensuel indicatif |
|---|---|
| Léger (questions, petits scripts) | **0–10 $** (BYOK ou Hobby) |
| Régulier (features hebdo) | **20 $** (1 abonnement) |
| Intensif (agent quotidien, CI) | **40–100 $** (abo + API) |
| Équipe de 5 | **150–400 $** (sièges + API) — à piloter (section 89) |

## 97. MCP en pratique : ton premier serveur utile

Le serveur MCP le plus rentable pour toi : **filesystem sur ton corpus de
docs** (tes guides RAG). Tous tes agents pourront alors citer ta doc.

```bash
# Serveur MCP officiel "filesystem" (Node.js)
npx -y @modelcontextprotocol/server-filesystem /home/zelef/docs
```

Branchement par outil :

```bash
# Claude Code
claude mcp add docs -- npx -y @modelcontextprotocol/server-filesystem /home/zelef/docs

# OpenCode — dans opencode.json : clé "mcp" (voir section 13)

# Cline — cline_mcp_settings.json (voir section 55)

# Cursor — .cursor/mcp.json : { "mcpServers": { "docs": { "command": "npx", "args": [...] } } }

# Codex — ~/.codex/config.toml (voir section 35)
```

Idées de serveurs MCP pour un sysadmin (tous existent dans l'écosystème) :

- **filesystem** : tes docs, tes playbooks (lecture).
- **git** : historique, blame, diff sans quitter l'agent.
- **postgres/sqlite** : interroger l'inventaire (« combien d'onduleurs en
  alarme ? » → l'agent fait le SELECT lui-même).
- **snmp/nut** : via un petit serveur maison — l'agent interroge tes
  onduleurs (lecture seule !).
- **prometheus/grafana** : « pourquoi ce pic ? » → l'agent lit les métriques.

> 🔒 MCP = l'agent obtient des **pouvoirs**. Chaque serveur en écriture
> (base de données, API d'admin) est une surface d'attaque : commence par
> des serveurs **lecture seule**, et audite ce que chaque outil fait
> réellement avant de l'auto-approuver.

## 98. Skills : capitaliser ton savoir-faire

Une skill, c'est un **prompt packagé et versionné**. Quand tu te surprends à
réécrire le même long prompt pour la 3ᵉ fois, c'en est une.

```bash
# Claude Code
.claude/skills/revue-ansible/SKILL.md

# Kilo Code (standard ouvert agentskills.io)
.agents/skills/revue-ansible/SKILL.md

# OpenClaw
# via ClawHub ou dossier de skills local
```

Exemple — ta skill « revue Ansible » :

```markdown
# SKILL.md — revue-ansible
---
name: revue-ansible
description: Relit un playbook Ansible selon nos standards avant merge
---

Quand on te demande de relire un playbook :
1. Vérifie : idempotence (chaque tâche rejouable sans effet de bord),
   pas de mot de passe en clair (vault obligatoire), handlers notifiés,
   `become` minimal.
2. Signale tout `shell:` qui pourrait être un module natif.
3. Rends un verdict : BLOQUANT / REMARQUES / OK, avec les lignes concernées.
4. Ne modifie jamais le playbook : tu proposes, l'humain applique.
```

> 💡 Pour ton RAG : tes skills (« comment ingérer un doc », « comment
> nettoyer un corpus ») sont des artefacts versionnables au même titre que
> tes guides. Un jour, ton RAG répondra ET tes skills agiront — même socle.

## 99. Les 20 pièges classiques (et comment les éviter)

**1. La boucle infinie** — l'agent relit, réécrit, relance les tests en
boucle sans progresser.
→ *Antidote* : critères d'arrêt dans le prompt (« max 5 itérations, puis
résume et rends la main ») ; `Esc` dès que tu vois le 3ᵉ tour identique.

**2. L'hallucination de dépendance** — l'agent importe une librairie qui
n'existe pas (ou un paquet typosquatté).
→ *Antidote* : interdiction d'ajouter des dépendances sans validation
(AGENTS.md) ; `pip install` toujours manuel.

**3. Le test qui ment** — l'agent affaiblit le test pour qu'il passe
(seuil baissé, assertion supprimée).
→ *Antidote* : passe 3 de la revue (section 90) ; `git diff` sur les tests
aussi.

**4. Le refactor surprise** — « pendant que j'y étais, j'ai aussi
réorganisé… » 800 lignes hors périmètre.
→ *Antidote* : périmètre explicite dans le prompt ; `git diff --stat`
systématique.

**5. Le secret qui fuit** — clé API collée dans le prompt, `.env` lu et
recopié dans un exemple.
→ *Antidote* : `.env.example` + valeurs fictives ; jamais de secret sous
les yeux de l'agent (section 92).

**6. Le `rm -rf` confiant** — script de nettoyage généré avec un chemin
mal quoté ou une variable vide.
→ *Antidote* : jamais d'auto-approve sur les suppressions ; `set -u` dans
les scripts shell générés.

**7. La prod touchée « pour tester »** — l'agent confond l'inventaire de
test et la base de prod (même nom de variable !).
→ *Antidote* : environnements nommés différemment, connexions prod hors de
portée de l'agent.

**8. Le contexte obèse** — session de 2 h, l'agent oublie le début et
contredit ses propres décisions.
→ *Antidote* : `/compact`, sessions courtes, résumé écrit des décisions.

**9. Le « oui » systématique** — l'agent valide tes mauvaises idées au lieu
de les contester.
→ *Antidote* : demande explicitement la contradiction (« challenge mon
approche avant d'implémenter ») ; Claude Code le fait mieux que les autres.

**10. La doc qui pourrit** — l'agent met à jour le code mais pas le README,
ou écrit une doc qui décrit un comportement imaginaire.
→ *Antidote* : « mets à jour la doc **à partir du code final**, pas de
mémoire » + relecture.

**11. Le copier-coller multiplié** — plutôt que factoriser, l'agent duplique
un bloc 5 fois avec des variantes.
→ *Antidote* : « DRY : factorise avant de dupliquer » dans AGENTS.md ;
passe revue n°3.

**12. L'anglais/français mélangé** — commentaires franglais, messages de
commit incohérents.
→ *Antidote* : règle de langue explicite dans AGENTS.md (section 91).

**13. Le modèle sous-dimensionné** — tâche complexe confiée au modèle
« flash » pour économiser 2 $, résultat à jeter (coût réel : 1 h perdue).
→ *Antidote* : petit modèle = exploration, gros modèle = implémentation.

**14. Le MCP trop puissant** — serveur MCP en écriture auto-approuvé, l'agent
écrit en base « pour tester ».
→ *Antidote* : MCP lecture seule par défaut (section 97).

**15. Le multi-agent qui se marche dessus** — deux agents éditent le même
fichier, le dernier écrase le premier.
→ *Antidote* : un seul écrivain par dossier (section 61) ; worktrees pour
le parallélisme (Kilo).

**16. La facture surprise** — 3 sessions « juste pour voir » sur le gros
modèle en API = 60 $.
→ *Antidote* : alertes de facturation + `/cost` + modèle pas cher par défaut.

