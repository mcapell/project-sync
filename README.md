# Git repository sync

Check if local git repositories are in sync with upstream.

## How to use it

Create a configuration file on `~/.config/project_sync.yaml` with the following structure:

	projects:
		- name: VimRC
		  repo: github.com/mcapell/vimrc
		  path: ~/.vim
