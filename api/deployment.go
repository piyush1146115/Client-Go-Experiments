package api

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateDeployment() {
	clientset := CreateClientSet()
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "go-client-api-server",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "go-rest-api",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "go-rest-api",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "go-rest-api",
							Image: "nginx",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("Creating deploymenr...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Created deployment %q\n", result.GetObjectMeta().GetName())
}

func GetDeployment() {
	fmt.Println("Listing all deployment objects ...")
	clientset := CreateClientSet()
	deploymentClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	list, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range list.Items {
		fmt.Printf("%s (%d replicas)\n", item.Name, *item.Spec.Replicas)
	}
}

func DeleteDeployment() {
	fmt.Printf("Deleting a deployment object .... %s \n", deploymentName)
	clientSet := CreateClientSet()
	deploymentClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
	err := deploymentClient.Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s is deleted successfully.\n", deploymentName)
}

func UpdateDeployment() {
	fmt.Printf("Updating deployment: Scaling the %s deployment to %d replicas \n", deploymentName, replicas)
	clientSet := CreateClientSet()
	deploymentClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
	result, getErr := deploymentClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if getErr != nil {
		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
	}
	result.Spec.Replicas = int32Ptr(replicas)
	_, updateErr := deploymentClient.Update(context.TODO(), result, metav1.UpdateOptions{})

	if updateErr != nil {
		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", updateErr))
	}
	fmt.Println("Updated deployment...")
}
