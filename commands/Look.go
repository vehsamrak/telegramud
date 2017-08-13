package commands

type Look struct{}

func (command Look) GetNames() []string {
	return []string{"look", "смотреть"}
}

func (command Look) Execute() (string) {
	return "Ты видишь контуры этого мира."
}
