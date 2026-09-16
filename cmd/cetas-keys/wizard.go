package main

import (
	"fmt"

	"cetas-lite/internal/keysetup"
	"cetas-lite/internal/securevault"
)

// runWizard — assistant interactif façon setup.py : on lance cetas-keys sans
// argument, il guide pas à pas (création du coffre, clés, validation,
// export .env, synchro). Aucune ligne de commande à retenir.
func (c *cli) runWizard() int {
	fmt.Println()
	fmt.Println("  ╔══════════════════════════════════════════════════╗")
	fmt.Println("  ║        CETAS LITE — Gestion des clés API         ║")
	fmt.Println("  ╚══════════════════════════════════════════════════╝")
	fmt.Println()

	pw, ok := c.wizardOpenOrCreate()
	if !ok {
		return exitErr
	}

	for {
		fmt.Println()
		fmt.Println("  ┌─ Que voulez-vous faire ? ───────────────────────┐")
		fmt.Println("  │  1. Ajouter / modifier une clé API               │")
		fmt.Println("  │  2. Tester les clés enregistrées                 │")
		fmt.Println("  │  3. Voir les clés (masquées)                     │")
		fmt.Println("  │  4. Supprimer une clé                            │")
		fmt.Println("  │  5. Synchroniser vers l'application              │")
		fmt.Println("  │  6. Régénérer le .env scellé                     │")
		fmt.Println("  │  7. Changer le mot de passe maître               │")
		fmt.Println("  │  0. Quitter                                      │")
		fmt.Println("  └──────────────────────────────────────────────────┘")
		fmt.Print("  Votre choix : ")
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
				fmt.Printf("  Erreur d'export .env : %v\n", err)
				return exitErr
			}
			fmt.Println("  Au revoir.")
			return exitOK
		default:
			fmt.Println("  Choix invalide.")
		}
	}
}

// wizardOpenOrCreate — crée le coffre s'il n'existe pas (avec configuration
// rapide des providers principaux), sinon demande le mot de passe maître
// (3 tentatives).
func (c *cli) wizardOpenOrCreate() (string, bool) {
	if !c.st.Exists() {
		fmt.Println("  ── Création du coffre ──")
		if !keysetup.PepperSet() {
			fmt.Println("  Note : CETAS_PEPPER n'est pas défini (recommandé en production).")
		}
		if suggest, err := securevault.GeneratePassword(20); err == nil {
			fmt.Printf("  Suggestion de mot de passe fort : %s\n\n", suggest)
		}
		pw, ok := c.readPasswordTwice("  Nouveau mot de passe maître")
		if !ok {
			return "", false
		}
		if err := c.st.Init(pw); err != nil {
			fmt.Printf("  Erreur : %v\n", err)
			return "", false
		}
		fmt.Printf("  Coffre créé : %s\n", c.st.Path())
		fmt.Println("  Chiffrement : AES-256-GCM | Scrypt (N=2^16) | pepper CETAS_PEPPER")
		c.wizardQuickSetup(pw)
		return pw, true
	}
	for attempt := 1; attempt <= 3; attempt++ {
		pw := c.readSecret("  Mot de passe maître : ")
		if _, err := c.st.Load(pw); err == nil {
			fmt.Println("  Coffre ouvert.")
			return pw, true
		} else {
			fmt.Printf("  Mot de passe incorrect (%d tentative(s) restante(s)).\n", 3-attempt)
		}
	}
	fmt.Println("  Accès refusé.")
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
	key := c.readSecret(fmt.Sprintf("  %-15s (%s) : ", p.Label, p.Hint))
	defer wipe([]byte(key))
	if key == "" {
		fmt.Printf("  %s ignoré.\n", p.Label)
		return
	}
	fmt.Printf("  Test de la clé %s ... ", p.Label)
	ok, reason := keysetup.TestKey(p, key)
	if !ok {
		fmt.Printf("ÉCHEC — %s\n", reason)
		fmt.Printf("  %s ignoré (recommencez via le menu 1).\n", p.Label)
		return
	}
	fmt.Println("OK")
	if err := c.st.SetAPIKey(pw, p.ID, key); err != nil {
		fmt.Printf("  Erreur d'enregistrement : %v\n", err)
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
		fmt.Printf("  Erreur : %v\n", err)
		return
	}
	keys := keysetup.APIKeys(data)
	if len(keys) == 0 {
		fmt.Println("  Aucune clé enregistrée.")
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
		fmt.Println("  Toutes les clés sont valides.")
	} else {
		fmt.Printf("  %d clé(s) en échec.\n", failed)
	}
}

func (c *cli) wizardList(pw string) {
	data, err := c.st.Load(pw)
	if err != nil {
		fmt.Printf("  Erreur : %v\n", err)
		return
	}
	keys := keysetup.APIKeys(data)
	if len(keys) == 0 {
		fmt.Println("  Aucune clé API enregistrée.")
		return
	}
	fmt.Println("  Clés API enregistrées :")
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
		fmt.Println("  Aucune clé à supprimer.")
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
	fmt.Print("  Numéro à supprimer (Entrée=annuler) : ")
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
		fmt.Println("  Choix invalide.")
		return
	}
	id := ids[idx]
	label := id
	if p, found := keysetup.ByID(id); found {
		label = p.Label
	}
	if !c.confirm(fmt.Sprintf("  Supprimer %s ? [o/N] : ", label)) {
		fmt.Println("  Annulé.")
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
		fmt.Printf("  %s supprimé.\n", label)
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
	fmt.Println("  ── Changement du mot de passe maître ──")
	newPw, ok := c.readPasswordTwice("  Nouveau mot de passe")
	if !ok {
		return "", false
	}
	if err := c.st.ChangePassword(pw, newPw); err != nil {
		fmt.Printf("  Erreur : %v\n", err)
		return "", false
	}
	fmt.Println("  Mot de passe modifié, coffre re-chiffré.")
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
