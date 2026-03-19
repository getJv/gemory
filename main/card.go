components {
  id: "card_script"
  component: "/main/scripts/card.script"
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
  id: "back"
  type: "sprite"
  data: "tile_set: \"/assets/game.atlas\"\ndefault_animation: \"card_back\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ALPHA\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "front"
  type: "sprite"
  data: "tile_set: \"/assets/game.atlas\"\ndefault_animation: \"card_front\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ALPHA\n"
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
  id: "shape"
  type: "sprite"
  data: "tile_set: \"/assets/game.atlas\"\ndefault_animation: \"shape1\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ALPHA\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.05
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
