# This makefile builds all the .yaml sources to .png

SOURCES := $(wildcard *.yaml)

.PHONY: all
all: $(SOURCES:.yaml=.png)

%.png : %.yaml
	../comprehend $^ --mime-type image/png > $@

.PHONY: clean
clean:
	rm -f *.png
