package chat

func resolveEffort(agent, think bool, text, requested string) string {
	switch requested {
	case "low", "medium", "high":
		return requested
	}
	if !agent && !think {
		return ""
	}
	if len([]rune(text)) > 400 {
		return "medium"
	}
	return "low"
}
