package util

import (
    "regexp"
    "strings"
)

// ValidateCPS проверяет синтаксис CPS-сигнатуры
func ValidateCPS(cps string) error {
    if cps == "" {
        return nil
    }
    
    // Паттерны для тегов
    patterns := []string{
        `<b\s+0x[0-9a-fA-F]+>`,           // static bytes
        `<t>`,                             // timestamp
        `<r\s+\d+>`,                       // random bytes
        `<rc\s+\d+>`,                      // random chars
        `<rd\s+\d+>`,                      // random digits
    }
    
    combined := "(" + strings.Join(patterns, ")|(") + ")"
    re := regexp.MustCompile(combined)
    
    // Проверяем, что вся строка состоит из валидных тегов
    remaining := cps
    for re.MatchString(remaining) {
        remaining = re.ReplaceAllString(remaining, "")
    }
    
    // Удаляем пробелы для финальной проверки
    remaining = strings.TrimSpace(remaining)
    if remaining != "" {
        return fmt.Errorf("invalid CPS syntax at: %s", remaining)
    }
    
    return nil
}

// ValidateHeaderRange проверяет диапазон заголовков
func ValidateHeaderRange(min, max uint32) error {
    if min > max {
        return fmt.Errorf("min (%d) cannot be greater than max (%d)", min, max)
    }
    return nil
}
