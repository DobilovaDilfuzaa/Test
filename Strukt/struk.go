

package main

import "fmt"

type Equipment struct {
    On    bool
    Ammo  int
    Power int
}

func (a *Equipment) Shoot() bool {
    if a.On && a.Ammo > 0 {
        a.Ammo--
        return true
    }
    return false
}

func (a *Equipment) RideBike() bool {
    if a.On && a.Power > 0 {
        return true
    }
    return false
}

func main() {
    testStruct := &Equipment{On: true, Ammo: 3, Power: 1}
    fmt.Println(testStruct.Shoot())
    fmt.Println(testStruct.RideBike())
}