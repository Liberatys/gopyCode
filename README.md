# gopyCode

An easy to use code summary tool, that collects all source code for a given file ending in a directory.
All source code is written into a single file you can copy your code from.
# usage

gopyCode outputFileName filending

gopyCode javaSummary.txt .java


it is also possible to add multiple endings
gopyCode codeSummary.txt .java .go


At the beginning of the collected file, there will be a table with all files located and displayed by their extension.

    ".go": [
        "gopyCode/files.go",
        "gopyCode/main.go",
        "gopyCode/writer.go"
    ]

# Future

I hacked this for a project, where I had to send all my source code in a single file. So I just wrote this little helper and it was done in about 5s with all files.

Maybe, I will expand on this idea in the future.


# Thanks