<p align="center">
    <img src="GopyKitty.png" height="400" width="400">
</p>

# GopyCode

An easy to use code summary tool, that collects all source code for a given file ending in a directory.
All source code is written into a single file you can copy your code from. The tool also provides an overview of all collected files boundled by the extension they have.

At the moment it is just a basic implementation, but in the future, I can see a good adaption of such a programm. I added flags, in order to make the configuration easier and cleaner.

# Usage

    [-o --o -out]    output
            Define the ouput file for gopyCode like
                -> -o output.txt

    [-s --s -src]        start
            Set the starting folder for gopyCode like
                -> -src .

    [-ex -e --e] extensions
            Define the extennsions, gopyCode should look for like
                -> -ex .java .go


Example:

    gopyCode -o output.txt -s . -ex .go .java

This will generate an output.txt file, that cointains all file contents for .go and .java.

At the head of an output file, the application will place short overview over files it found and indexed.
This looks something like this:

    ".go": [
        "gopyCode/files.go",
        "gopyCode/flags.go",
        "gopyCode/main.go",
        "gopyCode/writer.go"
    ],
    ".java": []

# Future

I hacked this for a project, where I had to send all my source code in a single file. So I just wrote this little helper and it was done in about 5s with all files.
