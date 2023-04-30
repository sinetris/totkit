# GIT Tips

## Work with pre-commit

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
Saved working directory and index state WIP on <<your-branch>>: <<hash>> <<message>>
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
