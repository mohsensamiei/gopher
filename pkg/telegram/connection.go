package telegram

func Connect(configs Configs) (*Connection, error) {
	var res Response[User]
	if err := request(configs.TelegramToken, getMe, nil, &res); err != nil {
		return nil, err
	}
	return &Connection{
		Configs: configs,
		User:    res.Result,
	}, nil
}

type Connection struct {
	Configs
	User
}
