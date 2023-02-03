package internal

type CloudFactory struct {
	Client CloudClient
	Prompt CloudPrompt
	Config UserConfig
}
