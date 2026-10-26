#!/usr/bin/env python3
"""Verify the oral deck, its timing contract, layout, and running LP example."""
from itertools import combinations
from pathlib import Path
import re
import subprocess
import tempfile
import xml.etree.ElementTree as ET

import numpy as np

ROOT = Path(__file__).resolve().parent


def run(*args):
    return subprocess.check_output(args, text=True)


def vertices(A, b):
    """Enumerate feasible vertices of this pointed 2-D example, not a general solver."""
    for pair in combinations(range(len(b)), 2):
        M = A[list(pair)]
        if abs(np.linalg.det(M)) < 1e-10:
            continue
        q = np.linalg.solve(M, b[list(pair)])
        if np.all(A @ q <= b + 1e-8):
            yield q


def expected_value(d):
    return max(0.0, d, 2.0 * d - 1.0)


def main():
    source = (ROOT / "qualification.tex").read_text()
    frames = re.findall(r"\\begin\{frame\}", source)
    assert len(frames) == 27, f"Expected 27 frames, found {len(frames)}"
    info = run("pdfinfo", str(ROOT / "qualification.pdf"))
    pages = int(re.search(r"Pages:\s+(\d+)", info).group(1))
    assert pages == 27, f"Expected 27 PDF pages, found {pages}"
    assert not re.search(r"\\(?:pause|only|uncover|onslide)\b", source), "Unexpected overlays"
    print("PASS: 27 frames, 27 PDF pages, no overlay commands")

    timing = re.findall(r"% (\d{2}) \| (\d+) s", source)
    assert [int(n) for n, _ in timing] == list(range(1, 28))
    assert sum(int(t) for _, t in timing) == 24 * 60
    guide = (ROOT / "speaking-guide.md").read_text()
    rows = re.findall(
        r"^\| (\d+) \| [^\n|]+ \| (\d+):(\d+) \| (\d+):(\d+) \|$",
        guide, re.M
    )
    assert len(rows) == 27, f"Expected 27 guide timing rows, got {len(rows)}"
    total = 0
    for (n, sec), (gn, m, s, cm, cs) in zip(timing, rows):
        assert int(n) == int(gn)
        assert int(sec) == 60 * int(m) + int(s)
        total += int(sec)
        assert total == 60 * int(cm) + int(cs)
    print("PASS: source and speaking-guide timings agree: 24:00 + 1:00 buffer")

    log = (ROOT / "qualification.log").read_text()
    problems = re.findall(
        r"(?:Overfull \\[hv]box|Missing character|LaTeX Warning|Package \S+ Warning)[^\n]*",
        log
    )
    assert not problems, "\n".join(problems)
    print("PASS: no overfull boxes, missing characters, or LaTeX/package warnings")

    with tempfile.TemporaryDirectory(prefix="glp-deck-check-") as temp:
        bbox = Path(temp) / "bbox.html"
        subprocess.run(
            ["pdftotext", "-bbox", str(ROOT / "qualification.pdf"), str(bbox)],
            check=True
        )
        # Poppler can emit XML-forbidden control bytes for legacy math glyphs.
        xml = re.sub(r"[\x00-\x08\x0b\x0c\x0e-\x1f]", "", bbox.read_text())
        tree = ET.fromstring(xml)
        ns = {"x": "http://www.w3.org/1999/xhtml"}
        rendered_pages = tree.findall(".//x:page", ns)
        assert len(rendered_pages) == 27
        for n, page in enumerate(rendered_pages, 1):
            w, h = float(page.attrib["width"]), float(page.attrib["height"])
            words = page.findall(".//x:word", ns)
            assert words, f"No text on page {n}"
            for word in words:
                x0, x1 = float(word.attrib["xMin"]), float(word.attrib["xMax"])
                y0, y1 = float(word.attrib["yMin"]), float(word.attrib["yMax"])
                assert x0 >= -0.5 and x1 <= w + 0.5, (n, word.text, x0, x1)
                assert y0 >= -0.5 and y1 <= h + 0.5, (n, word.text, y0, y1)
    print("PASS: extracted text lies inside all 27 PDF page boundaries")

    # Original constraints: x1 >= 0, x1 <= 1, x2 >= 0, x1+x2 >= d.
    A = np.array([[-1., 0.], [1., 0.], [0., -1.], [-1., -1.]])
    c = np.array([1., 2.])
    demands = np.unique(np.r_[np.linspace(-2, 4, 121), 0, 1])
    count = 0
    for d in demands:
        b = np.array([0., 1., 0., -d])
        candidates = list(vertices(A, b))
        assert candidates, f"No feasible vertices for d={d}"
        computed = min(float(c @ q) for q in candidates)
        expected = expected_value(d)
        assert abs(computed - expected) < 1e-8, (d, computed, expected)
        # Verify the projected epigraph against feasibility of the original
        # constraints augmented by c^T x <= t.
        for offset in [-1.0, -0.01, 0.0, 0.01, 1.0]:
            t = expected + offset
            augmented_A = np.vstack([A, c])
            augmented_b = np.r_[b, t]
            feasible = bool(list(vertices(augmented_A, augmented_b)))
            halfspaces = t >= max(0., d, 2 * d - 1) - 1e-8
            assert feasible == halfspaces, (d, t, feasible, halfspaces)
            count += 1
    print(f"PASS: running LP values for {len(demands)} demands; {count} epigraph checks")

    # Intermediate GPA projection: s=x1+x2 and c=2s-x1 <= t.
    projection_checks = 0
    for s in np.linspace(-2, 4, 61):
        for t in np.linspace(-3, 9, 121):
            interval_feasible = max(0., 2 * s - t) <= min(1., s) + 1e-8
            projected = s >= -1e-8 and t >= s - 1e-8 and t >= 2 * s - 1 - 1e-8
            assert interval_feasible == projected, (s, t)
            # Generator diagram: t>=0, t>=d, (t+1)/2>=d.
            for d in [-2., 0., 0.5, 1., 3.]:
                diagram = t >= -1e-8 and t >= d - 1e-8 and (t + 1) / 2 >= d - 1e-8
                assert diagram == (t >= expected_value(d) - 1e-8), (d, t)
            projection_checks += 1
    print(f"PASS: {projection_checks} intermediate projections and reduced-generator checks")
    print("NOTE: these checks validate the example and deck, not the general completeness theorem.")


if __name__ == "__main__":
    main()
