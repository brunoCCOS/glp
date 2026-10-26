# Master's qualification presentation

An English, 16:9 **27-slide** presentation designed to be spoken in **25 minutes**, based on the active dissertation source.

## Files

| File | Purpose |
|---|---|
| [qualification.pdf](qualification.pdf) | Project this deck. Exactly 27 pages, without overlays or backup slides. |
| [qualification.tex](qualification.tex) | Editable Beamer content. |
| [theme.tex](theme.tex) | Classic Madrid Beamer theme (short running header) and diagram styles. |
| [speaking-guide.md](speaking-guide.md) | Content selection, slide timings, spoken cues, transitions, and recovery cuts. |
| [qualification-notes.md](qualification-notes.md) | Source mapping, mathematical qualifications, likely committee questions, and proposed next work. |
| [Makefile](Makefile) | Reproducible build/check commands. |
| [check.py](check.py) | Page/frame/timing/layout checks and independent running-example verification. |

The talk gives the main time to the running LP, composition, the GLP generator, and epigraph-based completeness. Quantales and the inherited languages are context, not a chapter-by-chapter survey.

The source assumes it remains in this repository: it reuses the shared generator definitions and two existing TikZ figures. No dissertation source has been changed.

## Build

Required: LuaLaTeX, latexmk, Beamer, PGF/TikZ, and pgfplots. Checks additionally use Python 3, NumPy, pdfinfo, and pdftotext.

Run from the repository root:

~~~sh
make -C presentation compile
make -C presentation check
~~~

Or from this directory:

~~~sh
make compile
make check
~~~

## Use

1. Read the topic-selection rationale in the speaking guide.
2. Rehearse using the cumulative time targets.
3. Resolve the semantic qualifications in the private notes with the advisor.
4. Project only qualification.pdf.

The planned speech is 24 minutes, leaving one minute for pauses. Q&A is additional. Author, advisor, and year come from the existing dissertation metadata; an exact qualification date has not been invented.
