LATEXMK ?= latexmk

.PHONY: all compile check clean
all: qualification.pdf

compile: all

qualification.pdf: qualification.tex theme.tex ../material/generator.tikz ../material/generator.tikzstyles ../tikz/AxiomsGLP.tex ../tikz/NormalFormGLP.tex
	$(LATEXMK) -lualatex -interaction=nonstopmode -halt-on-error qualification.tex

check: all
	python3 check.py

clean:
	$(LATEXMK) -c qualification.tex
