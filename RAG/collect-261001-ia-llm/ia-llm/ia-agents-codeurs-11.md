---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-11
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agents", "copilot", "mcp", "open source"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [1766, 1956]
sha256: e6d2ec7447563de7adf2114d9854d55ac7adaa40014652394e42cad8f3b5af1e
---

# Les agents codeurs IA

⚠️ **Le revers de la viralité** : l'attaque **« ClawHavoc »** (documentée
début 2026) — 800+ skills malveillantes uploadées sur ClawHub (prompt
injection, payloads malveillants), des dizaines de milliers d'instances
exposées (défaut d'écoute sur `0.0.0.0:18789`), plusieurs CVE. Moralité :
**OpenClaw est puissant ET c'est l'outil de ce guide qui exige le plus de
rigueur sécurité** (section 72).

## 67. OpenClaw : installation pas à pas

Prérequis : une machine **qui reste allumée** (serveur Debian, Mini-PC,
VPS), Node.js, et **une clé API** (Anthropic ou OpenAI — c'est « le cerveau »).

```bash
# Méthode officielle — vérifiée sept. 2026
curl -fsSL https://openclaw.ai/install.sh | bash

# Alternative npm
npm install -g openclaw@latest

# Assistant d'onboarding (wizard) + installation en daemon (service)
openclaw onboard --install-daemon
```

Le wizard te demande :

1. **Le fournisseur de modèle** + ta clé API (Anthropic ou OpenAI ; modèles
   locaux via Ollama possibles).
2. **Le canal de messagerie** : Telegram (le plus simple pour débuter —
   crée un bot via @BotFather, colle le token), WhatsApp, Discord, Slack…
3. **Le workspace** : ton répertoire home devient son espace de travail
   (identité, mémoire, config y vivent).

```bash
# Vérifier que le daemon tourne
openclaw status

# Voir les logs
openclaw logs --tail 50
```

> 🔒 **Dès l'install** : ne laisse PAS l'écoute réseau sur `0.0.0.0`
> (c'était le défaut impliqué dans ClawHavoc). Lie l'interface à
> `127.0.0.1` ou protège par firewall + authentification. Voir section 72.

Alternative Docker (si tu préfères) :

```bash
git clone https://github.com/openclaw/openclaw.git
cd openclaw
docker compose up -d
```

## 68. OpenClaw : premier workflow commenté

Scénario : **briefing matinal automatique** + une tâche à la demande.

**Étape 1 — la tâche à la demande (depuis Telegram) :**

```
Toi (Telegram) : « Liste les fichiers modifiés cette semaine dans
                 ~/projets/inventaire et résume les changements. »

OpenClaw : [exécute git log, lit les diffs] « 4 fichiers modifiés :
           - app/collecte.py : ajout du retry SNMP (+2 fonctions)
           - tests/test_collecte.py : 3 nouveaux tests
           - README.md : section installation mise à jour
           - requirements.txt : pysnmp ajouté
           Veux-tu que je génère le changelog ? »
```

**Étape 2 — le cron (le vrai pouvoir d'OpenClaw) :**

```
Toi (Telegram) : « Tous les jours à 7h00, envoie-moi :
                 1) la météo du jour,
                 2) mes 3 rendez-vous du calendrier,
                 3) l'état des sauvegardes (vérifie /var/log/backup.log). »

OpenClaw : « C'est noté. Cron 'briefing-matin' créé : 0 7 * * *.
           Je te l'enverrai ici. »
```

**Étape 3 — la mémoire persistante :**

```
Toi : « Retiens que mes onduleurs sont des Easy UPS 3S et que je veux
       toujours un résumé en français, concis. »

OpenClaw : [écrit dans sa mémoire] « Noté. »
# → Trois jours plus tard, il s'en souvient sans qu'on lui redise.
```

C'est ça, la différence : **continuité** (mémoire), **initiative** (crons),
**profondeur** (il connaît ton contexte) — pas juste un chatbot.

## 69. OpenClaw : skills, ClawHub, MCP

- **Skills** : le format d'extension d'OpenClaw. Une skill = un dossier avec
  une description + les outils qu'elle apporte (naviguer sur le web, envoyer
  un mail, piloter ta domotique…).
- **ClawHub** (clawhub.ai) : le registre communautaire (des milliers de
  skills). Installation d'une skill :

```
Toi (Telegram) : « Installe la skill 'github-releases' depuis ClawHub. »
```

⚠️ **Règle d'or post-ClawHavoc** : n'installe une skill que si (1) elle est
populaire et auditée, (2) tu as lu ce qu'elle fait, (3) l'analyse VirusTotal
— désormais intégrée — est propre. **Jamais de skill obscure en prod.**

- **MCP** : OpenClaw supporte MCP — tu peux brancher les mêmes serveurs que
  tes autres agents (ton serveur « docs » du RAG, par exemple).
- **Modèles locaux** : Ollama/LM Studio supportés — un OpenClaw 100 % local
  (cerveau local + machine locale) est possible pour les usages sensibles,
  au prix d'un cerveau moins brillant.

## 70. OpenClaw : ce qu'il ne faut JAMAIS lui confier (garde-fous)

OpenClaw est l'agent le plus **autonome** de ce guide — donc celui où les
garde-fous comptent le plus :

1. 🔒 **Jamais d'accès direct à la prod** sans validation humaine : pas de
   redémarrage de service, pas de `apt upgrade`, pas de modification réseau
   en autonome.
2. 🔒 **Jamais de secrets en mémoire en clair** au-delà du nécessaire : la
   clé API du « cerveau » et les tokens de messagerie sont le minimum vital ;
   tout le reste (mots de passe, clés SSH) passe par des accès limités.
3. 🔒 **Jamais d'instance exposée sur Internet sans auth** : le défaut
   `0.0.0.0:18789` a causé des prises de contrôle (RCE). Bind local +
   VPN/Tailscale ou reverse proxy avec auth forte.
4. 🔒 **Crons en lecture seule d'abord** : un cron qui *lit* et *résume*,
   oui ; un cron qui *modifie* ou *envoie*, seulement après des semaines de
   comportement sain — et avec notification avant action.
5. 🔒 **Principe du moindre privilège** : l'utilisateur système qui fait
   tourner le daemon n'a ni sudo sans mot de passe, ni accès aux clés de
   prod, ni écriture hors de son workspace.

Checklist de durcissement (à cocher après install) :

- [ ] Écoute réseau liée à 127.0.0.1 ou derrière auth (pas 0.0.0.0 nu)
- [ ] Utilisateur dédié sans sudo
- [ ] Skills installées : liste relue, sources connues
- [ ] Crons : tous en lecture/notification pour commencer
- [ ] Sauvegarde du workspace (mémoire + config) chiffrée

## 71. OpenClaw : prix

| Poste | Coût |
|---|---|
| Le logiciel | **0 €** — MIT, gratuit |
| Le « cerveau » | Ta clé API (Anthropic/OpenAI) au token, ou modèle local gratuit |
| La machine | Ton serveur existant (coût marginal ~nul) ou VPS (~5 €/mois) |
| La messagerie | 0 € (bots Telegram gratuits) |

C'est l'agent **le moins cher** de ce guide en coût logiciel — le coût est
dans les tokens du cerveau et dans **ton temps de durcissement**. Ne
sous-estime pas ce dernier.

## 72. OpenClaw : verdict Zelef

**Le plus excitant ET le plus dangereux de ce guide.** Pour un chef de
service systèmes : un OpenClaw bien durci qui t'envoie le briefing infra du
matin sur Telegram, qui surveille tes logs et te prévient — c'est un vrai
multiplicateur. Mais c'est aussi une **surface d'attaque** ( ClawHavoc l'a
prouvé) et un agent qui agit **sans toi**. Ma recommandation : installe-le
sur une machine dédiée, en **lecture seule + notifications** pendant un
mois, durcis selon la checklist (section 70), et ne lui donne des droits
d'écriture que progressivement. Pour coder, garde les outils des parties A
et B — OpenClaw n'est pas un agent codeur, c'est un **agent de vie**.

---

# PARTIE D — PLUGINS & COMPLÉTION : LE PANORAMA

---

## 73. Les trois métiers des plugins (ne pas confondre)

| Métier | Ce que ça fait | Exemples (sept. 2026) |
|---|---|---|
| **Complétion** | Suggère la suite pendant que tu tapes | Supermaven, Copilot, Tabnine, Codeium, Tabby |
| **Chat in-IDE** | Répond/explique dans l'éditeur | Continue (chat), Copilot Chat, JetBrains AI |
| **Agent en extension** | Agit de façon autonome | Cline, Kilo Code, Continue (mode Agent) |

Un plugin peut cumuler (Continue fait les trois). La règle d'achat : **paie
pour UN métier à la fois** — payer deux complétions concurrentes, c'est
payer deux fois pour des suggestions qui se marchent dessus.

## 74. Continue : l'extension open source pragmatique

