# Load antigen
source $HOME/.shell/antigen.zsh

# Environment
export ZSH_AUTOSUGGEST_STRATEGY=(history completion)
export PATH="$PATH:$HOME/.local/bin"
export MANGOHUD=1
export TERM=xterm
export STEAMLIBRARY="/mnt/storage/Games/Steam/"
export PROTON="/mnt/storage/Games/Steam/steamapps/common/Proton - Experimental/files/"

# OhMyPosh theme
THEME_DIR=$HOME/.cache/oh-my-posh/themes
OMP_THEME="powerlevel10k_classic"

# Load OhMyPosh and other tools
eval "$(oh-my-posh init zsh --config $THEME_DIR/$OMP_THEME.omp.json)"
eval "$(mcfly init zsh)"
eval "$(zoxide init zsh)"

# Command history config
HISTFILE=~/.zsh_history
HISTSIZE=10000
SAVEHIST=10000
setopt appendhistory

# Antigen shell setup
antigen bundle zsh-users/zsh-completions
antigen bundle zsh-users/zsh-autosuggestions
antigen bundle zsh-users/zsh-syntax-highlighting
antigen apply

# Alias all the things
alias ls="pls -a -d perms -d user -d group -d size -d mtime -d git"
alias cp="cpg -g"
alias mv="mvg -g"
alias cat="bat"
alias codium="flatpak run com.vscodium.codium "

# Key bindings
bindkey '^[[H' beginning-of-line
bindkey '^[[F' end-of-line
bindkey '^[[3~' delete-char

# Sexy fetch
fastfetch