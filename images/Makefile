# Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

UML_FILES=$(wildcard ./*.plantuml)
UML_SVG_FILES=$(patsubst %.plantuml,%.uml.svg,$(UML_FILES))
UML_PNG_FILES=$(patsubst %.plantuml,%.uml.png,$(UML_FILES))

DOT_FILES=$(wildcard ./*.dot)
DOT_PNG_FILES=$(patsubst %.dot,%.dot.png,$(DOT_FILES))

DITAA_FILES=$(wildcard ./*.ditaa.txt)
DITAA_PNG_FILES=$(patsubst %.ditaa.txt,%.ditaa.png,$(DITAA_FILES))

default: $(UML_PNG_FILES) $(DOT_PNG_FILES) $(DITAA_PNG_FILES)
	@echo "ok"

clean:
	-rm *.dot.png
	-rm *.uml.png
	-rm *.uml.svg
	-rm *.ditaa.png

%.uml.svg: %.plantuml
	cat $< | docker run --rm -i chai2010/ibox:plantuml >  $@

%.uml.png: %.plantuml
	cat $< | docker run --rm -i chai2010/ibox:plantuml -tpng >  $@

%.dot.png: %.dot
	dot -Tpng -o $@ $<

%.ditaa.png: %.ditaa.txt
	docker run --rm -it  -v `pwd`:/root chai2010/ibox:ditaa $< $@
