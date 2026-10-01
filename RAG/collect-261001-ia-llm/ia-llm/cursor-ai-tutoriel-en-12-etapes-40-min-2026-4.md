---
id: collect-261001-ia-llm/ia-llm/cursor-ai-tutoriel-en-12-etapes-40-min-2026-4
title: "Windows (winget)"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft", "xAI"]
dates: []
keywords: ["agent", "agents", "attention", "claude", "copilot", "gemini", "grok", "mcp"]
source: docs/RAG/collect-261001-ia-llm/cursor-ai-tutoriel-en-12-etapes-40-min-2026.md
source_anchor: ""
source_lines: [263, 339]
sha256: 70bd55362ea3ee22fa3d047fae34366947d4efef8f48afa93180fb2a8f586ff7
---

# Windows (winget)

La configuration se fait dans un fichier `mcp.json`, au niveau du projet (`.cursor/mcp.json`) ou global. Voici un exemple minimal branchant un serveur MCP de système de fichiers et un serveur PostgreSQL :

```
// .cursor/mcp.json
{
  "mcpServers": {
    "fichiers": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-filesystem", "./"]
    },
    "postgres": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-postgres",
               "postgresql://user:pass@localhost:5432/ma_base"]
    }
  }
}
```
Une fois le serveur déclaré, ses outils apparaissent dans les paramètres de Cursor et l’agent peut les utiliser dans ses raisonnements. C’est ce qui permet, par exemple, de demander « interroge la base et génère un endpoint qui renvoie les tâches en retard » sans quitter l’éditeur. **Piège n°4 :** un serveur MCP exécute du code sur votre machine et peut accéder à des données sensibles – n’installez que des serveurs de sources fiables.

## Étape 12 – Automatiser avec Cursor CLI

Dernière brique : **Cursor CLI**, l’interface en ligne de commande qui apporte l’agent hors de l’éditeur graphique – dans un terminal, un script ou un pipeline d’intégration continue (CI). Elle ouvre la voie à des automatisations : correction de bugs déclenchée par un workflow, génération de documentation, revue de code dans la CI.

```
# Installation de l'agent CLI
curl https://cursor.com/install -fsS | bash
# Lancer une tâche agentique en une commande
cursor-agent "Ajoute une route GET /taches/stats qui renvoie
le nombre de tâches terminées et en cours"
# Mode non interactif, utile en CI (sortie exploitable)
cursor-agent -p "Corrige les avertissements de typage" --output-format text
```
Dans un pipeline, vous pouvez ainsi faire réviser automatiquement chaque *pull request* par un agent, ou générer un changelog. Votre projet de gestion de tâches est désormais complet : structuré, testé, documenté et automatisable. Versionnez-le avec `git init && git add . && git commit -m "API de tâches générée avec Cursor"` et vous disposez d’un dépôt propre.

## Sécurité et confidentialité : maîtriser ses données avec Cursor

Puisque les modèles de **Cursor** s’exécutent dans le cloud, une question revient sans cesse en entreprise : mon code part-il chez un tiers ? La réponse tient dans le **mode Privacy** (confidentialité), que nous vous recommandons d’activer dès le départ, surtout sur du code propriétaire. Lorsqu’il est activé, votre code n’est pas conservé sur les serveurs après traitement et n’est pas utilisé pour entraîner les modèles. C’est un critère décisif pour les équipes soumises à des obligations de conformité, notamment en Europe où le RGPD impose une vigilance particulière sur les transferts de données.

Concrètement, vérifiez trois points avant tout déploiement en équipe. D’abord, activez le mode Privacy dans les paramètres et confirmez qu’il s’applique à tous les modèles utilisés – y compris les modèles d’API tierces. Ensuite, exploitez le fichier `.cursorignore` pour exclure de l’indexation les fichiers sensibles (clés, secrets, données personnelles) : ce qui n’est pas indexé n’est pas envoyé. Enfin, pour les organisations, l’offre Enterprise ajoute l’authentification unique (SSO), des contrôles d’administration et des garanties contractuelles renforcées.

