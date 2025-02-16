#!/bin/bash

# Получаем процент загрузки CPU с помощью mpstat (из пакета sysstat)
cpu_idle=$(mpstat 1 1 | awk '/Average/ {print $NF}')
cpu_usage=$(echo "100 - $cpu_idle" | bc)

# Получаем количество логических процессоров
cpu_cores=$(nproc)

# Выводим информацию
echo "CPU Usage: $cpu_usage%"
echo "CPU Cores: $cpu_cores"
