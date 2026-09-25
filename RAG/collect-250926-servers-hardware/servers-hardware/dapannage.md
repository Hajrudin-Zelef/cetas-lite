---
id: collect-250926-servers-hardware/servers-hardware/dapannage
title: "DÃ©pannage"
domain: servers-hardware
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft", "OpenAI"]
dates: []
keywords: ["gemini", "kimi"]
source: docs/RAG/clean4/dapannage.md
source_anchor: ""
source_lines: [1, 178]
sha256: fcb70a9fa1456069b2800829dbdc8959e1edffe83ae2a94195e22f76cc5a5d0a
---

# DÃ©pannage

ProblÃ¨mes courants et comment les rÃ©soudre.

Pour dÃ©boguer les problÃ¨mes avec OpenCode, commencez par vÃ©rifier les journaux et les donnÃ©es locales quâil stocke sur le disque.

Les fichiers journaux sont Ã©crits dansÂ :

- **macOS/Linux** Â :`~/.local/share/opencode/log/`
- **Windows** Â : appuyez sur`WIN+R` et collez`%USERPROFILE%\.local\share\opencode\log`

Les fichiers journaux sont nommÃ©s avec des horodatages (par exemple, `2025-01-09T123456.log`) et les 10 fichiers journaux les plus rÃ©cents sont conservÃ©s.

Vous pouvez dÃ©finir le niveau de journalisation avec lâoption de ligne de commande `--log-level` pour obtenir des informations de dÃ©bogage plus dÃ©taillÃ©es. Par exemple, `opencode --log-level DEBUG`.

opencode stocke les donnÃ©es de session et autres donnÃ©es dâapplication sur le disque Ã lâemplacementÂ :

- **macOS/Linux** Â :`~/.local/share/opencode/`
- **Windows** Â : appuyez sur`WIN+R` et collez`%USERPROFILE%\.local\share\opencode`

Ce rÃ©pertoire contient :

- `auth.json` - DonnÃ©es dâauthentification telles que les clÃ©s API, les jetons OAuth
- `log/` - Journaux dâapplications
- `project/` - DonnÃ©es spÃ©cifiques au projet telles que les donnÃ©es de session et de message
  - Si le projet se trouve dans un dÃ©pÃ´t Git, il est stockÃ© dans `./<project-slug>/storage/`
  - Sâil ne sâagit pas dâun dÃ©pÃ´t Git, il est stockÃ© dans `./global/storage/`
- Si le projet se trouve dans un dÃ©pÃ´t Git, il est stockÃ© dans 

OpenCode Desktop exÃ©cute un serveur OpenCode local (le side-car `opencode-cli`) en arriÃ¨re-plan. La plupart des problÃ¨mes sont causÃ©s par un plugin qui se comporte mal, un cache corrompu ou un mauvais paramÃ¨tre du serveur.

- Quittez complÃ¨tement et relancez lâapplication.
- Si lâapplication affiche un Ã©cran dâerreur, cliquez sur **RedÃ©marrer** et copiez les dÃ©tails de lâerreur.
- macOS uniquementÂ : menu `OpenCode` ->**Recharger la vue Web** (aide si lâinterface utilisateur est vide/gelÃ©e).

Si lâapplication de bureau plante au lancement, se bloque ou se comporte Ã©trangement, commencez par dÃ©sactiver les plugins.

Ouvrez votre fichier de configuration global et recherchez une clÃ© `plugin`.

- **macOS/Linux** Â :`~/.config/opencode/opencode.jsonc` (ou`~/.config/opencode/opencode.json` )
- **macOS/Linux** (anciennes installations)Â :`~/.local/share/opencode/opencode.jsonc`
- **Windows** Â : appuyez sur`WIN+R` et collez`%USERPROFILE%\.config\opencode\opencode.jsonc`

Si vous avez configurÃ© des plugins, dÃ©sactivez-les temporairement en supprimant la clÃ© ou en la dÃ©finissant sur un tableau videÂ :

OpenCode peut Ã©galement charger des plugins locaux Ã partir du disque. Ãcartez-les temporairement (ou renommez le dossier) et redÃ©marrez lâapplication de bureauÂ :

- **Plugins mondiaux**  - **macOS/Linux** Â :`~/.config/opencode/plugins/`
  - **Windows** Â : appuyez sur`WIN+R` et collez`%USERPROFILE%\.config\opencode\plugins`
