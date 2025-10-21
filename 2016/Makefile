DOTNET = dotnet

all: build run

build:
	@$(DOTNET) build --nologo --no-restore --verbosity quiet > /dev/null 2>&1

run:
	@$(DOTNET) run --no-restore --verbosity quiet
