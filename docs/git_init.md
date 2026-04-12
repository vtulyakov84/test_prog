## Quick setup — if you’ve done this kind of thing before

### ssh: `git@github.com:vtulyakov84/test_prog.git`

#### …or create a new repository on the command line
```bash
echo "# test_prog" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:vtulyakov84/test_prog.git
git push -u origin main
```

#### …or push an existing repository from the command line
```bash
git remote add origin git@github.com:vtulyakov84/test_prog.git
git branch -M main
git push -u origin main
```


### https: https://github.com/vtulyakov84/test_prog.git

#### …or create a new repository on the command line
```bash
echo "# test_prog" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/vtulyakov84/test_prog.git
git push -u origin main
```
#### …or push an existing repository from the command line
```bash
git remote add origin https://github.com/vtulyakov84/test_prog.git
git branch -M main
git push -u origin main
```