- **Plugins de projet** (uniquement si vous utilisez une configuration par projet)
  - `<your-project>/.opencode/plugins/`

Si lâapplication recommence Ã fonctionner, rÃ©activez les plugins un par un pour trouver celui Ã lâorigine du problÃ¨me.

Si la dÃ©sactivation des plugins ne rÃ©sout pas le problÃ¨me (ou si lâinstallation dâun plugin est bloquÃ©e), videz le cache afin que OpenCode puisse le reconstruire.

1. Quittez complÃ¨tement OpenCode Desktop.
2. Supprimez le rÃ©pertoire cacheÂ :

- **macOS** Â : Finder ->`Cmd+Shift+G` -> coller`~/.cache/opencode`
- **Linux** Â : supprimez`~/.cache/opencode` (ou exÃ©cutez`rm -rf ~/.cache/opencode` )
- **Windows** Â : appuyez sur`WIN+R` et collez`%USERPROFILE%\.cache\opencode`

1. RedÃ©marrez le bureau OpenCode.

OpenCode Desktop peut soit dÃ©marrer son propre serveur local (par dÃ©faut), soit se connecter Ã un serveur URL que vous avez configurÃ©.

Si vous voyez une boÃ®te de dialogue **Â« Ãchec de la connexion Â»** (ou si lâapplication ne dÃ©passe jamais lâÃ©cran de dÃ©marrage), recherchez un serveur personnalisÃ© URL.

Depuis lâÃ©cran dâaccueil, cliquez sur le nom du serveur (avec le point dâÃ©tat) pour ouvrir le sÃ©lecteur de serveur. Dans la section **Serveur par dÃ©faut**, cliquez sur **Effacer**.

Si votre `opencode.json(c)` contient une section `server`, supprimez-la temporairement et redÃ©marrez lâapplication de bureau.

Si `OPENCODE_PORT` est dÃ©fini dans votre environnement, lâapplication de bureau tentera dâutiliser ce port pour le serveur local.

- DÃ©sactivez `OPENCODE_PORT` (ou choisissez un port libre) et redÃ©marrez.

Sur Linux, certaines configurations Wayland peuvent provoquer des fenÃªtres vides ou des erreurs de composition.

- Si vous Ãªtes sur Wayland et que lâapplication est vide/plante, essayez de la lancer avec `OC_ALLOW_WAYLAND=1` .
- Si cela aggrave les choses, supprimez-le et essayez plutÃ´t de le lancer sous une session X11.

Sur Windows, OpenCode Desktop nÃ©cessite Microsoft Edge **WebView2 Runtime**. Si lâapplication sâouvre sur une fenÃªtre vide ou ne dÃ©marre pas, installez/mettez Ã  jour WebView2 et rÃ©essayez.

Si vous rencontrez des performances lentes, des problÃ¨mes dâaccÃ¨s aux fichiers ou des problÃ¨mes de terminal sur Windows, essayez dâutiliser WSL (Windows Sous-systÃ¨me pour Linux). WSL fournit un environnement Linux qui fonctionne de maniÃ¨re plus transparente avec les fonctionnalitÃ©s de OpenCode.

OpenCode Desktop affiche uniquement les notifications systÃ¨me lorsqueÂ :

- les notifications sont activÃ©es pour OpenCode dans les paramÃ¨tres de votre systÃ¨me dâexploitation, et
- la fenÃªtre de lâapplication nâest pas ciblÃ©e.

Si lâapplication ne dÃ©marre pas et que vous ne pouvez pas effacer les paramÃ¨tres depuis lâinterface utilisateur, rÃ©initialisez lâÃ©tat enregistrÃ© de lâapplication de bureau.

1. Quittez le bureau OpenCode.
2. Recherchez et supprimez ces fichiers (ils se trouvent dans le rÃ©pertoire de donnÃ©es de lâapplication OpenCode Desktop)Â :

- `opencode.settings.dat` (serveur de bureau par dÃ©faut URL)
- `opencode.global.dat` et`opencode.workspace.*.dat` (Ã©tat de lâinterface utilisateur comme les serveurs/projets rÃ©cents)

Pour trouver rapidement le rÃ©pertoire :

- **macOS** Â : Finder ->`Cmd+Shift+G` ->`~/Library/Application Support` (puis recherchez les noms de fichiers ci-dessus)
- **Linux** Â : recherchez sous`~/.local/share` les noms de fichiers ci-dessus
- **Windows** Â : appuyez sur`WIN+R` ->`%APPDATA%` (puis recherchez les noms de fichiers ci-dessus)

