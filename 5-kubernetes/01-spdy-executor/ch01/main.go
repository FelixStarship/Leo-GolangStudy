//package main
//
//import (
//	"flag"
//	"fmt"
//	corev1 "k8s.io/api/core/v1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/client-go/kubernetes"
//	"k8s.io/client-go/tools/clientcmd"
//	"k8s.io/client-go/tools/remotecommand"
//	"log"
//)
//
//func main() {
//	// 获取 kubeconfig 文件路径
//	kubeconfig := flag.String("kubeconfig", "", "Path to the kubeconfig file")
//	flag.Parse()
//
//	// 使用 kubeconfig 创建 Kubernetes 客户端
//	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	if err != nil {
//		log.Fatal(err)
//	}
//	clientset, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 指定要进入的 Pod 和容器
//	podName := "appship-test-application-center-6558bdf6cc-x6bbg"
//	containerName := "application-center"
//
//	// 创建执行命令的请求
//	req := clientset.CoreV1().RESTClient().Post().
//		Resource("pods").
//		Name(podName).
//		Namespace("appship-test").
//		SubResource("exec")
//
//	// 设置执行命令的参数
//	req.VersionedParams(&corev1.PodExecOptions{
//		Container: containerName,
//		Command:   []string{"ls"},
//		Stdin:     false,
//		Stdout:    true,
//		Stderr:    true,
//	}, metav1.ParameterCodec{})
//
//	// 执行命令并获取响应
//	executor, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
//	if err != nil {
//		log.Fatal(err)
//	}
//	var stdout, stderr []byte
//	err = executor.Stream(remotecommand.StreamOptions{
//		Stdout: &stdout,
//		Stderr: &stderr,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 输出命令执行结果
//	fmt.Printf("Command output:\n%s\n", string(stdout))
//	fmt.Printf("Command error:\n%s\n", string(stderr))
//}
