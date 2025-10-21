JAVAC = javac
JAVA = java

JAVA_CLASSPATH = ./build/
JAVA_SOURCE_PATH = ../
JAVA_MAIN = Main

GO = go
GO_MAIN = main

INPUT_FILE = input.txt

go: $(GO_MAIN).go $(INPUT_FILE)
	$(GO) build -o $(GO_MAIN).out $(GO_MAIN).go

java: $(JAVA_MAIN).java $(INPUT_FILE)
	@$(JAVAC) -sourcepath $(JAVA_SOURCE_PATH) -d $(JAVA_CLASSPATH) $(JAVA_MAIN).java
	@$(JAVA) -classpath $(JAVA_CLASSPATH) $(JAVA_MAIN)