Si vous rencontrez des problÃ¨mes avec OpenCodeÂ :

1. **Signaler les problÃ¨mes le GitHub**

La meilleure faÃ§on de signaler des bogues ou de demander des fonctionnalitÃ©s consiste Ã utiliser notre rÃ©fÃ©rentiel GitHubÂ :

Avant de crÃ©er un nouveau problÃ¨me, recherchez les problÃ¨mes existants pour voir si votre problÃ¨me a dÃ©jÃ Ã©tÃ© signalÃ©.

1. **Rejoignez notre Discord**

Pour obtenir de lâaide en temps rÃ©el et une discussion communautaire, rejoignez notre serveur DiscordÂ :

Voici quelques problÃ¨mes courants et comment les rÃ©soudre.

1. VÃ©rifiez les journaux pour les messages dâerreur
2. Essayez dâexÃ©cuter avec `--print-logs` pour voir la sortie dans le terminal
3. Assurez-vous dâavoir la derniÃ¨re version avec `opencode upgrade`

1. Essayez de vous rÃ©authentifier avec la commande `/connect` dans le TUI
2. VÃ©rifiez que vos clÃ©s API sont valides
3. Assurez-vous que votre rÃ©seau autorise les connexions au API du fournisseur.

1. VÃ©rifiez que vous Ãªtes authentifiÃ© auprÃ¨s du fournisseur
2. VÃ©rifiez que le nom du modÃ¨le dans votre configuration est correct
3. Certains modÃ¨les peuvent nÃ©cessiter un accÃ¨s ou des abonnements spÃ©cifiques

Si vous rencontrez `ProviderModelNotFoundError`, vous avez probablement tort
faire rÃ©fÃ©rence Ã  un modÃ¨le quelque part.
Les modÃ¨les doivent Ãªtre rÃ©fÃ©rencÃ©s comme suitÂ : `<providerId>/<modelId>`

ExemplesÂ :

- `openai/gpt-4.1`
- `openrouter/google/gemini-2.5-flash`
- `opencode/kimi-k2`

Pour dÃ©terminer Ã  quels modÃ¨les vous avez accÃ¨s, exÃ©cutez `opencode models`

Si vous rencontrez une ProviderInitError, vous avez probablement une configuration non valide ou corrompue.

Pour rÃ©soudre ce problÃ¨meÂ :

1. 
Tout dâabord, vÃ©rifiez que votre fournisseur est correctement configurÃ© en suivant le guide du fournisseur
2. 
Si le problÃ¨me persiste, essayez dâeffacer votre configuration stockÃ©eÂ :

Sur Windows, appuyez sur `WIN+R` et supprimezÂ : `%USERPROFILE%\.local\share\opencode`

1. RÃ©-authentifiez-vous auprÃ¨s de votre fournisseur Ã  lâaide de la commande `/connect` dans le TUI.

Si vous rencontrez des erreurs dâappel API, cela peut Ãªtre dÃ» Ã des packages de fournisseurs obsolÃ¨tes. opencode installe dynamiquement les packages du fournisseur (OpenAI, Anthropic, Google, etc.) selon les besoins et les met en cache localement.

Pour rÃ©soudre les problÃ¨mes liÃ©s au package du fournisseurÂ :

1. 
Videz le cache du package du fournisseurÂ :

Sur Windows, appuyez sur `WIN+R` et supprimezÂ : `%USERPROFILE%\.cache\opencode`

1. RedÃ©marrez opencode pour rÃ©installer les derniers packages du fournisseur

Cela forcera opencode Ã tÃ©lÃ©charger les versions les plus rÃ©centes des packages du fournisseur, ce qui rÃ©sout souvent les problÃ¨mes de compatibilitÃ© avec les paramÃ¨tres du modÃ¨le et les modifications de API.

Les utilisateurs de Linux doivent disposer de lâun des utilitaires de presse-papiers suivants installÃ©s pour que la fonctionnalitÃ© copier/coller fonctionneÂ :

**Pour les systÃ¨mes X11Â :**

**Pour les systÃ¨mes WaylandÂ :**

**Pour les environnements sans tÃªteÂ :**

opencode dÃ©tectera si vous utilisez Wayland et prÃ©fÃ©rez `wl-clipboard`, sinon il essaiera de trouver les outils du presse-papiers dans lâordreÂ : `xclip` et `xsel`.
