/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package mongo

func CollectionUserTokenName() string {
	return "user.token"
}

func CollectionUserPolicyName() string {
	return "user.policy"
}

func CollectionUserRSAName(service string) string {
	return "user.rsa." + service
}

func CollectionCloudTokenName(cloud string) string {
	return "cloud.token." + cloud
}

func CollectionTemplateName(id string) string {
	return "template." + id
}

func CollectionAssetCloudName(cloud, id string) string {
	return "asset." + id + "." + cloud
}

func CollectionHookName() string {
	return "hook"
}

func CollectionHookReceiveName(service string) string {
	return "hook.receive." + service
}

func CollectionOrchestrationName() string {
	return "orchestration"
}
