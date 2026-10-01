---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-16
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "incident", "mcp", "qwen", "valuation"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [2618, 2761]
sha256: f20ea969464153157bf5ebee1f870032fae08655c1c3265633f06adf2129febd
---

# Les agents codeurs IA

**17. Le skill/plugin vérolé** — skill ClawHub ou plugin VS Code malveillant
(prompt injection, exfiltration).
→ *Antidote* : sources connues uniquement, lecture du code/description,
post-ClawHavoc : paranoïa saine (sections 66, 72).

**18. La confiance aveugle dans le vert** — « les tests passent » alors que
l'agent a mocké la fonction testée.
→ *Antidote* : lis les tests générés comme le code (section 90, passe 2).

**19. Le vendor lock-in silencieux** — 6 mois de règles `.cursor/`, de
skills `.claude/` et d'habitudes : changer d'outil coûte une blinde.
→ *Antidote* : **AGENTS.md comme source de vérité** (standard multi-outils),
skills au format ouvert quand possible.

**20. L'atrophie** — à force de déléguer, tu ne sais plus écrire le code que
tu supervises. Pour un chef de service responsable d'une prod, c'est le
piège le plus grave.
→ *Antidote* : impose-toi des plages « sans agent » ; fais expliquer par
l'agent (mode Ask) au lieu de faire faire ; relis vraiment les diffs.

## 100. Cas pratique n°1 : script de supervision d'onduleur (commenté)

Le scénario sysadmin par excellence : un check SNMP d'onduleur, format
Nagios, **avec** les garde-fous.

```bash
mkdir -p /tmp/lab-agent/ups && cd /tmp/lab-agent/ups && git init -q
cat > AGENTS.md <<'EOF'
# AGENTS.md — check_ups
- Python 3.12, zéro dépendance externe (stdlib uniquement).
- SNMP : paramètres via variables d'environnement (UPS_HOST, UPS_COMMUNITY),
  valeurs fictives dans les exemples et les tests.
- Sortie : code retour Nagios (0/1/2/3) + 1 ligne de texte.
- Tests : pytest, 100 % des branches critiques.
- INTERDIT : écrire un vrai hostname ou une vraie communauté, même en exemple.
EOF
```

Prompt (Claude Code, plan mode) :

```
[Plan mode] Besoin : check_ups.py — interroge un onduleur en SNMP
(OID fictifs documentés dans le README), sort un code Nagios selon
l'état batterie (0 OK / 1 WARNING <50 % / 2 CRITICAL <25 % / 3 UNKNOWN
si SNMP injoignable). Contraintes : stdlib uniquement, timeout SNMP
de 5 s, aucun secret en dur, tests pytest avec un faux serveur SNMP.
Propose le plan (fichiers, fonctions, cas de test). N'écris rien.
```

Ce que tu valides dans le plan : le découpage (collecte / évaluation /
sortie), les cas limites (timeout, OID absent), les tests (faux SNMP, pas
de vrai équipement). Puis exécution, puis revue en 4 passes (section 90).
**Le point pédagogique** : sur un sujet que tu maîtrises (les onduleurs),
tu détectes immédiatement les approximations de l'agent — c'est comme ça
qu'on calibre sa confiance.

## 101. Cas pratique n°2 : revue d'un playbook Ansible par l'agent

```bash
cd /tmp/lab-agent && mkdir -p ansible && cd ansible && git init -q
# Place un playbook existant (expurgé de tout secret !) dans playbook.yml
```

Prompt (Codex CLI, non-interactif — la voie CI) :

```bash
codex exec --full-auto \
  "Relis playbook.yml avec la skill implicite suivante : 1) idempotence
   de chaque tâche, 2) aucun secret en clair (signale toute variable
   suspecte), 3) tout 'shell:' remplaçable par un module natif,
   4) handlers correctement notifiés, 5) 'become' minimal.
   Rends un rapport markdown : BLOQUANT / REMARQUES / OK par tâche.
   Ne modifie AUCUN fichier."
```

Ensuite, en CI (GitHub Actions) : ce `codex exec` tourne à chaque PR qui
touche `ansible/` — ton **relecteur automatique** (voir l'exemple YAML
section 34). Le verdict reste humain, mais les oublis bêtes sont attrapés
par la machine.

## 102. Cas pratique n°3 : ingestion RAG assistée (ton chantier)

Ton pipeline RAG (collecte → nettoyage → indexation) est un terrain de jeu
idéal pour un agent **en lecture/écriture cadrée** :

```
Projet : ~/rag-pipeline (fictif, à adapter)
Agent : Claude Code (serveur MCP 'docs' branché sur ton corpus)

Session type :
> [Plan mode] Dans corpus/brut/, identifie les 10 fichiers avec le plus
> de résidus (headers/footers, sommaires, HTML) en échantillonnant.
> Propose une stratégie de nettoyage par type de résidu. N'écris rien.

> Implémente la stratégie validée sous forme de scripts dans tools/clean/,
> un script par type de résidu, avec tests. Ne touche JAMAIS à corpus/brut/
> (lecture seule) : les sorties vont dans corpus/propre/.
```

Garde-fous spécifiques RAG :

- **Source brute intouchable** (lecture seule, idéalement en snapshot).
- **Idempotence** : relancer le nettoyage donne le même résultat.
- **Échantillonnage de contrôle** : l'agent montre 5 exemples avant/après
  avant de traiter 10 000 fichiers.
- **Traçabilité** : chaque fichier propre garde la référence de sa source.

## 103. Cas pratique n°4 : débogage assisté (méthode)

Quand quelque chose casse et que tu veux l'agent en **détective**, pas en
chirurgien :

```
> [Mode Ask / Plan] Le service X échoue au démarrage depuis ce matin.
> Indices : journalctl -u X (je te le colle ci-dessous), config dans
> /etc/X/ (lisible). NE MODIFIE RIEN. Pose-moi des questions si un
> indice manque, puis propose 3 hypothèses classées par probabilité
> avec, pour chacune, LA commande qui la confirmerait.
[coller les logs]
```

Pourquoi ça marche : tu forces l'agent à **raisonner avant d'agir**, tu
gardes la main sur les commandes (c'est toi qui les lances), et tu obtiens
un raisonnement écrit — utile pour ton rapport d'incident.

## 104. Cas pratique n°5 : documentation d'un parc (batch nocturne)

```bash
# Cron hebdo : régénère la doc d'inventaire à partir des données
0 2 * * 0  cd /home/zelef/inventaire && \
  opencode run -m openrouter/qwen/qwen3-coder-flash \
  "à partir de data/inventaire.csv, régénère docs/parc.md (tableaux par
   site, alertes de fin de garantie). Ne modifie rien d'autre." \
  && git diff --exit-code -- docs/ || \
  git commit -am "docs: maj auto parc $(date +%F)"
```

Points de vigilance : modèle **pas cher** (tâche mécanique), `git diff`
pour ne commiter que si ça a changé, **relecture humaine le lundi matin**
(le cron propose, tu disposes). Jamais de push auto vers `main` sans
relecture — le commit local suffit, tu pousses après contrôle.

## 105. Glossaire (les 30 termes à connaître)

