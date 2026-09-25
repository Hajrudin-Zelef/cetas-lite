---
id: collect-250926-servers-hardware/servers-hardware/github
title: "GitHub"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "valuation"]
source: docs/RAG/clean4/github.md
source_anchor: ""
source_lines: [1, 97]
sha256: 65a3ae3140843abc9d5fe2ffd15821a4fb0db2b1d9f5311e07f240edaa586efb
---

# GitHub

Utilisez OpenCode dans les issues et les pull-requests GitHub.

OpenCode sâintÃ¨gre Ã  votre flux de travail GitHub. Mentionnez `/opencode` ou `/oc` dans votre commentaire, et OpenCode exÃ©cutera des tÃ¢ches dans votre runner GitHub Actions.

- **Triage des issues** : demandez Ã  OpenCode dâexaminer une issue et de vous lâexpliquer.
- **Correction et implÃ©mentation** : demandez Ã  OpenCode de rÃ©soudre un problÃ¨me ou dâimplÃ©menter une fonctionnalitÃ©. Il travaillera dans une nouvelle branche et soumettra une PR avec tous les changements.
- **SÃ©curisÃ©** : OpenCode sâexÃ©cute Ã  lâintÃ©rieur de vos runners GitHub.

ExÃ©cutez la commande suivante dans un projet qui se trouve dans un dÃ©pÃ´t GitHub :

Cela vous guidera dans lâinstallation de lâapplication GitHub, la crÃ©ation du workflow et la configuration des secrets.

Ou vous pouvez le configurer manuellement.

1. **Installez lâapplication GitHub**

Rendez-vous sur **github.com/apps/opencode-agent**. Assurez-vous quâil est installÃ© sur le dÃ©pÃ´t cible.

1. **Ajouter le workflow**

Ajoutez le fichier de workflow suivant Ã  `.github/workflows/opencode.yml` dans votre dÃ©pÃ´t. Assurez-vous de dÃ©finir les clÃ©s `model` appropriÃ©es et les clÃ©s API requises dans `env`.

1. **Stockez les clÃ©s API dans les secrets**

Dans les **paramÃ¨tres** de votre organisation ou de votre projet, dÃ©veloppez **Secrets et variables** sur la gauche et sÃ©lectionnez **Actions**. Puis ajoutez les clÃ©s API requises.

- `model` : Le modÃ¨le Ã  utiliser avec OpenCode. Prend le format`provider/model` . Ceci est**obligatoire** .
- `agent` : lâagent Ã  utiliser. Doit Ãªtre un agent primaire. Revient Ã`default_agent` Ã  partir de la configuration ou Ã`"build"` sâil nâest pas trouvÃ©.
- `share` : sâil faut partager la session OpenCode. La valeur par dÃ©faut est**true** pour les rÃ©fÃ©rentiels publics.
- `prompt` : prompt personnalisÃ© facultatif pour remplacer le comportement par dÃ©faut. Utilisez-le pour personnaliser la faÃ§on dont OpenCode traite les demandes.
- `token` : jeton dâaccÃ¨s GitHub facultatif pour effectuer des opÃ©rations telles que la crÃ©ation de commentaires, le commit de modifications et lâouverture de pull requests. Par dÃ©faut, OpenCode utilise le jeton dâaccÃ¨s Ã  lâinstallation de lâapplication OpenCode GitHub, de sorte que les commits, les commentaires et les pull requests apparaissent comme provenant de lâapplication.

