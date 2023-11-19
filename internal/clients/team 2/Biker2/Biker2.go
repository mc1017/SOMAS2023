package biker2 

import (
    "github.com/google/uuid"
    utils "SOMAS2023/internal/common/utils"
    objects "SOMAS2023/internal/common/objects"
    baseAgent "github.com/MattSScott/basePlatformSOMAS/BaseAgent"
)

type Biker2 struct {
    *objects.BaseBiker
}


func (cb *Biker2) DecideAction() BikerAction {
    
    return Pedal // or ChangeBike
}

func (cb *Biker2) DecideForce() {
    // Custom logic to decide the force
}

func (cb *Biker2) GetForces() utils.Forces {
    // Return the current forces
    return cb.forces
}