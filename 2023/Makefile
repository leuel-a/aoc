JAVAC = javac
JAVA = java

JAVA_CLASSPATH = ./build
JAVA_MAIN = Main

INPUT_FILE = input.txt

java: $(JAVA_MAIN).java $(INPUT_FILE)
	$(JAVAC) -d $(JAVA_CLASSPATH) $(JAVA_MAIN).java
	$(JAVA) -classpath $(JAVA_CLASSPATH) $(JAVA_MAIN)
