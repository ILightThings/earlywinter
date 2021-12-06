A small golang binary made to search for files on a (potentially compromosed) host. Useful for a quick priv esc via hard coded credients.
Will print out file paths of discovered files.


```
Usage of C:\dev\earlywinter.exe:
  -ext string
        Extentions to search for (default ".ps1,.bat,.sh,.py")
  -h    Displays this help file
  -path string
        Path to check. Network example("\\client.local\sharename") (default "c:")
```


If the binary is searching locally, will ignore common system folders:
```
C:\ProgramData
C:\Program Files (x86)
C:\Program Files
C:\$WinREAgent
C:\$Windows.~WS
C:\$WINDOWS.~BT
C:\Windows"
C:\Windows10Upgrade
C:\Documents and Settings
C:\Python27
```



Check a network share. Requires a network share (`\\server.local\share`) directly. Server name alone (`\\server.local`) will not work. 
```
earlywinter.exe -path "\\client.local\sharename"
```

Change the files to search for. Format is `.ext` with comma separation.
```
earlywinter.exe -ext ".jpeg,.config,.html"
```

