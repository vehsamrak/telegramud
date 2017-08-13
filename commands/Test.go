package commands

type Test struct{}

func (command Test) GetNames() []string {
    return []string{"test", "тест"}
}

func (command Test) Execute() (string) {
    return "Тест пройден!"
}
