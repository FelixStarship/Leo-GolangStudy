package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("helm", "template", "deploy-test", "/home/gzsl/kruise-2023.8.2.tgz", "-n", "hxz")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("获取标准输出管道失败:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("启动命令失败:", err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	var yamlDoc strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "---") {
			// 解析前一个YAML文档
			if err := parseYAML(yamlDoc.String()); err != nil {
				fmt.Println("解析错误:", err)
			}
			yamlDoc.Reset()
		}
		yamlDoc.WriteString(line)
		yamlDoc.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取命令输出失败:", err)
	}

	// 解析最后一个YAML文档
	if err := parseYAML(yamlDoc.String()); err != nil {
		fmt.Println("解析错误:", err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("等待命令完成失败:", err)
	}
}

func parseYAML(yamlStr string) error {
	// 使用map解析YAML
	var data map[string]interface{}
	if err := yaml.Unmarshal([]byte(yamlStr), &data); err != nil {
		return err
	}

	kind := data["kind"]
	metadata, ok := data["metadata"].(map[interface{}]interface{})
	name := ""
	if ok {
		nameValue, ok := metadata["name"]
		if ok {
			name = nameValue.(string)
		}
	}

	fmt.Println("kind:", kind)
	fmt.Println("metadata.name:", name)

	return nil
}
