components {
  id: "camera"
  component: "/rendercam/camera.script"
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
  properties {
    id: "orthoScale"
    value: "7.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "fixedAspectRatio"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
  properties {
    id: "aspectRatio"
    value: "720.0, 1280.0, 0.0"
    type: PROPERTY_TYPE_VECTOR3
  }
  properties {
    id: "useViewArea"
    value: "false"
    type: PROPERTY_TYPE_BOOLEAN
  }
  properties {
    id: "fixedArea"
    value: "false"
    type: PROPERTY_TYPE_BOOLEAN
  }
  properties {
    id: "fixedWidth"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
}
components {
  id: "camera_movement"
  component: "/assets/camera/camera_movement.script"
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
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"camera\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 310.0\n"
  "  data: 590.0\n"
  "  data: 10.0\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  "bullet: false\n"
  ""
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
