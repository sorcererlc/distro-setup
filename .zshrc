source $HOME/.shell/antigen.zsh

THEME_DIR=$HOME/.cache/oh-my-posh/themes
OMP_THEME="powerlevel10k_classic"
export ZSH_AUTOSUGGEST_STRATEGY=(history completion)

export PATH="$PATH:$HOME/.local/bin"

eval "$(oh-my-posh init zsh --config $THEME_DIR/$OMP_THEME.omp.json)"
eval "$(mcfly init zsh)"

alias ls="pls -a -d perms -d user -d group -d size -d mtime -d git"
alias cp="cpg -g"
alias mv="mvg -g"
alias cat="bat"
alias codium="flatpak run com.vscodium.codium "

antigen bundle zsh-users/zsh-completions
antigen bundle zsh-users/zsh-autosuggestions
antigen bundle zsh-users/zsh-syntax-highlighting
antigen apply

HISTFILE=~/.zsh_history
HISTSIZE=10000
SAVEHIST=10000
setopt appendhistory

# Key Bindings
bindkey '^[[H' beginning-of-line
bindkey '^[[F' end-of-line
bindkey '^[[3~' delete-char