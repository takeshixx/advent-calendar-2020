import sys
import random
import os
import lief

arglen = len(sys.argv)

if arglen != 2:
  print("Usage: " + sys.argv[0] + " <string>")

strarg = sys.argv[1]
strlen = len(strarg)

# Obfuscate string a bit such that it is not directly found via strings command
with open("./elf.c","w") as f:
  f.write("#include <stdio.h>\n\nint main() {\n")
  for i in range(0,strlen):
    rnd = random.randint(0,255)
    obfchar = ord(strarg[i]) ^ rnd
    f.write('printf("%c", ' + str(obfchar) + ' ^ ' + str(rnd) + ');'+"\n")
  f.write('}')

# Create binary
os.system('gcc ./elf.c -o ./elf.orig')

# Modify entry point
binary = lief.parse('./elf.orig')
header = binary.header
orig_entrypoint = header.entrypoint

header.entrypoint = 0xdeadbeef

binary.write('/usr/local/apache2/htdocs/elf')

content = """
<html>
<style>
@font-face {
  font-family: "xmas";
  src: url("./BeyondWonderland.ttf");
}

.content {
  font-family: xmas;
  font-size: 30pt;
  max-width: 500px;
  margin: auto;
}
</style>
<body>

<div class="content">

<img src='.\elf.png' height='50%'>

<br><br>

Santa's little <a href='./elf'>elf</a> has so much to do and doesn't know where to start. He tried to clean up the dead beef that somehow piles up. Santa told him to replace it with
""" + hex(orig_entrypoint) + """, but he doesn't know what that means...

</div>

</body></html>
"""

with open("/usr/local/apache2/htdocs/index.html","w") as f:
  f.write(content)
