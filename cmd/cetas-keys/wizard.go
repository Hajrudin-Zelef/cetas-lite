package main

import (
	"fmt"
	"strings"

	"cetas-lite/internal/keysetup"
	"cetas-lite/internal/securevault"
)

// runWizard — assistant interactif façon setup.py : on lance cetas-keys sans
// argument, il guide pas à pas (création du coffre, clés, validation,
// export .env, synchro). Aucune ligne de commande à retenir.
func (c *cli) runWizard() int {
	printSetupBanner()
	pw, ok := c.wizardOpenOrCreate()
	if !ok {
		return exitErr
	}

	for {
		fmt.Println()
		frame("QUE VOULEZ-VOUS FAIRE ?")
		fmt.Println(white("  1. Ajouter / Modifier une clé API"))
		fmt.Println(white("  2. Tester les clés enregistrées"))
		fmt.Println(white("  3. Voir les clés (masquées)"))
		fmt.Println(white("  4. Supprimer une clé"))
		fmt.Println(white("  5. Synchroniser vers l'application"))
		fmt.Println(white("  6. Régénérer le .env scellé"))
		fmt.Println(white("  7. Changer le mot de passe maître"))
		fmt.Println(white("  0. Quitter"))
		fmt.Println()
		fmt.Print(white("  Votre choix : "))
		switch c.readLine() {
		case "1":
			c.wizardAdd(pw)
		case "2":
			c.wizardTest(pw)
		case "3":
			c.wizardList(pw)
		case "4":
			c.wizardDelete(pw)
		case "5":
			c.wizardSync(pw)
		case "6":
			if err := c.exportEnv(pw); err != nil {
				fmt.Printf("  Erreur : %v\n", err)
			}
		case "7":
			newPw, changed := c.changePassword(pw)
			if changed {
				pw = newPw
			}
		case "0":
			if err := c.exportEnvQuiet(pw); err != nil {
				errLine(fmt.Sprintf("Erreur d'export .env : %v", err))
				return exitErr
			}
			fmt.Println(gray("  Au revoir.\n"))
			return exitOK
		default:
			errLine("Choix invalide.")
		}
	}
}

// wizardOpenOrCreate — crée le coffre s'il n'existe pas (avec configuration
// rapide des providers principaux), sinon demande le mot de passe maître
// (3 tentatives).
func (c *cli) wizardOpenOrCreate() (string, bool) {
	if !c.st.Exists() {
		frame("ETAPE 1 : Coffre chiffré")
		fmt.Println()
		if !keysetup.PepperSet() {
			fmt.Println(yellow("  CETAS_PEPPER non défini — protection réduite."))
		}
		if suggest, err := securevault.GeneratePassword(20); err == nil {
			fmt.Printf("  Suggestion de mot de passe fort : %s\n\n", suggest)
		}
		pw, ok := c.readPasswordTwice("  Nouveau mot de passe maître")
		if !ok {
			return "", false
		}
		if err := c.st.Init(pw); err != nil {
			errLine(fmt.Sprintf("Erreur : %v\n", err))
			return "", false
		}
		okLine(fmt.Sprintf("Coffre créé : %s", c.st.Path()))
		infoLine("Chiffrement : AES-256-GCM | Scrypt (N=2^16) | pepper CETAS_PEPPER")
		c.wizardQuickSetup(pw)
		return pw, true
	}
	for attempt := 1; attempt <= 3; attempt++ {
		pw := c.readSecret(white("  Mot de passe du coffre : "))
		if _, err := c.st.Load(pw); err == nil {
			okLine("Coffre ouvert.\n")
			return pw, true
		}
		errLine(fmt.Sprintf("Mot de passe incorrect (%d tentative(s) restante(s)).\n", 3-attempt))
	}
	errLine("Accès refusé.\n")
	return "", false
}

// wizardQuickSetup — à la première création : propose les providers principaux
// un par un (Entrée = passer), chacun testé en direct avant stockage.
func (c *cli) wizardQuickSetup(pw string) {
	fmt.Println()
	fmt.Println("  ── Providers principaux ──")
	fmt.Println("  (Appuyez sur Entrée pour passer, chaque clé est testée avant stockage)")
	fmt.Println()
	if !c.confirm("  Configurer maintenant ? [O/n] : ") {
		return
	}
	for _, id := range []string{"openrouter", "deepseek", "groq", "openai", "anthropic", "mistral", "google", "opencode", "opencode-go"} {
		p, found := keysetup.ByID(id)
		if !found {
			continue
		}
		fmt.Println()
		c.configureProviderQuick(pw, p)
	}
	if err := c.exportEnv(pw); err != nil {
		fmt.Printf("  Erreur d'export .env : %v\n", err)
	}
}

