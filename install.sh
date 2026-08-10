#!/bin/sh
set -e

BINARY=./sca
MANPAGE=./sca.1.gz
PREFIX=${PREFIX:-/usr/local}
BINDIR=${DESTDIR}${PREFIX}/bin
MANDIR=${DESTDIR}${PREFIX}/share/man/man1

install -Dm755 "$BINARY" "$BINDIR/sca"

install -Dm644 "$MANPAGE" "$MANDIR/sca.1.gz"

if command -v mandb >/dev/null; then
    mandb -q || true
fi
