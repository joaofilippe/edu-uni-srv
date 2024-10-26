package enums

type Disability int

const (
    Visual Disability = iota
    Auditory
    Motor
    Cognitive
    Speech
)

func (d Disability) String() string {
    return [...]string{"Visual", "Auditory", "Motor", "Cognitive", "Speech"}[d]
}

func ParseDisability(s string) Disability {
    disabilities := map[string]Disability{
        "Visual":    Visual,
        "Auditory":  Auditory,
        "Motor":     Motor,
        "Cognitive": Cognitive,
        "Speech":    Speech,
    }

    if val, ok := disabilities[s]; ok {
        return val
    }
    return -1
}