// configureProviderQuick — variante silencieuse : une seule saisie, Entrée = passer.
func (c *cli) configureProviderQuick(pw string, p keysetup.Provider) {
	key := strings.TrimSpace(c.readSecret(fmt.Sprintf("  %-15s (%s) : ", p.Label, p.Hint)))
	defer wipe([]byte(key))
	if key == "" {
		fmt.Println(gray(fmt.Sprintf("  %s ignoré.", p.Label)))
		return
	}
	fmt.Printf("  Test de la clé %s ... ", p.Label)
	ok, reason := keysetup.TestKey(p, key)
	if !ok {
		fmt.Println(red(fmt.Sprintf("  Échec : %s", reason)))
		fmt.Println(gray(fmt.Sprintf("  %s ignoré (recommencez via le menu 1).", p.Label)))
		return
	}
	fmt.Println(green("OK"))
	if err := c.st.SetAPIKey(pw, p.ID, key); err != nil {
		errLine(fmt.Sprintf("Erreur d'enregistrement : %v", err))
	}
}

func (c *cli) wizardAdd(pw string) {
	for {
		id := c.chooseProvider(pw)
		if id == "" {
			return
		}
		p, _ := keysetup.ByID(id)
		c.configureProvider(pw, p)
	}
}

func (c *cli) wizardTest(pw string) {
	data, err := c.st.Load(pw)
	if err != nil {
		errLine(fmt.Sprintf("Erreur : %v", err))
		return
	}
	keys := keysetup.APIKeys(data)
	if len(keys) == 0 {
		infoLine("Aucune clé enregistrée.")
		return
	}
	failed := 0
	for _, id := range sortedKeys(keys) {
		p, found := keysetup.ByID(id)
		if !found {
			continue
		}
		if c.testOne(p, keys[id]) != exitOK {
			failed++
		}
	}
	if failed == 0 {
		okLine("Toutes les clés sont valides.")
	} else {
		errLine(fmt.Sprintf("%d clé(s) en échec.", failed))
	}
}

func (c *cli) wizardList(pw string) {
	data, err := c.st.Load(pw)
	if err != nil {
		errLine(fmt.Sprintf("Erreur : %v", err))
		return
	}
	keys := keysetup.APIKeys(data)
	if len(keys) == 0 {
		infoLine("Aucune clé API enregistrée.")
		return
	}
	subtitle("Clés enregistrées")
	for _, id := range sortedKeys(keys) {
		label := id
		if p, found := keysetup.ByID(id); found {
			label = p.Label
		}
		fmt.Printf("    %-15s %s\n", label, keysetup.Mask(keys[id]))
	}
}

func (c *cli) wizardDelete(pw string) {
	data, err := c.st.Load(pw)
	if err != nil {
		fmt.Printf("  Erreur : %v\n", err)
		return
	}
	keys := keysetup.APIKeys(data)
	if len(keys) == 0 {
		infoLine("Aucune clé à supprimer.")
		return
	}
	ids := sortedKeys(keys)
	fmt.Println("  Clés :")
	for i, id := range ids {
		label := id
		if p, found := keysetup.ByID(id); found {
			label = p.Label
		}
		fmt.Printf("    %d. %s\n", i+1, label)
	}
	fmt.Print(white("  Numéro à supprimer (Entrée=annuler) : "))
	sel := c.readLine()
	if sel == "" {
		return
	}
	idx := -1
	for i, id := range ids {
		_ = id
		if fmt.Sprint(i+1) == sel {
			idx = i
			break
		}
	}
	if idx < 0 {
		errLine("Choix invalide.")
		return
	}
	id := ids[idx]
	label := id
	if p, found := keysetup.ByID(id); found {
		label = p.Label
	}
	if !c.confirm(red(fmt.Sprintf("  Supprimer %s ? [o/N] : ", label))) {
		infoLine("Annulé.")
		return
	}
	deleted, err := c.st.DeleteAPIKey(pw, id)
	if err != nil {
		fmt.Printf("  Erreur : %v\n", err)
		return
	}
	if deleted {
		if err := c.exportEnv(pw); err != nil {
			fmt.Printf("  Erreur d'export .env : %v\n", err)
		}
		okLine(fmt.Sprintf("%s supprimé.", label))
	}
}

func (c *cli) wizardSync(pw string) {
	if c.syncToApp(pw, syncAPIURL()) != exitOK {
		// Le détail de l'erreur est déjà affiché par syncToApp.
	}
}

// changePassword — variante réutilisable (wizard + commande passwd).
// Retourne le nouveau mot de passe en cas de succès.
func (c *cli) changePassword(pw string) (string, bool) {
	frame("Changement du mot de passe maître")
	newPw, ok := c.readPasswordTwice("  Nouveau mot de passe")
	if !ok {
		return "", false
	}
	if err := c.st.ChangePassword(pw, newPw); err != nil {
		errLine(fmt.Sprintf("Erreur : %v", err))
		return "", false
	}
	okLine("Mot de passe modifié, coffre re-chiffré.")
	return newPw, true
}

// exportEnvQuiet — export sans affichage (sortie de l'assistant).
func (c *cli) exportEnvQuiet(pw string) error {
	data, err := c.st.Load(pw)
	if err != nil {
		return err
	}
	proxyKey, err := c.st.ProxyKey(pw)
	if err != nil {
		return err
	}
	_, err = keysetup.WriteEnvFile(c.envPath, keysetup.APIKeys(data), proxyKey)
	return err
}

// syncAPIURL — URL de l'application pour la synchronisation.
func syncAPIURL() string {
	return apiFlag(nil)
}
