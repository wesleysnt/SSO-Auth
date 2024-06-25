package seeders

func Execute() error {
	adminSeeder := AdminSeeder{}
	adminSeeder.Run()

	return nil
}
