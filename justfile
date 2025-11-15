# List "Just" commands
default:
  just --list

test:
  gotestsum --format testdox --format-icons hivis --watch --watch-clear
