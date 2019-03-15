TITLE = high-performance-go-workshop
LANG = en
OUTPUT = $(CURDIR)/output
COMMON = $(CURDIR)/common
SRCDIR = $(CURDIR)/src
SRCFILES = $(wildcard $(SRCDIR)/*/*.go)
SITE = $(OUTPUT)/$(TITLE)
DOCBOOK = /opt/local/share/xsl/docbook-xsl-ns
STYLESHEET = $(DOCBOOK)/xhtml5/chunk.xsl
ICONS = $(DOCBOOK)/images
IMAGES = $(SITE)/images
DIRS = $(SITE) $(IMAGES) $(SITE)/$(LANG) $(OUTPUT)
GIT_VERSION := $(shell git describe --abbrev=6 --dirty --always)
GIT_DATE := $(shell git log -1 --format=%cd)

site: $(SITE)/$(TITLE).html $(IMAGES)/image-20180818145606919.png $(IMAGES)/jalopnik.png $(IMAGES)/cpu-performance.png $(IMAGES)/int_graph.png $(IMAGES)/stuttering.png $(IMAGES)/mandelbrot.png $(IMAGES)/cmos-inverter.png $(IMAGES)/power-density.png $(IMAGES)/gate-length.png $(IMAGES)/highrescpudies_fullyc_020-1105.png $(IMAGES)/AmdahlsLaw.svg
	rsync -az -e "ssh -o StrictHostKeyChecking=no -o ControlMaster=auto -o 'ControlPath=~/.ssh/cm_socket/%r@%h:%p' -o ControlPersist=yes" $(SITE) dave.cheney.net:/export/sites/cheney.net/dave/htdocs/

pdf: $(OUTPUT)/$(TITLE).pdf

epub: $(OUTPUT)/$(TITLE).epub

$(SITE)/%.html: $(LANG)/%.asciidoc $(CURDIR)/Makefile
	asciidoctor -b html5 \
		--failure-level=WARN \
		-d article \
		-a revnumber=$(GIT_VERSION) \
		-o $@ \
		$<

$(OUTPUT)/%.pdf: $(LANG)/%.asciidoc $(wildcard $(COMMON)/*.asciidoc) $(CURDIR)/Makefile | $(OUTPUT)
	docker run -it -v $(CURDIR):/documents -v $(OUTPUT):$(OUTPUT) asciidoctor/docker-asciidoctor \
		asciidoctor-pdf \
		-a revnumber=$(GIT_VERSION) \
		-o $@ \
		$<

$(OUTPUT)/%.epub: $(LANG)/%.asciidoc $(wildcard $(COMMON)/*.asciidoc) $(CURDIR)/Makefile | $(OUTPUT)
	docker run -it -v $(CURDIR):/documents -v $(OUTPUT):$(OUTPUT) asciidoctor/docker-asciidoctor \
		asciidoctor-epub3 \
		-a revnumber=$(GIT_VERSION) \
		-o $@ \
		$<

$(SITE)/%.css: $(COMMON)/%.css | $(SITE)
	install $< $@
	
$(IMAGES)/%.png: images/%.png | $(IMAGES)
	install $< $@

$(IMAGES)/%.svg: images/%.svg | $(IMAGES)
	install $< $@

fmt: $(SRCFILES)
	gofmt -w -s $^

build: fmt | $(SRCFILES:.go=.o)

%.o: %.go
	go tool compile -o $@ $< 

$(DIRS):
	install -d $@

clean:
	rm -rf $(OUTPUT) $(SRCFILES:.go=.o)

.PHONY: rsync epub pdf
