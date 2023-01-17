# bash oneliner to get repo list from gh utility
`gh repo list --limit 1000 | awk '{print $1}'`

# gh oneliner to fetch 'quiz' repo url to use it as a name for my go module

`gh repo url quiz | awk '{print $2}'
gh repo view quiz | grep -i 'clone url' | awk '{print $3}'`


# git commad to see changes made before commit

`git diff --cached`


# git command to see current user and email
`git config --global user.name && git config --global user.email
`
# gh command to read wiki related to 'quiz' repo
gh repo view quiz | grep -i 'wiki' | awk '{print $2}'
```

# gh command to manage github wiki 
gh wiki clone quiz

# bash command to save this text file to a specific folder