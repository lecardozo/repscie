package helper

import (
    "bytes"
)

func Cat (args ...string) string{
    var buffer bytes.Buffer

    for _, arg := range args {
        buffer.WriteString(arg)
    }

    return buffer.String()
}
