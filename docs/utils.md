### bash oneliner to get repo list from gh utility
`gh repo list --limit 1000 | awk '{print $1}'`
___
### gh oneliner to fetch url for 'quiz' repo to use it as a name for my go module

`gh repo view quiz ---json url --jq '.url' | sed 's/https:\/\/github.com\///g'`
___
### git commad to see changes made before commit

`git diff --cached`
___

### git command to see current user and email
`git config --global user.name && git config --global user.email
`
___