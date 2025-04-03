#!/bin/bash

: # This is a bash//cmd hybrid script which runs on both Windows & Linux and installs the dependencies for building chrxer.
: # See https://stackoverflow.com/a/17623721 for details.
: # following line is executed in bash
:; set -e;printf "\e[0;31m[Running %s]\033[0m sudo %s\n" "$(date +'%m-%d %T')" "scripts/deps.sh $@"; sudo $(dirname "$0")/deps.sh $@; exit 0
@ECHO off
cls

set "pv=3.11.0"
set "ppath=%TEMP%/pyinstaller_%pv%.exe"
set "log=%TEMP%/pyinstaller_%pv%.log"

:: Extract major (first number)
for /f "tokens=1 delims=." %%a in ("%pv%") do set "major=%%a"
:: Extract minor (second number)
for /f "tokens=2 delims=." %%b in ("%pv%") do set "minor=%%b"

for /f "tokens=*" %%i in ('where "python" 2^>nul') do set "pexc0=%%i"
set "pexc1=%LocalAppData%/Programs/Python/Python%major%%minor%/python.exe"
set "pexc2=%LocalAppData%/Programs/Python/Python%major%%minor%-32/python.exe"

if exist "%pexc0%" goto :python_installed
if exist "%pexc1%" goto :python_installed
if exist "%pexc2%" goto :python_installed

echo "Installing Python %pv% .."
set "purl=https://www.python.org/ftp/python/%pv%/python-%pv%-amd64.exe"
curl.exe --output "%ppath%" --url "%purl%"
%ppath% /passive Include_doc=0 Include_pip=1 Include_tcltk=1 Include_test=1 CompileAll=1 Include_symbols=1 Include_debug=1 PrependPath=1 Shortcuts=1 Include_launcher=1

:python_installed
for /f "tokens=*" %%i in ('python -c "import sys; print(sys.executable)" 2^>nul') do set "pexc=%%i"
for %%i in ("%pexc%") do set "pexc_dir=%%~dpi"
2>NUL mklink "%pexc_dir%python%major%%minor%.exe" "%pexc%"
2>NUL mklink "%pexc_dir%python%major%.exe" "%pexc%"
for /f "tokens=*" %%i in ('python%major% -c "import sys; print(sys.executable)"') do set "pexc=%%i"
echo "Python %pv% installed at %pexc%, exiting"
echo "WARNING: installing deps not implemented completely yet!"
