SHELL=/bin/bash

.PHONY: fileParserCompile
fileParserCompile:
	cd srv && go run CloverHealth/srv