Gardez aussi à l’esprit le rythme d’évolution de l’outil : Anysphere publie des mises à jour très fréquentes, documentées dans son journal des versions. Les fonctions de sécurité, les modèles disponibles et les modalités de facturation peuvent donc changer d’un mois à l’autre. Pour situer l’entreprise dans le paysage plus large de l’IA – son histoire, ses levées de fonds et sa place face à Microsoft – la fiche de synthèse de Cursor (entreprise) sur Wikipédia offre un panorama utile. Dans tous les cas, considérez la sécurité comme une brique à configurer *avant* d’ouvrir votre premier dépôt professionnel, et non après.

## Tarifs de Cursor en 2026 : Free, Pro, Pro+, Ultra et Teams

La tarification de **Cursor** repose désormais sur un système de crédits d’usage. L’offre gratuite *Hobby* permet de découvrir l’outil avec environ 2 000 complétions par mois, selon MobileAppDaily en juillet 2026, tandis que les offres payantes débloquent davantage d’utilisation des modèles. Cursor propose au total quatre grandes formules – Hobby, Individual, Teams et Enterprise –, l’offre Individual s’échelonnant de 20 à 200 dollars par mois selon Superblocks, une fourchette confirmée par Emergent.sh en juillet 2026. Les prix officiels sont libellés en dollars ; les montants en euros ci-dessous sont des estimations indicatives.

| Offre | Prix mensuel | Prix annuel (−20 %) | Pour qui ? | 
|---|---|---|---|
| Hobby (gratuit) | 0 $ | 0 $ | Découverte, usage occasionnel | 
| Pro | 20 $ (≈ 18 €) | 16 $/mois | Développeur individuel | 
| Pro+ | 60 $ (≈ 55 €) | 48 $/mois | Usage intensif | 
| Ultra | 200 $ (≈ 185 €) | 160 $/mois | Power users, gros volumes | 
| Teams | 40 $/utilisateur | – | Équipes, gestion centralisée | 
| Enterprise | Sur devis | Sur devis | Grandes organisations, SSO, sécurité | 

Point d’attention : en juin 2026, Cursor a fait évoluer son offre Teams en séparant l’usage en deux pools distincts – l’un pour les modèles maison (Composer / Auto), l’autre pour les modèles d’API tierces (Claude, GPT, Gemini). Depuis, l’offre Teams se scinde elle-même en deux niveaux, Standard à 40 dollars par utilisateur et par mois et Premium à 120 dollars, selon NoCode MBA en août 2026 ; à titre de comparaison, une configuration Pro pour trois postes revient à 720 dollars la première année, d’après Credit for Startups en septembre 2026. Ce système de crédits a suscité une certaine confusion dans la communauté, certains utilisateurs regrettant le manque de lisibilité par rapport à un abonnement à prix fixe. Surveillez votre consommation dans le tableau de bord pour éviter les dépassements. La documentation détaillée est disponible sur la page modèles et tarifs de Cursor.

## Cursor face à GitHub Copilot et Claude Code

Faut-il choisir **Cursor** plutôt que ses concurrents ? Tout dépend de votre flux de travail. Copilot reste une extension légère intégrée à votre éditeur existant ; Claude Code privilégie le terminal ; Cursor mise sur un IDE complet et l’orchestration multi-agents. Le tableau suivant synthétise les différences clés.

| Critère | Cursor | GitHub Copilot | Claude Code | 
|---|---|---|---|
| Forme | IDE autonome (fork VS Code) | Extension d’éditeur | Agent en ligne de commande | 
| Prix d’entrée payant | 20 $/mois (Pro) | 10 $/mois (Pro) | Selon offre Claude | 
| Modèles multiples | Oui (Claude, GPT, Gemini, Grok, Composer) | Oui (Claude, GPT, Gemini) | Modèles Claude | 
| Édition multi-fichiers | Excellente (Composer) | Bonne | Excellente | 
| Modèle maison | Composer 2 | Non | Non | 
| Agents en parallèle | Oui (jusqu’à 8) | Partiel | Oui | 

Pour un comparatif détaillé face à ses rivaux directs, consultez nos analyses Claude Code vs Cursor et Windsurf vs Cursor. En résumé : Cursor brille lorsque l’on veut un environnement unifié et agentique, quand Copilot séduit par sa simplicité et son intégration native à l’écosystème GitHub.

## 5 pièges fréquents à éviter avec Cursor

Au-delà des pièges déjà signalés au fil des étapes, voici les erreurs les plus courantes qui plombent l’expérience des nouveaux utilisateurs de **Cursor AI**.

