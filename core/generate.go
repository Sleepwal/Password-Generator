package core

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Generator 结构用于生成密码
type Generator struct {
	CharSets map[string]string
}

// NewGenerator 创建新的密码生成器实例
func NewGenerator(charSets map[string]string) *Generator {
	return &Generator{
		CharSets: charSets,
	}
}

// Generate 生成指定长度和模式的密码
func (g *Generator) Generate(length int, mode string) (string, error) {
	charSet, ok := g.CharSets[mode]
	if !ok {
		return "", fmt.Errorf("无效的密码模式：%s", mode)
	}

	password := make([]byte, length)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", err
		}
		password[i] = charSet[num.Int64()]
	}
	return string(password), nil
}
