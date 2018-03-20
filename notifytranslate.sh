#!/bin/bash
SELECTION=$(xsel -o)
notify-send --icon=info "Translate: $SELECTION" "$(translate ${SELECTION})"