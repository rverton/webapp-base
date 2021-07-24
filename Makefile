tailwind_cmd = npx tailwindcss -c ./public/tailwind.config.js -i ./public/tailwind.css -o ./public/dist/styles.css

all: app tailwind-prod

watch:
	$(tailwind_cmd)	--watch

tailwind-prod:
	NODE_ENV=production $(tailwind_cmd) 

app:
	go build -ldflags "-X main.commitHash=$$(git rev-parse --short HEAD) -X main.commitDate=$$(git log -1 --format=%ct)"
