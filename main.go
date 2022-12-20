package main

func main() {
	ctx, cancel, interval := loadConfig()
	defer cancel()

	jiggling(ctx, interval)
}