Vous pouvez Ã©galement utiliser le `GITHUB_TOKEN`](https://docs.github.com/en/actions/tutorials/authenticate-with-github_token) intÃ©grÃ© du runner GitHub Actions sans installer lâapplication OpenCode GitHub. Assurez-vous simplement dâaccorder les autorisations requises dans votre workflow :

Vous pouvez Ã©galement utiliser un jeton dâaccÃ¨s personnel(PAT) si vous prÃ©fÃ©rez.

OpenCode peut Ãªtre dÃ©clenchÃ© par les Ã©vÃ©nements GitHub suivants :

| Type dâÃ©vÃ©nement | DÃ©clenchÃ© par | DÃ©tails | 
|---|---|---|
| `issue_comment` | Commentaire sur une issue ou une PR | Mentionnez `/opencode` ou`/oc` dans votre commentaire. OpenCode lit le contexte et peut crÃ©er des branches, ouvrir des PR ou rÃ©pondre. | 
| `pull_request_review_comment` | Commentaire sur des lignes de code spÃ©cifiques dans une PR | Mentionnez `/opencode` ou`/oc` lors de la rÃ©vision du code. OpenCode reÃ§oit le chemin du fichier, les numÃ©ros de ligne et le contexte de comparaison. | 
| `issues` | Issue ouverte ou modifiÃ©e | DÃ©clenchez automatiquement OpenCode lorsque des issues sont crÃ©Ã©es ou modifiÃ©es. NÃ©cessite une entrÃ©e `prompt` . | 
| `pull_request` | PR ouverte ou mise Ã jour | DÃ©clenchez automatiquement OpenCode lorsque les PR sont ouvertes, synchronisÃ©es ou rouvertes. Utile pour les revues automatisÃ©es. | 
| `schedule` | Planification basÃ©e sur Cron | ExÃ©cutez OpenCode selon un planning. NÃ©cessite une entrÃ©e `prompt` . La sortie va aux journaux et aux PR (pas de commentaire sur les issues). | 
| `workflow_dispatch` | DÃ©clenchement manuel depuis lâinterface utilisateur GitHub | DÃ©clenchez OpenCode Ã  la demande via lâonglet Actions. NÃ©cessite une entrÃ©e `prompt` . La sortie va aux journaux et aux PR. | 

ExÃ©cutez OpenCode selon un planning pour effectuer des tÃ¢ches automatisÃ©es :

Pour les Ã©vÃ©nements planifiÃ©s, lâentrÃ©e `prompt` est **obligatoire** car il nây a aucun commentaire pour extraire les instructions. Les workflows planifiÃ©s sâexÃ©cutent sans contexte utilisateur pour vÃ©rifier les autorisations. Le workflow doit donc accorder `contents: write` et `pull-requests: write` si vous vous attendez Ã  ce que OpenCode crÃ©e des branches ou des PR.

Examinez automatiquement les PR lorsquâils sont ouverts ou mis Ã jour :

Pour les Ã©vÃ©nements `pull_request`, si aucun `prompt` nâest fourni, OpenCode examine par dÃ©faut la pull request.

Triez automatiquement les nouvelles issues. Cet exemple filtre les comptes datant de plus de 30 jours pour rÃ©duire le spam :

Pour les Ã©vÃ©nements `issues`, lâentrÃ©e `prompt` est **obligatoire** car il nây a aucun commentaire Ã  partir duquel extraire les instructions.

Remplacez lâinvite par dÃ©faut pour personnaliser le comportement de OpenCode pour votre workflow.

Ceci est utile pour appliquer des critÃ¨res dâÃ©valuation spÃ©cifiques, des normes de codage ou des domaines dâintervention pertinents pour votre projet.

Voici quelques exemples de la faÃ§on dont vous pouvez utiliser OpenCode dans GitHub.

- **Expliquer une issue**

Ajoutez ce commentaire dans une issue GitHub.

OpenCode lira lâintÃ©gralitÃ© du fil de discussion, y compris tous les commentaires, et rÃ©pondra avec une explication claire.

- **RÃ©soudre une issue**

Dans une issue GitHub, dites :

Et OpenCode crÃ©era une nouvelle branche, mettra en Åuvre les modifications et ouvrira une PR avec les modifications.

- **Examinez les PR et apportez des modifications**

Laissez le commentaire suivant sur une PR GitHub.

OpenCode mettra en Åuvre la modification demandÃ©e et la validera dans la mÃªme PR.

- **Revue de lignes de code spÃ©cifiques**

Laissez un commentaire directement sur les lignes de code dans lâonglet âFichiersâ de la PR. OpenCode dÃ©tecte automatiquement le fichier, les numÃ©ros de ligne et le contexte de comparaison pour fournir des rÃ©ponses prÃ©cises.

Lorsquâil commente des lignes spÃ©cifiques, OpenCode reÃ§oit :

- Le fichier exact en cours dâexamen
- Les lignes de code spÃ©cifiques
- Le contexte diffÃ©rentiel environnant
- Informations sur le numÃ©ro de ligne

Cela permet des requÃªtes plus ciblÃ©es sans avoir besoin de spÃ©cifier manuellement les chemins de fichiers ou les numÃ©ros de ligne.
