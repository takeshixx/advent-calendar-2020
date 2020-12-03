web-build:
	sudo docker build --tag advent-website -f _website/Dockerfile _website

web-up: web-build
	sudo docker save advent-website | pv | ssh xmas sudo docker load

web: web-up
	ssh xmas "sudo docker-compose stop website && sudo docker-compose up -d website"

run:
	sudo docker-compose up -d

sync:
	scp docker-compose.yml xmas:

day01-build:
	sudo docker build --tag day01 -f xmas-cookie2/Dockerfile xmas-cookie2

day01-up: day01-build
	sudo docker save day01 | pv | ssh xmas sudo docker load

day02-build:
	sudo docker build --tag day02 -f dtls/Dockerfile dtls

day02-up: day02-build
	sudo docker save day02 | pv | ssh xmas sudo docker load

day03-build:
	sudo docker build --tag day03 -f elf/Dockerfile elf

day03-up: day03-build
	sudo docker save day03 | pv | ssh xmas sudo docker load

day04-build:
	sudo docker build --tag day04 -f xmas-socks/Dockerfile xmas-socks

day04-up: day04-build
	sudo docker save day04 | pv | ssh xmas sudo docker load



day10-build:
	sudo docker build --tag day10 -f redstar/Dockerfile redstar

day10-up: day10-build
	sudo docker save day10 | pv | ssh xmas sudo docker load
