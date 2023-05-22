# GIT Tips

## Work with pre-commit

### What to do for 'Rolling back fixes'

In case you get:

> [WARNING] Stashed changes conflicted with hook auto-fixes... Rolling back fixes...

follow those steps:

```sh
# Add files we want to commit
git add some_file.ext some_dir/
# Stash the rest
git stash --keep-index --include-untracked --message "a message useful in git stash list"
# Commit changes
git commit --message="the commit message"
# Check that status is clean
git status
# Bring back stashed files
git stash pop
```

Example:

```console
~$ # Check current status
~$ git status
On branch virtualbox-plugin
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   README.md
        modified:   my_file.ext

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        new_file1.ext
        new_file2.ext
        new_folder/

no changes added to commit (use "git add" and/or "git commit -a")
~$ # Add files we want to commit
~$ git add my_file.ext new_file1.ext
~$ # Stash all other files
~$ git stash --keep-index --include-untracked --message <message>
Saved working directory and index state WIP on <<your-branch>>: <<message>>
~$ # Check we still have the staged files we want to commit
~$ git status
On branch virtualbox-plugin
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   my_file.ext
        new file:   new_file1.ext
~$ # Commit files
~$ git commit --message=<msg>
~$ # Check that status is clean
~$ git status
On branch virtualbox-plugin
nothing to commit, working tree clean
~$ # Bring back stashed files
~$ git stash pop
On branch virtualbox-plugin
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   README.md

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        new_file2.ext
        new_folder/

no changes added to commit (use "git add" and/or "git commit -a")
Dropped refs/stash@{0} (<<hash>>)
```

### How to edit commits (for example, split a commit)

Rebase to a commit that contain the commits you want to edit.
For example:

- if it's 3 commits back, use `HEAD~3`
- if you are on a new branch and want to rebase to 'main' use `main`

    > Example

    ```sh
    git rebase -i main
    ```

Mark as **edit** the commits you want to change (replace '**pick**' with '**edit**'
for the commits), and then modify each commit.

```sh
# For each commit marked as 'edit'
# (reset) to unstage chenges so we can work on them
git reset HEAD~
# (change) files and then commit
# - make the changes you want
# - add the needed files
# - stash the other files
# - commit the changes
# - bring back stashed files
git add one-file-I-changed.txt another-file-I-want-in-this-commit.txt
# stash files you want to keep for later (or skip if you added all the files)
git stash --keep-index --include-untracked
# commit the changes
git commit -m "a useful message"
# bring back stashed files (skip also this if you skipped the previous
# 'stash' command)
git stash pop
# Repeat from (change) until your working tree is clean and
# all changes are done for the commit.
# Tell git to continue to the next rebase step
git rebase --continue
# Repeat from (reset) for each commit you want to change.
```
