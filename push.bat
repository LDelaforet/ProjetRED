@echo off

:a
if "%1" == "" (
    set /p commitMessage="Description du commit : "
) else (
    set commitMessage=%1
)

if "%commitMessage%" == "" (
    echo Non mais faut mettre une description en fait
    goto a
)

git add .

git commit -m "%commitMessage%"

git push