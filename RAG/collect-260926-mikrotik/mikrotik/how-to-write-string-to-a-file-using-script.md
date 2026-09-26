---
id: collect-260926-mikrotik/mikrotik/how-to-write-string-to-a-file-using-script
title: "Create a file with specific content."
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/how-to-write-string-to-a-file-using-script.md
source_anchor: ""
source_lines: [1, 75]
sha256: 69cad3643e07b1fb8050c6c439012e2e711521835b9b21d61074e8e56b612197
---

# Create a file with specific content.

How to write string to a file using script ? Please provide a command or function, THX

             
                
            
           
          
            
            
              ```
# Create a file with specific content.
{
# Variables.
    :local filename "test.txt"
    :local content "This is a test."
    
# Create file.
    /file print file=$filename
# Set file's content.
    /file set $filename contents=$content
}
```

Be aware that there is a maximum file size that RouterOS will allow you to access using scripts.

I forget the size and method though, but at some point **[/file get $filename contents]** just starts either giving blank results or throw errors.

             
            
           
          
            
            
              your script is not working

             
            
           
          
            
            
              
Your comment is not helping.

Jokes aside. Helps if you tell me how it isn’t working or what is going wrong.

You might need a delay between creating the file and changing the content.

```
# Create a file with specific content.
{
# Variables.
    :local filename "test.txt"
    :local content "This is a test."
    
# Create file.
    /file print file=$filename where name=""
# Wait for the file to be created.
    :delay 1
# Set file's content.
    /file set $filename contents=$content
}
```

Works for me, but there are times when this will stop working and if it takes too long to create the file it won’t be able to write.

             
            
           
          
            
            
              also note, due to mikrotik’s implementation, variable value sizes can be up to 4096 bytes.

this ultimately limits the file size you create this way.
