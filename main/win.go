components {
  id: "win_gui"
  component: "/main/win.gui"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "victory_sound"
  type: "sound"
  data: "sound: \"/assets/sounds/victory.wav\"\nlooping: false\ngroup: \"master\"\ngain: 1.0\npan: 0.0\nspeed: 1.0\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
