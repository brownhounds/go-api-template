#!/bin/bash

log_message_success() {
    local text="$1"
    echo -e "\e[32m$text\e[0m"
}

log_message_info() {
    local text="$1"
    echo -e "\e[34m$text\e[0m"
}
