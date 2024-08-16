function port-kill {
  lsof -i tcp:$1 | awk 'NR!=1 {print $2}' | xargs kill;
  sudo kill $(sudo lsof -i:$1 -t) &> /dev/null;
}

function killport {
  sudo -S fuser -k ${1}/tcp
  sudo -S fuser -k ${1}/udp
}

function drivers {
  # Check video drivers in use
  lspci -n -n -k | grep -A 2 -e VGA -e 3D

}

function gpu-info {
  # Check active GPU driver
  glxinfo | grep -e OpenGL.vendor -e OpenGL.renderer
}

function gpu-list {
  # List available and default GPU
  switcherooctl list
}

function gpuls {
  echo "[GPU Drivers]"
  drivers
  echo "\n"

  echo "[GPU Info]"
  gpu-info
  echo "\n"

  echo "[GPU List]"
  gpu-list
  echo "\n"
}
