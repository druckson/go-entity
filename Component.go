package entity;

type Component interface {
    Encode()
    Decode() []byte
    GetID() string
}
