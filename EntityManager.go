package entity

type EntityManager struct {
    entities []Entity
    componentData map[string] map[Entity] []Component
}

func (em *EntityManager) AddComponent(entity Entity, component Component) {
    componentID := component.GetID()
    if em.componentData[componentID] == nil {
        em.componentData[componentID] = component
    }
}

func (*EntityManager) Query([]string) {
    
}
