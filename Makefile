SHELL=/bin/bash

BASEDIR=$(PWD)
OUT="$(BASEDIR)/.out/bin/jsc"
SRC="$(BASEDIR)/jsc"

# Color codes
Green      := \033[0;32m
Red        := \033[0;31m
Yellow     := \033[0;33m
Blue       := \033[0;34m
Magenta    := \033[0;35m
Cyan       := \033[0;36m
Orange     := \033[0;38;5;208m
White      := \033[0;37m
Purple     := \033[0;35m
Black      := \033[0;30m
Gray       := \033[0;90m

Reset      := \033[0m

# Run the application
run: force-build
	@echo -e "${Yellow}🚧Running Just Simple Cloud (jsc)...${Reset}"
	$(OUT)

# Display help message
help-message:
	@echo -e "${Green}Makefile for Just Simple Cloud (jsc)${Reset}"
	@echo ""
	@echo -e "${Yellow}Available targets:${Reset}"
	@echo -e "  ${Green}build${Reset}     		- 🚧Build the Go application"
	@echo -e "  ${Green}force-build${Reset} 	- 🚧Clean and build the Go application"
	@echo -e "  ${Green}install${Reset}   		- 🚧Install the Go application"
	@echo -e "  ${Green}clean${Reset}     		- 🚧Clean the built application"
	@echo -e "  ${Green}run${Reset}       		- 🚧Run the application"
	@echo -e "  ${Green}version${Reset}   		- 🚧Display version information"
	@echo -e "  ${Green}test${Reset}      		- 🚧Test the application"
	@echo -e "  ${Green}format${Reset}    		- 🚧Format the code"
	@echo -e "  ${Green}help-message${Reset} 	- 📋Display this help message"
	@echo ""

# Build the Go application
build:
	@echo -e "${Yellow}🚧Building Just Simple Cloud (jsc)...${Reset}"
	@go build -C $(SRC) -o $(OUT) main.go
	@echo -e "${Green}Output binary: ${Reset}$(OUT)"
# Force rebuild by cleaning first
force-build: clean build

# Install the Go application
install:
	@echo -e "${Yellow}🚧Installing Just Simple Cloud (jsc)...${Reset}"
	@go install $(SRC)
# Clean the built application
clean:
	@echo -e "${Yellow}🚧Cleaning built files...${Reset}"
	@rm -f $(OUT)

# Display version information
version:
	@echo -e "${Yellow}🚧Just Simple Cloud (jsc) version:${Reset}"
	$(OUT) --version
# Test the application
test:
	@echo -e "${Yellow}🚧Testing Just Simple Cloud (jsc)...${Reset}"
	go test ./...
# Format the code
format:
	@echo -e "${Yellow}🚧Formatting code...${Reset}"
	go fmt ./...
	@echo -e "${Green}Code formatting complete.${Reset}"


